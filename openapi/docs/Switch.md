# Switch

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**OdataContext** | **string** | The OData description of a payload. | [optional] 
**OdataEtag** | **string** | The current ETag of the resource. | [optional] 
**OdataId** | **string** | The unique identifier for a resource. | 
**OdataType** | **string** | The type of a resource. | 
**Actions** | [**Actions2**](Actions_2.md) |  | [optional] 
**AssetTag** | **string** | The user assigned asset tag for this switch. | [optional] 
**Description** | **string** | Provides a description of this resource and is used for commonality  in the schema definitions. | [optional] 
**DomainID** | **int32** | The Domain ID for this switch. | [optional] 
**Id** | **string** | Uniquely identifies the resource within the collection of like resources. | 
**IndicatorLED** | [**IndicatorLed2**](IndicatorLED_2.md) |  | [optional] 
**IsManaged** | **bool** | This indicates whether the switch is in a managed or unmanaged state. | [optional] 
**Links** | [**Links2**](Links_2.md) |  | [optional] 
**Location** | [**Location**](Location.md) |  | [optional] 
**LogServices** | [**IdRef**](idRef.md) |  | [optional] 
**Manufacturer** | **string** | This is the manufacturer of this switch. | [optional] 
**Model** | **string** | The product model number of this switch. | [optional] 
**Name** | **string** | The name of the resource or array element. | 
**Oem** | [**Oem**](Oem.md) |  | [optional] 
**PartNumber** | **string** | The part number for this switch. | [optional] 
**Ports** | [**IdRef**](idRef.md) |  | [optional] 
**PowerState** | [**PowerState2**](PowerState_2.md) |  | [optional] 
**Redundancy** | [**[]IdRef**](idRef.md) | Redundancy information for the switches. | [optional] 
**RedundancyodataCount** | **int32** | The number of items in a collection. | [optional] 
**SKU** | **string** | This is the SKU for this switch. | [optional] 
**SerialNumber** | **string** | The serial number for this switch. | [optional] 
**Status** | [**Status**](Status.md) |  | [optional] 
**SwitchType** | [**IdRef**](idRef.md) |  | [optional] 
**TotalSwitchWidth** | **int32** | The total number of lanes, phys, or other physical transport links that this switch contains. | [optional] 

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


