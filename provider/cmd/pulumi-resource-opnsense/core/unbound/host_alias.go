package unbound

import (
	"fmt"
	gooverrides "github.com/oss4u/go-opnsense/opnsense/core/unbound/overrides"
	"github.com/oss4u/pulumi-opnsense-native/cmd/pulumi-resource-opnsense/core/config"
	p "github.com/pulumi/pulumi-go-provider"
	"github.com/pulumi/pulumi-go-provider/infer"
	"github.com/pulumi/pulumi/sdk/v3/go/common/diag"
)

// Each resource has a controlling struct.
// Resource behavior is determined by implementing methods on the controlling struct.
// The `Create` method is mandatory, but other methods are optional.
// - Check: Remap inputs before they are typed.
// - Diff: Change how instances of a resource are compared.
// - Update: Mutate a resource in place.
// - Read: Get the state of a resource from the backing provider.
// - Delete: Custom logic when the resource is deleted.
// - Annotate: Describe fields and set defaults for a resource.
// - WireDependencies: Control how outputs and secrets flows through values.
type HostAliasOverride struct{}

// Each resource has in input struct, defining what arguments it accepts.
type HostAliasOverrideArgs struct {
	Enabled     *bool   `pulumi:"enabled"`
	Host        *string `pulumi:"host"`
	Hostname    *string `pulumi:"hostname"`
	Domain      *string `pulumi:"domain"`
	Description *string `pulumi:"description"`
}

func NewHostAliasOverrideArgs(enabled bool, host string, hostname string, domain string, description string) *HostAliasOverrideArgs {
	return &HostAliasOverrideArgs{
		Enabled:     &enabled,
		Host:        &host,
		Hostname:    &hostname,
		Domain:      &domain,
		Description: &description,
	}
}

var _ = (infer.CustomRead[HostAliasOverrideArgs, HostAliasOverrideState])((*HostAliasOverride)(nil))
var _ = (infer.CustomUpdate[HostAliasOverrideArgs, HostAliasOverrideState])((*HostAliasOverride)(nil))
var _ = (infer.CustomDelete[HostAliasOverrideState])((*HostAliasOverride)(nil))
var _ = (infer.CustomDiff[HostAliasOverrideArgs, HostAliasOverrideState])((*HostAliasOverride)(nil))

// Each resource has a state, describing the fields that exist on the created resource.
type HostAliasOverrideState struct {
	// It is generally a good idea to embed args in outputs, but it isn't strictly necessary.
	HostAliasOverrideArgs
	// Here we define a required output called result.
	Id string `pulumi:"result"`
}

func (HostAliasOverride) GetApi(ctx p.Context) gooverrides.OverridesAliasesApi {
	cfg := infer.GetConfig[config.Config](ctx)

	return gooverrides.GetAliasesOverrideApi(cfg.Api)
}

// All resources must implement Create at a minumum.
func (h HostAliasOverride) Create(ctx p.Context, name string, input HostAliasOverrideArgs, preview bool) (string, HostAliasOverrideState, error) {
	ctx.Log(diag.Info, "Running CREATE")
	state := HostAliasOverrideState{HostAliasOverrideArgs: input}
	if preview {
		return name, state, nil
	}
	var err error
	state.Id, err = h.createHostAlias(ctx, &input)
	return state.Id, state, err
}

func (h HostAliasOverride) Delete(ctx p.Context, id string, args HostAliasOverrideState) error {
	ctx.Log(diag.Info, "Running DELETE")
	err := h.deleteHostAliasOverride(ctx, id)
	return err
}

func (h HostAliasOverride) Update(ctx p.Context, id string, old HostAliasOverrideState, news HostAliasOverrideArgs, preview bool) (HostAliasOverrideState, error) {
	ctx.Log(diag.Info, "Running UPDATE")
	if preview {
		return HostAliasOverrideState{
			HostAliasOverrideArgs: news,
		}, nil
	}
	overrides := h.GetApi(ctx)
	host := HostAliasOverrideArgsToOverridesAlias(&news)
	host.Alias.Uuid = id
	_, err := overrides.Update(&host)
	return HostAliasOverrideState{
		HostAliasOverrideArgs: news,
	}, err
}

func (h HostAliasOverride) Read(ctx p.Context, id string, inputs HostAliasOverrideArgs, state HostAliasOverrideState) (canonicalID string, normalizedInputs HostAliasOverrideArgs, normalizedState HostAliasOverrideState, err error) {
	ctx.Log(diag.Info, "Running READ")
	overrides := h.GetApi(ctx)
	host, err := overrides.Read(id)
	newArgs := OverridesAliasToHostAliasOverrideArgs(host)
	return id, inputs, HostAliasOverrideState{
		HostAliasOverrideArgs: newArgs,
		Id:                    id,
	}, err
}

func (h HostAliasOverride) Diff(ctx p.Context, id string, old HostAliasOverrideState, new HostAliasOverrideArgs) (p.DiffResponse, error) {
	ctx.Log(diag.Info, "Running DIFF")
	overrides := h.GetApi(ctx)
	result, err := overrides.Read(id)
	ctx.Log(diag.Info, fmt.Sprintf("Retval: %+v", result))
	details := result.Alias
	if result == nil || details.Host == "" {
		return p.DiffResponse{
			DeleteBeforeReplace: true,
			HasChanges:          true,
			DetailedDiff:        nil,
		}, err
	}
	diffs := map[string]p.PropertyDiff{}
	if details.Enabled.Bool() != *new.Enabled {
		ctx.Log(diag.Info, fmt.Sprintf("Enabled differs: %t/%t", details.Enabled.Bool(), *new.Enabled))
		diffs["enabled"] = p.PropertyDiff{
			Kind: p.Update,
		}
	}
	if details.Hostname != *new.Hostname {
		ctx.Log(diag.Info, fmt.Sprintf("Hostname differs: %s/%s", details.Hostname, *new.Hostname))
		diffs["hostname"] = p.PropertyDiff{
			Kind: p.Update,
		}
	}
	if details.Host != *new.Host {
		ctx.Log(diag.Info, fmt.Sprintf("Host differs: %s/%s", details.Host, *new.Host))
		diffs["host"] = p.PropertyDiff{
			Kind: p.Update,
		}
	}
	if details.Domain != *new.Domain {
		ctx.Log(diag.Info, fmt.Sprintf("Domain differs: %s/%s", details.Domain, *new.Domain))
		diffs["domain"] = p.PropertyDiff{
			Kind: p.Update,
		}
	}
	if details.Description != *new.Description {
		ctx.Log(diag.Info, fmt.Sprintf("Description differs: %s/%s", details.Description, *new.Description))
		diffs["description"] = p.PropertyDiff{
			Kind: p.Update,
		}
	}
	diff := p.DiffResponse{
		DeleteBeforeReplace: false,
		HasChanges:          len(diffs) > 0,
		DetailedDiff:        diffs,
	}
	return diff, nil

}

func (h HostAliasOverride) deleteHostAliasOverride(ctx p.Context, id string) error {
	overrides := h.GetApi(ctx)
	err := overrides.DeleteByID(id)
	return err
}

func (h HostAliasOverride) createHostAlias(ctx p.Context, args *HostAliasOverrideArgs) (string, error) {
	overrides := h.GetApi(ctx)
	newHost := HostAliasOverrideArgsToOverridesAlias(args)
	createdAlias, err := overrides.Create(&newHost)
	if err != nil {
		return "", err
	}
	return createdAlias.Alias.Uuid, nil
}
