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

    public sealed class ZPAApplicationSegmentInspectionCommonAppsDtoArgs : global::Pulumi.ResourceArgs
    {
        [Input("appsConfigs")]
        private InputList<Inputs.ZPAApplicationSegmentInspectionCommonAppsDtoAppsConfigArgs>? _appsConfigs;
        public InputList<Inputs.ZPAApplicationSegmentInspectionCommonAppsDtoAppsConfigArgs> AppsConfigs
        {
            get => _appsConfigs ?? (_appsConfigs = new InputList<Inputs.ZPAApplicationSegmentInspectionCommonAppsDtoAppsConfigArgs>());
            set => _appsConfigs = value;
        }

        public ZPAApplicationSegmentInspectionCommonAppsDtoArgs()
        {
        }
        public static new ZPAApplicationSegmentInspectionCommonAppsDtoArgs Empty => new ZPAApplicationSegmentInspectionCommonAppsDtoArgs();
    }
}
