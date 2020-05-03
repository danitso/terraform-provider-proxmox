package proxmox

type NetworkType string

const (
	NETWORK_TYPE_BRIDGE     NetworkType = "bridge"
	NETWORK_TYPE_BOND       NetworkType = "bond"
	NETWORK_TYPE_ETH        NetworkType = "eth"
	NETWORK_TYPE_ALIAS      NetworkType = "alias"
	NETWORK_TYPE_VLAN       NetworkType = "vlan"
	NETWORK_TYPE_OVSBRIDGE  NetworkType = "OVSBridge"
	NETWORK_TYPE_OVSBOND    NetworkType = "OVSBond"
	NETWORK_TYPE_OVSPORT    NetworkType = "OVSPort"
	NETWORK_TYPE_OVSINTPORT NetworkType = "OVSIntPort"
	NETWORK_TYPE_ANYBRIDGE  NetworkType = "any_bridge"
)

// VirtualEnvironmentNodeNetworkDeviceListResponseBody contains the body from a node network device list response.
type VirtualEnvironmentNodeNetworkDeviceListResponseBody struct {
	Data []*VirtualEnvironmentNodeNetworkDeviceResponseData `json:"data,omitempty"`
}

// VirtualEnvironmentNodeNetworkGetResponseBody contains the body from a node network interface response.
type VirtualEnvironmentNodeNetworkGetResponseBody struct {
	Data *VirtualEnvironmentNodeNetworkDeviceResponseData `json:"data,omitempty"`
}

// VirtualEnvironmentNodeNetworkDeviceResponseData contains the data from a node network device list response.
type VirtualEnvironmentNodeNetworkDeviceResponseData struct {
	Active      *CustomBool `json:"active,omitempty"`
	Address     *string     `json:"address,omitempty"`
	Autostart   *CustomBool `json:"autostart,omitempty"`
	BridgeFD    *string     `json:"bridge_fd,omitempty"`
	BridgePorts *string     `json:"bridge_ports,omitempty"`
	BridgeSTP   *string     `json:"bridge_stp,omitempty"`
	CIDR        *string     `json:"cidr,omitempty"`
	Exists      *CustomBool `json:"exists,omitempty"`
	Families    *[]string   `json:"families,omitempty"`
	Gateway     *string     `json:"gateway,omitempty"`
	Iface       string      `json:"iface"`
	MethodIPv4  *string     `json:"method,omitempty"`
	MethodIPv6  *string     `json:"method6,omitempty"`
	Netmask     *string     `json:"netmask,omitempty"`
	Priority    int         `json:"priority"`
	Type        string      `json:"type"`
}

type VirtualEnvironmentNetworkInterfaceCreateRequestBody struct {
	Iface              string  `json:"iface,omitempty" url:"iface,omitempty"`
	Type               string  `json:"type" url:"type"`
	Address            *string `json:"address,omitempty" url:"address,omitempty"`
	Address6           *string `json:"address6,omitempty" url:"address6,omitempty"`
	Autostart          *bool   `json:"autostart,omitempty" url:"autostart,omitempty"`
	BondPrimary        *string `json:"bond-primary,omitempty" url:"bond-primary,omitempty"`
	BondMode           *string `json:"bond_mode,omitempty" url:"bond_mode,omitempty"`
	BondXmitHashPolicy *string `json:"bond_xmit_hash_policy,omitempty" url:"bond_xmit_hash_policy,omitempty"`
	BridgePorts        *string `json:"bridge_ports,omitempty" url:"bridge_ports,omitempty"`
	BridgeVlanAware    *bool   `json:"bridge_vlan_aware,omitempty" url:"bridge_vlan_aware,omitempty"`
	Cidr               *string `json:"cidr,omitempty" url:"cidr,omitempty"`
	Cidr6              *string `json:"cidr6,omitempty" url:"cidr6,omitempty"`
	Comments           *string `json:"comments,omitempty" url:"comments,omitempty"`
	Comments6          *string `json:"comments6,omitempty" url:"comments6,omitempty"`
	Gateway            *string `json:"gateway,omitempty" url:"gateway,omitempty"`
	Gateway6           *string `json:"gateway6,omitempty" url:"gateway6,omitempty"`
	Mtu                *int    `json:"mtu,omitempty" url:"mtu,omitempty"`
	NetMask            *string `json:"netmask,omitempty" url:"netmask,omitempty"`
	NetMask6           *string `json:"netmask6,omitempty" url:"netmask6,omitempty"`
	OvsBond            *string `json:"ovs_bonds,omitempty" url:"ovs_bonds,omitempty"`
	OvsBridge          *string `json:"ovs_bridge,omitempty" url:"ovs_bridge,omitempty"`
	OvsOptions         *string `json:"ovs_options,omitempty" url:"ovs_options,omitempty"`
	OvsPorts           *string `json:"ovs_ports,omitempty" url:"ovs_ports,omitempty"`
	OvsTags            *int    `json:"ovs_tag,omitempty" url:"ovs_tag,omitempty"`
	OvsSlaves          *string `json:"slaves,omitempty" url:"slaves,omitempty"`
	VlanId             *int    `json:"vlan-id,omitempty" url:"vlan-id,omitempty"`
	VlanRawDevide      *string `json:"vlan-raw-device,omitempty" url:"vlan-raw-device,omitempty"`
}
