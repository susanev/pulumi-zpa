package main

import (
	"fmt"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	zpa "github.com/zscaler/pulumi-zpa/sdk/go/zpa"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		applicationServer, err := zpa.NewZPAApplicationServer(ctx, "application-server-example", &zpa.ZPAApplicationServerArgs{
			Name:        pulumi.String("Pulumi Application Server"),
			Description: pulumi.String("Pulumi Application Server"),
			Enabled:     pulumi.Bool(true),
			Address:     pulumi.String("192.168.1.1"),
		})
		if err != nil {
			return fmt.Errorf("error creating application server: %v", err)
		}

		ctx.Export("server", applicationServer.Name)

		return nil
	})
}
