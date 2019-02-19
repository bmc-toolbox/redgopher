# TaskService

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**OdataContext** | **string** | The OData description of a payload. | [optional] 
**OdataEtag** | **string** | The current ETag of the resource. | [optional] 
**OdataId** | **string** | The unique identifier for a resource. | 
**OdataType** | **string** | The type of a resource. | 
**Actions** | [**Actions2**](Actions_2.md) |  | [optional] 
**CompletedTaskOverWritePolicy** | [**OverWritePolicy2**](OverWritePolicy_2.md) |  | [optional] 
**DateTime** | [**time.Time**](time.Time.md) | The current DateTime (with offset) setting that the task service is using. | [optional] 
**Description** | **string** | Provides a description of this resource and is used for commonality  in the schema definitions. | [optional] 
**Id** | **string** | Uniquely identifies the resource within the collection of like resources. | 
**LifeCycleEventOnTaskStateChange** | **bool** | Send an Event upon Task State Change. | [optional] 
**Name** | **string** | The name of the resource or array element. | 
**Oem** | [**Oem**](Oem.md) |  | [optional] 
**ServiceEnabled** | **bool** | This indicates whether this service is enabled. | [optional] 
**Status** | [**Status**](Status.md) |  | [optional] 
**Tasks** | [**IdRef**](idRef.md) |  | [optional] 

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


