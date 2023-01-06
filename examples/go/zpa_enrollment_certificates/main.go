package main

import (
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"github.com/zscaler/pulumi-zpa/sdk/go/zpa"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := zpa.GetZPAEnrollmentCert(ctx, &zpa.GetZPAEnrollmentCertArgs{
			Name: pulumi.StringRef("Root"),
		}, nil)
		if err != nil {
			return err
		}
		_, err = zpa.GetZPAEnrollmentCert(ctx, &zpa.GetZPAEnrollmentCertArgs{
			Name: pulumi.StringRef("Client"),
		}, nil)
		if err != nil {
			return err
		}
		_, err = zpa.GetZPAEnrollmentCert(ctx, &zpa.GetZPAEnrollmentCertArgs{
			Name: pulumi.StringRef("Connector"),
		}, nil)
		if err != nil {
			return err
		}
		_, err = zpa.GetZPAEnrollmentCert(ctx, &zpa.GetZPAEnrollmentCertArgs{
			Name: pulumi.StringRef("Service Edge"),
		}, nil)
		if err != nil {
			return err
		}
		_, err = zpa.GetZPAEnrollmentCert(ctx, &zpa.GetZPAEnrollmentCertArgs{
			Name: pulumi.StringRef("Isolation Client"),
		}, nil)
		if err != nil {
			return err
		}
		return nil
	})
}
