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
    public static class GetZPASegmentGroup
    {
        /// <summary>
        /// {{% examples %}}
        /// ## Example Usage
        /// {{% example %}}
        /// 
        /// ```csharp
        /// using System.Collections.Generic;
        /// using Pulumi;
        /// using Zpa = Pulumi.Zpa;
        /// 
        /// return await Deployment.RunAsync(() =&gt; 
        /// {
        ///     var example = Zpa.GetZPASegmentGroup.Invoke(new()
        ///     {
        ///         Name = "segment_group_name",
        ///     });
        /// 
        /// });
        /// ```
        /// {{% /example %}}
        /// {{% /examples %}}
        /// </summary>
        public static Task<GetZPASegmentGroupResult> InvokeAsync(GetZPASegmentGroupArgs? args = null, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.InvokeAsync<GetZPASegmentGroupResult>("zpa:index/getZPASegmentGroup:getZPASegmentGroup", args ?? new GetZPASegmentGroupArgs(), options.WithDefaults());

        /// <summary>
        /// {{% examples %}}
        /// ## Example Usage
        /// {{% example %}}
        /// 
        /// ```csharp
        /// using System.Collections.Generic;
        /// using Pulumi;
        /// using Zpa = Pulumi.Zpa;
        /// 
        /// return await Deployment.RunAsync(() =&gt; 
        /// {
        ///     var example = Zpa.GetZPASegmentGroup.Invoke(new()
        ///     {
        ///         Name = "segment_group_name",
        ///     });
        /// 
        /// });
        /// ```
        /// {{% /example %}}
        /// {{% /examples %}}
        /// </summary>
        public static Output<GetZPASegmentGroupResult> Invoke(GetZPASegmentGroupInvokeArgs? args = null, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.Invoke<GetZPASegmentGroupResult>("zpa:index/getZPASegmentGroup:getZPASegmentGroup", args ?? new GetZPASegmentGroupInvokeArgs(), options.WithDefaults());
    }


    public sealed class GetZPASegmentGroupArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// The ID of the segment group to be exported.
        /// </summary>
        [Input("id")]
        public string? Id { get; set; }

        /// <summary>
        /// The name of the segment group to be exported.
        /// </summary>
        [Input("name")]
        public string? Name { get; set; }

        /// <summary>
        /// (bool)
        /// </summary>
        [Input("policyMigrated")]
        public bool? PolicyMigrated { get; set; }

        public GetZPASegmentGroupArgs()
        {
        }
        public static new GetZPASegmentGroupArgs Empty => new GetZPASegmentGroupArgs();
    }

    public sealed class GetZPASegmentGroupInvokeArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// The ID of the segment group to be exported.
        /// </summary>
        [Input("id")]
        public Input<string>? Id { get; set; }

        /// <summary>
        /// The name of the segment group to be exported.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// (bool)
        /// </summary>
        [Input("policyMigrated")]
        public Input<bool>? PolicyMigrated { get; set; }

        public GetZPASegmentGroupInvokeArgs()
        {
        }
        public static new GetZPASegmentGroupInvokeArgs Empty => new GetZPASegmentGroupInvokeArgs();
    }


    [OutputType]
    public sealed class GetZPASegmentGroupResult
    {
        /// <summary>
        /// (Computed)
        /// </summary>
        public readonly ImmutableArray<Outputs.GetZPASegmentGroupApplicationResult> Applications;
        /// <summary>
        /// (string)
        /// </summary>
        public readonly string ConfigSpace;
        /// <summary>
        /// (string)
        /// </summary>
        public readonly string CreationTime;
        /// <summary>
        /// (string)
        /// </summary>
        public readonly string Description;
        /// <summary>
        /// (bool)
        /// </summary>
        public readonly bool Enabled;
        /// <summary>
        /// (string)
        /// </summary>
        public readonly string? Id;
        /// <summary>
        /// (string)
        /// </summary>
        public readonly string ModifiedBy;
        /// <summary>
        /// (string)
        /// </summary>
        public readonly string ModifiedTime;
        /// <summary>
        /// (string)
        /// </summary>
        public readonly string? Name;
        /// <summary>
        /// (bool)
        /// </summary>
        public readonly bool PolicyMigrated;
        /// <summary>
        /// (string)
        /// </summary>
        public readonly string TcpKeepAliveEnabled;

        [OutputConstructor]
        private GetZPASegmentGroupResult(
            ImmutableArray<Outputs.GetZPASegmentGroupApplicationResult> applications,

            string configSpace,

            string creationTime,

            string description,

            bool enabled,

            string? id,

            string modifiedBy,

            string modifiedTime,

            string? name,

            bool policyMigrated,

            string tcpKeepAliveEnabled)
        {
            Applications = applications;
            ConfigSpace = configSpace;
            CreationTime = creationTime;
            Description = description;
            Enabled = enabled;
            Id = id;
            ModifiedBy = modifiedBy;
            ModifiedTime = modifiedTime;
            Name = name;
            PolicyMigrated = policyMigrated;
            TcpKeepAliveEnabled = tcpKeepAliveEnabled;
        }
    }
}
