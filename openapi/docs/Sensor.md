# Sensor

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**OdataContext** | **string** | The OData description of a payload. | [optional] 
**OdataEtag** | **string** | The current ETag of the resource. | [optional] 
**OdataId** | **string** | The unique identifier for a resource. | 
**OdataType** | **string** | The type of a resource. | 
**Accuracy** | **float32** | Estimated percent error of measured vs. actual values. | [optional] 
**Actions** | [**Actions2**](Actions_2.md) |  | [optional] 
**AdjustedMaxAllowableOperatingValue** | **float32** | Adjusted maximum allowable operating value for this equipment based on the current environmental conditions present. | [optional] 
**AdjustedMinAllowableOperatingValue** | **float32** | Adjusted minimum allowable operating value for this equipment based on the current environmental conditions present. | [optional] 
**ApparentVA** | **float32** | The product of Voltage and Current for an AC circuit, in Volt-Amperes units. | [optional] 
**Description** | **string** | Provides a description of this resource and is used for commonality  in the schema definitions. | [optional] 
**ElectricalContext** | [**ElectricalContext**](ElectricalContext.md) |  | [optional] 
**Id** | **string** | Uniquely identifies the resource within the collection of like resources. | 
**LoadPercent** | **float32** | The power load utilization for this Sensor. | [optional] 
**Location** | [**Location**](Location.md) |  | [optional] 
**MaxAllowableOperatingValue** | **float32** | Maximum allowable operating value for this equipment. | [optional] 
**MinAllowableOperatingValue** | **float32** | Minimum allowable operating value for this equipment. | [optional] 
**Name** | **string** | The name of the resource or array element. | 
**Oem** | [**Oem**](Oem.md) |  | [optional] 
**PeakReading** | **float32** | The peak reading value for this sensor. | [optional] 
**PeakReadingTime** | [**time.Time**](time.Time.md) | The time at which the Peak Reading value occurred. | [optional] 
**PhysicalContext** | [**IdRef**](idRef.md) |  | [optional] 
**PhysicalSubContext** | [**PhysicalSubContext**](PhysicalSubContext.md) |  | [optional] 
**PowerFactor** | **float32** | The power factor for this Sensor. | [optional] 
**Precision** | **float32** | Number of significant digits in the Reading. | [optional] 
**ReactiveVAR** | **float32** | The square root of the difference term of squared ApparentVA and squared Power (Reading) for a circuit, expressed in VAR units. | [optional] 
**Reading** | **float32** | The present value for this Sensor. | [optional] 
**ReadingRangeMax** | **float32** | The maximum value of Reading possible for this Sensor. | [optional] 
**ReadingRangeMin** | **float32** | The minimum value of Reading possible for this Sensor. | [optional] 
**ReadingType** | [**ReadingType**](ReadingType.md) |  | [optional] 
**ReadingUnits** | **string** | Units in which the reading and thresholds are measured. | [optional] 
**SensingFrequency** | **float32** | The time interval between readings of the physical sensor. | [optional] 
**SensorResetTime** | [**time.Time**](time.Time.md) | The time at which the time-based properties were last reset. | [optional] 
**Status** | [**Status**](Status.md) |  | [optional] 
**Thresholds** | [**Thresholds**](Thresholds.md) |  | [optional] 
**VoltageType** | [**VoltageType**](VoltageType.md) |  | [optional] 

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


