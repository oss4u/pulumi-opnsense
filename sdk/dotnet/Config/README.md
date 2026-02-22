# Provider Configuration (.NET)

This namespace contains strongly typed configuration accessors for the OPNsense provider.

Common configuration keys:

- `opnsense:address`
- `opnsense:key`
- `opnsense:secret`

Set values with Pulumi CLI, for example:

```bash
pulumi config set --secret opnsense:address https://opnsense.example.local
pulumi config set --secret opnsense:key <api-key>
pulumi config set --secret opnsense:secret <api-secret>
```
