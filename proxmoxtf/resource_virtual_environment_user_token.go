/* This Source Code Form is subject to the terms of the Mozilla Public
 * License, v. 2.0. If a copy of the MPL was not distributed with this
 * file, You can obtain one at https://mozilla.org/MPL/2.0/. */

package proxmoxtf

import (
	"errors"
	"strings"
	"time"

	"github.com/danitso/terraform-provider-proxmox/proxmox"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/helper/validation"
)

const (
	dvResourceVirtualEnvironmentUserTokenComment             = ""
	dvResourceVirtualEnvironmentUserTokenPrivilegeSeparation = true

	mkResourceVirtualEnvironmentUserTokenACL                 = "acl"
	mkResourceVirtualEnvironmentUserTokenACLPath             = "path"
	mkResourceVirtualEnvironmentUserTokenACLPropagate        = "propagate"
	mkResourceVirtualEnvironmentUserTokenACLRoleID           = "role_id"
	mkResourceVirtualEnvironmentUserTokenComment             = "comment"
	mkResourceVirtualEnvironmentUserTokenExpirationDate      = "expiration_date"
	mkResourceVirtualEnvironmentUserTokenName                = "name"
	mkResourceVirtualEnvironmentUserTokenPrivilegeSeparation = "privilege_separation"
	mkResourceVirtualEnvironmentUserTokenSecret              = "secret"
	mkResourceVirtualEnvironmentUserTokenUserID              = "user_id"
)

var (
	dvResourceVirtualEnvironmentUserTokenExpirationDate = time.Unix(0, 0).UTC().Format(time.RFC3339)
)

func resourceVirtualEnvironmentUserToken() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			mkResourceVirtualEnvironmentUserTokenACL: {
				Type:        schema.TypeSet,
				Description: "The access control list",
				Optional:    true,
				DefaultFunc: func() (interface{}, error) {
					return []interface{}{}, nil
				},
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						mkResourceVirtualEnvironmentUserTokenACLPath: {
							Type:        schema.TypeString,
							Required:    true,
							Description: "The path",
						},
						mkResourceVirtualEnvironmentUserTokenACLPropagate: {
							Type:        schema.TypeBool,
							Optional:    true,
							Description: "Whether to propagate to child paths",
							Default:     false,
						},
						mkResourceVirtualEnvironmentUserTokenACLRoleID: {
							Type:        schema.TypeString,
							Required:    true,
							Description: "The role id",
						},
					},
				},
			},
			mkResourceVirtualEnvironmentUserTokenComment: {
				Type:        schema.TypeString,
				Description: "The user comment",
				Optional:    true,
				Default:     dvResourceVirtualEnvironmentUserTokenComment,
			},
			mkResourceVirtualEnvironmentUserTokenExpirationDate: {
				Type:         schema.TypeString,
				Description:  "The user account's expiration date",
				Optional:     true,
				Default:      dvResourceVirtualEnvironmentUserTokenExpirationDate,
				ValidateFunc: validation.ValidateRFC3339TimeString,
			},
			mkResourceVirtualEnvironmentUserTokenName: {
				Type:        schema.TypeString,
				Description: "The token name",
				Required:    true,
				ForceNew:    true,
			},
			mkResourceVirtualEnvironmentUserTokenPrivilegeSeparation: {
				Type:        schema.TypeBool,
				Optional:    true,
				Description: "Whether the privileges for the token differs from the account privileges",
				Default:     dvResourceVirtualEnvironmentUserTokenPrivilegeSeparation,
			},
			mkResourceVirtualEnvironmentUserTokenSecret: {
				Type:        schema.TypeString,
				Description: "The token secret",
				Computed:    true,
			},
			mkResourceVirtualEnvironmentUserTokenUserID: {
				Type:        schema.TypeString,
				Description: "The user id",
				Required:    true,
				ForceNew:    true,
			},
		},
		Create: resourceVirtualEnvironmentUserTokenCreate,
		Read:   resourceVirtualEnvironmentUserTokenRead,
		Update: resourceVirtualEnvironmentUserTokenUpdate,
		Delete: resourceVirtualEnvironmentUserTokenDelete,
	}
}

func resourceVirtualEnvironmentUserTokenCreate(d *schema.ResourceData, m interface{}) error {
	config := m.(providerConfiguration)
	veClient, err := config.GetVEClient()

	if err != nil {
		return err
	}

	comment := d.Get(mkResourceVirtualEnvironmentUserTokenComment).(string)
	expirationDate, err := time.Parse(time.RFC3339, d.Get(mkResourceVirtualEnvironmentUserTokenExpirationDate).(string))

	if err != nil {
		return err
	}

	expirationDateCustom := proxmox.CustomTimestamp(expirationDate)
	name := d.Get(mkResourceVirtualEnvironmentUserTokenName).(string)
	privilegeSeparation := proxmox.CustomBool(d.Get(mkResourceVirtualEnvironmentUserTokenPrivilegeSeparation).(bool))
	userID := d.Get(mkResourceVirtualEnvironmentUserTokenUserID).(string)

	body := &proxmox.VirtualEnvironmentUserTokenCreateRequestBody{
		Comment:             &comment,
		ExpirationDate:      &expirationDateCustom,
		PrivilegeSeperation: &privilegeSeparation,
	}

	res, err := veClient.CreateUserToken(userID, name, body)

	if err != nil {
		return err
	}

	if res.FullTokenID == nil || *res.FullTokenID == "" {
		return errors.New("The server did not include the full token identifier in the response")
	} else if res.Value == nil || *res.Value == "" {
		return errors.New("The server did not include the token secret in the response")
	}

	d.SetId(*res.FullTokenID)

	d.Set(mkResourceVirtualEnvironmentUserTokenSecret, res.Value)

	aclParsed := d.Get(mkResourceVirtualEnvironmentUserTokenACL).(*schema.Set).List()

	for _, v := range aclParsed {
		aclDelete := proxmox.CustomBool(false)
		aclEntry := v.(map[string]interface{})
		aclPropagate := proxmox.CustomBool(aclEntry[mkResourceVirtualEnvironmentUserTokenACLPropagate].(bool))

		aclBody := &proxmox.VirtualEnvironmentACLUpdateRequestBody{
			Delete:    &aclDelete,
			Path:      aclEntry[mkResourceVirtualEnvironmentUserTokenACLPath].(string),
			Propagate: &aclPropagate,
			Roles:     []string{aclEntry[mkResourceVirtualEnvironmentUserTokenACLRoleID].(string)},
			Tokens:    []string{d.Id()},
		}

		err := veClient.UpdateACL(aclBody)

		if err != nil {
			return err
		}
	}

	return resourceVirtualEnvironmentUserTokenRead(d, m)
}

func resourceVirtualEnvironmentUserTokenRead(d *schema.ResourceData, m interface{}) error {
	config := m.(providerConfiguration)
	veClient, err := config.GetVEClient()

	if err != nil {
		return err
	}

	tokenID := d.Id()
	tokenParts := strings.Split(tokenID, "!")
	token, err := veClient.GetUserToken(tokenParts[0], tokenParts[1])

	if err != nil {
		if strings.Contains(err.Error(), "HTTP 404") || (strings.Contains(err.Error(), "HTTP 500") && strings.Contains(err.Error(), "no such token")) {
			d.SetId("")

			return nil
		}

		return err
	}

	acl, err := veClient.GetACL()

	if err != nil {
		return err
	}

	aclParsed := []interface{}{}

	for _, v := range acl {
		if v.Type == "token" && v.UserOrGroupID == tokenID {
			aclEntry := map[string]interface{}{}

			aclEntry[mkResourceVirtualEnvironmentUserTokenACLPath] = v.Path

			if v.Propagate != nil {
				aclEntry[mkResourceVirtualEnvironmentUserTokenACLPropagate] = bool(*v.Propagate)
			} else {
				aclEntry[mkResourceVirtualEnvironmentUserTokenACLPropagate] = false
			}

			aclEntry[mkResourceVirtualEnvironmentUserTokenACLRoleID] = v.RoleID

			aclParsed = append(aclParsed, aclEntry)
		}
	}

	d.Set(mkResourceVirtualEnvironmentUserTokenACL, aclParsed)

	if token.Comment != nil {
		d.Set(mkResourceVirtualEnvironmentUserTokenComment, token.Comment)
	} else {
		d.Set(mkResourceVirtualEnvironmentUserTokenComment, "")
	}

	if token.ExpirationDate != nil {
		d.Set(mkResourceVirtualEnvironmentUserTokenExpirationDate, time.Time(*token.ExpirationDate).Format(time.RFC3339))
	} else {
		d.Set(mkResourceVirtualEnvironmentUserTokenExpirationDate, time.Unix(0, 0).UTC().Format(time.RFC3339))
	}

	if token.PrivilegeSeperation != nil {
		d.Set(mkResourceVirtualEnvironmentUserTokenPrivilegeSeparation, bool(*token.PrivilegeSeperation))
	} else {
		d.Set(mkResourceVirtualEnvironmentUserTokenPrivilegeSeparation, "")
	}

	return nil
}

func resourceVirtualEnvironmentUserTokenUpdate(d *schema.ResourceData, m interface{}) error {
	config := m.(providerConfiguration)
	veClient, err := config.GetVEClient()

	if err != nil {
		return err
	}

	tokenID := d.Id()
	tokenParts := strings.Split(tokenID, "!")

	comment := d.Get(mkResourceVirtualEnvironmentUserTokenComment).(string)
	expirationDate, err := time.Parse(time.RFC3339, d.Get(mkResourceVirtualEnvironmentUserTokenExpirationDate).(string))

	if err != nil {
		return err
	}

	expirationDateCustom := proxmox.CustomTimestamp(expirationDate)
	privilegeSeparation := proxmox.CustomBool(d.Get(mkResourceVirtualEnvironmentUserTokenPrivilegeSeparation).(bool))

	body := &proxmox.VirtualEnvironmentUserTokenUpdateRequestBody{
		Comment:             &comment,
		ExpirationDate:      &expirationDateCustom,
		PrivilegeSeperation: &privilegeSeparation,
	}

	_, err = veClient.UpdateUserToken(tokenParts[0], tokenParts[1], body)

	if err != nil {
		return err
	}

	aclArgOld, aclArg := d.GetChange(mkResourceVirtualEnvironmentUserTokenACL)
	aclParsedOld := aclArgOld.(*schema.Set).List()

	for _, v := range aclParsedOld {
		aclDelete := proxmox.CustomBool(true)
		aclEntry := v.(map[string]interface{})
		aclPropagate := proxmox.CustomBool(aclEntry[mkResourceVirtualEnvironmentUserTokenACLPropagate].(bool))

		aclBody := &proxmox.VirtualEnvironmentACLUpdateRequestBody{
			Delete:    &aclDelete,
			Path:      aclEntry[mkResourceVirtualEnvironmentUserTokenACLPath].(string),
			Propagate: &aclPropagate,
			Roles:     []string{aclEntry[mkResourceVirtualEnvironmentUserTokenACLRoleID].(string)},
			Tokens:    []string{tokenID},
		}

		err := veClient.UpdateACL(aclBody)

		if err != nil {
			return err
		}
	}

	aclParsed := aclArg.(*schema.Set).List()

	for _, v := range aclParsed {
		aclDelete := proxmox.CustomBool(false)
		aclEntry := v.(map[string]interface{})
		aclPropagate := proxmox.CustomBool(aclEntry[mkResourceVirtualEnvironmentUserTokenACLPropagate].(bool))

		aclBody := &proxmox.VirtualEnvironmentACLUpdateRequestBody{
			Delete:    &aclDelete,
			Path:      aclEntry[mkResourceVirtualEnvironmentUserTokenACLPath].(string),
			Propagate: &aclPropagate,
			Roles:     []string{aclEntry[mkResourceVirtualEnvironmentUserTokenACLRoleID].(string)},
			Tokens:    []string{tokenID},
		}

		err := veClient.UpdateACL(aclBody)

		if err != nil {
			return err
		}
	}

	return resourceVirtualEnvironmentUserTokenRead(d, m)
}

func resourceVirtualEnvironmentUserTokenDelete(d *schema.ResourceData, m interface{}) error {
	config := m.(providerConfiguration)
	veClient, err := config.GetVEClient()

	if err != nil {
		return err
	}

	tokenID := d.Id()
	tokenParts := strings.Split(tokenID, "!")

	aclParsed := d.Get(mkResourceVirtualEnvironmentUserTokenACL).(*schema.Set).List()

	for _, v := range aclParsed {
		aclDelete := proxmox.CustomBool(true)
		aclEntry := v.(map[string]interface{})
		aclPropagate := proxmox.CustomBool(aclEntry[mkResourceVirtualEnvironmentUserTokenACLPropagate].(bool))

		aclBody := &proxmox.VirtualEnvironmentACLUpdateRequestBody{
			Delete:    &aclDelete,
			Path:      aclEntry[mkResourceVirtualEnvironmentUserTokenACLPath].(string),
			Propagate: &aclPropagate,
			Roles:     []string{aclEntry[mkResourceVirtualEnvironmentUserTokenACLRoleID].(string)},
			Tokens:    []string{tokenID},
		}

		err := veClient.UpdateACL(aclBody)

		if err != nil {
			return err
		}
	}

	err = veClient.DeleteUserToken(tokenParts[0], tokenParts[1])

	if err != nil {
		if strings.Contains(err.Error(), "HTTP 404") || (strings.Contains(err.Error(), "HTTP 500") && strings.Contains(err.Error(), "no such token")) {
			d.SetId("")

			return nil
		}

		return err
	}

	d.SetId("")

	return nil
}
