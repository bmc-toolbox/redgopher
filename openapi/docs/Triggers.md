# Triggers

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**OdataId** | **string** | The unique identifier for a resource. | 
**OdataContext** | **string** | The OData description of a payload. | [optional] 
**OdataEtag** | **string** | The current ETag of the resource. | [optional] 
**OdataType** | **string** | The type of a resource. | 
**Actions** | [**Actions2**](Actions_2.md) |  | [optional] 
**Description** | **string** | Provides a description of this resource and is used for commonality  in the schema definitions. | [optional] 
**DiscreteTriggerCondition** | [**DiscreteTriggerConditionEnum**](DiscreteTriggerConditionEnum.md) |  | [optional] 
**DiscreteTriggers** | [**[]DiscreteTrigger**](DiscreteTrigger.md) | List of discrete triggers. | [optional] 
**Id** | **string** | Uniquely identifies the resource within the collection of like resources. | 
**MetricProperties** | **[]string** | A collection of URI for the properties on which this metric definition is defined. | [optional] 
**MetricType** | [**MetricTypeEnum**](MetricTypeEnum.md) |  | [optional] 
**Name** | **string** | The name of the resource or array element. | 
**NumericThresholds** | [**Thresholds2**](Thresholds_2.md) |  | [optional] 
**Oem** | [**Oem**](Oem.md) |  | [optional] 
**Status** | [**Status**](Status.md) |  | [optional] 
**TriggerActions** | [**[]TriggerActionEnum**](TriggerActionEnum.md) | This property specifies the actions to perform when the trigger occurs. | [optional] 
**Wildcards** | [**[]Wildcard**](Wildcard.md) | Wildcards used to replace values in MetricProperties array property. | [optional] 

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


