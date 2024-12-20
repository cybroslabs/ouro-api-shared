package driver

import (
	"fmt"

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

	if dv := attrDef.DefaultValue; dv != nil {
		var tmp interface{}
		switch dvx := dv.Value.(type) {
		case *pbdriver.AttributeValue_StrValue:
			tmp = dvx.StrValue
		case *pbdriver.AttributeValue_IntValue:
			tmp = dvx.IntValue
		case *pbdriver.AttributeValue_DoubleValue:
			tmp = dvx.DoubleValue
		case *pbdriver.AttributeValue_BinaryValue:
			tmp = dvx.BinaryValue
		case *pbdriver.AttributeValue_BoolValue:
			tmp = dvx.BoolValue
		case nil:
			tmp = nil
		default:
			return nil, fmt.Errorf("unknown default value type: %v", dv.Value)
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
		lp := dl_data.LinkProtocol.String()
		datalink_tpls[dl_idx].LinkProtocol = &lp
		if dl_app_cnt := len(dl_data.AppProtocolRefs); dl_app_cnt > 0 {
			// Let's use the same slice for the app protocols
			datalink_tpls[dl_idx].AppProtocolRefs = &dl_data.AppProtocolRefs
		}
	}

	return result, nil
}

func G2RAppProtocolTemplate(appProtocolTemplate *pbdriver.ApplicationProtocolTemplate) (*DriverAppProtocolSchema, error) {
	if appProtocolTemplate == nil {
		return nil, nil
	}

	ap := appProtocolTemplate.Protocol.String()
	result := &DriverAppProtocolSchema{
		Protocol: &ap,
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
