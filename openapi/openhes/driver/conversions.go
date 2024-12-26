package driver

import (
	"fmt"

	"github.com/cybroslabs/hes-2-apis/openapi/openhes/job"
	"github.com/cybroslabs/hes-2-apis/protobuf/pbdriver"
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

func G2RCommunicationTemplate(commTemp *pbdriver.CommunicationTemplate) (*DriverCommunicationTemplateSchema, error) {
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
		dl_out.Attributes = &attr_out

		lp := dl_data.LinkProtocol.String()
		dl_out.LinkProtocol = &lp
		if dl_app_cnt := len(dl_data.AppProtocolRefs); dl_app_cnt > 0 {
			// Let's use the same slice for the app protocols
			dl_out.AppProtocolRefs = &dl_data.AppProtocolRefs
		}
	}

	return result, nil
}

var (
	appProtoANSIC12    = job.APPPROTOANSIC12
	appProtoDLMSLN     = job.APPPROTODLMSLN
	appProtoDLMSSN     = job.APPPROTODLMSSN
	appProtoIEC6205621 = job.APPPROTOIEC6205621
	appProtoLIS200     = job.APPPROTOLIS200
	appProtoMQTT       = job.APPPROTOMQTT
	appProtoSCTM       = job.APPPROTOSCTM
)

func G2RAppProtocolTemplate(appProtocolTemplate *pbdriver.ApplicationProtocolTemplate) (*DriverAppProtocolSchema, error) {
	if appProtocolTemplate == nil {
		return nil, nil
	}

	result := &DriverAppProtocolSchema{}

	var proto = appProtocolTemplate.Protocol
	if proto == pbdriver.ApplicationProtocol_APPPROTO_DLMS_LN {
		result.Protocol = &appProtoDLMSLN
	} else if proto == pbdriver.ApplicationProtocol_APPPROTO_DLMS_SN {
		result.Protocol = &appProtoDLMSSN
	} else if proto == pbdriver.ApplicationProtocol_APPPROTO_IEC_62056_21 {
		result.Protocol = &appProtoIEC6205621
	} else if proto == pbdriver.ApplicationProtocol_APPPROTO_LIS200 {
		result.Protocol = &appProtoLIS200
	} else if proto == pbdriver.ApplicationProtocol_APPPROTO_MQTT {
		result.Protocol = &appProtoMQTT
	} else if proto == pbdriver.ApplicationProtocol_APPPROTO_SCTM {
		result.Protocol = &appProtoSCTM
	} else if proto == pbdriver.ApplicationProtocol_APPPROTO_ANSI_C12 {
		result.Protocol = &appProtoANSIC12
	} else {
		return nil, fmt.Errorf("unknown application protocol: %v", proto)
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
