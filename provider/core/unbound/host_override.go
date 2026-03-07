package unbound

import (
	"context"

	"github.com/oss4u/go-opnsense/opnsense"
	gooverrides "github.com/oss4u/go-opnsense/opnsense/core/unbound/overrides"
	"github.com/oss4u/pulumi-opnsense/provider/core/config"
	p "github.com/pulumi/pulumi-go-provider"
	"github.com/pulumi/pulumi-go-provider/infer"
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
var _ = (infer.CustomDiff[HostOverrideArgs, HostOverrideState])((*HostOverride)(nil))

type HostOverrideState struct {
	// It is generally a good idea to embed args in outputs, but it isn't strictly necessary.
	HostOverrideArgs
	// Here we define a required output called result.
	Id string `pulumi:"result"`
}

func (HostOverride) GetApi(ctx context.Context) gooverrides.OverridesHostsApi {
	cfg := infer.GetConfig[config.Config](ctx)
	if cfg.Api == nil {
		cfg.Api = opnsense.GetOpnSenseClient(cfg.Address, cfg.Key, cfg.Secret)
	}

	return gooverrides.GetHostsOverrideApi(cfg.Api)
}

func (h HostOverride) Create(ctx context.Context, req infer.CreateRequest[HostOverrideArgs]) (infer.CreateResponse[HostOverrideState], error) {
	state := HostOverrideState{HostOverrideArgs: req.Inputs}
	if req.DryRun {
		return infer.CreateResponse[HostOverrideState]{
			ID:     req.Name,
			Output: state,
		}, nil
	}
	var err error
	state.Id, err = h.createHostOverride(ctx, &req.Inputs)
	return infer.CreateResponse[HostOverrideState]{
		ID:     state.Id,
		Output: state,
	}, err
}

func (h HostOverride) Delete(ctx context.Context, req infer.DeleteRequest[HostOverrideState]) (infer.DeleteResponse, error) {
	err := h.deleteHostOverride(ctx, req.ID)
	return infer.DeleteResponse{}, err
}

func (h HostOverride) Update(ctx context.Context, req infer.UpdateRequest[HostOverrideArgs, HostOverrideState]) (infer.UpdateResponse[HostOverrideState], error) {
	if req.DryRun {
		return infer.UpdateResponse[HostOverrideState]{
			Output: HostOverrideState{
				HostOverrideArgs: req.Inputs,
			},
		}, nil
	}
	overrides := h.GetApi(ctx)
	host := HostOverrideArgsToOverridesHost(&req.Inputs)
	host.Host.Uuid = req.ID
	_, err := overrides.Update(host)
	return infer.UpdateResponse[HostOverrideState]{
		Output: HostOverrideState{
			HostOverrideArgs: req.Inputs,
			Id:               req.ID,
		},
	}, err
}

func (h HostOverride) Read(ctx context.Context, req infer.ReadRequest[HostOverrideArgs, HostOverrideState]) (infer.ReadResponse[HostOverrideArgs, HostOverrideState], error) {
	overrides := h.GetApi(ctx)
	host, err := overrides.Read(req.ID)
	if err != nil {
		return infer.ReadResponse[HostOverrideArgs, HostOverrideState]{}, err
	}
	newArgs := OverridesHostToHostOverrideArgs(host)
	return infer.ReadResponse[HostOverrideArgs, HostOverrideState]{
		ID:     req.ID,
		Inputs: req.Inputs,
		State: HostOverrideState{
			HostOverrideArgs: *newArgs,
			Id:               req.ID,
		},
	}, nil
}

func (h HostOverride) Diff(ctx context.Context, req infer.DiffRequest[HostOverrideArgs, HostOverrideState]) (infer.DiffResponse, error) {
	overrides := h.GetApi(ctx)
	result, err := overrides.Read(req.ID)
	if result == nil || result.Host.Hostname == "" {
		return infer.DiffResponse{
			DeleteBeforeReplace: true,
			HasChanges:          true,
			DetailedDiff:        nil,
		}, err
	}
	diffs := map[string]p.PropertyDiff{}
	if result.Host.Hostname != *req.Inputs.Hostname {
		diffs["hostname"] = p.PropertyDiff{
			Kind: p.Update,
		}
	}
	if result.Host.Domain != *req.Inputs.Domain {
		diffs["domain"] = p.PropertyDiff{
			Kind: p.Update,
		}
	}
	if result.Host.Description != *req.Inputs.Description {
		diffs["description"] = p.PropertyDiff{
			Kind: p.Update,
		}
	}
	if result.Host.Enabled.Bool() != *req.Inputs.Enabled {
		diffs["enabled"] = p.PropertyDiff{
			Kind: p.Update,
		}
	}
	if result.Host.Rr.String() == "A" {
		if result.Host.Server != *req.Inputs.Server {
			diffs["server"] = p.PropertyDiff{
				Kind: p.Update,
			}
		}
	} else if result.Host.Rr.String() == "MX" {
		if result.Host.Mx != *req.Inputs.Mx {
			diffs["mx"] = p.PropertyDiff{
				Kind: p.Update,
			}
		}
		if result.Host.Mxprio.Int() != *req.Inputs.MxPrio {
			diffs["mxprio"] = p.PropertyDiff{
				Kind: p.Update,
			}
		}

	}
	diff := infer.DiffResponse{
		DeleteBeforeReplace: false,
		HasChanges:          len(diffs) > 0,
		DetailedDiff:        diffs,
	}
	return diff, nil
}

func (h HostOverride) deleteHostOverride(ctx context.Context, id string) error {
	overrides := h.GetApi(ctx)
	err := overrides.DeleteByID(id)
	return err
}

func (h HostOverride) createHostOverride(ctx context.Context, args *HostOverrideArgs) (string, error) {
	overrides := h.GetApi(ctx)
	newHost := HostOverrideArgsToOverridesHost(args)
	createdHost, err := overrides.Create(newHost)
	if err != nil {
		return "", err
	}
	return createdHost.Host.GetUUID(), err
}
