# Drive

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**OdataContext** | **string** | The OData description of a payload. | [optional] 
**OdataEtag** | **string** | The current ETag of the resource. | [optional] 
**OdataId** | **string** | The unique identifier for a resource. | 
**OdataType** | **string** | The type of a resource. | 
**Actions** | [**Actions2**](Actions_2.md) |  | [optional] 
**Assembly** | [**IdRef**](idRef.md) |  | [optional] 
**AssetTag** | **string** | The user assigned asset tag for this drive. | [optional] 
**BlockSizeBytes** | **int32** | The size of the smallest addressible unit (Block) of this drive in bytes. | [optional] 
**CapableSpeedGbs** | **float32** | The speed which this drive can communicate to a storage controller in ideal conditions in Gigabits per second. | [optional] 
**CapacityBytes** | **int32** | The size in bytes of this Drive. | [optional] 
**Description** | **string** | Provides a description of this resource and is used for commonality  in the schema definitions. | [optional] 
**EncryptionAbility** | [**EncryptionAbility**](EncryptionAbility.md) |  | [optional] 
**EncryptionStatus** | [**EncryptionStatus**](EncryptionStatus.md) |  | [optional] 
**FailurePredicted** | **bool** | Is this drive currently predicting a failure in the near future. | [optional] 
**HotspareReplacementMode** | [**HotspareReplacementModeType**](HotspareReplacementModeType.md) |  | [optional] 
**HotspareType** | [**HotspareType**](HotspareType.md) |  | [optional] 
**Id** | **string** | Uniquely identifies the resource within the collection of like resources. | 
**Identifiers** | [**[]Identifier2**](Identifier_2.md) | The Durable names for the drive. | [optional] 
**IndicatorLED** | [**IndicatorLed2**](IndicatorLED_2.md) |  | [optional] 
**Links** | [**Links2**](Links_2.md) |  | [optional] 
**Location** | [**[]Location**](Location.md) | The Location of the drive. | [optional] 
**Manufacturer** | **string** | This is the manufacturer of this drive. | [optional] 
**MediaType** | [**MediaType**](MediaType.md) |  | [optional] 
**Model** | **string** | This is the model number for the drive. | [optional] 
**Name** | **string** | The name of the resource or array element. | 
**NegotiatedSpeedGbs** | **float32** | The speed which this drive is currently communicating to the storage controller in Gigabits per second. | [optional] 
**Oem** | [**Oem**](Oem.md) |  | [optional] 
**Operations** | [**[]Operations**](Operations.md) | The operations currently running on the Drive. | [optional] 
**PartNumber** | **string** | The part number for this drive. | [optional] 
**PhysicalLocation** | [**Location**](Location.md) |  | [optional] 
**PredictedMediaLifeLeftPercent** | **float32** | The percentage of reads and writes that are predicted to still be available for the media. | [optional] 
**Protocol** | [**IdRef**](idRef.md) |  | [optional] 
**Revision** | **string** | The revision of this Drive. This is typically the firmware/hardware version of the drive. | [optional] 
**RotationSpeedRPM** | **float32** | The rotation speed of this Drive in Revolutions per Minute (RPM). | [optional] 
**SKU** | **string** | This is the SKU for this drive. | [optional] 
**SerialNumber** | **string** | The serial number for this drive. | [optional] 
**Status** | [**Status**](Status.md) |  | [optional] 
**StatusIndicator** | [**StatusIndicator**](StatusIndicator.md) |  | [optional] 

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


