# YAML Example Program

This directory contains the source Pulumi YAML example for the OPNsense provider.

## Purpose

- Acts as the canonical example used for local provider testing.
- Serves as the source input for generated language examples (`go`, `nodejs`, `python`, `dotnet`).

## Run Locally

From the repository root:

```bash
make up
```

To destroy resources and clean up:

```bash
make down
```

## Configuration

The example in `Pulumi.yaml` includes placeholder values for:

- `opnsense:address`
- `opnsense:key`
- `opnsense:secret`

Replace them with valid OPNsense API credentials before running the stack.

## Regenerate Language Examples

Generated examples should not be edited manually. Regenerate them from this YAML source:

```bash
make gen_examples
```
