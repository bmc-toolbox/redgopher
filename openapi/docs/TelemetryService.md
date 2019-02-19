# TelemetryService

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**OdataContext** | **string** | The OData description of a payload. | [optional] 
**OdataEtag** | **string** | The current ETag of the resource. | [optional] 
**OdataId** | **string** | The unique identifier for a resource. | 
**OdataType** | **string** | The type of a resource. | 
**Actions** | [**Actions2**](Actions_2.md) |  | [optional] 
**Description** | **string** | Provides a description of this resource and is used for commonality  in the schema definitions. | [optional] 
**Id** | **string** | Uniquely identifies the resource within the collection of like resources. | 
**LogService** | [**IdRef**](idRef.md) |  | [optional] 
**MaxReports** | **int32** | The maximum number of metric reports supported by this service. | [optional] 
**MetricDefinitions** | [**IdRef**](idRef.md) |  | [optional] 
**MetricReportDefinitions** | [**IdRef**](idRef.md) |  | [optional] 
**MetricReports** | [**IdRef**](idRef.md) |  | [optional] 
**MinCollectionInterval** | **string** | The minimum time interval between collections supported by this service. | [optional] 
**Name** | **string** | The name of the resource or array element. | 
**Oem** | [**Oem**](Oem.md) |  | [optional] 
**Status** | [**Status**](Status.md) |  | [optional] 
**SupportedCollectionFunctions** | [**[]CollectionFunction**](CollectionFunction.md) | The functions that can be performed over each metric. | [optional] 
**Triggers** | [**IdRef**](idRef.md) |  | [optional] 

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


