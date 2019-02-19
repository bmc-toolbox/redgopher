# EthernetInterface

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**OdataContext** | **string** | The OData description of a payload. | [optional] 
**OdataEtag** | **string** | The current ETag of the resource. | [optional] 
**OdataId** | **string** | The unique identifier for a resource. | 
**OdataType** | **string** | The type of a resource. | 
**Actions** | [**Actions2**](Actions_2.md) |  | [optional] 
**AutoNeg** | **bool** | This indicates if the speed and duplex are automatically negotiated and configured on this interface. | [optional] 
**DHCPv4** | [**DhcPv4Configuration**](DHCPv4Configuration.md) |  | [optional] 
**DHCPv6** | [**DhcPv6Configuration**](DHCPv6Configuration.md) |  | [optional] 
**Description** | **string** | Provides a description of this resource and is used for commonality  in the schema definitions. | [optional] 
**FQDN** | **string** | This is the complete, fully qualified domain name obtained by DNS for this interface. | [optional] 
**FullDuplex** | **bool** | This indicates if the interface is in Full Duplex mode or not. | [optional] 
**HostName** | **string** | The DNS Host Name, without any domain information. | [optional] 
**IPv4Addresses** | [**[]IPv4Address**](IPv4Address.md) | The IPv4 addresses currently assigned to this interface. | [optional] 
**IPv4StaticAddresses** | [**[]IPv4Address**](IPv4Address.md) | The IPv4 static addresses assigned to this interface. | [optional] 
**IPv6AddressPolicyTable** | [**[]IPv6AddressPolicyEntry**](IPv6AddressPolicyEntry.md) | An array representing the RFC 6724 Address Selection Policy Table. | [optional] 
**IPv6Addresses** | [**[]IPv6Address**](IPv6Address.md) | Enumerates in an array all of the currently assigned IPv6 addresses on this interface. | [optional] 
**IPv6DefaultGateway** | **string** | This is the IPv6 default gateway address that is currently in use on this interface. | [optional] 
**IPv6StaticAddresses** | [**[]IPv6StaticAddress**](IPv6StaticAddress.md) | Represents in an array all of the IPv6 static addresses to be assigned on this interface. | [optional] 
**IPv6StaticDefaultGateways** | [**[]IPv6GatewayStaticAddress**](IPv6GatewayStaticAddress.md) | The IPv6 static default gateways for this interface. | [optional] 
**Id** | **string** | Uniquely identifies the resource within the collection of like resources. | 
**InterfaceEnabled** | **bool** | This indicates whether this interface is enabled. | [optional] 
**LinkStatus** | [**LinkStatus2**](LinkStatus_2.md) |  | [optional] 
**Links** | [**Links2**](Links_2.md) |  | [optional] 
**MACAddress** | **string** |  | [optional] 
**MTUSize** | **int32** | This is the currently configured Maximum Transmission Unit (MTU) in bytes on this interface. | [optional] 
**MaxIPv6StaticAddresses** | **int32** | This indicates the maximum number of Static IPv6 addresses that can be configured on this interface. | [optional] 
**Name** | **string** | The name of the resource or array element. | 
**NameServers** | **[]string** | This represents DNS name servers that are currently in use on this interface. | [optional] 
**Oem** | [**Oem**](Oem.md) |  | [optional] 
**PermanentMACAddress** | **string** |  | [optional] 
**SpeedMbps** | **int32** | This is the current speed in Mbps of this interface. | [optional] 
**StatelessAddressAutoConfig** | [**StatelessAddressAutoConfiguration**](StatelessAddressAutoConfiguration.md) |  | [optional] 
**StaticNameServers** | **[]string** | A statically defined set of DNS server IP addresses (both IPv4 and IPv6). | [optional] 
**Status** | [**Status**](Status.md) |  | [optional] 
**UefiDevicePath** | **string** | The UEFI device path for this interface. | [optional] 
**VLAN** | [**Vlan**](VLAN.md) |  | [optional] 
**VLANs** | [**IdRef**](idRef.md) |  | [optional] 

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


