package driver

import (
	"fmt"

	"github.com/cybroslabs/hes-2-apis/protobuf/pbdriver"
)

var (
	default_COSEM_Wrapper_Version = pbdriver.AttributeValue{Value: &pbdriver.AttributeValue_IntValue{IntValue: 1}}
)

func GetDataLinkAttributes(linkProtocol pbdriver.DataLinkProtocol) ([]pbdriver.AttributeDefinition, error) {
	if linkProtocol == pbdriver.DataLinkProtocol_LINKPROTO_COSEM_WRAPPER {
		// COSEM Wrapper attributes
		return []pbdriver.AttributeDefinition{
			{
				Name:        "cw_source_address",
				Description: "Source address",
				Type:        pbdriver.AttributeType_ATTRIBUTE_TYPE_INT,
				Mandatory:   true,
			},
			{
				Name:        "cw_destination_address",
				Description: "Destination address",
				Type:        pbdriver.AttributeType_ATTRIBUTE_TYPE_INT,
				Mandatory:   true,
			},
		}, nil
	} else if linkProtocol == pbdriver.DataLinkProtocol_LINKPROTO_HDLC {
		// HDLC attributes
		return []pbdriver.AttributeDefinition{
			{
				Name:        "hdlc_client_address",
				Description: "Client address",
				Type:        pbdriver.AttributeType_ATTRIBUTE_TYPE_INT,
				Mandatory:   true,
			},
			{
				Name:        "hdlc_server_address",
				Description: "Server address",
				Type:        pbdriver.AttributeType_ATTRIBUTE_TYPE_INT,
				Mandatory:   true,
			},
		}, nil
	} else if linkProtocol == pbdriver.DataLinkProtocol_LINKPROTO_IEC_62056_21 {
		// IEC 62056-21 attributes
		return []pbdriver.AttributeDefinition{
			{
				Name:        "iec62056_client_address",
				Description: "IEC serial",
				Type:        pbdriver.AttributeType_ATTRIBUTE_TYPE_STRING,
				Mandatory:   false,
			},
		}, nil
	} else if linkProtocol == pbdriver.DataLinkProtocol_LINKPROTO_NOT_APPLICABLE {
		return []pbdriver.AttributeDefinition{}, nil
	} else {
		return nil, fmt.Errorf("unknown link protocol '%v'", linkProtocol)
	}
}
