package testutils

import (
	"fmt"
	"github.com/danitso/terraform-provider-proxmox/proxmox"
	"github.com/danitso/terraform-provider-proxmox/proxmoxtf"
	"github.com/hashicorp/terraform/helper/resource"
	"github.com/hashicorp/terraform/terraform"
	"sort"
)

// CheckRole Given the name of role, this will return a function that will check
// whether or not a role
// - (1) exists in the state
// - (2) exist in Proxmox VE
// - (3) has correct privileges
func CheckRole(expectedPrivileges []string) resource.TestCheckFunc  {
	return func (s *terraform.State) error {
		res, ok := s.RootModule().Resources["proxmox_virtual_environment_role.role"]

		if !ok {
			return fmt.Errorf("Did not find the role in the TF state")
		}

		clients := GetProvider().Meta().(proxmoxtf.ProviderConfiguration)
		id := res.Primary.ID
		role, err := readRole(clients, id)

		if err != nil {
			return fmt.Errorf("Role with Name=`%s` cannot be found. Error %v", id, err)
		}

		if len(*role) != len(expectedPrivileges) {
			return fmt.Errorf("Role with Name=`%s` should have `%d` privileges, but got `%d`", id, len(*role), len(expectedPrivileges))
		}

		sort.Strings(expectedPrivileges)

		for i, v := range *role {
			if v != expectedPrivileges[i] {
				return fmt.Errorf("Role with Name=`%s` has privilege=`%s`, but expected `%s`", id, v, expectedPrivileges[i])
			}
		}

		return nil

	}
}

// readRole is a helper function that reads a role based on a given name
func readRole(clients proxmoxtf.ProviderConfiguration, identifier string) (*proxmox.CustomPrivileges, error) {
	conn, err := clients.GetVEClient()

	if err != nil {
		return nil, err
	}

	response, err := conn.GetRole(identifier)

	if err != nil {
		return nil, err
	}

	return response, nil
}

// HclRoleResource HCL describing of a PVE role resource
func HclRoleResource(name string, privileges []string) string  {

	if name == "" {
		panic("Parameter: `name` cannot be empty")
	}

	if len(privileges) <= 0 {
		panic("Parameter: `privileges` cannot be empty")
	}

	p := ""

	for _, v := range privileges {
		p = p + fmt.Sprintf("%q,", v)
	}

	return fmt.Sprintf(`
resource "proxmox_virtual_environment_role" "role" {
  role_id = "%[1]s"

  privileges = [
	%[2]s
  ]
}
`, name, p)
}