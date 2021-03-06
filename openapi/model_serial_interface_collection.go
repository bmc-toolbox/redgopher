/*
 * Redfish API
 *
 * This contains the definition of a Redfish service.
 *
 * API version: 2018.3
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package openapi

// A Collection of SerialInterface resource instances.
type SerialInterfaceCollection struct {
	// The OData description of a payload.
	OdataContext string `json:"@odata.context,omitempty"`
	// The current ETag of the resource.
	OdataEtag string `json:"@odata.etag,omitempty"`
	// The unique identifier for a resource.
	OdataId string `json:"@odata.id"`
	// The type of a resource.
	OdataType string `json:"@odata.type"`
	// Provides a description of this resource and is used for commonality  in the schema definitions.
	Description string `json:"Description,omitempty"`
	// Contains the members of this collection.
	Members []IdRef `json:"Members"`
	// The number of items in a collection.
	MembersodataCount int32 `json:"Members@odata.count,omitempty"`
	// The URI to the resource containing the next set of partial members.
	MembersodataNextLink string `json:"Members@odata.nextLink,omitempty"`
	// The name of the resource or array element.
	Name string `json:"Name"`
	Oem Oem `json:"Oem,omitempty"`
}
