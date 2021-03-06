/*
 * Redfish API
 *
 * This contains the definition of a Redfish service.
 *
 * API version: 2018.3
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package openapi

// This is the schema definition for the Power Metrics.  It represents the properties for Power Consumption and Power Limiting.
type Power struct {
	// The OData description of a payload.
	OdataContext string `json:"@odata.context,omitempty"`
	// The current ETag of the resource.
	OdataEtag string `json:"@odata.etag,omitempty"`
	// The unique identifier for a resource.
	OdataId string `json:"@odata.id"`
	// The type of a resource.
	OdataType string `json:"@odata.type"`
	Actions Actions2 `json:"Actions,omitempty"`
	// Provides a description of this resource and is used for commonality  in the schema definitions.
	Description string `json:"Description,omitempty"`
	// Uniquely identifies the resource within the collection of like resources.
	Id string `json:"Id"`
	// The name of the resource or array element.
	Name string `json:"Name"`
	Oem Oem `json:"Oem,omitempty"`
	// This is the definition for power control function (power reading/limiting).
	PowerControl []PowerControl `json:"PowerControl,omitempty"`
	// The number of items in a collection.
	PowerControlodataCount int32 `json:"PowerControl@odata.count,omitempty"`
	// Details of the power supplies associated with this system or device.
	PowerSupplies []PowerSupply `json:"PowerSupplies,omitempty"`
	// The number of items in a collection.
	PowerSuppliesodataCount int32 `json:"PowerSupplies@odata.count,omitempty"`
	// Redundancy information for the power subsystem of this system or device.
	Redundancy []IdRef `json:"Redundancy,omitempty"`
	// The number of items in a collection.
	RedundancyodataCount int32 `json:"Redundancy@odata.count,omitempty"`
	// This is the definition for voltage sensors.
	Voltages []Voltage `json:"Voltages,omitempty"`
	// The number of items in a collection.
	VoltagesodataCount int32 `json:"Voltages@odata.count,omitempty"`
}
