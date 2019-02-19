# NetworkPort

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**OdataContext** | **string** | The OData description of a payload. | [optional] 
**OdataEtag** | **string** | The current ETag of the resource. | [optional] 
**OdataId** | **string** | The unique identifier for a resource. | 
**OdataType** | **string** | The type of a resource. | 
**Actions** | [**Actions2**](Actions_2.md) |  | [optional] 
**ActiveLinkTechnology** | [**LinkNetworkTechnology**](LinkNetworkTechnology.md) |  | [optional] 
**AssociatedNetworkAddresses** | **[]string** | The array of configured network addresses (MAC or WWN) that are associated with this Network Port, including the programmed address of the lowest numbered Network Device Function, the configured but not active address if applicable, the address for hardware port teaming, or other network addresses. | [optional] 
**CurrentLinkSpeedMbps** | **int32** | Network Port Current Link Speed. | [optional] 
**Description** | **string** | Provides a description of this resource and is used for commonality  in the schema definitions. | [optional] 
**EEEEnabled** | **bool** | Whether IEEE 802.3az Energy Efficient Ethernet (EEE) is enabled for this network port. | [optional] 
**FCFabricName** | **string** | The FC Fabric Name provided by the switch. | [optional] 
**FCPortConnectionType** | [**PortConnectionType**](PortConnectionType.md) |  | [optional] 
**FlowControlConfiguration** | [**FlowControl**](FlowControl.md) |  | [optional] 
**FlowControlStatus** | [**FlowControl**](FlowControl.md) |  | [optional] 
**Id** | **string** | Uniquely identifies the resource within the collection of like resources. | 
**LinkStatus** | [**LinkStatus**](LinkStatus.md) |  | [optional] 
**MaxFrameSize** | **int32** | The maximum frame size supported by the port. | [optional] 
**Name** | **string** | The name of the resource or array element. | 
**NetDevFuncMaxBWAlloc** | [**[]NetDevFuncMaxBwAlloc**](NetDevFuncMaxBWAlloc.md) | The array of maximum bandwidth allocation percentages for the Network Device Functions associated with this port. | [optional] 
**NetDevFuncMinBWAlloc** | [**[]NetDevFuncMinBwAlloc**](NetDevFuncMinBWAlloc.md) | The array of minimum bandwidth allocation percentages for the Network Device Functions associated with this port. | [optional] 
**NumberDiscoveredRemotePorts** | **int32** | The number of ports not on this adapter that this port has discovered. | [optional] 
**Oem** | [**Oem**](Oem.md) |  | [optional] 
**PhysicalPortNumber** | **string** | The physical port number label for this port. | [optional] 
**PortMaximumMTU** | **int32** | The largest maximum transmission unit (MTU) that can be configured for this network port. | [optional] 
**SignalDetected** | **bool** | Whether or not the port has detected enough signal on enough lanes to establish link. | [optional] 
**Status** | [**Status**](Status.md) |  | [optional] 
**SupportedEthernetCapabilities** | [**[]SupportedEthernetCapabilities**](SupportedEthernetCapabilities.md) | The set of Ethernet capabilities that this port supports. | [optional] 
**SupportedLinkCapabilities** | [**[]SupportedLinkCapabilities**](SupportedLinkCapabilities.md) | The self-described link capabilities of this port. | [optional] 
**VendorId** | **string** | The Vendor Identification for this port. | [optional] 
**WakeOnLANEnabled** | **bool** | Whether Wake on LAN (WoL) is enabled for this network port. | [optional] 

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


