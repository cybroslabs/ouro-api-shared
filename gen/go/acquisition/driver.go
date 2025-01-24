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

// Encodes the driver descriptor to a string.
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

// Decodes the driver descriptor from an io reader. Returns the decoded descriptor, encoded base64 proto buf form of a descriptor and an error.
func DecodeDriverDescriptor(data io.ReadCloser) (*Driver, string, error) {
	var descriptor_holder driverDescriptor
	err := json.NewDecoder(data).Decode(&descriptor_holder)
	if err != nil {
		return nil, "", err
	}

	var descriptor *Driver
	descriptor, err = DecodeDriverDescriptorFromString(descriptor_holder.Descriptor)
	if err != nil {
		return nil, "", err
	}

	return descriptor, descriptor_holder.Descriptor, nil
}

// Decodes the driver descriptor from a string. Returns the decoded descriptor and an error.
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
