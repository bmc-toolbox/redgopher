/*
 * Redfish API
 *
 * This contains the definition of a Redfish service.
 *
 * API version: 2018.3
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package openapi
type EventType2 string

// List of EventType_2
const (
	EventType2STATUS_CHANGE EventType2 = "StatusChange"
	EventType2RESOURCE_UPDATED EventType2 = "ResourceUpdated"
	EventType2RESOURCE_ADDED EventType2 = "ResourceAdded"
	EventType2RESOURCE_REMOVED EventType2 = "ResourceRemoved"
	EventType2ALERT EventType2 = "Alert"
)