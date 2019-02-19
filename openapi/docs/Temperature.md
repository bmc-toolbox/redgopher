# Temperature

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**OdataId** | **string** | The unique identifier for a resource. | 
**Actions** | [**TemperatureActions**](TemperatureActions.md) |  | [optional] 
**AdjustedMaxAllowableOperatingValue** | **int32** | Adjusted maximum allowable operating temperature for this equipment based on the current environmental conditions present. | [optional] 
**AdjustedMinAllowableOperatingValue** | **int32** | Adjusted minimum allowable operating temperature for this equipment based on the current environmental conditions present. | [optional] 
**DeltaPhysicalContext** | [**IdRef**](idRef.md) |  | [optional] 
**DeltaReadingCelsius** | **float32** | Delta Temperature reading. | [optional] 
**LowerThresholdCritical** | **float32** | Below normal range but not yet fatal. | [optional] 
**LowerThresholdFatal** | **float32** | Below normal range and is fatal. | [optional] 
**LowerThresholdNonCritical** | **float32** | Below normal range. | [optional] 
**MaxAllowableOperatingValue** | **int32** | Maximum allowable operating temperature for this equipment. | [optional] 
**MaxReadingRangeTemp** | **float32** | Maximum value for ReadingCelsius. | [optional] 
**MemberId** | **string** | This is the identifier for the member within the collection. | 
**MinAllowableOperatingValue** | **int32** | Minimum allowable operating temperature for this equipment. | [optional] 
**MinReadingRangeTemp** | **float32** | Minimum value for ReadingCelsius. | [optional] 
**Name** | **string** | Temperature sensor name. | [optional] 
**Oem** | [**Oem**](Oem.md) |  | [optional] 
**PhysicalContext** | [**IdRef**](idRef.md) |  | [optional] 
**ReadingCelsius** | **float32** | Temperature. | [optional] 
**RelatedItem** | [**[]IdRef**](idRef.md) | Describes the areas or devices to which this temperature measurement applies. | [optional] 
**RelatedItemodataCount** | **int32** | The number of items in a collection. | [optional] 
**SensorNumber** | **int32** | A numerical identifier to represent the temperature sensor. | [optional] 
**Status** | [**Status**](Status.md) |  | [optional] 
**UpperThresholdCritical** | **float32** | Above normal range but not yet fatal. | [optional] 
**UpperThresholdFatal** | **float32** | Above normal range and is fatal. | [optional] 
**UpperThresholdNonCritical** | **float32** | Above normal range. | [optional] 

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


