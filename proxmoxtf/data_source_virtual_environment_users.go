/* This Source Code Form is subject to the terms of the Mozilla Public
 * License, v. 2.0. If a copy of the MPL was not distributed with this
 * file, You can obtain one at https://mozilla.org/MPL/2.0/. */

package proxmoxtf

import (
	"time"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
)

const (
	mkDataSourceVirtualEnvironmentUsersComments                  = "comments"
	mkDataSourceVirtualEnvironmentUsersEmails                    = "emails"
	mkDataSourceVirtualEnvironmentUsersEnabled                   = "enabled"
	mkDataSourceVirtualEnvironmentUsersExpirationDates           = "expiration_dates"
	mkDataSourceVirtualEnvironmentUsersFirstNames                = "first_names"
	mkDataSourceVirtualEnvironmentUsersGroups                    = "groups"
	mkDataSourceVirtualEnvironmentUsersKeys                      = "keys"
	mkDataSourceVirtualEnvironmentUsersLastNames                 = "last_names"
	mkDataSourceVirtualEnvironmentUsersTokens                    = "tokens"
	mkDataSourceVirtualEnvironmentUsersTokensComment             = "comment"
	mkDataSourceVirtualEnvironmentUsersTokensExpirationDate      = "expiration_date"
	mkDataSourceVirtualEnvironmentUsersTokensID                  = "id"
	mkDataSourceVirtualEnvironmentUsersTokensPrivilegeSeparation = "privilege_separation"
	mkDataSourceVirtualEnvironmentUsersUserIDs                   = "user_ids"
)

func dataSourceVirtualEnvironmentUsers() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			mkDataSourceVirtualEnvironmentUsersComments: {
				Type:        schema.TypeList,
				Description: "The user comments",
				Computed:    true,
				Elem:        &schema.Schema{Type: schema.TypeString},
			},
			mkDataSourceVirtualEnvironmentUsersEmails: {
				Type:        schema.TypeList,
				Description: "The users' email addresses",
				Computed:    true,
				Elem:        &schema.Schema{Type: schema.TypeString},
			},
			mkDataSourceVirtualEnvironmentUsersEnabled: {
				Type:        schema.TypeList,
				Description: "Whether a user account is enabled",
				Computed:    true,
				Elem:        &schema.Schema{Type: schema.TypeBool},
			},
			mkDataSourceVirtualEnvironmentUsersExpirationDates: {
				Type:        schema.TypeList,
				Description: "The user accounts' expiration dates",
				Computed:    true,
				Elem:        &schema.Schema{Type: schema.TypeString},
			},
			mkDataSourceVirtualEnvironmentUsersFirstNames: {
				Type:        schema.TypeList,
				Description: "The users' first names",
				Computed:    true,
				Elem:        &schema.Schema{Type: schema.TypeString},
			},
			mkDataSourceVirtualEnvironmentUsersGroups: {
				Type:        schema.TypeList,
				Description: "The users' groups",
				Computed:    true,
				Elem: &schema.Schema{
					Type: schema.TypeList,
					Elem: &schema.Schema{Type: schema.TypeString},
				},
			},
			mkDataSourceVirtualEnvironmentUsersKeys: {
				Type:        schema.TypeList,
				Description: "The users' keys",
				Computed:    true,
				Elem:        &schema.Schema{Type: schema.TypeString},
			},
			mkDataSourceVirtualEnvironmentUsersLastNames: {
				Type:        schema.TypeList,
				Description: "The users' last names",
				Computed:    true,
				Elem:        &schema.Schema{Type: schema.TypeString},
			},
			mkDataSourceVirtualEnvironmentUsersTokens: {
				Type:        schema.TypeList,
				Description: "The users' API tokens",
				Computed:    true,
				Elem: &schema.Schema{
					Type: schema.TypeSet,
					Elem: &schema.Resource{
						Schema: map[string]*schema.Schema{
							mkDataSourceVirtualEnvironmentUsersTokensComment: {
								Type:        schema.TypeString,
								Computed:    true,
								Description: "The token's comment",
							},
							mkDataSourceVirtualEnvironmentUsersTokensExpirationDate: {
								Type:        schema.TypeString,
								Computed:    true,
								Description: "The token's expiration date",
							},
							mkDataSourceVirtualEnvironmentUsersTokensID: {
								Type:        schema.TypeString,
								Computed:    true,
								Description: "The token's identifier",
							},
							mkDataSourceVirtualEnvironmentUsersTokensPrivilegeSeparation: {
								Type:        schema.TypeBool,
								Computed:    true,
								Description: "Whether the privileges for the token differs from the account privileges",
							},
						},
					},
				},
			},
			mkDataSourceVirtualEnvironmentUsersUserIDs: {
				Type:        schema.TypeList,
				Description: "The user ids",
				Computed:    true,
				Elem:        &schema.Schema{Type: schema.TypeString},
			},
		},
		Read: dataSourceVirtualEnvironmentUsersRead,
	}
}

func dataSourceVirtualEnvironmentUsersRead(d *schema.ResourceData, m interface{}) error {
	config := m.(providerConfiguration)
	veClient, err := config.GetVEClient()

	if err != nil {
		return err
	}

	list, err := veClient.ListUsers(true, true)

	if err != nil {
		return err
	}

	comments := make([]interface{}, len(list))
	emails := make([]interface{}, len(list))
	enabled := make([]interface{}, len(list))
	expirationDates := make([]interface{}, len(list))
	firstNames := make([]interface{}, len(list))
	groups := make([]interface{}, len(list))
	keys := make([]interface{}, len(list))
	lastNames := make([]interface{}, len(list))
	tokens := make([]interface{}, len(list))
	userIDs := make([]interface{}, len(list))

	for i, v := range list {
		if v.Comment != nil {
			comments[i] = v.Comment
		} else {
			comments[i] = ""
		}

		if v.Email != nil {
			emails[i] = v.Email
		} else {
			emails[i] = ""
		}

		if v.Enabled != nil {
			enabled[i] = v.Enabled
		} else {
			enabled[i] = true
		}

		if v.ExpirationDate != nil {
			t := time.Time(*v.ExpirationDate)

			if t.Unix() > 0 {
				expirationDates[i] = t.UTC().Format(time.RFC3339)
			} else {
				expirationDates[i] = time.Unix(0, 0).UTC().Format(time.RFC3339)
			}
		} else {
			expirationDates[i] = time.Unix(0, 0).UTC().Format(time.RFC3339)
		}

		if v.FirstName != nil {
			firstNames[i] = v.FirstName
		} else {
			firstNames[i] = ""
		}

		if v.Groups != nil {
			groups[i] = v.Groups
		} else {
			groups[i] = []string{}
		}

		if v.Keys != nil {
			keys[i] = v.Keys
		} else {
			keys[i] = ""
		}

		if v.LastName != nil {
			lastNames[i] = v.LastName
		} else {
			lastNames[i] = ""
		}

		if v.Tokens != nil {
			userTokens := make([]interface{}, len(*v.Tokens))

			for ti, tv := range *v.Tokens {
				userToken := map[string]interface{}{}

				if tv.Comment != nil {
					userToken[mkDataSourceVirtualEnvironmentUsersTokensComment] = *tv.Comment
				} else {
					userToken[mkDataSourceVirtualEnvironmentUsersTokensComment] = ""
				}

				if tv.ExpirationDate != nil {
					t := time.Time(*tv.ExpirationDate)

					if t.Unix() > 0 {
						userToken[mkDataSourceVirtualEnvironmentUsersTokensExpirationDate] = t.UTC().Format(time.RFC3339)
					} else {
						userToken[mkDataSourceVirtualEnvironmentUsersTokensExpirationDate] = time.Unix(0, 0).UTC().Format(time.RFC3339)
					}
				} else {
					userToken[mkDataSourceVirtualEnvironmentUsersTokensExpirationDate] = time.Unix(0, 0).UTC().Format(time.RFC3339)
				}

				if tv.ID != nil {
					userToken[mkDataSourceVirtualEnvironmentUsersTokensID] = *tv.ID
				} else {
					userToken[mkDataSourceVirtualEnvironmentUsersTokensID] = ""
				}

				if tv.PrivilegeSeperation != nil {
					userToken[mkDataSourceVirtualEnvironmentUsersTokensPrivilegeSeparation] = *tv.PrivilegeSeperation
				} else {
					userToken[mkDataSourceVirtualEnvironmentUsersTokensPrivilegeSeparation] = true
				}

				userTokens[ti] = userToken
			}

			tokens[i] = userTokens
		} else {
			tokens[i] = []interface{}{}
		}

		userIDs[i] = v.ID
	}

	d.SetId("users")

	d.Set(mkDataSourceVirtualEnvironmentUsersComments, comments)
	d.Set(mkDataSourceVirtualEnvironmentUsersEmails, emails)
	d.Set(mkDataSourceVirtualEnvironmentUsersEnabled, enabled)
	d.Set(mkDataSourceVirtualEnvironmentUsersExpirationDates, expirationDates)
	d.Set(mkDataSourceVirtualEnvironmentUsersFirstNames, firstNames)
	d.Set(mkDataSourceVirtualEnvironmentUsersGroups, groups)
	d.Set(mkDataSourceVirtualEnvironmentUsersKeys, keys)
	d.Set(mkDataSourceVirtualEnvironmentUsersLastNames, lastNames)
	d.Set(mkDataSourceVirtualEnvironmentUsersTokens, tokens)
	d.Set(mkDataSourceVirtualEnvironmentUsersUserIDs, userIDs)

	return nil
}
