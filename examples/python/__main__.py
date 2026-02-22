import pulumi
import pulumi_opnsense as opnsense

config = pulumi.Config()
opnsense_address = config.require_object("opnsense:address")
opnsense_key = config.require_object("opnsense:key")
opnsense_secret = config.require_object("opnsense:secret")
my_host_alias_override = opnsense.unbound.HostAliasOverride("myHostAliasOverride",
    description="Pulumi test",
    domain="example.com",
    enabled=True,
    host="host",
    hostname="hostname")
pulumi.export("output", {
    "value": my_host_alias_override.result,
})
