package nagios

import (
	"os"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/terraform"
)

var testAccProviders map[string]terraform.ResourceProvider
var testAccProvider *schema.Provider
var requiredEnvVariables = []string{
	"API_TOKEN",
	"NAGIOS_URL",
}

func TestNagiosProvider(t *testing.T) {
	if err := testAccProvider.InternalValidate(); err != nil {
		t.Fatalf("err: %s", err)
	}
}

func testAccPreCheck(t *testing.T) {
	for _, variable := range requiredEnvVariables {
		if value := os.Getenv(variable); value == "" {
			t.Fatalf("%s must be set before running acceptance tests.", variable)
		}
	}
}

func init() {
	testAccProvider = NagiosProvider()
	testAccProviders = map[string]terraform.ResourceProvider{
		"nagios": testAccProvider,
	}
}
