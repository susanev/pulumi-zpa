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
 * import * as zpa from "@pulumi/zpa";
 *
 * const example = zpa.getZPAServerGroup({
 *     name: "server_group_name",
 * });
 * ```
 */
export function getZPAServerGroup(args?: GetZPAServerGroupArgs, opts?: pulumi.InvokeOptions): Promise<GetZPAServerGroupResult> {
    args = args || {};

    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts || {});
    return pulumi.runtime.invoke("zpa:index/getZPAServerGroup:getZPAServerGroup", {
        "id": args.id,
        "name": args.name,
    }, opts);
}

/**
 * A collection of arguments for invoking getZPAServerGroup.
 */
export interface GetZPAServerGroupArgs {
    /**
     * The ID of the server group to be exported.
     */
    id?: string;
    /**
     * The name of the server group to be exported.
     */
    name?: string;
}

/**
 * A collection of values returned by getZPAServerGroup.
 */
export interface GetZPAServerGroupResult {
    /**
     * (string)This field is a json array of app-connector-id only.
     */
    readonly appConnectorGroups: outputs.GetZPAServerGroupAppConnectorGroup[];
    readonly applications: outputs.GetZPAServerGroupApplication[];
    /**
     * (string)
     */
    readonly configSpace: string;
    readonly creationTime: string;
    /**
     * (string) This field is the description of the server group.
     */
    readonly description: string;
    /**
     * (bool) This field controls dynamic discovery of the servers.
     */
    readonly dynamicDiscovery: boolean;
    /**
     * (bool) This field defines if the server group is enabled or disabled.
     */
    readonly enabled: boolean;
    readonly id?: string;
    /**
     * (bool)
     */
    readonly ipAnchored: boolean;
    readonly modifiedTime: string;
    readonly modifiedby: string;
    readonly name?: string;
    readonly servers: outputs.GetZPAServerGroupServer[];
}
/**
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as zpa from "@pulumi/zpa";
 *
 * const example = zpa.getZPAServerGroup({
 *     name: "server_group_name",
 * });
 * ```
 */
export function getZPAServerGroupOutput(args?: GetZPAServerGroupOutputArgs, opts?: pulumi.InvokeOptions): pulumi.Output<GetZPAServerGroupResult> {
    return pulumi.output(args).apply((a: any) => getZPAServerGroup(a, opts))
}

/**
 * A collection of arguments for invoking getZPAServerGroup.
 */
export interface GetZPAServerGroupOutputArgs {
    /**
     * The ID of the server group to be exported.
     */
    id?: pulumi.Input<string>;
    /**
     * The name of the server group to be exported.
     */
    name?: pulumi.Input<string>;
}
