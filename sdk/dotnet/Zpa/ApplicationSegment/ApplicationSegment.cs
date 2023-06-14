// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;
using Pulumi;

namespace zscaler.PulumiPackage.Zpa.ApplicationSegment
{
    /// <summary>
    /// ## Import
    /// 
    /// Zscaler offers a dedicated tool called Zscaler-Terraformer to allow the automated import of ZPA configurations into Terraform-compliant HashiCorp Configuration Language. [Visit](https://github.com/zscaler/zscaler-terraformer) Application Segment can be imported by using `&lt;APPLICATION SEGMENT ID&gt;` or `&lt;APPLICATION SEGMENT NAME&gt;` as the import ID.
    /// 
    /// ```sh
    ///  $ pulumi import zpa:ApplicationSegment/applicationSegment:ApplicationSegment example &lt;application_segment_id&gt;
    /// ```
    /// 
    ///  or
    /// 
    /// ```sh
    ///  $ pulumi import zpa:ApplicationSegment/applicationSegment:ApplicationSegment example &lt;application_segment_name&gt;
    /// ```
    /// </summary>
    [ZpaResourceType("zpa:ApplicationSegment/applicationSegment:ApplicationSegment")]
    public partial class ApplicationSegment : global::Pulumi.CustomResource
    {
        /// <summary>
        /// (Optional) Indicates whether users can bypass ZPA to access applications.
        /// </summary>
        [Output("bypassType")]
        public Output<string> BypassType { get; private set; } = null!;

        /// <summary>
        /// (Optional)
        /// </summary>
        [Output("configSpace")]
        public Output<string?> ConfigSpace { get; private set; } = null!;

        [Output("defaultIdleTimeout")]
        public Output<string> DefaultIdleTimeout { get; private set; } = null!;

        /// <summary>
        /// (Optional) Description of the application.
        /// </summary>
        [Output("description")]
        public Output<string?> Description { get; private set; } = null!;

        /// <summary>
        /// List of domains and IPs.
        /// </summary>
        [Output("domainNames")]
        public Output<ImmutableArray<string>> DomainNames { get; private set; } = null!;

        /// <summary>
        /// (Optional) Whether Double Encryption is enabled or disabled for the app.
        /// </summary>
        [Output("doubleEncrypt")]
        public Output<bool?> DoubleEncrypt { get; private set; } = null!;

        /// <summary>
        /// (Optional) Whether this application is enabled or not.
        /// </summary>
        [Output("enabled")]
        public Output<bool?> Enabled { get; private set; } = null!;

        /// <summary>
        /// (Optional)
        /// </summary>
        [Output("healthCheckType")]
        public Output<string?> HealthCheckType { get; private set; } = null!;

        /// <summary>
        /// (Optional) Whether health reporting for the app is Continuous or On Access. Supported values: `NONE`, `ON_ACCESS`, `CONTINUOUS`.
        /// </summary>
        [Output("healthReporting")]
        public Output<string?> HealthReporting { get; private set; } = null!;

        /// <summary>
        /// (Optional)
        /// </summary>
        [Output("icmpAccessType")]
        public Output<string?> IcmpAccessType { get; private set; } = null!;

        /// <summary>
        /// (Optional)
        /// </summary>
        [Output("ipAnchored")]
        public Output<bool?> IpAnchored { get; private set; } = null!;

        /// <summary>
        /// (Optional) Indicates if the Zscaler Client Connector (formerly Zscaler App or Z App) receives CNAME DNS records from the connectors.
        /// </summary>
        [Output("isCnameEnabled")]
        public Output<bool> IsCnameEnabled { get; private set; } = null!;

        /// <summary>
        /// Name. The name of the App Connector Group to be exported.
        /// </summary>
        [Output("name")]
        public Output<string> Name { get; private set; } = null!;

        /// <summary>
        /// (Optional)
        /// </summary>
        [Output("passiveHealthEnabled")]
        public Output<bool> PassiveHealthEnabled { get; private set; } = null!;

        /// <summary>
        /// List of Segment Group IDs
        /// </summary>
        [Output("segmentGroupId")]
        public Output<string> SegmentGroupId { get; private set; } = null!;

        [Output("segmentGroupName")]
        public Output<string> SegmentGroupName { get; private set; } = null!;

        /// <summary>
        /// List of Server Group IDs
        /// </summary>
        [Output("serverGroups")]
        public Output<ImmutableArray<Outputs.ApplicationSegmentServerGroup>> ServerGroups { get; private set; } = null!;

        /// <summary>
        /// TCP port ranges used to access the app.
        /// </summary>
        [Output("tcpPortRanges")]
        public Output<ImmutableArray<string>> TcpPortRanges { get; private set; } = null!;

        /// <summary>
        /// UDP port ranges used to access the app.
        /// </summary>
        [Output("udpPortRanges")]
        public Output<ImmutableArray<string>> UdpPortRanges { get; private set; } = null!;


        /// <summary>
        /// Create a ApplicationSegment resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public ApplicationSegment(string name, ApplicationSegmentArgs args, CustomResourceOptions? options = null)
            : base("zpa:ApplicationSegment/applicationSegment:ApplicationSegment", name, args ?? new ApplicationSegmentArgs(), MakeResourceOptions(options, ""))
        {
        }

        private ApplicationSegment(string name, Input<string> id, ApplicationSegmentState? state = null, CustomResourceOptions? options = null)
            : base("zpa:ApplicationSegment/applicationSegment:ApplicationSegment", name, state, MakeResourceOptions(options, id))
        {
        }

        private static CustomResourceOptions MakeResourceOptions(CustomResourceOptions? options, Input<string>? id)
        {
            var defaultOptions = new CustomResourceOptions
            {
                Version = Utilities.Version,
                PluginDownloadURL = "github://api.github.com/zscaler",
            };
            var merged = CustomResourceOptions.Merge(defaultOptions, options);
            // Override the ID if one was specified for consistency with other language SDKs.
            merged.Id = id ?? merged.Id;
            return merged;
        }
        /// <summary>
        /// Get an existing ApplicationSegment resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static ApplicationSegment Get(string name, Input<string> id, ApplicationSegmentState? state = null, CustomResourceOptions? options = null)
        {
            return new ApplicationSegment(name, id, state, options);
        }
    }

    public sealed class ApplicationSegmentArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// (Optional) Indicates whether users can bypass ZPA to access applications.
        /// </summary>
        [Input("bypassType")]
        public Input<string>? BypassType { get; set; }

        /// <summary>
        /// (Optional)
        /// </summary>
        [Input("configSpace")]
        public Input<string>? ConfigSpace { get; set; }

        [Input("defaultIdleTimeout")]
        public Input<string>? DefaultIdleTimeout { get; set; }

        /// <summary>
        /// (Optional) Description of the application.
        /// </summary>
        [Input("description")]
        public Input<string>? Description { get; set; }

        [Input("domainNames", required: true)]
        private InputList<string>? _domainNames;

        /// <summary>
        /// List of domains and IPs.
        /// </summary>
        public InputList<string> DomainNames
        {
            get => _domainNames ?? (_domainNames = new InputList<string>());
            set => _domainNames = value;
        }

        /// <summary>
        /// (Optional) Whether Double Encryption is enabled or disabled for the app.
        /// </summary>
        [Input("doubleEncrypt")]
        public Input<bool>? DoubleEncrypt { get; set; }

        /// <summary>
        /// (Optional) Whether this application is enabled or not.
        /// </summary>
        [Input("enabled")]
        public Input<bool>? Enabled { get; set; }

        /// <summary>
        /// (Optional)
        /// </summary>
        [Input("healthCheckType")]
        public Input<string>? HealthCheckType { get; set; }

        /// <summary>
        /// (Optional) Whether health reporting for the app is Continuous or On Access. Supported values: `NONE`, `ON_ACCESS`, `CONTINUOUS`.
        /// </summary>
        [Input("healthReporting")]
        public Input<string>? HealthReporting { get; set; }

        /// <summary>
        /// (Optional)
        /// </summary>
        [Input("icmpAccessType")]
        public Input<string>? IcmpAccessType { get; set; }

        /// <summary>
        /// (Optional)
        /// </summary>
        [Input("ipAnchored")]
        public Input<bool>? IpAnchored { get; set; }

        /// <summary>
        /// (Optional) Indicates if the Zscaler Client Connector (formerly Zscaler App or Z App) receives CNAME DNS records from the connectors.
        /// </summary>
        [Input("isCnameEnabled")]
        public Input<bool>? IsCnameEnabled { get; set; }

        /// <summary>
        /// Name. The name of the App Connector Group to be exported.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// (Optional)
        /// </summary>
        [Input("passiveHealthEnabled")]
        public Input<bool>? PassiveHealthEnabled { get; set; }

        /// <summary>
        /// List of Segment Group IDs
        /// </summary>
        [Input("segmentGroupId")]
        public Input<string>? SegmentGroupId { get; set; }

        [Input("segmentGroupName")]
        public Input<string>? SegmentGroupName { get; set; }

        [Input("serverGroups", required: true)]
        private InputList<Inputs.ApplicationSegmentServerGroupArgs>? _serverGroups;

        /// <summary>
        /// List of Server Group IDs
        /// </summary>
        public InputList<Inputs.ApplicationSegmentServerGroupArgs> ServerGroups
        {
            get => _serverGroups ?? (_serverGroups = new InputList<Inputs.ApplicationSegmentServerGroupArgs>());
            set => _serverGroups = value;
        }

        [Input("tcpPortRanges")]
        private InputList<string>? _tcpPortRanges;

        /// <summary>
        /// TCP port ranges used to access the app.
        /// </summary>
        public InputList<string> TcpPortRanges
        {
            get => _tcpPortRanges ?? (_tcpPortRanges = new InputList<string>());
            set => _tcpPortRanges = value;
        }

        [Input("udpPortRanges")]
        private InputList<string>? _udpPortRanges;

        /// <summary>
        /// UDP port ranges used to access the app.
        /// </summary>
        public InputList<string> UdpPortRanges
        {
            get => _udpPortRanges ?? (_udpPortRanges = new InputList<string>());
            set => _udpPortRanges = value;
        }

        public ApplicationSegmentArgs()
        {
        }
        public static new ApplicationSegmentArgs Empty => new ApplicationSegmentArgs();
    }

    public sealed class ApplicationSegmentState : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// (Optional) Indicates whether users can bypass ZPA to access applications.
        /// </summary>
        [Input("bypassType")]
        public Input<string>? BypassType { get; set; }

        /// <summary>
        /// (Optional)
        /// </summary>
        [Input("configSpace")]
        public Input<string>? ConfigSpace { get; set; }

        [Input("defaultIdleTimeout")]
        public Input<string>? DefaultIdleTimeout { get; set; }

        /// <summary>
        /// (Optional) Description of the application.
        /// </summary>
        [Input("description")]
        public Input<string>? Description { get; set; }

        [Input("domainNames")]
        private InputList<string>? _domainNames;

        /// <summary>
        /// List of domains and IPs.
        /// </summary>
        public InputList<string> DomainNames
        {
            get => _domainNames ?? (_domainNames = new InputList<string>());
            set => _domainNames = value;
        }

        /// <summary>
        /// (Optional) Whether Double Encryption is enabled or disabled for the app.
        /// </summary>
        [Input("doubleEncrypt")]
        public Input<bool>? DoubleEncrypt { get; set; }

        /// <summary>
        /// (Optional) Whether this application is enabled or not.
        /// </summary>
        [Input("enabled")]
        public Input<bool>? Enabled { get; set; }

        /// <summary>
        /// (Optional)
        /// </summary>
        [Input("healthCheckType")]
        public Input<string>? HealthCheckType { get; set; }

        /// <summary>
        /// (Optional) Whether health reporting for the app is Continuous or On Access. Supported values: `NONE`, `ON_ACCESS`, `CONTINUOUS`.
        /// </summary>
        [Input("healthReporting")]
        public Input<string>? HealthReporting { get; set; }

        /// <summary>
        /// (Optional)
        /// </summary>
        [Input("icmpAccessType")]
        public Input<string>? IcmpAccessType { get; set; }

        /// <summary>
        /// (Optional)
        /// </summary>
        [Input("ipAnchored")]
        public Input<bool>? IpAnchored { get; set; }

        /// <summary>
        /// (Optional) Indicates if the Zscaler Client Connector (formerly Zscaler App or Z App) receives CNAME DNS records from the connectors.
        /// </summary>
        [Input("isCnameEnabled")]
        public Input<bool>? IsCnameEnabled { get; set; }

        /// <summary>
        /// Name. The name of the App Connector Group to be exported.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// (Optional)
        /// </summary>
        [Input("passiveHealthEnabled")]
        public Input<bool>? PassiveHealthEnabled { get; set; }

        /// <summary>
        /// List of Segment Group IDs
        /// </summary>
        [Input("segmentGroupId")]
        public Input<string>? SegmentGroupId { get; set; }

        [Input("segmentGroupName")]
        public Input<string>? SegmentGroupName { get; set; }

        [Input("serverGroups")]
        private InputList<Inputs.ApplicationSegmentServerGroupGetArgs>? _serverGroups;

        /// <summary>
        /// List of Server Group IDs
        /// </summary>
        public InputList<Inputs.ApplicationSegmentServerGroupGetArgs> ServerGroups
        {
            get => _serverGroups ?? (_serverGroups = new InputList<Inputs.ApplicationSegmentServerGroupGetArgs>());
            set => _serverGroups = value;
        }

        [Input("tcpPortRanges")]
        private InputList<string>? _tcpPortRanges;

        /// <summary>
        /// TCP port ranges used to access the app.
        /// </summary>
        public InputList<string> TcpPortRanges
        {
            get => _tcpPortRanges ?? (_tcpPortRanges = new InputList<string>());
            set => _tcpPortRanges = value;
        }

        [Input("udpPortRanges")]
        private InputList<string>? _udpPortRanges;

        /// <summary>
        /// UDP port ranges used to access the app.
        /// </summary>
        public InputList<string> UdpPortRanges
        {
            get => _udpPortRanges ?? (_udpPortRanges = new InputList<string>());
            set => _udpPortRanges = value;
        }

        public ApplicationSegmentState()
        {
        }
        public static new ApplicationSegmentState Empty => new ApplicationSegmentState();
    }
}