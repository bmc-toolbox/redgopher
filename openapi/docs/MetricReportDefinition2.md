# MetricReportDefinition2

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**OdataContext** | **string** | The OData description of a payload. | [optional] 
**OdataEtag** | **string** | The current ETag of the resource. | [optional] 
**OdataId** | **string** | The unique identifier for a resource. | 
**OdataType** | **string** | The type of a resource. | 
**Actions** | [**Actions2**](Actions_2.md) |  | [optional] 
**AppendLimit** | **int32** | Indicates the maximum number of entries that can be appended to a metric report.  When the metric report reaches its limit, its behavior is dictated by the ReportUpdates property. | [optional] 
**Description** | **string** | Provides a description of this resource and is used for commonality  in the schema definitions. | [optional] 
**Id** | **string** | Uniquely identifies the resource within the collection of like resources. | 
**MetricProperties** | **[]string** | Specifies a list of metric properties to include in the metric report. | [optional] 
**MetricReport** | [**IdRef**](idRef.md) |  | [optional] 
**MetricReportDefinitionType** | [**MetricReportDefinitionType**](MetricReportDefinitionType.md) |  | [optional] 
**Metrics** | [**[]Metric**](Metric.md) | Specifies a list of metrics to include in the metric report.  The metrics may include metric properties or calculations applied to a metric property. | [optional] 
**Name** | **string** | The name of the resource or array element. | 
**Oem** | [**Oem**](Oem.md) |  | [optional] 
**ReportActions** | [**[]ReportActionsEnum**](ReportActionsEnum.md) | Specifies the actions to perform when a metric report is generated. | [optional] 
**ReportUpdates** | [**ReportUpdatesEnum**](ReportUpdatesEnum.md) |  | [optional] 
**Schedule** | [**IdRef**](idRef.md) |  | [optional] 
**Status** | [**Status**](Status.md) |  | [optional] 
**Wildcards** | [**[]Wildcard**](Wildcard.md) | Specifies the strings used to replace wildcards in the paths in MetricProperties array property. | [optional] 

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


