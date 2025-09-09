# ApiService - Devices

## CreateCommunicationUnit

Creates a new communication unit. Returns the identifier of the newly created register.

```proto
CreateCommunicationUnit(io.clbs.openhes.models.acquisition.CreateCommunicationUnitRequest) returns (google.protobuf.StringValue)
```

- Input: [`io.clbs.openhes.models.acquisition.CreateCommunicationUnitRequest`](model-io-clbs-openhes-models-acquisition-createcommunicationunitrequest.md)
- Output: `google.protobuf.StringValue`

## UpdateCommunicationUnit

Updates an exiting communication unit. Requires the full unit specification; partial updates are not supported.

```proto
UpdateCommunicationUnit(io.clbs.openhes.models.acquisition.CommunicationUnit)
```

- Input: [`io.clbs.openhes.models.acquisition.CommunicationUnit`](model-io-clbs-openhes-models-acquisition-communicationunit.md)

## ListCommunicationUnits

Retrieve a paginated list of communication units based on the specified criteria. The page size and page number (zero-based) can be defined in the request.

```proto
ListCommunicationUnits(io.clbs.openhes.models.common.ListSelector) returns (io.clbs.openhes.models.acquisition.ListOfCommunicationUnit)
```

- Input: [`io.clbs.openhes.models.common.ListSelector`](model-io-clbs-openhes-models-common-listselector.md)
- Output: [`io.clbs.openhes.models.acquisition.ListOfCommunicationUnit`](model-io-clbs-openhes-models-acquisition-listofcommunicationunit.md)

## GetCommunicationUnit

Retrieves the details of the specified communication unit.

```proto
GetCommunicationUnit(google.protobuf.StringValue) returns (io.clbs.openhes.models.acquisition.CommunicationUnit)
```

- Input: `google.protobuf.StringValue`
- Output: [`io.clbs.openhes.models.acquisition.CommunicationUnit`](model-io-clbs-openhes-models-acquisition-communicationunit.md)

## DeleteCommunicationUnit

Deletes the specified communication unit.

```proto
DeleteCommunicationUnit(google.protobuf.StringValue)
```

- Input: `google.protobuf.StringValue`

## GetCommunicationUnitNetworkMap

Retrieves the network map (topology) reported by the data concentrator reports for the specified communication unit.

```proto
GetCommunicationUnitNetworkMap(google.protobuf.StringValue) returns (io.clbs.openhes.models.acquisition.NetworkMap)
```

- Input: `google.protobuf.StringValue`
- Output: [`io.clbs.openhes.models.acquisition.NetworkMap`](model-io-clbs-openhes-models-acquisition-networkmap.md)

## ListCommunicationUnitLogRecords

Retrieves a paginated list of communication unit log records based on the specified criteria. The page size and page number (zero-based) can be defined in the request.

```proto
ListCommunicationUnitLogRecords(io.clbs.openhes.models.common.ListSelector) returns (io.clbs.openhes.models.acquisition.ListOfCommunicationUnitLogRecord)
```

- Input: [`io.clbs.openhes.models.common.ListSelector`](model-io-clbs-openhes-models-common-listselector.md)
- Output: [`io.clbs.openhes.models.acquisition.ListOfCommunicationUnitLogRecord`](model-io-clbs-openhes-models-acquisition-listofcommunicationunitlogrecord.md)

## CreateCommunicationBus

Creates a new communication bus. Returns the identifier of the newly created communication bus.

```proto
CreateCommunicationBus(io.clbs.openhes.models.acquisition.CreateCommunicationBusRequest) returns (google.protobuf.StringValue)
```

- Input: [`io.clbs.openhes.models.acquisition.CreateCommunicationBusRequest`](model-io-clbs-openhes-models-acquisition-createcommunicationbusrequest.md)
- Output: `google.protobuf.StringValue`

## ListCommunicationBuses

Retrieves a paginated list of communication buses. The page size and page number (zero-based) can be defined in the request.

```proto
ListCommunicationBuses(io.clbs.openhes.models.common.ListSelector) returns (io.clbs.openhes.models.acquisition.ListOfCommunicationBus)
```

- Input: [`io.clbs.openhes.models.common.ListSelector`](model-io-clbs-openhes-models-common-listselector.md)
- Output: [`io.clbs.openhes.models.acquisition.ListOfCommunicationBus`](model-io-clbs-openhes-models-acquisition-listofcommunicationbus.md)

## DeleteCommunicationBus

Deletes the specified communication bus.

```proto
DeleteCommunicationBus(google.protobuf.StringValue)
```

- Input: `google.protobuf.StringValue`

## AddCommunicationUnitsToCommunicationBus

Adds a specified communication unit to an existing communication bus.

```proto
AddCommunicationUnitsToCommunicationBus(io.clbs.openhes.models.acquisition.AddCommunicationUnitsToCommunicationBusRequest)
```

- Input: [`io.clbs.openhes.models.acquisition.AddCommunicationUnitsToCommunicationBusRequest`](model-io-clbs-openhes-models-acquisition-addcommunicationunitstocommunicationbusrequest.md)

## RemoveCommunicationUnitsFromCommunicationBus

Removes a specified communication bus from an existing communication bus.

```proto
RemoveCommunicationUnitsFromCommunicationBus(io.clbs.openhes.models.acquisition.RemoveCommunicationUnitsFromCommunicationBusRequest)
```

- Input: [`io.clbs.openhes.models.acquisition.RemoveCommunicationUnitsFromCommunicationBusRequest`](model-io-clbs-openhes-models-acquisition-removecommunicationunitsfromcommunicationbusrequest.md)

## CreateDevice

Creates a new device. Returns the identifier of the newly created device.

```proto
CreateDevice(io.clbs.openhes.models.acquisition.CreateDeviceRequest) returns (google.protobuf.StringValue)
```

- Input: [`io.clbs.openhes.models.acquisition.CreateDeviceRequest`](model-io-clbs-openhes-models-acquisition-createdevicerequest.md)
- Output: `google.protobuf.StringValue`

## UpdateDevice

Updates the details of an existing device. Fields that are omitted from the request will be left unchanged.

```proto
UpdateDevice(io.clbs.openhes.models.acquisition.Device)
```

- Input: [`io.clbs.openhes.models.acquisition.Device`](model-io-clbs-openhes-models-acquisition-device.md)

## ListDevices

Retrieves a paginated list of devices based on the specified criteria. The page size and page number (zero-based) can be defined in the request.

```proto
ListDevices(io.clbs.openhes.models.common.ListSelector) returns (io.clbs.openhes.models.acquisition.ListOfDevice)
```

- Input: [`io.clbs.openhes.models.common.ListSelector`](model-io-clbs-openhes-models-common-listselector.md)
- Output: [`io.clbs.openhes.models.acquisition.ListOfDevice`](model-io-clbs-openhes-models-acquisition-listofdevice.md)

## GetDevice

Retrieves the details of the specified device.

```proto
GetDevice(google.protobuf.StringValue) returns (io.clbs.openhes.models.acquisition.Device)
```

- Input: `google.protobuf.StringValue`
- Output: [`io.clbs.openhes.models.acquisition.Device`](model-io-clbs-openhes-models-acquisition-device.md)

## DeleteDevice

Deletes the specified device.

```proto
DeleteDevice(google.protobuf.StringValue)
```

- Input: `google.protobuf.StringValue`

## GetDeviceInfo

Retrieves the profile-typed info of the specified device.

```proto
GetDeviceInfo(google.protobuf.StringValue) returns (io.clbs.openhes.models.acquisition.DeviceInfo)
```

- Input: `google.protobuf.StringValue`
- Output: [`io.clbs.openhes.models.acquisition.DeviceInfo`](model-io-clbs-openhes-models-acquisition-deviceinfo.md)

## SetDeviceCommunicationUnits

Sets or replaces an ordered set of communication units linked to the specified device.

```proto
SetDeviceCommunicationUnits(io.clbs.openhes.models.acquisition.SetDeviceCommunicationUnitsRequest)
```

- Input: [`io.clbs.openhes.models.acquisition.SetDeviceCommunicationUnitsRequest`](model-io-clbs-openhes-models-acquisition-setdevicecommunicationunitsrequest.md)

## GetDeviceCommunicationUnits

Retrieves a list of communication units linked to the specified device.

```proto
GetDeviceCommunicationUnits(google.protobuf.StringValue) returns (io.clbs.openhes.models.acquisition.ListOfDeviceCommunicationUnit)
```

- Input: `google.protobuf.StringValue`
- Output: [`io.clbs.openhes.models.acquisition.ListOfDeviceCommunicationUnit`](model-io-clbs-openhes-models-acquisition-listofdevicecommunicationunit.md)

## ListDeviceCommunicationUnitChanges

Retrieves a paginated list of changes to device communication units based on the specified criteria. The page size and page number (zero-based) can be defined in the request.

```proto
ListDeviceCommunicationUnitChanges(io.clbs.openhes.models.common.ListSelector) returns (io.clbs.openhes.models.acquisition.ListOfDeviceCommunicationUnitChange)
```

- Input: [`io.clbs.openhes.models.common.ListSelector`](model-io-clbs-openhes-models-common-listselector.md)
- Output: [`io.clbs.openhes.models.acquisition.ListOfDeviceCommunicationUnitChange`](model-io-clbs-openhes-models-acquisition-listofdevicecommunicationunitchange.md)

## GetDeviceDeviceGroups

Retrieves a list of device groups that contain the specified device.

```proto
GetDeviceDeviceGroups(google.protobuf.StringValue) returns (io.clbs.openhes.models.acquisition.ListOfDeviceGroup)
```

- Input: `google.protobuf.StringValue`
- Output: [`io.clbs.openhes.models.acquisition.ListOfDeviceGroup`](model-io-clbs-openhes-models-acquisition-listofdevicegroup.md)

## GetDeviceNetworkMap

Retrieves the network map (topology) reported by the data concentrator reports for the specified  device.

```proto
GetDeviceNetworkMap(google.protobuf.StringValue) returns (io.clbs.openhes.models.acquisition.NetworkMap)
```

- Input: `google.protobuf.StringValue`
- Output: [`io.clbs.openhes.models.acquisition.NetworkMap`](model-io-clbs-openhes-models-acquisition-networkmap.md)

## CreateDeviceGroup

Creates a new device group. Returns the identifier of the newly created device group.

```proto
CreateDeviceGroup(io.clbs.openhes.models.acquisition.CreateDeviceGroupRequest) returns (google.protobuf.StringValue)
```

- Input: [`io.clbs.openhes.models.acquisition.CreateDeviceGroupRequest`](model-io-clbs-openhes-models-acquisition-createdevicegrouprequest.md)
- Output: `google.protobuf.StringValue`

## UpdateDeviceGroup

Updates the details of an existing device group. Fields that are omitted from the request will be left unchanged.

```proto
UpdateDeviceGroup(io.clbs.openhes.models.acquisition.DeviceGroup)
```

- Input: [`io.clbs.openhes.models.acquisition.DeviceGroup`](model-io-clbs-openhes-models-acquisition-devicegroup.md)

## ListDeviceGroups

Retrieves a paginated list of devices groups based on the specified criteria. The page size and page number (zero-based) can be defined in the request.

```proto
ListDeviceGroups(io.clbs.openhes.models.common.ListSelector) returns (io.clbs.openhes.models.acquisition.ListOfDeviceGroup)
```

- Input: [`io.clbs.openhes.models.common.ListSelector`](model-io-clbs-openhes-models-common-listselector.md)
- Output: [`io.clbs.openhes.models.acquisition.ListOfDeviceGroup`](model-io-clbs-openhes-models-acquisition-listofdevicegroup.md)

## GetDeviceGroup

Retrieves the details of the specified device group.

### Input
The device group identifier.

### Output
The device group specification.

```proto
GetDeviceGroup(google.protobuf.StringValue) returns (io.clbs.openhes.models.acquisition.DeviceGroup)
```

- Input: `google.protobuf.StringValue`
- Output: [`io.clbs.openhes.models.acquisition.DeviceGroup`](model-io-clbs-openhes-models-acquisition-devicegroup.md)

## DeleteDeviceGroup

Deletes the specified device group.

```proto
DeleteDeviceGroup(google.protobuf.StringValue)
```

- Input: `google.protobuf.StringValue`

## AddDevicesToGroup

Adds the specified devices to an existing device group.

```proto
AddDevicesToGroup(io.clbs.openhes.models.acquisition.AddDevicesToGroupRequest)
```

- Input: [`io.clbs.openhes.models.acquisition.AddDevicesToGroupRequest`](model-io-clbs-openhes-models-acquisition-adddevicestogrouprequest.md)

## RemoveDevicesFromGroup

Removes the specified devices from an existing device group.

```proto
RemoveDevicesFromGroup(io.clbs.openhes.models.acquisition.RemoveDevicesFromGroupRequest)
```

- Input: [`io.clbs.openhes.models.acquisition.RemoveDevicesFromGroupRequest`](model-io-clbs-openhes-models-acquisition-removedevicesfromgrouprequest.md)

## ListDeviceGroupDevices

Retrieves a paginated list of devices in the specified device group. The page size and page number (zero-based) can be defined in the request.

```proto
ListDeviceGroupDevices(io.clbs.openhes.models.acquisition.ListDeviceGroupDevicesRequest) returns (io.clbs.openhes.models.acquisition.ListOfDevice)
```

- Input: [`io.clbs.openhes.models.acquisition.ListDeviceGroupDevicesRequest`](model-io-clbs-openhes-models-acquisition-listdevicegroupdevicesrequest.md)
- Output: [`io.clbs.openhes.models.acquisition.ListOfDevice`](model-io-clbs-openhes-models-acquisition-listofdevice.md)

## ListModemPools

Retrieves a paginated list of modem pools. The page size and page number (zero-based) can be defined in the request.

```proto
ListModemPools(io.clbs.openhes.models.common.ListSelector) returns (io.clbs.openhes.models.acquisition.ListOfModemPool)
```

- Input: [`io.clbs.openhes.models.common.ListSelector`](model-io-clbs-openhes-models-common-listselector.md)
- Output: [`io.clbs.openhes.models.acquisition.ListOfModemPool`](model-io-clbs-openhes-models-acquisition-listofmodempool.md)

## GetModemPool

Retrieves the details of the specified modem pool.

```proto
GetModemPool(google.protobuf.StringValue) returns (io.clbs.openhes.models.acquisition.ModemPool)
```

- Input: `google.protobuf.StringValue`
- Output: [`io.clbs.openhes.models.acquisition.ModemPool`](model-io-clbs-openhes-models-acquisition-modempool.md)

## CreateModemPool

Creates a new modem pool. Returns the identifier of the newly craeted modem pool.

```proto
CreateModemPool(io.clbs.openhes.models.acquisition.SetModemPoolRequest) returns (google.protobuf.StringValue)
```

- Input: [`io.clbs.openhes.models.acquisition.SetModemPoolRequest`](model-io-clbs-openhes-models-acquisition-setmodempoolrequest.md)
- Output: `google.protobuf.StringValue`

## UpdateModemPool

Updates the details of an existing modem pool. Fields that are omitted from the request will be left unchanged.

```proto
UpdateModemPool(io.clbs.openhes.models.acquisition.SetModemPoolRequest)
```

- Input: [`io.clbs.openhes.models.acquisition.SetModemPoolRequest`](model-io-clbs-openhes-models-acquisition-setmodempoolrequest.md)

## DeleteModemPool

Deletes the specified modem pool.

```proto
DeleteModemPool(google.protobuf.StringValue)
```

- Input: `google.protobuf.StringValue`

## CreateModem

Creates a new modem within an existing modem pool. Returns the identifier of the newly created modem.

```proto
CreateModem(io.clbs.openhes.models.acquisition.SetModemRequest) returns (google.protobuf.StringValue)
```

- Input: [`io.clbs.openhes.models.acquisition.SetModemRequest`](model-io-clbs-openhes-models-acquisition-setmodemrequest.md)
- Output: `google.protobuf.StringValue`

## UpdateModem

Updates the details of an existing modem within the specified modem pool.

```proto
UpdateModem(io.clbs.openhes.models.acquisition.SetModemRequest)
```

- Input: [`io.clbs.openhes.models.acquisition.SetModemRequest`](model-io-clbs-openhes-models-acquisition-setmodemrequest.md)

## DeleteModem

Deletes th specified modem.

```proto
DeleteModem(google.protobuf.StringValue)
```

- Input: `google.protobuf.StringValue`

