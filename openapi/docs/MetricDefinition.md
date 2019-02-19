# MetricDefinition

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**OdataId** | **string** | The unique identifier for a resource. | 
**OdataContext** | **string** | The OData description of a payload. | [optional] 
**OdataEtag** | **string** | The current ETag of the resource. | [optional] 
**OdataType** | **string** | The type of a resource. | 
**Accuracy** | **float32** | Estimated percent error of measured vs. actual values. | [optional] 
**Actions** | [**Actions2**](Actions_2.md) |  | [optional] 
**Calculable** | [**Calculable**](Calculable.md) |  | [optional] 
**CalculationAlgorithm** | [**CalculationAlgorithmEnum**](CalculationAlgorithmEnum.md) |  | [optional] 
**CalculationParameters** | [**[]CalculationParamsType**](CalculationParamsType.md) | Specifies the metric properties which are part of the synthesis calculation.  This property is present when the MetricType property has the value &#39;Synthesized&#39;. | [optional] 
**CalculationTimeInterval** | **string** | The time interval over which the metric calculation is performed. | [optional] 
**Calibration** | **float32** | Specifies the calibration offset added to the metric reading. | [optional] 
**Description** | **string** | Provides a description of this resource and is used for commonality  in the schema definitions. | [optional] 
**DiscreteValues** | **[]string** | This array property specifies possible values of a discrete metric. | [optional] 
**Id** | **string** | Uniquely identifies the resource within the collection of like resources. | 
**Implementation** | [**ImplementationType**](ImplementationType.md) |  | [optional] 
**IsLinear** | **bool** | Indicates whether the metric values are linear (vs non-linear). | [optional] 
**MaxReadingRange** | **float32** | Maximum value for metric reading. | [optional] 
**MetricDataType** | [**MetricDataType**](MetricDataType.md) |  | [optional] 
**MetricProperties** | **[]string** | A collection of URI for the properties on which this metric definition is defined. | [optional] 
**MetricType** | [**MetricType**](MetricType.md) |  | [optional] 
**MinReadingRange** | **float32** | Minimum value for metric reading. | [optional] 
**Name** | **string** | The name of the resource or array element. | 
**Oem** | [**Oem**](Oem.md) |  | [optional] 
**PhysicalContext** | [**IdRef**](idRef.md) |  | [optional] 
**Precision** | **int32** | Number of significant digits in the metric reading. | [optional] 
**SensingInterval** | **string** | The time interval between when a metric is updated. | [optional] 
**TimestampAccuracy** | **string** | Accuracy of the timestamp. | [optional] 
**Units** | **string** | The units of measure for this metric. | [optional] 
**Wildcards** | [**[]Wildcard**](Wildcard.md) | Wildcards used to replace values in AppliesTo and Calculates metric property arrays. | [optional] 

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


