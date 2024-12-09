package driver

import (
	"errors"
	"fmt"

	"github.com/cybroslabs/hes-2-apis/protobuf/pbdriver"
)

var (
	ErrUnknownDataLinkProtocol = errors.Join(errors.ErrUnsupported, errors.New("unknown data link protocol"))
)

// GetDataLinkAttributes returns the attributes for the given data link protocol.
// Those are typically device addresses or other configuration parameters.
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
		return nil, errors.Join(ErrUnknownDataLinkProtocol, fmt.Errorf("protocol '%v'", linkProtocol))
	}
}
