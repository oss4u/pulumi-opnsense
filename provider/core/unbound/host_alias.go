package unbound

import (
	"context"
	"encoding/json"
	"fmt"
	"strings"

	"github.com/oss4u/go-opnsense/opnsense"
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
	if cfg.Api == nil {
		cfg.Api = opnsense.GetOpnSenseClient(cfg.Address, cfg.Key, cfg.Secret)
	}

	return gooverrides.GetAliasesOverrideApi(cfg.Api)
}

func (h HostAliasOverride) Create(ctx context.Context, req infer.CreateRequest[HostAliasOverrideArgs]) (infer.CreateResponse[HostAliasOverrideState], error) {
	state := HostAliasOverrideState{HostAliasOverrideArgs: req.Inputs}
	if req.DryRun {
		return infer.CreateResponse[HostAliasOverrideState]{
			ID:     req.Name,
			Output: state,
		}, nil
	}
	var err error
	state.Id, err = h.createHostAlias(ctx, &req.Inputs)
	return infer.CreateResponse[HostAliasOverrideState]{
		ID:     state.Id,
		Output: state,
	}, err
}

func (h HostAliasOverride) Delete(ctx context.Context, req infer.DeleteRequest[HostAliasOverrideState]) (infer.DeleteResponse, error) {
	err := h.deleteHostAliasOverride(ctx, req.ID)
	return infer.DeleteResponse{}, err
}

func (h HostAliasOverride) Update(ctx context.Context, req infer.UpdateRequest[HostAliasOverrideArgs, HostAliasOverrideState]) (infer.UpdateResponse[HostAliasOverrideState], error) {
	if req.DryRun {
		return infer.UpdateResponse[HostAliasOverrideState]{
			Output: HostAliasOverrideState{
				HostAliasOverrideArgs: req.Inputs,
			},
		}, nil
	}
	overrides := h.GetApi(ctx)
	host := HostAliasOverrideArgsToOverridesAlias(&req.Inputs)
	host.Alias.Uuid = req.ID
	_, err := overrides.Update(&host)
	return infer.UpdateResponse[HostAliasOverrideState]{
		Output: HostAliasOverrideState{
			HostAliasOverrideArgs: req.Inputs,
			Id:                    req.ID,
		},
	}, err
}

func (h HostAliasOverride) Read(ctx context.Context, req infer.ReadRequest[HostAliasOverrideArgs, HostAliasOverrideState]) (infer.ReadResponse[HostAliasOverrideArgs, HostAliasOverrideState], error) {
	overrides := h.GetApi(ctx)
	host, err := overrides.Read(req.ID)
	if err != nil {
		return infer.ReadResponse[HostAliasOverrideArgs, HostAliasOverrideState]{}, err
	}
	newArgs := OverridesAliasToHostAliasOverrideArgs(host)
	return infer.ReadResponse[HostAliasOverrideArgs, HostAliasOverrideState]{
		ID:     req.ID,
		Inputs: req.Inputs,
		State: HostAliasOverrideState{
			HostAliasOverrideArgs: newArgs,
			Id:                    req.ID,
		},
	}, nil
}

func (h HostAliasOverride) Diff(ctx context.Context, req infer.DiffRequest[HostAliasOverrideArgs, HostAliasOverrideState]) (infer.DiffResponse, error) {
	overrides := h.GetApi(ctx)
	result, err := overrides.Read(req.ID)
	details := result.Alias
	if result == nil || details.Host == "" {
		return infer.DiffResponse{
			DeleteBeforeReplace: true,
			HasChanges:          true,
			DetailedDiff:        nil,
		}, err
	}
	diffs := map[string]p.PropertyDiff{}
	if details.Enabled.Bool() != *req.Inputs.Enabled {
		diffs["enabled"] = p.PropertyDiff{
			Kind: p.Update,
		}
	}
	if details.Hostname != *req.Inputs.Hostname {
		diffs["hostname"] = p.PropertyDiff{
			Kind: p.Update,
		}
	}
	if details.Host != *req.Inputs.Host {
		diffs["host"] = p.PropertyDiff{
			Kind: p.Update,
		}
	}
	if details.Domain != *req.Inputs.Domain {
		diffs["domain"] = p.PropertyDiff{
			Kind: p.Update,
		}
	}
	if details.Description != *req.Inputs.Description {
		diffs["description"] = p.PropertyDiff{
			Kind: p.Update,
		}
	}
	diff := infer.DiffResponse{
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

	if createdAlias != nil {
		if id := strings.TrimSpace(createdAlias.Alias.Uuid); id != "" {
			return id, nil
		}
	}

	resolvedID, resolveErr := h.findHostAliasUUIDByHostDomain(ctx, *args.Hostname, *args.Domain)
	if resolveErr != nil {
		return "", resolveErr
	}
	if resolvedID == "" {
		return "", fmt.Errorf("host alias was created but no uuid was returned for %s.%s", *args.Hostname, *args.Domain)
	}

	return resolvedID, nil
}

func (h HostAliasOverride) findHostAliasUUIDByHostDomain(ctx context.Context, hostname, domain string) (string, error) {
	cfg := infer.GetConfig[config.Config](ctx)
	if cfg.Api == nil {
		cfg.Api = opnsense.GetOpnSenseClient(cfg.Address, cfg.Key, cfg.Secret)
	}

	searchPayload := map[string]any{
		"current":  1,
		"rowCount": 500,
	}

	payloadRaw, err := json.Marshal(searchPayload)
	if err != nil {
		return "", err
	}

	raw, err := cfg.Api.ModifyingRequest("unbound", "settings", "search_host_alias", string(payloadRaw), []string{})
	if err != nil {
		return "", err
	}

	var decoded map[string]any
	if err := json.Unmarshal([]byte(raw), &decoded); err != nil {
		return "", err
	}

	rows, ok := decoded["rows"].([]any)
	if !ok {
		return "", nil
	}

	for _, row := range rows {
		rowMap, ok := row.(map[string]any)
		if !ok {
			continue
		}

		candidateHost, _ := rowMap["hostname"].(string)
		candidateDomain, _ := rowMap["domain"].(string)
		uuid, _ := rowMap["uuid"].(string)

		if strings.EqualFold(strings.TrimSpace(candidateHost), strings.TrimSpace(hostname)) &&
			strings.EqualFold(strings.TrimSpace(candidateDomain), strings.TrimSpace(domain)) {
			return strings.TrimSpace(uuid), nil
		}
	}

	return "", nil
}
