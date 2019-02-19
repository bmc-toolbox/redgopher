# LogEntry

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**OdataContext** | **string** | The OData description of a payload. | [optional] 
**OdataEtag** | **string** | The current ETag of the resource. | [optional] 
**OdataId** | **string** | The unique identifier for a resource. | 
**OdataType** | **string** | The type of a resource. | 
**Actions** | [**Actions2**](Actions_2.md) |  | [optional] 
**Created** | [**time.Time**](time.Time.md) | The time the log entry was created. | [optional] 
**Description** | **string** | Provides a description of this resource and is used for commonality  in the schema definitions. | [optional] 
**EntryCode** | [**LogEntryCode**](LogEntryCode.md) |  | [optional] 
**EntryType** | [**LogEntryType**](LogEntryType.md) |  | 
**EventGroupId** | **int32** | This value is the identifier used to correlate events that came from the same cause. | [optional] 
**EventId** | **string** | This is a unique instance identifier of an event. | [optional] 
**EventTimestamp** | [**time.Time**](time.Time.md) | This is time the event occurred. | [optional] 
**EventType** | [**EventType**](EventType.md) |  | [optional] 
**Id** | **string** | Uniquely identifies the resource within the collection of like resources. | 
**Links** | [**Links2**](Links_2.md) |  | [optional] 
**Message** | **string** | This property decodes from EntryType.  If the EntryType is Event, then it is a message string.  Otherwise, it is SEL or OEM specific.  In most cases, this will be the actual Log Entry. | [optional] 
**MessageArgs** | **[]string** | The values of this property shall be any arguments for the message. | [optional] 
**MessageId** | **string** | This property decodes from EntryType.  If the EntryType is Event, then it is a message id.  If the EntryType is SEL, then it contains the Event Data.  Otherwise, it is OEM specific.  This value is only used for registries - for more information, see the specification. | [optional] 
**Name** | **string** | The name of the resource or array element. | 
**Oem** | [**Oem**](Oem.md) |  | [optional] 
**OemLogEntryCode** | **string** | If the LogEntryCode type is OEM, this will contain the OEM-specific entry code. | [optional] 
**OemRecordFormat** | **string** | If the EntryType is Oem, this will contain more information about the record format from the Oem. | [optional] 
**OemSensorType** | **string** | If the Sensor Type is OEM, this will contain the OEM-specific sensor type. | [optional] 
**SensorNumber** | **int32** | This property decodes from EntryType.  If the EntryType is SEL, it is the sensor number.  If the EntryType is Event, then the count of events.  Otherwise, it is OEM specific. | [optional] 
**SensorType** | [**SensorType**](SensorType.md) |  | [optional] 
**Severity** | [**EventSeverity**](EventSeverity.md) |  | [optional] 

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


