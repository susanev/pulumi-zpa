// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;
using Pulumi;

namespace zscaler.PulumiPackage.Zpa.Inputs
{

    public sealed class ZPAPolicyAccessInspectionRuleConditionGetArgs : global::Pulumi.ResourceArgs
    {
        [Input("id")]
        public Input<string>? Id { get; set; }

        [Input("negated")]
        public Input<bool>? Negated { get; set; }

        [Input("operands")]
        private InputList<Inputs.ZPAPolicyAccessInspectionRuleConditionOperandGetArgs>? _operands;
        public InputList<Inputs.ZPAPolicyAccessInspectionRuleConditionOperandGetArgs> Operands
        {
            get => _operands ?? (_operands = new InputList<Inputs.ZPAPolicyAccessInspectionRuleConditionOperandGetArgs>());
            set => _operands = value;
        }

        [Input("operator", required: true)]
        public Input<string> Operator { get; set; } = null!;

        public ZPAPolicyAccessInspectionRuleConditionGetArgs()
        {
        }
        public static new ZPAPolicyAccessInspectionRuleConditionGetArgs Empty => new ZPAPolicyAccessInspectionRuleConditionGetArgs();
    }
}
