// *** WARNING: this file was generated by pulumi. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;
using Pulumi;

namespace Oss4u.Opnsense.Unbound
{
    [OpnsenseResourceType("opnsense:unbound:HostOverride")]
    public partial class HostOverride : global::Pulumi.CustomResource
    {
        [Output("description")]
        public Output<string> Description { get; private set; } = null!;

        [Output("domain")]
        public Output<string> Domain { get; private set; } = null!;

        [Output("enabled")]
        public Output<bool> Enabled { get; private set; } = null!;

        [Output("hostname")]
        public Output<string> Hostname { get; private set; } = null!;

        [Output("mx")]
        public Output<string?> Mx { get; private set; } = null!;

        [Output("mx_prio")]
        public Output<int?> Mx_prio { get; private set; } = null!;

        [Output("result")]
        public Output<string> Result { get; private set; } = null!;

        [Output("rr")]
        public Output<string> Rr { get; private set; } = null!;

        [Output("server")]
        public Output<string?> Server { get; private set; } = null!;


        /// <summary>
        /// Create a HostOverride resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public HostOverride(string name, HostOverrideArgs args, CustomResourceOptions? options = null)
            : base("opnsense:unbound:HostOverride", name, args ?? new HostOverrideArgs(), MakeResourceOptions(options, ""))
        {
        }

        private HostOverride(string name, Input<string> id, CustomResourceOptions? options = null)
            : base("opnsense:unbound:HostOverride", name, null, MakeResourceOptions(options, id))
        {
        }

        private static CustomResourceOptions MakeResourceOptions(CustomResourceOptions? options, Input<string>? id)
        {
            var defaultOptions = new CustomResourceOptions
            {
                Version = Utilities.Version,
                PluginDownloadURL = "github://api.github.com/oss4u/pulumi-opnsense-native",
            };
            var merged = CustomResourceOptions.Merge(defaultOptions, options);
            // Override the ID if one was specified for consistency with other language SDKs.
            merged.Id = id ?? merged.Id;
            return merged;
        }
        /// <summary>
        /// Get an existing HostOverride resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static HostOverride Get(string name, Input<string> id, CustomResourceOptions? options = null)
        {
            return new HostOverride(name, id, options);
        }
    }

    public sealed class HostOverrideArgs : global::Pulumi.ResourceArgs
    {
        [Input("description", required: true)]
        public Input<string> Description { get; set; } = null!;

        [Input("domain", required: true)]
        public Input<string> Domain { get; set; } = null!;

        [Input("enabled", required: true)]
        public Input<bool> Enabled { get; set; } = null!;

        [Input("hostname", required: true)]
        public Input<string> Hostname { get; set; } = null!;

        [Input("mx")]
        public Input<string>? Mx { get; set; }

        [Input("mx_prio")]
        public Input<int>? Mx_prio { get; set; }

        [Input("rr", required: true)]
        public Input<string> Rr { get; set; } = null!;

        [Input("server")]
        public Input<string>? Server { get; set; }

        public HostOverrideArgs()
        {
        }
        public static new HostOverrideArgs Empty => new HostOverrideArgs();
    }
}
