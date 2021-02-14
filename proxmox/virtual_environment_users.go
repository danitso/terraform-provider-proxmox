/* This Source Code Form is subject to the terms of the Mozilla Public
 * License, v. 2.0. If a copy of the MPL was not distributed with this
 * file, You can obtain one at https://mozilla.org/MPL/2.0/. */

package proxmox

import (
	"errors"
	"fmt"
	"net/url"
	"sort"
	"time"
)

// ChangeUserPassword changes a user's password.
func (c *VirtualEnvironmentClient) ChangeUserPassword(id, password string) error {
	d := VirtualEnvironmentUserChangePasswordRequestBody{
		ID:       id,
		Password: password,
	}

	return c.DoRequest(hmPUT, "access/password", d, nil)
}

// CreateUser creates an user.
func (c *VirtualEnvironmentClient) CreateUser(d *VirtualEnvironmentUserCreateRequestBody) error {
	return c.DoRequest(hmPOST, "access/users", d, nil)
}

// CreateUserToken creates a user token.
func (c *VirtualEnvironmentClient) CreateUserToken(id, tokenID string, d *VirtualEnvironmentUserTokenCreateRequestBody) (*VirtualEnvironmentUserTokenCreateResponseData, error) {
	resBody := &VirtualEnvironmentUserTokenCreateResponseBody{}
	err := c.DoRequest(hmPOST, fmt.Sprintf("access/users/%s/token/%s", url.PathEscape(id), url.PathEscape(tokenID)), d, resBody)

	if err != nil {
		return nil, err
	}

	if resBody.Data == nil {
		return nil, errors.New("The server did not include a data object in the response")
	}

	return resBody.Data, nil
}

// DeleteUser deletes an user.
func (c *VirtualEnvironmentClient) DeleteUser(id string) error {
	return c.DoRequest(hmDELETE, fmt.Sprintf("access/users/%s", url.PathEscape(id)), nil, nil)
}

// DeleteUserToken deletes a user token.
func (c *VirtualEnvironmentClient) DeleteUserToken(id, tokenID string) error {
	return c.DoRequest(hmDELETE, fmt.Sprintf("access/users/%s/token/%s", url.PathEscape(id), url.PathEscape(tokenID)), nil, nil)
}

// GetUser retrieves an user.
func (c *VirtualEnvironmentClient) GetUser(id string) (*VirtualEnvironmentUserGetResponseData, error) {
	resBody := &VirtualEnvironmentUserGetResponseBody{}
	err := c.DoRequest(hmGET, fmt.Sprintf("access/users/%s", url.PathEscape(id)), nil, resBody)

	if err != nil {
		return nil, err
	}

	if resBody.Data == nil {
		return nil, errors.New("The server did not include a data object in the response")
	}

	if resBody.Data.ExpirationDate != nil {
		expirationDate := CustomTimestamp(time.Time(*resBody.Data.ExpirationDate).UTC())
		resBody.Data.ExpirationDate = &expirationDate
	}

	if resBody.Data.Groups != nil {
		sort.Strings(*resBody.Data.Groups)
	}

	if resBody.Data.Tokens != nil {
		for _, tv := range *resBody.Data.Tokens {
			if tv.ExpirationDate != nil {
				expirationDate := CustomTimestamp(time.Time(*tv.ExpirationDate).UTC())
				tv.ExpirationDate = &expirationDate
			}
		}
	}

	return resBody.Data, nil
}

// GetUserToken retrieves a user token.
func (c *VirtualEnvironmentClient) GetUserToken(id, tokenID string) (*VirtualEnvironmentUserTokenGetResponseData, error) {
	resBody := &VirtualEnvironmentUserTokenGetResponseBody{}
	err := c.DoRequest(hmGET, fmt.Sprintf("access/users/%s/token/%s", url.PathEscape(id), url.PathEscape(tokenID)), nil, resBody)

	if err != nil {
		return nil, err
	}

	if resBody.Data == nil {
		return nil, errors.New("The server did not include a data object in the response")
	}

	return resBody.Data, nil
}

// ListUsers retrieves a list of users.
func (c *VirtualEnvironmentClient) ListUsers(enabled, full bool) ([]*VirtualEnvironmentUserListResponseData, error) {
	cbEnabled := CustomBool(enabled)
	cbFull := CustomBool(full)

	d := VirtualEnvironmentUserListRequestBody{
		Enabled: &cbEnabled,
		Full:    &cbFull,
	}

	resBody := &VirtualEnvironmentUserListResponseBody{}
	err := c.DoRequest(hmGET, "access/users", d, resBody)

	if err != nil {
		return nil, err
	}

	if resBody.Data == nil {
		return nil, errors.New("The server did not include a data object in the response")
	}

	sort.Slice(resBody.Data, func(i, j int) bool {
		return resBody.Data[i].ID < resBody.Data[j].ID
	})

	for _, v := range resBody.Data {
		if v.ExpirationDate != nil {
			expirationDate := CustomTimestamp(time.Time(*v.ExpirationDate).UTC())
			v.ExpirationDate = &expirationDate
		}

		if v.Groups != nil {
			sort.Strings(*v.Groups)
		}

		if v.Tokens != nil {
			for _, tv := range *v.Tokens {
				if tv.ExpirationDate != nil {
					expirationDate := CustomTimestamp(time.Time(*tv.ExpirationDate).UTC())
					tv.ExpirationDate = &expirationDate
				}
			}
		}
	}

	return resBody.Data, nil
}

// ListUserTokens retrieves information about API tokens for a user.
func (c *VirtualEnvironmentClient) ListUserTokens(id string) ([]*VirtualEnvironmentUserTokenListResponseData, error) {
	resBody := &VirtualEnvironmentUserTokenListResponseBody{}
	err := c.DoRequest(hmGET, fmt.Sprintf("access/users/%s/token", url.PathEscape(id)), nil, resBody)

	if err != nil {
		return nil, err
	}

	if resBody.Data == nil {
		return nil, errors.New("The server did not include a data object in the response")
	}

	for _, tv := range resBody.Data {
		if tv.ExpirationDate != nil {
			expirationDate := CustomTimestamp(time.Time(*tv.ExpirationDate).UTC())
			tv.ExpirationDate = &expirationDate
		}
	}

	return resBody.Data, nil
}

// UpdateUser updates an user.
func (c *VirtualEnvironmentClient) UpdateUser(id string, d *VirtualEnvironmentUserUpdateRequestBody) error {
	return c.DoRequest(hmPUT, fmt.Sprintf("access/users/%s", url.PathEscape(id)), d, nil)
}

// UpdateUserToken updates a user token.
func (c *VirtualEnvironmentClient) UpdateUserToken(id, tokenID string, d *VirtualEnvironmentUserTokenUpdateRequestBody) (*VirtualEnvironmentUserTokenUpdateResponseData, error) {
	resBody := &VirtualEnvironmentUserTokenUpdateResponseBody{}
	err := c.DoRequest(hmPUT, fmt.Sprintf("access/users/%s/token/%s", url.PathEscape(id), url.PathEscape(tokenID)), d, resBody)

	if err != nil {
		return nil, err
	}

	if resBody.Data == nil {
		return nil, errors.New("The server did not include a data object in the response")
	}

	return resBody.Data, nil
}
