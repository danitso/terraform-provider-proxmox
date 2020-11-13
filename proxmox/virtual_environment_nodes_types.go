/* This Source Code Form is subject to the terms of the Mozilla Public
 * License, v. 2.0. If a copy of the MPL was not distributed with this
 * file, You can obtain one at https://mozilla.org/MPL/2.0/. */

package proxmox

import (
	"encoding/json"
	"net/url"
)

// CustomNodeCommands contains an array of commands to execute.
type CustomNodeCommands []string

// VirtualEnvironmentNodeExecuteRequestBody contains the data for a node execute request.
type VirtualEnvironmentNodeExecuteRequestBody struct {
	Commands CustomNodeCommands `json:"commands" url:"commands"`
}

// VirtualEnvironmentNodeGetTimeResponseBody contains the body from a node time zone get response.
type VirtualEnvironmentNodeGetTimeResponseBody struct {
	Data *VirtualEnvironmentNodeGetTimeResponseData `json:"data,omitempty"`
}

// VirtualEnvironmentNodeGetTimeResponseData contains the data from a node list response.
type VirtualEnvironmentNodeGetTimeResponseData struct {
	LocalTime CustomTimestamp `json:"localtime"`
	TimeZone  string          `json:"timezone"`
	UTCTime   CustomTimestamp `json:"time"`
}

// VirtualEnvironmentNodeListResponseBody contains the body from a node list response.
type VirtualEnvironmentNodeListResponseBody struct {
	Data []*VirtualEnvironmentNodeListResponseData `json:"data,omitempty"`
}

// VirtualEnvironmentNodeListResponseData contains the data from a node list response.
type VirtualEnvironmentNodeListResponseData struct {
	CPUCount        *int     `json:"maxcpu,omitempty"`
	CPUUtilization  *float64 `json:"cpu,omitempty"`
	MemoryAvailable *int     `json:"maxmem,omitempty"`
	MemoryUsed      *int     `json:"mem,omitempty"`
	Name            string   `json:"node"`
	SSLFingerprint  *string  `json:"ssl_fingerprint,omitempty"`
	Status          *string  `json:"status"`
	SupportLevel    *string  `json:"level,omitempty"`
	Uptime          *int     `json:"uptime"`
}

// VirtualEnvironmentNodeUpdateTimeRequestBody contains the body for a node time update request.
type VirtualEnvironmentNodeUpdateTimeRequestBody struct {
	TimeZone string `json:"timezone" url:"timezone"`
}

// EncodeValues converts a CustomNodeCommands array to a JSON encoded URL vlaue.
func (r CustomNodeCommands) EncodeValues(key string, v *url.Values) error {
	jsonArrayBytes, err := json.Marshal(r)

	if err != nil {
		return err
	}

	v.Add(key, string(jsonArrayBytes))

	return nil
}
