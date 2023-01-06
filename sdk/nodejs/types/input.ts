// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../types/input";
import * as outputs from "../types/output";

export interface GetZPAApplicationSegmentBrowserAccessTcpPortRange {
    from?: string;
    to?: string;
}

export interface GetZPAApplicationSegmentBrowserAccessTcpPortRangeArgs {
    from?: pulumi.Input<string>;
    to?: pulumi.Input<string>;
}

export interface GetZPAApplicationSegmentBrowserAccessUdpPortRange {
    from?: string;
    to?: string;
}

export interface GetZPAApplicationSegmentBrowserAccessUdpPortRangeArgs {
    from?: pulumi.Input<string>;
    to?: pulumi.Input<string>;
}

export interface GetZPAApplicationSegmentInspectionTcpPortRange {
    from?: string;
    to?: string;
}

export interface GetZPAApplicationSegmentInspectionTcpPortRangeArgs {
    from?: pulumi.Input<string>;
    to?: pulumi.Input<string>;
}

export interface GetZPAApplicationSegmentInspectionUdpPortRange {
    from?: string;
    to?: string;
}

export interface GetZPAApplicationSegmentInspectionUdpPortRangeArgs {
    from?: pulumi.Input<string>;
    to?: pulumi.Input<string>;
}

export interface GetZPAApplicationSegmentPRATcpPortRange {
    from?: string;
    to?: string;
}

export interface GetZPAApplicationSegmentPRATcpPortRangeArgs {
    from?: pulumi.Input<string>;
    to?: pulumi.Input<string>;
}

export interface GetZPAApplicationSegmentPRAUdpPortRange {
    from?: string;
    to?: string;
}

export interface GetZPAApplicationSegmentPRAUdpPortRangeArgs {
    from?: pulumi.Input<string>;
    to?: pulumi.Input<string>;
}

export interface GetZPAApplicationSegmentTcpPortRange {
    from?: string;
    to?: string;
}

export interface GetZPAApplicationSegmentTcpPortRangeArgs {
    from?: pulumi.Input<string>;
    to?: pulumi.Input<string>;
}

export interface GetZPAApplicationSegmentUdpPortRange {
    from?: string;
    to?: string;
}

export interface GetZPAApplicationSegmentUdpPortRangeArgs {
    from?: pulumi.Input<string>;
    to?: pulumi.Input<string>;
}

export interface GetZPALSSConfigControllerConfig {
    /**
     * (string)
     */
    auditMessage?: string;
    /**
     * (string)
     */
    description?: string;
    /**
     * (bool)
     */
    enabled?: boolean;
    /**
     * (string)
     */
    filters?: string[];
    /**
     * (string)
     */
    format?: string;
    /**
     * This field defines the name of the log streaming resource.
     */
    id?: string;
    /**
     * (string)
     */
    lssHost?: string;
    /**
     * (string)
     */
    lssPort?: string;
    /**
     * This field defines the name of the log streaming resource.
     */
    name?: string;
    /**
     * (string)
     */
    sourceLogType?: string;
    useTls?: boolean;
}

export interface GetZPALSSConfigControllerConfigArgs {
    /**
     * (string)
     */
    auditMessage?: pulumi.Input<string>;
    /**
     * (string)
     */
    description?: pulumi.Input<string>;
    /**
     * (bool)
     */
    enabled?: pulumi.Input<boolean>;
    /**
     * (string)
     */
    filters?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * (string)
     */
    format?: pulumi.Input<string>;
    /**
     * This field defines the name of the log streaming resource.
     */
    id?: pulumi.Input<string>;
    /**
     * (string)
     */
    lssHost?: pulumi.Input<string>;
    /**
     * (string)
     */
    lssPort?: pulumi.Input<string>;
    /**
     * This field defines the name of the log streaming resource.
     */
    name?: pulumi.Input<string>;
    /**
     * (string)
     */
    sourceLogType?: pulumi.Input<string>;
    useTls?: pulumi.Input<boolean>;
}

export interface ZPAApplicationSegmentBrowserAccessClientlessApp {
    /**
     * If you want ZPA to forward unauthenticated HTTP preflight OPTIONS requests from the browser to the app.. Supported values: `true` and `false`
     */
    allowOptions?: pulumi.Input<boolean>;
    /**
     * Port for the BA app.
     */
    applicationPort: pulumi.Input<string>;
    /**
     * Protocol for the BA app. Supported values: `HTTP` and `HTTPS`
     */
    applicationProtocol: pulumi.Input<string>;
    certificateId: pulumi.Input<string>;
    /**
     * (Optional)
     */
    cname?: pulumi.Input<string>;
    /**
     * (Optional) Description of the application.
     */
    description?: pulumi.Input<string>;
    /**
     * Domain name or IP address of the BA app.
     */
    domain?: pulumi.Input<string>;
    /**
     * (Optional) - Whether this app is enabled or not.
     */
    enabled?: pulumi.Input<boolean>;
    /**
     * (Optional)
     */
    hidden?: pulumi.Input<boolean>;
    id?: pulumi.Input<string>;
    /**
     * (Optional)
     */
    localDomain?: pulumi.Input<string>;
    /**
     * Name of the application.
     */
    name: pulumi.Input<string>;
    /**
     * (Optional)
     */
    path?: pulumi.Input<string>;
    /**
     * (Optional)
     */
    trustUntrustedCert?: pulumi.Input<boolean>;
}

export interface ZPAApplicationSegmentBrowserAccessServerGroup {
    ids: pulumi.Input<pulumi.Input<string>[]>;
}

export interface ZPAApplicationSegmentInspectionCommonAppsDto {
    appsConfigs?: pulumi.Input<pulumi.Input<inputs.ZPAApplicationSegmentInspectionCommonAppsDtoAppsConfig>[]>;
}

export interface ZPAApplicationSegmentInspectionCommonAppsDtoAppsConfig {
    allowOptions?: pulumi.Input<boolean>;
    appId?: pulumi.Input<string>;
    /**
     * Indicates the type of application as inspection. Supported value: `INSPECT`
     */
    appTypes?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * Port for the Inspection Application Segment.
     */
    applicationPort?: pulumi.Input<string>;
    /**
     * Protocol for the Inspection Application Segment.. Supported values: `HTTP` and `HTTPS`
     */
    applicationProtocol?: pulumi.Input<string>;
    /**
     * ID of the signing certificate. This field is required if the applicationProtocol is set to `HTTPS`. The certificateId is not supported if the applicationProtocol is set to `HTTP`.
     */
    certificateId?: pulumi.Input<string>;
    certificateName?: pulumi.Input<string>;
    cname?: pulumi.Input<string>;
    /**
     * (Optional) Description of the application.
     */
    description?: pulumi.Input<string>;
    /**
     * Domain name of the Inspection Application Segment.
     */
    domain?: pulumi.Input<string>;
    /**
     * Whether this application is enabled or not
     */
    enabled?: pulumi.Input<boolean>;
    hidden?: pulumi.Input<boolean>;
    id?: pulumi.Input<string>;
    localDomain?: pulumi.Input<string>;
    /**
     * Name of the Inspection Application Segment.
     */
    name?: pulumi.Input<string>;
    portal?: pulumi.Input<boolean>;
    trustUntrustedCert?: pulumi.Input<boolean>;
}

export interface ZPAApplicationSegmentInspectionServerGroup {
    ids: pulumi.Input<pulumi.Input<string>[]>;
}

export interface ZPAApplicationSegmentPRACommonAppsDto {
    appsConfigs?: pulumi.Input<pulumi.Input<inputs.ZPAApplicationSegmentPRACommonAppsDtoAppsConfig>[]>;
}

export interface ZPAApplicationSegmentPRACommonAppsDtoAppsConfig {
    allowOptions?: pulumi.Input<boolean>;
    appId?: pulumi.Input<string>;
    /**
     * Indicates the type of application as Privileged Remote Access. Supported value: `SECURE_REMOTE_ACCESS`
     */
    appTypes?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * Port for the Privileged Remote Access
     */
    applicationPort?: pulumi.Input<string>;
    /**
     * Protocol for the Privileged Remote Access. Supported values: `RDP` and `SSH`
     */
    applicationProtocol?: pulumi.Input<string>;
    cname?: pulumi.Input<string>;
    /**
     * Parameter required when `applicationProtocol` is of type `RDP`
     */
    connectionSecurity?: pulumi.Input<string>;
    /**
     * (Optional) Description of the application.
     */
    description?: pulumi.Input<string>;
    /**
     * Domain name of the Privileged Remote Access
     */
    domain?: pulumi.Input<string>;
    /**
     * Whether this application is enabled or not
     */
    enabled?: pulumi.Input<boolean>;
    hidden?: pulumi.Input<boolean>;
    id?: pulumi.Input<string>;
    localDomain?: pulumi.Input<string>;
    /**
     * Name of the Privileged Remote Access
     */
    name?: pulumi.Input<string>;
    portal?: pulumi.Input<boolean>;
}

export interface ZPAApplicationSegmentPRAServerGroup {
    ids: pulumi.Input<pulumi.Input<string>[]>;
}

export interface ZPAApplicationSegmentServerGroup {
    ids: pulumi.Input<pulumi.Input<string>[]>;
}

export interface ZPABrowserAccessClientlessApp {
    allowOptions?: pulumi.Input<boolean>;
    applicationPort: pulumi.Input<string>;
    applicationProtocol: pulumi.Input<string>;
    certificateId: pulumi.Input<string>;
    cname?: pulumi.Input<string>;
    description?: pulumi.Input<string>;
    domain?: pulumi.Input<string>;
    enabled?: pulumi.Input<boolean>;
    hidden?: pulumi.Input<boolean>;
    id?: pulumi.Input<string>;
    localDomain?: pulumi.Input<string>;
    name: pulumi.Input<string>;
    path?: pulumi.Input<string>;
    trustUntrustedCert?: pulumi.Input<boolean>;
}

export interface ZPABrowserAccessServerGroup {
    ids: pulumi.Input<pulumi.Input<string>[]>;
}

export interface ZPAInspectionCustomControlsAssociatedInspectionProfileName {
    ids?: pulumi.Input<pulumi.Input<string>[]>;
}

export interface ZPAInspectionCustomControlsRule {
    conditions?: pulumi.Input<inputs.ZPAInspectionCustomControlsRuleConditions>;
    names?: pulumi.Input<pulumi.Input<string>[]>;
    type?: pulumi.Input<string>;
}

export interface ZPAInspectionCustomControlsRuleConditions {
    lhs?: pulumi.Input<string>;
    op?: pulumi.Input<string>;
    rhs?: pulumi.Input<string>;
}

export interface ZPAInspectionProfileControlsInfo {
    /**
     * (Optional) Control types. Supported Values: `CUSTOM`, `PREDEFINED`, `ZSCALER`
     */
    controlType?: pulumi.Input<string>;
    /**
     * (Optional) Control information counts `Long`
     */
    count?: pulumi.Input<string>;
}

export interface ZPAInspectionProfileCustomControl {
    /**
     * The action of the predefined control. Supported values: `PASS`, `BLOCK` and `REDIRECT`
     */
    action: pulumi.Input<string>;
    /**
     * Value for the predefined controls action. This field is only required if the action is set to REDIRECT. This field is only required if the action is set to `REDIRECT`.
     */
    actionValue?: pulumi.Input<string>;
    /**
     * ID of the predefined control
     */
    id: pulumi.Input<string>;
}

export interface ZPAInspectionProfilePredefinedControl {
    /**
     * The action of the predefined control. Supported values: `PASS`, `BLOCK` and `REDIRECT`
     */
    action: pulumi.Input<string>;
    /**
     * Value for the predefined controls action. This field is only required if the action is set to REDIRECT. This field is only required if the action is set to `REDIRECT`.
     */
    actionValue?: pulumi.Input<string>;
    /**
     * ID of the predefined control
     */
    id: pulumi.Input<string>;
}

export interface ZPALSSConfigControllerConfig {
    /**
     * (Optional)
     */
    auditMessage?: pulumi.Input<string>;
    /**
     * (Optional)
     */
    description?: pulumi.Input<string>;
    /**
     * (Optional)
     */
    enabled?: pulumi.Input<boolean>;
    /**
     * (Optional)
     */
    filters?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * The format of the LSS resource. The supported formats are: `JSON`, `CSV`, and `TSV`
     */
    format: pulumi.Input<string>;
    /**
     * App Connector Group ID(s) where logs will be forwarded to.
     */
    id?: pulumi.Input<string>;
    /**
     * The IP or FQDN of the SIEM (Log Receiver) where logs will be forwarded to.
     */
    lssHost: pulumi.Input<string>;
    /**
     * The destination port of the SIEM (Log Receiver) where logs will be forwarded to.
     */
    lssPort: pulumi.Input<string>;
    /**
     * (Optional)
     */
    name: pulumi.Input<string>;
    /**
     * Refer to the log type documentation
     */
    sourceLogType: pulumi.Input<string>;
    /**
     * (Optional)
     */
    useTls?: pulumi.Input<boolean>;
}

export interface ZPALSSConfigControllerConnectorGroup {
    /**
     * App Connector Group ID(s) where logs will be forwarded to.
     */
    ids?: pulumi.Input<pulumi.Input<string>[]>;
}

export interface ZPALSSConfigControllerPolicyRuleResource {
    /**
     * (Optional)
     */
    action?: pulumi.Input<string>;
    actionId?: pulumi.Input<string>;
    bypassDefaultRule?: pulumi.Input<boolean>;
    /**
     * (Optional)
     */
    conditions?: pulumi.Input<pulumi.Input<inputs.ZPALSSConfigControllerPolicyRuleResourceCondition>[]>;
    /**
     * (Optional)
     */
    customMsg?: pulumi.Input<string>;
    defaultRule?: pulumi.Input<boolean>;
    /**
     * (Optional)
     */
    description?: pulumi.Input<string>;
    /**
     * App Connector Group ID(s) where logs will be forwarded to.
     */
    id?: pulumi.Input<string>;
    lssDefaultRule?: pulumi.Input<boolean>;
    /**
     * (Optional)
     */
    name: pulumi.Input<string>;
    /**
     * (Optional) - Supported values are: `AND` or `OR`
     */
    operator?: pulumi.Input<string>;
    policySetId?: pulumi.Input<string>;
    policyType?: pulumi.Input<string>;
    priority?: pulumi.Input<string>;
    reauthDefaultRule?: pulumi.Input<boolean>;
    reauthIdleTimeout?: pulumi.Input<string>;
    reauthTimeout?: pulumi.Input<string>;
    ruleOrder?: pulumi.Input<string>;
    zpnInspectionProfileId?: pulumi.Input<string>;
}

export interface ZPALSSConfigControllerPolicyRuleResourceCondition {
    /**
     * (Optional)
     */
    negated?: pulumi.Input<boolean>;
    operands?: pulumi.Input<pulumi.Input<inputs.ZPALSSConfigControllerPolicyRuleResourceConditionOperand>[]>;
    /**
     * (Optional) - Supported values are: `AND` or `OR`
     */
    operator: pulumi.Input<string>;
}

export interface ZPALSSConfigControllerPolicyRuleResourceConditionOperand {
    /**
     * (Optional) This is for specifying the policy critiera. Supported values: `APP`, `APP_GROUP`, `CLIENT_TYPE`, `TRUSTED_NETWORK`, `SAML`, `SCIM`, `SCIM_GROUP`
     */
    objectType: pulumi.Input<string>;
    /**
     * (Optional) The below values are supported when choosing `objectType` of type `CLIENT_TYPE`.
     */
    values?: pulumi.Input<pulumi.Input<string>[]>;
}

export interface ZPAPolicyAccessForwardingRuleCondition {
    id?: pulumi.Input<string>;
    negated?: pulumi.Input<boolean>;
    operands?: pulumi.Input<pulumi.Input<inputs.ZPAPolicyAccessForwardingRuleConditionOperand>[]>;
    operator: pulumi.Input<string>;
}

export interface ZPAPolicyAccessForwardingRuleConditionOperand {
    id?: pulumi.Input<string>;
    idpId?: pulumi.Input<string>;
    lhs: pulumi.Input<string>;
    name?: pulumi.Input<string>;
    objectType: pulumi.Input<string>;
    rhs?: pulumi.Input<string>;
    rhsLists?: pulumi.Input<pulumi.Input<string>[]>;
}

export interface ZPAPolicyAccessInspectionRuleCondition {
    id?: pulumi.Input<string>;
    negated?: pulumi.Input<boolean>;
    operands?: pulumi.Input<pulumi.Input<inputs.ZPAPolicyAccessInspectionRuleConditionOperand>[]>;
    operator: pulumi.Input<string>;
}

export interface ZPAPolicyAccessInspectionRuleConditionOperand {
    id?: pulumi.Input<string>;
    idpId?: pulumi.Input<string>;
    lhs: pulumi.Input<string>;
    name?: pulumi.Input<string>;
    objectType: pulumi.Input<string>;
    rhs?: pulumi.Input<string>;
    rhsLists?: pulumi.Input<pulumi.Input<string>[]>;
}

export interface ZPAPolicyAccessRuleAppConnectorGroup {
    /**
     * (Optional) The ID of a server group resource
     */
    ids?: pulumi.Input<pulumi.Input<string>[]>;
}

export interface ZPAPolicyAccessRuleAppServerGroup {
    /**
     * (Optional) The ID of a server group resource
     */
    ids?: pulumi.Input<pulumi.Input<string>[]>;
}

export interface ZPAPolicyAccessRuleCondition {
    /**
     * (Optional) The ID of a server group resource
     */
    id?: pulumi.Input<string>;
    /**
     * (Optional) Supported values: ``true`` or ``false``
     */
    negated?: pulumi.Input<boolean>;
    /**
     * (Optional) - Operands block must be repeated if multiple per `objectType` conditions are to be added to the rule.
     */
    operands?: pulumi.Input<pulumi.Input<inputs.ZPAPolicyAccessRuleConditionOperand>[]>;
    /**
     * (Optional) Supported values: ``AND``, and ``OR``
     */
    operator: pulumi.Input<string>;
}

export interface ZPAPolicyAccessRuleConditionOperand {
    /**
     * (Optional) The ID of a server group resource
     */
    id?: pulumi.Input<string>;
    /**
     * (Optional)
     */
    idpId?: pulumi.Input<string>;
    /**
     * (Optional) LHS must always carry the string value ``id`` or the attribute ID of the resource being associated with the rule.
     */
    lhs: pulumi.Input<string>;
    /**
     * (Optional)
     */
    name?: pulumi.Input<string>;
    /**
     * (Optional) This is for specifying the policy critiera. Supported values: `APP`, `APP_GROUP`, `SAML`, `IDP`, `CLIENT_TYPE`, `TRUSTED_NETWORK`, `POSTURE`, `SCIM`, `SCIM_GROUP`, and `CLOUD_CONNECTOR_GROUP`. `TRUSTED_NETWORK`, and `CLIENT_TYPE`.
     */
    objectType: pulumi.Input<string>;
    /**
     * (Optional) RHS is either the ID attribute of a resource or fixed string value. Refer to the chart below for further details.
     */
    rhs?: pulumi.Input<string>;
    rhsLists?: pulumi.Input<pulumi.Input<string>[]>;
}

export interface ZPAPolicyAccessTimeOutRuleCondition {
    id?: pulumi.Input<string>;
    negated?: pulumi.Input<boolean>;
    operands?: pulumi.Input<pulumi.Input<inputs.ZPAPolicyAccessTimeOutRuleConditionOperand>[]>;
    operator: pulumi.Input<string>;
}

export interface ZPAPolicyAccessTimeOutRuleConditionOperand {
    id?: pulumi.Input<string>;
    idpId?: pulumi.Input<string>;
    lhs: pulumi.Input<string>;
    name?: pulumi.Input<string>;
    objectType: pulumi.Input<string>;
    rhs?: pulumi.Input<string>;
    rhsLists?: pulumi.Input<pulumi.Input<string>[]>;
}

export interface ZPASegmentGroupApplication {
    id?: pulumi.Input<string>;
}

export interface ZPAServerGroupAppConnectorGroup {
    ids?: pulumi.Input<pulumi.Input<string>[]>;
}

export interface ZPAServerGroupApplication {
    ids?: pulumi.Input<pulumi.Input<string>[]>;
}

export interface ZPAServerGroupServer {
    ids?: pulumi.Input<pulumi.Input<string>[]>;
}

export interface ZPAServiceEdgeGroupServiceEdge {
    ids?: pulumi.Input<pulumi.Input<string>[]>;
}

export interface ZPAServiceEdgeGroupTrustedNetwork {
    ids?: pulumi.Input<pulumi.Input<string>[]>;
}
