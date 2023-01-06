package main

import (
	"fmt"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	zpa "github.com/zscaler/pulumi-zpa/sdk/go/zpa"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		segmentGroup, err := zpa.NewZPASegmentGroup(ctx, "segment-group-example", &zpa.ZPASegmentGroupArgs{
			Name:                pulumi.String("Pulumi Segment Group Golang"),
			Description:         pulumi.String("Pulumi Segment Group Golang"),
			Enabled:             pulumi.Bool(true),
			TcpKeepAliveEnabled: pulumi.String("1"),
		})
		if err != nil {
			return fmt.Errorf("error creating segment group: %v", err)
		}

		ctx.Export("segment_group", segmentGroup.Name)

		return nil
	})
}
