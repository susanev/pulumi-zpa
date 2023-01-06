// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;
using Pulumi;

namespace zscaler.PulumiPackage.Zpa
{
    /// <summary>
    /// The **zpa_lss_config_controller** resource creates and manages Log Streaming Service (LSS) in the Zscaler Private Access cloud.
    /// 
    /// ## Example 1 Usage
    /// 
    /// ```csharp
    /// using System.Collections.Generic;
    /// using Pulumi;
    /// using Zpa = Pulumi.Zpa;
    /// using Zpa = zscaler.PulumiPackage.Zpa;
    /// 
    /// return await Deployment.RunAsync(() =&gt; 
    /// {
    ///     var zpnAstAuthLog = Zpa.GetZPALSSLogTypeFormats.Invoke(new()
    ///     {
    ///         LogType = "zpn_ast_auth_log",
    ///     });
    /// 
    ///     // Create Log Receiver Configuration
    ///     var example = new Zpa.ZPALSSConfigController("example", new()
    ///     {
    ///         Config = new Zpa.Inputs.ZPALSSConfigControllerConfigArgs
    ///         {
    ///             Name = "Example",
    ///             Description = "Example",
    ///             Enabled = true,
    ///             Format = zpnAstAuthLog.Apply(getZPALSSLogTypeFormatsResult =&gt; getZPALSSLogTypeFormatsResult.Json),
    ///             LssHost = "splunk.acme.com",
    ///             LssPort = "11000",
    ///             SourceLogType = "zpn_ast_auth_log",
    ///             UseTls = true,
    ///             Filters = new[]
    ///             {
    ///                 "ZPN_STATUS_AUTH_FAILED",
    ///                 "ZPN_STATUS_DISCONNECTED",
    ///                 "ZPN_STATUS_AUTHENTICATED",
    ///             },
    ///         },
    ///         ConnectorGroups = new[]
    ///         {
    ///             new Zpa.Inputs.ZPALSSConfigControllerConnectorGroupArgs
    ///             {
    ///                 Ids = new[]
    ///                 {
    ///                     zpa_app_connector_group.Example.Id,
    ///                 },
    ///             },
    ///         },
    ///     });
    /// 
    /// });
    /// ```
    /// 
    /// ## Example 2 Usage
    /// 
    /// ```csharp
    /// using System.Collections.Generic;
    /// using Pulumi;
    /// using Zpa = Pulumi.Zpa;
    /// using Zpa = zscaler.PulumiPackage.Zpa;
    /// 
    /// return await Deployment.RunAsync(() =&gt; 
    /// {
    ///     var zpnTransLog = Zpa.GetZPALSSLogTypeFormats.Invoke(new()
    ///     {
    ///         LogType = "zpn_trans_log",
    ///     });
    /// 
    ///     var lssSiemPolicy = Zpa.GetZPAPolicyType.Invoke(new()
    ///     {
    ///         PolicyType = "SIEM_POLICY",
    ///     });
    /// 
    ///     var lssUserActivity = new Zpa.ZPALSSConfigController("lssUserActivity", new()
    ///     {
    ///         Config = new Zpa.Inputs.ZPALSSConfigControllerConfigArgs
    ///         {
    ///             Name = "LSS User Activity",
    ///             Description = "LSS User Activity",
    ///             Enabled = true,
    ///             Format = zpnTransLog.Apply(getZPALSSLogTypeFormatsResult =&gt; getZPALSSLogTypeFormatsResult.Json),
    ///             LssHost = "splunk.acme.com",
    ///             LssPort = "11001",
    ///             SourceLogType = "zpn_trans_log",
    ///             UseTls = true,
    ///         },
    ///         PolicyRuleResource = new Zpa.Inputs.ZPALSSConfigControllerPolicyRuleResourceArgs
    ///         {
    ///             Name = "policy_rule_resource-lss_user_activity",
    ///             Action = "ALLOW",
    ///             PolicySetId = lssSiemPolicy.Apply(getZPAPolicyTypeResult =&gt; getZPAPolicyTypeResult.Id),
    ///             Conditions = new[]
    ///             {
    ///                 new Zpa.Inputs.ZPALSSConfigControllerPolicyRuleResourceConditionArgs
    ///                 {
    ///                     Negated = false,
    ///                     Operator = "OR",
    ///                     Operands = new[]
    ///                     {
    ///                         new Zpa.Inputs.ZPALSSConfigControllerPolicyRuleResourceConditionOperandArgs
    ///                         {
    ///                             ObjectType = "CLIENT_TYPE",
    ///                             Values = new[]
    ///                             {
    ///                                 "zpn_client_type_exporter",
    ///                             },
    ///                         },
    ///                         new Zpa.Inputs.ZPALSSConfigControllerPolicyRuleResourceConditionOperandArgs
    ///                         {
    ///                             ObjectType = "CLIENT_TYPE",
    ///                             Values = new[]
    ///                             {
    ///                                 "zpn_client_type_ip_anchoring",
    ///                             },
    ///                         },
    ///                         new Zpa.Inputs.ZPALSSConfigControllerPolicyRuleResourceConditionOperandArgs
    ///                         {
    ///                             ObjectType = "CLIENT_TYPE",
    ///                             Values = new[]
    ///                             {
    ///                                 "zpn_client_type_zapp",
    ///                             },
    ///                         },
    ///                         new Zpa.Inputs.ZPALSSConfigControllerPolicyRuleResourceConditionOperandArgs
    ///                         {
    ///                             ObjectType = "CLIENT_TYPE",
    ///                             Values = new[]
    ///                             {
    ///                                 "zpn_client_type_edge_connector",
    ///                             },
    ///                         },
    ///                         new Zpa.Inputs.ZPALSSConfigControllerPolicyRuleResourceConditionOperandArgs
    ///                         {
    ///                             ObjectType = "CLIENT_TYPE",
    ///                             Values = new[]
    ///                             {
    ///                                 "zpn_client_type_machine_tunnel",
    ///                             },
    ///                         },
    ///                         new Zpa.Inputs.ZPALSSConfigControllerPolicyRuleResourceConditionOperandArgs
    ///                         {
    ///                             ObjectType = "CLIENT_TYPE",
    ///                             Values = new[]
    ///                             {
    ///                                 "zpn_client_type_browser_isolation",
    ///                             },
    ///                         },
    ///                         new Zpa.Inputs.ZPALSSConfigControllerPolicyRuleResourceConditionOperandArgs
    ///                         {
    ///                             ObjectType = "CLIENT_TYPE",
    ///                             Values = new[]
    ///                             {
    ///                                 "zpn_client_type_slogger",
    ///                             },
    ///                         },
    ///                     },
    ///                 },
    ///             },
    ///         },
    ///         ConnectorGroups = new[]
    ///         {
    ///             new Zpa.Inputs.ZPALSSConfigControllerConnectorGroupArgs
    ///             {
    ///                 Ids = new[]
    ///                 {
    ///                     zpa_app_connector_group.Example.Id,
    ///                 },
    ///             },
    ///         },
    ///     });
    /// 
    /// });
    /// ```
    /// 
    /// ## Import
    /// 
    /// Zscaler offers a dedicated tool called Zscaler-Terraformer to allow the automated import of ZPA configurations into Terraform-compliant HashiCorp Configuration Language. [Visit](https://github.com/zscaler/zscaler-terraformer)
    /// </summary>
    [ZpaResourceType("zpa:index/zPALSSConfigController:ZPALSSConfigController")]
    public partial class ZPALSSConfigController : global::Pulumi.CustomResource
    {
        /// <summary>
        /// (Required)
        /// </summary>
        [Output("config")]
        public Output<Outputs.ZPALSSConfigControllerConfig?> Config { get; private set; } = null!;

        /// <summary>
        /// (Optional)
        /// </summary>
        [Output("connectorGroups")]
        public Output<ImmutableArray<Outputs.ZPALSSConfigControllerConnectorGroup>> ConnectorGroups { get; private set; } = null!;

        [Output("policyRuleId")]
        public Output<string> PolicyRuleId { get; private set; } = null!;

        /// <summary>
        /// (Optional)
        /// </summary>
        [Output("policyRuleResource")]
        public Output<Outputs.ZPALSSConfigControllerPolicyRuleResource?> PolicyRuleResource { get; private set; } = null!;


        /// <summary>
        /// Create a ZPALSSConfigController resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public ZPALSSConfigController(string name, ZPALSSConfigControllerArgs? args = null, CustomResourceOptions? options = null)
            : base("zpa:index/zPALSSConfigController:ZPALSSConfigController", name, args ?? new ZPALSSConfigControllerArgs(), MakeResourceOptions(options, ""))
        {
        }

        private ZPALSSConfigController(string name, Input<string> id, ZPALSSConfigControllerState? state = null, CustomResourceOptions? options = null)
            : base("zpa:index/zPALSSConfigController:ZPALSSConfigController", name, state, MakeResourceOptions(options, id))
        {
        }

        private static CustomResourceOptions MakeResourceOptions(CustomResourceOptions? options, Input<string>? id)
        {
            var defaultOptions = new CustomResourceOptions
            {
                Version = Utilities.Version,
            };
            var merged = CustomResourceOptions.Merge(defaultOptions, options);
            // Override the ID if one was specified for consistency with other language SDKs.
            merged.Id = id ?? merged.Id;
            return merged;
        }
        /// <summary>
        /// Get an existing ZPALSSConfigController resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static ZPALSSConfigController Get(string name, Input<string> id, ZPALSSConfigControllerState? state = null, CustomResourceOptions? options = null)
        {
            return new ZPALSSConfigController(name, id, state, options);
        }
    }

    public sealed class ZPALSSConfigControllerArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// (Required)
        /// </summary>
        [Input("config")]
        public Input<Inputs.ZPALSSConfigControllerConfigArgs>? Config { get; set; }

        [Input("connectorGroups")]
        private InputList<Inputs.ZPALSSConfigControllerConnectorGroupArgs>? _connectorGroups;

        /// <summary>
        /// (Optional)
        /// </summary>
        public InputList<Inputs.ZPALSSConfigControllerConnectorGroupArgs> ConnectorGroups
        {
            get => _connectorGroups ?? (_connectorGroups = new InputList<Inputs.ZPALSSConfigControllerConnectorGroupArgs>());
            set => _connectorGroups = value;
        }

        /// <summary>
        /// (Optional)
        /// </summary>
        [Input("policyRuleResource")]
        public Input<Inputs.ZPALSSConfigControllerPolicyRuleResourceArgs>? PolicyRuleResource { get; set; }

        public ZPALSSConfigControllerArgs()
        {
        }
        public static new ZPALSSConfigControllerArgs Empty => new ZPALSSConfigControllerArgs();
    }

    public sealed class ZPALSSConfigControllerState : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// (Required)
        /// </summary>
        [Input("config")]
        public Input<Inputs.ZPALSSConfigControllerConfigGetArgs>? Config { get; set; }

        [Input("connectorGroups")]
        private InputList<Inputs.ZPALSSConfigControllerConnectorGroupGetArgs>? _connectorGroups;

        /// <summary>
        /// (Optional)
        /// </summary>
        public InputList<Inputs.ZPALSSConfigControllerConnectorGroupGetArgs> ConnectorGroups
        {
            get => _connectorGroups ?? (_connectorGroups = new InputList<Inputs.ZPALSSConfigControllerConnectorGroupGetArgs>());
            set => _connectorGroups = value;
        }

        [Input("policyRuleId")]
        public Input<string>? PolicyRuleId { get; set; }

        /// <summary>
        /// (Optional)
        /// </summary>
        [Input("policyRuleResource")]
        public Input<Inputs.ZPALSSConfigControllerPolicyRuleResourceGetArgs>? PolicyRuleResource { get; set; }

        public ZPALSSConfigControllerState()
        {
        }
        public static new ZPALSSConfigControllerState Empty => new ZPALSSConfigControllerState();
    }
}
