/*
 * Redfish API
 *
 * This contains the definition of a Redfish service.
 *
 * API version: 2018.3
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package openapi
// HostingRole : The enumerations of HostingRoles specify different features that the hosting ComputerSystem supports.
type HostingRole string

// List of HostingRole
const (
	HostingRoleAPPLICATION_SERVER HostingRole = "ApplicationServer"
	HostingRoleSTORAGE_SERVER HostingRole = "StorageServer"
	HostingRoleSWITCH HostingRole = "Switch"
)