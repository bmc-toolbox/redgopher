# Manager

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**OdataContext** | **string** | The OData description of a payload. | [optional] 
**OdataEtag** | **string** | The current ETag of the resource. | [optional] 
**OdataId** | **string** | The unique identifier for a resource. | 
**OdataType** | **string** | The type of a resource. | 
**Actions** | [**Actions2**](Actions_2.md) |  | [optional] 
**AutoDSTEnabled** | **bool** | Indicates whether the manager is configured for automatic DST adjustment. | [optional] 
**CommandShell** | [**CommandShell**](CommandShell.md) |  | [optional] 
**DateTime** | [**time.Time**](time.Time.md) | The current DateTime (with offset) for the manager, used to set or read time. | [optional] 
**DateTimeLocalOffset** | **string** | The time offset from UTC that the DateTime property is set to in format: +06:00 . | [optional] 
**Description** | **string** | Provides a description of this resource and is used for commonality  in the schema definitions. | [optional] 
**EthernetInterfaces** | [**IdRef**](idRef.md) |  | [optional] 
**FirmwareVersion** | **string** | The firmware version of this Manager. | [optional] 
**GraphicalConsole** | [**GraphicalConsole**](GraphicalConsole.md) |  | [optional] 
**HostInterfaces** | [**IdRef**](idRef.md) |  | [optional] 
**Id** | **string** | Uniquely identifies the resource within the collection of like resources. | 
**Links** | [**Links2**](Links_2.md) |  | [optional] 
**LogServices** | [**IdRef**](idRef.md) |  | [optional] 
**ManagerType** | [**ManagerType**](ManagerType.md) |  | [optional] 
**Model** | **string** | The model information of this Manager as defined by the manufacturer. | [optional] 
**Name** | **string** | The name of the resource or array element. | 
**NetworkProtocol** | [**IdRef**](idRef.md) |  | [optional] 
**Oem** | [**Oem**](Oem.md) |  | [optional] 
**PowerState** | [**PowerState2**](PowerState_2.md) |  | [optional] 
**Redundancy** | [**[]IdRef**](idRef.md) | Redundancy information for the managers of this system. | [optional] 
**RedundancyodataCount** | **int32** | The number of items in a collection. | [optional] 
**RemoteAccountService** | [**IdRef**](idRef.md) |  | [optional] 
**RemoteRedfishServiceUri** | **string** | This property contains the URI of the Redfish Service Root for the remote Manager represented by this resource. | [optional] 
**SerialConsole** | [**SerialConsole**](SerialConsole.md) |  | [optional] 
**SerialInterfaces** | [**IdRef**](idRef.md) |  | [optional] 
**ServiceEntryPointUUID** | **string** |  | [optional] 
**Status** | [**Status**](Status.md) |  | [optional] 
**UUID** | **string** |  | [optional] 
**VirtualMedia** | [**IdRef**](idRef.md) |  | [optional] 

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


