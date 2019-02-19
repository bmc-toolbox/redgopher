# Thermal

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**OdataContext** | **string** | The OData description of a payload. | [optional] 
**OdataEtag** | **string** | The current ETag of the resource. | [optional] 
**OdataId** | **string** | The unique identifier for a resource. | 
**OdataType** | **string** | The type of a resource. | 
**Actions** | [**ThermalActions**](ThermalActions.md) |  | [optional] 
**Description** | **string** | Provides a description of this resource and is used for commonality  in the schema definitions. | [optional] 
**Fans** | [**[]Fan**](Fan.md) | This is the definition for fans. | [optional] 
**FansodataCount** | **int32** | The number of items in a collection. | [optional] 
**Id** | **string** | Uniquely identifies the resource within the collection of like resources. | 
**Name** | **string** | The name of the resource or array element. | 
**Oem** | [**Oem**](Oem.md) |  | [optional] 
**Redundancy** | [**[]IdRef**](idRef.md) | This structure is used to show redundancy for fans.  The Component ids will reference the members of the redundancy groups. | [optional] 
**RedundancyodataCount** | **int32** | The number of items in a collection. | [optional] 
**Status** | [**Status**](Status.md) |  | [optional] 
**Temperatures** | [**[]Temperature**](Temperature.md) | This is the definition for temperature sensors. | [optional] 
**TemperaturesodataCount** | **int32** | The number of items in a collection. | [optional] 

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


