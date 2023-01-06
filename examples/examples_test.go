package examples

import (
	"os"
	"testing"

	"github.com/pulumi/pulumi/pkg/v3/testing/integration"
)

func skipIfNoZPACreds(t *testing.T) {
	zpa_client_id := os.Getenv("ZPA_CLIENT_ID")
	if zpa_client_id == "" {
		t.Skipf("Skipping test due to missing ZPA_CLIENT_ID variable")
	}
	zpa_client_secret := os.Getenv("ZPA_CLIENT_SECRET")
	if zpa_client_secret == "" {
		t.Skipf("Skipping test due to missing ZPA_CLIENT_SECRET variable")
	}
	zpa_customer_id := os.Getenv("ZPA_CUSTOMER_ID")
	if zpa_customer_id == "" {
		t.Skipf("Skipping test due to missing ZPA_CUSTOMER_ID variable")
	}
}

func getCwd(t *testing.T) string {
	cwd, err := os.Getwd()
	if err != nil {
		t.FailNow()
	}

	return cwd
}

func getBaseOptions() integration.ProgramTestOptions {
	return integration.ProgramTestOptions{
		RunUpdateTest:        false,
		ExpectRefreshChanges: true,
	}
}
