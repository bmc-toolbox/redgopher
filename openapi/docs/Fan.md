# Fan

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**OdataId** | **string** | The unique identifier for a resource. | 
**Actions** | [**FanActions**](FanActions.md) |  | [optional] 
**Assembly** | [**IdRef**](idRef.md) |  | [optional] 
**FanName** | **string** | Name of the fan. | [optional] 
**HotPluggable** | **bool** | Indicates if this device can be inserted or removed while the equipment is in operation. | [optional] 
**IndicatorLED** | [**IndicatorLed2**](IndicatorLED_2.md) |  | [optional] 
**Location** | [**Location**](Location.md) |  | [optional] 
**LowerThresholdCritical** | **int32** | Below normal range but not yet fatal. | [optional] 
**LowerThresholdFatal** | **int32** | Below normal range and is fatal. | [optional] 
**LowerThresholdNonCritical** | **int32** | Below normal range. | [optional] 
**Manufacturer** | **string** | This is the manufacturer of this Fan. | [optional] 
**MaxReadingRange** | **int32** | Maximum value for Reading. | [optional] 
**MemberId** | **string** | This is the identifier for the member within the collection. | 
**MinReadingRange** | **int32** | Minimum value for Reading. | [optional] 
**Model** | **string** | The model number for this Fan. | [optional] 
**Name** | **string** | Name of the fan. | [optional] 
**Oem** | [**Oem**](Oem.md) |  | [optional] 
**PartNumber** | **string** | The part number for this Fan. | [optional] 
**PhysicalContext** | [**IdRef**](idRef.md) |  | [optional] 
**Reading** | **int32** | Current fan speed. | [optional] 
**ReadingUnits** | [**ReadingUnits**](ReadingUnits.md) |  | [optional] 
**Redundancy** | [**[]IdRef**](idRef.md) | This structure is used to show redundancy for fans.  The Component ids will reference the members of the redundancy groups. | [optional] 
**RedundancyodataCount** | **int32** | The number of items in a collection. | [optional] 
**RelatedItem** | [**[]IdRef**](idRef.md) | The ID(s) of the resources serviced with this fan. | [optional] 
**RelatedItemodataCount** | **int32** | The number of items in a collection. | [optional] 
**SensorNumber** | **int32** | A numerical identifier to represent the fan speed sensor. | [optional] 
**SerialNumber** | **string** | The serial number for this Fan. | [optional] 
**SparePartNumber** | **string** | The spare part number for this Fan. | [optional] 
**Status** | [**Status**](Status.md) |  | [optional] 
**UpperThresholdCritical** | **int32** | Above normal range but not yet fatal. | [optional] 
**UpperThresholdFatal** | **int32** | Above normal range and is fatal. | [optional] 
**UpperThresholdNonCritical** | **int32** | Above normal range. | [optional] 

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


