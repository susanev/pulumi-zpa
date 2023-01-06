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
    /// The  **zpa_inspection_profile** resource creates an inspection profile in the Zscaler Private Access cloud. This resource can then be referenced in an inspection custom control resource.
    /// 
    /// ## Import
    /// 
    /// Zscaler offers a dedicated tool called Zscaler-Terraformer to allow the automated import of ZPA configurations into Terraform-compliant HashiCorp Configuration Language. [Visit](https://github.com/zscaler/zscaler-terraformer)
    /// </summary>
    [ZpaResourceType("zpa:index/zPAInspectionProfile:ZPAInspectionProfile")]
    public partial class ZPAInspectionProfile : global::Pulumi.CustomResource
    {
        [Output("associateAllControls")]
        public Output<bool?> AssociateAllControls { get; private set; } = null!;

        /// <summary>
        /// (Optional)
        /// </summary>
        [Output("commonGlobalOverrideActionsConfig")]
        public Output<ImmutableDictionary<string, string>> CommonGlobalOverrideActionsConfig { get; private set; } = null!;

        /// <summary>
        /// (Optional) Types for custom controls
        /// </summary>
        [Output("controlsInfos")]
        public Output<ImmutableArray<Outputs.ZPAInspectionProfileControlsInfo>> ControlsInfos { get; private set; } = null!;

        /// <summary>
        /// (Optional) Types for custom controls
        /// * `type` (Optional) Types for custom controls
        /// * `control_rule_json` (Optional) Custom controls string in JSON format
        /// </summary>
        [Output("customControls")]
        public Output<ImmutableArray<Outputs.ZPAInspectionProfileCustomControl>> CustomControls { get; private set; } = null!;

        /// <summary>
        /// Description of the inspection profile.
        /// </summary>
        [Output("description")]
        public Output<string> Description { get; private set; } = null!;

        [Output("globalControlActions")]
        public Output<ImmutableArray<string>> GlobalControlActions { get; private set; } = null!;

        [Output("incarnationNumber")]
        public Output<string> IncarnationNumber { get; private set; } = null!;

        /// <summary>
        /// The name of the inspection profile.
        /// </summary>
        [Output("name")]
        public Output<string> Name { get; private set; } = null!;

        /// <summary>
        /// OWASP Predefined Paranoia Level. Range: [1-4], inclusive
        /// </summary>
        [Output("paranoiaLevel")]
        public Output<string> ParanoiaLevel { get; private set; } = null!;

        /// <summary>
        /// The predefined controls. The default predefined control `Preprocessors` are mandatory and injected in the request by default. Individual `predefined_controls` can be set by using the data source `data_source_zpa_predefined_controls` or by group using the data source `zpa.getZPAInspectionAllPredefinedControls`.
        /// </summary>
        [Output("predefinedControls")]
        public Output<ImmutableArray<Outputs.ZPAInspectionProfilePredefinedControl>> PredefinedControls { get; private set; } = null!;

        [Output("predefinedControlsVersion")]
        public Output<string> PredefinedControlsVersion { get; private set; } = null!;


        /// <summary>
        /// Create a ZPAInspectionProfile resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public ZPAInspectionProfile(string name, ZPAInspectionProfileArgs? args = null, CustomResourceOptions? options = null)
            : base("zpa:index/zPAInspectionProfile:ZPAInspectionProfile", name, args ?? new ZPAInspectionProfileArgs(), MakeResourceOptions(options, ""))
        {
        }

        private ZPAInspectionProfile(string name, Input<string> id, ZPAInspectionProfileState? state = null, CustomResourceOptions? options = null)
            : base("zpa:index/zPAInspectionProfile:ZPAInspectionProfile", name, state, MakeResourceOptions(options, id))
        {
        }

        private static CustomResourceOptions MakeResourceOptions(CustomResourceOptions? options, Input<string>? id)
        {
            var defaultOptions = new CustomResourceOptions
            {
                Version = Utilities.Version,
                PluginDownloadURL = "github://api.github.com/zscaler",
            };
            var merged = CustomResourceOptions.Merge(defaultOptions, options);
            // Override the ID if one was specified for consistency with other language SDKs.
            merged.Id = id ?? merged.Id;
            return merged;
        }
        /// <summary>
        /// Get an existing ZPAInspectionProfile resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static ZPAInspectionProfile Get(string name, Input<string> id, ZPAInspectionProfileState? state = null, CustomResourceOptions? options = null)
        {
            return new ZPAInspectionProfile(name, id, state, options);
        }
    }

    public sealed class ZPAInspectionProfileArgs : global::Pulumi.ResourceArgs
    {
        [Input("associateAllControls")]
        public Input<bool>? AssociateAllControls { get; set; }

        [Input("commonGlobalOverrideActionsConfig")]
        private InputMap<string>? _commonGlobalOverrideActionsConfig;

        /// <summary>
        /// (Optional)
        /// </summary>
        public InputMap<string> CommonGlobalOverrideActionsConfig
        {
            get => _commonGlobalOverrideActionsConfig ?? (_commonGlobalOverrideActionsConfig = new InputMap<string>());
            set => _commonGlobalOverrideActionsConfig = value;
        }

        [Input("controlsInfos")]
        private InputList<Inputs.ZPAInspectionProfileControlsInfoArgs>? _controlsInfos;

        /// <summary>
        /// (Optional) Types for custom controls
        /// </summary>
        public InputList<Inputs.ZPAInspectionProfileControlsInfoArgs> ControlsInfos
        {
            get => _controlsInfos ?? (_controlsInfos = new InputList<Inputs.ZPAInspectionProfileControlsInfoArgs>());
            set => _controlsInfos = value;
        }

        [Input("customControls")]
        private InputList<Inputs.ZPAInspectionProfileCustomControlArgs>? _customControls;

        /// <summary>
        /// (Optional) Types for custom controls
        /// * `type` (Optional) Types for custom controls
        /// * `control_rule_json` (Optional) Custom controls string in JSON format
        /// </summary>
        public InputList<Inputs.ZPAInspectionProfileCustomControlArgs> CustomControls
        {
            get => _customControls ?? (_customControls = new InputList<Inputs.ZPAInspectionProfileCustomControlArgs>());
            set => _customControls = value;
        }

        /// <summary>
        /// Description of the inspection profile.
        /// </summary>
        [Input("description")]
        public Input<string>? Description { get; set; }

        [Input("globalControlActions")]
        private InputList<string>? _globalControlActions;
        public InputList<string> GlobalControlActions
        {
            get => _globalControlActions ?? (_globalControlActions = new InputList<string>());
            set => _globalControlActions = value;
        }

        [Input("incarnationNumber")]
        public Input<string>? IncarnationNumber { get; set; }

        /// <summary>
        /// The name of the inspection profile.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// OWASP Predefined Paranoia Level. Range: [1-4], inclusive
        /// </summary>
        [Input("paranoiaLevel")]
        public Input<string>? ParanoiaLevel { get; set; }

        [Input("predefinedControls")]
        private InputList<Inputs.ZPAInspectionProfilePredefinedControlArgs>? _predefinedControls;

        /// <summary>
        /// The predefined controls. The default predefined control `Preprocessors` are mandatory and injected in the request by default. Individual `predefined_controls` can be set by using the data source `data_source_zpa_predefined_controls` or by group using the data source `zpa.getZPAInspectionAllPredefinedControls`.
        /// </summary>
        public InputList<Inputs.ZPAInspectionProfilePredefinedControlArgs> PredefinedControls
        {
            get => _predefinedControls ?? (_predefinedControls = new InputList<Inputs.ZPAInspectionProfilePredefinedControlArgs>());
            set => _predefinedControls = value;
        }

        [Input("predefinedControlsVersion")]
        public Input<string>? PredefinedControlsVersion { get; set; }

        public ZPAInspectionProfileArgs()
        {
        }
        public static new ZPAInspectionProfileArgs Empty => new ZPAInspectionProfileArgs();
    }

    public sealed class ZPAInspectionProfileState : global::Pulumi.ResourceArgs
    {
        [Input("associateAllControls")]
        public Input<bool>? AssociateAllControls { get; set; }

        [Input("commonGlobalOverrideActionsConfig")]
        private InputMap<string>? _commonGlobalOverrideActionsConfig;

        /// <summary>
        /// (Optional)
        /// </summary>
        public InputMap<string> CommonGlobalOverrideActionsConfig
        {
            get => _commonGlobalOverrideActionsConfig ?? (_commonGlobalOverrideActionsConfig = new InputMap<string>());
            set => _commonGlobalOverrideActionsConfig = value;
        }

        [Input("controlsInfos")]
        private InputList<Inputs.ZPAInspectionProfileControlsInfoGetArgs>? _controlsInfos;

        /// <summary>
        /// (Optional) Types for custom controls
        /// </summary>
        public InputList<Inputs.ZPAInspectionProfileControlsInfoGetArgs> ControlsInfos
        {
            get => _controlsInfos ?? (_controlsInfos = new InputList<Inputs.ZPAInspectionProfileControlsInfoGetArgs>());
            set => _controlsInfos = value;
        }

        [Input("customControls")]
        private InputList<Inputs.ZPAInspectionProfileCustomControlGetArgs>? _customControls;

        /// <summary>
        /// (Optional) Types for custom controls
        /// * `type` (Optional) Types for custom controls
        /// * `control_rule_json` (Optional) Custom controls string in JSON format
        /// </summary>
        public InputList<Inputs.ZPAInspectionProfileCustomControlGetArgs> CustomControls
        {
            get => _customControls ?? (_customControls = new InputList<Inputs.ZPAInspectionProfileCustomControlGetArgs>());
            set => _customControls = value;
        }

        /// <summary>
        /// Description of the inspection profile.
        /// </summary>
        [Input("description")]
        public Input<string>? Description { get; set; }

        [Input("globalControlActions")]
        private InputList<string>? _globalControlActions;
        public InputList<string> GlobalControlActions
        {
            get => _globalControlActions ?? (_globalControlActions = new InputList<string>());
            set => _globalControlActions = value;
        }

        [Input("incarnationNumber")]
        public Input<string>? IncarnationNumber { get; set; }

        /// <summary>
        /// The name of the inspection profile.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// OWASP Predefined Paranoia Level. Range: [1-4], inclusive
        /// </summary>
        [Input("paranoiaLevel")]
        public Input<string>? ParanoiaLevel { get; set; }

        [Input("predefinedControls")]
        private InputList<Inputs.ZPAInspectionProfilePredefinedControlGetArgs>? _predefinedControls;

        /// <summary>
        /// The predefined controls. The default predefined control `Preprocessors` are mandatory and injected in the request by default. Individual `predefined_controls` can be set by using the data source `data_source_zpa_predefined_controls` or by group using the data source `zpa.getZPAInspectionAllPredefinedControls`.
        /// </summary>
        public InputList<Inputs.ZPAInspectionProfilePredefinedControlGetArgs> PredefinedControls
        {
            get => _predefinedControls ?? (_predefinedControls = new InputList<Inputs.ZPAInspectionProfilePredefinedControlGetArgs>());
            set => _predefinedControls = value;
        }

        [Input("predefinedControlsVersion")]
        public Input<string>? PredefinedControlsVersion { get; set; }

        public ZPAInspectionProfileState()
        {
        }
        public static new ZPAInspectionProfileState Empty => new ZPAInspectionProfileState();
    }
}
