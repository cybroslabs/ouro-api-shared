# ApiService - Devices

## CreateCommunicationUnit

The method called by the RestAPI to register a new communication unit. The parameter contains the communication unit specification.

```proto
CreateCommunicationUnit(io.clbs.openhes.models.acquisition.CreateCommunicationUnitRequest) returns (google.protobuf.StringValue)
```

- Input: [io.clbs.openhes.models.acquisition.CreateCommunicationUnitRequest](model-io-clbs-openhes-models-acquisition-createcommunicationunitrequest.md)
- Output: google.protobuf.StringValue

## ListCommunicationUnits

The method called by the RestAPI to get the information about the communication unit. The parameter contains the search criteria.

```proto
ListCommunicationUnits(io.clbs.openhes.models.common.ListSelector) returns (io.clbs.openhes.models.acquisition.ListOfCommunicationUnit)
```

- Input: [io.clbs.openhes.models.common.ListSelector](model-io-clbs-openhes-models-common-listselector.md)
- Output: [io.clbs.openhes.models.acquisition.ListOfCommunicationUnit](model-io-clbs-openhes-models-acquisition-listofcommunicationunit.md)

## GetCommunicationUnit

The method called by the RestAPI to get the information about the communication unit. The parameter contains the search criteria.

```proto
GetCommunicationUnit(google.protobuf.StringValue) returns (io.clbs.openhes.models.acquisition.CommunicationUnit)
```

- Input: google.protobuf.StringValue
- Output: [io.clbs.openhes.models.acquisition.CommunicationUnit](model-io-clbs-openhes-models-acquisition-communicationunit.md)

## CreateCommunicationBus

```proto
CreateCommunicationBus(io.clbs.openhes.models.acquisition.CreateCommunicationBusRequest) returns (google.protobuf.StringValue)
```

- Input: [io.clbs.openhes.models.acquisition.CreateCommunicationBusRequest](model-io-clbs-openhes-models-acquisition-createcommunicationbusrequest.md)
- Output: google.protobuf.StringValue

## ListCommunicationBuses

```proto
ListCommunicationBuses(io.clbs.openhes.models.common.ListSelector) returns (io.clbs.openhes.models.acquisition.ListOfCommunicationBus)
```

- Input: [io.clbs.openhes.models.common.ListSelector](model-io-clbs-openhes-models-common-listselector.md)
- Output: [io.clbs.openhes.models.acquisition.ListOfCommunicationBus](model-io-clbs-openhes-models-acquisition-listofcommunicationbus.md)

## AddCommunicationUnitsToCommunicationBus

```proto
AddCommunicationUnitsToCommunicationBus(io.clbs.openhes.models.acquisition.AddCommunicationUnitsToCommunicationBusRequest)
```

- Input: [io.clbs.openhes.models.acquisition.AddCommunicationUnitsToCommunicationBusRequest](model-io-clbs-openhes-models-acquisition-addcommunicationunitstocommunicationbusrequest.md)

## RemoveCommunicationUnitsFromCommunicationBus

```proto
RemoveCommunicationUnitsFromCommunicationBus(io.clbs.openhes.models.acquisition.RemoveCommunicationUnitsFromCommunicationBusRequest)
```

- Input: [io.clbs.openhes.models.acquisition.RemoveCommunicationUnitsFromCommunicationBusRequest](model-io-clbs-openhes-models-acquisition-removecommunicationunitsfromcommunicationbusrequest.md)

## CreateDevice

The method called by the RestAPI to register a new device. The parameter contains the device specification.

```proto
CreateDevice(io.clbs.openhes.models.acquisition.CreateDeviceRequest) returns (google.protobuf.StringValue)
```

- Input: [io.clbs.openhes.models.acquisition.CreateDeviceRequest](model-io-clbs-openhes-models-acquisition-createdevicerequest.md)
- Output: google.protobuf.StringValue

## ListDevices

The method called by the RestAPI to get the information about the device. The parameter contains the search criteria.

```proto
ListDevices(io.clbs.openhes.models.common.ListSelector) returns (io.clbs.openhes.models.acquisition.ListOfDevice)
```

- Input: [io.clbs.openhes.models.common.ListSelector](model-io-clbs-openhes-models-common-listselector.md)
- Output: [io.clbs.openhes.models.acquisition.ListOfDevice](model-io-clbs-openhes-models-acquisition-listofdevice.md)

## GetDevice

The method called by the RestAPI to get the information about the device. The parameter contains the search criteria.

```proto
GetDevice(google.protobuf.StringValue) returns (io.clbs.openhes.models.acquisition.Device)
```

- Input: google.protobuf.StringValue
- Output: [io.clbs.openhes.models.acquisition.Device](model-io-clbs-openhes-models-acquisition-device.md)

## SetDeviceCommunicationUnits

The method called by the RestAPI to replace ordered set of linked communication units.

```proto
SetDeviceCommunicationUnits(io.clbs.openhes.models.acquisition.SetDeviceCommunicationUnitsRequest)
```

- Input: [io.clbs.openhes.models.acquisition.SetDeviceCommunicationUnitsRequest](model-io-clbs-openhes-models-acquisition-setdevicecommunicationunitsrequest.md)

## GetDeviceCommunicationUnits

The method called by the RestAPI to get communication units definitions linked to the device(s).

```proto
GetDeviceCommunicationUnits(google.protobuf.StringValue) returns (io.clbs.openhes.models.acquisition.ListOfDeviceCommunicationUnit)
```

- Input: google.protobuf.StringValue
- Output: [io.clbs.openhes.models.acquisition.ListOfDeviceCommunicationUnit](model-io-clbs-openhes-models-acquisition-listofdevicecommunicationunit.md)

## CreateDeviceGroup

The method called by the RestAPI to create a new device group. The parameter contains the device group specification.

```proto
CreateDeviceGroup(io.clbs.openhes.models.acquisition.CreateDeviceGroupRequest) returns (google.protobuf.StringValue)
```

- Input: [io.clbs.openhes.models.acquisition.CreateDeviceGroupRequest](model-io-clbs-openhes-models-acquisition-createdevicegrouprequest.md)
- Output: google.protobuf.StringValue

## ListDeviceGroups

The method returns a list of device groups.

```proto
ListDeviceGroups(io.clbs.openhes.models.common.ListSelector) returns (io.clbs.openhes.models.acquisition.ListOfDeviceGroup)
```

- Input: [io.clbs.openhes.models.common.ListSelector](model-io-clbs-openhes-models-common-listselector.md)
- Output: [io.clbs.openhes.models.acquisition.ListOfDeviceGroup](model-io-clbs-openhes-models-acquisition-listofdevicegroup.md)

## GetDeviceGroup

The method returns single device group.

```proto
GetDeviceGroup(google.protobuf.StringValue) returns (io.clbs.openhes.models.acquisition.DeviceGroup)
```

- Input: google.protobuf.StringValue
- Output: [io.clbs.openhes.models.acquisition.DeviceGroup](model-io-clbs-openhes-models-acquisition-devicegroup.md)

## AddDevicesToGroup

The method called by the RestAPI to add a new device to the device group. The parameter contains the device group specification.

```proto
AddDevicesToGroup(io.clbs.openhes.models.acquisition.AddDevicesToGroupRequest)
```

- Input: [io.clbs.openhes.models.acquisition.AddDevicesToGroupRequest](model-io-clbs-openhes-models-acquisition-adddevicestogrouprequest.md)

## RemoveDevicesFromGroup

The method called by the RestAPI to remove a device from the device group. The parameter contains the device group specification.

```proto
RemoveDevicesFromGroup(io.clbs.openhes.models.acquisition.RemoveDevicesFromGroupRequest)
```

- Input: [io.clbs.openhes.models.acquisition.RemoveDevicesFromGroupRequest](model-io-clbs-openhes-models-acquisition-removedevicesfromgrouprequest.md)

## ListDeviceGroupDevices

```proto
ListDeviceGroupDevices(io.clbs.openhes.models.common.ListSelector) returns (io.clbs.openhes.models.acquisition.ListOfDevice)
```

- Input: [io.clbs.openhes.models.common.ListSelector](model-io-clbs-openhes-models-common-listselector.md)
- Output: [io.clbs.openhes.models.acquisition.ListOfDevice](model-io-clbs-openhes-models-acquisition-listofdevice.md)

## ListModemPools

The method to get list of the modem pools.

```proto
ListModemPools(io.clbs.openhes.models.common.ListSelector) returns (io.clbs.openhes.models.acquisition.ListOfModemPool)
```

- Input: [io.clbs.openhes.models.common.ListSelector](model-io-clbs-openhes-models-common-listselector.md)
- Output: [io.clbs.openhes.models.acquisition.ListOfModemPool](model-io-clbs-openhes-models-acquisition-listofmodempool.md)

## GetModemPool

The method to get the information about the modem pool. The method returns the modem pool information.

```proto
GetModemPool(google.protobuf.StringValue) returns (io.clbs.openhes.models.acquisition.ModemPool)
```

- Input: google.protobuf.StringValue
- Output: [io.clbs.openhes.models.acquisition.ModemPool](model-io-clbs-openhes-models-acquisition-modempool.md)

## CreateModemPool

The method to create a new modem pool.

```proto
CreateModemPool(io.clbs.openhes.models.acquisition.SetModemPoolRequest) returns (google.protobuf.StringValue)
```

- Input: [io.clbs.openhes.models.acquisition.SetModemPoolRequest](model-io-clbs-openhes-models-acquisition-setmodempoolrequest.md)
- Output: google.protobuf.StringValue

## UpdateModemPool

The method to update the modem pool.

```proto
UpdateModemPool(io.clbs.openhes.models.acquisition.SetModemPoolRequest)
```

- Input: [io.clbs.openhes.models.acquisition.SetModemPoolRequest](model-io-clbs-openhes-models-acquisition-setmodempoolrequest.md)

## DeleteModemPool

The method to delete the modem pool.

```proto
DeleteModemPool(google.protobuf.StringValue)
```

- Input: google.protobuf.StringValue

## CreateModem

The method to create a new modem within the pool.

```proto
CreateModem(io.clbs.openhes.models.acquisition.SetModemRequest) returns (google.protobuf.StringValue)
```

- Input: [io.clbs.openhes.models.acquisition.SetModemRequest](model-io-clbs-openhes-models-acquisition-setmodemrequest.md)
- Output: google.protobuf.StringValue

## UpdateModem

The method to update the modem within the pool.

```proto
UpdateModem(io.clbs.openhes.models.acquisition.SetModemRequest)
```

- Input: [io.clbs.openhes.models.acquisition.SetModemRequest](model-io-clbs-openhes-models-acquisition-setmodemrequest.md)

## DeleteModem

The method to delete the modem.

```proto
DeleteModem(google.protobuf.StringValue)
```

- Input: google.protobuf.StringValue

