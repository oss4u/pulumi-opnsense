# Pulumi OPNsense Native Provider

Pulumi provider for managing OPNsense resources, currently focused on Unbound DNS overrides.

## Scope

Current provider resources:

- `opnsense:unbound:HostOverride`
- `opnsense:unbound:HostAliasOverride`

Provider configuration keys:

- `opnsense:address`
- `opnsense:key`
- `opnsense:secret`

## Repository Layout

- `provider/` - provider implementation and schema source
- `sdk/` - generated SDKs for Go, Node.js, Python, and .NET
- `examples/` - Pulumi example programs (`yaml` is the source for generated examples)
- `deployment-templates/` - release workflow and GoReleaser templates

## Prerequisites

- Go
- Pulumi CLI
- `pulumictl`
- Node.js + Yarn
- Python 3
- .NET SDK
- `jq`

You can use `mise` to install pinned local tools:

```bash
mise install
```

## Build

Build provider + all SDKs:

```bash
make build
```

Build provider binary only:

```bash
make provider
```

Install provider and local SDK links:

```bash
make install
```

## Development

### Keep dependencies tidy

```bash
make tidy
```

### Generate schema and SDKs

```bash
make codegen
```

### Lint and tests

Run lint for provider code:

```bash
make lint
```

Run provider tests:

```bash
make test_provider
```

### Debug build

Build an unoptimized provider binary for local debugging:

```bash
make provider_debug
```

## Examples

The canonical example is `examples/yaml`. Language-specific examples are generated from it.

Generate all language examples:

```bash
make gen_examples
```

Run the YAML example locally:

```bash
make up
```

Destroy example resources:

```bash
make down
```

## CI and Releases

- `develop` workflow builds and validates the provider on pull requests and on schedule.
- `main-tag` workflow creates release tags from commits on `main`.
- `release` workflow builds provider artifacts, generates SDK packages, and publishes release outputs.

Release setup details are documented in [deployment-templates/README-DEPLOYMENT.md](./deployment-templates/README-DEPLOYMENT.md).

## Notes

- `PULUMI_IGNORE_AMBIENT_PLUGINS=true` is set in the `Makefile` to prefer local, pinned plugins.
- Example config values in `examples/yaml` are placeholders and must be replaced with valid OPNsense API credentials.

This SDK is generated from the provider schema. Regenerate SDKs from the repository root:

- [Pulumi Go Provider](https://github.com/pulumi/pulumi-go-provider)
- [Pulumi Provider Development](https://www.pulumi.com/docs/iac/extending-pulumi/)
