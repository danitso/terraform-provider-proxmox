package acceptancetests

import (
	"fmt"
	"github.com/danitso/terraform-provider-proxmox/proxmoxtf"
	"github.com/danitso/terraform-provider-proxmox/proxmoxtf/acceptancetests/testutils"
	"github.com/hashicorp/terraform/helper/resource"
	"github.com/hashicorp/terraform/terraform"
	"strings"
	"testing"
)


// Verifies that a role can be created and updated
func TestAccResourceVirtualEnvironmentRole_CreateAndUpdate(t *testing.T) {
	roleId := testutils.GenerateResourceName()
	privilegesFirst := []string{
		"Datastore.Allocate",
		"Datastore.AllocateSpace",
	}

	privilegesSecond := []string{
		"Datastore.Audit",
	}

	tfNode := "proxmox_virtual_environment_role.role"

	resource.Test(t, resource.TestCase{
		PreCheck: func() { testutils.PreCheck(t, nil) },
		Providers: testutils.GetProviders(),
		CheckDestroy: CheckRoleDestroyed,
		Steps: []resource.TestStep{
			// Create role
			{
				Config: testutils.HclRoleResource(roleId, privilegesFirst),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(tfNode, "role_id", roleId),
					testutils.CheckRole(privilegesFirst),
				),
			},
			// Update role
			{
				Config: testutils.HclRoleResource(roleId, privilegesSecond),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(tfNode, "role_id", roleId),
					testutils.CheckRole(privilegesSecond),
				),
			},
		},
	})
}

// CheckRoleDestroyed verifies that all roles referenced in the state
// are destroyed. This will be invoked *after* terraform destroys
// the resource but *before* the state is wiped clean
func CheckRoleDestroyed(s *terraform.State) error {
	config := testutils.GetProvider().Meta().(proxmoxtf.ProviderConfiguration)

	conn, err := config.GetVEClient()

	if err != nil {
		return err
	}

	// loop through the resource state
	for _, rs := range s.RootModule().Resources {
		if rs.Type != "proxmox_virtual_environment_role" {
			continue
		}

		role, err := conn.GetRole(rs.Primary.ID)

		if err == nil {
			// if we still receiving role with privileges
			if len(*role) != 0 {
				return fmt.Errorf("Role with Name=`%s` should not exist", rs.Primary.ID)
			}

			return nil
		}

		if !strings.Contains(err.Error(), "does not exist") {
			return err
		}
	}

	return nil
}
