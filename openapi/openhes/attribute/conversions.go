package attribute

import (
	"encoding/json"
	"fmt"
	"log"
	"reflect"

	"github.com/cybroslabs/hes-2-apis/protobuf/pbdriver"
)

var (
	ErrUnknownAttributeType    = fmt.Errorf("unknown attribute type")
	ErrUnknownAPIAttributeType = fmt.Errorf("unknown API attribute type")
)

func R2GAttributes(attrs Attributes) (map[string]*pbdriver.AttributeValue, error) {
	apiAttrs := make(map[string]*pbdriver.AttributeValue)
	for k, v := range attrs {

		av, err := _R2GAttributeValue(v)
		if err != nil {
			return nil, err
		}

		apiAttrs[k] = av
	}
	return apiAttrs, nil
}

func _R2GAttributeValue(attr Value) (*pbdriver.AttributeValue, error) {
	t := reflect.TypeOf(attr)

	switch typed_attr := attr.(type) {
	case string:
		tpv := &pbdriver.AttributeValue_StrValue{
			StrValue: typed_attr,
		}
		return &pbdriver.AttributeValue{Value: tpv}, nil

	case json.Number:
		dt, err := typed_attr.Int64()
		if err == nil {
			tpv := &pbdriver.AttributeValue_IntValue{
				IntValue: dt,
			}
			return &pbdriver.AttributeValue{Value: tpv}, nil
		}

		dtf, err := typed_attr.Float64()
		if err == nil {
			tpv := &pbdriver.AttributeValue_DoubleValue{
				DoubleValue: dtf,
			}
			return &pbdriver.AttributeValue{Value: tpv}, nil
		}

		return nil, ErrUnknownAttributeType

	default:
		log.Default().Printf("attribute: %v %v %v", t.Kind(), attr, typed_attr)

		switch t.Kind() {
		case reflect.String:
			tpv := &pbdriver.AttributeValue_StrValue{
				StrValue: attr.(string),
			}

			return &pbdriver.AttributeValue{Value: tpv}, nil

		case reflect.Int64:
			tpv := &pbdriver.AttributeValue_IntValue{
				IntValue: attr.(int64),
			}

			return &pbdriver.AttributeValue{Value: tpv}, nil

		case reflect.Float64:
			tpv := &pbdriver.AttributeValue_DoubleValue{
				DoubleValue: attr.(float64),
			}

			return &pbdriver.AttributeValue{Value: tpv}, nil

		case reflect.Slice:
			if t.Elem().Kind() != reflect.Uint8 {
				return nil, ErrUnknownAttributeType
			}

			tpv := &pbdriver.AttributeValue_BinaryValue{
				BinaryValue: attr.([]byte),
			}

			return &pbdriver.AttributeValue{Value: tpv}, nil

		default:
			return nil, ErrUnknownAttributeType
		}
	}
}

func G2RAttributes(attrs map[string]*pbdriver.AttributeValue) Attributes {
	attributes := make(Attributes)
	for k, v := range attrs {
		attributes[k] = _G2RAttributeValue(v)
	}
	return attributes
}

func _G2RAttributeValue(attr *pbdriver.AttributeValue) Value {
	switch typed_attr := attr.Value.(type) {
	case *pbdriver.AttributeValue_StrValue:
		return typed_attr.StrValue

	case *pbdriver.AttributeValue_IntValue:
		return typed_attr.IntValue

	case *pbdriver.AttributeValue_DoubleValue:
		return typed_attr.DoubleValue

	case *pbdriver.AttributeValue_BinaryValue:
		return typed_attr.BinaryValue

	default:
		log.Default().Printf("attribute: %v", attr)
		return nil
	}
}
