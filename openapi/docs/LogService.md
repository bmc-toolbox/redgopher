# LogService

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**OdataContext** | **string** | The OData description of a payload. | [optional] 
**OdataEtag** | **string** | The current ETag of the resource. | [optional] 
**OdataId** | **string** | The unique identifier for a resource. | 
**OdataType** | **string** | The type of a resource. | 
**Actions** | [**Actions2**](Actions_2.md) |  | [optional] 
**DateTime** | [**time.Time**](time.Time.md) | The current DateTime (with offset) for the log service, used to set or read time. | [optional] 
**DateTimeLocalOffset** | **string** | The time offset from UTC that the DateTime property is set to in format: +06:00 . | [optional] 
**Description** | **string** | Provides a description of this resource and is used for commonality  in the schema definitions. | [optional] 
**Entries** | [**IdRef**](idRef.md) |  | [optional] 
**Id** | **string** | Uniquely identifies the resource within the collection of like resources. | 
**LogEntryType** | [**LogEntryTypes**](LogEntryTypes.md) |  | [optional] 
**MaxNumberOfRecords** | **int32** | The maximum number of log entries this service can have. | [optional] 
**Name** | **string** | The name of the resource or array element. | 
**Oem** | [**Oem**](Oem.md) |  | [optional] 
**OverWritePolicy** | [**OverWritePolicy**](OverWritePolicy.md) |  | [optional] 
**ServiceEnabled** | **bool** | This indicates whether this service is enabled. | [optional] 
**Status** | [**Status**](Status.md) |  | [optional] 

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


