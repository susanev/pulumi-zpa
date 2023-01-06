// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "./types/input";
import * as outputs from "./types/output";
import * as utilities from "./utilities";

/**
 * ## Import
 *
 * Zscaler offers a dedicated tool called Zscaler-Terraformer to allow the automated import of ZPA configurations into Terraform-compliant HashiCorp Configuration Language. [Visit](https://github.com/zscaler/zscaler-terraformer) Policy access rule can be imported by using `<POLICY ACCESS RULE ID>` as the import ID. For example
 *
 * ```sh
 *  $ pulumi import zpa:index/zPAPolicyAccessRule:ZPAPolicyAccessRule example <policy_access_rule_id>
 * ```
 */
export class ZPAPolicyAccessRule extends pulumi.CustomResource {
    /**
     * Get an existing ZPAPolicyAccessRule resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: ZPAPolicyAccessRuleState, opts?: pulumi.CustomResourceOptions): ZPAPolicyAccessRule {
        return new ZPAPolicyAccessRule(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'zpa:index/zPAPolicyAccessRule:ZPAPolicyAccessRule';

    /**
     * Returns true if the given object is an instance of ZPAPolicyAccessRule.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is ZPAPolicyAccessRule {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === ZPAPolicyAccessRule.__pulumiType;
    }

    /**
     * This is for providing the rule action.
     */
    public readonly action!: pulumi.Output<string | undefined>;
    /**
     * This field defines the description of the server.
     */
    public readonly actionId!: pulumi.Output<string | undefined>;
    /**
     * List of app-connector IDs.
     */
    public readonly appConnectorGroups!: pulumi.Output<outputs.ZPAPolicyAccessRuleAppConnectorGroup[]>;
    /**
     * List of the server group IDs.
     */
    public readonly appServerGroups!: pulumi.Output<outputs.ZPAPolicyAccessRuleAppServerGroup[]>;
    public readonly bypassDefaultRule!: pulumi.Output<boolean | undefined>;
    /**
     * (Optional)
     */
    public readonly conditions!: pulumi.Output<outputs.ZPAPolicyAccessRuleCondition[]>;
    /**
     * This is for providing a customer message for the user.
     */
    public readonly customMsg!: pulumi.Output<string | undefined>;
    /**
     * This is for providing a customer message for the user.
     */
    public readonly defaultRule!: pulumi.Output<boolean | undefined>;
    /**
     * This is the description of the access policy.
     */
    public readonly description!: pulumi.Output<string | undefined>;
    public readonly lssDefaultRule!: pulumi.Output<boolean | undefined>;
    /**
     * This is the name of the policy.
     */
    public readonly name!: pulumi.Output<string>;
    public readonly operator!: pulumi.Output<string>;
    public readonly policySetId!: pulumi.Output<string>;
    public readonly policyType!: pulumi.Output<string>;
    public readonly priority!: pulumi.Output<string>;
    public readonly reauthDefaultRule!: pulumi.Output<boolean | undefined>;
    public readonly reauthIdleTimeout!: pulumi.Output<string | undefined>;
    public readonly reauthTimeout!: pulumi.Output<string | undefined>;
    public readonly ruleOrder!: pulumi.Output<string>;
    public readonly zpnInspectionProfileId!: pulumi.Output<string | undefined>;

    /**
     * Create a ZPAPolicyAccessRule resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args?: ZPAPolicyAccessRuleArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: ZPAPolicyAccessRuleArgs | ZPAPolicyAccessRuleState, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as ZPAPolicyAccessRuleState | undefined;
            resourceInputs["action"] = state ? state.action : undefined;
            resourceInputs["actionId"] = state ? state.actionId : undefined;
            resourceInputs["appConnectorGroups"] = state ? state.appConnectorGroups : undefined;
            resourceInputs["appServerGroups"] = state ? state.appServerGroups : undefined;
            resourceInputs["bypassDefaultRule"] = state ? state.bypassDefaultRule : undefined;
            resourceInputs["conditions"] = state ? state.conditions : undefined;
            resourceInputs["customMsg"] = state ? state.customMsg : undefined;
            resourceInputs["defaultRule"] = state ? state.defaultRule : undefined;
            resourceInputs["description"] = state ? state.description : undefined;
            resourceInputs["lssDefaultRule"] = state ? state.lssDefaultRule : undefined;
            resourceInputs["name"] = state ? state.name : undefined;
            resourceInputs["operator"] = state ? state.operator : undefined;
            resourceInputs["policySetId"] = state ? state.policySetId : undefined;
            resourceInputs["policyType"] = state ? state.policyType : undefined;
            resourceInputs["priority"] = state ? state.priority : undefined;
            resourceInputs["reauthDefaultRule"] = state ? state.reauthDefaultRule : undefined;
            resourceInputs["reauthIdleTimeout"] = state ? state.reauthIdleTimeout : undefined;
            resourceInputs["reauthTimeout"] = state ? state.reauthTimeout : undefined;
            resourceInputs["ruleOrder"] = state ? state.ruleOrder : undefined;
            resourceInputs["zpnInspectionProfileId"] = state ? state.zpnInspectionProfileId : undefined;
        } else {
            const args = argsOrState as ZPAPolicyAccessRuleArgs | undefined;
            resourceInputs["action"] = args ? args.action : undefined;
            resourceInputs["actionId"] = args ? args.actionId : undefined;
            resourceInputs["appConnectorGroups"] = args ? args.appConnectorGroups : undefined;
            resourceInputs["appServerGroups"] = args ? args.appServerGroups : undefined;
            resourceInputs["bypassDefaultRule"] = args ? args.bypassDefaultRule : undefined;
            resourceInputs["conditions"] = args ? args.conditions : undefined;
            resourceInputs["customMsg"] = args ? args.customMsg : undefined;
            resourceInputs["defaultRule"] = args ? args.defaultRule : undefined;
            resourceInputs["description"] = args ? args.description : undefined;
            resourceInputs["lssDefaultRule"] = args ? args.lssDefaultRule : undefined;
            resourceInputs["name"] = args ? args.name : undefined;
            resourceInputs["operator"] = args ? args.operator : undefined;
            resourceInputs["policySetId"] = args ? args.policySetId : undefined;
            resourceInputs["policyType"] = args ? args.policyType : undefined;
            resourceInputs["priority"] = args ? args.priority : undefined;
            resourceInputs["reauthDefaultRule"] = args ? args.reauthDefaultRule : undefined;
            resourceInputs["reauthIdleTimeout"] = args ? args.reauthIdleTimeout : undefined;
            resourceInputs["reauthTimeout"] = args ? args.reauthTimeout : undefined;
            resourceInputs["ruleOrder"] = args ? args.ruleOrder : undefined;
            resourceInputs["zpnInspectionProfileId"] = args ? args.zpnInspectionProfileId : undefined;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(ZPAPolicyAccessRule.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering ZPAPolicyAccessRule resources.
 */
export interface ZPAPolicyAccessRuleState {
    /**
     * This is for providing the rule action.
     */
    action?: pulumi.Input<string>;
    /**
     * This field defines the description of the server.
     */
    actionId?: pulumi.Input<string>;
    /**
     * List of app-connector IDs.
     */
    appConnectorGroups?: pulumi.Input<pulumi.Input<inputs.ZPAPolicyAccessRuleAppConnectorGroup>[]>;
    /**
     * List of the server group IDs.
     */
    appServerGroups?: pulumi.Input<pulumi.Input<inputs.ZPAPolicyAccessRuleAppServerGroup>[]>;
    bypassDefaultRule?: pulumi.Input<boolean>;
    /**
     * (Optional)
     */
    conditions?: pulumi.Input<pulumi.Input<inputs.ZPAPolicyAccessRuleCondition>[]>;
    /**
     * This is for providing a customer message for the user.
     */
    customMsg?: pulumi.Input<string>;
    /**
     * This is for providing a customer message for the user.
     */
    defaultRule?: pulumi.Input<boolean>;
    /**
     * This is the description of the access policy.
     */
    description?: pulumi.Input<string>;
    lssDefaultRule?: pulumi.Input<boolean>;
    /**
     * This is the name of the policy.
     */
    name?: pulumi.Input<string>;
    operator?: pulumi.Input<string>;
    policySetId?: pulumi.Input<string>;
    policyType?: pulumi.Input<string>;
    priority?: pulumi.Input<string>;
    reauthDefaultRule?: pulumi.Input<boolean>;
    reauthIdleTimeout?: pulumi.Input<string>;
    reauthTimeout?: pulumi.Input<string>;
    ruleOrder?: pulumi.Input<string>;
    zpnInspectionProfileId?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a ZPAPolicyAccessRule resource.
 */
export interface ZPAPolicyAccessRuleArgs {
    /**
     * This is for providing the rule action.
     */
    action?: pulumi.Input<string>;
    /**
     * This field defines the description of the server.
     */
    actionId?: pulumi.Input<string>;
    /**
     * List of app-connector IDs.
     */
    appConnectorGroups?: pulumi.Input<pulumi.Input<inputs.ZPAPolicyAccessRuleAppConnectorGroup>[]>;
    /**
     * List of the server group IDs.
     */
    appServerGroups?: pulumi.Input<pulumi.Input<inputs.ZPAPolicyAccessRuleAppServerGroup>[]>;
    bypassDefaultRule?: pulumi.Input<boolean>;
    /**
     * (Optional)
     */
    conditions?: pulumi.Input<pulumi.Input<inputs.ZPAPolicyAccessRuleCondition>[]>;
    /**
     * This is for providing a customer message for the user.
     */
    customMsg?: pulumi.Input<string>;
    /**
     * This is for providing a customer message for the user.
     */
    defaultRule?: pulumi.Input<boolean>;
    /**
     * This is the description of the access policy.
     */
    description?: pulumi.Input<string>;
    lssDefaultRule?: pulumi.Input<boolean>;
    /**
     * This is the name of the policy.
     */
    name?: pulumi.Input<string>;
    operator?: pulumi.Input<string>;
    policySetId?: pulumi.Input<string>;
    policyType?: pulumi.Input<string>;
    priority?: pulumi.Input<string>;
    reauthDefaultRule?: pulumi.Input<boolean>;
    reauthIdleTimeout?: pulumi.Input<string>;
    reauthTimeout?: pulumi.Input<string>;
    ruleOrder?: pulumi.Input<string>;
    zpnInspectionProfileId?: pulumi.Input<string>;
}
