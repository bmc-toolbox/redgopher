# EventDestination

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**OdataId** | **string** | The unique identifier for a resource. | [optional] 
**OdataContext** | **string** | The OData description of a payload. | [optional] 
**OdataType** | **string** | The type of a resource. | [optional] 
**Context** | **string** | A client-supplied string that is stored with the event destination subscription. | [optional] 
**Description** | **string** | Provides a description of this resource and is used for commonality  in the schema definitions. | [optional] 
**Destination** | **string** | The URI of the destination Event Service. | [optional] 
**EventTypes** | [**[]EventType2**](EventType_2.md) | This property shall contain the types of events that shall be sent to the desination. | [optional] 
**HttpHeaders** | [**[]HttpHeaderProperty**](HttpHeaderProperty.md) | This is for setting HTTP headers, such as authorization information.  This object will be null on a GET. | [optional] 
**Id** | **string** | Uniquely identifies the resource within the collection of like resources. | 
**Name** | **string** | The name of the resource or array element. | 
**Oem** | [**Oem**](Oem.md) |  | [optional] 
**Protocol** | [**EventDestinationProtocol**](EventDestinationProtocol.md) |  | [optional] 

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


