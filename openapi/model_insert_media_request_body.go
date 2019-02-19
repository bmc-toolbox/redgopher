/*
 * Redfish API
 *
 * This contains the definition of a Redfish service.
 *
 * API version: 2018.3
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package openapi

// This action is used to attach remote media to virtual media.
type InsertMediaRequestBody struct {
	// The URI of the remote media to attach to the virtual media.
	Image string `json:"Image"`
	// Indicates if the image is to be treated as inserted upon completion of the action.
	Inserted bool `json:"Inserted,omitempty"`
	// The password to be used when accessing the URI specified by the Image parameter.
	Password string `json:"Password,omitempty"`
	TransferMethod TransferMethod `json:"TransferMethod,omitempty"`
	TransferProtocolType TransferProtocolType `json:"TransferProtocolType,omitempty"`
	// The username to be used when accessing the URI specified by the Image parameter.
	UserName string `json:"UserName,omitempty"`
	// Indicates if the remote media is supposed to be treated as write protected.
	WriteProtected bool `json:"WriteProtected,omitempty"`
}