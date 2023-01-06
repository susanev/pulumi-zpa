package main

import (
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"github.com/zscaler/pulumi-zpa/sdk/go/zpa"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		exampleAppConnectorGroup, err := zpa.NewZPAAppConnectorGroup(ctx, "app-connector-group-example", &zpa.ZPAAppConnectorGroupArgs{
			Name:                   pulumi.String("Pulumi App Connector Group"),
			Description:            pulumi.String("Pulumi App Connector Group"),
			Enabled:                pulumi.Bool(true),
			CountryCode:            pulumi.String("US"),
			Latitude:               pulumi.String("37.338"),
			Longitude:              pulumi.String("-121.8863"),
			Location:               pulumi.String("San Jose, CA, US"),
			UpgradeDay:             pulumi.String("SUNDAY"),
			UpgradeTimeInSecs:      pulumi.String("66600"),
			OverrideVersionProfile: pulumi.Bool(true),
			VersionProfileId:       pulumi.String("0"),
			DnsQueryType:           pulumi.String("IPV4_IPV6"),
		})
		if err != nil {
			return err
		}
		_, err = zpa.NewZPAServerGroup(ctx, "server-group-example", &zpa.ZPAServerGroupArgs{
			Name:             pulumi.String("Pulumi Server Group"),
			Description:      pulumi.String("Pulumi Server Group"),
			Enabled:          pulumi.Bool(true),
			DynamicDiscovery: pulumi.Bool(true),
			AppConnectorGroups: zpa.ZPAServerGroupAppConnectorGroupArray{
				&zpa.ZPAServerGroupAppConnectorGroupArgs{
					Ids: pulumi.StringArray{
						exampleAppConnectorGroup.ID(),
					},
				},
			},
		}, pulumi.DependsOn([]pulumi.Resource{
			exampleAppConnectorGroup,
		}))
		if err != nil {
			return err
		}
		return nil
	})
}
