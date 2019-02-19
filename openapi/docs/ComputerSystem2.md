# ComputerSystem2

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**OdataId** | **string** | The unique identifier for a resource. | 
**OdataContext** | **string** | The OData description of a payload. | [optional] 
**OdataEtag** | **string** | The current ETag of the resource. | [optional] 
**OdataType** | **string** | The type of a resource. | 
**Actions** | [**Actions2**](Actions_2.md) |  | [optional] 
**AssetTag** | **string** | The user definable tag that can be used to track this computer system for inventory or other client purposes. | [optional] 
**Bios** | [**IdRef**](idRef.md) |  | [optional] 
**BiosVersion** | **string** | The version of the system BIOS or primary system firmware. | [optional] 
**Boot** | [**Boot**](Boot.md) |  | [optional] 
**Description** | **string** | Provides a description of this resource and is used for commonality  in the schema definitions. | [optional] 
**EthernetInterfaces** | [**IdRef**](idRef.md) |  | [optional] 
**HostName** | **string** | The DNS Host Name, without any domain information. | [optional] 
**HostWatchdogTimer** | [**WatchdogTimer**](WatchdogTimer.md) |  | [optional] 
**HostedServices** | [**HostedServices**](HostedServices.md) |  | [optional] 
**HostingRoles** | [**[]HostingRole**](HostingRole.md) | The hosing roles that this computer system supports. | [optional] 
**Id** | **string** | Uniquely identifies the resource within the collection of like resources. | 
**IndicatorLED** | [**IndicatorLed2**](IndicatorLED_2.md) |  | [optional] 
**Links** | [**Links2**](Links_2.md) |  | [optional] 
**LogServices** | [**IdRef**](idRef.md) |  | [optional] 
**Manufacturer** | **string** | The manufacturer or OEM of this system. | [optional] 
**Memory** | [**IdRef**](idRef.md) |  | [optional] 
**MemoryDomains** | [**IdRef**](idRef.md) |  | [optional] 
**MemorySummary** | [**MemorySummary**](MemorySummary.md) |  | [optional] 
**Model** | **string** | The product name for this system, without the manufacturer name. | [optional] 
**Name** | **string** | The name of the resource or array element. | 
**NetworkInterfaces** | [**IdRef**](idRef.md) |  | [optional] 
**Oem** | [**Oem**](Oem.md) |  | [optional] 
**PCIeDevices** | [**[]IdRef**](idRef.md) | A reference to a collection of PCIe Devices used by this computer system. | [optional] 
**PCIeDevicesodataCount** | **int32** | The number of items in a collection. | [optional] 
**PCIeFunctions** | [**[]IdRef**](idRef.md) | A reference to a collection of PCIe Functions used by this computer system. | [optional] 
**PCIeFunctionsodataCount** | **int32** | The number of items in a collection. | [optional] 
**PartNumber** | **string** | The part number for this system. | [optional] 
**PowerRestorePolicy** | [**PowerRestorePolicyTypes**](PowerRestorePolicyTypes.md) |  | [optional] 
**PowerState** | [**PowerState2**](PowerState_2.md) |  | [optional] 
**ProcessorSummary** | [**ProcessorSummary**](ProcessorSummary.md) |  | [optional] 
**Processors** | [**IdRef**](idRef.md) |  | [optional] 
**Redundancy** | [**[]IdRef**](idRef.md) | A reference to a collection of Redundancy entities that each name a set of computer systems that provide redundancy for this ComputerSystem. | [optional] 
**RedundancyodataCount** | **int32** | The number of items in a collection. | [optional] 
**SKU** | **string** | The manufacturer SKU for this system. | [optional] 
**SecureBoot** | [**IdRef**](idRef.md) |  | [optional] 
**SerialNumber** | **string** | The serial number for this system. | [optional] 
**SimpleStorage** | [**IdRef**](idRef.md) |  | [optional] 
**Status** | [**Status**](Status.md) |  | [optional] 
**Storage** | [**IdRef**](idRef.md) |  | [optional] 
**SubModel** | **string** | The sub-model for this system. | [optional] 
**SystemType** | [**SystemType**](SystemType.md) |  | [optional] 
**TrustedModules** | [**[]TrustedModules**](TrustedModules.md) | This object describes the array of Trusted Modules in the system. | [optional] 
**UUID** | **string** |  | [optional] 

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


