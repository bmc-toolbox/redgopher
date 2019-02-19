# Certificate

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**OdataId** | **string** | The unique identifier for a resource. | 
**OdataContext** | **string** | The OData description of a payload. | [optional] 
**OdataEtag** | **string** | The current ETag of the resource. | [optional] 
**OdataType** | **string** | The type of a resource. | 
**Actions** | [**Actions2**](Actions_2.md) |  | [optional] 
**CertificateString** | **string** | The string for the certificate. | [optional] 
**CertificateType** | [**CertificateType**](CertificateType.md) |  | [optional] 
**Description** | **string** | Provides a description of this resource and is used for commonality  in the schema definitions. | [optional] 
**Id** | **string** | Uniquely identifies the resource within the collection of like resources. | 
**Issuer** | [**Identifier**](Identifier.md) |  | [optional] 
**KeyUsage** | [**[]KeyUsage**](KeyUsage.md) | The usage of the key contained in the certificate. | [optional] 
**Name** | **string** | The name of the resource or array element. | 
**Oem** | [**Oem**](Oem.md) |  | [optional] 
**Subject** | [**Identifier**](Identifier.md) |  | [optional] 
**ValidNotAfter** | [**time.Time**](time.Time.md) | The date when the certificate is no longer valid. | [optional] 
**ValidNotBefore** | [**time.Time**](time.Time.md) | The date when the certificate becomes valid. | [optional] 

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


