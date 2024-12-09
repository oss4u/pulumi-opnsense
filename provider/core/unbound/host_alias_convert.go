package unbound

import (
	"github.com/oss4u/go-opnsense/opnsense/core/unbound/overrides"
	"github.com/oss4u/go-opnsense/opnsense/types"
)

func OverridesAliasToHostAliasOverrideArgs(alias *overrides.OverridesAlias) HostAliasOverrideArgs {
	enabled := alias.Alias.Enabled.Bool()
	result := HostAliasOverrideArgs{
		Enabled:     &enabled,
		Host:        &alias.Alias.Host,
		Hostname:    &alias.Alias.Hostname,
		Domain:      &alias.Alias.Domain,
		Description: &alias.Alias.Description,
	}
	return result

}

func HostAliasOverrideArgsToOverridesAlias(args *HostAliasOverrideArgs) overrides.OverridesAlias {
	result := overrides.OverridesAlias{Alias: overrides.OverridesAliasDetails{
		Enabled:     types.Bool(*args.Enabled),
		Host:        *args.Host,
		Hostname:    *args.Hostname,
		Domain:      *args.Domain,
		Description: *args.Description,
	}}
	return result
}
