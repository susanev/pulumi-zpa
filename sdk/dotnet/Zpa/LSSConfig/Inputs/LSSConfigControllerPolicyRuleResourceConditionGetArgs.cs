// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;
using Pulumi;

namespace zscaler.PulumiPackage.Zpa.LSSConfig.Inputs
{

    public sealed class LSSConfigControllerPolicyRuleResourceConditionGetArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// (Optional)
        /// </summary>
        [Input("negated")]
        public Input<bool>? Negated { get; set; }

        [Input("operands")]
        private InputList<Inputs.LSSConfigControllerPolicyRuleResourceConditionOperandGetArgs>? _operands;
        public InputList<Inputs.LSSConfigControllerPolicyRuleResourceConditionOperandGetArgs> Operands
        {
            get => _operands ?? (_operands = new InputList<Inputs.LSSConfigControllerPolicyRuleResourceConditionOperandGetArgs>());
            set => _operands = value;
        }

        /// <summary>
        /// (Optional) - Supported values are: `AND` or `OR`
        /// </summary>
        [Input("operator", required: true)]
        public Input<string> Operator { get; set; } = null!;

        public LSSConfigControllerPolicyRuleResourceConditionGetArgs()
        {
        }
        public static new LSSConfigControllerPolicyRuleResourceConditionGetArgs Empty => new LSSConfigControllerPolicyRuleResourceConditionGetArgs();
    }
}