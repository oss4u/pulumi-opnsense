package main

import (
	"example.com/pulumi-opnsense/sdk/go/opnsense/unbound"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		myHostAliasOverride, err := unbound.NewHostAliasOverride(ctx, "myHostAliasOverride", &unbound.HostAliasOverrideArgs{
			Description: pulumi.String("Pulumi test"),
			Domain:      pulumi.String("example.com"),
			Enabled:     pulumi.Bool(true),
			Host:        pulumi.String("host"),
			Hostname:    pulumi.String("hostname"),
		})
		if err != nil {
			return err
		}
		ctx.Export("output", pulumi.StringMap{
			"value": myHostAliasOverride.Result,
		})
		return nil
	})
}
