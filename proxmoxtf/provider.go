/* This Source Code Form is subject to the terms of the Mozilla Public
 * License, v. 2.0. If a copy of the MPL was not distributed with this
 * file, You can obtain one at https://mozilla.org/MPL/2.0/. */

package proxmoxtf

import (
	"errors"
	"fmt"
	"github.com/danitso/terraform-provider-proxmox/proxmox"
	"github.com/hashicorp/terraform/helper/schema"
	"net/url"
	"os"
	"strings"
)

const (
	dvProviderVirtualEnvironmentEndpoint = ""
	dvProviderVirtualEnvironmentOTP      = ""
	dvProviderVirtualEnvironmentPassword = ""
	dvProviderVirtualEnvironmentUsername = ""

	mkProviderVirtualEnvironment         = "virtual_environment"
	mkProviderVirtualEnvironmentEndpoint = "endpoint"
	mkProviderVirtualEnvironmentInsecure = "insecure"
	mkProviderVirtualEnvironmentOTP      = "otp"
	mkProviderVirtualEnvironmentPassword = "password"
	mkProviderVirtualEnvironmentUsername = "username"
)

type ProviderConfiguration struct {
	veClient *proxmox.VirtualEnvironmentClient
}

// Provider returns the object for this provider.
func Provider() *schema.Provider {
	return &schema.Provider{
		ConfigureFunc: providerConfigure,
		DataSourcesMap: map[string]*schema.Resource{
			"proxmox_virtual_environment_datastores": dataSourceVirtualEnvironmentDatastores(),
			"proxmox_virtual_environment_dns":        dataSourceVirtualEnvironmentDNS(),
			"proxmox_virtual_environment_group":      dataSourceVirtualEnvironmentGroup(),
			"proxmox_virtual_environment_groups":     dataSourceVirtualEnvironmentGroups(),
			"proxmox_virtual_environment_hosts":      dataSourceVirtualEnvironmentHosts(),
			"proxmox_virtual_environment_nodes":      dataSourceVirtualEnvironmentNodes(),
			"proxmox_virtual_environment_pool":       dataSourceVirtualEnvironmentPool(),
			"proxmox_virtual_environment_pools":      dataSourceVirtualEnvironmentPools(),
			"proxmox_virtual_environment_role":       dataSourceVirtualEnvironmentRole(),
			"proxmox_virtual_environment_roles":      dataSourceVirtualEnvironmentRoles(),
			"proxmox_virtual_environment_time":       dataSourceVirtualEnvironmentTime(),
			"proxmox_virtual_environment_user":       dataSourceVirtualEnvironmentUser(),
			"proxmox_virtual_environment_users":      dataSourceVirtualEnvironmentUsers(),
			"proxmox_virtual_environment_version":    dataSourceVirtualEnvironmentVersion(),
		},
		ResourcesMap: map[string]*schema.Resource{
			"proxmox_virtual_environment_certificate": resourceVirtualEnvironmentCertificate(),
			"proxmox_virtual_environment_container":   resourceVirtualEnvironmentContainer(),
			"proxmox_virtual_environment_dns":         resourceVirtualEnvironmentDNS(),
			"proxmox_virtual_environment_file":        resourceVirtualEnvironmentFile(),
			"proxmox_virtual_environment_group":       resourceVirtualEnvironmentGroup(),
			"proxmox_virtual_environment_hosts":       resourceVirtualEnvironmentHosts(),
			"proxmox_virtual_environment_pool":        resourceVirtualEnvironmentPool(),
			"proxmox_virtual_environment_role":        resourceVirtualEnvironmentRole(),
			"proxmox_virtual_environment_time":        resourceVirtualEnvironmentTime(),
			"proxmox_virtual_environment_user":        resourceVirtualEnvironmentUser(),
			"proxmox_virtual_environment_vm":          resourceVirtualEnvironmentVM(),
		},
		Schema: map[string]*schema.Schema{
			mkProviderVirtualEnvironment: {
				Type:     schema.TypeList,
				Optional: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						mkProviderVirtualEnvironmentEndpoint: {
							Type:        schema.TypeString,
							Optional:    true,
							Description: "The endpoint for the Proxmox Virtual Environment API",
							DefaultFunc: schema.MultiEnvDefaultFunc(
								[]string{"PROXMOX_VE_ENDPOINT", "PM_VE_ENDPOINT"},
								dvProviderVirtualEnvironmentEndpoint,
							),
							ValidateFunc: func(v interface{}, k string) (warns []string, errs []error) {
								value := v.(string)

								if value == "" {
									return []string{}, []error{
										errors.New("You must specify an endpoint for the Proxmox Virtual Environment API (valid: https://host:port)"),
									}
								}

								_, err := url.ParseRequestURI(value)

								if err != nil {
									return []string{}, []error{
										errors.New("You must specify a valid endpoint for the Proxmox Virtual Environment API (valid: https://host:port)"),
									}
								}

								return []string{}, []error{}
							},
						},
						mkProviderVirtualEnvironmentInsecure: {
							Type:        schema.TypeBool,
							Optional:    true,
							Description: "Whether to skip the TLS verification step",
							DefaultFunc: func() (interface{}, error) {
								for _, k := range []string{"PROXMOX_VE_INSECURE", "PM_VE_INSECURE"} {
									v := os.Getenv(k)

									if v == "true" || v == "1" {
										return true, nil
									}
								}

								return false, nil
							},
						},
						mkProviderVirtualEnvironmentOTP: {
							Type:        schema.TypeString,
							Optional:    true,
							Description: "The one-time password for the Proxmox Virtual Environment API",
							DefaultFunc: schema.MultiEnvDefaultFunc(
								[]string{"PROXMOX_VE_OTP", "PM_VE_OTP"},
								dvProviderVirtualEnvironmentOTP,
							),
						},
						mkProviderVirtualEnvironmentPassword: {
							Type:        schema.TypeString,
							Optional:    true,
							Description: "The password for the Proxmox Virtual Environment API",
							DefaultFunc: schema.MultiEnvDefaultFunc(
								[]string{"PROXMOX_VE_PASSWORD", "PM_VE_PASSWORD"},
								dvProviderVirtualEnvironmentPassword,
							),
							ValidateFunc: func(v interface{}, k string) (warns []string, errs []error) {
								value := v.(string)

								if value == "" {
									return []string{}, []error{
										errors.New("You must specify a password for the Proxmox Virtual Environment API"),
									}
								}

								return []string{}, []error{}
							},
						},
						mkProviderVirtualEnvironmentUsername: {
							Type:        schema.TypeString,
							Optional:    true,
							Description: "The username for the Proxmox Virtual Environment API",
							DefaultFunc: schema.MultiEnvDefaultFunc(
								[]string{"PROXMOX_VE_USERNAME", "PM_VE_USERNAME"},
								dvProviderVirtualEnvironmentUsername,
							),
							ValidateFunc: func(v interface{}, k string) (warns []string, errs []error) {
								value := v.(string)

								if value == "" {
									return []string{}, []error{
										errors.New("You must specify a username for the Proxmox Virtual Environment API (valid: username@realm)"),
									}
								}

								return []string{}, []error{}
							},
						},
					},
				},
				MaxItems: 1,
			},
		},
	}
}

func providerConfigure(d *schema.ResourceData) (interface{}, error) {
	var err error
	var veClient *proxmox.VirtualEnvironmentClient

	// Initialize the client for the Virtual Environment, if required.
	veConfigBlock := d.Get(mkProviderVirtualEnvironment).([]interface{})
	_, AcceptanceTestFlag := os.LookupEnv("TF_ACC")

	// Initialize provider from environmental variables if `TF_ACC` (acceptance testing flag) was set
	if AcceptanceTestFlag {
		envVeConfig := make(map[string]interface{})

		for _, e := range []string{
			mkProviderVirtualEnvironmentEndpoint,
			mkProviderVirtualEnvironmentUsername,
			mkProviderVirtualEnvironmentPassword,
			mkProviderVirtualEnvironmentInsecure,
			mkProviderVirtualEnvironmentOTP,
		} {
			envVeConfig[e] = ""
			upEnv := strings.ToUpper(e)
			if v, ok := os.LookupEnv(fmt.Sprintf("PROXMOX_VE_%s", upEnv)); ok {
				envVeConfig[e] = v
			} else if v, ok := os.LookupEnv(fmt.Sprintf("PM_VE_%s", upEnv)); ok {
				envVeConfig[e] = v
			}
		}

		if insecureEnv := os.Getenv("PROXMOX_VE_INSECURE"); insecureEnv == "true" || insecureEnv == "1" {
			envVeConfig[mkProviderVirtualEnvironmentInsecure] = true
		} else if insecureEnv := os.Getenv("PM_VE_INSECURE"); insecureEnv == "true" || insecureEnv == "1" {
			envVeConfig[mkProviderVirtualEnvironmentInsecure] = true
		} else {
			envVeConfig[mkProviderVirtualEnvironmentInsecure] = false
		}

		veConfigBlock = append(veConfigBlock, envVeConfig)
	}

	if len(veConfigBlock) > 0 {
		veConfig := veConfigBlock[0].(map[string]interface{})

		veClient, err = proxmox.NewVirtualEnvironmentClient(
			veConfig[mkProviderVirtualEnvironmentEndpoint].(string),
			veConfig[mkProviderVirtualEnvironmentUsername].(string),
			veConfig[mkProviderVirtualEnvironmentPassword].(string),
			veConfig[mkProviderVirtualEnvironmentOTP].(string),
			veConfig[mkProviderVirtualEnvironmentInsecure].(bool),
		)

		if err != nil {
			return nil, err
		}

	}

	config := ProviderConfiguration{
		veClient: veClient,
	}

	return config, nil
}

func (c *ProviderConfiguration) GetVEClient() (*proxmox.VirtualEnvironmentClient, error) {
	if c.veClient == nil {
		return nil, errors.New("You must specify the virtual environment details in the provider configuration")
	}

	return c.veClient, nil
}
