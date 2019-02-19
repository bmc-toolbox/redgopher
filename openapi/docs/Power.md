# Power

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**OdataContext** | **string** | The OData description of a payload. | [optional] 
**OdataEtag** | **string** | The current ETag of the resource. | [optional] 
**OdataId** | **string** | The unique identifier for a resource. | 
**OdataType** | **string** | The type of a resource. | 
**Actions** | [**Actions2**](Actions_2.md) |  | [optional] 
**Description** | **string** | Provides a description of this resource and is used for commonality  in the schema definitions. | [optional] 
**Id** | **string** | Uniquely identifies the resource within the collection of like resources. | 
**Name** | **string** | The name of the resource or array element. | 
**Oem** | [**Oem**](Oem.md) |  | [optional] 
**PowerControl** | [**[]PowerControl**](PowerControl.md) | This is the definition for power control function (power reading/limiting). | [optional] 
**PowerControlodataCount** | **int32** | The number of items in a collection. | [optional] 
**PowerSupplies** | [**[]PowerSupply**](PowerSupply.md) | Details of the power supplies associated with this system or device. | [optional] 
**PowerSuppliesodataCount** | **int32** | The number of items in a collection. | [optional] 
**Redundancy** | [**[]IdRef**](idRef.md) | Redundancy information for the power subsystem of this system or device. | [optional] 
**RedundancyodataCount** | **int32** | The number of items in a collection. | [optional] 
**Voltages** | [**[]Voltage**](Voltage.md) | This is the definition for voltage sensors. | [optional] 
**VoltagesodataCount** | **int32** | The number of items in a collection. | [optional] 

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


