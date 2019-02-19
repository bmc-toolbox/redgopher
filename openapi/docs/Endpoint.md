# Endpoint

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**OdataContext** | **string** | The OData description of a payload. | [optional] 
**OdataEtag** | **string** | The current ETag of the resource. | [optional] 
**OdataId** | **string** | The unique identifier for a resource. | 
**OdataType** | **string** | The type of a resource. | 
**Actions** | [**Actions2**](Actions_2.md) |  | [optional] 
**ConnectedEntities** | [**[]ConnectedEntity**](ConnectedEntity.md) | All the entities connected to this endpoint. | [optional] 
**Description** | **string** | Provides a description of this resource and is used for commonality  in the schema definitions. | [optional] 
**EndpointProtocol** | [**IdRef**](idRef.md) |  | [optional] 
**HostReservationMemoryBytes** | **int32** | The amount of memory in Bytes that the Host should allocate to connect to this endpoint. | [optional] 
**IPTransportDetails** | [**[]IpTransportDetails**](IPTransportDetails.md) | This array contains details for each IP transport supported by this endpoint. The array structure can be used to model multiple IP addresses for this endpoint. | [optional] 
**Id** | **string** | Uniquely identifies the resource within the collection of like resources. | 
**Identifiers** | [**[]Identifier2**](Identifier_2.md) | Identifiers for this endpoint. | [optional] 
**Links** | [**Links2**](Links_2.md) |  | [optional] 
**Name** | **string** | The name of the resource or array element. | 
**Oem** | [**Oem**](Oem.md) |  | [optional] 
**PciId** | [**PciId**](PciId.md) |  | [optional] 
**Redundancy** | [**[]IdRef**](idRef.md) | Redundancy information for the lower level endpoints supporting this endpoint. | [optional] 
**RedundancyodataCount** | **int32** | The number of items in a collection. | [optional] 
**Status** | [**Status**](Status.md) |  | [optional] 

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


