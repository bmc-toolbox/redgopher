/*
 * Redfish API
 *
 * This contains the definition of a Redfish service.
 *
 * API version: 2018.3
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package openapi

// Memory region information within a Memory entity.
type RegionSet struct {
	MemoryClassification MemoryClassification `json:"MemoryClassification,omitempty"`
	// Offset with in the Memory that corresponds to the starting of this memory region in mebibytes (MiB).
	OffsetMiB int32 `json:"OffsetMiB,omitempty"`
	// Indicates if the passphrase is enabled for this region.
	PassphraseEnabled bool `json:"PassphraseEnabled,omitempty"`
	// State of the passphrase for this region.
	PassphraseState bool `json:"PassphraseState,omitempty"`
	// Unique region ID representing a specific region within the Memory.
	RegionId string `json:"RegionId,omitempty"`
	// Size of this memory region in mebibytes (MiB).
	SizeMiB int32 `json:"SizeMiB,omitempty"`
}
