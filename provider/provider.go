// Copyright 2016-2023, Pulumi Corporation.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package provider

import (
	"github.com/oss4u/pulumi-opnsense/provider/core/config"
	"github.com/oss4u/pulumi-opnsense/provider/core/unbound"
	p "github.com/pulumi/pulumi-go-provider"
	"github.com/pulumi/pulumi-go-provider/infer"
	"github.com/pulumi/pulumi-go-provider/middleware/schema"
	"github.com/pulumi/pulumi/sdk/v3/go/common/tokens"
)

// Version is initialized by the Go linker to contain the semver of this build.
var Version string

const Name string = "opnsense"

func Provider() p.Provider {
	// We tell the provider what resources it needs to support.
	// In this case, a single resource and component
	return infer.Provider(infer.Options{
		Metadata: schema.Metadata{
			DisplayName: "OpnSense",
			Description: "The Pulumi OpnSense provider is used to interact with the resources supported by OpnSense.",
			Keywords:    []string{"pulumi", "opnsense"},
			Homepage:    "https://github.com/oss4u/pulumi-opnsense",
			License:     "Apache-2.0",
			Repository:  "https://github.com/oss4u/pulumi-opnsense",
			Publisher:   "oss4u",
			LanguageMap: map[string]any{
				"csharp": map[string]any{
					"respectSchemaVersion": true,
					"packageReferences": map[string]string{
						"Pulumi": "3.*",
					},
				},
				"go": map[string]any{
					"respectSchemaVersion":           true,
					"generateResourceContainerTypes": true,
					"importBasePath":                 "github.com/oss4u/pulumi-opnsense/sdk/go/oensense",
				},
				"nodejs": map[string]any{
					"respectSchemaVersion": true,
					"packageName":          "@oss4u/opnsense",
				},
				"python": map[string]any{
					"respectSchemaVersion": true,
					"pyproject": map[string]bool{
						"enabled": true,
					},
				},
			},
		},
		Resources: []infer.InferredResource{
			infer.Resource[unbound.HostAliasOverride, unbound.HostAliasOverrideArgs, unbound.HostAliasOverrideState](),
			infer.Resource[unbound.HostOverride, unbound.HostOverrideArgs, unbound.HostOverrideState](),
		},
		Components: []infer.InferredComponent{
			infer.Component[*RandomComponent, RandomComponentArgs, *RandomComponentState](),
		},
		Config: infer.Config[config.Config](),
		ModuleMap: map[tokens.ModuleName]tokens.ModuleName{
			"provider": "index",
		},
	})
}

// Define some provider-level configuration
type Config struct {
	Scream *bool `pulumi:"itsasecret,optional"`
}
