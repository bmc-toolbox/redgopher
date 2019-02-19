# PowerControl

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**OdataId** | **string** | The unique identifier for a resource. | 
**Actions** | [**PowerControlActions**](PowerControlActions.md) |  | [optional] 
**MemberId** | **string** | This is the identifier for the member within the collection. | 
**Name** | **string** | Power Control Function name. | [optional] 
**Oem** | [**Oem**](Oem.md) |  | [optional] 
**PhysicalContext** | [**IdRef**](idRef.md) |  | [optional] 
**PowerAllocatedWatts** | **float32** | The total amount of power that has been allocated (or budegeted)to  chassis resources. | [optional] 
**PowerAvailableWatts** | **float32** | The amount of power not already budgeted and therefore available for additional allocation. (powerCapacity - powerAllocated).  This indicates how much reserve power capacity is left. | [optional] 
**PowerCapacityWatts** | **float32** | The total amount of power available to the chassis for allocation. This may the power supply capacity, or power budget assigned to the chassis from an up-stream chassis. | [optional] 
**PowerConsumedWatts** | **float32** | The actual power being consumed by the chassis. | [optional] 
**PowerLimit** | [**PowerLimit**](PowerLimit.md) |  | [optional] 
**PowerMetrics** | [**PowerMetric**](PowerMetric.md) |  | [optional] 
**PowerRequestedWatts** | **float32** | The potential power that the chassis resources are requesting which may be higher than the current level being consumed since requested power includes budget that the chassis resource wants for future use. | [optional] 
**RelatedItem** | [**[]IdRef**](idRef.md) | The ID(s) of the resources associated with this Power Limit. | [optional] 
**RelatedItemodataCount** | **int32** | The number of items in a collection. | [optional] 
**Status** | [**Status**](Status.md) |  | [optional] 

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


