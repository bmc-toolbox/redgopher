# CompositionService

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**OdataContext** | **string** | The OData description of a payload. | [optional] 
**OdataEtag** | **string** | The current ETag of the resource. | [optional] 
**OdataId** | **string** | The unique identifier for a resource. | 
**OdataType** | **string** | The type of a resource. | 
**Actions** | [**Actions2**](Actions_2.md) |  | [optional] 
**AllowOverprovisioning** | **bool** | This indicates whether this service is allowed to overprovision a composition relative to the composition request. | [optional] 
**AllowZoneAffinity** | **bool** | This indicates whether a client is allowed to request that given composition request is fulfilled by a specified Resource Zone. | [optional] 
**Description** | **string** | Provides a description of this resource and is used for commonality  in the schema definitions. | [optional] 
**Id** | **string** | Uniquely identifies the resource within the collection of like resources. | 
**Name** | **string** | The name of the resource or array element. | 
**Oem** | [**Oem**](Oem.md) |  | [optional] 
**ResourceBlocks** | [**IdRef**](idRef.md) |  | [optional] 
**ResourceZones** | [**IdRef**](idRef.md) |  | [optional] 
**ServiceEnabled** | **bool** | This indicates whether this service is enabled. | [optional] 
**Status** | [**Status**](Status.md) |  | [optional] 

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


