/*
 * Redfish API
 *
 * This contains the definition of a Redfish service.
 *
 * API version: 2018.3
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package openapi

// This object contains the Memory metrics since last reset or ClearCurrentPeriod action.
type CurrentPeriod struct {
	// Number of blocks read since reset.
	BlocksRead int32 `json:"BlocksRead,omitempty"`
	// Number of blocks written since reset.
	BlocksWritten int32 `json:"BlocksWritten,omitempty"`
}