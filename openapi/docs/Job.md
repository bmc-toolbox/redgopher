# Job

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**OdataId** | **string** | The unique identifier for a resource. | 
**OdataContext** | **string** | The OData description of a payload. | [optional] 
**OdataEtag** | **string** | The current ETag of the resource. | [optional] 
**OdataType** | **string** | The type of a resource. | 
**Actions** | [**Actions2**](Actions_2.md) |  | [optional] 
**CreatedBy** | **string** | The person or program that created this job entry. | [optional] 
**Description** | **string** | Provides a description of this resource and is used for commonality  in the schema definitions. | [optional] 
**EndTime** | [**time.Time**](time.Time.md) | The date-time stamp that the job was completed. | [optional] 
**HidePayload** | **bool** | Indicates that the contents of the Payload should be hidden from view after the Job has been created.  When set to True, the Payload object will not be returned on GET. | [optional] 
**Id** | **string** | Uniquely identifies the resource within the collection of like resources. | 
**JobState** | [**JobState**](JobState.md) |  | [optional] 
**JobStatus** | [**Health**](Health.md) |  | [optional] 
**MaxExecutionTime** | **string** | The maximum amount of time the job is allowed to execute. | [optional] 
**Messages** | [**[]IdRef**](idRef.md) | This is an array of messages associated with the job. | [optional] 
**Name** | **string** | The name of the resource or array element. | 
**Oem** | [**Oem**](Oem.md) |  | [optional] 
**Payload** | [**Payload2**](Payload_2.md) |  | [optional] 
**PercentComplete** | **int32** | The completion percentage of this job. | [optional] 
**Schedule** | [**IdRef**](idRef.md) |  | [optional] 
**StartTime** | [**time.Time**](time.Time.md) | The date-time stamp that the job was started or is scheduled to start. | [optional] 
**StepOrder** | **[]string** | This represents the serialized execution order of the Job Steps. | [optional] 
**Steps** | [**IdRef**](idRef.md) |  | [optional] 

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


