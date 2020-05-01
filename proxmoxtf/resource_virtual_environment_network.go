/* This Source Code Form is subject to the terms of the Mozilla Public
 * License, v. 2.0. If a copy of the MPL was not distributed with this
 * file, You can obtain one at https://mozilla.org/MPL/2.0/. */

package proxmoxtf

import (
	"fmt"
	"github.com/hashicorp/terraform/helper/schema"
)

const (
	mkResourceVirtualEnvironmentNetworkNodeName    = "node_name"
	mkResourceVirtualEnvironmentNetworkActive      = "active"
	mkResourceVirtualEnvironmentNetworkAddress     = "address"
	mkResourceVirtualEnvironmentNetworkAutostart   = "autostart"
	mkResourceVirtualEnvironmentNetworkBridgeFD    = "bridge_fd"
	mkResourceVirtualEnvironmentNetworkBridgePorts = "bridge_ports"
	mkResourceVirtualEnvironmentNetworkBridgeSTP   = "bridge_stp"
	mkResourceVirtualEnvironmentNetworkCIDR        = "cidr"
	mkResourceVirtualEnvironmentNetworkExists      = "exists"
	mkResourceVirtualEnvironmentNetworkFamilies    = "families"
	mkResourceVirtualEnvironmentNetworkGateway     = "gateway"
	mkResourceVirtualEnvironmentNetworkIface       = "iface"
	mkResourceVirtualEnvironmentNetworkMethodIPv4  = "method_ipv4"
	mkResourceVirtualEnvironmentNetworkMethodIPv6  = "method_ipv6"
	mkResourceVirtualEnvironmentNetworkNetmask     = "netmask"
	mkResourceVirtualEnvironmentNetworkPriority    = "priority"
	mkResourceVirtualEnvironmentNetworkType        = "type"
)

func resourceVirtualEnvironmentNetwork() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			mkResourceVirtualEnvironmentNetworkNodeName: {
				Type:        schema.TypeString,
				Description: "The network active status",
				Required:    true,
			},
			mkResourceVirtualEnvironmentNetworkType: {
				Type:        schema.TypeString,
				Description: "The network type",
				Required:    true,
			},
			mkResourceVirtualEnvironmentNetworkIface: {
				Type:        schema.TypeString,
				Description: "The network interface",
				Required:    true,
			},
			mkResourceVirtualEnvironmentNetworkActive: {
				Type:        schema.TypeBool,
				Description: "The network active status",
				Optional:    true,
			},
			mkResourceVirtualEnvironmentNetworkAddress: {
				Type:        schema.TypeString,
				Description: "The network IP address",
				Optional:    true,
			},
			mkResourceVirtualEnvironmentNetworkCIDR: {
				Type:        schema.TypeString,
				Description: "The network CIDR",
				Optional:    true,
			},
			mkResourceVirtualEnvironmentNetworkNetmask: {
				Type:        schema.TypeString,
				Description: "The network netmask",
				Optional:    true,
			},
			mkResourceVirtualEnvironmentNetworkAutostart: {
				Type:        schema.TypeBool,
				Description: "The network autostart status",
				Optional:    true,
			},
			mkResourceVirtualEnvironmentNetworkBridgeFD: {
				Type:        schema.TypeInt,
				Description: "The network bridge forwarding delay",
				Optional:    true,
			},
			mkResourceVirtualEnvironmentNetworkBridgePorts: {
				Type:        schema.TypeString,
				Description: "The network bridge ports",
				Optional:    true,
			},
			mkResourceVirtualEnvironmentNetworkBridgeSTP: {
				Type:        schema.TypeString,
				Description: "The network bridge spanning tree protocol status",
				Optional:    true,
			},
			mkResourceVirtualEnvironmentNetworkPriority: {
				Type:        schema.TypeInt,
				Description: "The network priority",
				Optional:    true,
			},
			mkResourceVirtualEnvironmentNetworkMethodIPv4: {
				Type:        schema.TypeString,
				Description: "The network method for ipv4",
				Optional:    true,
			},
			mkResourceVirtualEnvironmentNetworkMethodIPv6: {
				Type:        schema.TypeString,
				Description: "The network method for ipv6",
				Optional:    true,
			},
			mkResourceVirtualEnvironmentNetworkExists: {
				Type:        schema.TypeBool,
				Description: "The network existed prior the request",
				Optional:    true,
			},
			mkResourceVirtualEnvironmentNetworkFamilies: {
				Type:        schema.TypeList,
				Description: "The network families",
				Optional:    true,
				Elem:        &schema.Schema{Type: schema.TypeString},
			},
			mkResourceVirtualEnvironmentNetworkGateway: {
				Type:        schema.TypeString,
				Description: "The network gateway",
				Computed:    true,
			},
		},
		Read: resourceVirtualEnvironmentNetworkRead,
		Create: resourceVirtualEnvironmentNetworkCreate,
		Update: resourceVirtualEnvironmentNetworkUpdate,
		Delete: resourceVirtualEnvironmentNetworkDelete,
	}
}

func resourceVirtualEnvironmentNetworkDelete(d *schema.ResourceData, m interface{}) error {

}

func resourceVirtualEnvironmentNetworkUpdate(d *schema.ResourceData, m interface{}) error {

}

func resourceVirtualEnvironmentNetworkCreate(d *schema.ResourceData, m interface{}) error {

}

func resourceVirtualEnvironmentNetworkRead(d *schema.ResourceData, m interface{}) error {
	config := m.(providerConfiguration)
	veClient, err := config.GetVEClient()

	if err != nil {
		return err
	}

	nodeName := d.Get(mkResourceVirtualEnvironmentNetworkNodeName).(string)
	iface := d.Get(mkResourceVirtualEnvironmentNetworkIface).(string)

	network, err := veClient.GetNetworkInterface(nodeName, iface)

	if err != nil {
		return err
	}
	d.SetId(fmt.Sprintf("%s_network_%s", nodeName, network.Iface))

	d.Set(mkResourceVirtualEnvironmentNetworkType, network.Type)
	d.Set(mkResourceVirtualEnvironmentContainerNetworkInterface, network.Iface)
	d.Set(mkResourceVirtualEnvironmentNetworkActive, network.Active)
	d.Set(mkResourceVirtualEnvironmentNetworkAddress, network.Address)
	d.Set(mkResourceVirtualEnvironmentNetworkCIDR, network.CIDR)
	d.Set(mkResourceVirtualEnvironmentNetworkNetmask, network.Netmask)
	d.Set(mkResourceVirtualEnvironmentNetworkAutostart, network.Autostart)
	d.Set(mkResourceVirtualEnvironmentNetworkBridgeFD, network.BridgeFD)
	d.Set(mkResourceVirtualEnvironmentNetworkBridgePorts, network.BridgePorts)
	d.Set(mkResourceVirtualEnvironmentNetworkBridgeSTP, network.BridgeSTP)
	d.Set(mkResourceVirtualEnvironmentNetworkPriority, network.Priority)
	d.Set(mkResourceVirtualEnvironmentNetworkMethodIPv4, network.MethodIPv4)
	d.Set(mkResourceVirtualEnvironmentNetworkMethodIPv6, network.MethodIPv6)

	d.Set(mkResourceVirtualEnvironmentNetworkExists, network.Exists)
	d.Set(mkResourceVirtualEnvironmentNetworkFamilies, network.Families)
	d.Set(mkResourceVirtualEnvironmentNetworkGateway, network.Gateway)

	return nil
}
