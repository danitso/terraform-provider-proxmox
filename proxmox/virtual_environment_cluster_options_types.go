/* This Source Code Form is subject to the terms of the Mozilla Public
 * License, v. 2.0. If a copy of the MPL was not distributed with this
 * file, You can obtain one at https://mozilla.org/MPL/2.0/. */

package proxmox

import (
	"strconv"
	"strings"
)

// VirtualEnvironmentCertificateDeleteRequestBody contains the data for a custom certificate delete request.
type VirtualEnvironmentClusterOptionsDeleteRequestBody struct {
	Delete DeleteOptions `json:"delete" url:"delete"`
}

type DeleteOptions []string

func (dr *DeleteOptions) UnmarshalJSON(b []byte) error {
	unquoted, err := strconv.Unquote(string(b))
	if err != nil {
		return nil
	}

	list := strings.Split(unquoted, ",")
	*dr = list

	return nil
}

func (dr DeleteOptions) MarshalJSON() ([]byte, error) {
	return []byte(strconv.Quote(strings.Join(dr, ","))), nil
}

// VirtualEnvironmentCertificateListResponseBody contains the body from a certificate list response.
type VirtualEnvironmentClusterOptionsGetResponseBody struct {
	Data *VirtualEnvironmentClusterOptionsGetResponseData `json:"data,omitempty"`
}

// VirtualEnvironmentCertificateListResponseData contains the data from a certificate list response.
type VirtualEnvironmentClusterOptionsGetResponseData struct {
	BandwidthLimit string `json:"bwlimit"`
	Console        string `json:"console"`
	EmailFrom      string `json:"email_from"`
	Fencing        string `json:"fencing"`
	HttpProxy      string `json:"http_proxy"`
	Keyboard       string `json:"keyboard"`
	Language       string `json:"language"`
	MacPrefix      string `json:"mac_prefix"`
	MaxWorkers     int    `json:"max_workers"`

	// todo ha and possible migration, migration_unsecure, u2f
}

// VirtualEnvironmentCertificateUpdateRequestBody contains the body for a custom certificate update request.
type VirtualEnvironmentClusterOptionsUpdateRequestBody struct {
	BandwidthLimit *string `json:"bwlimit" url:"bwlimit,omitempty"`
	Console        *string `json:"console" url:"console,omitempty"`
	EmailFrom      *string `json:"email_from" url:"email_from,omitempty"`
	Fencing        *string `json:"fencing" url:"fencing,omitempty"`
	HttpProxy      *string `json:"http_proxy" url:"http_proxy,omitempty"`
	Keyboard       *string `json:"keyboard" url:"keyboard,omitempty"`
	Language       *string `json:"language" url:"language,omitempty"`
	MacPrefix      *string `json:"mac_prefix" url:"mac_prefix,omitempty"`
	MaxWorkers     *int    `json:"max_workers" url:"max_workers,omitempty"`
}
