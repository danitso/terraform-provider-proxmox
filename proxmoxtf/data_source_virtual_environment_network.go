/* This Source Code Form is subject to the terms of the Mozilla Public
 * License, v. 2.0. If a copy of the MPL was not distributed with this
 * file, You can obtain one at https://mozilla.org/MPL/2.0/. */

package proxmoxtf

import (
	"fmt"
	"github.com/hashicorp/terraform/helper/schema"
)

const (
	mkDataSourceVirtualEnvironmentNetworkNodeName    = "node_name"
	mkDataSourceVirtualEnvironmentNetworkActive      = "active"
	mkDataSourceVirtualEnvironmentNetworkAddress     = "address"
	mkDataSourceVirtualEnvironmentNetworkAutostart   = "autostart"
	mkDataSourceVirtualEnvironmentNetworkBridgeFD    = "bridge_fd"
	mkDataSourceVirtualEnvironmentNetworkBridgePorts = "bridge_ports"
	mkDataSourceVirtualEnvironmentNetworkBridgeSTP   = "bridge_stp"
	mkDataSourceVirtualEnvironmentNetworkCIDR        = "cidr"
	mkDataSourceVirtualEnvironmentNetworkExists      = "exists"
	mkDataSourceVirtualEnvironmentNetworkFamilies    = "families"
	mkDataSourceVirtualEnvironmentNetworkGateway     = "gateway"
	mkDataSourceVirtualEnvironmentNetworkIface       = "iface"
	mkDataSourceVirtualEnvironmentNetworkMethodIPv4  = "method_ipv4"
	mkDataSourceVirtualEnvironmentNetworkMethodIPv6  = "method_ipv6"
	mkDataSourceVirtualEnvironmentNetworkNetmask     = "netmask"
	mkDataSourceVirtualEnvironmentNetworkPriority    = "priority"
	mkDataSourceVirtualEnvironmentNetworkType        = "type"
)

func dataSourceVirtualEnvironmentNetwork() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			mkDataSourceVirtualEnvironmentNetworkNodeName: {
				Type:        schema.TypeString,
				Description: "The network active status",
				Required: true,
			},
			mkDataSourceVirtualEnvironmentNetworkType: {
				Type:        schema.TypeString,
				Description: "The network type",
				Optional: true,
			},
			mkDataSourceVirtualEnvironmentNetworkIface: {
				Type: schema.TypeString,
				Description: "The network interface",
				Optional: true,
				Computed: true,
			},
			mkDataSourceVirtualEnvironmentNetworkActive: {
				Type:        schema.TypeBool,
				Description: "The network active status",
				Computed:    true,
			},
			mkDataSourceVirtualEnvironmentNetworkAddress: {
				Type:        schema.TypeString,
				Description: "The network IP address",
				Computed:    true,
			},
			mkDataSourceVirtualEnvironmentNetworkCIDR: {
				Type:        schema.TypeString,
				Description: "The network CIDR",
				Computed:    true,
			},
			mkDataSourceVirtualEnvironmentNetworkNetmask:{
				Type:        schema.TypeString,
				Description: "The network netmask",
				Computed:    true,
			},
			mkDataSourceVirtualEnvironmentNetworkAutostart: {
				Type:        schema.TypeBool,
				Description: "The network autostart status",
				Computed:    true,
			},
			mkDataSourceVirtualEnvironmentNetworkBridgeFD:{
				Type:        schema.TypeInt,
				Description: "The network bridge forwarding delay",
				Computed:    true,
			},
			mkDataSourceVirtualEnvironmentNetworkBridgePorts:{
				Type:        schema.TypeString,
				Description: "The network bridge ports",
				Computed:    true,
			},
			mkDataSourceVirtualEnvironmentNetworkBridgeSTP:{
				Type:        schema.TypeString,
				Description: "The network bridge spanning tree protocol status",
				Computed:    true,
			},
			mkDataSourceVirtualEnvironmentNetworkPriority:{
				Type:        schema.TypeInt,
				Description: "The network priority",
				Computed:    true,
			},
			mkDataSourceVirtualEnvironmentNetworkMethodIPv4:{
				Type:        schema.TypeString,
				Description: "The network method for ipv4",
				Computed:    true,
			},
			mkDataSourceVirtualEnvironmentNetworkMethodIPv6:{
				Type:        schema.TypeString,
				Description: "The network method for ipv6",
				Computed:    true,
			},
			mkDataSourceVirtualEnvironmentNetworkExists:{
				Type:        schema.TypeBool,
				Description: "The network existed prior the request",
				Computed:    true,
			},
			mkDataSourceVirtualEnvironmentNetworkFamilies :{
				Type:        schema.TypeList,
				Description: "The network families",
				Computed:    true,
				Elem: &schema.Schema{Type: schema.TypeString},
			},
			mkDataSourceVirtualEnvironmentNetworkGateway:{
				Type:        schema.TypeString,
				Description: "The network gateway",
				Computed:    true,
			},
		},
		Read: dataSourceVirtualEnvironmentNetworkRead,
	}
}

func dataSourceVirtualEnvironmentNetworkRead(d *schema.ResourceData, m interface{}) error {
	config := m.(providerConfiguration)
	veClient, err := config.GetVEClient()

	if err != nil {
		return err
	}

	nodeName := d.Get(mkDataSourceVirtualEnvironmentNetworkNodeName).(string)
	iface := d.Get(mkDataSourceVirtualEnvironmentNetworkIface).(string)

	network, err := veClient.GetNetworkInterface(nodeName,iface)

	if err != nil {
		return err
	}
	d.SetId(fmt.Sprintf("%s_network_%s", nodeName,network.Iface))

	d.Set(mkDataSourceVirtualEnvironmentNetworkType,network.Type)
	d.Set(mkResourceVirtualEnvironmentContainerNetworkInterface,network.Iface)
	d.Set(mkDataSourceVirtualEnvironmentNetworkActive,network.Active)
	d.Set(mkDataSourceVirtualEnvironmentNetworkAddress,network.Address)
	d.Set(mkDataSourceVirtualEnvironmentNetworkCIDR,network.CIDR)
	d.Set(mkDataSourceVirtualEnvironmentNetworkNetmask,network.Netmask)
	d.Set(mkDataSourceVirtualEnvironmentNetworkAutostart,network.Autostart)
	d.Set(mkDataSourceVirtualEnvironmentNetworkBridgeFD,network.BridgeFD)
	d.Set(mkDataSourceVirtualEnvironmentNetworkBridgePorts,network.BridgePorts)
	d.Set(mkDataSourceVirtualEnvironmentNetworkBridgeSTP,network.BridgeSTP)
	d.Set(mkDataSourceVirtualEnvironmentNetworkPriority,network.Priority)
	d.Set(mkDataSourceVirtualEnvironmentNetworkMethodIPv4,network.MethodIPv4)
	d.Set(mkDataSourceVirtualEnvironmentNetworkMethodIPv6,network.MethodIPv6)

	d.Set(mkDataSourceVirtualEnvironmentNetworkExists,network.Exists)
	d.Set(mkDataSourceVirtualEnvironmentNetworkFamilies,network.Families)
	d.Set(mkDataSourceVirtualEnvironmentNetworkGateway,network.Gateway)


	return nil
}

