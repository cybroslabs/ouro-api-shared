package device

import (
	"errors"

	"github.com/cybroslabs/hes-2-apis/openapi/openhes/attribute"
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

	ci := communicationUnit.ConnectionInfo

	connection_info := &pbdriver.ConnectionInfo{}

	result := &pbdeviceregistry.CommunicationUnitSpec{
		Id:             communicationUnit.Id.String(),
		ExternalId:     communicationUnit.ExternalID,
		Name:           communicationUnit.Name,
		ConnectionInfo: connection_info,
	}

	if tcp, err := ci.AsConnectionTypeTcpIpSchema(); err == nil {
		connection_info.Connection = &pbdriver.ConnectionInfo_Tcpip{
			Tcpip: &pbdriver.ConnectionTypeDirectTcpIp{
				Host: tcp.Host,
				Port: uint32(tcp.Port),
			},
		}
	} else if modem, err := ci.AsConnectionTypePhoneLineSchema(); err == nil {
		connection_info.Connection = &pbdriver.ConnectionInfo_ModemPool{
			ModemPool: &pbdriver.ConnectionTypeModemPool{
				Number: modem.Number,
			},
		}
	} else if serial_moxa, err := ci.AsConnectionTypeSerialMoxaSchema(); err == nil {
		connection_info.Connection = &pbdriver.ConnectionInfo_SerialOverIp{
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
	} else if serial_direct, err := ci.AsConnectionTypeSerialDirectSchema(); err == nil {
		connection_info.Connection = &pbdriver.ConnectionInfo_SerialOverIp{
			SerialOverIp: &pbdriver.ConnectionTypeControlledSerial{
				Converter: &pbdriver.ConnectionTypeControlledSerial_Direct{
					Direct: &pbdriver.ConnectionTypeSerialDirect{
						Host: serial_direct.Host,
						Port: uint32(serial_direct.Port),
					},
				},
			},
		}
	} else {
		return nil, ErrInvalidConnectionInfo
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
			err = result.ConnectionInfo.FromConnectionTypeTcpIpSchema(ConnectionTypeTcpIpSchema{
				Host: tcpip.Host,
				Port: tcpip.Port,
			})
		} else if modem := ci.GetModemPool(); modem != nil {
			err = result.ConnectionInfo.FromConnectionTypePhoneLineSchema(ConnectionTypePhoneLineSchema{
				Number: modem.Number,
			})
		} else if controlled_serial := ci.GetSerialOverIp(); controlled_serial != nil {
			if serial_moxa := controlled_serial.GetMoxa(); serial_moxa != nil {
				err = result.ConnectionInfo.FromConnectionTypeSerialMoxaSchema(ConnectionTypeSerialMoxaSchema{
					Host:        serial_moxa.Host,
					DataPort:    serial_moxa.DataPort,
					CommandPort: serial_moxa.CommandPort,
				})
			} else if serial_direct := controlled_serial.GetDirect(); serial_direct != nil {
				err = result.ConnectionInfo.FromConnectionTypeSerialDirectSchema(ConnectionTypeSerialDirectSchema{
					Host: serial_direct.Host,
					Port: serial_direct.Port,
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

// Converts the modem pool - Rest API to gRPC
func R2GModemPool(modemPool *ModemPoolSchema) (*pbdeviceregistry.ModemPoolSpec, error) {
	if modemPool == nil {
		return nil, nil
	}
	result := &pbdeviceregistry.ModemPoolSpec{
		PoolId: modemPool.Id.String(),
		Name:   modemPool.Name,
	}
	return result, nil
}

// Converts the modem pool - gRPC to Rest API
func G2RModemPool(modemPool *pbdeviceregistry.ModemPoolSpec) (*ModemPoolSchema, error) {
	if modemPool == nil {
		return nil, nil
	}
	pool_id, err := uuid.Parse(modemPool.PoolId)
	if err != nil {
		return nil, err
	}
	result := &ModemPoolSchema{
		Id:   pool_id,
		Name: modemPool.Name,
	}
	return result, nil
}

// Converts the modem - Rest API to gRPC
func R2GModem(modem *ModemSchema) (*pbdriver.ModemInfo, error) {
	if modem == nil {
		return nil, nil
	}

	result := &pbdriver.ModemInfo{
		Id:             modem.Id.String(),
		Name:           modem.Name,
		AtInit:         modem.AtInit,
		AtTest:         modem.AtTest,
		AtConfig:       modem.AtConfig,
		AtDial:         modem.AtDial,
		AtHangup:       modem.AtHangup,
		AtEscape:       modem.AtEscape,
		AtDsr:          modem.AtDsr,
		ConnectTimeout: modem.ConnectTimeout,
	}

	if tcp_ip, err := modem.Connection.AsConnectionTypeTcpIpSchema(); err == nil {
		result.ModemConnection = &pbdriver.ModemInfo_Tcpip{
			Tcpip: &pbdriver.ConnectionTypeDirectTcpIp{
				Host: tcp_ip.Host,
				Port: uint32(tcp_ip.Port),
			},
		}
	} else {
		return nil, ErrInvalidConnectionInfo
	}

	return result, nil
}

// Converts the modem - gRPC to Rest API
func G2RModem(modem *pbdriver.ModemInfo) (*ModemSchema, error) {
	if modem == nil {
		return nil, nil
	}
	modem_id, err := uuid.Parse(modem.Id)
	if err != nil {
		return nil, err
	}
	result := &ModemSchema{
		Id:             modem_id,
		Name:           modem.Name,
		AtInit:         modem.AtInit,
		AtTest:         modem.AtTest,
		AtConfig:       modem.AtConfig,
		AtDial:         modem.AtDial,
		AtHangup:       modem.AtHangup,
		AtEscape:       modem.AtEscape,
		AtDsr:          modem.AtDsr,
		ConnectTimeout: modem.ConnectTimeout,
	}

	if tcpip := modem.GetTcpip(); tcpip != nil {
		err = result.Connection.FromConnectionTypeTcpIpSchema(ConnectionTypeTcpIpSchema{
			Host: tcpip.Host,
			Port: tcpip.Port,
		})
		if err != nil {
			return nil, err
		}
	} else {
		return nil, ErrInvalidConnectionInfo
	}

	return result, nil
}
