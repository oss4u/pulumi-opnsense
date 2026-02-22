package main

import (
	"github.com/oss4u/pulumi-opnsense/sdk/go/oensense/unbound"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi/config"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		cfg := config.New(ctx, "")
		opnsenseAddress := cfg.RequireObject("opnsense:address")
		opnsenseKey := cfg.RequireObject("opnsense:key")
		opnsenseSecret := cfg.RequireObject("opnsense:secret")
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
