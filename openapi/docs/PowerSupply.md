# PowerSupply

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**OdataId** | **string** | The unique identifier for a resource. | 
**Actions** | [**PowerSupplyActions**](PowerSupplyActions.md) |  | [optional] 
**Assembly** | [**IdRef**](idRef.md) |  | [optional] 
**EfficiencyPercent** | **float32** | The measured efficiency of this Power Supply as a percentage. | [optional] 
**FirmwareVersion** | **string** | The firmware version for this Power Supply. | [optional] 
**HotPluggable** | **bool** | Indicates if this device can be inserted or removed while the equipment is in operation. | [optional] 
**IndicatorLED** | [**IndicatorLed2**](IndicatorLED_2.md) |  | [optional] 
**InputRanges** | [**[]InputRange**](InputRange.md) | This is the input ranges that the power supply can use. | [optional] 
**LastPowerOutputWatts** | **float32** | The average power output of this Power Supply. | [optional] 
**LineInputVoltage** | **float32** | The line input voltage at which the Power Supply is operating. | [optional] 
**LineInputVoltageType** | [**LineInputVoltageType**](LineInputVoltageType.md) |  | [optional] 
**Location** | [**Location**](Location.md) |  | [optional] 
**Manufacturer** | **string** | This is the manufacturer of this power supply. | [optional] 
**MemberId** | **string** | This is the identifier for the member within the collection. | 
**Model** | **string** | The model number for this Power Supply. | [optional] 
**Name** | **string** | The name of the Power Supply. | [optional] 
**Oem** | [**Oem**](Oem.md) |  | [optional] 
**PartNumber** | **string** | The part number for this Power Supply. | [optional] 
**PowerCapacityWatts** | **float32** | The maximum capacity of this Power Supply. | [optional] 
**PowerInputWatts** | **float32** | The measured input power of this Power Supply. | [optional] 
**PowerOutputWatts** | **float32** | The measured output power of this Power Supply. | [optional] 
**PowerSupplyType** | [**PowerSupplyType**](PowerSupplyType.md) |  | [optional] 
**Redundancy** | [**[]IdRef**](idRef.md) | This structure is used to show redundancy for power supplies.  The Component ids will reference the members of the redundancy groups. | [optional] 
**RedundancyodataCount** | **int32** | The number of items in a collection. | [optional] 
**RelatedItem** | [**[]IdRef**](idRef.md) | The ID(s) of the resources associated with this Power Limit. | [optional] 
**RelatedItemodataCount** | **int32** | The number of items in a collection. | [optional] 
**SerialNumber** | **string** | The serial number for this Power Supply. | [optional] 
**SparePartNumber** | **string** | The spare part number for this Power Supply. | [optional] 
**Status** | [**Status**](Status.md) |  | [optional] 

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


