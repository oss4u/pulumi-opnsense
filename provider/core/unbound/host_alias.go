package unbound

import (
	"context"
	gooverrides "github.com/oss4u/go-opnsense/opnsense/core/unbound/overrides"
	"github.com/oss4u/pulumi-opnsense/provider/core/config"
	p "github.com/pulumi/pulumi-go-provider"
	"github.com/pulumi/pulumi-go-provider/infer"
)

type HostAliasOverride struct{}

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

type HostAliasOverrideState struct {
	// It is generally a good idea to embed args in outputs, but it isn't strictly necessary.
	HostAliasOverrideArgs
	// Here we define a required output called result.
	Id string `pulumi:"result"`
}

func (HostAliasOverride) GetApi(ctx context.Context) gooverrides.OverridesAliasesApi {
	cfg := infer.GetConfig[config.Config](ctx)

	return gooverrides.GetAliasesOverrideApi(cfg.Api)
}

func (h HostAliasOverride) Create(ctx context.Context, name string, input HostAliasOverrideArgs, preview bool) (string, HostAliasOverrideState, error) {
	state := HostAliasOverrideState{HostAliasOverrideArgs: input}
	if preview {
		return name, state, nil
	}
	var err error
	state.Id, err = h.createHostAlias(ctx, &input)
	return state.Id, state, err
}

func (h HostAliasOverride) Delete(ctx context.Context, id string, _ HostAliasOverrideState) error {
	err := h.deleteHostAliasOverride(ctx, id)
	return err
}

func (h HostAliasOverride) Update(ctx context.Context, id string, _ HostAliasOverrideState, news HostAliasOverrideArgs, preview bool) (HostAliasOverrideState, error) {
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

func (h HostAliasOverride) Read(ctx context.Context, id string, inputs HostAliasOverrideArgs, _ HostAliasOverrideState) (canonicalID string, normalizedInputs HostAliasOverrideArgs, normalizedState HostAliasOverrideState, err error) {
	overrides := h.GetApi(ctx)
	host, err := overrides.Read(id)
	newArgs := OverridesAliasToHostAliasOverrideArgs(host)
	return id, inputs, HostAliasOverrideState{
		HostAliasOverrideArgs: newArgs,
		Id:                    id,
	}, err
}

func (h HostAliasOverride) Diff(ctx context.Context, id string, _ HostAliasOverrideState, new HostAliasOverrideArgs) (p.DiffResponse, error) {
	overrides := h.GetApi(ctx)
	result, err := overrides.Read(id)
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
		diffs["enabled"] = p.PropertyDiff{
			Kind: p.Update,
		}
	}
	if details.Hostname != *new.Hostname {
		diffs["hostname"] = p.PropertyDiff{
			Kind: p.Update,
		}
	}
	if details.Host != *new.Host {
		diffs["host"] = p.PropertyDiff{
			Kind: p.Update,
		}
	}
	if details.Domain != *new.Domain {
		diffs["domain"] = p.PropertyDiff{
			Kind: p.Update,
		}
	}
	if details.Description != *new.Description {
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

func (h HostAliasOverride) deleteHostAliasOverride(ctx context.Context, id string) error {
	overrides := h.GetApi(ctx)
	err := overrides.DeleteByID(id)
	return err
}

func (h HostAliasOverride) createHostAlias(ctx context.Context, args *HostAliasOverrideArgs) (string, error) {
	overrides := h.GetApi(ctx)
	newHost := HostAliasOverrideArgsToOverridesAlias(args)
	createdAlias, err := overrides.Create(&newHost)
	if err != nil {
		return "", err
	}
	return createdAlias.Alias.Uuid, nil
}
