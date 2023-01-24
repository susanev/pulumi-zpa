// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;
using Pulumi;

namespace zscaler.PulumiPackage.Zpa
{
    /// <summary>
    /// The provider type for the zpa package. By default, resources use package-wide configuration
    /// settings, however an explicit `Provider` instance may be created and passed during resource
    /// construction to achieve fine-grained programmatic control over provider settings. See the
    /// [documentation](https://www.pulumi.com/docs/reference/programming-model/#providers) for more information.
    /// </summary>
    [ZpaResourceType("pulumi:providers:zpa")]
    public partial class Provider : global::Pulumi.ProviderResource
    {
        /// <summary>
        /// zpa client id
        /// </summary>
        [Output("zpaClientId")]
        public Output<string?> ZpaClientId { get; private set; } = null!;

        /// <summary>
        /// zpa client secret
        /// </summary>
        [Output("zpaClientSecret")]
        public Output<string?> ZpaClientSecret { get; private set; } = null!;

        /// <summary>
        /// Cloud to use PRODUCTION, BETA, GOV or PREVIEW
        /// </summary>
        [Output("zpaCloud")]
        public Output<string?> ZpaCloud { get; private set; } = null!;

        /// <summary>
        /// zpa customer id
        /// </summary>
        [Output("zpaCustomerId")]
        public Output<string?> ZpaCustomerId { get; private set; } = null!;


        /// <summary>
        /// Create a Provider resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public Provider(string name, ProviderArgs? args = null, CustomResourceOptions? options = null)
            : base("zpa", name, args ?? new ProviderArgs(), MakeResourceOptions(options, ""))
        {
        }

        private static CustomResourceOptions MakeResourceOptions(CustomResourceOptions? options, Input<string>? id)
        {
            var defaultOptions = new CustomResourceOptions
            {
                Version = Utilities.Version,
                PluginDownloadURL = "github://api.github.com/zscaler",
                AdditionalSecretOutputs =
                {
                    "zpaClientSecret",
                    "zpaCustomerId",
                },
            };
            var merged = CustomResourceOptions.Merge(defaultOptions, options);
            // Override the ID if one was specified for consistency with other language SDKs.
            merged.Id = id ?? merged.Id;
            return merged;
        }
    }

    public sealed class ProviderArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// zpa client id
        /// </summary>
        [Input("zpaClientId")]
        public Input<string>? ZpaClientId { get; set; }

        [Input("zpaClientSecret")]
        private Input<string>? _zpaClientSecret;

        /// <summary>
        /// zpa client secret
        /// </summary>
        public Input<string>? ZpaClientSecret
        {
            get => _zpaClientSecret;
            set
            {
                var emptySecret = Output.CreateSecret(0);
                _zpaClientSecret = Output.Tuple<Input<string>?, int>(value, emptySecret).Apply(t => t.Item1);
            }
        }

        /// <summary>
        /// Cloud to use PRODUCTION, BETA, GOV or PREVIEW
        /// </summary>
        [Input("zpaCloud")]
        public Input<string>? ZpaCloud { get; set; }

        [Input("zpaCustomerId")]
        private Input<string>? _zpaCustomerId;

        /// <summary>
        /// zpa customer id
        /// </summary>
        public Input<string>? ZpaCustomerId
        {
            get => _zpaCustomerId;
            set
            {
                var emptySecret = Output.CreateSecret(0);
                _zpaCustomerId = Output.Tuple<Input<string>?, int>(value, emptySecret).Apply(t => t.Item1);
            }
        }

        public ProviderArgs()
        {
            ZpaClientId = Utilities.GetEnv("ZPA_CLIENT_ID");
            ZpaClientSecret = Utilities.GetEnv("ZPA_CLIENT_SECRET");
            ZpaCloud = Utilities.GetEnv("ZPA_CLOUD");
            ZpaCustomerId = Utilities.GetEnv("ZPA_CUSTOMER_ID");
        }
        public static new ProviderArgs Empty => new ProviderArgs();
    }
}
