// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../utilities";

/**
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as zpa from "@zscaler/pulumi-zpa";
 *
 * // Create a App Connector Group
 * const example = new zpa.appconnectorgroup.ConnectorGroup("example", {
 *     cityCountry: "San Jose, CA",
 *     countryCode: "US",
 *     description: "Example",
 *     dnsQueryType: "IPV4_IPV6",
 *     enabled: true,
 *     latitude: "37.338",
 *     location: "San Jose, CA, US",
 *     longitude: "-121.8863",
 *     overrideVersionProfile: true,
 *     upgradeDay: "SUNDAY",
 *     upgradeTimeInSecs: "66600",
 *     versionProfileName: "New Release",
 * });
 * ```
 *
 * ## Import
 *
 * Zscaler offers a dedicated tool called Zscaler-Terraformer to allow the automated import of ZPA configurations into Terraform-compliant HashiCorp Configuration Language. [Visit](https://github.com/zscaler/zscaler-terraformer) App Connector Group can be imported by using `<APP CONNECTOR GROUP ID>` or `<APP CONNECTOR GROUP NAME>`as the import ID.
 *
 * ```sh
 *  $ pulumi import zpa:AppConnectorGroup/connectorGroup:ConnectorGroup example <app_connector_group_id>
 * ```
 *
 *  or
 *
 * ```sh
 *  $ pulumi import zpa:AppConnectorGroup/connectorGroup:ConnectorGroup example <app_connector_group_name>
 * ```
 */
export class ConnectorGroup extends pulumi.CustomResource {
    /**
     * Get an existing ConnectorGroup resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: ConnectorGroupState, opts?: pulumi.CustomResourceOptions): ConnectorGroup {
        return new ConnectorGroup(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'zpa:AppConnectorGroup/connectorGroup:ConnectorGroup';

    /**
     * Returns true if the given object is an instance of ConnectorGroup.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is ConnectorGroup {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === ConnectorGroup.__pulumiType;
    }

    /**
     * Whether Double Encryption is enabled or disabled for the app. i.e ``"San Jose, US"``
     */
    public readonly cityCountry!: pulumi.Output<string>;
    /**
     * i.e ``"US"``, ``"CA"``
     */
    public readonly countryCode!: pulumi.Output<string>;
    /**
     * Description of the App Connector Group.
     */
    public readonly description!: pulumi.Output<string | undefined>;
    /**
     * Supported values are:
     */
    public readonly dnsQueryType!: pulumi.Output<string | undefined>;
    /**
     * Whether this App Connector Group is enabled or not. Default value: `true`. Supported values: `true`, `false`
     */
    public readonly enabled!: pulumi.Output<boolean>;
    /**
     * Latitude of the App Connector Group. Integer or decimal. With values in the range of `-90` to `90`
     */
    public readonly latitude!: pulumi.Output<string>;
    /**
     * Location of the App Connector Group. i.e ``"San Jose, CA, USA"``
     */
    public readonly location!: pulumi.Output<string>;
    /**
     * Longitude of the App Connector Group. Integer or decimal. With values in the range of `-180` to `180`
     */
    public readonly longitude!: pulumi.Output<string>;
    public readonly lssAppConnectorGroup!: pulumi.Output<boolean>;
    /**
     * Name of the App Connector Group.
     */
    public readonly name!: pulumi.Output<string>;
    /**
     * Whether the default version profile of the App Connector Group is applied or overridden. Default: `false` Supported values: `true`, `false`
     */
    public readonly overrideVersionProfile!: pulumi.Output<boolean>;
    /**
     * Whether TCP Quick Acknowledgement is enabled or disabled for the application. The tcpQuickAckApp, tcpQuickAckAssistant, and tcpQuickAckReadAssistant fields must all share the same value. Supported values: `true`, `false`
     */
    public readonly tcpQuickAckApp!: pulumi.Output<boolean>;
    /**
     * Whether TCP Quick Acknowledgement is enabled or disabled for the application. The tcpQuickAckApp, tcpQuickAckAssistant, and tcpQuickAckReadAssistant fields must all share the same value. Supported values: `true`, `false`
     */
    public readonly tcpQuickAckAssistant!: pulumi.Output<boolean>;
    /**
     * Whether TCP Quick Acknowledgement is enabled or disabled for the application. The tcpQuickAckApp, tcpQuickAckAssistant, and tcpQuickAckReadAssistant fields must all share the same value. Supported values: `true`, `false`
     */
    public readonly tcpQuickAckReadAssistant!: pulumi.Output<boolean>;
    /**
     * App Connectors in this group will attempt to update to a newer version of the software during this specified day i.e ``SUNDAY``
     */
    public readonly upgradeDay!: pulumi.Output<string | undefined>;
    /**
     * App Connectors in this group will attempt to update to a newer version of the software during this specified time. Default value: `66600`. Integer in seconds (i.e., `-66600`). The integer should be greater than or equal to `0` and less than `86400`, in `15` minute intervals
     */
    public readonly upgradeTimeInSecs!: pulumi.Output<string | undefined>;
    /**
     * Supported values: `true`, `false`
     */
    public readonly useInDrMode!: pulumi.Output<boolean>;
    /**
     * ID of the version profile. To learn more, see Version Profile Use Cases. Supported values are:
     */
    public readonly versionProfileId!: pulumi.Output<string>;
    /**
     * Name of the version profile. To learn more, see Version Profile Use Cases. This value is required, if the value for
     * overrideVersionProfile is set to true
     */
    public readonly versionProfileName!: pulumi.Output<string>;

    /**
     * Create a ConnectorGroup resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: ConnectorGroupArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: ConnectorGroupArgs | ConnectorGroupState, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as ConnectorGroupState | undefined;
            resourceInputs["cityCountry"] = state ? state.cityCountry : undefined;
            resourceInputs["countryCode"] = state ? state.countryCode : undefined;
            resourceInputs["description"] = state ? state.description : undefined;
            resourceInputs["dnsQueryType"] = state ? state.dnsQueryType : undefined;
            resourceInputs["enabled"] = state ? state.enabled : undefined;
            resourceInputs["latitude"] = state ? state.latitude : undefined;
            resourceInputs["location"] = state ? state.location : undefined;
            resourceInputs["longitude"] = state ? state.longitude : undefined;
            resourceInputs["lssAppConnectorGroup"] = state ? state.lssAppConnectorGroup : undefined;
            resourceInputs["name"] = state ? state.name : undefined;
            resourceInputs["overrideVersionProfile"] = state ? state.overrideVersionProfile : undefined;
            resourceInputs["tcpQuickAckApp"] = state ? state.tcpQuickAckApp : undefined;
            resourceInputs["tcpQuickAckAssistant"] = state ? state.tcpQuickAckAssistant : undefined;
            resourceInputs["tcpQuickAckReadAssistant"] = state ? state.tcpQuickAckReadAssistant : undefined;
            resourceInputs["upgradeDay"] = state ? state.upgradeDay : undefined;
            resourceInputs["upgradeTimeInSecs"] = state ? state.upgradeTimeInSecs : undefined;
            resourceInputs["useInDrMode"] = state ? state.useInDrMode : undefined;
            resourceInputs["versionProfileId"] = state ? state.versionProfileId : undefined;
            resourceInputs["versionProfileName"] = state ? state.versionProfileName : undefined;
        } else {
            const args = argsOrState as ConnectorGroupArgs | undefined;
            if ((!args || args.latitude === undefined) && !opts.urn) {
                throw new Error("Missing required property 'latitude'");
            }
            if ((!args || args.location === undefined) && !opts.urn) {
                throw new Error("Missing required property 'location'");
            }
            if ((!args || args.longitude === undefined) && !opts.urn) {
                throw new Error("Missing required property 'longitude'");
            }
            resourceInputs["cityCountry"] = args ? args.cityCountry : undefined;
            resourceInputs["countryCode"] = args ? args.countryCode : undefined;
            resourceInputs["description"] = args ? args.description : undefined;
            resourceInputs["dnsQueryType"] = args ? args.dnsQueryType : undefined;
            resourceInputs["enabled"] = args ? args.enabled : undefined;
            resourceInputs["latitude"] = args ? args.latitude : undefined;
            resourceInputs["location"] = args ? args.location : undefined;
            resourceInputs["longitude"] = args ? args.longitude : undefined;
            resourceInputs["lssAppConnectorGroup"] = args ? args.lssAppConnectorGroup : undefined;
            resourceInputs["name"] = args ? args.name : undefined;
            resourceInputs["overrideVersionProfile"] = args ? args.overrideVersionProfile : undefined;
            resourceInputs["tcpQuickAckApp"] = args ? args.tcpQuickAckApp : undefined;
            resourceInputs["tcpQuickAckAssistant"] = args ? args.tcpQuickAckAssistant : undefined;
            resourceInputs["tcpQuickAckReadAssistant"] = args ? args.tcpQuickAckReadAssistant : undefined;
            resourceInputs["upgradeDay"] = args ? args.upgradeDay : undefined;
            resourceInputs["upgradeTimeInSecs"] = args ? args.upgradeTimeInSecs : undefined;
            resourceInputs["useInDrMode"] = args ? args.useInDrMode : undefined;
            resourceInputs["versionProfileId"] = args ? args.versionProfileId : undefined;
            resourceInputs["versionProfileName"] = args ? args.versionProfileName : undefined;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(ConnectorGroup.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering ConnectorGroup resources.
 */
export interface ConnectorGroupState {
    /**
     * Whether Double Encryption is enabled or disabled for the app. i.e ``"San Jose, US"``
     */
    cityCountry?: pulumi.Input<string>;
    /**
     * i.e ``"US"``, ``"CA"``
     */
    countryCode?: pulumi.Input<string>;
    /**
     * Description of the App Connector Group.
     */
    description?: pulumi.Input<string>;
    /**
     * Supported values are:
     */
    dnsQueryType?: pulumi.Input<string>;
    /**
     * Whether this App Connector Group is enabled or not. Default value: `true`. Supported values: `true`, `false`
     */
    enabled?: pulumi.Input<boolean>;
    /**
     * Latitude of the App Connector Group. Integer or decimal. With values in the range of `-90` to `90`
     */
    latitude?: pulumi.Input<string>;
    /**
     * Location of the App Connector Group. i.e ``"San Jose, CA, USA"``
     */
    location?: pulumi.Input<string>;
    /**
     * Longitude of the App Connector Group. Integer or decimal. With values in the range of `-180` to `180`
     */
    longitude?: pulumi.Input<string>;
    lssAppConnectorGroup?: pulumi.Input<boolean>;
    /**
     * Name of the App Connector Group.
     */
    name?: pulumi.Input<string>;
    /**
     * Whether the default version profile of the App Connector Group is applied or overridden. Default: `false` Supported values: `true`, `false`
     */
    overrideVersionProfile?: pulumi.Input<boolean>;
    /**
     * Whether TCP Quick Acknowledgement is enabled or disabled for the application. The tcpQuickAckApp, tcpQuickAckAssistant, and tcpQuickAckReadAssistant fields must all share the same value. Supported values: `true`, `false`
     */
    tcpQuickAckApp?: pulumi.Input<boolean>;
    /**
     * Whether TCP Quick Acknowledgement is enabled or disabled for the application. The tcpQuickAckApp, tcpQuickAckAssistant, and tcpQuickAckReadAssistant fields must all share the same value. Supported values: `true`, `false`
     */
    tcpQuickAckAssistant?: pulumi.Input<boolean>;
    /**
     * Whether TCP Quick Acknowledgement is enabled or disabled for the application. The tcpQuickAckApp, tcpQuickAckAssistant, and tcpQuickAckReadAssistant fields must all share the same value. Supported values: `true`, `false`
     */
    tcpQuickAckReadAssistant?: pulumi.Input<boolean>;
    /**
     * App Connectors in this group will attempt to update to a newer version of the software during this specified day i.e ``SUNDAY``
     */
    upgradeDay?: pulumi.Input<string>;
    /**
     * App Connectors in this group will attempt to update to a newer version of the software during this specified time. Default value: `66600`. Integer in seconds (i.e., `-66600`). The integer should be greater than or equal to `0` and less than `86400`, in `15` minute intervals
     */
    upgradeTimeInSecs?: pulumi.Input<string>;
    /**
     * Supported values: `true`, `false`
     */
    useInDrMode?: pulumi.Input<boolean>;
    /**
     * ID of the version profile. To learn more, see Version Profile Use Cases. Supported values are:
     */
    versionProfileId?: pulumi.Input<string>;
    /**
     * Name of the version profile. To learn more, see Version Profile Use Cases. This value is required, if the value for
     * overrideVersionProfile is set to true
     */
    versionProfileName?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a ConnectorGroup resource.
 */
export interface ConnectorGroupArgs {
    /**
     * Whether Double Encryption is enabled or disabled for the app. i.e ``"San Jose, US"``
     */
    cityCountry?: pulumi.Input<string>;
    /**
     * i.e ``"US"``, ``"CA"``
     */
    countryCode?: pulumi.Input<string>;
    /**
     * Description of the App Connector Group.
     */
    description?: pulumi.Input<string>;
    /**
     * Supported values are:
     */
    dnsQueryType?: pulumi.Input<string>;
    /**
     * Whether this App Connector Group is enabled or not. Default value: `true`. Supported values: `true`, `false`
     */
    enabled?: pulumi.Input<boolean>;
    /**
     * Latitude of the App Connector Group. Integer or decimal. With values in the range of `-90` to `90`
     */
    latitude: pulumi.Input<string>;
    /**
     * Location of the App Connector Group. i.e ``"San Jose, CA, USA"``
     */
    location: pulumi.Input<string>;
    /**
     * Longitude of the App Connector Group. Integer or decimal. With values in the range of `-180` to `180`
     */
    longitude: pulumi.Input<string>;
    lssAppConnectorGroup?: pulumi.Input<boolean>;
    /**
     * Name of the App Connector Group.
     */
    name?: pulumi.Input<string>;
    /**
     * Whether the default version profile of the App Connector Group is applied or overridden. Default: `false` Supported values: `true`, `false`
     */
    overrideVersionProfile?: pulumi.Input<boolean>;
    /**
     * Whether TCP Quick Acknowledgement is enabled or disabled for the application. The tcpQuickAckApp, tcpQuickAckAssistant, and tcpQuickAckReadAssistant fields must all share the same value. Supported values: `true`, `false`
     */
    tcpQuickAckApp?: pulumi.Input<boolean>;
    /**
     * Whether TCP Quick Acknowledgement is enabled or disabled for the application. The tcpQuickAckApp, tcpQuickAckAssistant, and tcpQuickAckReadAssistant fields must all share the same value. Supported values: `true`, `false`
     */
    tcpQuickAckAssistant?: pulumi.Input<boolean>;
    /**
     * Whether TCP Quick Acknowledgement is enabled or disabled for the application. The tcpQuickAckApp, tcpQuickAckAssistant, and tcpQuickAckReadAssistant fields must all share the same value. Supported values: `true`, `false`
     */
    tcpQuickAckReadAssistant?: pulumi.Input<boolean>;
    /**
     * App Connectors in this group will attempt to update to a newer version of the software during this specified day i.e ``SUNDAY``
     */
    upgradeDay?: pulumi.Input<string>;
    /**
     * App Connectors in this group will attempt to update to a newer version of the software during this specified time. Default value: `66600`. Integer in seconds (i.e., `-66600`). The integer should be greater than or equal to `0` and less than `86400`, in `15` minute intervals
     */
    upgradeTimeInSecs?: pulumi.Input<string>;
    /**
     * Supported values: `true`, `false`
     */
    useInDrMode?: pulumi.Input<boolean>;
    /**
     * ID of the version profile. To learn more, see Version Profile Use Cases. Supported values are:
     */
    versionProfileId?: pulumi.Input<string>;
    /**
     * Name of the version profile. To learn more, see Version Profile Use Cases. This value is required, if the value for
     * overrideVersionProfile is set to true
     */
    versionProfileName?: pulumi.Input<string>;
}