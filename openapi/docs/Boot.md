# Boot

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AliasBootOrder** | [**[]BootSource**](BootSource.md) | Ordered array of boot source aliases representing the persistent Boot Order associated with this computer system. | [optional] 
**BootNext** | **string** | This property is the BootOptionReference of the Boot Option to perform a one time boot from when BootSourceOverrideTarget is UefiBootNext. | [optional] 
**BootOptions** | [**IdRef**](idRef.md) |  | [optional] 
**BootOrder** | **[]string** | Ordered array of BootOptionReference strings representing the persistent Boot Order associated with this computer system. | [optional] 
**BootOrderPropertySelection** | [**BootOrderTypes**](BootOrderTypes.md) |  | [optional] 
**BootSourceOverrideEnabled** | [**BootSourceOverrideEnabled**](BootSourceOverrideEnabled.md) |  | [optional] 
**BootSourceOverrideMode** | [**BootSourceOverrideMode**](BootSourceOverrideMode.md) |  | [optional] 
**BootSourceOverrideTarget** | [**BootSource**](BootSource.md) |  | [optional] 
**UefiTargetBootSourceOverride** | **string** | This property is the UEFI Device Path of the device to boot from when BootSourceOverrideTarget is UefiTarget. | [optional] 

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


