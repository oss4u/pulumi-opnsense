package unbound

import (
	"fmt"
	gooverrides "github.com/oss4u/go-opnsense/opnsense/core/unbound/overrides"
	"github.com/oss4u/pulumi-opnsense-native/cmd/pulumi-resource-opnsense/core/config"
	p "github.com/pulumi/pulumi-go-provider"
	"github.com/pulumi/pulumi-go-provider/infer"
	"github.com/pulumi/pulumi/sdk/v3/go/common/diag"
)

type HostOverride struct{}

type HostOverrideArgs struct {
	Enabled     *bool   `pulumi:"enabled"`
	Hostname    *string `pulumi:"hostname"`
	Domain      *string `pulumi:"domain"`
	Rr          *string `pulumi:"rr"`
	MxPrio      *int    `pulumi:"mx_prio,optional"`
	Mx          *string `pulumi:"mx,optional"`
	Server      *string `pulumi:"server,optional"`
	Description *string `pulumi:"description"`
	//Aliases     *[]HostAliasOverrideArgs `pulumi:"aliases,optional"`
}

var _ = (infer.CustomRead[HostOverrideArgs, HostOverrideState])((*HostOverride)(nil))
var _ = (infer.CustomUpdate[HostOverrideArgs, HostOverrideState])((*HostOverride)(nil))
var _ = (infer.CustomDelete[HostOverrideState])((*HostOverride)(nil))

type HostOverrideState struct {
	// It is generally a good idea to embed args in outputs, but it isn't strictly necessary.
	HostOverrideArgs
	// Here we define a required output called result.
	Id string `pulumi:"result"`
}

func (HostOverride) GetApi(ctx p.Context) gooverrides.OverridesHostsApi {
	cfg := infer.GetConfig[config.Config](ctx)

	return gooverrides.GetHostsOverrideApi(cfg.Api)
}

func (h HostOverride) Create(ctx p.Context, name string, input HostOverrideArgs, preview bool) (string, HostOverrideState, error) {
	ctx.Log(diag.Debug, "Running CREATE")
	state := HostOverrideState{HostOverrideArgs: input}
	if preview {
		return name, state, nil
	}
	var err error
	state.Id, err = h.createHostOverride(ctx, &input)
	return state.Id, state, err
}

func (h HostOverride) Delete(ctx p.Context, id string, _ HostOverrideState) error {
	ctx.Log(diag.Debug, "Running DELETE")
	err := h.deleteHostOverride(ctx, id)
	return err
}

func (h HostOverride) Update(ctx p.Context, id string, _ HostOverrideState, news HostOverrideArgs, preview bool) (HostOverrideState, error) {
	ctx.Log(diag.Debug, "Running UPDATE")
	if preview {
		return HostOverrideState{
			HostOverrideArgs: news,
		}, nil
	}
	overrides := h.GetApi(ctx)
	host := HostOverrideArgsToOverridesHost(&news)
	host.Host.Uuid = id
	_, err := overrides.Update(host)
	return HostOverrideState{
		HostOverrideArgs: news,
	}, err
}

func (h HostOverride) Read(ctx p.Context, id string, inputs HostOverrideArgs, _ HostOverrideState) (canonicalID string, normalizedInputs HostOverrideArgs, normalizedState HostOverrideState, err error) {
	ctx.Log(diag.Debug, "Running READ")
	overrides := h.GetApi(ctx)
	host, err := overrides.Read(id)
	newArgs := OverridesHostToHostOverrideArgs(host)
	return id, inputs, HostOverrideState{
		HostOverrideArgs: *newArgs,
		Id:               id,
	}, err
}

func (h HostOverride) Diff(ctx p.Context, id string, _ HostOverrideState, new HostOverrideArgs) (p.DiffResponse, error) {
	ctx.Log(diag.Debug, "Running DIFF")
	overrides := h.GetApi(ctx)
	result, err := overrides.Read(id)
	if result == nil || result.Host.Hostname == "" {
		return p.DiffResponse{
			DeleteBeforeReplace: true,
			HasChanges:          true,
			DetailedDiff:        nil,
		}, err
	}
	diffs := map[string]p.PropertyDiff{}
	if result.Host.Hostname != *new.Hostname {
		ctx.Log(diag.Debug, fmt.Sprintf("Hostname differs: %s/%s", result.Host.Hostname, *new.Hostname))
		diffs["hostname"] = p.PropertyDiff{
			Kind: p.Update,
		}
	}
	if result.Host.Domain != *new.Domain {
		ctx.Log(diag.Debug, fmt.Sprintf("Domain differs: %s/%s", result.Host.Domain, *new.Domain))
		diffs["domain"] = p.PropertyDiff{
			Kind: p.Update,
		}
	}
	if result.Host.Description != *new.Description {
		ctx.Log(diag.Debug, fmt.Sprintf("Description differs: %s/%s", result.Host.Description, *new.Description))
		diffs["description"] = p.PropertyDiff{
			Kind: p.Update,
		}
	}
	if result.Host.Enabled.Bool() != *new.Enabled {
		ctx.Log(diag.Debug, fmt.Sprintf("Enabled differs: %t/%t", result.Host.Enabled.Bool(), *new.Enabled))
		diffs["enabled"] = p.PropertyDiff{
			Kind: p.Update,
		}
	}
	if result.Host.Rr.String() == "A" {
		if result.Host.Server != *new.Server {
			ctx.Log(diag.Debug, fmt.Sprintf("Server differs: %s/%s", result.Host.Server, *new.Server))
			diffs["server"] = p.PropertyDiff{
				Kind: p.Update,
			}
		}
	} else if result.Host.Rr.String() == "MX" {
		if result.Host.Mx != *new.Mx {
			ctx.Log(diag.Debug, fmt.Sprintf("Mx differs: %s/%s", result.Host.Mx, *new.Mx))
			diffs["mx"] = p.PropertyDiff{
				Kind: p.Update,
			}
		}
		if result.Host.Mxprio.Int() != *new.MxPrio {
			ctx.Log(diag.Debug, fmt.Sprintf("MxPrio differs: %d/%d", result.Host.Mxprio.Int(), *new.MxPrio))
			diffs["mxprio"] = p.PropertyDiff{
				Kind: p.Update,
			}
		}

	}
	diff := p.DiffResponse{
		DeleteBeforeReplace: false,
		HasChanges:          len(diffs) > 0,
		DetailedDiff:        diffs,
	}
	return diff, nil
}

func (h HostOverride) deleteHostOverride(ctx p.Context, id string) error {
	overrides := h.GetApi(ctx)
	err := overrides.DeleteByID(id)
	return err
}

func (h HostOverride) createHostOverride(ctx p.Context, args *HostOverrideArgs) (string, error) {
	overrides := h.GetApi(ctx)
	newHost := HostOverrideArgsToOverridesHost(args)
	createdHost, err := overrides.Create(newHost)
	if err != nil {
		return "", err
	}
	return createdHost.Host.GetUUID(), err
}
