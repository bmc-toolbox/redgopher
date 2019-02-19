# Task

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**OdataContext** | **string** | The OData description of a payload. | [optional] 
**OdataEtag** | **string** | The current ETag of the resource. | [optional] 
**OdataId** | **string** | The unique identifier for a resource. | 
**OdataType** | **string** | The type of a resource. | 
**Actions** | [**Actions2**](Actions_2.md) |  | [optional] 
**Description** | **string** | Provides a description of this resource and is used for commonality  in the schema definitions. | [optional] 
**EndTime** | [**time.Time**](time.Time.md) | The date-time stamp that the task was last completed. | [optional] 
**HidePayload** | **bool** | Indicates that the contents of the Payload should be hidden from view after the Task has been created.  When set to True, the Payload object will not be returned on GET. | [optional] 
**Id** | **string** | Uniquely identifies the resource within the collection of like resources. | 
**Messages** | [**[]IdRef**](idRef.md) | This is an array of messages associated with the task. | [optional] 
**Name** | **string** | The name of the resource or array element. | 
**Oem** | [**Oem**](Oem.md) |  | [optional] 
**Payload** | [**Payload**](Payload.md) |  | [optional] 
**PercentComplete** | **int32** | The completion percentage of this Task. | [optional] 
**StartTime** | [**time.Time**](time.Time.md) | The date-time stamp that the task was last started. | [optional] 
**TaskMonitor** | **string** | The URI of the Task Monitor for this task. | [optional] 
**TaskState** | [**TaskState**](TaskState.md) |  | [optional] 
**TaskStatus** | [**Health**](Health.md) |  | [optional] 

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


