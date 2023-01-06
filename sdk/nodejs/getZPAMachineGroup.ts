// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "./types/input";
import * as outputs from "./types/output";
import * as utilities from "./utilities";

/**
 * Use the **zpa_machine_group** data source to get information about a machine group created in the Zscaler Private Access cloud. This data source can then be referenced in an Access Policy, Timeout policy, Forwarding Policy, Inspection Policy or Isolation Policy.
 *
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as zpa from "@pulumi/zpa";
 *
 * const example = zpa.getZPAMachineGroup({
 *     name: "MGR01",
 * });
 * ```
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as zpa from "@pulumi/zpa";
 *
 * const example = zpa.getZPAMachineGroup({
 *     id: "1234567890",
 * });
 * ```
 */
export function getZPAMachineGroup(args?: GetZPAMachineGroupArgs, opts?: pulumi.InvokeOptions): Promise<GetZPAMachineGroupResult> {
    args = args || {};

    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts || {});
    return pulumi.runtime.invoke("zpa:index/getZPAMachineGroup:getZPAMachineGroup", {
        "id": args.id,
        "name": args.name,
    }, opts);
}

/**
 * A collection of arguments for invoking getZPAMachineGroup.
 */
export interface GetZPAMachineGroupArgs {
    /**
     * The ID of the machine group to be exported.
     */
    id?: string;
    /**
     * The name of the machine group to be exported.
     */
    name?: string;
}

/**
 * A collection of values returned by getZPAMachineGroup.
 */
export interface GetZPAMachineGroupResult {
    /**
     * (string)
     */
    readonly creationTime: string;
    /**
     * (string)
     */
    readonly description: string;
    /**
     * (bool)
     */
    readonly enabled: boolean;
    /**
     * (string)
     */
    readonly id?: string;
    /**
     * (string)
     */
    readonly machines: outputs.GetZPAMachineGroupMachine[];
    /**
     * (string)
     */
    readonly modifiedBy: string;
    /**
     * (string)
     */
    readonly modifiedTime: string;
    /**
     * (string)
     */
    readonly name?: string;
}
/**
 * Use the **zpa_machine_group** data source to get information about a machine group created in the Zscaler Private Access cloud. This data source can then be referenced in an Access Policy, Timeout policy, Forwarding Policy, Inspection Policy or Isolation Policy.
 *
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as zpa from "@pulumi/zpa";
 *
 * const example = zpa.getZPAMachineGroup({
 *     name: "MGR01",
 * });
 * ```
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as zpa from "@pulumi/zpa";
 *
 * const example = zpa.getZPAMachineGroup({
 *     id: "1234567890",
 * });
 * ```
 */
export function getZPAMachineGroupOutput(args?: GetZPAMachineGroupOutputArgs, opts?: pulumi.InvokeOptions): pulumi.Output<GetZPAMachineGroupResult> {
    return pulumi.output(args).apply((a: any) => getZPAMachineGroup(a, opts))
}

/**
 * A collection of arguments for invoking getZPAMachineGroup.
 */
export interface GetZPAMachineGroupOutputArgs {
    /**
     * The ID of the machine group to be exported.
     */
    id?: pulumi.Input<string>;
    /**
     * The name of the machine group to be exported.
     */
    name?: pulumi.Input<string>;
}
