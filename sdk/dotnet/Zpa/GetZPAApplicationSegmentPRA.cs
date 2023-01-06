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
    public static class GetZPAApplicationSegmentPRA
    {
        /// <summary>
        /// Use the **zpa_application_segment_pra** data source to get information about an application segment for Privileged Remote Access in the Zscaler Private Access cloud. This resource can then be referenced in an access policy rule, access policy timeout rule, access policy client forwarding rule and inspection policy. This resource supports Privileged Remote Access for both `RDP` and `SSH`.
        /// 
        /// {{% examples %}}
        /// ## Example Usage
        /// {{% example %}}
        /// 
        /// ```csharp
        /// using System.Collections.Generic;
        /// using Pulumi;
        /// using Zpa = Pulumi.Zpa;
        /// 
        /// return await Deployment.RunAsync(() =&gt; 
        /// {
        ///     var @this = Zpa.GetZPAApplicationSegmentPRA.Invoke(new()
        ///     {
        ///         Name = "PRA_Example",
        ///     });
        /// 
        /// });
        /// ```
        /// 
        /// ```csharp
        /// using System.Collections.Generic;
        /// using Pulumi;
        /// using Zpa = Pulumi.Zpa;
        /// 
        /// return await Deployment.RunAsync(() =&gt; 
        /// {
        ///     var @this = Zpa.GetZPAApplicationSegmentPRA.Invoke(new()
        ///     {
        ///         Id = "123456789",
        ///     });
        /// 
        /// });
        /// ```
        /// {{% /example %}}
        /// {{% /examples %}}
        /// </summary>
        public static Task<GetZPAApplicationSegmentPRAResult> InvokeAsync(GetZPAApplicationSegmentPRAArgs? args = null, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.InvokeAsync<GetZPAApplicationSegmentPRAResult>("zpa:index/getZPAApplicationSegmentPRA:getZPAApplicationSegmentPRA", args ?? new GetZPAApplicationSegmentPRAArgs(), options.WithDefaults());

        /// <summary>
        /// Use the **zpa_application_segment_pra** data source to get information about an application segment for Privileged Remote Access in the Zscaler Private Access cloud. This resource can then be referenced in an access policy rule, access policy timeout rule, access policy client forwarding rule and inspection policy. This resource supports Privileged Remote Access for both `RDP` and `SSH`.
        /// 
        /// {{% examples %}}
        /// ## Example Usage
        /// {{% example %}}
        /// 
        /// ```csharp
        /// using System.Collections.Generic;
        /// using Pulumi;
        /// using Zpa = Pulumi.Zpa;
        /// 
        /// return await Deployment.RunAsync(() =&gt; 
        /// {
        ///     var @this = Zpa.GetZPAApplicationSegmentPRA.Invoke(new()
        ///     {
        ///         Name = "PRA_Example",
        ///     });
        /// 
        /// });
        /// ```
        /// 
        /// ```csharp
        /// using System.Collections.Generic;
        /// using Pulumi;
        /// using Zpa = Pulumi.Zpa;
        /// 
        /// return await Deployment.RunAsync(() =&gt; 
        /// {
        ///     var @this = Zpa.GetZPAApplicationSegmentPRA.Invoke(new()
        ///     {
        ///         Id = "123456789",
        ///     });
        /// 
        /// });
        /// ```
        /// {{% /example %}}
        /// {{% /examples %}}
        /// </summary>
        public static Output<GetZPAApplicationSegmentPRAResult> Invoke(GetZPAApplicationSegmentPRAInvokeArgs? args = null, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.Invoke<GetZPAApplicationSegmentPRAResult>("zpa:index/getZPAApplicationSegmentPRA:getZPAApplicationSegmentPRA", args ?? new GetZPAApplicationSegmentPRAInvokeArgs(), options.WithDefaults());
    }


    public sealed class GetZPAApplicationSegmentPRAArgs : global::Pulumi.InvokeArgs
    {
        [Input("id")]
        public string? Id { get; set; }

        /// <summary>
        /// The name of the PRA Application Segment to be exported.
        /// </summary>
        [Input("name")]
        public string? Name { get; set; }

        [Input("tcpPortRanges")]
        private List<Inputs.GetZPAApplicationSegmentPRATcpPortRangeArgs>? _tcpPortRanges;

        /// <summary>
        /// (string) TCP port ranges used to access the app.
        /// * `from:`
        /// * `to:`
        /// </summary>
        public List<Inputs.GetZPAApplicationSegmentPRATcpPortRangeArgs> TcpPortRanges
        {
            get => _tcpPortRanges ?? (_tcpPortRanges = new List<Inputs.GetZPAApplicationSegmentPRATcpPortRangeArgs>());
            set => _tcpPortRanges = value;
        }

        [Input("udpPortRanges")]
        private List<Inputs.GetZPAApplicationSegmentPRAUdpPortRangeArgs>? _udpPortRanges;

        /// <summary>
        /// (string) UDP port ranges used to access the app.
        /// * `from:`
        /// * `to:`
        /// </summary>
        public List<Inputs.GetZPAApplicationSegmentPRAUdpPortRangeArgs> UdpPortRanges
        {
            get => _udpPortRanges ?? (_udpPortRanges = new List<Inputs.GetZPAApplicationSegmentPRAUdpPortRangeArgs>());
            set => _udpPortRanges = value;
        }

        public GetZPAApplicationSegmentPRAArgs()
        {
        }
        public static new GetZPAApplicationSegmentPRAArgs Empty => new GetZPAApplicationSegmentPRAArgs();
    }

    public sealed class GetZPAApplicationSegmentPRAInvokeArgs : global::Pulumi.InvokeArgs
    {
        [Input("id")]
        public Input<string>? Id { get; set; }

        /// <summary>
        /// The name of the PRA Application Segment to be exported.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        [Input("tcpPortRanges")]
        private InputList<Inputs.GetZPAApplicationSegmentPRATcpPortRangeInputArgs>? _tcpPortRanges;

        /// <summary>
        /// (string) TCP port ranges used to access the app.
        /// * `from:`
        /// * `to:`
        /// </summary>
        public InputList<Inputs.GetZPAApplicationSegmentPRATcpPortRangeInputArgs> TcpPortRanges
        {
            get => _tcpPortRanges ?? (_tcpPortRanges = new InputList<Inputs.GetZPAApplicationSegmentPRATcpPortRangeInputArgs>());
            set => _tcpPortRanges = value;
        }

        [Input("udpPortRanges")]
        private InputList<Inputs.GetZPAApplicationSegmentPRAUdpPortRangeInputArgs>? _udpPortRanges;

        /// <summary>
        /// (string) UDP port ranges used to access the app.
        /// * `from:`
        /// * `to:`
        /// </summary>
        public InputList<Inputs.GetZPAApplicationSegmentPRAUdpPortRangeInputArgs> UdpPortRanges
        {
            get => _udpPortRanges ?? (_udpPortRanges = new InputList<Inputs.GetZPAApplicationSegmentPRAUdpPortRangeInputArgs>());
            set => _udpPortRanges = value;
        }

        public GetZPAApplicationSegmentPRAInvokeArgs()
        {
        }
        public static new GetZPAApplicationSegmentPRAInvokeArgs Empty => new GetZPAApplicationSegmentPRAInvokeArgs();
    }


    [OutputType]
    public sealed class GetZPAApplicationSegmentPRAResult
    {
        /// <summary>
        /// (string) Indicates whether users can bypass ZPA to access applications.
        /// </summary>
        public readonly string BypassType;
        /// <summary>
        /// (string)
        /// </summary>
        public readonly string ConfigSpace;
        /// <summary>
        /// (string) Description of the application.
        /// </summary>
        public readonly string Description;
        /// <summary>
        /// (string) List of domains and IPs.
        /// </summary>
        public readonly ImmutableArray<string> DomainNames;
        /// <summary>
        /// (bool) Whether Double Encryption is enabled or disabled for the app.
        /// </summary>
        public readonly bool DoubleEncrypt;
        /// <summary>
        /// (bool) Whether this application is enabled or not
        /// </summary>
        public readonly bool Enabled;
        /// <summary>
        /// (string)
        /// </summary>
        public readonly string HealthCheckType;
        /// <summary>
        /// (string) Whether health reporting for the app is Continuous or On Access. Supported values: NONE, ON_ACCESS, CONTINUOUS.
        /// </summary>
        public readonly string HealthReporting;
        public readonly string? Id;
        /// <summary>
        /// (bool)
        /// </summary>
        public readonly bool IpAnchored;
        /// <summary>
        /// (bool) Indicates if the Zscaler Client Connector (formerly Zscaler App or Z App) receives CNAME DNS records from the connectors.
        /// </summary>
        public readonly bool IsCnameEnabled;
        public readonly string? Name;
        /// <summary>
        /// (bool)
        /// </summary>
        public readonly bool PassiveHealthEnabled;
        /// <summary>
        /// (String) Segment Group IDs
        /// </summary>
        public readonly string SegmentGroupId;
        public readonly string SegmentGroupName;
        /// <summary>
        /// (string) List of Server Group IDs
        /// * `id:` - (string) List of Server Group IDs
        /// </summary>
        public readonly ImmutableArray<Outputs.GetZPAApplicationSegmentPRAServerGroupResult> ServerGroups;
        /// <summary>
        /// (string) TCP port ranges used to access the app.
        /// * `app_id:` - (string)
        /// * `name:` - (string) Name of the Privileged Remote Access
        /// * `description:` - (string) Description of the Privileged Remote Access
        /// * `domain:` - (string) Domain name of the Privileged Remote Access
        /// </summary>
        public readonly ImmutableArray<Outputs.GetZPAApplicationSegmentPRASraAppResult> SraApps;
        /// <summary>
        /// (string) TCP port ranges used to access the app.
        /// </summary>
        public readonly ImmutableArray<string> TcpPortRanges;
        /// <summary>
        /// (string) UDP port ranges used to access the app.
        /// </summary>
        public readonly ImmutableArray<string> UdpPortRanges;

        [OutputConstructor]
        private GetZPAApplicationSegmentPRAResult(
            string bypassType,

            string configSpace,

            string description,

            ImmutableArray<string> domainNames,

            bool doubleEncrypt,

            bool enabled,

            string healthCheckType,

            string healthReporting,

            string? id,

            bool ipAnchored,

            bool isCnameEnabled,

            string? name,

            bool passiveHealthEnabled,

            string segmentGroupId,

            string segmentGroupName,

            ImmutableArray<Outputs.GetZPAApplicationSegmentPRAServerGroupResult> serverGroups,

            ImmutableArray<Outputs.GetZPAApplicationSegmentPRASraAppResult> sraApps,

            ImmutableArray<string> tcpPortRanges,

            ImmutableArray<string> udpPortRanges)
        {
            BypassType = bypassType;
            ConfigSpace = configSpace;
            Description = description;
            DomainNames = domainNames;
            DoubleEncrypt = doubleEncrypt;
            Enabled = enabled;
            HealthCheckType = healthCheckType;
            HealthReporting = healthReporting;
            Id = id;
            IpAnchored = ipAnchored;
            IsCnameEnabled = isCnameEnabled;
            Name = name;
            PassiveHealthEnabled = passiveHealthEnabled;
            SegmentGroupId = segmentGroupId;
            SegmentGroupName = segmentGroupName;
            ServerGroups = serverGroups;
            SraApps = sraApps;
            TcpPortRanges = tcpPortRanges;
            UdpPortRanges = udpPortRanges;
        }
    }
}
