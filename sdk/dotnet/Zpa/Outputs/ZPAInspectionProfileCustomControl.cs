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
    public sealed class ZPAInspectionProfileCustomControl
    {
        /// <summary>
        /// The action of the predefined control. Supported values: `PASS`, `BLOCK` and `REDIRECT`
        /// </summary>
        public readonly string Action;
        /// <summary>
        /// Value for the predefined controls action. This field is only required if the action is set to REDIRECT. This field is only required if the action is set to `REDIRECT`.
        /// </summary>
        public readonly string? ActionValue;
        /// <summary>
        /// ID of the predefined control
        /// </summary>
        public readonly string Id;

        [OutputConstructor]
        private ZPAInspectionProfileCustomControl(
            string action,

            string? actionValue,

            string id)
        {
            Action = action;
            ActionValue = actionValue;
            Id = id;
        }
    }
}
