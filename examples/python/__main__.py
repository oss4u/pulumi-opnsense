import pulumi
import pulumi_opnsense as opnsense

my_host_alias_override = opnsense.unbound.HostAliasOverride(
    "myHostAliasOverride",
    description="Pulumi test",
    domain="example.com",
    enabled=True,
    host="host",
    hostname="hostname",
)
pulumi.export("output", {
    "value": my_host_alias_override.result,
})
