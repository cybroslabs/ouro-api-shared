package driver

import (
	"errors"
	"fmt"
	"strings"

	"github.com/cybroslabs/hes-2-apis/protobuf/pbdriver"
)

var (
	ErrInvalidActionType = errors.New("invalid action type")
)

// Converts the attribute type - gRPC to REST
func G2RAttributeType(attributeType pbdriver.AttributeType) AttributeDefinitionSchemaType {
	switch attributeType {
	case pbdriver.AttributeType_ATTRIBUTE_TYPE_STRING:
		return STRING
	case pbdriver.AttributeType_ATTRIBUTE_TYPE_INT:
		return INTEGER
	case pbdriver.AttributeType_ATTRIBUTE_TYPE_DOUBLE:
		return NUMBER
	case pbdriver.AttributeType_ATTRIBUTE_TYPE_BINARY:
		return BINARY
	case pbdriver.AttributeType_ATTRIBUTE_TYPE_TIMESTAMP:
		return TIMESTAMP
	case pbdriver.AttributeType_ATTRIBUTE_TYPE_TIMESTAMP_TZ:
		return TIMESTAMPTZ
	case pbdriver.AttributeType_ATTRIBUTE_TYPE_BOOL:
		return BOOLEAN
	default:
		return STRING
	}
}

// Converts the attribute definition - gRPC to REST
func G2RAttributeDefinition(attrDef *pbdriver.AttributeDefinition) (*AttributeDefinitionSchema, error) {
	if attrDef == nil {
		return nil, nil
	}

	attr_type := G2RAttributeType(attrDef.Type)

	result := &AttributeDefinitionSchema{
		Name:        &attrDef.Name,
		Description: &attrDef.Description,
		Type:        &attr_type,
		Mandatory:   &attrDef.Mandatory,
	}

	// Some default-value content validation
	if attrDef.DefaultValue != nil {
		var tmp interface{}
		switch attr_type {
		case STRING:
			if d, ok := attrDef.DefaultValue.Value.(*pbdriver.AttributeValue_StrValue); !ok {
				return nil, fmt.Errorf("attribute %s is of type STRING but default value is not of type StrValue", attrDef.Name)
			} else {
				tmp = d.StrValue
			}
		case INTEGER:
			if d, ok := attrDef.DefaultValue.Value.(*pbdriver.AttributeValue_IntValue); !ok {
				return nil, fmt.Errorf("attribute %s is of type INTEGER but default value is not of type IntValue", attrDef.Name)
			} else {
				tmp = d.IntValue
			}
		case NUMBER:
			if d, ok := attrDef.DefaultValue.Value.(*pbdriver.AttributeValue_DoubleValue); !ok {
				return nil, fmt.Errorf("attribute %s is of type NUMBER but default value is not of type DoubleValue", attrDef.Name)
			} else {
				tmp = d.DoubleValue
			}
		case BINARY:
			if d, ok := attrDef.DefaultValue.Value.(*pbdriver.AttributeValue_BinaryValue); !ok {
				return nil, fmt.Errorf("attribute %s is of type BINARY but default value is not of type BinaryValue", attrDef.Name)
			} else {
				tmp = d.BinaryValue
			}
		case BOOLEAN:
			if d, ok := attrDef.DefaultValue.Value.(*pbdriver.AttributeValue_BoolValue); !ok {
				return nil, fmt.Errorf("attribute %s is of type BOOLEAN but default value is not of type BoolValue", attrDef.Name)
			} else {
				tmp = d.BoolValue
			}
		case TIMESTAMP:
			if d, ok := attrDef.DefaultValue.Value.(*pbdriver.AttributeValue_IntValue); !ok {
				return nil, fmt.Errorf("attribute %s is of type TIMESTAMP but default value is not of type IntValue", attrDef.Name)
			} else {
				tmp = d.IntValue
			}
		case TIMESTAMPTZ:
			if d, ok := attrDef.DefaultValue.Value.(*pbdriver.AttributeValue_StrValue); !ok {
				return nil, fmt.Errorf("attribute %s is of type TIMESTAMPTZ but default value is not of type StrValue", attrDef.Name)
			} else {
				tmp = d.StrValue
			}
		default:
			return nil, fmt.Errorf("unknown attribute type: %v", attrDef.Type)
		}
		if tmp != nil {
			result.DefaultValue = &tmp
		}
	}

	return result, nil
}

func G2RCommunicationTemplate(commTemp *pbdriver.CommunicationTemplate, appProtocols []*pbdriver.ApplicationProtocolTemplate) (*DriverCommunicationTemplateSchema, error) {
	if commTemp == nil {
		return nil, nil
	}

	t := commTemp.Type.String()
	datalink_tpls := make([]DriverDatalinkTemplateSchema, len(commTemp.Datalinks))
	result := &DriverCommunicationTemplateSchema{
		Type:              &t,
		DatalinkTemplates: &datalink_tpls,
	}

	for dl_idx, dl_data := range commTemp.Datalinks {
		if dl_data == nil {
			return nil, fmt.Errorf("datalink template contains nil")
		}

		dl_out := &datalink_tpls[dl_idx]

		attr_out := make([]AttributeDefinitionSchema, 0, len(dl_data.Attributes))
		for _, attr := range dl_data.Attributes {
			attr_def, err := G2RAttributeDefinition(attr)
			if err != nil {
				return nil, err
			}
			if attr_def != nil {
				attr_out = append(attr_out, *attr_def)
			}
		}
		dl_out.Attributes = attr_out

		lp, err := G2RDataLinkProtocol(dl_data.LinkProtocol)
		if err != nil {
			return nil, err
		}
		dl_out.LinkProtocol = lp

		if dl_app_cnt := len(dl_data.AppProtocolRefs); dl_app_cnt > 0 {
			// Let's use the same slice for the app protocols
			ap := make(DriverAppProtocolRefsSchema, dl_app_cnt)
			for ap_idx, ap_data := range dl_data.AppProtocolRefs {
				ap_out, err := G2RAppProtocol(ap_data)
				if err != nil {
					return nil, err
				}
				ap[ap_idx] = *ap_out
			}
			dl_out.AppProtocolRefs = ap
		}
	}

	return result, nil
}

// Converts the action type - Rest API to gRPC
func R2GActionType(actionType ActionTypeSchema) (pbdriver.ActionType, error) {
	action_name := "ACTION_TYPE_" + string(actionType)
	no, ok := pbdriver.ActionType_value[action_name]
	if !ok {
		return -1, ErrInvalidActionType
	}
	return pbdriver.ActionType(no), nil
}

// Converts the action type - gRPC to Rest API
func G2RActionType(actionType pbdriver.ActionType) (ActionTypeSchema, error) {
	no := int32(actionType.Number())
	action_name, ok := pbdriver.ActionType_name[no]
	if !ok {
		return "", ErrInvalidActionType
	}
	result, ok := strings.CutPrefix(action_name, "ACTION_TYPE_")
	if !ok {
		return "", ErrInvalidActionType
	}
	return ActionTypeSchema(result), nil
}

var (
	appProtoANSIC12    = ApplicationProtocolSchemaANSIC12
	appProtoDLMSLN     = ApplicationProtocolSchemaDLMSLN
	appProtoDLMSSN     = ApplicationProtocolSchemaDLMSSN
	appProtoIEC6205621 = ApplicationProtocolSchemaIEC6205621
	appProtoLIS200     = ApplicationProtocolSchemaLIS200
	appProtoMQTT       = ApplicationProtocolSchemaMQTT
	appProtoSCTM       = ApplicationProtocolSchemaSCTM
)

var (
	dlProtoCOSEMWRAPPER  = DataLinkProtocolSchemaCOSEMWRAPPER
	dlProtoHDLC          = DataLinkProtocolSchemaHDLC
	dlProtoIEC6205621    = DataLinkProtocolSchemaIEC6205621
	dlProtoMBUS          = DataLinkProtocolSchemaMBUS
	dlProtoMODBUS        = DataLinkProtocolSchemaMODBUS
	dlProtoNOTAPPLICABLE = DataLinkProtocolSchemaNOTAPPLICABLE
)

func R2GAppProtocol(proto *ApplicationProtocolSchema) (pbdriver.ApplicationProtocol, error) {
	if proto == nil {
		return 0, fmt.Errorf("application protocol is nil")
	} else if *proto == ApplicationProtocolSchemaDLMSLN {
		return pbdriver.ApplicationProtocol_APPPROTO_DLMS_LN, nil
	} else if *proto == ApplicationProtocolSchemaDLMSSN {
		return pbdriver.ApplicationProtocol_APPPROTO_DLMS_SN, nil
	} else if *proto == ApplicationProtocolSchemaANSIC12 {
		return pbdriver.ApplicationProtocol_APPPROTO_ANSI_C12, nil
	} else if *proto == ApplicationProtocolSchemaIEC6205621 {
		return pbdriver.ApplicationProtocol_APPPROTO_IEC_62056_21, nil
	} else if *proto == ApplicationProtocolSchemaLIS200 {
		return pbdriver.ApplicationProtocol_APPPROTO_LIS200, nil
	} else if *proto == ApplicationProtocolSchemaMQTT {
		return pbdriver.ApplicationProtocol_APPPROTO_MQTT, nil
	} else if *proto == ApplicationProtocolSchemaSCTM {
		return pbdriver.ApplicationProtocol_APPPROTO_SCTM, nil
	} else {
		return 0, fmt.Errorf("invalid application protocol %s", *proto)
	}
}

func G2RAppProtocol(proto pbdriver.ApplicationProtocol) (*ApplicationProtocolSchema, error) {
	if proto == pbdriver.ApplicationProtocol_APPPROTO_DLMS_LN {
		return &appProtoDLMSLN, nil
	} else if proto == pbdriver.ApplicationProtocol_APPPROTO_DLMS_SN {
		return &appProtoDLMSSN, nil
	} else if proto == pbdriver.ApplicationProtocol_APPPROTO_IEC_62056_21 {
		return &appProtoIEC6205621, nil
	} else if proto == pbdriver.ApplicationProtocol_APPPROTO_LIS200 {
		return &appProtoLIS200, nil
	} else if proto == pbdriver.ApplicationProtocol_APPPROTO_MQTT {
		return &appProtoMQTT, nil
	} else if proto == pbdriver.ApplicationProtocol_APPPROTO_SCTM {
		return &appProtoSCTM, nil
	} else if proto == pbdriver.ApplicationProtocol_APPPROTO_ANSI_C12 {
		return &appProtoANSIC12, nil
	} else {
		return nil, fmt.Errorf("unknown application protocol: %v", proto)
	}
}

func R2GDataLinkProtocol(proto DataLinkProtocolSchema) (pbdriver.DataLinkProtocol, error) {
	if proto == DataLinkProtocolSchemaCOSEMWRAPPER {
		return pbdriver.DataLinkProtocol_LINKPROTO_COSEM_WRAPPER, nil
	} else if proto == DataLinkProtocolSchemaHDLC {
		return pbdriver.DataLinkProtocol_LINKPROTO_HDLC, nil
	} else if proto == DataLinkProtocolSchemaIEC6205621 {
		return pbdriver.DataLinkProtocol_LINKPROTO_IEC_62056_21, nil
	} else if proto == DataLinkProtocolSchemaMBUS {
		return pbdriver.DataLinkProtocol_LINKPROTO_MBUS, nil
	} else if proto == DataLinkProtocolSchemaMODBUS {
		return pbdriver.DataLinkProtocol_LINKPROTO_MODBUS, nil
	} else if proto == dlProtoNOTAPPLICABLE {
		return pbdriver.DataLinkProtocol_LINKPROTO_NOT_APPLICABLE, nil
	} else {
		return 0, fmt.Errorf("invalid data link protocol %s", proto)
	}
}

func G2RDataLinkProtocol(proto pbdriver.DataLinkProtocol) (DataLinkProtocolSchema, error) {
	if proto == pbdriver.DataLinkProtocol_LINKPROTO_COSEM_WRAPPER {
		return dlProtoCOSEMWRAPPER, nil
	} else if proto == pbdriver.DataLinkProtocol_LINKPROTO_HDLC {
		return dlProtoHDLC, nil
	} else if proto == pbdriver.DataLinkProtocol_LINKPROTO_IEC_62056_21 {
		return dlProtoIEC6205621, nil
	} else if proto == pbdriver.DataLinkProtocol_LINKPROTO_MBUS {
		return dlProtoMBUS, nil
	} else if proto == pbdriver.DataLinkProtocol_LINKPROTO_MODBUS {
		return dlProtoMODBUS, nil
	} else if proto == pbdriver.DataLinkProtocol_LINKPROTO_NOT_APPLICABLE {
		return dlProtoNOTAPPLICABLE, nil
	} else {
		return "", fmt.Errorf("unknown application protocol: %v", proto)
	}
}

func G2RAppProtocolTemplate(appProtocolTemplate *pbdriver.ApplicationProtocolTemplate) (*DriverAppProtocolSchema, error) {
	if appProtocolTemplate == nil {
		return nil, nil
	}

	result := &DriverAppProtocolSchema{}

	var err error
	result.Protocol, err = G2RAppProtocol(appProtocolTemplate.Protocol)
	if err != nil {
		return nil, err
	}

	if ap_attr_cnt := len(appProtocolTemplate.Attributes); ap_attr_cnt > 0 {
		attrs := make([]AttributeDefinitionSchema, 0, ap_attr_cnt)
		result.Attributes = &attrs
		for _, attr := range appProtocolTemplate.Attributes {
			attr_def, err := G2RAttributeDefinition(attr)
			if err != nil {
				return nil, err
			}
			if attr_def == nil {
				continue
			}
			attrs = append(attrs, *attr_def)
		}
	}

	return result, nil
}
