import pulumi
import pulumi_hcloud as hcloud
import pulumi_cloudflare as cloudflare
import re
from datetime import datetime
import json
import urllib.parse
import urllib.request

config = pulumi.Config()

hcloud_token = config.require_secret("hcloudToken")
cloudflare_api_token = config.require_secret("cloudflareApiToken")

cloudflare_zone_name = config.get("cloudflareZoneName") or "sys-int.de"
fqdn = config.get("fqdn") or "fw-test.sys-int.de"

if not fqdn.endswith(f".{cloudflare_zone_name}"):
    raise ValueError(f"fqdn '{fqdn}' must end with '.{cloudflare_zone_name}'")

record_name = fqdn[: -(len(cloudflare_zone_name) + 1)]
if not record_name:
    raise ValueError("record name derived from fqdn is empty")

server_name = config.get("serverName") or "fw-test-sys-int"
server_type = config.get("serverType") or "cx23"
snapshot_name_prefix = config.get("snapshotNamePrefix") or "opnsense-"
snapshot_id_override = config.get("snapshotIdOverride")
location = config.get("location")
ssh_key_ids = config.get_object("sshKeyIds") or []
dns_ttl = config.get_int("dnsTtl") if config.get_int("dnsTtl") is not None else 120
proxied = config.get_bool("proxied") if config.get_bool("proxied") is not None else False

hcloud_provider = hcloud.Provider("hcloud", token=hcloud_token)

cloudflare_provider = cloudflare.Provider(
    "cloudflare",
    api_token=cloudflare_api_token,
)

def select_newest_snapshot_id_from_api(token: str) -> str:
    strict_pattern = re.compile(
        rf"{re.escape(snapshot_name_prefix)}(\d{{8}})[-_]?(\d{{4,6}})",
        re.IGNORECASE,
    )
    loose_pattern = re.compile(r"opnsense[-_ ]?(\d{8})[-_ ]?(\d{4,6})", re.IGNORECASE)

    def parse_snapshot_datetime(name: str) -> datetime | None:
        value = name or ""
        match = strict_pattern.search(value) or loose_pattern.search(value)
        if not match:
            return None
        date_part, time_part = match.groups()
        if len(time_part) == 4:
            time_part = f"{time_part}00"
        try:
            return datetime.strptime(f"{date_part}{time_part}", "%Y%m%d%H%M%S")
        except ValueError:
            return None

    candidates: list[tuple[int, datetime, str]] = []
    seen_snapshots: list[str] = []
    page = 1
    while True:
        query = urllib.parse.urlencode({"type": "snapshot", "per_page": 50, "page": page})
        request = urllib.request.Request(
            f"https://api.hetzner.cloud/v1/images?{query}",
            headers={"Authorization": f"Bearer {token}"},
        )
        with urllib.request.urlopen(request) as response:
            payload = json.loads(response.read().decode("utf-8"))

        images = payload.get("images", [])
        for image in images:
            name = str(image.get("name") or "")
            description = str(image.get("description") or "")
            seen_snapshots.append(f"id={image.get('id')} name='{name}' description='{description}'")

            parsed = parse_snapshot_datetime(name) or parse_snapshot_datetime(description)
            if parsed is None:
                continue
            image_id = int(image.get("id") or 0)
            candidates.append((image_id, parsed, name or description))

        next_page = ((payload.get("meta") or {}).get("pagination") or {}).get("next_page")
        if not next_page:
            break
        page = int(next_page)

    if not candidates:
        sample = "; ".join(seen_snapshots[:10])
        raise Exception(
            f"No snapshot name matches prefix '{snapshot_name_prefix}' in Hetzner API results. "
            f"Seen snapshots: {sample}"
        )

    newest_id = max(candidates, key=lambda item: (item[1], item[0]))[0]
    return str(newest_id)


server_image = (
    pulumi.Output.from_input(snapshot_id_override)
    if snapshot_id_override
    else hcloud_token.apply(select_newest_snapshot_id_from_api)
)

zone = cloudflare.get_zone_output(
    filter=cloudflare.GetZoneFilterArgs(
        name=cloudflare_zone_name,
        match="all",
    ),
    opts=pulumi.InvokeOutputOptions(provider=cloudflare_provider),
)

server = hcloud.Server(
    "fwServer",
    name=server_name,
    server_type=server_type,
    image=server_image,
    ssh_keys=ssh_key_ids,
    public_nets=[
        hcloud.ServerPublicNetArgs(
            ipv4_enabled=True,
            ipv6_enabled=True,
        )
    ],
    location=location if location else None,
    opts=pulumi.ResourceOptions(provider=hcloud_provider),
)

dns_record = cloudflare.DnsRecord(
    "fwDnsRecord",
    zone_id=zone.id,
    name=record_name,
    type="A",
    content=server.ipv4_address,
    ttl=dns_ttl,
    proxied=proxied,
    opts=pulumi.ResourceOptions(provider=cloudflare_provider),
)

pulumi.export("serverId", server.id)
pulumi.export("serverIpv4", server.ipv4_address)
pulumi.export("dnsRecordId", dns_record.id)
pulumi.export("boundFqdn", pulumi.Output.concat(record_name, ".", cloudflare_zone_name))
pulumi.export("snapshotNamePrefix", pulumi.Output.from_input(snapshot_name_prefix))
pulumi.export("snapshotImageId", server_image)
pulumi.export("snapshotIdOverride", pulumi.Output.from_input(snapshot_id_override or ""))
