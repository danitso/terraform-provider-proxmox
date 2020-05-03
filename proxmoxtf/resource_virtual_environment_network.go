/* This Source Code Form is subject to the terms of the Mozilla Public
 * License, v. 2.0. If a copy of the MPL was not distributed with this
 * file, You can obtain one at https://mozilla.org/MPL/2.0/. */

package proxmoxtf

import (
	"fmt"
	"github.com/danitso/terraform-provider-proxmox/proxmox"
	"github.com/hashicorp/terraform/helper/schema"
	"strings"
)

const (
	mkResourceVirtualEnvironmentNetworkIface              = "iface"
	mkResourceVirtualEnvironmentNetworkNodeName           = "node_name"
	mkResourceVirtualEnvironmentNetworkType               = "type"
	mkResourceVirtualEnvironmentNetworkAddress            = "address"
	mkResourceVirtualEnvironmentNetworkAddress6           = "address6"
	mkResourceVirtualEnvironmentNetworkAutostart          = "autostart"
	mkResourceVirtualEnvironmentNetworkBondPrimary        = "bond-primary"
	mkResourceVirtualEnvironmentNetworkBondMode           = "bond_mode"
	mkResourceVirtualEnvironmentNetworkBondXmitHashPolicy = "bond_xmit_hash_policy"
	mkResourceVirtualEnvironmentNetworkBridgePorts        = "bridge_ports"
	mkResourceVirtualEnvironmentNetworkBridgeVlanAware    = "bridge_vlan_aware"
	mkResourceVirtualEnvironmentNetworkCIDR               = "cidr"
	mkResourceVirtualEnvironmentNetworkCIDR6              = "cidr6"
	mkResourceVirtualEnvironmentNetworkComments           = "comments"
	mkResourceVirtualEnvironmentNetworkComments6          = "comments6"
	mkResourceVirtualEnvironmentNetworkGateway            = "gateway"
	mkResourceVirtualEnvironmentNetworkGateway6           = "gateway6"
	mkResourceVirtualEnvironmentNetworkMtu                = "mtu"
	mkResourceVirtualEnvironmentNetworkNetmask            = "netmask"
	mkResourceVirtualEnvironmentNetworkNetmask6           = "netmask6"
	mkResourceVirtualEnvironmentNetworkOvsBond            = "ovs_bonds"
	mkResourceVirtualEnvironmentNetworkOvsBridge          = "ovs_bridge"
	mkResourceVirtualEnvironmentNetworkOvsOptions         = "ovs_options"
	mkResourceVirtualEnvironmentNetworkOvsPorts           = "ovs_ports"
	mkResourceVirtualEnvironmentNetworkOvsTag             = "ovs_tag"
	mkResourceVirtualEnvironmentNetworkOvsSlaves          = "slaves"
	mkResourceVirtualEnvironmentNetworkVlanID             = "vlan-id"
	mkResourceVirtualEnvironmentNetworkVlanRawDevice      = "vlan-raw-device"
)

func resourceVirtualEnvironmentNetwork() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			mkResourceVirtualEnvironmentNetworkIface: {
				Type:        schema.TypeString,
				Description: "The network interface",
				Optional:    true,
				ForceNew:    true,
			},
			mkResourceVirtualEnvironmentNetworkNodeName: {
				Type:        schema.TypeString,
				Description: "The network active status",
				Optional:    true,
			},
			mkResourceVirtualEnvironmentNetworkType: {
				Type:        schema.TypeString,
				Description: "The network type",
				Required:    true,
			},
			mkResourceVirtualEnvironmentNetworkAddress: {
				Type:        schema.TypeString,
				Description: "The network IP address",
				Optional:    true,
			},
			mkResourceVirtualEnvironmentNetworkAddress6: {
				Type:        schema.TypeString,
				Description: "The network IPv6 address",
				Optional:    true,
			},
			mkResourceVirtualEnvironmentNetworkAutostart: {
				Type:        schema.TypeBool,
				Description: "The network autostart status",
				Optional:    true,
			},
			mkResourceVirtualEnvironmentNetworkBondPrimary: {
				Type:        schema.TypeString,
				Description: "Specify the primary interface for the active-backup bond",
				Optional:    true,
			},
			mkResourceVirtualEnvironmentNetworkBondMode: {
				Type:        schema.TypeString,
				Description: "Bonding mode",
				Optional:    true,
			},
			mkResourceVirtualEnvironmentNetworkBondXmitHashPolicy: {
				Type:        schema.TypeString,
				Description: "Selects the transmit hash policy to use for slave selection in balance-xor and 802.3ad modes.",
				Optional:    true,
			},
			mkResourceVirtualEnvironmentNetworkBridgePorts: {
				Type:        schema.TypeString,
				Description: "The network bridge ports",
				Optional:    true,
			},
			mkResourceVirtualEnvironmentNetworkBridgeVlanAware: {
				Type:        schema.TypeBool,
				Description: "Enable bridge vlan support",
				Optional:    true,
			},
			mkResourceVirtualEnvironmentNetworkCIDR: {
				Type:        schema.TypeString,
				Description: "The network IPv4 CIDR",
				Optional:    true,
			},
			mkResourceVirtualEnvironmentNetworkCIDR6: {
				Type:        schema.TypeString,
				Description: "The network IPv6 CIDR",
				Optional:    true,
			},
			mkResourceVirtualEnvironmentNetworkComments: {
				Type:        schema.TypeString,
				Description: "Comments",
				Optional:    true,
			},
			mkResourceVirtualEnvironmentNetworkComments6: {
				Type:        schema.TypeString,
				Description: "Comments",
				Optional:    true,
			},
			mkResourceVirtualEnvironmentNetworkGateway: {
				Type:        schema.TypeString,
				Description: "Default gateway address",
				Computed:    true,
			},
			mkResourceVirtualEnvironmentNetworkGateway6: {
				Type:        schema.TypeString,
				Description: "Default ipv6 gateway address",
				Computed:    true,
			},
			mkResourceVirtualEnvironmentNetworkMtu: {
				Type:        schema.TypeInt,
				Description: "MTU",
				Computed:    true,
			},
			mkResourceVirtualEnvironmentNetworkNetmask: {
				Type:        schema.TypeString,
				Description: "The network netmask",
				Optional:    true,
			},
			mkResourceVirtualEnvironmentNetworkNetmask6: {
				Type:        schema.TypeString,
				Description: "The network ipv6 netmask",
				Optional:    true,
			},
			mkResourceVirtualEnvironmentNetworkOvsBond: {
				Type:        schema.TypeString,
				Description: "Specify the interfaces used by the bonding device",
				Optional:    true,
			},
			mkResourceVirtualEnvironmentNetworkOvsBridge: {
				Type:        schema.TypeString,
				Description: "The OVS bridge associated with a OVS port. This is required when you create an OVS port",
				Optional:    true,
			},
			mkResourceVirtualEnvironmentNetworkOvsOptions: {
				Type:        schema.TypeString,
				Description: "OVS interface options",
				Optional:    true,
			},
			mkResourceVirtualEnvironmentNetworkOvsPorts: {
				Type:        schema.TypeString,
				Description: "Specify the interfaces you want to add to your bridge",
				Optional:    true,
			},
			mkResourceVirtualEnvironmentNetworkOvsTag: {
				Type:        schema.TypeInt,
				Description: "Specify the Vlan tag (used by OVSPort, OVSIntPort, OVSBond)",
				Optional:    true,
			},
			mkResourceVirtualEnvironmentNetworkOvsSlaves: {
				Type:        schema.TypeString,
				Description: "Specify the interfaces used byt the bonding device",
				Optional:    true,
			},
			mkResourceVirtualEnvironmentNetworkVlanID: {
				Type:        schema.TypeInt,
				Description: "vlan-id for a custom named vlan interface (ifupdown2 only)",
				Optional:    true,
			},
			mkResourceVirtualEnvironmentNetworkVlanRawDevice: {
				Type:        schema.TypeString,
				Description: "Specify the raw interface for the vlan interface",
				Optional:    true,
			},
		},
		Read:   resourceVirtualEnvironmentNetworkRead,
		Create: resourceVirtualEnvironmentNetworkCreate,
		Update: resourceVirtualEnvironmentNetworkUpdate,
		Delete: resourceVirtualEnvironmentNetworkDelete,
	}
}

func resourceVirtualEnvironmentNetworkDelete(d *schema.ResourceData, m interface{}) error {
	config := m.(providerConfiguration)
	veClient, err := config.GetVEClient()

	if err != nil {
		return err
	}
	nodeName := d.Get(mkResourceVirtualEnvironmentNetworkNodeName).(string)
	iface := d.Get(mkResourceVirtualEnvironmentNetworkIface).(string)

	d.SetId(fmt.Sprintf("%s_network_%s", nodeName, strings.ReplaceAll(iface, ".", "_")))

	err = veClient.DeleteNetworkInterface(nodeName, iface)

	if err != nil {
		return err
	}

	d.SetId("")

	return nil
}

func resourceVirtualEnvironmentNetworkUpdate(d *schema.ResourceData, m interface{}) error {
	config := m.(providerConfiguration)
	veClient, err := config.GetVEClient()

	if err != nil {
		return err
	}

	nodeName := d.Get(mkResourceVirtualEnvironmentNetworkNodeName).(string)
	iface := d.Get(mkResourceVirtualEnvironmentNetworkIface).(string)

	d.SetId(fmt.Sprintf("%s_network_%s", nodeName, strings.ReplaceAll(iface, ".", "_")))

	body := getBody(d, true)
	err = veClient.UpdateNetworkInterface(nodeName, iface, body)

	if err != nil {
		return err
	}

	return nil
}

func resourceVirtualEnvironmentNetworkCreate(d *schema.ResourceData, m interface{}) error {
	config := m.(providerConfiguration)
	veClient, err := config.GetVEClient()

	if err != nil {
		return err
	}
	nodeName := d.Get(mkResourceVirtualEnvironmentNetworkNodeName).(string)
	iface := d.Get(mkResourceVirtualEnvironmentNetworkIface).(string)

	d.SetId(fmt.Sprintf("%s_network_%s", nodeName, strings.ReplaceAll(iface, ".", "_")))

	body := getBody(d, false)
	err = veClient.CreateNetworkInterface(nodeName, body)

	if err != nil {
		return err
	}

	return resourceVirtualEnvironmentNetworkRead(d, m)
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

	d.SetId(fmt.Sprintf("%s_network_%s", nodeName, strings.ReplaceAll(iface, ".", "_")))

	d.Set(mkResourceVirtualEnvironmentNetworkType, network.Type)
	d.Set(mkResourceVirtualEnvironmentContainerNetworkInterface, network.Iface)
	d.Set(mkResourceVirtualEnvironmentNetworkAddress, network.Address)
	d.Set(mkResourceVirtualEnvironmentNetworkCIDR, network.CIDR)
	d.Set(mkResourceVirtualEnvironmentNetworkNetmask, network.Netmask)
	d.Set(mkResourceVirtualEnvironmentNetworkAutostart, network.Autostart)
	d.Set(mkResourceVirtualEnvironmentNetworkBridgePorts, network.BridgePorts)
	d.Set(mkResourceVirtualEnvironmentNetworkGateway, network.Gateway)

	return nil
}

func getBody(d *schema.ResourceData, isUpdate bool) *proxmox.VirtualEnvironmentNetworkInterfaceCreateRequestBody {
	body := &proxmox.VirtualEnvironmentNetworkInterfaceCreateRequestBody{
		Type: d.Get(mkResourceVirtualEnvironmentNetworkType).(string),
	}
	if isUpdate == false {
		body.Iface = d.Get(mkResourceVirtualEnvironmentNetworkIface).(string)
	}

	assignIfStringExists(d, &body.Address, mkResourceVirtualEnvironmentNetworkAddress)
	assignIfStringExists(d, &body.Address6, mkResourceVirtualEnvironmentNetworkAddress6)

	assignIfBoolExists(d, &body.Autostart, mkResourceVirtualEnvironmentNetworkAutostart)

	assignIfStringExists(d, &body.BondPrimary, mkResourceVirtualEnvironmentNetworkBondPrimary)
	assignIfStringExists(d, &body.BondMode, mkResourceVirtualEnvironmentNetworkBondMode)
	assignIfStringExists(d, &body.BondXmitHashPolicy, mkResourceVirtualEnvironmentNetworkBondXmitHashPolicy)
	assignIfStringExists(d, &body.BridgePorts, mkResourceVirtualEnvironmentNetworkBridgePorts)

	assignIfBoolExists(d, &body.BridgeVlanAware, mkResourceVirtualEnvironmentNetworkBridgeVlanAware)

	assignIfStringExists(d, &body.Cidr, mkResourceVirtualEnvironmentNetworkCIDR)
	assignIfStringExists(d, &body.Cidr6, mkResourceVirtualEnvironmentNetworkCIDR6)
	assignIfStringExists(d, &body.Comments, mkResourceVirtualEnvironmentNetworkComments)
	assignIfStringExists(d, &body.Comments6, mkResourceVirtualEnvironmentNetworkComments6)
	assignIfStringExists(d, &body.Gateway, mkResourceVirtualEnvironmentNetworkGateway)
	assignIfStringExists(d, &body.Gateway6, mkResourceVirtualEnvironmentNetworkGateway6)

	assignIfIntExists(d, &body.Mtu, mkResourceVirtualEnvironmentNetworkMtu)

	assignIfStringExists(d, &body.NetMask, mkResourceVirtualEnvironmentNetworkNetmask)
	assignIfStringExists(d, &body.NetMask6, mkResourceVirtualEnvironmentNetworkNetmask6)

	// Ovs Params
	assignIfStringExists(d, &body.OvsBond, mkResourceVirtualEnvironmentNetworkOvsBond)
	assignIfStringExists(d, &body.OvsBridge, mkResourceVirtualEnvironmentNetworkOvsBridge)
	assignIfStringExists(d, &body.OvsOptions, mkResourceVirtualEnvironmentNetworkOvsOptions)
	assignIfStringExists(d, &body.OvsPorts, mkResourceVirtualEnvironmentNetworkOvsPorts)

	assignIfIntExists(d, &body.OvsTags, mkResourceVirtualEnvironmentNetworkOvsTag)
	assignIfStringExists(d, &body.OvsSlaves, mkResourceVirtualEnvironmentNetworkOvsPorts)

	// Vlan
	assignIfIntExists(d, &body.VlanId, mkResourceVirtualEnvironmentNetworkVlanID)
	assignIfStringExists(d, &body.VlanRawDevide, mkResourceVirtualEnvironmentNetworkVlanRawDevice)

	return body
}

func assignIfStringExists(d *schema.ResourceData, p **string, key string) {
	stringVal := d.Get(key).(string)

	if stringVal == "" {
		return
	}

	*p = &stringVal
}

func assignIfBoolExists(d *schema.ResourceData, p **bool, key string) {
	stringVal := d.Get(key).(bool)

	if stringVal == false {
		return
	}

	*p = &stringVal
}

func assignIfIntExists(d *schema.ResourceData, p **int, key string) {
	stringVal := d.Get(key).(int)

	if stringVal == 0 {
		return
	}

	*p = &stringVal
}