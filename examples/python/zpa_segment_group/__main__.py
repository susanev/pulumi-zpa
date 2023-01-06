"""A Python Pulumi program"""

import zscaler_zpa as zpa

segmentGroup = zpa.ZPASegmentGroup("segment-group-example",
    name = "Pulumi Segment Group",
    description = "Pulumi Segment Group",
    enabled = True,
    policy_migrated = True,
    tcp_keep_alive_enabled = "1"
)