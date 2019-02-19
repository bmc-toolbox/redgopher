# ProcessorMetrics

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**OdataContext** | **string** | The OData description of a payload. | [optional] 
**OdataEtag** | **string** | The current ETag of the resource. | [optional] 
**OdataId** | **string** | The unique identifier for a resource. | 
**OdataType** | **string** | The type of a resource. | 
**Actions** | [**Actions2**](Actions_2.md) |  | [optional] 
**AverageFrequencyMHz** | **float32** | The average frequency of the processor. | [optional] 
**BandwidthPercent** | **float32** | The CPU bandwidth as a percentage. | [optional] 
**Cache** | [**[]CacheMetrics**](CacheMetrics.md) | The processor cache metrics. | [optional] 
**ConsumedPowerWatt** | **float32** | The power consumed by the processor. | [optional] 
**CoreMetrics** | [**[]CoreMetrics**](CoreMetrics.md) | The processor core metrics. | [optional] 
**Description** | **string** | Provides a description of this resource and is used for commonality  in the schema definitions. | [optional] 
**FrequencyRatio** | **float32** | The frequency relative to the nominal processor frequency ratio. | [optional] 
**Id** | **string** | Uniquely identifies the resource within the collection of like resources. | 
**KernelPercent** | **float32** | The percentage of time spent in kernel mode. | [optional] 
**LocalMemoryBandwidthBytes** | **int32** | The local memory bandwidth usage in bytes. | [optional] 
**Name** | **string** | The name of the resource or array element. | 
**Oem** | [**Oem**](Oem.md) |  | [optional] 
**RemoteMemoryBandwidthBytes** | **int32** | The remote memory bandwidth usage in bytes. | [optional] 
**TemperatureCelsius** | **float32** | The temperature of the processor. | [optional] 
**ThrottlingCelsius** | **float32** | The CPU margin to throttle (temperature offset in degree Celsius). | [optional] 
**UserPercent** | **float32** | The percentage of time spent in user mode. | [optional] 

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


