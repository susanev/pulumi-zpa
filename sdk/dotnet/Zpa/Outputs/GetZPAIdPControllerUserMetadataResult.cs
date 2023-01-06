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
    public sealed class GetZPAIdPControllerUserMetadataResult
    {
        /// <summary>
        /// (string)
        /// </summary>
        public readonly string CertificateUrl;
        public readonly string SpBaseUrl;
        /// <summary>
        /// (string)
        /// </summary>
        public readonly string SpEntityId;
        /// <summary>
        /// (string)
        /// </summary>
        public readonly string SpMetadataUrl;
        /// <summary>
        /// (string)
        /// </summary>
        public readonly string SpPostUrl;

        [OutputConstructor]
        private GetZPAIdPControllerUserMetadataResult(
            string certificateUrl,

            string spBaseUrl,

            string spEntityId,

            string spMetadataUrl,

            string spPostUrl)
        {
            CertificateUrl = certificateUrl;
            SpBaseUrl = spBaseUrl;
            SpEntityId = spEntityId;
            SpMetadataUrl = spMetadataUrl;
            SpPostUrl = spPostUrl;
        }
    }
}
