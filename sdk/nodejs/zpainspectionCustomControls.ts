// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "./types/input";
import * as outputs from "./types/output";
import * as utilities from "./utilities";

export class ZPAInspectionCustomControls extends pulumi.CustomResource {
    /**
     * Get an existing ZPAInspectionCustomControls resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: ZPAInspectionCustomControlsState, opts?: pulumi.CustomResourceOptions): ZPAInspectionCustomControls {
        return new ZPAInspectionCustomControls(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'zpa:index/zPAInspectionCustomControls:ZPAInspectionCustomControls';

    /**
     * Returns true if the given object is an instance of ZPAInspectionCustomControls.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is ZPAInspectionCustomControls {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === ZPAInspectionCustomControls.__pulumiType;
    }

    /**
     * The performed action
     */
    public readonly action!: pulumi.Output<string>;
    public readonly actionValue!: pulumi.Output<string>;
    /**
     * Name of the inspection profile
     */
    public readonly associatedInspectionProfileNames!: pulumi.Output<outputs.ZPAInspectionCustomControlsAssociatedInspectionProfileName[]>;
    public readonly controlNumber!: pulumi.Output<string>;
    /**
     * The control rule in JSON format that has the conditions and type of control for the inspection control
     */
    public readonly controlRuleJson!: pulumi.Output<string>;
    /**
     * The performed action
     */
    public readonly defaultAction!: pulumi.Output<string>;
    /**
     * This is used to provide the redirect URL if the default action is set to REDIRECT
     */
    public readonly defaultActionValue!: pulumi.Output<string>;
    /**
     * Description of the custom control
     */
    public readonly description!: pulumi.Output<string>;
    public readonly name!: pulumi.Output<string>;
    /**
     * OWASP Predefined Paranoia Level. Range: [1-4], inclusive
     */
    public readonly paranoiaLevel!: pulumi.Output<string>;
    /**
     * Rules of the custom controls applied as conditions (JSON)
     */
    public readonly rules!: pulumi.Output<outputs.ZPAInspectionCustomControlsRule[]>;
    /**
     * Severity of the control number
     */
    public readonly severity!: pulumi.Output<string>;
    /**
     * Rules to be applied to the request or response type
     */
    public readonly type!: pulumi.Output<string>;
    public readonly version!: pulumi.Output<string>;

    /**
     * Create a ZPAInspectionCustomControls resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: ZPAInspectionCustomControlsArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: ZPAInspectionCustomControlsArgs | ZPAInspectionCustomControlsState, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as ZPAInspectionCustomControlsState | undefined;
            resourceInputs["action"] = state ? state.action : undefined;
            resourceInputs["actionValue"] = state ? state.actionValue : undefined;
            resourceInputs["associatedInspectionProfileNames"] = state ? state.associatedInspectionProfileNames : undefined;
            resourceInputs["controlNumber"] = state ? state.controlNumber : undefined;
            resourceInputs["controlRuleJson"] = state ? state.controlRuleJson : undefined;
            resourceInputs["defaultAction"] = state ? state.defaultAction : undefined;
            resourceInputs["defaultActionValue"] = state ? state.defaultActionValue : undefined;
            resourceInputs["description"] = state ? state.description : undefined;
            resourceInputs["name"] = state ? state.name : undefined;
            resourceInputs["paranoiaLevel"] = state ? state.paranoiaLevel : undefined;
            resourceInputs["rules"] = state ? state.rules : undefined;
            resourceInputs["severity"] = state ? state.severity : undefined;
            resourceInputs["type"] = state ? state.type : undefined;
            resourceInputs["version"] = state ? state.version : undefined;
        } else {
            const args = argsOrState as ZPAInspectionCustomControlsArgs | undefined;
            if ((!args || args.defaultAction === undefined) && !opts.urn) {
                throw new Error("Missing required property 'defaultAction'");
            }
            if ((!args || args.severity === undefined) && !opts.urn) {
                throw new Error("Missing required property 'severity'");
            }
            if ((!args || args.type === undefined) && !opts.urn) {
                throw new Error("Missing required property 'type'");
            }
            resourceInputs["action"] = args ? args.action : undefined;
            resourceInputs["actionValue"] = args ? args.actionValue : undefined;
            resourceInputs["associatedInspectionProfileNames"] = args ? args.associatedInspectionProfileNames : undefined;
            resourceInputs["controlNumber"] = args ? args.controlNumber : undefined;
            resourceInputs["controlRuleJson"] = args ? args.controlRuleJson : undefined;
            resourceInputs["defaultAction"] = args ? args.defaultAction : undefined;
            resourceInputs["defaultActionValue"] = args ? args.defaultActionValue : undefined;
            resourceInputs["description"] = args ? args.description : undefined;
            resourceInputs["name"] = args ? args.name : undefined;
            resourceInputs["paranoiaLevel"] = args ? args.paranoiaLevel : undefined;
            resourceInputs["rules"] = args ? args.rules : undefined;
            resourceInputs["severity"] = args ? args.severity : undefined;
            resourceInputs["type"] = args ? args.type : undefined;
            resourceInputs["version"] = args ? args.version : undefined;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(ZPAInspectionCustomControls.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering ZPAInspectionCustomControls resources.
 */
export interface ZPAInspectionCustomControlsState {
    /**
     * The performed action
     */
    action?: pulumi.Input<string>;
    actionValue?: pulumi.Input<string>;
    /**
     * Name of the inspection profile
     */
    associatedInspectionProfileNames?: pulumi.Input<pulumi.Input<inputs.ZPAInspectionCustomControlsAssociatedInspectionProfileName>[]>;
    controlNumber?: pulumi.Input<string>;
    /**
     * The control rule in JSON format that has the conditions and type of control for the inspection control
     */
    controlRuleJson?: pulumi.Input<string>;
    /**
     * The performed action
     */
    defaultAction?: pulumi.Input<string>;
    /**
     * This is used to provide the redirect URL if the default action is set to REDIRECT
     */
    defaultActionValue?: pulumi.Input<string>;
    /**
     * Description of the custom control
     */
    description?: pulumi.Input<string>;
    name?: pulumi.Input<string>;
    /**
     * OWASP Predefined Paranoia Level. Range: [1-4], inclusive
     */
    paranoiaLevel?: pulumi.Input<string>;
    /**
     * Rules of the custom controls applied as conditions (JSON)
     */
    rules?: pulumi.Input<pulumi.Input<inputs.ZPAInspectionCustomControlsRule>[]>;
    /**
     * Severity of the control number
     */
    severity?: pulumi.Input<string>;
    /**
     * Rules to be applied to the request or response type
     */
    type?: pulumi.Input<string>;
    version?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a ZPAInspectionCustomControls resource.
 */
export interface ZPAInspectionCustomControlsArgs {
    /**
     * The performed action
     */
    action?: pulumi.Input<string>;
    actionValue?: pulumi.Input<string>;
    /**
     * Name of the inspection profile
     */
    associatedInspectionProfileNames?: pulumi.Input<pulumi.Input<inputs.ZPAInspectionCustomControlsAssociatedInspectionProfileName>[]>;
    controlNumber?: pulumi.Input<string>;
    /**
     * The control rule in JSON format that has the conditions and type of control for the inspection control
     */
    controlRuleJson?: pulumi.Input<string>;
    /**
     * The performed action
     */
    defaultAction: pulumi.Input<string>;
    /**
     * This is used to provide the redirect URL if the default action is set to REDIRECT
     */
    defaultActionValue?: pulumi.Input<string>;
    /**
     * Description of the custom control
     */
    description?: pulumi.Input<string>;
    name?: pulumi.Input<string>;
    /**
     * OWASP Predefined Paranoia Level. Range: [1-4], inclusive
     */
    paranoiaLevel?: pulumi.Input<string>;
    /**
     * Rules of the custom controls applied as conditions (JSON)
     */
    rules?: pulumi.Input<pulumi.Input<inputs.ZPAInspectionCustomControlsRule>[]>;
    /**
     * Severity of the control number
     */
    severity: pulumi.Input<string>;
    /**
     * Rules to be applied to the request or response type
     */
    type: pulumi.Input<string>;
    version?: pulumi.Input<string>;
}
