package main

import (
	"fmt"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	zpa "github.com/zscaler/pulumi-zpa/sdk/go/zpa"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		appConnectorGroup, err := zpa.NewZPAAppConnectorGroup(ctx, "app-connector-group-example", &zpa.ZPAAppConnectorGroupArgs{
			Name:                   pulumi.String("Pulumi App Connector Group"),
			Description:            pulumi.String("Pulumi App Connector Group"),
			Enabled:                pulumi.Bool(true),
			CountryCode:            pulumi.String("US"),
			Latitude:               pulumi.String("37.3382082"),
			Longitude:              pulumi.String("-121.8863286"),
			Location:               pulumi.String("San Jose, CA, USA"),
			UpgradeDay:             pulumi.String("SUNDAY"),
			UpgradeTimeInSecs:      pulumi.String("66600"),
			OverrideVersionProfile: pulumi.Bool(true),
			VersionProfileId:       pulumi.String("2"),
			DnsQueryType:           pulumi.String("IPV4_IPV6"),
		})
		if err != nil {
			return fmt.Errorf("error creating app connector group: %v", err)
		}

		ctx.Export("group", appConnectorGroup.Name)

		return nil
	})
}
