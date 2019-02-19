# HostInterface

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**OdataContext** | **string** | The OData description of a payload. | [optional] 
**OdataEtag** | **string** | The current ETag of the resource. | [optional] 
**OdataId** | **string** | The unique identifier for a resource. | 
**OdataType** | **string** | The type of a resource. | 
**Actions** | [**Actions2**](Actions_2.md) |  | [optional] 
**AuthNoneRoleId** | **string** | The Role used when no authentication on this interface is used. | [optional] 
**AuthenticationModes** | [**[]AuthenticationMode**](AuthenticationMode.md) | Indicates the authentication modes available on this interface. | [optional] 
**Description** | **string** | Provides a description of this resource and is used for commonality  in the schema definitions. | [optional] 
**ExternallyAccessible** | **bool** | Indicates whether this interface is accessible by external entities. | [optional] 
**FirmwareAuthEnabled** | **bool** | Indicates whether this firmware authentication is enabled for this interface. | [optional] 
**FirmwareAuthRoleId** | **string** | The Role used for firmware authentication on this interface. | [optional] 
**HostEthernetInterfaces** | [**IdRef**](idRef.md) |  | [optional] 
**HostInterfaceType** | [**HostInterfaceType**](HostInterfaceType.md) |  | [optional] 
**Id** | **string** | Uniquely identifies the resource within the collection of like resources. | 
**InterfaceEnabled** | **bool** | Indicates whether this interface is enabled. | [optional] 
**KernelAuthEnabled** | **bool** | Indicates whether this kernel authentication is enabled for this interface. | [optional] 
**KernelAuthRoleId** | **string** | The Role used for kernel authentication on this interface. | [optional] 
**Links** | [**Links2**](Links_2.md) |  | [optional] 
**ManagerEthernetInterface** | [**IdRef**](idRef.md) |  | [optional] 
**Name** | **string** | The name of the resource or array element. | 
**NetworkProtocol** | [**IdRef**](idRef.md) |  | [optional] 
**Oem** | [**Oem**](Oem.md) |  | [optional] 
**Status** | [**Status**](Status.md) |  | [optional] 

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


