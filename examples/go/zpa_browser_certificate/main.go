package main

import (
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"github.com/zscaler/pulumi-zpa/sdk/go/zpa"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := zpa.GetZPABaCertificate(ctx, &zpa.GetZPABaCertificateArgs{
			Name: pulumi.StringRef("jenkins.securitygeek.io"),
		}, nil)
		if err != nil {
			return err
		}
		return nil
	})
}
