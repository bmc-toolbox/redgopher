/*
 * Redfish API
 *
 * This contains the definition of a Redfish service.
 *
 * API version: 2018.3
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package openapi

// Information about a Serial Console service provided by a manager.
type SerialConsole struct {
	// This object is used to enumerate the Serial Console connection types allowed by the implementation.
	ConnectTypesSupported []SerialConnectTypesSupported `json:"ConnectTypesSupported,omitempty"`
	// Indicates the maximum number of service sessions, regardless of protocol, this manager is able to support.
	MaxConcurrentSessions int32 `json:"MaxConcurrentSessions,omitempty"`
	// Indicates if the service is enabled for this manager.
	ServiceEnabled bool `json:"ServiceEnabled,omitempty"`
}
