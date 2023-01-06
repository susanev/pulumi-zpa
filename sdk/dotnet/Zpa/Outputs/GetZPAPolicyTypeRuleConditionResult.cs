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
    public sealed class GetZPAPolicyTypeRuleConditionResult
    {
        public readonly string CreationTime;
        public readonly string Id;
        public readonly string ModifiedBy;
        public readonly string ModifiedTime;
        public readonly bool Negated;
        public readonly ImmutableArray<Outputs.GetZPAPolicyTypeRuleConditionOperandResult> Operands;
        public readonly string Operator;

        [OutputConstructor]
        private GetZPAPolicyTypeRuleConditionResult(
            string creationTime,

            string id,

            string modifiedBy,

            string modifiedTime,

            bool negated,

            ImmutableArray<Outputs.GetZPAPolicyTypeRuleConditionOperandResult> operands,

            string @operator)
        {
            CreationTime = creationTime;
            Id = id;
            ModifiedBy = modifiedBy;
            ModifiedTime = modifiedTime;
            Negated = negated;
            Operands = operands;
            Operator = @operator;
        }
    }
}
