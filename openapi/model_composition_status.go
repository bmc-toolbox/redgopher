/*
 * Redfish API
 *
 * This contains the definition of a Redfish service.
 *
 * API version: 2018.3
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package openapi

// Composition status of the Resource Block.
type CompositionStatus struct {
	CompositionState CompositionState `json:"CompositionState"`
	// The maximum number of compositions in which this Resource Block is capable of participating simultaneously.
	MaxCompositions int32 `json:"MaxCompositions,omitempty"`
	// The number of compositions in which this Resource Block is currently participating.
	NumberOfCompositions int32 `json:"NumberOfCompositions,omitempty"`
	// This represents if the Resource Block is reserved by any client.
	Reserved bool `json:"Reserved,omitempty"`
	// Indicates if this Resource Block is capable of participating in multiple compositions simultaneously.
	SharingCapable bool `json:"SharingCapable,omitempty"`
	// Indicates if this Resource Block is allowed to participate in multiple compositions simultaneously.
	SharingEnabled bool `json:"SharingEnabled,omitempty"`
}
