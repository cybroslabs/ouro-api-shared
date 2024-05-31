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
		ci := &pbdeviceregistry.CommunicationUnitSpec_Tcp{
			Tcp: &pbdriver.ConnectionTypeTcp{
				Host: tcp.Host,
			},
		}
		result.ConnectionType = ci
	} else if phone, err := ci.AsConnectionTypePhoneSchema(); err == nil {
		ci := &pbdeviceregistry.CommunicationUnitSpec_Phone{
			Phone: &pbdriver.ConnectionTypePhone{
				Number: phone.Number,
			},
		}
		result.ConnectionType = ci
	} else if _, err := ci.AsConnectionTypeSerialSchema(); err == nil {
		ci := &pbdeviceregistry.CommunicationUnitSpec_Serial{
			Serial: &pbdriver.ConnectionTypeSerial{},
		}
		result.ConnectionType = ci
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

	if tcp := communicationUnit.GetTcp(); tcp != nil {
		result.ConnectionInfo.FromConnectionTypeTcpSchema(job.ConnectionTypeTcpSchema{
			Host: tcp.Host,
			Port: int(tcp.Port),
		})
	} else if phone := communicationUnit.GetPhone(); phone != nil {
		result.ConnectionInfo.FromConnectionTypePhoneSchema(job.ConnectionTypePhoneSchema{
			Number: phone.Number,
		})
	} else if serial := communicationUnit.GetSerial(); serial != nil {
		result.ConnectionInfo.FromConnectionTypeSerialSchema(job.ConnectionTypeSerialSchema{})
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

	result := &DeviceSchema{
		Id:                  uuid.MustParse(device.Id),
		ExternalID:          device.ExternalId,
		Name:                device.Name,
		Attributes:          attribute.G2RAttributes(device.Attributes),
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
