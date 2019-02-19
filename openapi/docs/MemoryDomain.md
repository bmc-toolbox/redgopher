# MemoryDomain

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**OdataContext** | **string** | The OData description of a payload. | [optional] 
**OdataEtag** | **string** | The current ETag of the resource. | [optional] 
**OdataId** | **string** | The unique identifier for a resource. | 
**OdataType** | **string** | The type of a resource. | 
**Actions** | [**Actions2**](Actions_2.md) |  | [optional] 
**AllowsBlockProvisioning** | **bool** | Indicates if this Memory Domain supports the provisioning of blocks of memory. | [optional] 
**AllowsMemoryChunkCreation** | **bool** | Indicates if this Memory Domain supports the creation of Memory Chunks. | [optional] 
**AllowsMirroring** | **bool** | Indicates if this Memory Domain supports the creation of Memory Chunks with mirroring enabled. | [optional] 
**AllowsSparing** | **bool** | Indicates if this Memory Domain supports the creation of Memory Chunks with sparing enabled. | [optional] 
**Description** | **string** | Provides a description of this resource and is used for commonality  in the schema definitions. | [optional] 
**Id** | **string** | Uniquely identifies the resource within the collection of like resources. | 
**InterleavableMemorySets** | [**[]MemorySet**](MemorySet.md) | This is the interleave sets for the memory chunk. | [optional] 
**MemoryChunks** | [**IdRef**](idRef.md) |  | [optional] 
**Name** | **string** | The name of the resource or array element. | 
**Oem** | [**Oem**](Oem.md) |  | [optional] 

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


