# Storage

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**OdataContext** | **string** | The OData description of a payload. | [optional] 
**OdataEtag** | **string** | The current ETag of the resource. | [optional] 
**OdataId** | **string** | The unique identifier for a resource. | 
**OdataType** | **string** | The type of a resource. | 
**Actions** | [**Actions2**](Actions_2.md) |  | [optional] 
**Description** | **string** | Provides a description of this resource and is used for commonality  in the schema definitions. | [optional] 
**Drives** | [**[]IdRef**](idRef.md) | The set of drives attached to the storage controllers represented by this resource. | [optional] 
**DrivesodataCount** | **int32** | The number of items in a collection. | [optional] 
**Id** | **string** | Uniquely identifies the resource within the collection of like resources. | 
**Links** | [**Links2**](Links_2.md) |  | [optional] 
**Name** | **string** | The name of the resource or array element. | 
**Oem** | [**Oem**](Oem.md) |  | [optional] 
**Redundancy** | [**[]IdRef**](idRef.md) | Redundancy information for the storage subsystem. | [optional] 
**RedundancyodataCount** | **int32** | The number of items in a collection. | [optional] 
**Status** | [**Status**](Status.md) |  | [optional] 
**StorageControllers** | [**[]StorageController**](StorageController.md) | The set of storage controllers represented by this resource. | [optional] 
**StorageControllersodataCount** | **int32** | The number of items in a collection. | [optional] 
**Volumes** | [**IdRef**](idRef.md) |  | [optional] 

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


