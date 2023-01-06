// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../types/input";
import * as outputs from "../types/output";

export interface GetZPAAppConnectorGroupConnector {
    appconnectorGroupId: string;
    appconnectorGroupName: string;
    applicationStartTime: string;
    controlChannelStatus: string;
    creationTime: string;
    ctrlBrokerName: string;
    currentVersion: string;
    description: string;
    /**
     * (String) Whether this App Connector Group is enabled or not. Default value: `true`. Supported values: `true`, `false`
     */
    enabled: boolean;
    enrollmentCert: {[key: string]: any};
    expectedUpgradeTime: string;
    expectedVersion: string;
    fingerprint: string;
    /**
     * ID of the App Connector Group.
     */
    id: string;
    ipacl: string;
    issuedCertId: string;
    lastBrokerConnectTime: string;
    lastBrokerConnectTimeDuration: string;
    lastBrokerDisconnectTime: string;
    lastBrokerDisconnectTimeDuration: string;
    lastUpgradeTime: string;
    /**
     * (String) Latitude of the App Connector Group. Integer or decimal. With values in the range of `-90` to `90`
     */
    latitude: string;
    /**
     * (String) Location of the App Connector Group.
     */
    location: string;
    /**
     * (String) Longitude of the App Connector Group. Integer or decimal. With values in the range of `-180` to `180`
     */
    longitude: string;
    modifiedTime: string;
    modifiedby: string;
    /**
     * Name of the App Connector Group.
     */
    name: string;
    platform: string;
    previousVersion: string;
    privateIp: string;
    provisioningKeyId: string;
    provisioningKeyName: string;
    publicIp: string;
    sargeVersion: string;
    upgradeAttempt: string;
    upgradeStatus: string;
}

export interface GetZPAAppConnectorGroupServerGroup {
    configSpace: string;
    creationTime: string;
    description: string;
    dynamicDiscovery: boolean;
    /**
     * (String) Whether this App Connector Group is enabled or not. Default value: `true`. Supported values: `true`, `false`
     */
    enabled: boolean;
    /**
     * ID of the App Connector Group.
     */
    id: string;
    modifiedTime: string;
    modifiedby: string;
    /**
     * Name of the App Connector Group.
     */
    name: string;
}

export interface GetZPAApplicationSegmentBrowserAccessClientlessApp {
    /**
     * (bool)
     * * `cname` (string)
     * * `description` (string)
     * * `enabled` (bool)
     * * `hidden` (bool)
     * * `localDomain` (string)
     * * `path` (string)
     * * `trustUntrustedCert` (bool)
     */
    allowOptions: boolean;
    appId: string;
    /**
     * (string)
     */
    applicationPort: string;
    /**
     * (string)
     */
    applicationProtocol: string;
    /**
     * (string)
     */
    certificateId: string;
    /**
     * (string)
     */
    certificateName: string;
    cname: string;
    /**
     * (string) Description of the application.
     */
    description: string;
    /**
     * (string)
     */
    domain: string;
    /**
     * (Boolean) Whether this application is enabled or not. Default: false. Supported values: `true`, `false`.
     */
    enabled: boolean;
    hidden: boolean;
    /**
     * This field defines the id of the application server.
     */
    id: string;
    localDomain: string;
    /**
     * This field defines the name of the server.
     */
    name: string;
    path: string;
    trustUntrustedCert: boolean;
}

export interface GetZPAApplicationSegmentBrowserAccessServerGroup {
    /**
     * This field defines the id of the application server.
     */
    ids: string[];
}

export interface GetZPAApplicationSegmentInspectionInspectionApp {
    appId: string;
    /**
     * (string) TCP/UDP Port for ZPA Inspection.
     */
    applicationPort: string;
    /**
     * (string) Protocol for the Inspection Application. Supported values: `HTTP` and `HTTPS`
     */
    applicationProtocol: string;
    /**
     * (string) - ID of the signing certificate. This field is required if the applicationProtocol is set to `HTTPS`. The certificateId is not supported if the applicationProtocol is set to `HTTP`.
     */
    certificateId: string;
    /**
     * (string) - Parameter required when `applicationProtocol` is of type `HTTPS`
     */
    certificateName: string;
    /**
     * (string) Description of the application.
     */
    description: string;
    domain: string;
    /**
     * (bool) Whether this application is enabled or not
     */
    enabled: boolean;
    /**
     * The ID of the Inspection Application Segment to be exported.
     */
    id: string;
    /**
     * The name of the Inspection Application Segment to be exported.
     */
    name: string;
}

export interface GetZPAApplicationSegmentInspectionServerGroup {
    /**
     * The ID of the Inspection Application Segment to be exported.
     */
    ids: string[];
}

export interface GetZPAApplicationSegmentPRAServerGroup {
    ids: string[];
}

export interface GetZPAApplicationSegmentPRASraApp {
    appId: string;
    /**
     * (string) Port for the Privileged Remote Accessvalues: `RDP` and `SSH`
     */
    applicationPort: string;
    /**
     * (string) Protocol for the Privileged Remote Access. Supported values: `RDP` and `SSH`
     */
    applicationProtocol: string;
    certificateId: string;
    certificateName: string;
    /**
     * (string) - Parameter required when `applicationProtocol` is of type `RDP`
     */
    connectionSecurity: string;
    /**
     * (string) Description of the application.
     */
    description: string;
    domain: string;
    /**
     * (bool) Whether this application is enabled or not
     */
    enabled: boolean;
    hidden: boolean;
    id: string;
    /**
     * The name of the PRA Application Segment to be exported.
     */
    name: string;
    portal: boolean;
}

export interface GetZPAApplicationSegmentServerGroup {
    configSpace: string;
    creationTime: string;
    /**
     * Description of the application.
     */
    description: string;
    dynamicDiscovery: boolean;
    /**
     * Whether this application is enabled or not. Default: false. Supported values: `true`, `false`.
     */
    enabled: boolean;
    id: string;
    modifiedTime: string;
    modifiedby: string;
    /**
     * Name of the application.
     */
    name: string;
}

export interface GetZPACloudConnectorGroupCloudConnector {
    creationTime: string;
    description: string;
    enabled: boolean;
    fingerprint: string;
    /**
     * This field defines the id of the cloud connector group.
     */
    id: string;
    ipacls: any[];
    issuedCertId: string;
    modifiedTime: string;
    modifiedby: string;
    /**
     * This field defines the name of the cloud connector group.
     */
    name: string;
    signingCert: {[key: string]: any};
}

export interface GetZPAIdPControllerAdminMetadata {
    certificateUrl: string;
    spBaseUrl: string;
    spEntityId: string;
    spMetadataUrl: string;
    spPostUrl: string;
}

export interface GetZPAIdPControllerUserMetadata {
    certificateUrl: string;
    spBaseUrl: string;
    spEntityId: string;
    spMetadataUrl: string;
    spPostUrl: string;
}

export interface GetZPAInspectionAllPredefinedControlsList {
    /**
     * (string)
     * * `PASS`
     * * `BLOCK`
     * * `REDIRECT`
     */
    action: string;
    /**
     * (string)
     */
    actionValue: string;
    /**
     * (string)
     * * `id`- (string)
     * * `name`- (string)
     */
    associatedInspectionProfileNames: outputs.GetZPAInspectionAllPredefinedControlsListAssociatedInspectionProfileName[];
    /**
     * (string)
     */
    attachment: string;
    /**
     * (string)
     */
    controlGroup: string;
    /**
     * (string)
     */
    controlNumber: string;
    /**
     * (string)
     */
    creationTime: string;
    /**
     * (string)
     * * `PASS`
     * * `BLOCK`
     * * `REDIRECT`
     */
    defaultAction: string;
    /**
     * (string)
     */
    defaultActionValue: string;
    /**
     * (string)
     */
    description: string;
    /**
     * (string)
     */
    id: string;
    /**
     * (string)
     */
    modifiedTime: string;
    modifiedby: string;
    name: string;
    /**
     * (string)
     */
    paranoiaLevel: string;
    /**
     * (string)
     * * `CRITICAL`
     * * `ERROR`
     * * `WARNING`
     * * `INFO`
     */
    severity: string;
    /**
     * The version of the predefined control, the default is: `OWASP_CRS/3.3.0`
     */
    version: string;
}

export interface GetZPAInspectionAllPredefinedControlsListAssociatedInspectionProfileName {
    /**
     * (string)
     */
    id: string;
    name: string;
}

export interface GetZPAInspectionCustomControlsRule {
    conditions: outputs.GetZPAInspectionCustomControlsRuleCondition[];
    names: string[];
    type: string;
}

export interface GetZPAInspectionCustomControlsRuleCondition {
    lhs: string;
    op: string;
    rhs: string;
}

export interface GetZPAInspectionPredefinedControlsAssociatedInspectionProfileName {
    /**
     * (Computed)
     */
    id: string;
    /**
     * The name of the predefined control.
     */
    name: string;
}

export interface GetZPAInspectionProfileControlsInfo {
    /**
     * (string) Control types. Supported Values: `CUSTOM`, `PREDEFINED`, `ZSCALER`
     */
    controlType: string;
    /**
     * (string) Control information counts `Long`
     */
    count: string;
}

export interface GetZPAInspectionProfileCustomControl {
    /**
     * (string) The action of the predefined control. Supported values: `PASS`, `BLOCK` and `REDIRECT`
     */
    action: string;
    /**
     * (string) Value for the predefined controls action. This field is only required if the action is set to REDIRECT. This field is only required if the action is set to `REDIRECT`.
     * * `attachment` (string) Control attachment
     * * `controlGroup` (string) Control group
     */
    actionValue: string;
    /**
     * (string) Name of the inspection profile
     * * `id`- (string)
     * * `name`- (string)
     */
    associatedInspectionProfileNames: outputs.GetZPAInspectionProfileCustomControlAssociatedInspectionProfileName[];
    controlNumber: string;
    controlRuleJson: string;
    creationTime: string;
    defaultAction: string;
    defaultActionValue: string;
    /**
     * (string) Description of the inspection profile.
     */
    description: string;
    /**
     * This field defines the id of the application server.
     */
    id: string;
    modifiedBy: string;
    modifiedTime: string;
    /**
     * This field defines the name of the server.
     */
    name: string;
    /**
     * (string) OWASP Predefined Paranoia Level. Range: [1-4], inclusive
     */
    paranoiaLevel: string;
    /**
     * (string) Rules of the custom controls applied as conditions `JSON`
     */
    rules: outputs.GetZPAInspectionProfileCustomControlRule[];
    severity: string;
    /**
     * (string) Type value for the rules
     */
    type: string;
    /**
     * (string) The version of the predefined control, the default is: `OWASP_CRS/3.3.0`
     */
    version: string;
}

export interface GetZPAInspectionProfileCustomControlAssociatedInspectionProfileName {
    /**
     * This field defines the id of the application server.
     */
    id: string;
    /**
     * This field defines the name of the server.
     */
    name: string;
}

export interface GetZPAInspectionProfileCustomControlRule {
    /**
     * (string)
     */
    conditions: outputs.GetZPAInspectionProfileCustomControlRuleCondition[];
    /**
     * (string) Name of the rules. If rules.type is set to `REQUEST_HEADERS`, `REQUEST_COOKIES`, or `RESPONSE_HEADERS`, the rules.name field is required.
     */
    names: string;
    /**
     * (string) Type value for the rules
     */
    type: string;
}

export interface GetZPAInspectionProfileCustomControlRuleCondition {
    /**
     * (string) Signifies the key for the object type Supported values: `SIZE`, `VALUE`
     */
    lhs: string;
    /**
     * (string) If lhs is set to SIZE, then the user may pass one of the following: `EQ: Equals`, `LE: Less than or equal to`, `GE: Greater than or equal to`. If the lhs is set to `VALUE`, then the user may pass one of the following: `CONTAINS`, `STARTS_WITH`, `ENDS_WITH`, `RX`.
     */
    op: string;
    /**
     * (string) Denotes the value for the given object type. Its value depends on the key. If rules.type is set to REQUEST_METHOD, the conditions.rhs field must have one of the following values: `GET`,`HEAD`, `POST`, `OPTIONS`, `PUT`, `DELETE`, `TRACE`
     */
    rhs: string;
}

export interface GetZPAInspectionProfilePredefinedControl {
    /**
     * (string) The action of the predefined control. Supported values: `PASS`, `BLOCK` and `REDIRECT`
     */
    action: string;
    /**
     * (string) Value for the predefined controls action. This field is only required if the action is set to REDIRECT. This field is only required if the action is set to `REDIRECT`.
     * * `attachment` (string) Control attachment
     * * `controlGroup` (string) Control group
     */
    actionValue: string;
    /**
     * (string) Name of the inspection profile
     * * `id`- (string)
     * * `name`- (string)
     */
    associatedInspectionProfileNames: outputs.GetZPAInspectionProfilePredefinedControlAssociatedInspectionProfileName[];
    attachment: string;
    controlGroup: string;
    controlNumber: string;
    creationTime: string;
    defaultAction: string;
    defaultActionValue: string;
    /**
     * (string) Description of the inspection profile.
     */
    description: string;
    /**
     * This field defines the id of the application server.
     */
    id: string;
    modifiedBy: string;
    modifiedTime: string;
    /**
     * This field defines the name of the server.
     */
    name: string;
    /**
     * (string) OWASP Predefined Paranoia Level. Range: [1-4], inclusive
     */
    paranoiaLevel: string;
    severity: string;
    /**
     * (string) The version of the predefined control, the default is: `OWASP_CRS/3.3.0`
     */
    version: string;
}

export interface GetZPAInspectionProfilePredefinedControlAssociatedInspectionProfileName {
    /**
     * This field defines the id of the application server.
     */
    id: string;
    /**
     * This field defines the name of the server.
     */
    name: string;
}

export interface GetZPALSSConfigControllerConfig {
    /**
     * (string)
     */
    auditMessage: string;
    /**
     * (string)
     */
    description: string;
    /**
     * (bool)
     */
    enabled: boolean;
    /**
     * (string)
     */
    filters: string[];
    /**
     * (string)
     */
    format: string;
    /**
     * This field defines the name of the log streaming resource.
     */
    id: string;
    /**
     * (string)
     */
    lssHost: string;
    /**
     * (string)
     */
    lssPort: string;
    /**
     * This field defines the name of the log streaming resource.
     */
    name: string;
    /**
     * (string)
     */
    sourceLogType: string;
    useTls: boolean;
}

export interface GetZPALSSConfigControllerConnectorGroup {
    /**
     * This field defines the name of the log streaming resource.
     */
    id: string;
}

export interface GetZPALSSConfigControllerPolicyRule {
    action: string;
    actionId: string;
    bypassDefaultRule: boolean;
    conditions: outputs.GetZPALSSConfigControllerPolicyRuleCondition[];
    /**
     * (string)
     */
    creationTime: string;
    customMsg: string;
    defaultRule: boolean;
    /**
     * (string)
     */
    description: string;
    /**
     * This field defines the name of the log streaming resource.
     */
    id: string;
    isolationDefaultRule: boolean;
    lssDefaultRule: boolean;
    /**
     * (string)
     */
    modifiedTime: string;
    modifiedby: string;
    /**
     * This field defines the name of the log streaming resource.
     */
    name: string;
    operator: string;
    policySetId: string;
    policyType: string;
    priority: string;
    reauthDefaultRule: boolean;
    reauthIdleTimeout: string;
    reauthTimeout: string;
    ruleOrder: string;
    zpnCbiProfileId: string;
    zpnInspectionProfileId: string;
    zpnInspectionProfileName: string;
}

export interface GetZPALSSConfigControllerPolicyRuleCondition {
    /**
     * (string)
     */
    creationTime: string;
    /**
     * This field defines the name of the log streaming resource.
     */
    id: string;
    /**
     * (string)
     */
    modifiedTime: string;
    modifiedby: string;
    negated: boolean;
    operands: outputs.GetZPALSSConfigControllerPolicyRuleConditionOperand[];
    operator: string;
}

export interface GetZPALSSConfigControllerPolicyRuleConditionOperand {
    /**
     * (string)
     */
    creationTime: string;
    /**
     * This field defines the name of the log streaming resource.
     */
    id: string;
    idpId: string;
    lhs: string;
    /**
     * (string)
     */
    modifiedTime: string;
    modifiedby: string;
    /**
     * This field defines the name of the log streaming resource.
     */
    name: string;
    objectType: string;
    operator: string;
    rhs: string;
}

export interface GetZPAMachineGroupMachine {
    creationTime: string;
    description: string;
    fingerprint: string;
    /**
     * The ID of the machine group to be exported.
     */
    id: string;
    issuedCertId: string;
    machineGroupId: string;
    machineGroupName: string;
    machineTokenId: string;
    modifiedBy: string;
    modifiedTime: string;
    /**
     * The name of the machine group to be exported.
     */
    name: string;
    signingCert: {[key: string]: string};
}

export interface GetZPAPolicyTypeRule {
    action: string;
    actionId: string;
    bypassDefaultRule: boolean;
    conditions: outputs.GetZPAPolicyTypeRuleCondition[];
    creationTime: string;
    customMsg: string;
    description: string;
    id: string;
    isolationDefaultRule: boolean;
    modifiedBy: string;
    modifiedTime: string;
    name: string;
    operator: string;
    /**
     * The ID of the global policy set.
     */
    policySetId: string;
    /**
     * The value for differentiating the policy types.
     */
    policyType: string;
    priority: string;
    reauthDefaultRule: boolean;
    reauthIdleTimeout: string;
    reauthTimeout: string;
    ruleOrder: string;
    zpnCbiProfileId: string;
    zpnInspectionProfileId: string;
}

export interface GetZPAPolicyTypeRuleCondition {
    creationTime: string;
    id: string;
    modifiedBy: string;
    modifiedTime: string;
    negated: boolean;
    operands: outputs.GetZPAPolicyTypeRuleConditionOperand[];
    operator: string;
}

export interface GetZPAPolicyTypeRuleConditionOperand {
    creationTime: string;
    id: string;
    idpId: string;
    lhs: string;
    modifiedBy: string;
    modifiedTime: string;
    name: string;
    objectType: string;
    operator: string;
    rhs: string;
}

export interface GetZPASegmentGroupApplication {
    /**
     * (string)
     */
    bypassType: string;
    /**
     * (string)
     */
    configSpace: string;
    /**
     * (string)
     */
    creationTime: string;
    /**
     * (string)
     */
    defaultIdleTimeout: string;
    /**
     * (string)
     */
    defaultMaxAge: string;
    /**
     * (string)
     */
    description: string;
    /**
     * (string)
     */
    domainName: string;
    /**
     * (string)
     */
    domainNames: string[];
    /**
     * (string)
     */
    doubleEncrypt: boolean;
    /**
     * (bool)
     */
    enabled: boolean;
    /**
     * (string)
     */
    healthCheckType: string;
    /**
     * The ID of the segment group to be exported.
     */
    id: string;
    /**
     * (bool)
     */
    ipAnchored: boolean;
    logFeatures: string[];
    /**
     * (string)
     */
    modifiedBy: string;
    /**
     * (string)
     */
    modifiedTime: string;
    /**
     * The name of the segment group to be exported.
     */
    name: string;
    /**
     * (bool)
     */
    passiveHealthEnabled: boolean;
    /**
     * (Computed)
     */
    serverGroups: outputs.GetZPASegmentGroupApplicationServerGroup[];
    /**
     * (string)
     */
    tcpPortRanges: string[];
    /**
     * (string)
     */
    tcpPortsIns: string[];
    tcpPortsOuts: string[];
    /**
     * (string)
     */
    udpPortRanges: string[];
}

export interface GetZPASegmentGroupApplicationServerGroup {
    /**
     * (string)
     */
    configSpace: string;
    /**
     * (string)
     */
    creationTime: string;
    /**
     * (string)
     */
    description: string;
    /**
     * (bool)
     */
    dynamicDiscovery: boolean;
    /**
     * (bool)
     */
    enabled: boolean;
    /**
     * The ID of the segment group to be exported.
     */
    id: string;
    /**
     * (string)
     */
    modifiedBy: string;
    /**
     * (string)
     */
    modifiedTime: string;
    /**
     * The name of the segment group to be exported.
     */
    name: string;
}

export interface GetZPAServerGroupAppConnectorGroup {
    cityCountry: string;
    connectors: outputs.GetZPAServerGroupAppConnectorGroupConnector[];
    countryCode: string;
    creationTime: string;
    /**
     * (string) This field is the description of the server group.
     */
    description: string;
    dnsQueryType: string;
    /**
     * (bool) This field defines if the server group is enabled or disabled.
     */
    enabled: boolean;
    geolocationId: string;
    /**
     * The ID of the server group to be exported.
     */
    id: string;
    latitude: string;
    location: string;
    longitude: string;
    modifiedTime: string;
    modifiedby: string;
    /**
     * The name of the server group to be exported.
     */
    name: string;
    serverGroups: outputs.GetZPAServerGroupAppConnectorGroupServerGroup[];
    siemAppConnectorGroup: boolean;
    upgradeDay: string;
    upgradeTimeInSecs: string;
    versionProfileId: string;
}

export interface GetZPAServerGroupAppConnectorGroupConnector {
    creationTime: string;
    /**
     * (string) This field is the description of the server group.
     */
    description: string;
    /**
     * (bool) This field defines if the server group is enabled or disabled.
     */
    enabled: boolean;
    fingerprint: string;
    /**
     * The ID of the server group to be exported.
     */
    id: string;
    issuedCertId: string;
    modifiedTime: string;
    modifiedby: string;
    /**
     * The name of the server group to be exported.
     */
    name?: string;
    upgradeAttempt: string;
}

export interface GetZPAServerGroupAppConnectorGroupServerGroup {
    /**
     * (string)
     */
    configSpace: string;
    creationTime: string;
    /**
     * (string) This field is the description of the server group.
     */
    description: string;
    /**
     * (bool) This field controls dynamic discovery of the servers.
     */
    dynamicDiscovery: boolean;
    /**
     * (bool) This field defines if the server group is enabled or disabled.
     */
    enabled: boolean;
    /**
     * The ID of the server group to be exported.
     */
    id: string;
    modifiedTime: string;
    modifiedby: string;
    /**
     * The name of the server group to be exported.
     */
    name?: string;
}

export interface GetZPAServerGroupApplication {
    /**
     * The ID of the server group to be exported.
     */
    id: string;
    /**
     * The name of the server group to be exported.
     */
    name: string;
}

export interface GetZPAServerGroupServer {
    address: string;
    appServerGroupIds: string[];
    /**
     * (string)
     */
    configSpace: string;
    creationTime: string;
    /**
     * (string) This field is the description of the server group.
     */
    description: string;
    /**
     * (bool) This field defines if the server group is enabled or disabled.
     */
    enabled: boolean;
    /**
     * The ID of the server group to be exported.
     */
    id: string;
    modifiedTime: string;
    modifiedby: string;
    /**
     * The name of the server group to be exported.
     */
    name: string;
}

export interface GetZPAServiceEdgeGroupServiceEdge {
    applicationStartTime: string;
    controlChannelStatus: string;
    /**
     * (string)
     */
    creationTime: string;
    ctrlBrokerName: string;
    currentVersion: string;
    description: string;
    /**
     * (bool) Whether this App Connector Group is enabled or not. Default value: `true`. Supported values: `true`, `false`
     */
    enabled: boolean;
    enrollmentCert: {[key: string]: any};
    expectedUpgradeTime: string;
    expectedVersion: string;
    fingerprint: string;
    /**
     * The ID of the service edge group to be exported.
     */
    id: string;
    ipacl: string;
    issuedCertId: string;
    lastBrokerConnectTime: string;
    lastBrokerConnectTimeDuration: string;
    lastBrokerDisconnectTime: string;
    lastBrokerDisconnectTimeDuration: string;
    lastUpgradeTime: string;
    /**
     * (string) Latitude of the Service Edge Group. Integer or decimal. With values in the range of `-90` to `90`
     */
    latitude: string;
    listenIps: string;
    /**
     * (string) Location of the Service Edge Group.
     */
    location: string;
    /**
     * (string) Longitude of the Service Edge Group.Integer or decimal. With values in the range of `-180` to `180`
     */
    longitude: string;
    /**
     * (string)
     */
    modifiedTime: string;
    modifiedby: string;
    /**
     * The name of the service edge group to be exported.
     */
    name: string;
    platform: string;
    previousVersion: string;
    privateIp: string;
    provisioningKeyId: string;
    provisioningKeyName: string;
    publicIp: string;
    publishIps: string;
    sargeVersion: string;
    serviceEdgeGroupId: string;
    serviceEdgeGroupName: string;
    upgradeAttempt: string;
    upgradeStatus: string;
}

export interface GetZPAServiceEdgeGroupTrustedNetwork {
    /**
     * (string)
     */
    creationTime: string;
    domain: string;
    /**
     * The ID of the service edge group to be exported.
     */
    id: string;
    masterCustomerId: string;
    /**
     * (string)
     */
    modifiedTime: string;
    modifiedby: string;
    /**
     * The name of the service edge group to be exported.
     */
    name: string;
    networkId: string;
    zscalerCloud: string;
}

export interface ZPAApplicationSegmentBrowserAccessClientlessApp {
    /**
     * - If you want ZPA to forward unauthenticated HTTP preflight OPTIONS requests from the browser to the app.. Supported values: `true` and `false`
     */
    allowOptions?: boolean;
    /**
     * - Port for the BA app.
     */
    applicationPort: string;
    /**
     * - Protocol for the BA app. Supported values: `HTTP` and `HTTPS`
     */
    applicationProtocol: string;
    certificateId: string;
    cname: string;
    description?: string;
    /**
     * - Domain name or IP address of the BA app.
     */
    domain?: string;
    enabled: boolean;
    hidden: boolean;
    id: string;
    localDomain?: string;
    /**
     * - Name of BA app.
     */
    name: string;
    path?: string;
    trustUntrustedCert: boolean;
}

export interface ZPAApplicationSegmentBrowserAccessServerGroup {
    ids: string[];
}

export interface ZPAApplicationSegmentInspectionCommonAppsDto {
    appsConfigs: outputs.ZPAApplicationSegmentInspectionCommonAppsDtoAppsConfig[];
}

export interface ZPAApplicationSegmentInspectionCommonAppsDtoAppsConfig {
    allowOptions: boolean;
    appId: string;
    /**
     * Indicates the type of application as inspection. Supported value: `INSPECT`
     */
    appTypes: string[];
    /**
     * Port for the Inspection Application Segment.
     */
    applicationPort: string;
    /**
     * Protocol for the Inspection Application Segment.. Supported values: `HTTP` and `HTTPS`
     */
    applicationProtocol: string;
    /**
     * - ID of the signing certificate. This field is required if the applicationProtocol is set to `HTTPS`. The certificateId is not supported if the applicationProtocol is set to `HTTP`.
     */
    certificateId: string;
    certificateName: string;
    cname: string;
    /**
     * (Optional) Description of the application.
     */
    description: string;
    /**
     * Domain name of the Inspection Application Segment.
     */
    domain: string;
    /**
     * Whether this application is enabled or not
     */
    enabled: boolean;
    hidden: boolean;
    id: string;
    localDomain: string;
    /**
     * Name of the Inspection Application Segment.
     */
    name: string;
    portal: boolean;
    trustUntrustedCert: boolean;
}

export interface ZPAApplicationSegmentInspectionServerGroup {
    ids: string[];
}

export interface ZPAApplicationSegmentPRACommonAppsDto {
    appsConfigs: outputs.ZPAApplicationSegmentPRACommonAppsDtoAppsConfig[];
}

export interface ZPAApplicationSegmentPRACommonAppsDtoAppsConfig {
    allowOptions: boolean;
    appId: string;
    /**
     * Indicates the type of application as Privileged Remote Access. Supported value: `SECURE_REMOTE_ACCESS`
     */
    appTypes: string[];
    /**
     * Port for the Privileged Remote Access
     */
    applicationPort: string;
    /**
     * Protocol for the Privileged Remote Access. Supported values: `RDP` and `SSH`
     */
    applicationProtocol: string;
    cname: string;
    /**
     * - Parameter required when `applicationProtocol` is of type `RDP`
     */
    connectionSecurity: string;
    /**
     * (Optional) Description of the application.
     */
    description: string;
    /**
     * Domain name of the Privileged Remote Access
     */
    domain: string;
    /**
     * Whether this application is enabled or not
     */
    enabled: boolean;
    hidden: boolean;
    id: string;
    localDomain: string;
    /**
     * Name of the Privileged Remote Access
     */
    name: string;
    portal: boolean;
}

export interface ZPAApplicationSegmentPRAServerGroup {
    ids: string[];
}

export interface ZPAApplicationSegmentServerGroup {
    ids: string[];
}

export interface ZPABrowserAccessClientlessApp {
    allowOptions?: boolean;
    applicationPort: string;
    applicationProtocol: string;
    certificateId: string;
    cname: string;
    description?: string;
    domain?: string;
    enabled: boolean;
    hidden: boolean;
    id: string;
    localDomain?: string;
    name: string;
    path?: string;
    trustUntrustedCert: boolean;
}

export interface ZPABrowserAccessServerGroup {
    ids: string[];
}

export interface ZPAInspectionCustomControlsAssociatedInspectionProfileName {
    ids: string[];
}

export interface ZPAInspectionCustomControlsRule {
    conditions: outputs.ZPAInspectionCustomControlsRuleConditions;
    names: string[];
    type: string;
}

export interface ZPAInspectionCustomControlsRuleConditions {
    lhs: string;
    op: string;
    rhs: string;
}

export interface ZPAInspectionProfileControlsInfo {
    /**
     * (Optional) Control types. Supported Values: `CUSTOM`, `PREDEFINED`, `ZSCALER`
     */
    controlType: string;
    /**
     * (Optional) Control information counts `Long`
     */
    count: string;
}

export interface ZPAInspectionProfileCustomControl {
    /**
     * The action of the predefined control. Supported values: `PASS`, `BLOCK` and `REDIRECT`
     */
    action: string;
    /**
     * Value for the predefined controls action. This field is only required if the action is set to REDIRECT. This field is only required if the action is set to `REDIRECT`.
     */
    actionValue?: string;
    /**
     * ID of the predefined control
     */
    id: string;
}

export interface ZPAInspectionProfilePredefinedControl {
    /**
     * The action of the predefined control. Supported values: `PASS`, `BLOCK` and `REDIRECT`
     */
    action: string;
    /**
     * Value for the predefined controls action. This field is only required if the action is set to REDIRECT. This field is only required if the action is set to `REDIRECT`.
     */
    actionValue?: string;
    /**
     * ID of the predefined control
     */
    id: string;
}

export interface ZPALSSConfigControllerConfig {
    /**
     * (Optional)
     */
    auditMessage: string;
    /**
     * (Optional)
     */
    description?: string;
    /**
     * (Optional)
     */
    enabled?: boolean;
    /**
     * (Optional)
     */
    filters?: string[];
    /**
     * The format of the LSS resource. The supported formats are: `JSON`, `CSV`, and `TSV`
     */
    format: string;
    /**
     * - App Connector Group ID(s) where logs will be forwarded to.
     */
    id: string;
    /**
     * The IP or FQDN of the SIEM (Log Receiver) where logs will be forwarded to.
     */
    lssHost: string;
    /**
     * The destination port of the SIEM (Log Receiver) where logs will be forwarded to.
     */
    lssPort: string;
    /**
     * (Optional)
     */
    name: string;
    /**
     * Refer to the log type documentation
     */
    sourceLogType: string;
    /**
     * (Optional)
     */
    useTls?: boolean;
}

export interface ZPALSSConfigControllerConnectorGroup {
    /**
     * - App Connector Group ID(s) where logs will be forwarded to.
     */
    ids?: string[];
}

export interface ZPALSSConfigControllerPolicyRuleResource {
    /**
     * (Optional)
     */
    action?: string;
    actionId?: string;
    bypassDefaultRule?: boolean;
    /**
     * (Optional)
     */
    conditions?: outputs.ZPALSSConfigControllerPolicyRuleResourceCondition[];
    /**
     * (Optional)
     */
    customMsg?: string;
    defaultRule?: boolean;
    /**
     * (Optional)
     */
    description?: string;
    /**
     * - App Connector Group ID(s) where logs will be forwarded to.
     */
    id: string;
    lssDefaultRule?: boolean;
    /**
     * (Optional)
     */
    name: string;
    operator: string;
    policySetId?: string;
    policyType: string;
    priority: string;
    reauthDefaultRule?: boolean;
    reauthIdleTimeout?: string;
    reauthTimeout?: string;
    ruleOrder: string;
    zpnInspectionProfileId?: string;
}

export interface ZPALSSConfigControllerPolicyRuleResourceCondition {
    /**
     * (Optional)
     * * `operator` (Optional) - Supported values are: `AND` or `OR`
     * * `operands`
     * * `objectType` (Optional) This is for specifying the policy critiera. Supported values: `APP`, `APP_GROUP`, `CLIENT_TYPE`, `TRUSTED_NETWORK`, `SAML`, `SCIM`, `SCIM_GROUP`
     * * `values` (Optional) The below values are supported when choosing `objectType` of type `CLIENT_TYPE`.
     * - `zpnClientTypeExporter`
     * - `zpnClientTypeBrowserIsolation`
     * - `zpnClientTypeMachineTunnel`
     * - `zpnClientTypeIpAnchoring`
     * - `zpnClientTypeEdgeConnector`
     * - `zpnClientTypeZapp`
     */
    negated?: boolean;
    operands?: outputs.ZPALSSConfigControllerPolicyRuleResourceConditionOperand[];
    operator: string;
}

export interface ZPALSSConfigControllerPolicyRuleResourceConditionOperand {
    objectType: string;
    values?: string[];
}

export interface ZPAPolicyAccessForwardingRuleCondition {
    id: string;
    negated: boolean;
    operands: outputs.ZPAPolicyAccessForwardingRuleConditionOperand[];
    operator: string;
}

export interface ZPAPolicyAccessForwardingRuleConditionOperand {
    id: string;
    idpId: string;
    lhs: string;
    name: string;
    objectType: string;
    rhs?: string;
    rhsLists?: string[];
}

export interface ZPAPolicyAccessInspectionRuleCondition {
    id: string;
    negated: boolean;
    operands: outputs.ZPAPolicyAccessInspectionRuleConditionOperand[];
    operator: string;
}

export interface ZPAPolicyAccessInspectionRuleConditionOperand {
    id: string;
    idpId: string;
    lhs: string;
    name: string;
    objectType: string;
    rhs?: string;
    rhsLists?: string[];
}

export interface ZPAPolicyAccessRuleAppConnectorGroup {
    /**
     * (Optional) The ID of a server group resource
     */
    ids?: string[];
}

export interface ZPAPolicyAccessRuleAppServerGroup {
    /**
     * (Optional) The ID of a server group resource
     */
    ids?: string[];
}

export interface ZPAPolicyAccessRuleCondition {
    /**
     * (Optional) The ID of a server group resource
     */
    id: string;
    /**
     * (Optional) Supported values: ``true`` or ``false``
     * * `operator` (Optional) Supported values: ``AND``, and ``OR``
     * * `operands` (Optional) - Operands block must be repeated if multiple per `objectType` conditions are to be added to the rule.
     * * `name` (Optional)
     * * `lhs` (Optional) LHS must always carry the string value ``id`` or the attribute ID of the resource being associated with the rule.
     * * `rhs` (Optional) RHS is either the ID attribute of a resource or fixed string value. Refer to the chart below for further details.
     * * `idpId` (Optional)
     * * `objectType` (Optional) This is for specifying the policy critiera. Supported values: `APP`, `APP_GROUP`, `SAML`, `IDP`, `CLIENT_TYPE`, `TRUSTED_NETWORK`, `POSTURE`, `SCIM`, `SCIM_GROUP`, and `CLOUD_CONNECTOR_GROUP`. `TRUSTED_NETWORK`, and `CLIENT_TYPE`.
     * * `CLIENT_TYPE` (Optional) - The below options are the only ones supported in an access policy rule.
     * * `zpnClientTypeExporter`
     * * `zpnClientTypeBrowserIsolation`
     * * `zpnClientTypeMachineTunnel`
     * * `zpnClientTypeIpAnchoring`
     * * `zpnClientTypeEdgeConnector`
     * * `zpnClientTypeZapp`
     */
    negated: boolean;
    operands: outputs.ZPAPolicyAccessRuleConditionOperand[];
    operator: string;
}

export interface ZPAPolicyAccessRuleConditionOperand {
    /**
     * (Optional) The ID of a server group resource
     */
    id: string;
    idpId: string;
    lhs: string;
    name: string;
    objectType: string;
    rhs?: string;
    rhsLists?: string[];
}

export interface ZPAPolicyAccessTimeOutRuleCondition {
    id: string;
    negated: boolean;
    operands: outputs.ZPAPolicyAccessTimeOutRuleConditionOperand[];
    operator: string;
}

export interface ZPAPolicyAccessTimeOutRuleConditionOperand {
    id: string;
    idpId: string;
    lhs: string;
    name: string;
    objectType: string;
    rhs?: string;
    rhsLists?: string[];
}

export interface ZPASegmentGroupApplication {
    id?: string;
}

export interface ZPAServerGroupAppConnectorGroup {
    ids?: string[];
}

export interface ZPAServerGroupApplication {
    ids?: string[];
}

export interface ZPAServerGroupServer {
    ids?: string[];
}

export interface ZPAServiceEdgeGroupServiceEdge {
    ids?: string[];
}

export interface ZPAServiceEdgeGroupTrustedNetwork {
    ids?: string[];
}

