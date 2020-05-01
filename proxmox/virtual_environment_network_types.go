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
	Node      string      `json:"node"`
	Iface     string      `json:"iface"`
	Type      NetworkType `json:"type"`
	Address   *string     `json:"address,omitempty"`
	Address6  *string     `json:"address6,omitempty"`
	Autostart *bool       `json:"autostart,omitempty"`
}
