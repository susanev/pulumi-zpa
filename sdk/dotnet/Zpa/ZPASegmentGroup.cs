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
    /// ## Example Usage
    /// 
    /// ```csharp
    /// using System.Collections.Generic;
    /// using Pulumi;
    /// using Zpa = zscaler.PulumiPackage.Zpa;
    /// 
    /// return await Deployment.RunAsync(() =&gt; 
    /// {
    ///     // ZPA Segment Group resource
    ///     var testSegmentGroup = new Zpa.ZPASegmentGroup("testSegmentGroup", new()
    ///     {
    ///         Description = "test1-segment-group",
    ///         Enabled = true,
    ///         TcpKeepAliveEnabled = "1",
    ///     });
    /// 
    /// });
    /// ```
    /// 
    /// ## Import
    /// 
    /// Zscaler offers a dedicated tool called Zscaler-Terraformer to allow the automated import of ZPA configurations into Terraform-compliant HashiCorp Configuration Language. [Visit](https://github.com/zscaler/zscaler-terraformer) **segment_group** can be imported by using `&lt;SEGMENT GROUP ID&gt;` or `&lt;SEGMENT GROUP NAME&gt;` as the import ID. For example
    /// 
    /// ```sh
    ///  $ pulumi import zpa:index/zPASegmentGroup:ZPASegmentGroup example &lt;segment_group_id&gt;
    /// ```
    /// 
    ///  or
    /// 
    /// ```sh
    ///  $ pulumi import zpa:index/zPASegmentGroup:ZPASegmentGroup example &lt;segment_group_name&gt;
    /// ```
    /// </summary>
    [ZpaResourceType("zpa:index/zPASegmentGroup:ZPASegmentGroup")]
    public partial class ZPASegmentGroup : global::Pulumi.CustomResource
    {
        [Output("applications")]
        public Output<ImmutableArray<Outputs.ZPASegmentGroupApplication>> Applications { get; private set; } = null!;

        /// <summary>
        /// (Optional)
        /// </summary>
        [Output("configSpace")]
        public Output<string?> ConfigSpace { get; private set; } = null!;

        /// <summary>
        /// (Optional) Description of the segment group.
        /// </summary>
        [Output("description")]
        public Output<string?> Description { get; private set; } = null!;

        /// <summary>
        /// (Optional) Whether this segment group is enabled or not.
        /// </summary>
        [Output("enabled")]
        public Output<bool?> Enabled { get; private set; } = null!;

        /// <summary>
        /// (Required) Name of the segment group.
        /// </summary>
        [Output("name")]
        public Output<string> Name { get; private set; } = null!;

        [Output("policyMigrated")]
        public Output<bool?> PolicyMigrated { get; private set; } = null!;

        /// <summary>
        /// (Optional)
        /// </summary>
        [Output("tcpKeepAliveEnabled")]
        public Output<string?> TcpKeepAliveEnabled { get; private set; } = null!;


        /// <summary>
        /// Create a ZPASegmentGroup resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public ZPASegmentGroup(string name, ZPASegmentGroupArgs? args = null, CustomResourceOptions? options = null)
            : base("zpa:index/zPASegmentGroup:ZPASegmentGroup", name, args ?? new ZPASegmentGroupArgs(), MakeResourceOptions(options, ""))
        {
        }

        private ZPASegmentGroup(string name, Input<string> id, ZPASegmentGroupState? state = null, CustomResourceOptions? options = null)
            : base("zpa:index/zPASegmentGroup:ZPASegmentGroup", name, state, MakeResourceOptions(options, id))
        {
        }

        private static CustomResourceOptions MakeResourceOptions(CustomResourceOptions? options, Input<string>? id)
        {
            var defaultOptions = new CustomResourceOptions
            {
                Version = Utilities.Version,
            };
            var merged = CustomResourceOptions.Merge(defaultOptions, options);
            // Override the ID if one was specified for consistency with other language SDKs.
            merged.Id = id ?? merged.Id;
            return merged;
        }
        /// <summary>
        /// Get an existing ZPASegmentGroup resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static ZPASegmentGroup Get(string name, Input<string> id, ZPASegmentGroupState? state = null, CustomResourceOptions? options = null)
        {
            return new ZPASegmentGroup(name, id, state, options);
        }
    }

    public sealed class ZPASegmentGroupArgs : global::Pulumi.ResourceArgs
    {
        [Input("applications")]
        private InputList<Inputs.ZPASegmentGroupApplicationArgs>? _applications;
        public InputList<Inputs.ZPASegmentGroupApplicationArgs> Applications
        {
            get => _applications ?? (_applications = new InputList<Inputs.ZPASegmentGroupApplicationArgs>());
            set => _applications = value;
        }

        /// <summary>
        /// (Optional)
        /// </summary>
        [Input("configSpace")]
        public Input<string>? ConfigSpace { get; set; }

        /// <summary>
        /// (Optional) Description of the segment group.
        /// </summary>
        [Input("description")]
        public Input<string>? Description { get; set; }

        /// <summary>
        /// (Optional) Whether this segment group is enabled or not.
        /// </summary>
        [Input("enabled")]
        public Input<bool>? Enabled { get; set; }

        /// <summary>
        /// (Required) Name of the segment group.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        [Input("policyMigrated")]
        public Input<bool>? PolicyMigrated { get; set; }

        /// <summary>
        /// (Optional)
        /// </summary>
        [Input("tcpKeepAliveEnabled")]
        public Input<string>? TcpKeepAliveEnabled { get; set; }

        public ZPASegmentGroupArgs()
        {
        }
        public static new ZPASegmentGroupArgs Empty => new ZPASegmentGroupArgs();
    }

    public sealed class ZPASegmentGroupState : global::Pulumi.ResourceArgs
    {
        [Input("applications")]
        private InputList<Inputs.ZPASegmentGroupApplicationGetArgs>? _applications;
        public InputList<Inputs.ZPASegmentGroupApplicationGetArgs> Applications
        {
            get => _applications ?? (_applications = new InputList<Inputs.ZPASegmentGroupApplicationGetArgs>());
            set => _applications = value;
        }

        /// <summary>
        /// (Optional)
        /// </summary>
        [Input("configSpace")]
        public Input<string>? ConfigSpace { get; set; }

        /// <summary>
        /// (Optional) Description of the segment group.
        /// </summary>
        [Input("description")]
        public Input<string>? Description { get; set; }

        /// <summary>
        /// (Optional) Whether this segment group is enabled or not.
        /// </summary>
        [Input("enabled")]
        public Input<bool>? Enabled { get; set; }

        /// <summary>
        /// (Required) Name of the segment group.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        [Input("policyMigrated")]
        public Input<bool>? PolicyMigrated { get; set; }

        /// <summary>
        /// (Optional)
        /// </summary>
        [Input("tcpKeepAliveEnabled")]
        public Input<string>? TcpKeepAliveEnabled { get; set; }

        public ZPASegmentGroupState()
        {
        }
        public static new ZPASegmentGroupState Empty => new ZPASegmentGroupState();
    }
}
