/*
 * Redfish API
 *
 * This contains the definition of a Redfish service.
 *
 * API version: 2018.3
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package openapi
type LogEntryTypes string

// List of LogEntryTypes
const (
	LogEntryTypesEVENT LogEntryTypes = "Event"
	LogEntryTypesSEL LogEntryTypes = "SEL"
	LogEntryTypesMULTIPLE LogEntryTypes = "Multiple"
	LogEntryTypesOEM LogEntryTypes = "OEM"
)