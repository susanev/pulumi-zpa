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
    /// ## Import
    /// 
    /// Zscaler offers a dedicated tool called Zscaler-Terraformer to allow the automated import of ZPA configurations into Terraform-compliant HashiCorp Configuration Language. [Visit](https://github.com/zscaler/zscaler-terraformer) Policy access rule can be imported by using `&lt;POLICY ACCESS RULE ID&gt;` as the import ID. For example
    /// 
    /// ```sh
    ///  $ pulumi import zpa:index/zPAPolicyAccessRule:ZPAPolicyAccessRule example &lt;policy_access_rule_id&gt;
    /// ```
    /// </summary>
    [ZpaResourceType("zpa:index/zPAPolicyAccessRule:ZPAPolicyAccessRule")]
    public partial class ZPAPolicyAccessRule : global::Pulumi.CustomResource
    {
        /// <summary>
        /// (Optional) This is for providing the rule action. Supported values: ``ALLOW``, ``DENY``
        /// </summary>
        [Output("action")]
        public Output<string?> Action { get; private set; } = null!;

        /// <summary>
        /// This field defines the description of the server.
        /// </summary>
        [Output("actionId")]
        public Output<string?> ActionId { get; private set; } = null!;

        /// <summary>
        /// List of app-connector IDs.
        /// </summary>
        [Output("appConnectorGroups")]
        public Output<ImmutableArray<Outputs.ZPAPolicyAccessRuleAppConnectorGroup>> AppConnectorGroups { get; private set; } = null!;

        /// <summary>
        /// List of the server group IDs.
        /// </summary>
        [Output("appServerGroups")]
        public Output<ImmutableArray<Outputs.ZPAPolicyAccessRuleAppServerGroup>> AppServerGroups { get; private set; } = null!;

        [Output("bypassDefaultRule")]
        public Output<bool?> BypassDefaultRule { get; private set; } = null!;

        /// <summary>
        /// (Optional)
        /// </summary>
        [Output("conditions")]
        public Output<ImmutableArray<Outputs.ZPAPolicyAccessRuleCondition>> Conditions { get; private set; } = null!;

        /// <summary>
        /// (Optional) This is for providing a customer message for the user.
        /// </summary>
        [Output("customMsg")]
        public Output<string?> CustomMsg { get; private set; } = null!;

        /// <summary>
        /// This is for providing a customer message for the user.
        /// </summary>
        [Output("defaultRule")]
        public Output<bool?> DefaultRule { get; private set; } = null!;

        /// <summary>
        /// (Optional) This is the description of the access policy rule.
        /// </summary>
        [Output("description")]
        public Output<string?> Description { get; private set; } = null!;

        [Output("lssDefaultRule")]
        public Output<bool?> LssDefaultRule { get; private set; } = null!;

        /// <summary>
        /// (Optional)
        /// </summary>
        [Output("name")]
        public Output<string> Name { get; private set; } = null!;

        /// <summary>
        /// (Optional) Supported values: ``AND``, and ``OR``
        /// </summary>
        [Output("operator")]
        public Output<string> Operator { get; private set; } = null!;

        [Output("policySetId")]
        public Output<string> PolicySetId { get; private set; } = null!;

        /// <summary>
        /// (Optional) Supported values: ``ACCESS_POLICY`` or ``GLOBAL_POLICY``
        /// </summary>
        [Output("policyType")]
        public Output<string> PolicyType { get; private set; } = null!;

        [Output("priority")]
        public Output<string> Priority { get; private set; } = null!;

        [Output("reauthDefaultRule")]
        public Output<bool?> ReauthDefaultRule { get; private set; } = null!;

        [Output("reauthIdleTimeout")]
        public Output<string?> ReauthIdleTimeout { get; private set; } = null!;

        [Output("reauthTimeout")]
        public Output<string?> ReauthTimeout { get; private set; } = null!;

        /// <summary>
        /// (Optional)
        /// </summary>
        [Output("ruleOrder")]
        public Output<string> RuleOrder { get; private set; } = null!;

        [Output("zpnInspectionProfileId")]
        public Output<string?> ZpnInspectionProfileId { get; private set; } = null!;


        /// <summary>
        /// Create a ZPAPolicyAccessRule resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public ZPAPolicyAccessRule(string name, ZPAPolicyAccessRuleArgs? args = null, CustomResourceOptions? options = null)
            : base("zpa:index/zPAPolicyAccessRule:ZPAPolicyAccessRule", name, args ?? new ZPAPolicyAccessRuleArgs(), MakeResourceOptions(options, ""))
        {
        }

        private ZPAPolicyAccessRule(string name, Input<string> id, ZPAPolicyAccessRuleState? state = null, CustomResourceOptions? options = null)
            : base("zpa:index/zPAPolicyAccessRule:ZPAPolicyAccessRule", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing ZPAPolicyAccessRule resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static ZPAPolicyAccessRule Get(string name, Input<string> id, ZPAPolicyAccessRuleState? state = null, CustomResourceOptions? options = null)
        {
            return new ZPAPolicyAccessRule(name, id, state, options);
        }
    }

    public sealed class ZPAPolicyAccessRuleArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// (Optional) This is for providing the rule action. Supported values: ``ALLOW``, ``DENY``
        /// </summary>
        [Input("action")]
        public Input<string>? Action { get; set; }

        /// <summary>
        /// This field defines the description of the server.
        /// </summary>
        [Input("actionId")]
        public Input<string>? ActionId { get; set; }

        [Input("appConnectorGroups")]
        private InputList<Inputs.ZPAPolicyAccessRuleAppConnectorGroupArgs>? _appConnectorGroups;

        /// <summary>
        /// List of app-connector IDs.
        /// </summary>
        public InputList<Inputs.ZPAPolicyAccessRuleAppConnectorGroupArgs> AppConnectorGroups
        {
            get => _appConnectorGroups ?? (_appConnectorGroups = new InputList<Inputs.ZPAPolicyAccessRuleAppConnectorGroupArgs>());
            set => _appConnectorGroups = value;
        }

        [Input("appServerGroups")]
        private InputList<Inputs.ZPAPolicyAccessRuleAppServerGroupArgs>? _appServerGroups;

        /// <summary>
        /// List of the server group IDs.
        /// </summary>
        public InputList<Inputs.ZPAPolicyAccessRuleAppServerGroupArgs> AppServerGroups
        {
            get => _appServerGroups ?? (_appServerGroups = new InputList<Inputs.ZPAPolicyAccessRuleAppServerGroupArgs>());
            set => _appServerGroups = value;
        }

        [Input("bypassDefaultRule")]
        public Input<bool>? BypassDefaultRule { get; set; }

        [Input("conditions")]
        private InputList<Inputs.ZPAPolicyAccessRuleConditionArgs>? _conditions;

        /// <summary>
        /// (Optional)
        /// </summary>
        public InputList<Inputs.ZPAPolicyAccessRuleConditionArgs> Conditions
        {
            get => _conditions ?? (_conditions = new InputList<Inputs.ZPAPolicyAccessRuleConditionArgs>());
            set => _conditions = value;
        }

        /// <summary>
        /// (Optional) This is for providing a customer message for the user.
        /// </summary>
        [Input("customMsg")]
        public Input<string>? CustomMsg { get; set; }

        /// <summary>
        /// This is for providing a customer message for the user.
        /// </summary>
        [Input("defaultRule")]
        public Input<bool>? DefaultRule { get; set; }

        /// <summary>
        /// (Optional) This is the description of the access policy rule.
        /// </summary>
        [Input("description")]
        public Input<string>? Description { get; set; }

        [Input("lssDefaultRule")]
        public Input<bool>? LssDefaultRule { get; set; }

        /// <summary>
        /// (Optional)
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// (Optional) Supported values: ``AND``, and ``OR``
        /// </summary>
        [Input("operator")]
        public Input<string>? Operator { get; set; }

        [Input("policySetId")]
        public Input<string>? PolicySetId { get; set; }

        /// <summary>
        /// (Optional) Supported values: ``ACCESS_POLICY`` or ``GLOBAL_POLICY``
        /// </summary>
        [Input("policyType")]
        public Input<string>? PolicyType { get; set; }

        [Input("priority")]
        public Input<string>? Priority { get; set; }

        [Input("reauthDefaultRule")]
        public Input<bool>? ReauthDefaultRule { get; set; }

        [Input("reauthIdleTimeout")]
        public Input<string>? ReauthIdleTimeout { get; set; }

        [Input("reauthTimeout")]
        public Input<string>? ReauthTimeout { get; set; }

        /// <summary>
        /// (Optional)
        /// </summary>
        [Input("ruleOrder")]
        public Input<string>? RuleOrder { get; set; }

        [Input("zpnInspectionProfileId")]
        public Input<string>? ZpnInspectionProfileId { get; set; }

        public ZPAPolicyAccessRuleArgs()
        {
        }
        public static new ZPAPolicyAccessRuleArgs Empty => new ZPAPolicyAccessRuleArgs();
    }

    public sealed class ZPAPolicyAccessRuleState : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// (Optional) This is for providing the rule action. Supported values: ``ALLOW``, ``DENY``
        /// </summary>
        [Input("action")]
        public Input<string>? Action { get; set; }

        /// <summary>
        /// This field defines the description of the server.
        /// </summary>
        [Input("actionId")]
        public Input<string>? ActionId { get; set; }

        [Input("appConnectorGroups")]
        private InputList<Inputs.ZPAPolicyAccessRuleAppConnectorGroupGetArgs>? _appConnectorGroups;

        /// <summary>
        /// List of app-connector IDs.
        /// </summary>
        public InputList<Inputs.ZPAPolicyAccessRuleAppConnectorGroupGetArgs> AppConnectorGroups
        {
            get => _appConnectorGroups ?? (_appConnectorGroups = new InputList<Inputs.ZPAPolicyAccessRuleAppConnectorGroupGetArgs>());
            set => _appConnectorGroups = value;
        }

        [Input("appServerGroups")]
        private InputList<Inputs.ZPAPolicyAccessRuleAppServerGroupGetArgs>? _appServerGroups;

        /// <summary>
        /// List of the server group IDs.
        /// </summary>
        public InputList<Inputs.ZPAPolicyAccessRuleAppServerGroupGetArgs> AppServerGroups
        {
            get => _appServerGroups ?? (_appServerGroups = new InputList<Inputs.ZPAPolicyAccessRuleAppServerGroupGetArgs>());
            set => _appServerGroups = value;
        }

        [Input("bypassDefaultRule")]
        public Input<bool>? BypassDefaultRule { get; set; }

        [Input("conditions")]
        private InputList<Inputs.ZPAPolicyAccessRuleConditionGetArgs>? _conditions;

        /// <summary>
        /// (Optional)
        /// </summary>
        public InputList<Inputs.ZPAPolicyAccessRuleConditionGetArgs> Conditions
        {
            get => _conditions ?? (_conditions = new InputList<Inputs.ZPAPolicyAccessRuleConditionGetArgs>());
            set => _conditions = value;
        }

        /// <summary>
        /// (Optional) This is for providing a customer message for the user.
        /// </summary>
        [Input("customMsg")]
        public Input<string>? CustomMsg { get; set; }

        /// <summary>
        /// This is for providing a customer message for the user.
        /// </summary>
        [Input("defaultRule")]
        public Input<bool>? DefaultRule { get; set; }

        /// <summary>
        /// (Optional) This is the description of the access policy rule.
        /// </summary>
        [Input("description")]
        public Input<string>? Description { get; set; }

        [Input("lssDefaultRule")]
        public Input<bool>? LssDefaultRule { get; set; }

        /// <summary>
        /// (Optional)
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// (Optional) Supported values: ``AND``, and ``OR``
        /// </summary>
        [Input("operator")]
        public Input<string>? Operator { get; set; }

        [Input("policySetId")]
        public Input<string>? PolicySetId { get; set; }

        /// <summary>
        /// (Optional) Supported values: ``ACCESS_POLICY`` or ``GLOBAL_POLICY``
        /// </summary>
        [Input("policyType")]
        public Input<string>? PolicyType { get; set; }

        [Input("priority")]
        public Input<string>? Priority { get; set; }

        [Input("reauthDefaultRule")]
        public Input<bool>? ReauthDefaultRule { get; set; }

        [Input("reauthIdleTimeout")]
        public Input<string>? ReauthIdleTimeout { get; set; }

        [Input("reauthTimeout")]
        public Input<string>? ReauthTimeout { get; set; }

        /// <summary>
        /// (Optional)
        /// </summary>
        [Input("ruleOrder")]
        public Input<string>? RuleOrder { get; set; }

        [Input("zpnInspectionProfileId")]
        public Input<string>? ZpnInspectionProfileId { get; set; }

        public ZPAPolicyAccessRuleState()
        {
        }
        public static new ZPAPolicyAccessRuleState Empty => new ZPAPolicyAccessRuleState();
    }
}
