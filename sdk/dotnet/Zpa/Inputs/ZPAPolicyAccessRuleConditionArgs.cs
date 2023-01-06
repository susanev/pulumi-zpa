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

    public sealed class ZPAPolicyAccessRuleConditionArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// (Optional) The ID of a server group resource
        /// </summary>
        [Input("id")]
        public Input<string>? Id { get; set; }

        /// <summary>
        /// (Optional) Supported values: ``true`` or ``false``
        /// * `operator` (Optional) Supported values: ``AND``, and ``OR``
        /// * `operands` (Optional) - Operands block must be repeated if multiple per `object_type` conditions are to be added to the rule.
        /// * `name` (Optional)
        /// * `lhs` (Optional) LHS must always carry the string value ``id`` or the attribute ID of the resource being associated with the rule.
        /// * `rhs` (Optional) RHS is either the ID attribute of a resource or fixed string value. Refer to the chart below for further details.
        /// * `idp_id` (Optional)
        /// * `object_type` (Optional) This is for specifying the policy critiera. Supported values: `APP`, `APP_GROUP`, `SAML`, `IDP`, `CLIENT_TYPE`, `TRUSTED_NETWORK`, `POSTURE`, `SCIM`, `SCIM_GROUP`, and `CLOUD_CONNECTOR_GROUP`. `TRUSTED_NETWORK`, and `CLIENT_TYPE`.
        /// * `CLIENT_TYPE` (Optional) - The below options are the only ones supported in an access policy rule.
        /// * `zpn_client_type_exporter`
        /// * `zpn_client_type_browser_isolation`
        /// * `zpn_client_type_machine_tunnel`
        /// * `zpn_client_type_ip_anchoring`
        /// * `zpn_client_type_edge_connector`
        /// * `zpn_client_type_zapp`
        /// </summary>
        [Input("negated")]
        public Input<bool>? Negated { get; set; }

        [Input("operands")]
        private InputList<Inputs.ZPAPolicyAccessRuleConditionOperandArgs>? _operands;
        public InputList<Inputs.ZPAPolicyAccessRuleConditionOperandArgs> Operands
        {
            get => _operands ?? (_operands = new InputList<Inputs.ZPAPolicyAccessRuleConditionOperandArgs>());
            set => _operands = value;
        }

        [Input("operator", required: true)]
        public Input<string> Operator { get; set; } = null!;

        public ZPAPolicyAccessRuleConditionArgs()
        {
        }
        public static new ZPAPolicyAccessRuleConditionArgs Empty => new ZPAPolicyAccessRuleConditionArgs();
    }
}
