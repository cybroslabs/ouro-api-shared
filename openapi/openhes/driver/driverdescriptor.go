package driver

// DriverDescriptor is a struct to hold the driver descriptor. It is rendered out
// when the driver is started with --descriptor=<format> argument.
// The <format> should be json, but it's reserved for future use.
type DriverDescriptor struct {
	// DriverDescriptor is a base64 encoded value protobuf of the driver descriptor, see pbdriver.NegotiateRequest.
	Descriptor string `json:"driver"`
}
