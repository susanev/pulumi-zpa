import * as zpa from "@zscaler/pulumi-zpa";

// Create a Segment Group
const segmentGroup = new zpa.ZPASegmentGroup("segment-group-example", {
    name: "Pulumi Segment Group",
    description: "Pulumi Segment Group",
    enabled: true,
    policyMigrated: true,
    tcpKeepAliveEnabled: "1"
 })

 export const group = segmentGroup.id