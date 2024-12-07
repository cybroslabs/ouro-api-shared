package driver

import "github.com/cybroslabs/hes-2-apis/protobuf/pbdriver"

// DriverDescriptor is a struct to hold the driver descriptor. It is rendered out
// when the driver is started with --descriptor=<format> argument.
// The <format> should be json, but it's reserved for future use.
type DriverDescriptor struct {
	Descriptor pbdriver.NegotiateRequest `json:"driver"`
}
