import * as zpa from "@zscaler/pulumi-zpa";

// Create a App Connector Group
const appConnectorGroup = new zpa.ZPAAppConnectorGroup("app-connector-group-example", {
    name: "Pulumi App Connector Group",
    description: "Pulumi App Connector Group",
    enabled: true,
    cityCountry: "San Jose, CA",
    countryCode: "US",
    latitude: "37.3382082",
    longitude: "-121.8863",
    location: "San Jose, CA, US",
    overrideVersionProfile: true,
    upgradeDay: "SUNDAY",
    upgradeTimeInSecs: "66600",
    versionProfileId: "2",
    dnsQueryType: "IPV4_IPV6",
 })

 export const groupId = appConnectorGroup.id

const certId = zpa.getZPAEnrollmentCert({
     name: "Connector",
 });

 export const connector = certId.then((connector => connector.id))

// Create a Provisioning Key
const provisioningKey = new zpa.ZPAProvisioningKey("provisioning-key-example", {
    name: "Pulumi Provisioning Key",
    associationType: "CONNECTOR_GRP",
    enrollmentCertId: "6573",
    zcomponentId: appConnectorGroup.id,
    maxUsage: "10"
 })

 export const keyId = provisioningKey.id