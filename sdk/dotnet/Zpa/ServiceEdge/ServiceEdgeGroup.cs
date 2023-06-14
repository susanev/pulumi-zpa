// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;
using Pulumi;

namespace zscaler.PulumiPackage.Zpa.ServiceEdge
{
    /// <summary>
    /// The **zpa_service_edge_group** resource creates a service edge group in the Zscaler Private Access cloud. This resource can then be referenced in a service edge connector.
    /// 
    /// ## Example Usage
    /// 
    /// ```csharp
    /// using System.Collections.Generic;
    /// using Pulumi;
    /// using Zpa = zscaler.PulumiPackage.Zpa;
    /// 
    /// return await Deployment.RunAsync(() =&gt; 
    /// {
    ///     // ZPA Service Edge Group resource - Trusted Network
    ///     var serviceEdgeGroupSjc = new Zpa.ServiceEdge.ServiceEdgeGroup("serviceEdgeGroupSjc", new()
    ///     {
    ///         Description = "Service Edge Group in San Jose",
    ///         Enabled = true,
    ///         IsPublic = true,
    ///         UpgradeDay = "SUNDAY",
    ///         UpgradeTimeInSecs = "66600",
    ///         Latitude = "37.3382082",
    ///         Longitude = "-121.8863286",
    ///         Location = "San Jose, CA, USA",
    ///         VersionProfileName = "New Release",
    ///         TrustedNetworks = new[]
    ///         {
    ///             new Zpa.ServiceEdge.Inputs.ServiceEdgeGroupTrustedNetworkArgs
    ///             {
    ///                 Ids = new[]
    ///                 {
    ///                     data.Zpa_trusted_network.Example.Id,
    ///                 },
    ///             },
    ///         },
    ///     });
    /// 
    /// });
    /// ```
    /// 
    /// ```csharp
    /// using System.Collections.Generic;
    /// using Pulumi;
    /// using Zpa = zscaler.PulumiPackage.Zpa;
    /// 
    /// return await Deployment.RunAsync(() =&gt; 
    /// {
    ///     // ZPA Service Edge Group resource - No Trusted Network
    ///     var serviceEdgeGroupNyc = new Zpa.ServiceEdge.ServiceEdgeGroup("serviceEdgeGroupNyc", new()
    ///     {
    ///         Description = "Service Edge Group in New York",
    ///         Enabled = true,
    ///         IsPublic = true,
    ///         Latitude = "40.7128",
    ///         Location = "New York, NY, USA",
    ///         Longitude = "-73.935242",
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
    /// Zscaler offers a dedicated tool called Zscaler-Terraformer to allow the automated import of ZPA configurations into Terraform-compliant HashiCorp Configuration Language. [Visit](https://github.com/zscaler/zscaler-terraformer) Service Edge Group can be imported; use `&lt;SERVER EDGE GROUP ID&gt;` or `&lt;SERVER EDGE GROUP NAME&gt;` as the import ID. For example
    /// 
    /// ```sh
    ///  $ pulumi import zpa:ServiceEdge/serviceEdgeGroup:ServiceEdgeGroup example &lt;service_edge_group_id&gt;
    /// ```
    /// 
    ///  or
    /// 
    /// ```sh
    ///  $ pulumi import zpa:ServiceEdge/serviceEdgeGroup:ServiceEdgeGroup example &lt;service_edge_group_name&gt;
    /// ```
    /// </summary>
    [ZpaResourceType("zpa:ServiceEdge/serviceEdgeGroup:ServiceEdgeGroup")]
    public partial class ServiceEdgeGroup : global::Pulumi.CustomResource
    {
        /// <summary>
        /// This field controls dynamic discovery of the servers.
        /// </summary>
        [Output("cityCountry")]
        public Output<string> CityCountry { get; private set; } = null!;

        /// <summary>
        /// This field is an array of app-connector-id only.
        /// </summary>
        [Output("countryCode")]
        public Output<string> CountryCode { get; private set; } = null!;

        /// <summary>
        /// Description of the Service Edge Group.
        /// </summary>
        [Output("description")]
        public Output<string?> Description { get; private set; } = null!;

        /// <summary>
        /// Whether this Service Edge Group is enabled or not. Default value: `true` Supported values: `true`, `false`
        /// </summary>
        [Output("enabled")]
        public Output<bool?> Enabled { get; private set; } = null!;

        /// <summary>
        /// Enable or disable public access for the Service Edge Group. Default value: `false` Supported values: `true`, `false`
        /// </summary>
        [Output("isPublic")]
        public Output<bool?> IsPublic { get; private set; } = null!;

        /// <summary>
        /// Latitude for the Service Edge Group. Integer or decimal with values in the range of `-90` to `90`
        /// </summary>
        [Output("latitude")]
        public Output<string> Latitude { get; private set; } = null!;

        /// <summary>
        /// Location for the Service Edge Group.
        /// </summary>
        [Output("location")]
        public Output<string> Location { get; private set; } = null!;

        /// <summary>
        /// Longitude for the Service Edge Group. Integer or decimal with values in the range of `-180` to `180`
        /// </summary>
        [Output("longitude")]
        public Output<string> Longitude { get; private set; } = null!;

        /// <summary>
        /// Name of the Service Edge Group.
        /// </summary>
        [Output("name")]
        public Output<string> Name { get; private set; } = null!;

        /// <summary>
        /// Whether the default version profile of the App Connector Group is applied or overridden. Default: `false` Supported values: `true`, `false`
        /// </summary>
        [Output("overrideVersionProfile")]
        public Output<bool?> OverrideVersionProfile { get; private set; } = null!;

        [Output("serviceEdges")]
        public Output<ImmutableArray<Outputs.ServiceEdgeGroupServiceEdge>> ServiceEdges { get; private set; } = null!;

        /// <summary>
        /// Trusted networks for this Service Edge Group. List of trusted network objects
        /// </summary>
        [Output("trustedNetworks")]
        public Output<ImmutableArray<Outputs.ServiceEdgeGroupTrustedNetwork>> TrustedNetworks { get; private set; } = null!;

        /// <summary>
        /// Service Edges in this group will attempt to update to a newer version of the software during this specified day. Default value: `SUNDAY` List of valid days (i.e., Sunday, Monday)
        /// </summary>
        [Output("upgradeDay")]
        public Output<string?> UpgradeDay { get; private set; } = null!;

        /// <summary>
        /// Service Edges in this group will attempt to update to a newer version of the software during this specified time. Default value: `66600` Integer in seconds (i..e, 66600). The integer must be greater than or equal to 0 and less than `86400`, in `15` minute intervals
        /// </summary>
        [Output("upgradeTimeInSecs")]
        public Output<string?> UpgradeTimeInSecs { get; private set; } = null!;

        /// <summary>
        /// ID of the version profile. To learn more, see Version Profile Use Cases. Supported values are:
        /// </summary>
        [Output("versionProfileId")]
        public Output<string> VersionProfileId { get; private set; } = null!;

        /// <summary>
        /// ID of the version profile.
        /// </summary>
        [Output("versionProfileName")]
        public Output<string> VersionProfileName { get; private set; } = null!;

        /// <summary>
        /// ID of the version profile.
        /// </summary>
        [Output("versionProfileVisibilityScope")]
        public Output<string> VersionProfileVisibilityScope { get; private set; } = null!;


        /// <summary>
        /// Create a ServiceEdgeGroup resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public ServiceEdgeGroup(string name, ServiceEdgeGroupArgs args, CustomResourceOptions? options = null)
            : base("zpa:ServiceEdge/serviceEdgeGroup:ServiceEdgeGroup", name, args ?? new ServiceEdgeGroupArgs(), MakeResourceOptions(options, ""))
        {
        }

        private ServiceEdgeGroup(string name, Input<string> id, ServiceEdgeGroupState? state = null, CustomResourceOptions? options = null)
            : base("zpa:ServiceEdge/serviceEdgeGroup:ServiceEdgeGroup", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing ServiceEdgeGroup resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static ServiceEdgeGroup Get(string name, Input<string> id, ServiceEdgeGroupState? state = null, CustomResourceOptions? options = null)
        {
            return new ServiceEdgeGroup(name, id, state, options);
        }
    }

    public sealed class ServiceEdgeGroupArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// This field controls dynamic discovery of the servers.
        /// </summary>
        [Input("cityCountry")]
        public Input<string>? CityCountry { get; set; }

        /// <summary>
        /// This field is an array of app-connector-id only.
        /// </summary>
        [Input("countryCode")]
        public Input<string>? CountryCode { get; set; }

        /// <summary>
        /// Description of the Service Edge Group.
        /// </summary>
        [Input("description")]
        public Input<string>? Description { get; set; }

        /// <summary>
        /// Whether this Service Edge Group is enabled or not. Default value: `true` Supported values: `true`, `false`
        /// </summary>
        [Input("enabled")]
        public Input<bool>? Enabled { get; set; }

        /// <summary>
        /// Enable or disable public access for the Service Edge Group. Default value: `false` Supported values: `true`, `false`
        /// </summary>
        [Input("isPublic")]
        public Input<bool>? IsPublic { get; set; }

        /// <summary>
        /// Latitude for the Service Edge Group. Integer or decimal with values in the range of `-90` to `90`
        /// </summary>
        [Input("latitude", required: true)]
        public Input<string> Latitude { get; set; } = null!;

        /// <summary>
        /// Location for the Service Edge Group.
        /// </summary>
        [Input("location", required: true)]
        public Input<string> Location { get; set; } = null!;

        /// <summary>
        /// Longitude for the Service Edge Group. Integer or decimal with values in the range of `-180` to `180`
        /// </summary>
        [Input("longitude", required: true)]
        public Input<string> Longitude { get; set; } = null!;

        /// <summary>
        /// Name of the Service Edge Group.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// Whether the default version profile of the App Connector Group is applied or overridden. Default: `false` Supported values: `true`, `false`
        /// </summary>
        [Input("overrideVersionProfile")]
        public Input<bool>? OverrideVersionProfile { get; set; }

        [Input("serviceEdges")]
        private InputList<Inputs.ServiceEdgeGroupServiceEdgeArgs>? _serviceEdges;
        public InputList<Inputs.ServiceEdgeGroupServiceEdgeArgs> ServiceEdges
        {
            get => _serviceEdges ?? (_serviceEdges = new InputList<Inputs.ServiceEdgeGroupServiceEdgeArgs>());
            set => _serviceEdges = value;
        }

        [Input("trustedNetworks")]
        private InputList<Inputs.ServiceEdgeGroupTrustedNetworkArgs>? _trustedNetworks;

        /// <summary>
        /// Trusted networks for this Service Edge Group. List of trusted network objects
        /// </summary>
        public InputList<Inputs.ServiceEdgeGroupTrustedNetworkArgs> TrustedNetworks
        {
            get => _trustedNetworks ?? (_trustedNetworks = new InputList<Inputs.ServiceEdgeGroupTrustedNetworkArgs>());
            set => _trustedNetworks = value;
        }

        /// <summary>
        /// Service Edges in this group will attempt to update to a newer version of the software during this specified day. Default value: `SUNDAY` List of valid days (i.e., Sunday, Monday)
        /// </summary>
        [Input("upgradeDay")]
        public Input<string>? UpgradeDay { get; set; }

        /// <summary>
        /// Service Edges in this group will attempt to update to a newer version of the software during this specified time. Default value: `66600` Integer in seconds (i..e, 66600). The integer must be greater than or equal to 0 and less than `86400`, in `15` minute intervals
        /// </summary>
        [Input("upgradeTimeInSecs")]
        public Input<string>? UpgradeTimeInSecs { get; set; }

        /// <summary>
        /// ID of the version profile. To learn more, see Version Profile Use Cases. Supported values are:
        /// </summary>
        [Input("versionProfileId")]
        public Input<string>? VersionProfileId { get; set; }

        /// <summary>
        /// ID of the version profile.
        /// </summary>
        [Input("versionProfileName")]
        public Input<string>? VersionProfileName { get; set; }

        public ServiceEdgeGroupArgs()
        {
        }
        public static new ServiceEdgeGroupArgs Empty => new ServiceEdgeGroupArgs();
    }

    public sealed class ServiceEdgeGroupState : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// This field controls dynamic discovery of the servers.
        /// </summary>
        [Input("cityCountry")]
        public Input<string>? CityCountry { get; set; }

        /// <summary>
        /// This field is an array of app-connector-id only.
        /// </summary>
        [Input("countryCode")]
        public Input<string>? CountryCode { get; set; }

        /// <summary>
        /// Description of the Service Edge Group.
        /// </summary>
        [Input("description")]
        public Input<string>? Description { get; set; }

        /// <summary>
        /// Whether this Service Edge Group is enabled or not. Default value: `true` Supported values: `true`, `false`
        /// </summary>
        [Input("enabled")]
        public Input<bool>? Enabled { get; set; }

        /// <summary>
        /// Enable or disable public access for the Service Edge Group. Default value: `false` Supported values: `true`, `false`
        /// </summary>
        [Input("isPublic")]
        public Input<bool>? IsPublic { get; set; }

        /// <summary>
        /// Latitude for the Service Edge Group. Integer or decimal with values in the range of `-90` to `90`
        /// </summary>
        [Input("latitude")]
        public Input<string>? Latitude { get; set; }

        /// <summary>
        /// Location for the Service Edge Group.
        /// </summary>
        [Input("location")]
        public Input<string>? Location { get; set; }

        /// <summary>
        /// Longitude for the Service Edge Group. Integer or decimal with values in the range of `-180` to `180`
        /// </summary>
        [Input("longitude")]
        public Input<string>? Longitude { get; set; }

        /// <summary>
        /// Name of the Service Edge Group.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// Whether the default version profile of the App Connector Group is applied or overridden. Default: `false` Supported values: `true`, `false`
        /// </summary>
        [Input("overrideVersionProfile")]
        public Input<bool>? OverrideVersionProfile { get; set; }

        [Input("serviceEdges")]
        private InputList<Inputs.ServiceEdgeGroupServiceEdgeGetArgs>? _serviceEdges;
        public InputList<Inputs.ServiceEdgeGroupServiceEdgeGetArgs> ServiceEdges
        {
            get => _serviceEdges ?? (_serviceEdges = new InputList<Inputs.ServiceEdgeGroupServiceEdgeGetArgs>());
            set => _serviceEdges = value;
        }

        [Input("trustedNetworks")]
        private InputList<Inputs.ServiceEdgeGroupTrustedNetworkGetArgs>? _trustedNetworks;

        /// <summary>
        /// Trusted networks for this Service Edge Group. List of trusted network objects
        /// </summary>
        public InputList<Inputs.ServiceEdgeGroupTrustedNetworkGetArgs> TrustedNetworks
        {
            get => _trustedNetworks ?? (_trustedNetworks = new InputList<Inputs.ServiceEdgeGroupTrustedNetworkGetArgs>());
            set => _trustedNetworks = value;
        }

        /// <summary>
        /// Service Edges in this group will attempt to update to a newer version of the software during this specified day. Default value: `SUNDAY` List of valid days (i.e., Sunday, Monday)
        /// </summary>
        [Input("upgradeDay")]
        public Input<string>? UpgradeDay { get; set; }

        /// <summary>
        /// Service Edges in this group will attempt to update to a newer version of the software during this specified time. Default value: `66600` Integer in seconds (i..e, 66600). The integer must be greater than or equal to 0 and less than `86400`, in `15` minute intervals
        /// </summary>
        [Input("upgradeTimeInSecs")]
        public Input<string>? UpgradeTimeInSecs { get; set; }

        /// <summary>
        /// ID of the version profile. To learn more, see Version Profile Use Cases. Supported values are:
        /// </summary>
        [Input("versionProfileId")]
        public Input<string>? VersionProfileId { get; set; }

        /// <summary>
        /// ID of the version profile.
        /// </summary>
        [Input("versionProfileName")]
        public Input<string>? VersionProfileName { get; set; }

        /// <summary>
        /// ID of the version profile.
        /// </summary>
        [Input("versionProfileVisibilityScope")]
        public Input<string>? VersionProfileVisibilityScope { get; set; }

        public ServiceEdgeGroupState()
        {
        }
        public static new ServiceEdgeGroupState Empty => new ServiceEdgeGroupState();
    }
}