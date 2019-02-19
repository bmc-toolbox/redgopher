# SoftwareInventory

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**OdataContext** | **string** | The OData description of a payload. | [optional] 
**OdataEtag** | **string** | The current ETag of the resource. | [optional] 
**OdataId** | **string** | The unique identifier for a resource. | 
**OdataType** | **string** | The type of a resource. | 
**Actions** | [**Actions2**](Actions_2.md) |  | [optional] 
**Description** | **string** | Provides a description of this resource and is used for commonality  in the schema definitions. | [optional] 
**Id** | **string** | Uniquely identifies the resource within the collection of like resources. | 
**LowestSupportedVersion** | **string** | A string representing the lowest supported version of this software. | [optional] 
**Manufacturer** | **string** | A string representing the manufacturer/producer of this software. | [optional] 
**Name** | **string** | The name of the resource or array element. | 
**Oem** | [**Oem**](Oem.md) |  | [optional] 
**RelatedItem** | [**[]IdRef**](idRef.md) | The ID(s) of the resources associated with this software inventory item. | [optional] 
**RelatedItemodataCount** | **int32** | The number of items in a collection. | [optional] 
**ReleaseDate** | [**time.Time**](time.Time.md) | Release date of this software. | [optional] 
**SoftwareId** | **string** | A string representing the implementation-specific ID for identifying this software. | [optional] 
**Status** | [**Status**](Status.md) |  | [optional] 
**UefiDevicePaths** | **[]string** | A list of strings representing the UEFI Device Path(s) of the component(s) associated with this software inventory item. | [optional] 
**Updateable** | **bool** | Indicates whether this software can be updated by the update service. | [optional] 
**Version** | **string** | A string representing the version of this software. | [optional] 

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


