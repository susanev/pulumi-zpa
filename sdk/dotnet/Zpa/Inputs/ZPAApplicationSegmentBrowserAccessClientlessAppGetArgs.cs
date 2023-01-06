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

    public sealed class ZPAApplicationSegmentBrowserAccessClientlessAppGetArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// - If you want ZPA to forward unauthenticated HTTP preflight OPTIONS requests from the browser to the app.. Supported values: `true` and `false`
        /// </summary>
        [Input("allowOptions")]
        public Input<bool>? AllowOptions { get; set; }

        /// <summary>
        /// - Port for the BA app.
        /// </summary>
        [Input("applicationPort", required: true)]
        public Input<string> ApplicationPort { get; set; } = null!;

        /// <summary>
        /// - Protocol for the BA app. Supported values: `HTTP` and `HTTPS`
        /// </summary>
        [Input("applicationProtocol", required: true)]
        public Input<string> ApplicationProtocol { get; set; } = null!;

        [Input("certificateId", required: true)]
        public Input<string> CertificateId { get; set; } = null!;

        [Input("cname")]
        public Input<string>? Cname { get; set; }

        [Input("description")]
        public Input<string>? Description { get; set; }

        /// <summary>
        /// - Domain name or IP address of the BA app.
        /// </summary>
        [Input("domain")]
        public Input<string>? Domain { get; set; }

        [Input("enabled")]
        public Input<bool>? Enabled { get; set; }

        [Input("hidden")]
        public Input<bool>? Hidden { get; set; }

        [Input("id")]
        public Input<string>? Id { get; set; }

        [Input("localDomain")]
        public Input<string>? LocalDomain { get; set; }

        /// <summary>
        /// - Name of BA app.
        /// </summary>
        [Input("name", required: true)]
        public Input<string> Name { get; set; } = null!;

        [Input("path")]
        public Input<string>? Path { get; set; }

        [Input("trustUntrustedCert")]
        public Input<bool>? TrustUntrustedCert { get; set; }

        public ZPAApplicationSegmentBrowserAccessClientlessAppGetArgs()
        {
        }
        public static new ZPAApplicationSegmentBrowserAccessClientlessAppGetArgs Empty => new ZPAApplicationSegmentBrowserAccessClientlessAppGetArgs();
    }
}
