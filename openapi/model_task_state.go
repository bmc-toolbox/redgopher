/*
 * Redfish API
 *
 * This contains the definition of a Redfish service.
 *
 * API version: 2018.3
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package openapi
type TaskState string

// List of TaskState
const (
	TaskStateNEW TaskState = "New"
	TaskStateSTARTING TaskState = "Starting"
	TaskStateRUNNING TaskState = "Running"
	TaskStateSUSPENDED TaskState = "Suspended"
	TaskStateINTERRUPTED TaskState = "Interrupted"
	TaskStatePENDING TaskState = "Pending"
	TaskStateSTOPPING TaskState = "Stopping"
	TaskStateCOMPLETED TaskState = "Completed"
	TaskStateKILLED TaskState = "Killed"
	TaskStateEXCEPTION TaskState = "Exception"
	TaskStateSERVICE TaskState = "Service"
	TaskStateCANCELLING TaskState = "Cancelling"
	TaskStateCANCELLED TaskState = "Cancelled"
)