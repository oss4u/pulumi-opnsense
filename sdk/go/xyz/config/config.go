// Code generated by pulumi-language-go DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package config

import (
	"example.com/pulumi-xyz/sdk/go/xyz/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi/config"
)

var _ = internal.GetEnvOrDefault

func GetItsasecret(ctx *pulumi.Context) bool {
	return config.GetBool(ctx, "xyz:itsasecret")
}
