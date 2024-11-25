package device

import (
	"github.com/cybroslabs/hes-2-apis/openapi/openhes/attribute"
	"github.com/cybroslabs/hes-2-apis/openapi/openhes/job"
	"github.com/cybroslabs/hes-2-apis/protobuf/pbdeviceregistry"
	"github.com/cybroslabs/hes-2-apis/protobuf/pbdriver"
	"github.com/google/uuid"
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
	if tcp, err := ci.AsConnectionTypeTcpSchema(); err == nil {
		ci := &pbdeviceregistry.ConnectionType_Tcp{
			Tcp: &pbdriver.ConnectionTypeTcp{
				Host: tcp.Host,
				Port: uint32(tcp.Port),
			},
		}
		result.ConnectionType = &pbdeviceregistry.ConnectionType{Type: ci}
	} else if modem, err := ci.AsConnectionTypeModemSchema(); err == nil {
		ci := &pbdeviceregistry.ConnectionType_Modem{
			Modem: &pbdriver.ConnectionTypeModem{
				Number: modem.Number,
			},
		}
		result.ConnectionType = &pbdeviceregistry.ConnectionType{Type: ci}
	} else if moxa, err := ci.AsConnectionTypeMoxaSchema(); err == nil {
		ci := &pbdeviceregistry.ConnectionType_Moxa{
			Moxa: &pbdriver.ConnectionTypeMoxa{
				Host:        moxa.Host,
				DataPort:    uint32(moxa.DataPort),
				CommandPort: uint32(moxa.CommandPort),
			},
		}
		result.ConnectionType = &pbdeviceregistry.ConnectionType{Type: ci}
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

	if ct := communicationUnit.ConnectionType; ct != nil {
		if tcp := ct.GetTcp(); tcp != nil {
			err := result.ConnectionInfo.FromConnectionTypeTcpSchema(job.ConnectionTypeTcpSchema{
				Host: tcp.Host,
				Port: int(tcp.Port),
			})
			if err != nil {
				return nil, err
			}
		} else if modem := ct.GetModem(); modem != nil {
			err := result.ConnectionInfo.FromConnectionTypeModemSchema(job.ConnectionTypeModemSchema{
				Number: modem.Number,
			})
			if err != nil {
				return nil, err
			}
		} else if moxa := ct.GetMoxa(); moxa != nil {
			err := result.ConnectionInfo.FromConnectionTypeMoxaSchema(job.ConnectionTypeMoxaSchema{
				Host:        moxa.Host,
				DataPort:    int(moxa.DataPort),
				CommandPort: int(moxa.CommandPort),
			})
			if err != nil {
				return nil, err
			}
		}
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
		Attributes:          &pbdeviceregistry.KeyValue{Kv: attributes},
		CommunicationUnitId: make([]string, cu_cnt),
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
		attrs = attribute.G2RAttributes(da.Kv)
	} else {
		attrs = nil
	}

	result := &DeviceSchema{
		Id:                  uuid.MustParse(device.Id),
		ExternalID:          device.ExternalId,
		Name:                device.Name,
		Attributes:          attrs,
		CommunicationUnitID: make([]uuid.UUID, cu_cnt),
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
