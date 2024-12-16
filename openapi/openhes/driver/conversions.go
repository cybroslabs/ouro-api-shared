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
		if dl_app_cnt := len(dl_data.AppProtocols); dl_app_cnt > 0 {
			dl_app_protocol := make([]DriverAppProtocolSchema, dl_app_cnt)
			datalink_tpls[dl_idx].AppProtocols = &dl_app_protocol
			for ap_idx, ap_data := range dl_data.AppProtocols {
				if ap_data == nil {
					return nil, fmt.Errorf("app protocol contains nil")
				}
				ap := ap_data.Protocol.String()
				dl_app_protocol[ap_idx].Protocol = &ap
				if ap_attr_cnt := len(ap_data.Attributes); ap_attr_cnt > 0 {
					attrs := make([]AttributeDefinitionSchema, 0, ap_attr_cnt)
					dl_app_protocol[ap_idx].Profile = &attrs
					for _, attr := range ap_data.Attributes {
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
			}
		}
	}

	return result, nil
}
