# UpdateService

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**OdataContext** | **string** | The OData description of a payload. | [optional] 
**OdataEtag** | **string** | The current ETag of the resource. | [optional] 
**OdataId** | **string** | The unique identifier for a resource. | 
**OdataType** | **string** | The type of a resource. | 
**Actions** | [**Actions2**](Actions_2.md) |  | [optional] 
**Description** | **string** | Provides a description of this resource and is used for commonality  in the schema definitions. | [optional] 
**FirmwareInventory** | [**IdRef**](idRef.md) |  | [optional] 
**HttpPushUri** | **string** | The URI used to perform an HTTP or HTTPS push update to the Update Service. | [optional] 
**HttpPushUriOptions** | [**HttpPushUriOptions**](HttpPushUriOptions.md) |  | [optional] 
**HttpPushUriOptionsBusy** | **bool** | This represents if the properties of HttpPushUriOptions are reserved by any client. | [optional] 
**HttpPushUriTargets** | **[]string** | The array of URIs indicating the target for applying the update image. | [optional] 
**HttpPushUriTargetsBusy** | **bool** | This represents if the HttpPushUriTargets property is reserved by any client. | [optional] 
**Id** | **string** | Uniquely identifies the resource within the collection of like resources. | 
**Name** | **string** | The name of the resource or array element. | 
**Oem** | [**Oem**](Oem.md) |  | [optional] 
**ServiceEnabled** | **bool** | This indicates whether this service is enabled. | [optional] 
**SoftwareInventory** | [**IdRef**](idRef.md) |  | [optional] 
**Status** | [**Status**](Status.md) |  | [optional] 

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


