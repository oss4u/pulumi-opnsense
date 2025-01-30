package config

import (
	"context"

	"github.com/oss4u/go-opnsense/opnsense"
	"github.com/pulumi/pulumi-go-provider/infer"
)

type Config struct {
	Address string `pulumi:"address" provider:"secret"`
	Key     string `pulumi:"key" provider:"secret"`
	Secret  string `pulumi:"secret" provider:"secret"`
	// Version string `pulumi:"version"`
	//  PluginDownloadURL string `pulumi:"pluginDownloadURL"`
	Api *opnsense.OpnSenseApi
}

var _ = (infer.Annotated)((*Config)(nil))

func (c *Config) Annotate(a infer.Annotator) {
	a.Describe(&c.Address, "The address of the fw. (without /api)")
	a.Describe(&c.Key, "The key to access the api of the fw.")
	a.Describe(&c.Secret, `The secret to access the api of the fw.`)
}

var _ = (infer.CustomConfigure)((*Config)(nil))

func (c *Config) Configure(ctx context.Context) error {
	c.Api = opnsense.GetOpnSenseClient(c.Address, c.Key, c.Secret)
	return nil
}
