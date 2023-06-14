// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;
using Pulumi;

namespace zscaler.PulumiPackage.Zpa.PostureProfile
{
    public static class GetPostureProfile
    {
        /// <summary>
        /// Use the **zpa_posture_profile** data source to get information about a posture profile created in the Zscaler Private Access Mobile Portal. This data source can then be referenced in an Access Policy, Timeout policy, Forwarding Policy, Inspection Policy or Isolation Policy.
        /// 
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
        ///     var example1 = Zpa.PostureProfile.GetPostureProfile.Invoke(new()
        ///     {
        ///         Name = "CrowdStrike_ZPA_ZTA_40",
        ///     });
        /// 
        /// });
        /// ```
        /// 
        /// ```csharp
        /// using System.Collections.Generic;
        /// using Pulumi;
        /// using Zpa = Pulumi.Zpa;
        /// 
        /// return await Deployment.RunAsync(() =&gt; 
        /// {
        ///     var example2 = Zpa.PostureProfile.GetPostureProfile.Invoke(new()
        ///     {
        ///         Name = "Detect SentinelOne",
        ///     });
        /// 
        /// });
        /// ```
        /// 
        /// ```csharp
        /// using System.Collections.Generic;
        /// using Pulumi;
        /// using Zpa = Pulumi.Zpa;
        /// 
        /// return await Deployment.RunAsync(() =&gt; 
        /// {
        ///     var example3 = Zpa.PostureProfile.GetPostureProfile.Invoke(new()
        ///     {
        ///         Name = "domain_joined",
        ///     });
        /// 
        /// });
        /// ```
        /// 
        /// &gt; **NOTE** To query posture profiles that are associated with a specific Zscaler cloud, it is required to append the cloud name to the name of the posture profile as the below example:
        /// 
        /// ```csharp
        /// using System.Collections.Generic;
        /// using Pulumi;
        /// using Zpa = Pulumi.Zpa;
        /// 
        /// return await Deployment.RunAsync(() =&gt; 
        /// {
        ///     var example1 = Zpa.PostureProfile.GetPostureProfile.Invoke(new()
        ///     {
        ///         Name = "CrowdStrike_ZPA_ZTA_40 (zscalertwo.net)",
        ///     });
        /// 
        /// });
        /// ```
        /// 
        /// &gt; **NOTE** When associating a posture profile with one of supported resources, the following parameter must be exported: ``posture_udid`` instead of the ``id`` of the resource.
        /// 
        /// ```csharp
        /// using System.Collections.Generic;
        /// using Pulumi;
        /// using Zpa = Pulumi.Zpa;
        /// 
        /// return await Deployment.RunAsync(() =&gt; 
        /// {
        ///     var example1 = Zpa.PostureProfile.GetPostureProfile.Invoke(new()
        ///     {
        ///         Name = "CrowdStrike_ZPA_ZTA_40 (zscalertwo.net)",
        ///     });
        /// 
        ///     return new Dictionary&lt;string, object?&gt;
        ///     {
        ///         ["zpaPostureProfile"] = example1.Apply(getPostureProfileResult =&gt; getPostureProfileResult.PostureUdid),
        ///     };
        /// });
        /// ```
        /// {{% /example %}}
        /// {{% /examples %}}
        /// </summary>
        public static Task<GetPostureProfileResult> InvokeAsync(GetPostureProfileArgs? args = null, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.InvokeAsync<GetPostureProfileResult>("zpa:postureProfile/getPostureProfile:getPostureProfile", args ?? new GetPostureProfileArgs(), options.WithDefaults());

        /// <summary>
        /// Use the **zpa_posture_profile** data source to get information about a posture profile created in the Zscaler Private Access Mobile Portal. This data source can then be referenced in an Access Policy, Timeout policy, Forwarding Policy, Inspection Policy or Isolation Policy.
        /// 
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
        ///     var example1 = Zpa.PostureProfile.GetPostureProfile.Invoke(new()
        ///     {
        ///         Name = "CrowdStrike_ZPA_ZTA_40",
        ///     });
        /// 
        /// });
        /// ```
        /// 
        /// ```csharp
        /// using System.Collections.Generic;
        /// using Pulumi;
        /// using Zpa = Pulumi.Zpa;
        /// 
        /// return await Deployment.RunAsync(() =&gt; 
        /// {
        ///     var example2 = Zpa.PostureProfile.GetPostureProfile.Invoke(new()
        ///     {
        ///         Name = "Detect SentinelOne",
        ///     });
        /// 
        /// });
        /// ```
        /// 
        /// ```csharp
        /// using System.Collections.Generic;
        /// using Pulumi;
        /// using Zpa = Pulumi.Zpa;
        /// 
        /// return await Deployment.RunAsync(() =&gt; 
        /// {
        ///     var example3 = Zpa.PostureProfile.GetPostureProfile.Invoke(new()
        ///     {
        ///         Name = "domain_joined",
        ///     });
        /// 
        /// });
        /// ```
        /// 
        /// &gt; **NOTE** To query posture profiles that are associated with a specific Zscaler cloud, it is required to append the cloud name to the name of the posture profile as the below example:
        /// 
        /// ```csharp
        /// using System.Collections.Generic;
        /// using Pulumi;
        /// using Zpa = Pulumi.Zpa;
        /// 
        /// return await Deployment.RunAsync(() =&gt; 
        /// {
        ///     var example1 = Zpa.PostureProfile.GetPostureProfile.Invoke(new()
        ///     {
        ///         Name = "CrowdStrike_ZPA_ZTA_40 (zscalertwo.net)",
        ///     });
        /// 
        /// });
        /// ```
        /// 
        /// &gt; **NOTE** When associating a posture profile with one of supported resources, the following parameter must be exported: ``posture_udid`` instead of the ``id`` of the resource.
        /// 
        /// ```csharp
        /// using System.Collections.Generic;
        /// using Pulumi;
        /// using Zpa = Pulumi.Zpa;
        /// 
        /// return await Deployment.RunAsync(() =&gt; 
        /// {
        ///     var example1 = Zpa.PostureProfile.GetPostureProfile.Invoke(new()
        ///     {
        ///         Name = "CrowdStrike_ZPA_ZTA_40 (zscalertwo.net)",
        ///     });
        /// 
        ///     return new Dictionary&lt;string, object?&gt;
        ///     {
        ///         ["zpaPostureProfile"] = example1.Apply(getPostureProfileResult =&gt; getPostureProfileResult.PostureUdid),
        ///     };
        /// });
        /// ```
        /// {{% /example %}}
        /// {{% /examples %}}
        /// </summary>
        public static Output<GetPostureProfileResult> Invoke(GetPostureProfileInvokeArgs? args = null, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.Invoke<GetPostureProfileResult>("zpa:postureProfile/getPostureProfile:getPostureProfile", args ?? new GetPostureProfileInvokeArgs(), options.WithDefaults());
    }


    public sealed class GetPostureProfileArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// The name of the posture profile to be exported.
        /// </summary>
        [Input("name")]
        public string? Name { get; set; }

        public GetPostureProfileArgs()
        {
        }
        public static new GetPostureProfileArgs Empty => new GetPostureProfileArgs();
    }

    public sealed class GetPostureProfileInvokeArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// The name of the posture profile to be exported.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        public GetPostureProfileInvokeArgs()
        {
        }
        public static new GetPostureProfileInvokeArgs Empty => new GetPostureProfileInvokeArgs();
    }


    [OutputType]
    public sealed class GetPostureProfileResult
    {
        /// <summary>
        /// (string)
        /// </summary>
        public readonly string CreationTime;
        /// <summary>
        /// (string)
        /// </summary>
        public readonly string Domain;
        public readonly string Id;
        /// <summary>
        /// (string)
        /// </summary>
        public readonly string MasterCustomerId;
        /// <summary>
        /// (string)
        /// </summary>
        public readonly string ModifiedTime;
        public readonly string Modifiedby;
        public readonly string? Name;
        /// <summary>
        /// (string)
        /// </summary>
        public readonly string PostureUdid;
        /// <summary>
        /// (string)
        /// </summary>
        public readonly string ZscalerCloud;
        /// <summary>
        /// (string)
        /// </summary>
        public readonly string ZscalerCustomerId;

        [OutputConstructor]
        private GetPostureProfileResult(
            string creationTime,

            string domain,

            string id,

            string masterCustomerId,

            string modifiedTime,

            string modifiedby,

            string? name,

            string postureUdid,

            string zscalerCloud,

            string zscalerCustomerId)
        {
            CreationTime = creationTime;
            Domain = domain;
            Id = id;
            MasterCustomerId = masterCustomerId;
            ModifiedTime = modifiedTime;
            Modifiedby = modifiedby;
            Name = name;
            PostureUdid = postureUdid;
            ZscalerCloud = zscalerCloud;
            ZscalerCustomerId = zscalerCustomerId;
        }
    }
}