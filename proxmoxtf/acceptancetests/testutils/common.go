package testutils

import (
	"github.com/danitso/terraform-provider-proxmox/proxmoxtf"
	"github.com/hashicorp/terraform/helper/acctest"
	"github.com/hashicorp/terraform/helper/schema"
	"github.com/hashicorp/terraform/terraform"
	"os"
	"testing"
)

// initialize once, so it can be shared by each acceptance test
var provider = proxmoxtf.Provider()

// GetProvider returns the proxmox provider
func GetProvider() *schema.Provider  {
	return provider
}

// GetProviders returns a map of all providers needed for the project
func GetProviders() map[string]terraform.ResourceProvider {
	return map[string]terraform.ResourceProvider{
		"proxmox": GetProvider(),
	}
}

// PreCheck checks that required environmental variables are set
func PreCheck(t *testing.T, extraEnvVars *[]string) {
	requiredEnvVars := []string{
		"PROXMOX_VE_USERNAME",
		"PROXMOX_VE_PASSWORD",
		"PROXMOX_VE_ENDPOINT",
	}

	if extraEnvVars != nil {
		requiredEnvVars = append(requiredEnvVars, *extraEnvVars...)
	}

	for _, variable := range requiredEnvVars {
		if _, ok := os.LookupEnv(variable); !ok {
			t.Fatalf("`%s` must be set for acceptance testing", variable)
		}
	}
}

// GenerateResourceName generates a random name with a constant prefix
func GenerateResourceName() string {
	return "test-" + acctest.RandStringFromCharSet(10, acctest.CharSetAlphaNum)
}

