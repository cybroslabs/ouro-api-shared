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
				Name:         "Version",
				Description:  "Version of the COSEM Wrapper",
				Type:         pbdriver.AttributeType_ATTRIBUTE_TYPE_INT,
				Mandatory:    true,
				DefaultValue: &default_COSEM_Wrapper_Version,
			},
			{
				Name:        "SourceAddress",
				Description: "Source address of the COSEM Wrapper",
				Type:        pbdriver.AttributeType_ATTRIBUTE_TYPE_INT,
				Mandatory:   true,
			},
			{
				Name:        "DestinationAddress",
				Description: "Destination address of the COSEM Wrapper",
				Type:        pbdriver.AttributeType_ATTRIBUTE_TYPE_INT,
				Mandatory:   true,
			},
		}, nil
	} else if linkProtocol == pbdriver.DataLinkProtocol_LINKPROTO_HDLC {
		// HDLC attributes
		return []pbdriver.AttributeDefinition{
			{
				Name:        "ClientAddress",
				Description: "Client address",
				Type:        pbdriver.AttributeType_ATTRIBUTE_TYPE_INT,
				Mandatory:   true,
			},
			{
				Name:        "ServerAddress",
				Description: "Server address",
				Type:        pbdriver.AttributeType_ATTRIBUTE_TYPE_INT,
				Mandatory:   true,
			},
		}, nil
	} else if linkProtocol == pbdriver.DataLinkProtocol_LINKPROTO_IEC_62056_21 {
		// IEC 62056-21 attributes
		return []pbdriver.AttributeDefinition{
			{
				Name:        "SerialNumber",
				Description: "Serial number",
				Type:        pbdriver.AttributeType_ATTRIBUTE_TYPE_STRING,
				Mandatory:   true,
			},
		}, nil
	} else {
		return nil, fmt.Errorf("unknown link protocol '%v'", linkProtocol)
	}
}
