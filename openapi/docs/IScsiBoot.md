# IScsiBoot

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AuthenticationMethod** | [**AuthenticationMethod**](AuthenticationMethod.md) |  | [optional] 
**CHAPSecret** | **string** | The shared secret for CHAP authentication. | [optional] 
**CHAPUsername** | **string** | The username for CHAP authentication. | [optional] 
**IPAddressType** | [**IpAddressType**](IPAddressType.md) |  | [optional] 
**IPMaskDNSViaDHCP** | **bool** | Whether the iSCSI boot initiator uses DHCP to obtain the iniator name, IP address, and netmask. | [optional] 
**InitiatorDefaultGateway** | **string** | The IPv6 or IPv4 iSCSI boot default gateway. | [optional] 
**InitiatorIPAddress** | **string** | The IPv6 or IPv4 address of the iSCSI initiator. | [optional] 
**InitiatorName** | **string** | The iSCSI initiator name. | [optional] 
**InitiatorNetmask** | **string** | The IPv6 or IPv4 netmask of the iSCSI boot initiator. | [optional] 
**MutualCHAPSecret** | **string** | The CHAP Secret for 2-way CHAP authentication. | [optional] 
**MutualCHAPUsername** | **string** | The CHAP Username for 2-way CHAP authentication. | [optional] 
**PrimaryDNS** | **string** | The IPv6 or IPv4 address of the primary DNS server for the iSCSI boot initiator. | [optional] 
**PrimaryLUN** | **int32** | The logical unit number (LUN) for the primary iSCSI boot target. | [optional] 
**PrimaryTargetIPAddress** | **string** | The IP address (IPv6 or IPv4) for the primary iSCSI boot target. | [optional] 
**PrimaryTargetName** | **string** | The name of the iSCSI primary boot target. | [optional] 
**PrimaryTargetTCPPort** | **int32** | The TCP port for the primary iSCSI boot target. | [optional] 
**PrimaryVLANEnable** | **bool** | This indicates if the primary VLAN is enabled. | [optional] 
**PrimaryVLANId** | **int32** | The 802.1q VLAN ID to use for iSCSI boot from the primary target. | [optional] 
**RouterAdvertisementEnabled** | **bool** | Whether IPv6 router advertisement is enabled for the iSCSI boot target. | [optional] 
**SecondaryDNS** | **string** | The IPv6 or IPv4 address of the secondary DNS server for the iSCSI boot initiator. | [optional] 
**SecondaryLUN** | **int32** | The logical unit number (LUN) for the secondary iSCSI boot target. | [optional] 
**SecondaryTargetIPAddress** | **string** | The IP address (IPv6 or IPv4) for the secondary iSCSI boot target. | [optional] 
**SecondaryTargetName** | **string** | The name of the iSCSI secondary boot target. | [optional] 
**SecondaryTargetTCPPort** | **int32** | The TCP port for the secondary iSCSI boot target. | [optional] 
**SecondaryVLANEnable** | **bool** | This indicates if the secondary VLAN is enabled. | [optional] 
**SecondaryVLANId** | **int32** | The 802.1q VLAN ID to use for iSCSI boot from the secondary target. | [optional] 
**TargetInfoViaDHCP** | **bool** | Whether the iSCSI boot target name, LUN, IP address, and netmask should be obtained from DHCP. | [optional] 

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


