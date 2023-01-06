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

    public sealed class ZPAInspectionProfileCustomControlArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// The action of the predefined control. Supported values: `PASS`, `BLOCK` and `REDIRECT`
        /// </summary>
        [Input("action", required: true)]
        public Input<string> Action { get; set; } = null!;

        /// <summary>
        /// Value for the predefined controls action. This field is only required if the action is set to REDIRECT. This field is only required if the action is set to `REDIRECT`.
        /// </summary>
        [Input("actionValue")]
        public Input<string>? ActionValue { get; set; }

        /// <summary>
        /// ID of the predefined control
        /// </summary>
        [Input("id", required: true)]
        public Input<string> Id { get; set; } = null!;

        public ZPAInspectionProfileCustomControlArgs()
        {
        }
        public static new ZPAInspectionProfileCustomControlArgs Empty => new ZPAInspectionProfileCustomControlArgs();
    }
}
