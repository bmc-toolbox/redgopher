# BootOption

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**OdataId** | **string** | The unique identifier for a resource. | [optional] 
**OdataContext** | **string** | The OData description of a payload. | [optional] 
**OdataType** | **string** | The type of a resource. | [optional] 
**Actions** | [**Actions2**](Actions_2.md) |  | [optional] 
**Alias** | [**BootSource**](BootSource.md) |  | [optional] 
**BootOptionEnabled** | **bool** | A flag that shows if the Boot Option is enabled. | [optional] 
**BootOptionReference** | **string** | The unique boot option string that is referenced in the BootOrder. | 
**Description** | **string** | Provides a description of this resource and is used for commonality  in the schema definitions. | [optional] 
**DisplayName** | **string** | The user-readable display string of the Boot Option. | [optional] 
**Id** | **string** | Uniquely identifies the resource within the collection of like resources. | 
**Name** | **string** | The name of the resource or array element. | 
**Oem** | [**Oem**](Oem.md) |  | [optional] 
**RelatedItem** | [**[]IdRef2**](idRef_2.md) | The ID(s) of the resources associated with this Boot Option. | [optional] 
**RelatedItemodataCount** | **float32** | The number of items in a collection. | [optional] 
**UefiDevicePath** | **string** | The UEFI device path used to access this UEFI Boot Option. | [optional] 

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


