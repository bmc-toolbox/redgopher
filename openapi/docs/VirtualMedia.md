# VirtualMedia

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**OdataContext** | **string** | The OData description of a payload. | [optional] 
**OdataEtag** | **string** | The current ETag of the resource. | [optional] 
**OdataId** | **string** | The unique identifier for a resource. | 
**OdataType** | **string** | The type of a resource. | 
**Actions** | [**Actions2**](Actions_2.md) |  | [optional] 
**ConnectedVia** | [**ConnectedVia**](ConnectedVia.md) |  | [optional] 
**Description** | **string** | Provides a description of this resource and is used for commonality  in the schema definitions. | [optional] 
**Id** | **string** | Uniquely identifies the resource within the collection of like resources. | 
**Image** | **string** | A URI providing the location of the selected image. | [optional] 
**ImageName** | **string** | The current image name. | [optional] 
**Inserted** | **bool** | Indicates if virtual media is inserted in the virtual device. | [optional] 
**MediaTypes** | [**[]MediaType**](MediaType.md) | This is the media types supported as virtual media. | [optional] 
**Name** | **string** | The name of the resource or array element. | 
**Oem** | [**Oem**](Oem.md) |  | [optional] 
**Password** | **string** | The password to be used when accessing the URI specified by the Image parameter. This property is null on a GET. | [optional] 
**TransferMethod** | [**TransferMethod**](TransferMethod.md) |  | [optional] 
**TransferProtocolType** | [**TransferProtocolType**](TransferProtocolType.md) |  | [optional] 
**UserName** | **string** | The username to be used when accessing the URI specified by the Image parameter. | [optional] 
**WriteProtected** | **bool** | Indicates the media is write protected. | [optional] 

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


