# ApiService - Devices

## CreateDeviceGroup

The method called by the RestAPI to create a new device group. The parameter contains the device group specification.

```proto
CreateDeviceGroup(io.clbs.openhes.models.acquisition.CreateDeviceGroupRequest) returns (google.protobuf.StringValue)
```

- Input: [`io.clbs.openhes.models.acquisition.CreateDeviceGroupRequest`](model-io-clbs-openhes-models-acquisition-createdevicegrouprequest.md)
- Output: `google.protobuf.StringValue`

## ListDeviceGroups

The method returns a list of device groups.

```proto
ListDeviceGroups(io.clbs.openhes.models.common.ListSelector) returns (io.clbs.openhes.models.acquisition.ListOfDeviceGroup)
```

- Input: [`io.clbs.openhes.models.common.ListSelector`](model-io-clbs-openhes-models-common-listselector.md)
- Output: [`io.clbs.openhes.models.acquisition.ListOfDeviceGroup`](model-io-clbs-openhes-models-acquisition-listofdevicegroup.md)

## GetDeviceGroup

The method returns single device group.

```proto
GetDeviceGroup(google.protobuf.StringValue) returns (io.clbs.openhes.models.acquisition.DeviceGroup)
```

- Input: `google.protobuf.StringValue`
- Output: [`io.clbs.openhes.models.acquisition.DeviceGroup`](model-io-clbs-openhes-models-acquisition-devicegroup.md)

## DeleteDeviceGroup

Deletes the device group. The parameter contains the device group identifier.

```proto
DeleteDeviceGroup(google.protobuf.StringValue)
```

- Input: `google.protobuf.StringValue`

## AddDevicesToGroup

The method called by the RestAPI to add a new device to the device group. The parameter contains the device group specification.

```proto
AddDevicesToGroup(io.clbs.openhes.models.acquisition.AddDevicesToGroupRequest)
```

- Input: [`io.clbs.openhes.models.acquisition.AddDevicesToGroupRequest`](model-io-clbs-openhes-models-acquisition-adddevicestogrouprequest.md)

## RemoveDevicesFromGroup

The method called by the RestAPI to remove a device from the device group. The parameter contains the device group specification.

```proto
RemoveDevicesFromGroup(io.clbs.openhes.models.acquisition.RemoveDevicesFromGroupRequest)
```

- Input: [`io.clbs.openhes.models.acquisition.RemoveDevicesFromGroupRequest`](model-io-clbs-openhes-models-acquisition-removedevicesfromgrouprequest.md)

## ListDeviceGroupDevices

```proto
ListDeviceGroupDevices(io.clbs.openhes.models.acquisition.ListDeviceGroupDevicesRequest) returns (io.clbs.openhes.models.acquisition.ListOfDevice)
```

- Input: [`io.clbs.openhes.models.acquisition.ListDeviceGroupDevicesRequest`](model-io-clbs-openhes-models-acquisition-listdevicegroupdevicesrequest.md)
- Output: [`io.clbs.openhes.models.acquisition.ListOfDevice`](model-io-clbs-openhes-models-acquisition-listofdevice.md)

## ListModemPools

The method to get list of the modem pools.

```proto
ListModemPools(io.clbs.openhes.models.common.ListSelector) returns (io.clbs.openhes.models.acquisition.ListOfModemPool)
```

- Input: [`io.clbs.openhes.models.common.ListSelector`](model-io-clbs-openhes-models-common-listselector.md)
- Output: [`io.clbs.openhes.models.acquisition.ListOfModemPool`](model-io-clbs-openhes-models-acquisition-listofmodempool.md)

## GetModemPool

The method to get the information about the modem pool. The method returns the modem pool information.

```proto
GetModemPool(google.protobuf.StringValue) returns (io.clbs.openhes.models.acquisition.ModemPool)
```

- Input: `google.protobuf.StringValue`
- Output: [`io.clbs.openhes.models.acquisition.ModemPool`](model-io-clbs-openhes-models-acquisition-modempool.md)

## CreateModemPool

The method to create a new modem pool.

```proto
CreateModemPool(io.clbs.openhes.models.acquisition.SetModemPoolRequest) returns (google.protobuf.StringValue)
```

- Input: [`io.clbs.openhes.models.acquisition.SetModemPoolRequest`](model-io-clbs-openhes-models-acquisition-setmodempoolrequest.md)
- Output: `google.protobuf.StringValue`

## UpdateModemPool

The method to update the modem pool.

```proto
UpdateModemPool(io.clbs.openhes.models.acquisition.SetModemPoolRequest)
```

- Input: [`io.clbs.openhes.models.acquisition.SetModemPoolRequest`](model-io-clbs-openhes-models-acquisition-setmodempoolrequest.md)

## DeleteModemPool

The method to delete the modem pool.

```proto
DeleteModemPool(google.protobuf.StringValue)
```

- Input: `google.protobuf.StringValue`

## CreateModem

The method to create a new modem within the pool.

```proto
CreateModem(io.clbs.openhes.models.acquisition.SetModemRequest) returns (google.protobuf.StringValue)
```

- Input: [`io.clbs.openhes.models.acquisition.SetModemRequest`](model-io-clbs-openhes-models-acquisition-setmodemrequest.md)
- Output: `google.protobuf.StringValue`

## UpdateModem

The method to update the modem within the pool.

```proto
UpdateModem(io.clbs.openhes.models.acquisition.SetModemRequest)
```

- Input: [`io.clbs.openhes.models.acquisition.SetModemRequest`](model-io-clbs-openhes-models-acquisition-setmodemrequest.md)

## DeleteModem

The method to delete the modem.

```proto
DeleteModem(google.protobuf.StringValue)
```

- Input: `google.protobuf.StringValue`

