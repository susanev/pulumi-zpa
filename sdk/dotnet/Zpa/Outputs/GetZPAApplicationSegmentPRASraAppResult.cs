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
    public sealed class GetZPAApplicationSegmentPRASraAppResult
    {
        public readonly string AppId;
        /// <summary>
        /// (string) Port for the Privileged Remote Accessvalues: `RDP` and `SSH`
        /// </summary>
        public readonly string ApplicationPort;
        /// <summary>
        /// (string) Protocol for the Privileged Remote Access. Supported values: `RDP` and `SSH`
        /// </summary>
        public readonly string ApplicationProtocol;
        public readonly string CertificateId;
        public readonly string CertificateName;
        /// <summary>
        /// (string) - Parameter required when `application_protocol` is of type `RDP`
        /// </summary>
        public readonly string ConnectionSecurity;
        /// <summary>
        /// (string) Description of the application.
        /// </summary>
        public readonly string Description;
        public readonly string Domain;
        /// <summary>
        /// (bool) Whether this application is enabled or not
        /// </summary>
        public readonly bool Enabled;
        public readonly bool Hidden;
        public readonly string Id;
        /// <summary>
        /// The name of the PRA Application Segment to be exported.
        /// </summary>
        public readonly string Name;
        public readonly bool Portal;

        [OutputConstructor]
        private GetZPAApplicationSegmentPRASraAppResult(
            string appId,

            string applicationPort,

            string applicationProtocol,

            string certificateId,

            string certificateName,

            string connectionSecurity,

            string description,

            string domain,

            bool enabled,

            bool hidden,

            string id,

            string name,

            bool portal)
        {
            AppId = appId;
            ApplicationPort = applicationPort;
            ApplicationProtocol = applicationProtocol;
            CertificateId = certificateId;
            CertificateName = certificateName;
            ConnectionSecurity = connectionSecurity;
            Description = description;
            Domain = domain;
            Enabled = enabled;
            Hidden = hidden;
            Id = id;
            Name = name;
            Portal = portal;
        }
    }
}
