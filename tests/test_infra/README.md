# Pulumi test project: Hetzner server + Cloudflare DNS

This Pulumi project creates:

- a Hetzner Cloud server
- a Cloudflare A record that points `fw.test.sys-int.de` to the server IPv4

Before server creation, it resolves the Hetzner image from snapshots using the label selector
`opnsense-<yyyymmdd-hhss>` naming convention. If multiple snapshots match, the
most recent one is used.

## Prerequisites

- Pulumi CLI installed and logged in
- Python 3.11+ and pip
- Hetzner Cloud API token
- Cloudflare API token (with DNS edit permissions for zone `sys-int.de`)

## Setup

```bash
python3 -m venv .venv
source .venv/bin/activate
pip install -r requirements.txt
pulumi stack init dev
pulumi config set --secret hcloudToken <HCLOUD_TOKEN>
pulumi config set --secret cloudflareApiToken <CLOUDFLARE_API_TOKEN>
pulumi config set cloudflareZoneName sys-int.de
```

Optional config:

```bash
pulumi config set fqdn fw.test.sys-int.de
pulumi config set cloudflareZoneName sys-int.de
pulumi config set serverName fw-test-sys-int
pulumi config set serverType cx23
pulumi config set snapshotNamePrefix opnsense-
pulumi config set snapshotIdOverride 360695484
# optional: force a location (e.g. nbg1/hel1); omit to let Hetzner auto-place
pulumi config set location nbg1
pulumi config set --path sshKeyIds[0] my-ssh-key-name-or-id
pulumi config set dnsTtl 120
pulumi config set proxied false
```

If `snapshotIdOverride` is set, it takes precedence over name-based discovery.

## Deploy

```bash
pulumi up
```

## Destroy

```bash
pulumi destroy
```
