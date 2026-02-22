# Pulumi OPNsense Provider for .NET

.NET SDK for the Pulumi OPNsense provider.

## Install

Add the NuGet package to your Pulumi .NET project:

```bash
dotnet add package Pulumi.Opnsense
```

## Supported Resources

- `Pulumi.Opnsense.Unbound.HostOverride`
- `Pulumi.Opnsense.Unbound.HostAliasOverride`

## Example

```csharp
using System.Collections.Generic;
using Pulumi;
using Pulumi.Opnsense.Unbound;

return await Deployment.RunAsync(() =>
{
	var hostOverride = new HostOverride("example-host", new HostOverrideArgs
	{
		Enabled = true,
		Hostname = "srv01",
		Domain = "example.local",
		Rr = "A",
		Server = "10.0.0.10",
		Description = "Managed by Pulumi",
	});

	return new Dictionary<string, object?>
	{
		["result"] = hostOverride.Result,
	};
});
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
