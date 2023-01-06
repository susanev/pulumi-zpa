package main

import (
	"fmt"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	zpa "github.com/zscaler/pulumi-zpa/sdk/go/zpa"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		serviceEdgeGroup, err := zpa.NewZPAServiceEdgeGroup(ctx, "service-edge-group-example", &zpa.ZPAServiceEdgeGroupArgs{
			Name:                   pulumi.String("Pulumi Service Edge Group"),
			Description:            pulumi.String("Pulumi Service Edge Group"),
			Enabled:                pulumi.Bool(true),
			IsPublic:               pulumi.String("TRUE"),
			CountryCode:            pulumi.String("US"),
			CityCountry:            pulumi.String("San Jose, US"),
			Location:               pulumi.String("San Jose, CA, USA"),
			Latitude:               pulumi.String("37.3382082"),
			Longitude:              pulumi.String("-121.8863286"),
			UpgradeDay:             pulumi.String("SUNDAY"),
			UpgradeTimeInSecs:      pulumi.String("66600"),
			OverrideVersionProfile: pulumi.Bool(true),
			VersionProfileId:       pulumi.String("2"),
		})
		if err != nil {
			return fmt.Errorf("error creating service edge group: %v", err)
		}

		ctx.Export("group", serviceEdgeGroup.Name)

		return nil
	})
}
