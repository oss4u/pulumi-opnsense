# Pulumi OPNsense Provider for Node.js

Node.js SDK for the Pulumi OPNsense provider.

## Install

```bash
npm install @oss4u/opnsense @pulumi/pulumi
```

## Supported Resources

- `opnsense.unbound.HostOverride`
- `opnsense.unbound.HostAliasOverride`

## Example

```typescript
import * as opnsense from "@oss4u/opnsense";

const hostOverride = new opnsense.unbound.HostOverride("example-host", {
	enabled: true,
	hostname: "srv01",
	domain: "example.local",
	rr: "A",
	server: "10.0.0.10",
	description: "Managed by Pulumi",
});

export const result = hostOverride.result;
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
