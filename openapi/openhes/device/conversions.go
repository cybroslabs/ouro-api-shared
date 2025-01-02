package device

import (
	"errors"

	"github.com/cybroslabs/hes-2-apis/openapi/openhes/attribute"
	"github.com/cybroslabs/hes-2-apis/openapi/openhes/job"
	"github.com/cybroslabs/hes-2-apis/protobuf/pbdeviceregistry"
	"github.com/cybroslabs/hes-2-apis/protobuf/pbdriver"
	"github.com/google/uuid"
)

var (
	ErrInvalidConnectionInfo = errors.New("invalid connection info")
)

// Converts the communication unit - Rest API to gRPC
func R2GCommunicationUnit(communicationUnit *CommunicationUnitSchema) (*pbdeviceregistry.CommunicationUnitSpec, error) {
	if communicationUnit == nil {
		return nil, nil
	}

	result := &pbdeviceregistry.CommunicationUnitSpec{
		Id:         communicationUnit.Id.String(),
		ExternalId: communicationUnit.ExternalID,
		Name:       communicationUnit.Name,
	}

	ci := communicationUnit.ConnectionInfo
	if tcp, err := ci.AsConnectionTypeTcpIpSchema(); err == nil {
		ci := &pbdriver.ConnectionInfo_Tcpip{
			Tcpip: &pbdriver.ConnectionTypeDirectTcpIp{
				Host: tcp.Host,
				Port: uint32(tcp.Port),
			},
		}
		result.ConnectionInfo = &pbdriver.ConnectionInfo{Connection: ci}
	} else if modem, err := ci.AsConnectionTypePhoneLineSchema(); err == nil {
		ci := &pbdriver.ConnectionInfo_ModemPool{
			ModemPool: &pbdriver.ConnectionTypeModemPool{
				Number: modem.Number,
			},
		}
		result.ConnectionInfo = &pbdriver.ConnectionInfo{Connection: ci}
	} else if serial_moxa, err := ci.AsConnectionTypeSerialMoxaSchema(); err == nil {
		ci := &pbdriver.ConnectionInfo_SerialOverIp{
			SerialOverIp: &pbdriver.ConnectionTypeControlledSerial{
				Converter: &pbdriver.ConnectionTypeControlledSerial_Moxa{
					Moxa: &pbdriver.ConnectionTypeSerialMoxa{
						Host:        serial_moxa.Host,
						DataPort:    uint32(serial_moxa.DataPort),
						CommandPort: uint32(serial_moxa.CommandPort),
					},
				},
			},
		}
		result.ConnectionInfo = &pbdriver.ConnectionInfo{Connection: ci}
	} else if serial_direct, err := ci.AsConnectionTypeSerialDirectSchema(); err == nil {
		ci := &pbdriver.ConnectionTypeControlledSerial_Direct{
			Direct: &pbdriver.ConnectionTypeSerialDirect{
				Host: serial_direct.Host,
				Port: uint32(serial_direct.Port),
			},
		}
		result.ConnectionInfo = &pbdriver.ConnectionInfo{Connection: ci}
	}

	return result, nil
}

// Converts the communication unit - gRPC to Rest API
func G2RCommunicationUnit(communicationUnit *pbdeviceregistry.CommunicationUnitSpec) (*CommunicationUnitSchema, error) {
	if communicationUnit == nil {
		return nil, nil
	}

	result := &CommunicationUnitSchema{
		Id:         uuid.MustParse(communicationUnit.Id),
		ExternalID: communicationUnit.ExternalId,
		Name:       communicationUnit.Name,
	}

	var err error

	if ci := communicationUnit.ConnectionInfo; ci != nil {
		if tcpip := ci.GetTcpip(); tcpip != nil {
			err = result.ConnectionInfo.FromConnectionTypeTcpIpSchema(job.ConnectionTypeTcpIpSchema{
				Host: tcpip.Host,
				Port: int(tcpip.Port),
			})
		} else if modem := ci.GetModemPool(); modem != nil {
			err = result.ConnectionInfo.FromConnectionTypePhoneLineSchema(job.ConnectionTypePhoneLineSchema{
				Number: modem.Number,
			})
		} else if controlled_serial := ci.GetSerialOverIp(); controlled_serial != nil {
			if serial_moxa := controlled_serial.GetMoxa(); serial_moxa != nil {
				err = result.ConnectionInfo.FromConnectionTypeSerialMoxaSchema(job.ConnectionTypeSerialMoxaSchema{
					Host:        serial_moxa.Host,
					DataPort:    int(serial_moxa.DataPort),
					CommandPort: int(serial_moxa.CommandPort),
				})
			} else if serial_direct := controlled_serial.GetDirect(); serial_direct != nil {
				err = result.ConnectionInfo.FromConnectionTypeSerialDirectSchema(job.ConnectionTypeSerialDirectSchema{
					Host: serial_direct.Host,
					Port: int(serial_direct.Port),
				})
			} else {
				err = ErrInvalidConnectionInfo
			}
		} else {
			err = ErrInvalidConnectionInfo
		}
	} else {
		err = ErrInvalidConnectionInfo
	}

	if err != nil {
		return nil, err
	}

	return result, nil
}

// Converts the device - Rest API to gRPC
func R2GDevice(device *DeviceSchema) (*pbdeviceregistry.DeviceSpec, error) {
	if device == nil {
		return nil, nil
	}

	attributes, err := attribute.R2GAttributes(device.Attributes)
	if err != nil {
		return nil, err
	}

	cu_cnt := len(device.CommunicationUnitID)

	result := &pbdeviceregistry.DeviceSpec{
		Id:                  device.Id.String(),
		ExternalId:          device.ExternalID,
		Name:                device.Name,
		Attributes:          attributes,
		CommunicationUnitId: make([]string, cu_cnt),
		Timezone:            device.Timezone,
	}

	tmp := result.CommunicationUnitId
	for i, id := range device.CommunicationUnitID {
		tmp[i] = id.String()
	}

	return result, nil
}

// Converts the device - gRPC to Rest API
func G2RDevice(device *pbdeviceregistry.DeviceSpec) (*DeviceSchema, error) {
	if device == nil {
		return nil, nil
	}

	cu_cnt := len(device.CommunicationUnitId)

	var attrs attribute.Attributes
	if da := device.Attributes; da != nil {
		attrs = attribute.G2RAttributes(da)
	} else {
		attrs = nil
	}

	result := &DeviceSchema{
		Id:                  uuid.MustParse(device.Id),
		ExternalID:          device.ExternalId,
		Name:                device.Name,
		Attributes:          attrs,
		CommunicationUnitID: make([]uuid.UUID, cu_cnt),
		Timezone:            device.Timezone,
	}

	tmp := result.CommunicationUnitID
	for i, id := range device.CommunicationUnitId {
		tmp[i] = uuid.MustParse(id)
	}

	return result, nil
}

// Converts the device group type - Rest API to gRPC
func R2GDeviceGroupType(deviceGroupType *DeviceGroupSchema) (*pbdeviceregistry.DeviceGroupSpec, error) {
	if deviceGroupType == nil {
		return nil, nil
	}

	result := &pbdeviceregistry.DeviceGroupSpec{
		Id:         deviceGroupType.Id.String(),
		ExternalId: deviceGroupType.ExternalID,
		Name:       deviceGroupType.Name,
	}

	return result, nil
}

// Converts the device group type - gRPC to Rest API
func G2RDeviceGroupType(deviceGroupType *pbdeviceregistry.DeviceGroupSpec) (*DeviceGroupSchema, error) {
	if deviceGroupType == nil {
		return nil, nil
	}

	result := &DeviceGroupSchema{
		Id:         uuid.MustParse(deviceGroupType.Id),
		ExternalID: deviceGroupType.ExternalId,
		Name:       deviceGroupType.Name,
	}

	return result, nil
}
