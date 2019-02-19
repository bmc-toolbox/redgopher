# MetricValue2

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**MetricDefinition** | [**IdRef**](idRef.md) |  | [optional] 
**MetricId** | **string** | The metric definitions identifier for this metric. | [optional] 
**MetricProperty** | **string** | The URI for the property from which this metric is derived. | [optional] 
**MetricValue** | **string** | The value identifies this resource. | [optional] 
**Timestamp** | [**time.Time**](time.Time.md) | The time when the value of the metric is obtained. A management application may establish a time series of metric data by retrieving the instances of metric value and sorting them according to their Timestamp. | [optional] 

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


