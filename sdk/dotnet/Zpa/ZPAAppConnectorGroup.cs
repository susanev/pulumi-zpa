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
    ///     // Create a App Connector Group
    ///     var example = new Zpa.ZPAAppConnectorGroup("example", new()
    ///     {
    ///         CityCountry = "San Jose, CA",
    ///         CountryCode = "US",
    ///         Description = "Example",
    ///         DnsQueryType = "IPV4_IPV6",
    ///         Enabled = true,
    ///         Latitude = "37.338",
    ///         Location = "San Jose, CA, US",
    ///         Longitude = "-121.8863",
    ///         OverrideVersionProfile = true,
    ///         UpgradeDay = "SUNDAY",
    ///         UpgradeTimeInSecs = "66600",
    ///         VersionProfileName = "New Release",
    ///     });
    /// 
    /// });
    /// ```
    /// 
    /// ## Import
    /// 
    /// Zscaler offers a dedicated tool called Zscaler-Terraformer to allow the automated import of ZPA configurations into Terraform-compliant HashiCorp Configuration Language. [Visit](https://github.com/zscaler/zscaler-terraformer) App Connector Group can be imported by using `&lt;APP CONNECTOR GROUP ID&gt;` or `&lt;APP CONNECTOR GROUP NAME&gt;`as the import ID.
    /// 
    /// ```sh
    ///  $ pulumi import zpa:index/zPAAppConnectorGroup:ZPAAppConnectorGroup example &lt;app_connector_group_id&gt;
    /// ```
    /// 
    ///  or
    /// 
    /// ```sh
    ///  $ pulumi import zpa:index/zPAAppConnectorGroup:ZPAAppConnectorGroup example &lt;app_connector_group_name&gt;
    /// ```
    /// </summary>
    [ZpaResourceType("zpa:index/zPAAppConnectorGroup:ZPAAppConnectorGroup")]
    public partial class ZPAAppConnectorGroup : global::Pulumi.CustomResource
    {
        /// <summary>
        /// Whether Double Encryption is enabled or disabled for the app. i.e ``"San Jose, US"``
        /// </summary>
        [Output("cityCountry")]
        public Output<string> CityCountry { get; private set; } = null!;

        /// <summary>
        /// i.e ``"US"``, ``"CA"``
        /// </summary>
        [Output("countryCode")]
        public Output<string> CountryCode { get; private set; } = null!;

        /// <summary>
        /// Description of the App Connector Group.
        /// </summary>
        [Output("description")]
        public Output<string?> Description { get; private set; } = null!;

        /// <summary>
        /// Supported values are:
        /// </summary>
        [Output("dnsQueryType")]
        public Output<string?> DnsQueryType { get; private set; } = null!;

        /// <summary>
        /// Whether this App Connector Group is enabled or not. Default value: `true`. Supported values: `true`, `false`
        /// </summary>
        [Output("enabled")]
        public Output<bool> Enabled { get; private set; } = null!;

        /// <summary>
        /// Latitude of the App Connector Group. Integer or decimal. With values in the range of `-90` to `90`
        /// </summary>
        [Output("latitude")]
        public Output<string> Latitude { get; private set; } = null!;

        /// <summary>
        /// Location of the App Connector Group. i.e ``"San Jose, CA, USA"``
        /// </summary>
        [Output("location")]
        public Output<string> Location { get; private set; } = null!;

        /// <summary>
        /// Longitude of the App Connector Group. Integer or decimal. With values in the range of `-180` to `180`
        /// </summary>
        [Output("longitude")]
        public Output<string> Longitude { get; private set; } = null!;

        [Output("lssAppConnectorGroup")]
        public Output<bool> LssAppConnectorGroup { get; private set; } = null!;

        /// <summary>
        /// Name of the App Connector Group.
        /// </summary>
        [Output("name")]
        public Output<string> Name { get; private set; } = null!;

        /// <summary>
        /// Whether the default version profile of the App Connector Group is applied or overridden. Default: `false` Supported values: `true`, `false`
        /// </summary>
        [Output("overrideVersionProfile")]
        public Output<bool> OverrideVersionProfile { get; private set; } = null!;

        /// <summary>
        /// Whether TCP Quick Acknowledgement is enabled or disabled for the application. The tcpQuickAckApp, tcpQuickAckAssistant, and tcpQuickAckReadAssistant fields must all share the same value. Supported values: `true`, `false`
        /// </summary>
        [Output("tcpQuickAckApp")]
        public Output<bool> TcpQuickAckApp { get; private set; } = null!;

        /// <summary>
        /// Whether TCP Quick Acknowledgement is enabled or disabled for the application. The tcpQuickAckApp, tcpQuickAckAssistant, and tcpQuickAckReadAssistant fields must all share the same value. Supported values: `true`, `false`
        /// </summary>
        [Output("tcpQuickAckAssistant")]
        public Output<bool> TcpQuickAckAssistant { get; private set; } = null!;

        /// <summary>
        /// Whether TCP Quick Acknowledgement is enabled or disabled for the application. The tcpQuickAckApp, tcpQuickAckAssistant, and tcpQuickAckReadAssistant fields must all share the same value. Supported values: `true`, `false`
        /// </summary>
        [Output("tcpQuickAckReadAssistant")]
        public Output<bool> TcpQuickAckReadAssistant { get; private set; } = null!;

        /// <summary>
        /// App Connectors in this group will attempt to update to a newer version of the software during this specified day i.e ``SUNDAY``
        /// </summary>
        [Output("upgradeDay")]
        public Output<string?> UpgradeDay { get; private set; } = null!;

        /// <summary>
        /// App Connectors in this group will attempt to update to a newer version of the software during this specified time. Default value: `66600`. Integer in seconds (i.e., `-66600`). The integer should be greater than or equal to `0` and less than `86400`, in `15` minute intervals
        /// </summary>
        [Output("upgradeTimeInSecs")]
        public Output<string?> UpgradeTimeInSecs { get; private set; } = null!;

        /// <summary>
        /// Supported values: `true`, `false`
        /// </summary>
        [Output("useInDrMode")]
        public Output<bool> UseInDrMode { get; private set; } = null!;

        /// <summary>
        /// ID of the version profile. To learn more, see Version Profile Use Cases. Supported values are:
        /// </summary>
        [Output("versionProfileId")]
        public Output<string> VersionProfileId { get; private set; } = null!;

        /// <summary>
        /// Name of the version profile. To learn more, see Version Profile Use Cases. This value is required, if the value for
        /// overrideVersionProfile is set to true
        /// </summary>
        [Output("versionProfileName")]
        public Output<string> VersionProfileName { get; private set; } = null!;


        /// <summary>
        /// Create a ZPAAppConnectorGroup resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public ZPAAppConnectorGroup(string name, ZPAAppConnectorGroupArgs args, CustomResourceOptions? options = null)
            : base("zpa:index/zPAAppConnectorGroup:ZPAAppConnectorGroup", name, args ?? new ZPAAppConnectorGroupArgs(), MakeResourceOptions(options, ""))
        {
        }

        private ZPAAppConnectorGroup(string name, Input<string> id, ZPAAppConnectorGroupState? state = null, CustomResourceOptions? options = null)
            : base("zpa:index/zPAAppConnectorGroup:ZPAAppConnectorGroup", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing ZPAAppConnectorGroup resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static ZPAAppConnectorGroup Get(string name, Input<string> id, ZPAAppConnectorGroupState? state = null, CustomResourceOptions? options = null)
        {
            return new ZPAAppConnectorGroup(name, id, state, options);
        }
    }

    public sealed class ZPAAppConnectorGroupArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Whether Double Encryption is enabled or disabled for the app. i.e ``"San Jose, US"``
        /// </summary>
        [Input("cityCountry")]
        public Input<string>? CityCountry { get; set; }

        /// <summary>
        /// i.e ``"US"``, ``"CA"``
        /// </summary>
        [Input("countryCode")]
        public Input<string>? CountryCode { get; set; }

        /// <summary>
        /// Description of the App Connector Group.
        /// </summary>
        [Input("description")]
        public Input<string>? Description { get; set; }

        /// <summary>
        /// Supported values are:
        /// </summary>
        [Input("dnsQueryType")]
        public Input<string>? DnsQueryType { get; set; }

        /// <summary>
        /// Whether this App Connector Group is enabled or not. Default value: `true`. Supported values: `true`, `false`
        /// </summary>
        [Input("enabled")]
        public Input<bool>? Enabled { get; set; }

        /// <summary>
        /// Latitude of the App Connector Group. Integer or decimal. With values in the range of `-90` to `90`
        /// </summary>
        [Input("latitude", required: true)]
        public Input<string> Latitude { get; set; } = null!;

        /// <summary>
        /// Location of the App Connector Group. i.e ``"San Jose, CA, USA"``
        /// </summary>
        [Input("location", required: true)]
        public Input<string> Location { get; set; } = null!;

        /// <summary>
        /// Longitude of the App Connector Group. Integer or decimal. With values in the range of `-180` to `180`
        /// </summary>
        [Input("longitude", required: true)]
        public Input<string> Longitude { get; set; } = null!;

        [Input("lssAppConnectorGroup")]
        public Input<bool>? LssAppConnectorGroup { get; set; }

        /// <summary>
        /// Name of the App Connector Group.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// Whether the default version profile of the App Connector Group is applied or overridden. Default: `false` Supported values: `true`, `false`
        /// </summary>
        [Input("overrideVersionProfile")]
        public Input<bool>? OverrideVersionProfile { get; set; }

        /// <summary>
        /// Whether TCP Quick Acknowledgement is enabled or disabled for the application. The tcpQuickAckApp, tcpQuickAckAssistant, and tcpQuickAckReadAssistant fields must all share the same value. Supported values: `true`, `false`
        /// </summary>
        [Input("tcpQuickAckApp")]
        public Input<bool>? TcpQuickAckApp { get; set; }

        /// <summary>
        /// Whether TCP Quick Acknowledgement is enabled or disabled for the application. The tcpQuickAckApp, tcpQuickAckAssistant, and tcpQuickAckReadAssistant fields must all share the same value. Supported values: `true`, `false`
        /// </summary>
        [Input("tcpQuickAckAssistant")]
        public Input<bool>? TcpQuickAckAssistant { get; set; }

        /// <summary>
        /// Whether TCP Quick Acknowledgement is enabled or disabled for the application. The tcpQuickAckApp, tcpQuickAckAssistant, and tcpQuickAckReadAssistant fields must all share the same value. Supported values: `true`, `false`
        /// </summary>
        [Input("tcpQuickAckReadAssistant")]
        public Input<bool>? TcpQuickAckReadAssistant { get; set; }

        /// <summary>
        /// App Connectors in this group will attempt to update to a newer version of the software during this specified day i.e ``SUNDAY``
        /// </summary>
        [Input("upgradeDay")]
        public Input<string>? UpgradeDay { get; set; }

        /// <summary>
        /// App Connectors in this group will attempt to update to a newer version of the software during this specified time. Default value: `66600`. Integer in seconds (i.e., `-66600`). The integer should be greater than or equal to `0` and less than `86400`, in `15` minute intervals
        /// </summary>
        [Input("upgradeTimeInSecs")]
        public Input<string>? UpgradeTimeInSecs { get; set; }

        /// <summary>
        /// Supported values: `true`, `false`
        /// </summary>
        [Input("useInDrMode")]
        public Input<bool>? UseInDrMode { get; set; }

        /// <summary>
        /// ID of the version profile. To learn more, see Version Profile Use Cases. Supported values are:
        /// </summary>
        [Input("versionProfileId")]
        public Input<string>? VersionProfileId { get; set; }

        /// <summary>
        /// Name of the version profile. To learn more, see Version Profile Use Cases. This value is required, if the value for
        /// overrideVersionProfile is set to true
        /// </summary>
        [Input("versionProfileName")]
        public Input<string>? VersionProfileName { get; set; }

        public ZPAAppConnectorGroupArgs()
        {
        }
        public static new ZPAAppConnectorGroupArgs Empty => new ZPAAppConnectorGroupArgs();
    }

    public sealed class ZPAAppConnectorGroupState : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Whether Double Encryption is enabled or disabled for the app. i.e ``"San Jose, US"``
        /// </summary>
        [Input("cityCountry")]
        public Input<string>? CityCountry { get; set; }

        /// <summary>
        /// i.e ``"US"``, ``"CA"``
        /// </summary>
        [Input("countryCode")]
        public Input<string>? CountryCode { get; set; }

        /// <summary>
        /// Description of the App Connector Group.
        /// </summary>
        [Input("description")]
        public Input<string>? Description { get; set; }

        /// <summary>
        /// Supported values are:
        /// </summary>
        [Input("dnsQueryType")]
        public Input<string>? DnsQueryType { get; set; }

        /// <summary>
        /// Whether this App Connector Group is enabled or not. Default value: `true`. Supported values: `true`, `false`
        /// </summary>
        [Input("enabled")]
        public Input<bool>? Enabled { get; set; }

        /// <summary>
        /// Latitude of the App Connector Group. Integer or decimal. With values in the range of `-90` to `90`
        /// </summary>
        [Input("latitude")]
        public Input<string>? Latitude { get; set; }

        /// <summary>
        /// Location of the App Connector Group. i.e ``"San Jose, CA, USA"``
        /// </summary>
        [Input("location")]
        public Input<string>? Location { get; set; }

        /// <summary>
        /// Longitude of the App Connector Group. Integer or decimal. With values in the range of `-180` to `180`
        /// </summary>
        [Input("longitude")]
        public Input<string>? Longitude { get; set; }

        [Input("lssAppConnectorGroup")]
        public Input<bool>? LssAppConnectorGroup { get; set; }

        /// <summary>
        /// Name of the App Connector Group.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// Whether the default version profile of the App Connector Group is applied or overridden. Default: `false` Supported values: `true`, `false`
        /// </summary>
        [Input("overrideVersionProfile")]
        public Input<bool>? OverrideVersionProfile { get; set; }

        /// <summary>
        /// Whether TCP Quick Acknowledgement is enabled or disabled for the application. The tcpQuickAckApp, tcpQuickAckAssistant, and tcpQuickAckReadAssistant fields must all share the same value. Supported values: `true`, `false`
        /// </summary>
        [Input("tcpQuickAckApp")]
        public Input<bool>? TcpQuickAckApp { get; set; }

        /// <summary>
        /// Whether TCP Quick Acknowledgement is enabled or disabled for the application. The tcpQuickAckApp, tcpQuickAckAssistant, and tcpQuickAckReadAssistant fields must all share the same value. Supported values: `true`, `false`
        /// </summary>
        [Input("tcpQuickAckAssistant")]
        public Input<bool>? TcpQuickAckAssistant { get; set; }

        /// <summary>
        /// Whether TCP Quick Acknowledgement is enabled or disabled for the application. The tcpQuickAckApp, tcpQuickAckAssistant, and tcpQuickAckReadAssistant fields must all share the same value. Supported values: `true`, `false`
        /// </summary>
        [Input("tcpQuickAckReadAssistant")]
        public Input<bool>? TcpQuickAckReadAssistant { get; set; }

        /// <summary>
        /// App Connectors in this group will attempt to update to a newer version of the software during this specified day i.e ``SUNDAY``
        /// </summary>
        [Input("upgradeDay")]
        public Input<string>? UpgradeDay { get; set; }

        /// <summary>
        /// App Connectors in this group will attempt to update to a newer version of the software during this specified time. Default value: `66600`. Integer in seconds (i.e., `-66600`). The integer should be greater than or equal to `0` and less than `86400`, in `15` minute intervals
        /// </summary>
        [Input("upgradeTimeInSecs")]
        public Input<string>? UpgradeTimeInSecs { get; set; }

        /// <summary>
        /// Supported values: `true`, `false`
        /// </summary>
        [Input("useInDrMode")]
        public Input<bool>? UseInDrMode { get; set; }

        /// <summary>
        /// ID of the version profile. To learn more, see Version Profile Use Cases. Supported values are:
        /// </summary>
        [Input("versionProfileId")]
        public Input<string>? VersionProfileId { get; set; }

        /// <summary>
        /// Name of the version profile. To learn more, see Version Profile Use Cases. This value is required, if the value for
        /// overrideVersionProfile is set to true
        /// </summary>
        [Input("versionProfileName")]
        public Input<string>? VersionProfileName { get; set; }

        public ZPAAppConnectorGroupState()
        {
        }
        public static new ZPAAppConnectorGroupState Empty => new ZPAAppConnectorGroupState();
    }
}
