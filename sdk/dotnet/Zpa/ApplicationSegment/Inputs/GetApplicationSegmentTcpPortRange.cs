// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;
using Pulumi;

namespace zscaler.PulumiPackage.Zpa.ApplicationSegment.Inputs
{

    public sealed class GetApplicationSegmentTcpPortRangeArgs : global::Pulumi.InvokeArgs
    {
        [Input("from")]
        public string? From { get; set; }

        [Input("to")]
        public string? To { get; set; }

        public GetApplicationSegmentTcpPortRangeArgs()
        {
        }
        public static new GetApplicationSegmentTcpPortRangeArgs Empty => new GetApplicationSegmentTcpPortRangeArgs();
    }
}