# NetworkDeviceFunction

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**OdataContext** | **string** | The OData description of a payload. | [optional] 
**OdataEtag** | **string** | The current ETag of the resource. | [optional] 
**OdataId** | **string** | The unique identifier for a resource. | 
**OdataType** | **string** | The type of a resource. | 
**Actions** | [**Actions2**](Actions_2.md) |  | [optional] 
**AssignablePhysicalPorts** | [**[]IdRef**](idRef.md) | The array of physical port references that this network device function may be assigned to. | [optional] 
**AssignablePhysicalPortsodataCount** | **int32** | The number of items in a collection. | [optional] 
**BootMode** | [**BootMode**](BootMode.md) |  | [optional] 
**Description** | **string** | Provides a description of this resource and is used for commonality  in the schema definitions. | [optional] 
**DeviceEnabled** | **bool** | Whether the network device function is enabled. | [optional] 
**Ethernet** | [**Ethernet**](Ethernet.md) |  | [optional] 
**FibreChannel** | [**FibreChannel**](FibreChannel.md) |  | [optional] 
**Id** | **string** | Uniquely identifies the resource within the collection of like resources. | 
**Links** | [**Links2**](Links_2.md) |  | [optional] 
**MaxVirtualFunctions** | **int32** | The number of virtual functions (VFs) that are available for this Network Device Function. | [optional] 
**Name** | **string** | The name of the resource or array element. | 
**NetDevFuncCapabilities** | [**[]NetworkDeviceTechnology**](NetworkDeviceTechnology.md) | Capabilities of this network device function. | [optional] 
**NetDevFuncType** | [**NetworkDeviceTechnology**](NetworkDeviceTechnology.md) |  | [optional] 
**Oem** | [**Oem**](Oem.md) |  | [optional] 
**PhysicalPortAssignment** | [**IdRef**](idRef.md) |  | [optional] 
**Status** | [**Status**](Status.md) |  | [optional] 
**VirtualFunctionsEnabled** | **bool** | Whether Single Root I/O Virtualization (SR-IOV) Virual Functions (VFs) are enabled for this Network Device Function. | [optional] 
**ISCSIBoot** | [**IScsiBoot**](iSCSIBoot.md) |  | [optional] 

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


