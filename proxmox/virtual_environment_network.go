package proxmox

import (
	"errors"
	"fmt"
	"net/url"
	"sort"
)

// ListNodeNetworkDevices retrieves a list of network devices for a specific nodes.
func (c *VirtualEnvironmentClient) ListNodeNetworkDevices(nodeName string) ([]*VirtualEnvironmentNodeNetworkDeviceResponseData, error) {
	resBody := &VirtualEnvironmentNodeNetworkDeviceListResponseBody{}
	err := c.DoRequest(hmGET, fmt.Sprintf("nodes/%s/network", url.PathEscape(nodeName)), nil, resBody)

	if err != nil {
		return nil, err
	}

	if resBody.Data == nil {
		return nil, errors.New("The server did not include a data object in the response")
	}

	sort.Slice(resBody.Data, func(i, j int) bool {
		return resBody.Data[i].Priority < resBody.Data[j].Priority
	})

	return resBody.Data, nil
}

func (c *VirtualEnvironmentClient) GetNetworkInterface(nodeName string, iface string) (*VirtualEnvironmentNodeNetworkDeviceResponseData, error) {
	resBody := &VirtualEnvironmentNodeNetworkGetResponseBody{}
	err := c.DoRequest(hmGET, fmt.Sprintf("nodes/%s/network/%s", url.PathEscape(nodeName), iface), nil, resBody)

	if err != nil {
		return nil, err
	}

	if resBody.Data == nil {
		return nil, errors.New("The server did not include a data object in the response")
	}

	return resBody.Data, nil
}

func (c *VirtualEnvironmentClient) CreateNetworkInterface(nodeName string, d *VirtualEnvironmentNetworkInterfaceCreateRequestBody) error {
	return c.DoRequest(hmPOST, fmt.Sprintf("nodes/%s/network", nodeName), d, nil)
}

func (c *VirtualEnvironmentClient) UpdateNetworkInterface(nodeName string, iface string, d *VirtualEnvironmentNetworkInterfaceCreateRequestBody) error {
	return c.DoRequest(hmPUT, fmt.Sprintf("nodes/%s/network/%s", nodeName, iface), d, nil)
}

func (c *VirtualEnvironmentClient) DeleteNetworkInterface(nodeName string, iface string) error {
	return c.DoRequest(hmDELETE, fmt.Sprintf("nodes/%s/network/%s", nodeName, iface), nil, nil)
}
