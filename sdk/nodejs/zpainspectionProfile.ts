// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "./types/input";
import * as outputs from "./types/output";
import * as utilities from "./utilities";

/**
 * The  **zpa_inspection_profile** resource creates an inspection profile in the Zscaler Private Access cloud. This resource can then be referenced in an inspection custom control resource.
 *
 * ## Import
 *
 * Zscaler offers a dedicated tool called Zscaler-Terraformer to allow the automated import of ZPA configurations into Terraform-compliant HashiCorp Configuration Language. [Visit](https://github.com/zscaler/zscaler-terraformer)
 */
export class ZPAInspectionProfile extends pulumi.CustomResource {
    /**
     * Get an existing ZPAInspectionProfile resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: ZPAInspectionProfileState, opts?: pulumi.CustomResourceOptions): ZPAInspectionProfile {
        return new ZPAInspectionProfile(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'zpa:index/zPAInspectionProfile:ZPAInspectionProfile';

    /**
     * Returns true if the given object is an instance of ZPAInspectionProfile.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is ZPAInspectionProfile {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === ZPAInspectionProfile.__pulumiType;
    }

    public readonly associateAllControls!: pulumi.Output<boolean | undefined>;
    /**
     * (Optional)
     */
    public readonly commonGlobalOverrideActionsConfig!: pulumi.Output<{[key: string]: string}>;
    /**
     * (Optional) Types for custom controls
     */
    public readonly controlsInfos!: pulumi.Output<outputs.ZPAInspectionProfileControlsInfo[]>;
    /**
     * (Optional) Types for custom controls
     * * `type` (Optional) Types for custom controls
     * * `controlRuleJson` (Optional) Custom controls string in JSON format
     */
    public readonly customControls!: pulumi.Output<outputs.ZPAInspectionProfileCustomControl[]>;
    /**
     * Description of the inspection profile.
     */
    public readonly description!: pulumi.Output<string>;
    public readonly globalControlActions!: pulumi.Output<string[]>;
    public readonly incarnationNumber!: pulumi.Output<string>;
    /**
     * The name of the inspection profile.
     */
    public readonly name!: pulumi.Output<string>;
    /**
     * OWASP Predefined Paranoia Level. Range: [1-4], inclusive
     */
    public readonly paranoiaLevel!: pulumi.Output<string>;
    /**
     * The predefined controls. The default predefined control `Preprocessors` are mandatory and injected in the request by default. Individual `predefinedControls` can be set by using the data source `dataSourceZpaPredefinedControls` or by group using the data source `zpa.getZPAInspectionAllPredefinedControls`.
     */
    public readonly predefinedControls!: pulumi.Output<outputs.ZPAInspectionProfilePredefinedControl[]>;
    public readonly predefinedControlsVersion!: pulumi.Output<string>;

    /**
     * Create a ZPAInspectionProfile resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args?: ZPAInspectionProfileArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: ZPAInspectionProfileArgs | ZPAInspectionProfileState, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as ZPAInspectionProfileState | undefined;
            resourceInputs["associateAllControls"] = state ? state.associateAllControls : undefined;
            resourceInputs["commonGlobalOverrideActionsConfig"] = state ? state.commonGlobalOverrideActionsConfig : undefined;
            resourceInputs["controlsInfos"] = state ? state.controlsInfos : undefined;
            resourceInputs["customControls"] = state ? state.customControls : undefined;
            resourceInputs["description"] = state ? state.description : undefined;
            resourceInputs["globalControlActions"] = state ? state.globalControlActions : undefined;
            resourceInputs["incarnationNumber"] = state ? state.incarnationNumber : undefined;
            resourceInputs["name"] = state ? state.name : undefined;
            resourceInputs["paranoiaLevel"] = state ? state.paranoiaLevel : undefined;
            resourceInputs["predefinedControls"] = state ? state.predefinedControls : undefined;
            resourceInputs["predefinedControlsVersion"] = state ? state.predefinedControlsVersion : undefined;
        } else {
            const args = argsOrState as ZPAInspectionProfileArgs | undefined;
            resourceInputs["associateAllControls"] = args ? args.associateAllControls : undefined;
            resourceInputs["commonGlobalOverrideActionsConfig"] = args ? args.commonGlobalOverrideActionsConfig : undefined;
            resourceInputs["controlsInfos"] = args ? args.controlsInfos : undefined;
            resourceInputs["customControls"] = args ? args.customControls : undefined;
            resourceInputs["description"] = args ? args.description : undefined;
            resourceInputs["globalControlActions"] = args ? args.globalControlActions : undefined;
            resourceInputs["incarnationNumber"] = args ? args.incarnationNumber : undefined;
            resourceInputs["name"] = args ? args.name : undefined;
            resourceInputs["paranoiaLevel"] = args ? args.paranoiaLevel : undefined;
            resourceInputs["predefinedControls"] = args ? args.predefinedControls : undefined;
            resourceInputs["predefinedControlsVersion"] = args ? args.predefinedControlsVersion : undefined;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(ZPAInspectionProfile.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering ZPAInspectionProfile resources.
 */
export interface ZPAInspectionProfileState {
    associateAllControls?: pulumi.Input<boolean>;
    /**
     * (Optional)
     */
    commonGlobalOverrideActionsConfig?: pulumi.Input<{[key: string]: pulumi.Input<string>}>;
    /**
     * (Optional) Types for custom controls
     */
    controlsInfos?: pulumi.Input<pulumi.Input<inputs.ZPAInspectionProfileControlsInfo>[]>;
    /**
     * (Optional) Types for custom controls
     * * `type` (Optional) Types for custom controls
     * * `controlRuleJson` (Optional) Custom controls string in JSON format
     */
    customControls?: pulumi.Input<pulumi.Input<inputs.ZPAInspectionProfileCustomControl>[]>;
    /**
     * Description of the inspection profile.
     */
    description?: pulumi.Input<string>;
    globalControlActions?: pulumi.Input<pulumi.Input<string>[]>;
    incarnationNumber?: pulumi.Input<string>;
    /**
     * The name of the inspection profile.
     */
    name?: pulumi.Input<string>;
    /**
     * OWASP Predefined Paranoia Level. Range: [1-4], inclusive
     */
    paranoiaLevel?: pulumi.Input<string>;
    /**
     * The predefined controls. The default predefined control `Preprocessors` are mandatory and injected in the request by default. Individual `predefinedControls` can be set by using the data source `dataSourceZpaPredefinedControls` or by group using the data source `zpa.getZPAInspectionAllPredefinedControls`.
     */
    predefinedControls?: pulumi.Input<pulumi.Input<inputs.ZPAInspectionProfilePredefinedControl>[]>;
    predefinedControlsVersion?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a ZPAInspectionProfile resource.
 */
export interface ZPAInspectionProfileArgs {
    associateAllControls?: pulumi.Input<boolean>;
    /**
     * (Optional)
     */
    commonGlobalOverrideActionsConfig?: pulumi.Input<{[key: string]: pulumi.Input<string>}>;
    /**
     * (Optional) Types for custom controls
     */
    controlsInfos?: pulumi.Input<pulumi.Input<inputs.ZPAInspectionProfileControlsInfo>[]>;
    /**
     * (Optional) Types for custom controls
     * * `type` (Optional) Types for custom controls
     * * `controlRuleJson` (Optional) Custom controls string in JSON format
     */
    customControls?: pulumi.Input<pulumi.Input<inputs.ZPAInspectionProfileCustomControl>[]>;
    /**
     * Description of the inspection profile.
     */
    description?: pulumi.Input<string>;
    globalControlActions?: pulumi.Input<pulumi.Input<string>[]>;
    incarnationNumber?: pulumi.Input<string>;
    /**
     * The name of the inspection profile.
     */
    name?: pulumi.Input<string>;
    /**
     * OWASP Predefined Paranoia Level. Range: [1-4], inclusive
     */
    paranoiaLevel?: pulumi.Input<string>;
    /**
     * The predefined controls. The default predefined control `Preprocessors` are mandatory and injected in the request by default. Individual `predefinedControls` can be set by using the data source `dataSourceZpaPredefinedControls` or by group using the data source `zpa.getZPAInspectionAllPredefinedControls`.
     */
    predefinedControls?: pulumi.Input<pulumi.Input<inputs.ZPAInspectionProfilePredefinedControl>[]>;
    predefinedControlsVersion?: pulumi.Input<string>;
}
