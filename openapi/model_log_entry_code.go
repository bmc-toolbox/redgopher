/*
 * Redfish API
 *
 * This contains the definition of a Redfish service.
 *
 * API version: 2018.3
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package openapi
type LogEntryCode string

// List of LogEntryCode
const (
	LogEntryCodeASSERT LogEntryCode = "Assert"
	LogEntryCodeDEASSERT LogEntryCode = "Deassert"
	LogEntryCodeLOWER_NON_CRITICAL___GOING_LOW LogEntryCode = "Lower Non-critical - going low"
	LogEntryCodeLOWER_NON_CRITICAL___GOING_HIGH LogEntryCode = "Lower Non-critical - going high"
	LogEntryCodeLOWER_CRITICAL___GOING_LOW LogEntryCode = "Lower Critical - going low"
	LogEntryCodeLOWER_CRITICAL___GOING_HIGH LogEntryCode = "Lower Critical - going high"
	LogEntryCodeLOWER_NON_RECOVERABLE___GOING_LOW LogEntryCode = "Lower Non-recoverable - going low"
	LogEntryCodeLOWER_NON_RECOVERABLE___GOING_HIGH LogEntryCode = "Lower Non-recoverable - going high"
	LogEntryCodeUPPER_NON_CRITICAL___GOING_LOW LogEntryCode = "Upper Non-critical - going low"
	LogEntryCodeUPPER_NON_CRITICAL___GOING_HIGH LogEntryCode = "Upper Non-critical - going high"
	LogEntryCodeUPPER_CRITICAL___GOING_LOW LogEntryCode = "Upper Critical - going low"
	LogEntryCodeUPPER_CRITICAL___GOING_HIGH LogEntryCode = "Upper Critical - going high"
	LogEntryCodeUPPER_NON_RECOVERABLE___GOING_LOW LogEntryCode = "Upper Non-recoverable - going low"
	LogEntryCodeUPPER_NON_RECOVERABLE___GOING_HIGH LogEntryCode = "Upper Non-recoverable - going high"
	LogEntryCodeTRANSITION_TO_IDLE LogEntryCode = "Transition to Idle"
	LogEntryCodeTRANSITION_TO_ACTIVE LogEntryCode = "Transition to Active"
	LogEntryCodeTRANSITION_TO_BUSY LogEntryCode = "Transition to Busy"
	LogEntryCodeSTATE_DEASSERTED LogEntryCode = "State Deasserted"
	LogEntryCodeSTATE_ASSERTED LogEntryCode = "State Asserted"
	LogEntryCodePREDICTIVE_FAILURE_DEASSERTED LogEntryCode = "Predictive Failure deasserted"
	LogEntryCodePREDICTIVE_FAILURE_ASSERTED LogEntryCode = "Predictive Failure asserted"
	LogEntryCodeLIMIT_NOT_EXCEEDED LogEntryCode = "Limit Not Exceeded"
	LogEntryCodeLIMIT_EXCEEDED LogEntryCode = "Limit Exceeded"
	LogEntryCodePERFORMANCE_MET LogEntryCode = "Performance Met"
	LogEntryCodePERFORMANCE_LAGS LogEntryCode = "Performance Lags"
	LogEntryCodeTRANSITION_TO_OK LogEntryCode = "Transition to OK"
	LogEntryCodeTRANSITION_TO_NON_CRITICAL_FROM_OK LogEntryCode = "Transition to Non-Critical from OK"
	LogEntryCodeTRANSITION_TO_CRITICAL_FROM_LESS_SEVERE LogEntryCode = "Transition to Critical from less severe"
	LogEntryCodeTRANSITION_TO_NON_RECOVERABLE_FROM_LESS_SEVERE LogEntryCode = "Transition to Non-recoverable from less severe"
	LogEntryCodeTRANSITION_TO_NON_CRITICAL_FROM_MORE_SEVERE LogEntryCode = "Transition to Non-Critical from more severe"
	LogEntryCodeTRANSITION_TO_CRITICAL_FROM_NON_RECOVERABLE LogEntryCode = "Transition to Critical from Non-recoverable"
	LogEntryCodeTRANSITION_TO_NON_RECOVERABLE LogEntryCode = "Transition to Non-recoverable"
	LogEntryCodeMONITOR LogEntryCode = "Monitor"
	LogEntryCodeINFORMATIONAL LogEntryCode = "Informational"
	LogEntryCodeDEVICE_REMOVED__DEVICE_ABSENT LogEntryCode = "Device Removed / Device Absent"
	LogEntryCodeDEVICE_INSERTED__DEVICE_PRESENT LogEntryCode = "Device Inserted / Device Present"
	LogEntryCodeDEVICE_DISABLED LogEntryCode = "Device Disabled"
	LogEntryCodeDEVICE_ENABLED LogEntryCode = "Device Enabled"
	LogEntryCodeTRANSITION_TO_RUNNING LogEntryCode = "Transition to Running"
	LogEntryCodeTRANSITION_TO_IN_TEST LogEntryCode = "Transition to In Test"
	LogEntryCodeTRANSITION_TO_POWER_OFF LogEntryCode = "Transition to Power Off"
	LogEntryCodeTRANSITION_TO_ON_LINE LogEntryCode = "Transition to On Line"
	LogEntryCodeTRANSITION_TO_OFF_LINE LogEntryCode = "Transition to Off Line"
	LogEntryCodeTRANSITION_TO_OFF_DUTY LogEntryCode = "Transition to Off Duty"
	LogEntryCodeTRANSITION_TO_DEGRADED LogEntryCode = "Transition to Degraded"
	LogEntryCodeTRANSITION_TO_POWER_SAVE LogEntryCode = "Transition to Power Save"
	LogEntryCodeINSTALL_ERROR LogEntryCode = "Install Error"
	LogEntryCodeFULLY_REDUNDANT LogEntryCode = "Fully Redundant"
	LogEntryCodeREDUNDANCY_LOST LogEntryCode = "Redundancy Lost"
	LogEntryCodeREDUNDANCY_DEGRADED LogEntryCode = "Redundancy Degraded"
	LogEntryCodeNON_REDUNDANTSUFFICIENT_RESOURCES_FROM_REDUNDANT LogEntryCode = "Non-redundant:Sufficient Resources from Redundant"
	LogEntryCodeNON_REDUNDANTSUFFICIENT_RESOURCES_FROM_INSUFFICIENT_RESOURCES LogEntryCode = "Non-redundant:Sufficient Resources from Insufficient Resources"
	LogEntryCodeNON_REDUNDANTINSUFFICIENT_RESOURCES LogEntryCode = "Non-redundant:Insufficient Resources"
	LogEntryCodeREDUNDANCY_DEGRADED_FROM_FULLY_REDUNDANT LogEntryCode = "Redundancy Degraded from Fully Redundant"
	LogEntryCodeREDUNDANCY_DEGRADED_FROM_NON_REDUNDANT LogEntryCode = "Redundancy Degraded from Non-redundant"
	LogEntryCodeD0_POWER_STATE LogEntryCode = "D0 Power State"
	LogEntryCodeD1_POWER_STATE LogEntryCode = "D1 Power State"
	LogEntryCodeD2_POWER_STATE LogEntryCode = "D2 Power State"
	LogEntryCodeD3_POWER_STATE LogEntryCode = "D3 Power State"
	LogEntryCodeOEM LogEntryCode = "OEM"
)