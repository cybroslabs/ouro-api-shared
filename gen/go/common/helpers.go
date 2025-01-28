package common

import structpb "google.golang.org/protobuf/types/known/structpb"

func (mf *MetadataFields) GetDouble(key string) float64 {
	if mf == nil {
		return 0
	}
	m := mf.GetFields().GetFields()
	if m == nil {
		return 0
	}
	raw, ok := m[key]
	if !ok {
		return 0
	}
	switch v := raw.GetKind().(type) {
	case *structpb.Value_NumberValue:
		return v.NumberValue
	case *structpb.Value_StringValue:
		return 0
	default:
		return 0
	}
}
