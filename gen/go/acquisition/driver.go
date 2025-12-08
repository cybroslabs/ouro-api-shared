package acquisition

import (
	"encoding/base64"
	"encoding/json"
	"errors"
	"io"

	"google.golang.org/protobuf/proto"
)

// DriverDescriptor is a struct to hold the driver descriptor. It is rendered out
// when the driver is started with --descriptor=<format> argument.
// The <format> should be json, but it's reserved for future use.
type driverDescriptor struct {
	// DriverDescriptor is a base64 encoded value protobuf of the driver descriptor, see pbdriver.NegotiateRequest.
	Descriptor string `json:"driver"`
}

// EncodeDriverDescriptor encodes the driver descriptor to a JSON string.
// The driver descriptor is serialized as a protobuf message, base64-encoded, and wrapped in a JSON object.
// This format is used when drivers report their capabilities during initialization.
func EncodeDriverDescriptor(descriptor *Driver) (string, error) {
	if descriptor == nil {
		return "", errors.New("descriptor is nil")
	}

	descriptor_holder := driverDescriptor{}
	if m, err := proto.Marshal(descriptor); err != nil {
		return "", err
	} else {
		descriptor_holder.Descriptor = base64.URLEncoding.EncodeToString(m)
	}
	json, err := json.Marshal(&descriptor_holder)
	return string(json), err
}

// DecodeDriverDescriptor decodes the driver descriptor from an io.Reader.
// It parses the JSON wrapper, extracts the base64-encoded protobuf, and deserializes it into a Driver object.
// Returns the decoded descriptor, the original base64-encoded string, and any error encountered.
func DecodeDriverDescriptor(data io.ReadCloser) (*Driver, string, error) {
	var descriptor_holder driverDescriptor
	err := json.NewDecoder(data).Decode(&descriptor_holder)
	if err != nil {
		return nil, "", err
	}

	var descriptor *Driver
	descriptor, err = DecodeDriverDescriptorFromString(descriptor_holder.Descriptor)
	if err != nil {
		return nil, descriptor_holder.Descriptor, err
	}

	return descriptor, descriptor_holder.Descriptor, nil
}

// DecodeDriverDescriptorFromString decodes the driver descriptor from a base64-encoded string.
// It decodes the base64 string and deserializes the protobuf message into a Driver object.
// Returns the decoded descriptor and any error encountered during decoding or deserialization.
func DecodeDriverDescriptorFromString(descriptor string) (*Driver, error) {
	bin, err := base64.URLEncoding.DecodeString(descriptor)
	if err != nil {
		return nil, err
	}

	result := &Driver{}
	err = proto.Unmarshal(bin, result)
	if err != nil {
		return nil, err
	}

	return result, nil
}
