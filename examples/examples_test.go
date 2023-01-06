// Copyright 2016-2017, Pulumi Corporation.  All rights reserved.

package examples

import (
	"os"
	"testing"

	"github.com/pulumi/pulumi/pkg/v3/testing/integration"
)

func getClientId(t *testing.T) string {
	zpa_client_id := os.Getenv("ZPA_CLIENT_ID")
	if zpa_client_id == "" {
		t.Skipf("Skipping test due to missing ZPA_CLIENT_ID environment variable")
	}

	return zpa_client_id
}

func getClientSecret(t *testing.T) string {
	zpa_client_secret := os.Getenv("ZPA_CLIENT_SECRET")
	if zpa_client_secret == "" {
		t.Skipf("Skipping test due to missing ZPA_CLIENT_SECRET environment variable")
	}

	return zpa_client_secret
}

func getCustomerId(t *testing.T) string {
	zpa_customer_id := os.Getenv("ZPA_CUSTOMER_ID")
	if zpa_customer_id == "" {
		t.Skipf("Skipping test due to missing ZPA_CUSTOMER_ID environment variable")
	}

	return zpa_customer_id
}

func getCwd(t *testing.T) string {
	cwd, err := os.Getwd()
	if err != nil {
		t.FailNow()
	}

	return cwd
}

func getBaseOptions(t *testing.T) integration.ProgramTestOptions {
	zpa_client_id := getClientId(t)
	zpa_client_secret := getClientSecret(t)
	zpa_customer_id := getCustomerId(t)
	return integration.ProgramTestOptions{
		Config: map[string]string{
			"zpa_client_id":     zpa_client_id,
			"zpa_client_secret": zpa_client_secret,
			"zpa_customer_id":   zpa_customer_id,
		},
	}
}
