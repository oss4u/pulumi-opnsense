package main

import (
	"example.com/pulumi-opnsense/sdk/go/opnsense"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		myRandomResource, err := opnsense.NewRandom(ctx, "myRandomResource", &opnsense.RandomArgs{
			Length: pulumi.Int(24),
		})
		if err != nil {
			return err
		}
		_, err = opnsense.NewRandomComponent(ctx, "myRandomComponent", &opnsense.RandomComponentArgs{
			Length: pulumi.Int(24),
		})
		if err != nil {
			return err
		}
		ctx.Export("output", pulumi.StringMap{
			"value": myRandomResource.Result,
		})
		return nil
	})
}
