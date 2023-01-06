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

    public sealed class ZPAApplicationSegmentInspectionCommonAppsDtoAppsConfigGetArgs : global::Pulumi.ResourceArgs
    {
        [Input("allowOptions")]
        public Input<bool>? AllowOptions { get; set; }

        [Input("appId")]
        public Input<string>? AppId { get; set; }

        [Input("appTypes")]
        private InputList<string>? _appTypes;

        /// <summary>
        /// Indicates the type of application as inspection. Supported value: `INSPECT`
        /// </summary>
        public InputList<string> AppTypes
        {
            get => _appTypes ?? (_appTypes = new InputList<string>());
            set => _appTypes = value;
        }

        /// <summary>
        /// Port for the Inspection Application Segment.
        /// </summary>
        [Input("applicationPort")]
        public Input<string>? ApplicationPort { get; set; }

        /// <summary>
        /// Protocol for the Inspection Application Segment.. Supported values: `HTTP` and `HTTPS`
        /// </summary>
        [Input("applicationProtocol")]
        public Input<string>? ApplicationProtocol { get; set; }

        /// <summary>
        /// ID of the signing certificate. This field is required if the applicationProtocol is set to `HTTPS`. The certificateId is not supported if the applicationProtocol is set to `HTTP`.
        /// </summary>
        [Input("certificateId")]
        public Input<string>? CertificateId { get; set; }

        [Input("certificateName")]
        public Input<string>? CertificateName { get; set; }

        [Input("cname")]
        public Input<string>? Cname { get; set; }

        /// <summary>
        /// (Optional) Description of the application.
        /// </summary>
        [Input("description")]
        public Input<string>? Description { get; set; }

        /// <summary>
        /// Domain name of the Inspection Application Segment.
        /// </summary>
        [Input("domain")]
        public Input<string>? Domain { get; set; }

        /// <summary>
        /// Whether this application is enabled or not
        /// </summary>
        [Input("enabled")]
        public Input<bool>? Enabled { get; set; }

        [Input("hidden")]
        public Input<bool>? Hidden { get; set; }

        [Input("id")]
        public Input<string>? Id { get; set; }

        [Input("localDomain")]
        public Input<string>? LocalDomain { get; set; }

        /// <summary>
        /// Name of the Inspection Application Segment.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        [Input("portal")]
        public Input<bool>? Portal { get; set; }

        [Input("trustUntrustedCert")]
        public Input<bool>? TrustUntrustedCert { get; set; }

        public ZPAApplicationSegmentInspectionCommonAppsDtoAppsConfigGetArgs()
        {
        }
        public static new ZPAApplicationSegmentInspectionCommonAppsDtoAppsConfigGetArgs Empty => new ZPAApplicationSegmentInspectionCommonAppsDtoAppsConfigGetArgs();
    }
}
