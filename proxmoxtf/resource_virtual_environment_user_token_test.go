/* This Source Code Form is subject to the terms of the Mozilla Public
 * License, v. 2.0. If a copy of the MPL was not distributed with this
 * file, You can obtain one at https://mozilla.org/MPL/2.0/. */

package proxmoxtf

import (
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
)

// TestResourceVirtualEnvironmentUserTokenInstantiation tests whether the ResourceVirtualEnvironmentUserToken instance can be instantiated.
func TestResourceVirtualEnvironmentUserTokenInstantiation(t *testing.T) {
	s := resourceVirtualEnvironmentUserToken()

	if s == nil {
		t.Fatalf("Cannot instantiate resourceVirtualEnvironmentUserToken")
	}
}

// TestResourceVirtualEnvironmentUserTokenSchema tests the resourceVirtualEnvironmentUserToken schema.
func TestResourceVirtualEnvironmentUserTokenSchema(t *testing.T) {
	s := resourceVirtualEnvironmentUserToken()

	testRequiredArguments(t, s, []string{
		mkResourceVirtualEnvironmentUserTokenName,
		mkResourceVirtualEnvironmentUserTokenUserID,
	})

	testOptionalArguments(t, s, []string{
		mkResourceVirtualEnvironmentUserTokenACL,
		mkResourceVirtualEnvironmentUserTokenComment,
		mkResourceVirtualEnvironmentUserTokenExpirationDate,
		mkResourceVirtualEnvironmentUserTokenPrivilegeSeparation,
	})

	testComputedAttributes(t, s, []string{
		mkResourceVirtualEnvironmentUserTokenSecret,
	})

	testValueTypes(t, s, map[string]schema.ValueType{
		mkResourceVirtualEnvironmentUserTokenACL:                 schema.TypeSet,
		mkResourceVirtualEnvironmentUserTokenComment:             schema.TypeString,
		mkResourceVirtualEnvironmentUserTokenExpirationDate:      schema.TypeString,
		mkResourceVirtualEnvironmentUserTokenName:                schema.TypeString,
		mkResourceVirtualEnvironmentUserTokenPrivilegeSeparation: schema.TypeBool,
		mkResourceVirtualEnvironmentUserTokenSecret:              schema.TypeString,
		mkResourceVirtualEnvironmentUserTokenUserID:              schema.TypeString,
	})

	aclSchema := testNestedSchemaExistence(t, s, mkResourceVirtualEnvironmentUserTokenACL)

	testRequiredArguments(t, aclSchema, []string{
		mkResourceVirtualEnvironmentUserTokenACLPath,
		mkResourceVirtualEnvironmentUserTokenACLRoleID,
	})

	testOptionalArguments(t, aclSchema, []string{
		mkResourceVirtualEnvironmentUserTokenACLPropagate,
	})

	testValueTypes(t, aclSchema, map[string]schema.ValueType{
		mkResourceVirtualEnvironmentUserTokenACLPath:      schema.TypeString,
		mkResourceVirtualEnvironmentUserTokenACLPropagate: schema.TypeBool,
		mkResourceVirtualEnvironmentUserTokenACLRoleID:    schema.TypeString,
	})
}
