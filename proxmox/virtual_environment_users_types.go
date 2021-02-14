/* This Source Code Form is subject to the terms of the Mozilla Public
 * License, v. 2.0. If a copy of the MPL was not distributed with this
 * file, You can obtain one at https://mozilla.org/MPL/2.0/. */

package proxmox

// VirtualEnvironmentUserChangePasswordRequestBody contains the data for a user password change request.
type VirtualEnvironmentUserChangePasswordRequestBody struct {
	ID       string `json:"userid" url:"userid"`
	Password string `json:"password" url:"password"`
}

// VirtualEnvironmentUserCreateRequestBody contains the data for an user create request.
type VirtualEnvironmentUserCreateRequestBody struct {
	Comment        *string          `json:"comment,omitempty" url:"comment,omitempty"`
	Email          *string          `json:"email,omitempty" url:"email,omitempty"`
	Enabled        *CustomBool      `json:"enable,omitempty" url:"enable,omitempty,int"`
	ExpirationDate *CustomTimestamp `json:"expire,omitempty" url:"expire,omitempty,unix"`
	FirstName      *string          `json:"firstname,omitempty" url:"firstname,omitempty"`
	Groups         []string         `json:"groups,omitempty" url:"groups,omitempty,comma"`
	ID             string           `json:"userid" url:"userid"`
	Keys           *string          `json:"keys,omitempty" url:"keys,omitempty"`
	LastName       *string          `json:"lastname,omitempty" url:"lastname,omitempty"`
	Password       string           `json:"password" url:"password"`
}

// VirtualEnvironmentUserGetResponseBody contains the body from an user get response.
type VirtualEnvironmentUserGetResponseBody struct {
	Data *VirtualEnvironmentUserGetResponseData `json:"data,omitempty"`
}

// VirtualEnvironmentUserGetResponseData contains the data from an user get response.
type VirtualEnvironmentUserGetResponseData struct {
	Comment        *string                                                `json:"comment,omitempty"`
	Email          *string                                                `json:"email,omitempty"`
	Enabled        *CustomBool                                            `json:"enable,omitempty"`
	ExpirationDate *CustomTimestamp                                       `json:"expire,omitempty"`
	FirstName      *string                                                `json:"firstname,omitempty"`
	Groups         *[]string                                              `json:"groups,omitempty"`
	Keys           *string                                                `json:"keys,omitempty"`
	LastName       *string                                                `json:"lastname,omitempty"`
	Tokens         *map[string]VirtualEnvironmentUserTokenGetResponseData `json:"tokens,omitempty"`
}

// VirtualEnvironmentUserListRequestBody contains the data for a user list request.
type VirtualEnvironmentUserListRequestBody struct {
	Enabled *CustomBool `json:"enabled,omitempty" url:"enabled,omitempty,int"`
	Full    *CustomBool `json:"full,omitempty" url:"full,omitempty,int"`
}

// VirtualEnvironmentUserListResponseBody contains the body from a user list response.
type VirtualEnvironmentUserListResponseBody struct {
	Data []*VirtualEnvironmentUserListResponseData `json:"data,omitempty"`
}

// VirtualEnvironmentUserListResponseData contains the data from a user list response.
type VirtualEnvironmentUserListResponseData struct {
	Comment        *string                                       `json:"comment,omitempty"`
	Email          *string                                       `json:"email,omitempty"`
	Enabled        *CustomBool                                   `json:"enable,omitempty"`
	ExpirationDate *CustomTimestamp                              `json:"expire,omitempty"`
	FirstName      *string                                       `json:"firstname,omitempty"`
	Groups         *CustomCommaSeparatedList                     `json:"groups,omitempty"`
	ID             string                                        `json:"userid"`
	Keys           *string                                       `json:"keys,omitempty"`
	LastName       *string                                       `json:"lastname,omitempty"`
	Tokens         *[]VirtualEnvironmentUserTokenGetResponseData `json:"tokens,omitempty"`
}

// VirtualEnvironmentUserTokenGetResponseBody contains the body from a user token get response.
type VirtualEnvironmentUserTokenGetResponseBody struct {
	Data *VirtualEnvironmentUserTokenGetResponseData `json:"data,omitempty"`
}

// VirtualEnvironmentUserTokenGetResponseData contains the data from a user token get response.
type VirtualEnvironmentUserTokenGetResponseData struct {
	Comment             *string          `json:"comment,omitempty"`
	ExpirationDate      *CustomTimestamp `json:"expire,omitempty"`
	ID                  *string          `json:"tokenid,omitempty"`
	PrivilegeSeperation *CustomBool      `json:"privsep,omitempty"`
}

// VirtualEnvironmentUserUpdateRequestBody contains the data for an user update request.
type VirtualEnvironmentUserUpdateRequestBody struct {
	Append         *CustomBool      `json:"append,omitempty" url:"append,omitempty"`
	Comment        *string          `json:"comment,omitempty" url:"comment,omitempty"`
	Email          *string          `json:"email,omitempty" url:"email,omitempty"`
	Enabled        *CustomBool      `json:"enable,omitempty" url:"enable,omitempty,int"`
	ExpirationDate *CustomTimestamp `json:"expire,omitempty" url:"expire,omitempty,unix"`
	FirstName      *string          `json:"firstname,omitempty" url:"firstname,omitempty"`
	Groups         []string         `json:"groups,omitempty" url:"groups,omitempty,comma"`
	Keys           *string          `json:"keys,omitempty" url:"keys,omitempty"`
	LastName       *string          `json:"lastname,omitempty" url:"lastname,omitempty"`
}
