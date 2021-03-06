/*
 * Redfish API
 *
 * This contains the definition of a Redfish service.
 *
 * API version: 2018.3
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package openapi

// This type describes Ethernet capabilities, status, and configuration of a network device function.
type Ethernet struct {
	// This is the currently configured MAC address of the (logical port) network device function.
	MACAddress string `json:"MACAddress,omitempty"`
	// The Maximum Transmission Unit (MTU) configured for this network device function.
	MTUSize int32 `json:"MTUSize,omitempty"`
	// This is the permanent MAC address assigned to this network device function (physical function).
	PermanentMACAddress string `json:"PermanentMACAddress,omitempty"`
	VLAN Vlan `json:"VLAN,omitempty"`
	VLANs IdRef `json:"VLANs,omitempty"`
}
