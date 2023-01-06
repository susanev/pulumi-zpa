// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "./types/input";
import * as outputs from "./types/output";
import * as utilities from "./utilities";

/**
 * Use the **zpa_inspection_predefined_controls** data source to get information about an OWASP predefined control and prefedined control version. This data source is required when creating an inspection profile.
 *
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as zpa from "@pulumi/zpa";
 *
 * const example = zpa.getZPAInspectionPredefinedControls({
 *     name: "Failed to parse request body",
 *     version: "OWASP_CRS/3.3.0",
 * });
 * export const zpaInspectionPredefinedControls = example;
 * ```
 */
export function getZPAInspectionPredefinedControls(args?: GetZPAInspectionPredefinedControlsArgs, opts?: pulumi.InvokeOptions): Promise<GetZPAInspectionPredefinedControlsResult> {
    args = args || {};

    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts || {});
    return pulumi.runtime.invoke("zpa:index/getZPAInspectionPredefinedControls:getZPAInspectionPredefinedControls", {
        "id": args.id,
        "name": args.name,
        "version": args.version,
    }, opts);
}

/**
 * A collection of arguments for invoking getZPAInspectionPredefinedControls.
 */
export interface GetZPAInspectionPredefinedControlsArgs {
    /**
     * (Computed)
     */
    id?: string;
    /**
     * The name of the predefined control.
     */
    name?: string;
    /**
     * The version of the predefined control, the default is: `OWASP_CRS/3.3.0`
     */
    version?: string;
}

/**
 * A collection of values returned by getZPAInspectionPredefinedControls.
 */
export interface GetZPAInspectionPredefinedControlsResult {
    /**
     * (Computed)
     */
    readonly action: string;
    /**
     * (Computed)
     */
    readonly actionValue: string;
    /**
     * (Computed)
     * * `id`- (Computed)
     * * `name`- (Computed)
     */
    readonly associatedInspectionProfileNames: outputs.GetZPAInspectionPredefinedControlsAssociatedInspectionProfileName[];
    /**
     * (Computed)
     */
    readonly attachment: string;
    /**
     * (Computed)
     */
    readonly controlGroup: string;
    /**
     * (Computed)
     */
    readonly controlNumber: string;
    /**
     * (Computed)
     */
    readonly creationTime: string;
    /**
     * (Computed)
     */
    readonly defaultAction: string;
    /**
     * (Computed)
     */
    readonly defaultActionValue: string;
    /**
     * (Computed)
     */
    readonly description: string;
    /**
     * (Computed)
     */
    readonly id: string;
    /**
     * (Computed)
     */
    readonly modifiedTime: string;
    readonly modifiedby: string;
    readonly name: string;
    /**
     * (Computed)
     */
    readonly paranoiaLevel: string;
    /**
     * (Computed)
     */
    readonly severity: string;
    readonly version?: string;
}
/**
 * Use the **zpa_inspection_predefined_controls** data source to get information about an OWASP predefined control and prefedined control version. This data source is required when creating an inspection profile.
 *
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as zpa from "@pulumi/zpa";
 *
 * const example = zpa.getZPAInspectionPredefinedControls({
 *     name: "Failed to parse request body",
 *     version: "OWASP_CRS/3.3.0",
 * });
 * export const zpaInspectionPredefinedControls = example;
 * ```
 */
export function getZPAInspectionPredefinedControlsOutput(args?: GetZPAInspectionPredefinedControlsOutputArgs, opts?: pulumi.InvokeOptions): pulumi.Output<GetZPAInspectionPredefinedControlsResult> {
    return pulumi.output(args).apply((a: any) => getZPAInspectionPredefinedControls(a, opts))
}

/**
 * A collection of arguments for invoking getZPAInspectionPredefinedControls.
 */
export interface GetZPAInspectionPredefinedControlsOutputArgs {
    /**
     * (Computed)
     */
    id?: pulumi.Input<string>;
    /**
     * The name of the predefined control.
     */
    name?: pulumi.Input<string>;
    /**
     * The version of the predefined control, the default is: `OWASP_CRS/3.3.0`
     */
    version?: pulumi.Input<string>;
}
