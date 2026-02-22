# Pulumi OPNsense Provider for Python

Python SDK for the Pulumi OPNsense provider.

## Install

```bash
pip install pulumi pulumi_opnsense
```

## Supported Resources

- `pulumi_opnsense.unbound.HostOverride`
- `pulumi_opnsense.unbound.HostAliasOverride`

## Example

```python
import pulumi
import pulumi_opnsense as opnsense

host_override = opnsense.unbound.HostOverride(
   "example-host",
   enabled=True,
   hostname="srv01",
   domain="example.local",
   rr="A",
   server="10.0.0.10",
   description="Managed by Pulumi",
)

pulumi.export("result", host_override.result)
```

## Configuration

Configure provider credentials in your Pulumi stack:

- `opnsense:address`
- `opnsense:key`
- `opnsense:secret`

## Development

This SDK is generated from the provider schema. Regenerate SDKs from the repository root:

```bash
make codegen
```
