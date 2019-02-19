# StorageController

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**OdataId** | **string** | The unique identifier for a resource. | 
**Actions** | [**StorageControllerActions**](StorageControllerActions.md) |  | [optional] 
**Assembly** | [**IdRef**](idRef.md) |  | [optional] 
**AssetTag** | **string** | The user assigned asset tag for this storage controller. | [optional] 
**CacheSummary** | [**CacheSummary**](CacheSummary.md) |  | [optional] 
**FirmwareVersion** | **string** | The firmware version of this storage Controller. | [optional] 
**Identifiers** | [**[]Identifier2**](Identifier_2.md) | The Durable names for the storage controller. | [optional] 
**Links** | [**StorageControllerLinks**](StorageControllerLinks.md) |  | [optional] 
**Location** | [**Location**](Location.md) |  | [optional] 
**Manufacturer** | **string** | This is the manufacturer of this storage controller. | [optional] 
**MemberId** | **string** | This is the identifier for the member within the collection. | 
**Model** | **string** | This is the model number for the storage controller. | [optional] 
**Name** | **string** | The name of the Storage Controller. | [optional] 
**Oem** | [**Oem**](Oem.md) |  | [optional] 
**PCIeInterface** | [**PcIeInterface**](PCIeInterface.md) |  | [optional] 
**PartNumber** | **string** | The part number for this storage controller. | [optional] 
**SKU** | **string** | This is the SKU for this storage controller. | [optional] 
**SerialNumber** | **string** | The serial number for this storage controller. | [optional] 
**SpeedGbps** | **float32** | The maximum speed of the storage controller&#39;s device interface. | [optional] 
**Status** | [**Status**](Status.md) |  | [optional] 
**SupportedControllerProtocols** | [**[]IdRef**](idRef.md) | This represents the protocols by which this storage controller can be communicated to. | [optional] 
**SupportedDeviceProtocols** | [**[]IdRef**](idRef.md) | This represents the protocols which the storage controller can use to communicate with attached devices. | [optional] 
**SupportedRAIDTypes** | [**[]RaidType**](RAIDType.md) | This object describes the RAID Types supported by the storage controller. | [optional] 

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


