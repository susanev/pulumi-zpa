// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;
using Pulumi;

namespace zscaler.PulumiPackage.Zpa.Outputs
{

    [OutputType]
    public sealed class GetZPALSSConfigControllerPolicyRuleResult
    {
        public readonly string Action;
        public readonly string ActionId;
        public readonly bool BypassDefaultRule;
        public readonly ImmutableArray<Outputs.GetZPALSSConfigControllerPolicyRuleConditionResult> Conditions;
        /// <summary>
        /// (string)
        /// </summary>
        public readonly string CreationTime;
        public readonly string CustomMsg;
        public readonly bool DefaultRule;
        /// <summary>
        /// (string)
        /// </summary>
        public readonly string Description;
        /// <summary>
        /// This field defines the name of the log streaming resource.
        /// </summary>
        public readonly string Id;
        public readonly bool IsolationDefaultRule;
        public readonly bool LssDefaultRule;
        /// <summary>
        /// (string)
        /// </summary>
        public readonly string ModifiedTime;
        public readonly string Modifiedby;
        /// <summary>
        /// This field defines the name of the log streaming resource.
        /// </summary>
        public readonly string Name;
        public readonly string Operator;
        public readonly string PolicySetId;
        public readonly string PolicyType;
        public readonly string Priority;
        public readonly bool ReauthDefaultRule;
        public readonly string ReauthIdleTimeout;
        public readonly string ReauthTimeout;
        public readonly string RuleOrder;
        public readonly string ZpnCbiProfileId;
        public readonly string ZpnInspectionProfileId;
        public readonly string ZpnInspectionProfileName;

        [OutputConstructor]
        private GetZPALSSConfigControllerPolicyRuleResult(
            string action,

            string actionId,

            bool bypassDefaultRule,

            ImmutableArray<Outputs.GetZPALSSConfigControllerPolicyRuleConditionResult> conditions,

            string creationTime,

            string customMsg,

            bool defaultRule,

            string description,

            string id,

            bool isolationDefaultRule,

            bool lssDefaultRule,

            string modifiedTime,

            string modifiedby,

            string name,

            string @operator,

            string policySetId,

            string policyType,

            string priority,

            bool reauthDefaultRule,

            string reauthIdleTimeout,

            string reauthTimeout,

            string ruleOrder,

            string zpnCbiProfileId,

            string zpnInspectionProfileId,

            string zpnInspectionProfileName)
        {
            Action = action;
            ActionId = actionId;
            BypassDefaultRule = bypassDefaultRule;
            Conditions = conditions;
            CreationTime = creationTime;
            CustomMsg = customMsg;
            DefaultRule = defaultRule;
            Description = description;
            Id = id;
            IsolationDefaultRule = isolationDefaultRule;
            LssDefaultRule = lssDefaultRule;
            ModifiedTime = modifiedTime;
            Modifiedby = modifiedby;
            Name = name;
            Operator = @operator;
            PolicySetId = policySetId;
            PolicyType = policyType;
            Priority = priority;
            ReauthDefaultRule = reauthDefaultRule;
            ReauthIdleTimeout = reauthIdleTimeout;
            ReauthTimeout = reauthTimeout;
            RuleOrder = ruleOrder;
            ZpnCbiProfileId = zpnCbiProfileId;
            ZpnInspectionProfileId = zpnInspectionProfileId;
            ZpnInspectionProfileName = zpnInspectionProfileName;
        }
    }
}
