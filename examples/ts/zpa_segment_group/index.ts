import * as zpa from "@zscaler/pulumi-zpa"


// Create a Segment Group
const segmentGroup = new zpa.ZPASegmentGroup("segment-group-example", {
    name: "Pulumi Segment Group TypeScript",
    description: "Pulumi Segment Group TypeScript",
    enabled: true,
    tcpKeepAliveEnabled: "1"
 })

 export const group = segmentGroup.id