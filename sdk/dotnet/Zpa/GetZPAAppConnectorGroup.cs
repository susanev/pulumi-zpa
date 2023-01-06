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
    public static class GetZPAAppConnectorGroup
    {
        /// <summary>
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
        ///     var foo = Zpa.GetZPAAppConnectorGroup.Invoke(new()
        ///     {
        ///         Name = "DataCenter",
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
        ///     var foo = Zpa.GetZPAAppConnectorGroup.Invoke(new()
        ///     {
        ///         Id = "123456789",
        ///     });
        /// 
        /// });
        /// ```
        /// {{% /example %}}
        /// {{% /examples %}}
        /// </summary>
        public static Task<GetZPAAppConnectorGroupResult> InvokeAsync(GetZPAAppConnectorGroupArgs? args = null, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.InvokeAsync<GetZPAAppConnectorGroupResult>("zpa:index/getZPAAppConnectorGroup:getZPAAppConnectorGroup", args ?? new GetZPAAppConnectorGroupArgs(), options.WithDefaults());

        /// <summary>
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
        ///     var foo = Zpa.GetZPAAppConnectorGroup.Invoke(new()
        ///     {
        ///         Name = "DataCenter",
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
        ///     var foo = Zpa.GetZPAAppConnectorGroup.Invoke(new()
        ///     {
        ///         Id = "123456789",
        ///     });
        /// 
        /// });
        /// ```
        /// {{% /example %}}
        /// {{% /examples %}}
        /// </summary>
        public static Output<GetZPAAppConnectorGroupResult> Invoke(GetZPAAppConnectorGroupInvokeArgs? args = null, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.Invoke<GetZPAAppConnectorGroupResult>("zpa:index/getZPAAppConnectorGroup:getZPAAppConnectorGroup", args ?? new GetZPAAppConnectorGroupInvokeArgs(), options.WithDefaults());
    }


    public sealed class GetZPAAppConnectorGroupArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// ID of the App Connector Group.
        /// </summary>
        [Input("id")]
        public string? Id { get; set; }

        /// <summary>
        /// Name of the App Connector Group.
        /// </summary>
        [Input("name")]
        public string? Name { get; set; }

        /// <summary>
        /// (bool) Whether the default version profile of the App Connector Group is applied or overridden. Default: `false` Supported values: `true`, `false`
        /// </summary>
        [Input("overrideVersionProfile")]
        public bool? OverrideVersionProfile { get; set; }

        public GetZPAAppConnectorGroupArgs()
        {
        }
        public static new GetZPAAppConnectorGroupArgs Empty => new GetZPAAppConnectorGroupArgs();
    }

    public sealed class GetZPAAppConnectorGroupInvokeArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// ID of the App Connector Group.
        /// </summary>
        [Input("id")]
        public Input<string>? Id { get; set; }

        /// <summary>
        /// Name of the App Connector Group.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// (bool) Whether the default version profile of the App Connector Group is applied or overridden. Default: `false` Supported values: `true`, `false`
        /// </summary>
        [Input("overrideVersionProfile")]
        public Input<bool>? OverrideVersionProfile { get; set; }

        public GetZPAAppConnectorGroupInvokeArgs()
        {
        }
        public static new GetZPAAppConnectorGroupInvokeArgs Empty => new GetZPAAppConnectorGroupInvokeArgs();
    }


    [OutputType]
    public sealed class GetZPAAppConnectorGroupResult
    {
        /// <summary>
        /// (String) Whether Double Encryption is enabled or disabled for the app.
        /// </summary>
        public readonly string CityCountry;
        public readonly ImmutableArray<Outputs.GetZPAAppConnectorGroupConnectorResult> Connectors;
        /// <summary>
        /// (String)
        /// </summary>
        public readonly string CountryCode;
        public readonly string CreationTime;
        /// <summary>
        /// (String) Description of the App Connector Group.
        /// </summary>
        public readonly string Description;
        /// <summary>
        /// (String)
        /// </summary>
        public readonly string DnsQueryType;
        /// <summary>
        /// (String) Whether this App Connector Group is enabled or not. Default value: `true`. Supported values: `true`, `false`
        /// </summary>
        public readonly bool Enabled;
        /// <summary>
        /// (String)
        /// </summary>
        public readonly string GeoLocationId;
        public readonly string? Id;
        /// <summary>
        /// (String) Latitude of the App Connector Group. Integer or decimal. With values in the range of `-90` to `90`
        /// </summary>
        public readonly string Latitude;
        /// <summary>
        /// (String) Location of the App Connector Group.
        /// </summary>
        public readonly string Location;
        /// <summary>
        /// (String) Longitude of the App Connector Group. Integer or decimal. With values in the range of `-180` to `180`
        /// </summary>
        public readonly string Longitude;
        public readonly bool LssAppConnectorGroup;
        public readonly string ModifiedTime;
        public readonly string Modifiedby;
        public readonly string? Name;
        /// <summary>
        /// (bool) Whether the default version profile of the App Connector Group is applied or overridden. Default: `false` Supported values: `true`, `false`
        /// </summary>
        public readonly bool? OverrideVersionProfile;
        public readonly ImmutableArray<Outputs.GetZPAAppConnectorGroupServerGroupResult> ServerGroups;
        public readonly bool TcpQuickAckApp;
        public readonly bool TcpQuickAckAssistant;
        public readonly bool TcpQuickAckReadAssistant;
        /// <summary>
        /// (String) App Connectors in this group will attempt to update to a newer version of the software during this specified day
        /// </summary>
        public readonly string UpgradeDay;
        /// <summary>
        /// (String) App Connectors in this group will attempt to update to a newer version of the software during this specified time. Default value: `66600`. Integer in seconds (i.e., `-66600`). The integer should be greater than or equal to `0` and less than `86400`, in `15` minute intervals
        /// </summary>
        public readonly string UpgradeTimeInSecs;
        public readonly bool UseInDrMode;
        /// <summary>
        /// (String) ID of the version profile.
        /// Exported values are:
        /// </summary>
        public readonly string VersionProfileId;
        /// <summary>
        /// (String)
        /// Exported values are:
        /// </summary>
        public readonly string VersionProfileName;
        /// <summary>
        /// (String)
        /// Exported values are:
        /// </summary>
        public readonly string VersionProfileVisibilityScope;

        [OutputConstructor]
        private GetZPAAppConnectorGroupResult(
            string cityCountry,

            ImmutableArray<Outputs.GetZPAAppConnectorGroupConnectorResult> connectors,

            string countryCode,

            string creationTime,

            string description,

            string dnsQueryType,

            bool enabled,

            string geoLocationId,

            string? id,

            string latitude,

            string location,

            string longitude,

            bool lssAppConnectorGroup,

            string modifiedTime,

            string modifiedby,

            string? name,

            bool? overrideVersionProfile,

            ImmutableArray<Outputs.GetZPAAppConnectorGroupServerGroupResult> serverGroups,

            bool tcpQuickAckApp,

            bool tcpQuickAckAssistant,

            bool tcpQuickAckReadAssistant,

            string upgradeDay,

            string upgradeTimeInSecs,

            bool useInDrMode,

            string versionProfileId,

            string versionProfileName,

            string versionProfileVisibilityScope)
        {
            CityCountry = cityCountry;
            Connectors = connectors;
            CountryCode = countryCode;
            CreationTime = creationTime;
            Description = description;
            DnsQueryType = dnsQueryType;
            Enabled = enabled;
            GeoLocationId = geoLocationId;
            Id = id;
            Latitude = latitude;
            Location = location;
            Longitude = longitude;
            LssAppConnectorGroup = lssAppConnectorGroup;
            ModifiedTime = modifiedTime;
            Modifiedby = modifiedby;
            Name = name;
            OverrideVersionProfile = overrideVersionProfile;
            ServerGroups = serverGroups;
            TcpQuickAckApp = tcpQuickAckApp;
            TcpQuickAckAssistant = tcpQuickAckAssistant;
            TcpQuickAckReadAssistant = tcpQuickAckReadAssistant;
            UpgradeDay = upgradeDay;
            UpgradeTimeInSecs = upgradeTimeInSecs;
            UseInDrMode = useInDrMode;
            VersionProfileId = versionProfileId;
            VersionProfileName = versionProfileName;
            VersionProfileVisibilityScope = versionProfileVisibilityScope;
        }
    }
}
