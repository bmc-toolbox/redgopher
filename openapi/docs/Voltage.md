# Voltage

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**OdataId** | **string** | The unique identifier for a resource. | 
**Actions** | [**VoltageActions**](VoltageActions.md) |  | [optional] 
**LowerThresholdCritical** | **float32** | Below normal range but not yet fatal. | [optional] 
**LowerThresholdFatal** | **float32** | Below normal range and is fatal. | [optional] 
**LowerThresholdNonCritical** | **float32** | Below normal range. | [optional] 
**MaxReadingRange** | **float32** | Maximum value for this Voltage sensor. | [optional] 
**MemberId** | **string** | This is the identifier for the member within the collection. | 
**MinReadingRange** | **float32** | Minimum value for this Voltage sensor. | [optional] 
**Name** | **string** | Voltage sensor name. | [optional] 
**Oem** | [**Oem**](Oem.md) |  | [optional] 
**PhysicalContext** | [**IdRef**](idRef.md) |  | [optional] 
**ReadingVolts** | **float32** | The present reading of the voltage sensor. | [optional] 
**RelatedItem** | [**[]IdRef**](idRef.md) | Describes the areas or devices to which this voltage measurement applies. | [optional] 
**RelatedItemodataCount** | **int32** | The number of items in a collection. | [optional] 
**SensorNumber** | **int32** | A numerical identifier to represent the voltage sensor. | [optional] 
**Status** | [**Status**](Status.md) |  | [optional] 
**UpperThresholdCritical** | **float32** | Above normal range but not yet fatal. | [optional] 
**UpperThresholdFatal** | **float32** | Above normal range and is fatal. | [optional] 
**UpperThresholdNonCritical** | **float32** | Above normal range. | [optional] 

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


