// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "./types/input";
import * as outputs from "./types/output";
import * as utilities from "./utilities";

/**
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as zpa from "@zscaler/pulumi-zpa";
 *
 * // Create a App Connector Group
 * const exampleZPAAppConnectorGroup = new zpa.ZPAAppConnectorGroup("exampleZPAAppConnectorGroup", {
 *     description: "Example",
 *     enabled: true,
 *     cityCountry: "San Jose, CA",
 *     countryCode: "US",
 *     latitude: "37.338",
 *     longitude: "-121.8863",
 *     location: "San Jose, CA, US",
 *     upgradeDay: "SUNDAY",
 *     upgradeTimeInSecs: "66600",
 *     overrideVersionProfile: true,
 *     versionProfileId: "0",
 *     dnsQueryType: "IPV4",
 * });
 * // Create a Server Group resource with Dynamic Discovery Enabled
 * const exampleZPAServerGroup = new zpa.ZPAServerGroup("exampleZPAServerGroup", {
 *     description: "Example",
 *     enabled: true,
 *     dynamicDiscovery: true,
 *     appConnectorGroups: [{
 *         ids: [exampleZPAAppConnectorGroup.id],
 *     }],
 * }, {
 *     dependsOn: [exampleZPAAppConnectorGroup],
 * });
 * ```
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as zpa from "@zscaler/pulumi-zpa";
 *
 * // Create an application server
 * const exampleZPAApplicationServer = new zpa.ZPAApplicationServer("exampleZPAApplicationServer", {
 *     description: "Example",
 *     address: "server.example.com",
 *     enabled: true,
 * });
 * // Create a App Connector Group
 * const exampleZPAAppConnectorGroup = new zpa.ZPAAppConnectorGroup("exampleZPAAppConnectorGroup", {
 *     description: "Example",
 *     enabled: true,
 *     cityCountry: "San Jose, CA",
 *     countryCode: "US",
 *     latitude: "37.338",
 *     longitude: "-121.8863",
 *     location: "San Jose, CA, US",
 *     upgradeDay: "SUNDAY",
 *     upgradeTimeInSecs: "66600",
 *     overrideVersionProfile: true,
 *     versionProfileId: "0",
 *     dnsQueryType: "IPV4",
 * });
 * // ZPA Server Group resource with Dynamic Discovery Disabled
 * const exampleZPAServerGroup = new zpa.ZPAServerGroup("exampleZPAServerGroup", {
 *     description: "Example",
 *     enabled: true,
 *     dynamicDiscovery: false,
 *     servers: [{
 *         ids: [exampleZPAApplicationServer.id],
 *     }],
 *     appConnectorGroups: [{
 *         ids: [exampleZPAAppConnectorGroup.id],
 *     }],
 * }, {
 *     dependsOn: [
 *         exampleZPAAppConnectorGroup,
 *         zpa_application_server.server,
 *     ],
 * });
 * ```
 * ### Required
 *
 * * `name` - (Required) This field defines the name of the server group.
 * * `appConnectorGroups` - (Required)
 *   * `id` - (Required) The ID of this resource.
 *
 * ## Import
 *
 * Zscaler offers a dedicated tool called Zscaler-Terraformer to allow the automated import of ZPA configurations into Terraform-compliant HashiCorp Configuration Language. [Visit](https://github.com/zscaler/zscaler-terraformer) Server Groups can be imported; use `<SERVER GROUP ID>` or `<SERVER GROUP NAME>` as the import ID. For example
 *
 * ```sh
 *  $ pulumi import zpa:index/zPAServerGroup:ZPAServerGroup example <server_group_id>
 * ```
 *
 *  or
 *
 * ```sh
 *  $ pulumi import zpa:index/zPAServerGroup:ZPAServerGroup example <server_group_name>
 * ```
 */
export class ZPAServerGroup extends pulumi.CustomResource {
    /**
     * Get an existing ZPAServerGroup resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: ZPAServerGroupState, opts?: pulumi.CustomResourceOptions): ZPAServerGroup {
        return new ZPAServerGroup(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'zpa:index/zPAServerGroup:ZPAServerGroup';

    /**
     * Returns true if the given object is an instance of ZPAServerGroup.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is ZPAServerGroup {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === ZPAServerGroup.__pulumiType;
    }

    /**
     * List of app-connector IDs.
     */
    public readonly appConnectorGroups!: pulumi.Output<outputs.ZPAServerGroupAppConnectorGroup[]>;
    /**
     * This field is a json array of app-connector-id only.
     */
    public readonly applications!: pulumi.Output<outputs.ZPAServerGroupApplication[]>;
    public readonly configSpace!: pulumi.Output<string | undefined>;
    /**
     * (Optional) This field is the description of the server group.
     */
    public readonly description!: pulumi.Output<string | undefined>;
    /**
     * (Optional) This field controls dynamic discovery of the servers.
     */
    public readonly dynamicDiscovery!: pulumi.Output<boolean | undefined>;
    /**
     * (Optional) This field defines if the server group is enabled or disabled.
     */
    public readonly enabled!: pulumi.Output<boolean | undefined>;
    public readonly ipAnchored!: pulumi.Output<boolean | undefined>;
    /**
     * This field defines the name of the server group.
     */
    public readonly name!: pulumi.Output<string>;
    /**
     * (Block List) This field is a list of servers that are applicable only when dynamic discovery is disabled. Server name is required only in cases where the new servers need to be created in this API.
     */
    public readonly servers!: pulumi.Output<outputs.ZPAServerGroupServer[]>;

    /**
     * Create a ZPAServerGroup resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args?: ZPAServerGroupArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: ZPAServerGroupArgs | ZPAServerGroupState, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as ZPAServerGroupState | undefined;
            resourceInputs["appConnectorGroups"] = state ? state.appConnectorGroups : undefined;
            resourceInputs["applications"] = state ? state.applications : undefined;
            resourceInputs["configSpace"] = state ? state.configSpace : undefined;
            resourceInputs["description"] = state ? state.description : undefined;
            resourceInputs["dynamicDiscovery"] = state ? state.dynamicDiscovery : undefined;
            resourceInputs["enabled"] = state ? state.enabled : undefined;
            resourceInputs["ipAnchored"] = state ? state.ipAnchored : undefined;
            resourceInputs["name"] = state ? state.name : undefined;
            resourceInputs["servers"] = state ? state.servers : undefined;
        } else {
            const args = argsOrState as ZPAServerGroupArgs | undefined;
            resourceInputs["appConnectorGroups"] = args ? args.appConnectorGroups : undefined;
            resourceInputs["applications"] = args ? args.applications : undefined;
            resourceInputs["configSpace"] = args ? args.configSpace : undefined;
            resourceInputs["description"] = args ? args.description : undefined;
            resourceInputs["dynamicDiscovery"] = args ? args.dynamicDiscovery : undefined;
            resourceInputs["enabled"] = args ? args.enabled : undefined;
            resourceInputs["ipAnchored"] = args ? args.ipAnchored : undefined;
            resourceInputs["name"] = args ? args.name : undefined;
            resourceInputs["servers"] = args ? args.servers : undefined;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(ZPAServerGroup.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering ZPAServerGroup resources.
 */
export interface ZPAServerGroupState {
    /**
     * List of app-connector IDs.
     */
    appConnectorGroups?: pulumi.Input<pulumi.Input<inputs.ZPAServerGroupAppConnectorGroup>[]>;
    /**
     * This field is a json array of app-connector-id only.
     */
    applications?: pulumi.Input<pulumi.Input<inputs.ZPAServerGroupApplication>[]>;
    configSpace?: pulumi.Input<string>;
    /**
     * (Optional) This field is the description of the server group.
     */
    description?: pulumi.Input<string>;
    /**
     * (Optional) This field controls dynamic discovery of the servers.
     */
    dynamicDiscovery?: pulumi.Input<boolean>;
    /**
     * (Optional) This field defines if the server group is enabled or disabled.
     */
    enabled?: pulumi.Input<boolean>;
    ipAnchored?: pulumi.Input<boolean>;
    /**
     * This field defines the name of the server group.
     */
    name?: pulumi.Input<string>;
    /**
     * (Block List) This field is a list of servers that are applicable only when dynamic discovery is disabled. Server name is required only in cases where the new servers need to be created in this API.
     */
    servers?: pulumi.Input<pulumi.Input<inputs.ZPAServerGroupServer>[]>;
}

/**
 * The set of arguments for constructing a ZPAServerGroup resource.
 */
export interface ZPAServerGroupArgs {
    /**
     * List of app-connector IDs.
     */
    appConnectorGroups?: pulumi.Input<pulumi.Input<inputs.ZPAServerGroupAppConnectorGroup>[]>;
    /**
     * This field is a json array of app-connector-id only.
     */
    applications?: pulumi.Input<pulumi.Input<inputs.ZPAServerGroupApplication>[]>;
    configSpace?: pulumi.Input<string>;
    /**
     * (Optional) This field is the description of the server group.
     */
    description?: pulumi.Input<string>;
    /**
     * (Optional) This field controls dynamic discovery of the servers.
     */
    dynamicDiscovery?: pulumi.Input<boolean>;
    /**
     * (Optional) This field defines if the server group is enabled or disabled.
     */
    enabled?: pulumi.Input<boolean>;
    ipAnchored?: pulumi.Input<boolean>;
    /**
     * This field defines the name of the server group.
     */
    name?: pulumi.Input<string>;
    /**
     * (Block List) This field is a list of servers that are applicable only when dynamic discovery is disabled. Server name is required only in cases where the new servers need to be created in this API.
     */
    servers?: pulumi.Input<pulumi.Input<inputs.ZPAServerGroupServer>[]>;
}
