/* This Source Code Form is subject to the terms of the Mozilla Public
 * License, v. 2.0. If a copy of the MPL was not distributed with this
 * file, You can obtain one at https://mozilla.org/MPL/2.0/. */

package proxmoxtf

import (
	"testing"

	"github.com/hashicorp/terraform/helper/schema"
)

// TestResourceVirtualEnvironmentCertificateInstantiation tests whether the ResourceVirtualEnvironmentCertificate instance can be instantiated.
func TestResourceVirtualEnvironmentClusterOptionsInstantiation(t *testing.T) {
	s := resourceVirtualEnvironmentClusterOptions()

	if s == nil {
		t.Fatalf("Cannot instantiate resourceVirtualEnvironmentClusterOptions")
	}
}

// TestResourceVirtualEnvironmentCertificateSchema tests the resourceVirtualEnvironmentCertificate schema.
func TestResourceVirtualEnvironmentClusterOptionsSchema(t *testing.T) {
	s := resourceVirtualEnvironmentClusterOptions()

	testOptionalArguments(t, s, []string{
		mkResourceVirtualEnvironmentClusterOptionsBandwidthLimit,
		mkResourceVirtualEnvironmentClusterOptionsConsole,
		mkResourceVirtualEnvironmentClusterOptionsEmailFrom,
		mkResourceVirtualEnvironmentClusterOptionsFencing,
		mkResourceVirtualEnvironmentClusterOptionsHttpProxy,
		mkResourceVirtualEnvironmentClusterOptionsKeyboard,
		mkResourceVirtualEnvironmentClusterOptionsLanguage,
		mkResourceVirtualEnvironmentClusterOptionsMacPrefix,
		mkResourceVirtualEnvironmentClusterOptionsMaxWorkers,
	})

	testValueTypes(t, s, map[string]schema.ValueType{
		mkResourceVirtualEnvironmentClusterOptionsBandwidthLimit: schema.TypeList,
		mkResourceVirtualEnvironmentClusterOptionsConsole:        schema.TypeString,
		mkResourceVirtualEnvironmentClusterOptionsEmailFrom:      schema.TypeString,
		mkResourceVirtualEnvironmentClusterOptionsFencing:        schema.TypeString,
		mkResourceVirtualEnvironmentClusterOptionsHttpProxy:      schema.TypeString,
		mkResourceVirtualEnvironmentClusterOptionsKeyboard:       schema.TypeString,
		mkResourceVirtualEnvironmentClusterOptionsLanguage:       schema.TypeString,
		mkResourceVirtualEnvironmentClusterOptionsMacPrefix:      schema.TypeString,
		mkResourceVirtualEnvironmentClusterOptionsMaxWorkers:     schema.TypeInt,
	})
}
