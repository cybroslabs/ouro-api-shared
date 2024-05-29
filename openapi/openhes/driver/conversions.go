package driver

import (
	"fmt"

	"github.com/cybroslabs/hes-2-apis/protobuf/pbdriver"
)

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

func G2RAttributeDefinition(attrDef *pbdriver.AttributeDefinition) (*AttributeDefinitionSchema, error) {
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
