// Copyright 2016-2018, Pulumi Corporation.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package zpa

import (
	"fmt"
	"path/filepath"

	"github.com/pulumi/pulumi-terraform-bridge/v3/pkg/tfbridge"
	shim "github.com/pulumi/pulumi-terraform-bridge/v3/pkg/tfshim"
	shimv2 "github.com/pulumi/pulumi-terraform-bridge/v3/pkg/tfshim/sdk-v2"
	"github.com/pulumi/pulumi/sdk/v3/go/common/resource"
	"github.com/zscaler/pulumi-zpa/provider/pkg/version"
	"github.com/zscaler/terraform-provider-zpa/v2/zpa"
)

// all of the token components used below.
const (
	// This variable controls the default name of the package in the package
	// registries for nodejs and python:
	zpaPkg = "zpa"
	// modules:
	segmentGroupMod          = "SegmentGroup"
	serverGroupMod           = "ServerGroup"
	applicationSegmentMod    = "ApplicationSegment"
	appServerMod             = "ApplicationServer"
	appConnectorGroupMod     = "AppConnectorGroup"
	serviceEdgeGroupMod      = "ServiceEdge"
	cloudConnectorGroupMod   = "CloudConnectorGroup"
	browserCertificateMod    = "BrowserCertificate"
	enrollmentCertificateMod = "EnrollmentCertificate"
	inspectionMod            = "Inspection"
	lssConfigMod             = "LSSConfig"
	policyTypeMod            = "PolicyType"
	acessPolicyMod           = "AccessPolicy"
	forwardPolicyMod         = "ForwardPolicy"
	inspectionPolicyMod      = "InspectionPolicy"
	timeoutPolicyMod         = "TimeoutPolicy"
	idpControllerMod         = "idpController"
	samlAttributeMod         = "SamlAttribute"
	scimAttributeHeaderMod   = "ScimAttribute"
	scimGroupMod             = "scimGroup"
	machineGroupMod          = "machineGroup"
	postureProfileMod        = "postureProfile"
	trustedNetworkMod        = "trustedNetwork"
	provisioningKeyMod       = "ProvisioningKey"
)

func preConfigureCallback(vars resource.PropertyMap, c shim.ResourceConfig) error {
	return nil
}

func refProviderLicense(license tfbridge.TFProviderLicense) *tfbridge.TFProviderLicense {
	return &license
}

// Provider returns additional overlaid schema and metadata associated with the provider..
func Provider() tfbridge.ProviderInfo {
	// Instantiate the Terraform provider
	p := shimv2.NewProvider(zpa.Provider())

	// Create a Pulumi provider mapping
	prov := tfbridge.ProviderInfo{
		P:                       p,
		Name:                    "zpa",
		Description:             "A Pulumi package for creating and managing zpa cloud resources.",
		Keywords:                []string{"pulumi", "zpa", "zscaler", "category/cloud"},
		TFProviderLicense:       refProviderLicense(tfbridge.MITLicenseType),
		License:                 "MIT",
		LogoURL:                 "https://www.zscaler.com/themes/custom/zscaler/logo.svg", //nolint:golint,lll
		Homepage:                "https://www.zscaler.com",
		Repository:              "https://github.com/zscaler/pulumi-zpa",
		PluginDownloadURL:       "github://api.github.com/zscaler",
		GitHubOrg:               "zscaler",
		Publisher:               "Zscaler",
		DisplayName:             "Zscaler Private Access",
		TFProviderModuleVersion: "v2",
		Config: map[string]*tfbridge.SchemaInfo{
			"zpa_client_id": {
				Default: &tfbridge.DefaultInfo{
					EnvVars: []string{"ZPA_CLIENT_ID"},
				},
			},
			"zpa_client_secret": {
				Default: &tfbridge.DefaultInfo{
					EnvVars: []string{"ZPA_CLIENT_SECRET"},
				},
			},
			"zpa_customer_id": {
				Default: &tfbridge.DefaultInfo{
					EnvVars: []string{"ZPA_CUSTOMER_ID"},
				},
			},
			"zpa_cloud": {
				Default: &tfbridge.DefaultInfo{
					EnvVars: []string{"ZPA_CLOUD"},
				},
			},
		},
		PreConfigureCallback: preConfigureCallback,
		Resources: map[string]*tfbridge.ResourceInfo{
			"zpa_app_connector_group": {Tok: tfbridge.MakeResource(zpaPkg, appConnectorGroupMod, "ConnectorGroup"),
				Docs: &tfbridge.DocInfo{Source: "zpa_app_connector_group.md"},
			},
			"zpa_service_edge_group": {Tok: tfbridge.MakeResource(zpaPkg, serviceEdgeGroupMod, "ServiceEdgeGroup"),
				Docs: &tfbridge.DocInfo{Source: "zpa_service_edge_group.md"},
			},
			"zpa_segment_group": {Tok: tfbridge.MakeResource(zpaPkg, segmentGroupMod, "SegmentGroup"),
				Docs: &tfbridge.DocInfo{Source: "zpa_segment_group.md"},
			},
			"zpa_server_group": {Tok: tfbridge.MakeResource(zpaPkg, serverGroupMod, "ServerGroup"),
				Docs: &tfbridge.DocInfo{Source: "zpa_server_group.md"},
			},
			"zpa_application_segment": {Tok: tfbridge.MakeResource(zpaPkg, applicationSegmentMod, "ApplicationSegment"),
				Docs: &tfbridge.DocInfo{Source: "zpa_application_segment.md"},
			},
			"zpa_application_segment_browser_access": {Tok: tfbridge.MakeResource(zpaPkg, applicationSegmentMod, "ApplicationSegmentBrowserAccess"),
				Docs: &tfbridge.DocInfo{Source: "zpa_application_segment_browser_access.md"},
			},
			"zpa_browser_access": {Tok: tfbridge.MakeResource(zpaPkg, applicationSegmentMod, "BrowserAccess"),
				Docs: noUpstreamDocs(),
			},
			"zpa_application_segment_inspection": {Tok: tfbridge.MakeResource(zpaPkg, applicationSegmentMod, "ApplicationSegmentInspection"),
				Docs: &tfbridge.DocInfo{Source: "zpa_application_segment_inspection.md"},
			},
			"zpa_application_segment_pra": {Tok: tfbridge.MakeResource(zpaPkg, applicationSegmentMod, "ApplicationSegmentPRA"),
				Docs: &tfbridge.DocInfo{Source: "zpa_application_segment_pra.md"},
			},
			"zpa_application_server": {Tok: tfbridge.MakeResource(zpaPkg, appServerMod, "ApplicationServer"),
				Docs: &tfbridge.DocInfo{Source: "zpa_application_server.md"},
			},
			"zpa_inspection_custom_controls": {Tok: tfbridge.MakeResource(zpaPkg, inspectionMod, "InspectionCustomControls"),
				Docs: &tfbridge.DocInfo{Source: "zpa_inspection_custom_controls.md"},
			},
			"zpa_inspection_profile": {Tok: tfbridge.MakeResource(zpaPkg, inspectionMod, "InspectionProfile"),
				Docs: &tfbridge.DocInfo{Source: "zpa_inspection_profile.md"},
			},
			"zpa_lss_config_controller": {Tok: tfbridge.MakeResource(zpaPkg, lssConfigMod, "LSSConfigController"),
				Docs: &tfbridge.DocInfo{Source: "zpa_lss_config_controller.md"},
			},
			"zpa_policy_access_rule": {Tok: tfbridge.MakeResource(zpaPkg, acessPolicyMod, "PolicyAccessRule"),
				Docs: &tfbridge.DocInfo{Source: "zpa_policy_access_rule.md"},
			},
			"zpa_policy_forwarding_rule": {Tok: tfbridge.MakeResource(zpaPkg, forwardPolicyMod, "PolicyAccessForwardingRule"),
				Docs: &tfbridge.DocInfo{Source: "zpa_policy_forwarding_rule.md"},
			},

			"zpa_policy_timeout_rule": {Tok: tfbridge.MakeResource(zpaPkg, timeoutPolicyMod, "PolicyAccessTimeOutRule"),
				Docs: &tfbridge.DocInfo{Source: "zpa_policy_timeout_rule.md"},
			},
			"zpa_policy_inspection_rule": {Tok: tfbridge.MakeResource(zpaPkg, inspectionPolicyMod, "PolicyAccessInspectionRule"),
				Docs: &tfbridge.DocInfo{Source: "zpa_policy_inspection_rule.md"},
			},
			"zpa_provisioning_key": {Tok: tfbridge.MakeResource(zpaPkg, provisioningKeyMod, "ProvisioningKey"),
				Fields: map[string]*tfbridge.SchemaInfo{
					// Rename field to prevent this error in the DotNet SDK generation:
					// sdk/dotnet/ProvisioningKey.cs(69,31): error CS0542: 'ProvisioningKey': member names cannot be the same as their enclosing type [sdk/dotnet/zscaler.Zpa.csproj]
					"provisioning_key": {
						Name: "key",
					},
				},
				Docs: &tfbridge.DocInfo{Source: "zpa_provisioning_key.md"},
			},
		},
		DataSources: map[string]*tfbridge.DataSourceInfo{
			"zpa_app_connector_group": {Tok: tfbridge.MakeDataSource(zpaPkg, appConnectorGroupMod, "getAppConnectorGroup"),
				Docs: &tfbridge.DocInfo{Source: "zpa_app_connector_group.md"},
			},
			"zpa_service_edge_group": {Tok: tfbridge.MakeDataSource(zpaPkg, serviceEdgeGroupMod, "getServiceEdgeGroup"),
				Docs: &tfbridge.DocInfo{Source: "zpa_service_edge_group.md"},
			},
			"zpa_segment_group": {Tok: tfbridge.MakeDataSource(zpaPkg, segmentGroupMod, "getSegmentGroup"),
				Docs: &tfbridge.DocInfo{Source: "zpa_segment_group.md"},
			},
			"zpa_server_group": {Tok: tfbridge.MakeDataSource(zpaPkg, serverGroupMod, "getServerGroup"),
				Docs: &tfbridge.DocInfo{Source: "zpa_server_group.md"},
			},
			"zpa_application_segment": {Tok: tfbridge.MakeDataSource(zpaPkg, applicationSegmentMod, "getApplicationSegment"),
				Docs: &tfbridge.DocInfo{Source: "zpa_application_segment.md"},
			},
			"zpa_application_segment_browser_access": {Tok: tfbridge.MakeDataSource(zpaPkg, applicationSegmentMod, "getApplicationSegmentBrowserAccess"),
				Docs: &tfbridge.DocInfo{Source: "zpa_application_segment_browser_access.md"},
			},
			"zpa_browser_access": {Tok: tfbridge.MakeDataSource(zpaPkg, applicationSegmentMod, "getBrowserAccess"),
				Docs: noUpstreamDocs(),
			},
			"zpa_application_segment_inspection": {Tok: tfbridge.MakeDataSource(zpaPkg, applicationSegmentMod, "getApplicationSegmentInspection"),
				Docs: &tfbridge.DocInfo{Source: "zpa_application_segment_inspection.md"},
			},
			"zpa_inspection_custom_controls": {Tok: tfbridge.MakeDataSource(zpaPkg, inspectionMod, "getInspectionCustomControls"),
				Docs: &tfbridge.DocInfo{Source: "zpa_inspection_custom_controls.md"},
			},
			"zpa_inspection_profile": {Tok: tfbridge.MakeDataSource(zpaPkg, inspectionMod, "getInspectionProfile"),
				Docs: &tfbridge.DocInfo{Source: "zpa_inspection_profile.md"},
			},
			"zpa_inspection_predefined_controls": {Tok: tfbridge.MakeDataSource(zpaPkg, inspectionMod, "getInspectionPredefinedControls"),
				Docs: &tfbridge.DocInfo{Source: "zpa_inspection_predefined_controls.md"},
			},
			"zpa_inspection_all_predefined_controls": {Tok: tfbridge.MakeDataSource(zpaPkg, inspectionMod, "getInspectionAllPredefinedControls"),
				Docs: &tfbridge.DocInfo{Source: "zpa_inspection_all_predefined_controls.md"},
			},
			"zpa_application_segment_pra": {Tok: tfbridge.MakeDataSource(zpaPkg, applicationSegmentMod, "getApplicationSegmentPRA"),
				Docs: &tfbridge.DocInfo{Source: "zpa_application_segment_pra.md"},
			},
			"zpa_application_server": {Tok: tfbridge.MakeDataSource(zpaPkg, appServerMod, "getApplicationServer"),
				Docs: &tfbridge.DocInfo{Source: "zpa_application_server.md"},
			},
			"zpa_lss_config_controller": {Tok: tfbridge.MakeDataSource(zpaPkg, lssConfigMod, "getLSSConfigController"),
				Docs: &tfbridge.DocInfo{Source: "zpa_lss_config_controller.md"},
			},
			"zpa_lss_config_client_types": {Tok: tfbridge.MakeDataSource(zpaPkg, lssConfigMod, "getLSSClientTypes"),
				Docs: &tfbridge.DocInfo{Source: "zpa_lss_config_client_types.md"},
			},
			"zpa_lss_config_log_type_formats": {Tok: tfbridge.MakeDataSource(zpaPkg, lssConfigMod, "getLSSLogTypeFormats"),
				Docs: &tfbridge.DocInfo{Source: "zpa_lss_config_log_type_formats.md"},
			},
			"zpa_lss_config_status_codes": {Tok: tfbridge.MakeDataSource(zpaPkg, lssConfigMod, "getLSSStatusCodes"),
				Docs: &tfbridge.DocInfo{Source: "zpa_lss_config_status_codes.md"},
			},
			"zpa_provisioning_key": {Tok: tfbridge.MakeDataSource(zpaPkg, provisioningKeyMod, "getProvisioningKey"),
				Docs: &tfbridge.DocInfo{Source: "zpa_provisioning_key.md"},
			},
			"zpa_ba_certificate": {Tok: tfbridge.MakeDataSource(zpaPkg, browserCertificateMod, "getBaCertificate"),
				Docs: &tfbridge.DocInfo{Source: "zpa_ba_certificate.md"},
			},
			"zpa_cloud_connector_group": {Tok: tfbridge.MakeDataSource(zpaPkg, cloudConnectorGroupMod, "getCloudConnectorGroup"),
				Docs: &tfbridge.DocInfo{Source: "zpa_cloud_connector_group.md"},
			},
			"zpa_enrollment_cert": {Tok: tfbridge.MakeDataSource(zpaPkg, enrollmentCertificateMod, "getEnrollmentCert"),
				Docs: &tfbridge.DocInfo{Source: "zpa_enrollment_cert.md"},
			},
			"zpa_idp_controller": {Tok: tfbridge.MakeDataSource(zpaPkg, idpControllerMod, "getIdPController"),
				Docs: &tfbridge.DocInfo{Source: "zpa_idp_controller.md"},
			},
			"zpa_policy_type": {Tok: tfbridge.MakeDataSource(zpaPkg, policyTypeMod, "getPolicyType"),
				Docs: &tfbridge.DocInfo{Source: "zpa_policy_type.md"},
			},
			"zpa_saml_attribute": {Tok: tfbridge.MakeDataSource(zpaPkg, samlAttributeMod, "getSAMLAttribute"),
				Docs: &tfbridge.DocInfo{Source: "zpa_saml_attribute.md"},
			},
			"zpa_scim_attribute_header": {Tok: tfbridge.MakeDataSource(zpaPkg, scimAttributeHeaderMod, "getSCIMAttributeHeader"),
				Docs: &tfbridge.DocInfo{Source: "zpa_scim_attribute_header.md"},
			},
			"zpa_scim_groups": {Tok: tfbridge.MakeDataSource(zpaPkg, scimGroupMod, "getSCIMGroups"),
				Docs: &tfbridge.DocInfo{Source: "zpa_scim_groups.md"},
			},
			"zpa_machine_group": {Tok: tfbridge.MakeDataSource(zpaPkg, machineGroupMod, "getMachineGroup"),
				Docs: &tfbridge.DocInfo{Source: "zpa_machine_group.md"},
			},
			"zpa_posture_profile": {Tok: tfbridge.MakeDataSource(zpaPkg, postureProfileMod, "getPostureProfile"),
				Docs: &tfbridge.DocInfo{Source: "zpa_posture_profile.md"},
			},
			"zpa_trusted_network": {Tok: tfbridge.MakeDataSource(zpaPkg, trustedNetworkMod, "getTrustedNetwork"),
				Docs: &tfbridge.DocInfo{Source: "zpa_trusted_network.md"},
			},
		},
		JavaScript: &tfbridge.JavaScriptInfo{
			// List any npm dependencies and their versions
			Dependencies: map[string]string{
				"@pulumi/pulumi": "^3.0.0",
			},
			DevDependencies: map[string]string{
				"@types/node": "^10.0.0", // so we can access strongly typed node definitions.
				"@types/mime": "^2.0.0",
			},
			PackageName: "@zscaler/pulumi-zpa",
			// See the documentation for tfbridge.OverlayInfo for how to lay out this
			// section, or refer to the AWS provider. Delete this section if there are
			// no overlay files.
			// Overlay: &tfbridge.OverlayInfo{},
		},
		Python: &tfbridge.PythonInfo{
			PackageName: "zscaler_pulumi_zpa",
			// List any Python dependencies and their version ranges
			Requires: map[string]string{
				"pulumi": ">=3.0.0,<4.0.0",
			},
		},
		Golang: &tfbridge.GolangInfo{
			ImportBasePath: filepath.Join(
				fmt.Sprintf("github.com/zscaler/pulumi-%[1]s/sdk/", zpaPkg),
				tfbridge.GetModuleMajorVersion(version.Version),
				"go",
				zpaPkg,
			),
			GenerateResourceContainerTypes: true,
		},
		CSharp: &tfbridge.CSharpInfo{
			RootNamespace: "zscaler.PulumiPackage",
			PackageReferences: map[string]string{
				"Pulumi": "3.*",
			},
		},
	}

	prov.SetAutonaming(255, "-")

	return prov
}

func noUpstreamDocs() *tfbridge.DocInfo {
	return &tfbridge.DocInfo{
		Markdown: []byte(" "),
	}
}
