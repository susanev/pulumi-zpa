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
    public sealed class GetZPAApplicationSegmentServerGroupResult
    {
        public readonly string ConfigSpace;
        public readonly string CreationTime;
        /// <summary>
        /// Description of the application.
        /// </summary>
        public readonly string Description;
        public readonly bool DynamicDiscovery;
        /// <summary>
        /// Whether this application is enabled or not. Default: false. Supported values: `true`, `false`.
        /// </summary>
        public readonly bool Enabled;
        public readonly string Id;
        public readonly string ModifiedTime;
        public readonly string Modifiedby;
        /// <summary>
        /// Name of the application.
        /// </summary>
        public readonly string Name;

        [OutputConstructor]
        private GetZPAApplicationSegmentServerGroupResult(
            string configSpace,

            string creationTime,

            string description,

            bool dynamicDiscovery,

            bool enabled,

            string id,

            string modifiedTime,

            string modifiedby,

            string name)
        {
            ConfigSpace = configSpace;
            CreationTime = creationTime;
            Description = description;
            DynamicDiscovery = dynamicDiscovery;
            Enabled = enabled;
            Id = id;
            ModifiedTime = modifiedTime;
            Modifiedby = modifiedby;
            Name = name;
        }
    }
}
