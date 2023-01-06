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
    public static class GetZPAApplicationSegmentInspection
    {
        /// <summary>
        /// Use the **zpa_application_segment_inspection** data source to get information about an inspection application segment in the Zscaler Private Access cloud. This resource can then be referenced in a ZPA access inspection policy. This resource supports ZPA Inspection for both `HTTP` and `HTTPS`.
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
        ///     var @this = Zpa.GetZPAApplicationSegmentInspection.Invoke(new()
        ///     {
        ///         Name = "ZPA_Inspection_Example",
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
        ///     var @this = Zpa.GetZPAApplicationSegmentInspection.Invoke(new()
        ///     {
        ///         Id = "123456789",
        ///     });
        /// 
        /// });
        /// ```
        /// {{% /example %}}
        /// {{% /examples %}}
        /// </summary>
        public static Task<GetZPAApplicationSegmentInspectionResult> InvokeAsync(GetZPAApplicationSegmentInspectionArgs? args = null, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.InvokeAsync<GetZPAApplicationSegmentInspectionResult>("zpa:index/getZPAApplicationSegmentInspection:getZPAApplicationSegmentInspection", args ?? new GetZPAApplicationSegmentInspectionArgs(), options.WithDefaults());

        /// <summary>
        /// Use the **zpa_application_segment_inspection** data source to get information about an inspection application segment in the Zscaler Private Access cloud. This resource can then be referenced in a ZPA access inspection policy. This resource supports ZPA Inspection for both `HTTP` and `HTTPS`.
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
        ///     var @this = Zpa.GetZPAApplicationSegmentInspection.Invoke(new()
        ///     {
        ///         Name = "ZPA_Inspection_Example",
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
        ///     var @this = Zpa.GetZPAApplicationSegmentInspection.Invoke(new()
        ///     {
        ///         Id = "123456789",
        ///     });
        /// 
        /// });
        /// ```
        /// {{% /example %}}
        /// {{% /examples %}}
        /// </summary>
        public static Output<GetZPAApplicationSegmentInspectionResult> Invoke(GetZPAApplicationSegmentInspectionInvokeArgs? args = null, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.Invoke<GetZPAApplicationSegmentInspectionResult>("zpa:index/getZPAApplicationSegmentInspection:getZPAApplicationSegmentInspection", args ?? new GetZPAApplicationSegmentInspectionInvokeArgs(), options.WithDefaults());
    }


    public sealed class GetZPAApplicationSegmentInspectionArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// The ID of the Inspection Application Segment to be exported.
        /// </summary>
        [Input("id")]
        public string? Id { get; set; }

        /// <summary>
        /// The name of the Inspection Application Segment to be exported.
        /// </summary>
        [Input("name")]
        public string? Name { get; set; }

        [Input("tcpPortRanges")]
        private List<Inputs.GetZPAApplicationSegmentInspectionTcpPortRangeArgs>? _tcpPortRanges;

        /// <summary>
        /// (string) TCP port ranges used to access the app.
        /// * `from:`
        /// * `to:`
        /// </summary>
        public List<Inputs.GetZPAApplicationSegmentInspectionTcpPortRangeArgs> TcpPortRanges
        {
            get => _tcpPortRanges ?? (_tcpPortRanges = new List<Inputs.GetZPAApplicationSegmentInspectionTcpPortRangeArgs>());
            set => _tcpPortRanges = value;
        }

        [Input("udpPortRanges")]
        private List<Inputs.GetZPAApplicationSegmentInspectionUdpPortRangeArgs>? _udpPortRanges;

        /// <summary>
        /// (string) UDP port ranges used to access the app.
        /// * `from:`
        /// * `to:`
        /// </summary>
        public List<Inputs.GetZPAApplicationSegmentInspectionUdpPortRangeArgs> UdpPortRanges
        {
            get => _udpPortRanges ?? (_udpPortRanges = new List<Inputs.GetZPAApplicationSegmentInspectionUdpPortRangeArgs>());
            set => _udpPortRanges = value;
        }

        public GetZPAApplicationSegmentInspectionArgs()
        {
        }
        public static new GetZPAApplicationSegmentInspectionArgs Empty => new GetZPAApplicationSegmentInspectionArgs();
    }

    public sealed class GetZPAApplicationSegmentInspectionInvokeArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// The ID of the Inspection Application Segment to be exported.
        /// </summary>
        [Input("id")]
        public Input<string>? Id { get; set; }

        /// <summary>
        /// The name of the Inspection Application Segment to be exported.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        [Input("tcpPortRanges")]
        private InputList<Inputs.GetZPAApplicationSegmentInspectionTcpPortRangeInputArgs>? _tcpPortRanges;

        /// <summary>
        /// (string) TCP port ranges used to access the app.
        /// * `from:`
        /// * `to:`
        /// </summary>
        public InputList<Inputs.GetZPAApplicationSegmentInspectionTcpPortRangeInputArgs> TcpPortRanges
        {
            get => _tcpPortRanges ?? (_tcpPortRanges = new InputList<Inputs.GetZPAApplicationSegmentInspectionTcpPortRangeInputArgs>());
            set => _tcpPortRanges = value;
        }

        [Input("udpPortRanges")]
        private InputList<Inputs.GetZPAApplicationSegmentInspectionUdpPortRangeInputArgs>? _udpPortRanges;

        /// <summary>
        /// (string) UDP port ranges used to access the app.
        /// * `from:`
        /// * `to:`
        /// </summary>
        public InputList<Inputs.GetZPAApplicationSegmentInspectionUdpPortRangeInputArgs> UdpPortRanges
        {
            get => _udpPortRanges ?? (_udpPortRanges = new InputList<Inputs.GetZPAApplicationSegmentInspectionUdpPortRangeInputArgs>());
            set => _udpPortRanges = value;
        }

        public GetZPAApplicationSegmentInspectionInvokeArgs()
        {
        }
        public static new GetZPAApplicationSegmentInspectionInvokeArgs Empty => new GetZPAApplicationSegmentInspectionInvokeArgs();
    }


    [OutputType]
    public sealed class GetZPAApplicationSegmentInspectionResult
    {
        /// <summary>
        /// (string) Indicates whether users can bypass ZPA to access applications.
        /// </summary>
        public readonly string BypassType;
        /// <summary>
        /// (String)
        /// </summary>
        public readonly string CreationTime;
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
        public readonly string HealthCheckType;
        /// <summary>
        /// (string) Whether health reporting for the app is Continuous or On Access. Supported values: `NONE`, `ON_ACCESS`, `CONTINUOUS`.
        /// * `health_check_type` (string)
        /// </summary>
        public readonly string HealthReporting;
        /// <summary>
        /// (string)
        /// </summary>
        public readonly string IcmpAccessType;
        public readonly string? Id;
        /// <summary>
        /// (string) TCP port ranges used to access the app.
        /// * `app_id:` - (string)
        /// * `name:` - (string) Name of the Inspection Application
        /// * `description:` - (string) Description of the Inspection Application
        /// * `domain:` - (string) Domain name of the inspection application
        /// </summary>
        public readonly ImmutableArray<Outputs.GetZPAApplicationSegmentInspectionInspectionAppResult> InspectionApps;
        /// <summary>
        /// (bool)
        /// </summary>
        public readonly bool IpAnchored;
        /// <summary>
        /// (bool) Indicates if the Zscaler Client Connector (formerly Zscaler App or Z App) receives CNAME DNS records from the connectors.
        /// </summary>
        public readonly bool IsCnameEnabled;
        public readonly string ModifiedBy;
        /// <summary>
        /// (String)
        /// </summary>
        public readonly string ModifiedTime;
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
        public readonly bool SelectConnectorCloseToApp;
        /// <summary>
        /// (string) List of Server Group IDs
        /// * `id:` - (string) List of Server Group IDs
        /// </summary>
        public readonly ImmutableArray<Outputs.GetZPAApplicationSegmentInspectionServerGroupResult> ServerGroups;
        /// <summary>
        /// (string) TCP port ranges used to access the app.
        /// </summary>
        public readonly ImmutableArray<string> TcpPortRanges;
        /// <summary>
        /// (string) UDP port ranges used to access the app.
        /// </summary>
        public readonly ImmutableArray<string> UdpPortRanges;

        [OutputConstructor]
        private GetZPAApplicationSegmentInspectionResult(
            string bypassType,

            string creationTime,

            string description,

            ImmutableArray<string> domainNames,

            bool doubleEncrypt,

            bool enabled,

            string healthCheckType,

            string healthReporting,

            string icmpAccessType,

            string? id,

            ImmutableArray<Outputs.GetZPAApplicationSegmentInspectionInspectionAppResult> inspectionApps,

            bool ipAnchored,

            bool isCnameEnabled,

            string modifiedBy,

            string modifiedTime,

            string? name,

            bool passiveHealthEnabled,

            string segmentGroupId,

            string segmentGroupName,

            bool selectConnectorCloseToApp,

            ImmutableArray<Outputs.GetZPAApplicationSegmentInspectionServerGroupResult> serverGroups,

            ImmutableArray<string> tcpPortRanges,

            ImmutableArray<string> udpPortRanges)
        {
            BypassType = bypassType;
            CreationTime = creationTime;
            Description = description;
            DomainNames = domainNames;
            DoubleEncrypt = doubleEncrypt;
            Enabled = enabled;
            HealthCheckType = healthCheckType;
            HealthReporting = healthReporting;
            IcmpAccessType = icmpAccessType;
            Id = id;
            InspectionApps = inspectionApps;
            IpAnchored = ipAnchored;
            IsCnameEnabled = isCnameEnabled;
            ModifiedBy = modifiedBy;
            ModifiedTime = modifiedTime;
            Name = name;
            PassiveHealthEnabled = passiveHealthEnabled;
            SegmentGroupId = segmentGroupId;
            SegmentGroupName = segmentGroupName;
            SelectConnectorCloseToApp = selectConnectorCloseToApp;
            ServerGroups = serverGroups;
            TcpPortRanges = tcpPortRanges;
            UdpPortRanges = udpPortRanges;
        }
    }
}
