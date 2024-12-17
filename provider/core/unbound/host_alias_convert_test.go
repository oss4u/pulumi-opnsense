package unbound

import (
	"github.com/oss4u/go-opnsense/opnsense/core/unbound/overrides"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestHostAliasOverrideArgsToOverridesAlias(t *testing.T) {
	type args struct {
		args *HostAliasOverrideArgs
	}
	tests := []struct {
		name string
		args args
		want overrides.OverridesAlias
	}{
		{
			name: "Working-Enabled=TRUE",
			args: args{NewHostAliasOverrideArgs(
				true,
				"host-uuid",
				"host",
				"domain.com",
				"Alias for host",
			)},
			want: *overrides.NewOverridesAlias(
				true,
				"host-uuid",
				"host",
				"domain.com",
				"Alias for host"),
		},
		{
			name: "Working-Enabled=FALSE",
			args: args{NewHostAliasOverrideArgs(
				false,
				"host-uuid",
				"host",
				"domain.com",
				"Alias for host",
			)},
			want: *overrides.NewOverridesAlias(
				false,
				"host-uuid",
				"host",
				"domain.com",
				"Alias for host"),
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, tt.want, HostAliasOverrideArgsToOverridesAlias(tt.args.args))
		})
	}
}

func TestOverridesAliasToHostAliasOverrideArgs(t *testing.T) {
	type args struct {
		alias *overrides.OverridesAlias
	}
	tests := []struct {
		name string
		args args
		want HostAliasOverrideArgs
	}{
		{
			name: "Working-Enabled=TRUE",
			args: args{overrides.NewOverridesAlias(
				true,
				"host-uuid",
				"host",
				"domain.com",
				"Alias for host",
			)},
			want: *NewHostAliasOverrideArgs(
				true,
				"host-uuid",
				"host",
				"domain.com",
				"Alias for host"),
		},
		{
			name: "Working-Enabled=FALSE",
			args: args{overrides.NewOverridesAlias(
				false,
				"host-uuid",
				"host",
				"domain.com",
				"Alias for host",
			)},
			want: *NewHostAliasOverrideArgs(
				false,
				"host-uuid",
				"host",
				"domain.com",
				"Alias for host"),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, tt.want, OverridesAliasToHostAliasOverrideArgs(tt.args.alias))
		})
	}
}
