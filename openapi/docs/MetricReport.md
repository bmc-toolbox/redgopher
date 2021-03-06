# MetricReport

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
**MetricReportDefinition** | [**IdRef**](idRef.md) |  | [optional] 
**MetricValues** | [**[]MetricValue**](MetricValue.md) | An array of metric values for the metered items of this Metric. | [optional] 
**Name** | **string** | The name of the resource or array element. | 
**Oem** | [**Oem**](Oem.md) |  | [optional] 
**ReportSequence** | **string** | The current sequence identifier for this metric report. | 
**Timestamp** | [**time.Time**](time.Time.md) | The time associated with the metric report in its entirety.  The time of the metric report may be relevant when the time of individual metrics are minimally different. | [optional] 

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


