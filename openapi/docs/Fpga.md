# Fpga

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ExternalInterfaces** | [**[]FpgaInterface**](FpgaInterface.md) | An array of the FPGA external interfaces. | [optional] 
**FirmwareId** | **string** | The FPGA firmware identifier. | [optional] 
**FirmwareManufacturer** | **string** | The FPGA firmware manufacturer. | [optional] 
**FirmwareVersion** | **string** | The FPGA firmware version. | [optional] 
**FpgaType** | [**FpgaType**](FpgaType.md) |  | [optional] 
**HostInterface** | [**FpgaInterface**](FpgaInterface.md) |  | [optional] 
**Model** | **string** | The FPGA model. | [optional] 
**Oem** | [**Oem**](Oem.md) |  | [optional] 
**PCIeVirtualFunctions** | **int32** | The number of the PCIe Virtual Functions. | [optional] 
**ProgrammableFromHost** | **bool** | This flag indicates if the FPGA firmware can be reprogrammed from the host using system software. | [optional] 
**ReconfigurationSlots** | [**[]FpgaReconfigurationSlot**](FpgaReconfigurationSlot.md) | An array of the FPGA reconfiguration slots.  A reconfiguration slot is used by an FPGA to contain an acceleration function that can change as the FPGA is being provisioned. | [optional] 

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


