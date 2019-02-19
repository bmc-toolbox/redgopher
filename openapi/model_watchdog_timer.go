/*
 * Redfish API
 *
 * This contains the definition of a Redfish service.
 *
 * API version: 2018.3
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package openapi

// This type describes the Host Watchdog Timer functionality for this system.
type WatchdogTimer struct {
	// This indicates if the Host Watchdog Timer functionality has been enabled. Additional host-based software is necessary to activate the timer function.
	FunctionEnabled bool `json:"FunctionEnabled"`
	Oem string `json:"Oem,omitempty"`
	Status Status `json:"Status,omitempty"`
	TimeoutAction WatchdogTimeoutActions `json:"TimeoutAction"`
	WarningAction WatchdogWarningActions `json:"WarningAction,omitempty"`
}