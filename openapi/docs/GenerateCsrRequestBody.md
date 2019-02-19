# GenerateCsrRequestBody

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AlternativeNames** | **[]string** | Additional hostnames of the component that is being secured. | [optional] 
**CertificateCollection** | [**IdRef**](idRef.md) |  | 
**ChallengePassword** | **string** | The challenge password to be applied to the certificate for revocation requests. | [optional] 
**City** | **string** | The city or locality of the organization making the request. | 
**CommonName** | **string** | The fully qualified domain name of the component that is being secured. | 
**ContactPerson** | **string** | The name of the user making the request. | [optional] 
**Country** | **string** | The country of the organization making the request. | 
**Email** | **string** | The email address of the contact within the organization making the request. | [optional] 
**GivenName** | **string** | The given name of the user making the request. | [optional] 
**Initials** | **string** | The initials of the user making the request. | [optional] 
**KeyBitLength** | **int32** | The length of the key in bits, if needed based on the value of the &#39;KeyPairAlgorithm&#39; parameter. | [optional] 
**KeyCurveId** | **string** | The curve ID to be used with the key, if needed based on the value of the &#39;KeyPairAlgorithm&#39; parameter. | [optional] 
**KeyPairAlgorithm** | **string** | The type of key pair for use with signing algorithms. | [optional] 
**KeyUsage** | [**[]KeyUsage**](KeyUsage.md) | The usage of the key contained in the certificate. | [optional] 
**Organization** | **string** | The name of the organization making the request. | 
**OrganizationalUnit** | **string** | The name of the unit or division of the organization making the request. | 
**State** | **string** | The state, province, or region of the organization making the request. | 
**Surname** | **string** | The surname of the user making the request. | [optional] 
**UnstructuredName** | **string** | The unstructured name of the subject. | [optional] 

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


