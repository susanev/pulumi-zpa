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

    public sealed class ZPALSSConfigControllerPolicyRuleResourceConditionOperandArgs : global::Pulumi.ResourceArgs
    {
        [Input("objectType", required: true)]
        public Input<string> ObjectType { get; set; } = null!;

        [Input("values")]
        private InputList<string>? _values;
        public InputList<string> Values
        {
            get => _values ?? (_values = new InputList<string>());
            set => _values = value;
        }

        public ZPALSSConfigControllerPolicyRuleResourceConditionOperandArgs()
        {
        }
        public static new ZPALSSConfigControllerPolicyRuleResourceConditionOperandArgs Empty => new ZPALSSConfigControllerPolicyRuleResourceConditionOperandArgs();
    }
}
