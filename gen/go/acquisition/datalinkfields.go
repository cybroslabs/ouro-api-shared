package acquisition

import (
	"github.com/cybroslabs/hes-2-apis/gen/go/common"
	"k8s.io/utils/ptr"
)

func GetDataLinkFields(dataLinkProtocol DataLinkProtocol) []*common.FieldDescriptor {
	switch dataLinkProtocol {
	case DataLinkProtocol_LINKPROTO_HDLC:
		// HDLC specific fields
		return []*common.FieldDescriptor{
			common.FieldDescriptor_builder{
				FieldId:  ptr.To("negotiate"),
				Label:    ptr.To("Negotiate HDLC frame size."),
				DataType: common.FieldDataType_BOOLEAN.Enum(),
				Required: ptr.To(false),
				DefaultValue: common.FieldValue_builder{
					BoolValue: ptr.To(false),
				}.Build(),
			}.Build(),
		}

	default:
		return []*common.FieldDescriptor{}
	}
}
