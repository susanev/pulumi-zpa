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
    public sealed class GetZPAInspectionCustomControlsRuleResult
    {
        public readonly ImmutableArray<Outputs.GetZPAInspectionCustomControlsRuleConditionResult> Conditions;
        public readonly ImmutableArray<string> Names;
        public readonly string Type;

        [OutputConstructor]
        private GetZPAInspectionCustomControlsRuleResult(
            ImmutableArray<Outputs.GetZPAInspectionCustomControlsRuleConditionResult> conditions,

            ImmutableArray<string> names,

            string type)
        {
            Conditions = conditions;
            Names = names;
            Type = type;
        }
    }
}
