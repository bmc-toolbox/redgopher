/*
 * Redfish API
 *
 * This contains the definition of a Redfish service.
 *
 * API version: 2018.3
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package openapi
type BootMode string

// List of BootMode
const (
	BootModeDISABLED BootMode = "Disabled"
	BootModePXE BootMode = "PXE"
	BootModeI_SCSI BootMode = "iSCSI"
	BootModeFIBRE_CHANNEL BootMode = "FibreChannel"
	BootModeFIBRE_CHANNEL_OVER_ETHERNET BootMode = "FibreChannelOverEthernet"
)