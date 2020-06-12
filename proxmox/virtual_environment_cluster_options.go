/* This Source Code Form is subject to the terms of the Mozilla Public
 * License, v. 2.0. If a copy of the MPL was not distributed with this
 * file, You can obtain one at https://mozilla.org/MPL/2.0/. */

package proxmox

import (
	"errors"
)

// DeleteCertificate deletes the custom certificate for a node.
func (c *VirtualEnvironmentClient) DeleteClusterOptions(d *VirtualEnvironmentClusterOptionsDeleteRequestBody) error {
	return c.DoRequest(hmPUT, "cluster/options", d, nil)
}

// UpdateCertificate updates the custom certificate for a node.
func (c *VirtualEnvironmentClient) GetClusterOptions() (*VirtualEnvironmentClusterOptionsGetResponseData, error) {
	resBody := &VirtualEnvironmentClusterOptionsGetResponseBody{}
	err := c.DoRequest(hmGET, "cluster/options", nil, resBody)

	if err != nil {
		return nil, err
	}

	if resBody.Data == nil {
		return nil, errors.New("the server did not include a data object in the response")
	}

	return resBody.Data, nil
}

// UpdateCertificate updates the custom certificate for a node.
func (c *VirtualEnvironmentClient) UpdateClusterOptions(d *VirtualEnvironmentClusterOptionsUpdateRequestBody) error {
	return c.DoRequest(hmPUT, "cluster/options", d, nil)
}
