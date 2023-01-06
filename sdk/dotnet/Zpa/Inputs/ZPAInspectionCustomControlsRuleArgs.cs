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

    public sealed class ZPAInspectionCustomControlsRuleArgs : global::Pulumi.ResourceArgs
    {
        [Input("conditions")]
        public Input<Inputs.ZPAInspectionCustomControlsRuleConditionsArgs>? Conditions { get; set; }

        [Input("names")]
        private InputList<string>? _names;
        public InputList<string> Names
        {
            get => _names ?? (_names = new InputList<string>());
            set => _names = value;
        }

        [Input("type")]
        public Input<string>? Type { get; set; }

        public ZPAInspectionCustomControlsRuleArgs()
        {
        }
        public static new ZPAInspectionCustomControlsRuleArgs Empty => new ZPAInspectionCustomControlsRuleArgs();
    }
}
