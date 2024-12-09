// Code generated by pulumi-language-go DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package opnsense

import (
	"context"
	"reflect"

	"errors"
	"example.com/pulumi-opnsense/sdk/go/opnsense/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type Provider struct {
	pulumi.ProviderResourceState

	// The address of the fw. (without /api)
	Address pulumi.StringOutput `pulumi:"address"`
	// The key to access the api of the fw.
	Key pulumi.StringOutput `pulumi:"key"`
	// The secret to access the api of the fw.
	Secret pulumi.StringOutput `pulumi:"secret"`
}

// NewProvider registers a new resource with the given unique name, arguments, and options.
func NewProvider(ctx *pulumi.Context,
	name string, args *ProviderArgs, opts ...pulumi.ResourceOption) (*Provider, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.Address == nil {
		return nil, errors.New("invalid value for required argument 'Address'")
	}
	if args.Key == nil {
		return nil, errors.New("invalid value for required argument 'Key'")
	}
	if args.Secret == nil {
		return nil, errors.New("invalid value for required argument 'Secret'")
	}
	if args.Address != nil {
		args.Address = pulumi.ToSecret(args.Address).(pulumi.StringInput)
	}
	if args.Key != nil {
		args.Key = pulumi.ToSecret(args.Key).(pulumi.StringInput)
	}
	if args.Secret != nil {
		args.Secret = pulumi.ToSecret(args.Secret).(pulumi.StringInput)
	}
	secrets := pulumi.AdditionalSecretOutputs([]string{
		"address",
		"key",
		"secret",
	})
	opts = append(opts, secrets)
	opts = internal.PkgResourceDefaultOpts(opts)
	var resource Provider
	err := ctx.RegisterResource("pulumi:providers:opnsense", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

type providerArgs struct {
	// The address of the fw. (without /api)
	Address string `pulumi:"address"`
	// The key to access the api of the fw.
	Key string `pulumi:"key"`
	// The secret to access the api of the fw.
	Secret string `pulumi:"secret"`
}

// The set of arguments for constructing a Provider resource.
type ProviderArgs struct {
	// The address of the fw. (without /api)
	Address pulumi.StringInput
	// The key to access the api of the fw.
	Key pulumi.StringInput
	// The secret to access the api of the fw.
	Secret pulumi.StringInput
}

func (ProviderArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*providerArgs)(nil)).Elem()
}

type ProviderInput interface {
	pulumi.Input

	ToProviderOutput() ProviderOutput
	ToProviderOutputWithContext(ctx context.Context) ProviderOutput
}

func (*Provider) ElementType() reflect.Type {
	return reflect.TypeOf((**Provider)(nil)).Elem()
}

func (i *Provider) ToProviderOutput() ProviderOutput {
	return i.ToProviderOutputWithContext(context.Background())
}

func (i *Provider) ToProviderOutputWithContext(ctx context.Context) ProviderOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ProviderOutput)
}

type ProviderOutput struct{ *pulumi.OutputState }

func (ProviderOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**Provider)(nil)).Elem()
}

func (o ProviderOutput) ToProviderOutput() ProviderOutput {
	return o
}

func (o ProviderOutput) ToProviderOutputWithContext(ctx context.Context) ProviderOutput {
	return o
}

// The address of the fw. (without /api)
func (o ProviderOutput) Address() pulumi.StringOutput {
	return o.ApplyT(func(v *Provider) pulumi.StringOutput { return v.Address }).(pulumi.StringOutput)
}

// The key to access the api of the fw.
func (o ProviderOutput) Key() pulumi.StringOutput {
	return o.ApplyT(func(v *Provider) pulumi.StringOutput { return v.Key }).(pulumi.StringOutput)
}

// The secret to access the api of the fw.
func (o ProviderOutput) Secret() pulumi.StringOutput {
	return o.ApplyT(func(v *Provider) pulumi.StringOutput { return v.Secret }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*ProviderInput)(nil)).Elem(), &Provider{})
	pulumi.RegisterOutputType(ProviderOutput{})
}
