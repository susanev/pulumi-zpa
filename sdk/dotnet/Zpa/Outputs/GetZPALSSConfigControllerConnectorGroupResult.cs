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
    public sealed class GetZPALSSConfigControllerConnectorGroupResult
    {
        /// <summary>
        /// This field defines the name of the log streaming resource.
        /// </summary>
        public readonly string Id;

        [OutputConstructor]
        private GetZPALSSConfigControllerConnectorGroupResult(string id)
        {
            Id = id;
        }
    }
}
