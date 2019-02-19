# Volume2

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**OdataContext** | **string** | The OData description of a payload. | [optional] 
**OdataEtag** | **string** | The current ETag of the resource. | [optional] 
**OdataId** | **string** | The unique identifier for a resource. | 
**OdataType** | **string** | The type of a resource. | 
**Actions** | [**Actions2**](Actions_2.md) |  | [optional] 
**BlockSizeBytes** | **int32** | The size of the smallest addressable unit (Block) of this volume in bytes. | [optional] 
**CapacityBytes** | **int32** | The size in bytes of this Volume. | [optional] 
**Description** | **string** | Provides a description of this resource and is used for commonality  in the schema definitions. | [optional] 
**Encrypted** | **bool** | Is this Volume encrypted. | [optional] 
**EncryptionTypes** | [**[]EncryptionTypes**](EncryptionTypes.md) | The types of encryption used by this Volume. | [optional] 
**Id** | **string** | Uniquely identifies the resource within the collection of like resources. | 
**Identifiers** | [**[]Identifier2**](Identifier_2.md) | The Durable names for the volume. | [optional] 
**Links** | [**Links2**](Links_2.md) |  | [optional] 
**Name** | **string** | The name of the resource or array element. | 
**Oem** | [**Oem**](Oem.md) |  | [optional] 
**Operations** | [**[]Operation**](Operation.md) | The operations currently running on the Volume. | [optional] 
**OptimumIOSizeBytes** | **int32** | The size in bytes of this Volume&#39;s optimum IO size. | [optional] 
**Status** | [**Status**](Status.md) |  | [optional] 
**VolumeType** | [**VolumeType**](VolumeType.md) |  | [optional] 

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


