package acquisition

import (
	"bytes"
	"encoding/base64"
	"encoding/json"
	"io"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"google.golang.org/protobuf/proto"
)

func TestEncodeDriverDescriptor_Success(t *testing.T) {
	version := "1.0.0"
	driverType := "test-driver"
	driver := Driver_builder{
		Spec: DriverSpec_builder{
			Version:    &version,
			DriverType: &driverType,
		}.Build(),
	}.Build()

	encoded, err := EncodeDriverDescriptor(driver)
	require.NoError(t, err)
	assert.NotEmpty(t, encoded)

	// Verify it's valid JSON
	var holder driverDescriptor
	err = json.Unmarshal([]byte(encoded), &holder)
	require.NoError(t, err)
	assert.NotEmpty(t, holder.Descriptor)

	// Verify the base64 content can be decoded
	decodedBytes, err := base64.URLEncoding.DecodeString(holder.Descriptor)
	require.NoError(t, err)
	assert.NotEmpty(t, decodedBytes)

	// Verify it unmarshals to a valid Driver
	var decodedDriver Driver
	err = proto.Unmarshal(decodedBytes, &decodedDriver)
	require.NoError(t, err)
	assert.Equal(t, version, decodedDriver.GetSpec().GetVersion())
	assert.Equal(t, driverType, decodedDriver.GetSpec().GetDriverType())
}

func TestEncodeDriverDescriptor_NilDescriptor(t *testing.T) {
	encoded, err := EncodeDriverDescriptor(nil)
	assert.Error(t, err)
	assert.Contains(t, err.Error(), "descriptor is nil")
	assert.Empty(t, encoded)
}

func TestEncodeDriverDescriptor_EmptyDriver(t *testing.T) {
	driver := &Driver{}
	encoded, err := EncodeDriverDescriptor(driver)
	require.NoError(t, err)
	assert.NotEmpty(t, encoded)
}

func TestDecodeDriverDescriptor_Success(t *testing.T) {
	version := "2.0.0"
	driverType := "test-driver"
	originalDriver := Driver_builder{
		Spec: DriverSpec_builder{
			Version:    &version,
			DriverType: &driverType,
		}.Build(),
	}.Build()

	// Encode first
	encoded, err := EncodeDriverDescriptor(originalDriver)
	require.NoError(t, err)

	// Decode using io.Reader
	reader := io.NopCloser(strings.NewReader(encoded))
	decodedDriver, descriptorStr, err := DecodeDriverDescriptor(reader)

	require.NoError(t, err)
	assert.NotNil(t, decodedDriver)
	assert.NotEmpty(t, descriptorStr)
	assert.Equal(t, version, decodedDriver.GetSpec().GetVersion())
	assert.Equal(t, driverType, decodedDriver.GetSpec().GetDriverType())
}

func TestDecodeDriverDescriptor_InvalidJSON(t *testing.T) {
	reader := io.NopCloser(strings.NewReader("not valid json"))
	driver, descriptorStr, err := DecodeDriverDescriptor(reader)

	assert.Error(t, err)
	assert.Nil(t, driver)
	assert.Empty(t, descriptorStr)
}

func TestDecodeDriverDescriptor_MissingDescriptorField(t *testing.T) {
	invalidJSON := `{"wrong_field": "value"}`
	reader := io.NopCloser(strings.NewReader(invalidJSON))
	driver, descriptorStr, err := DecodeDriverDescriptor(reader)

	// Empty base64 string decodes to empty but valid protobuf
	require.NoError(t, err)
	assert.NotNil(t, driver)
	assert.Empty(t, descriptorStr)
}

func TestDecodeDriverDescriptor_InvalidBase64(t *testing.T) {
	invalidJSON := `{"driver": "not-valid-base64!@#$"}`
	reader := io.NopCloser(strings.NewReader(invalidJSON))
	driver, descriptorStr, err := DecodeDriverDescriptor(reader)

	assert.Error(t, err)
	assert.Nil(t, driver)
	assert.NotEmpty(t, descriptorStr) // The base64 string is returned even on error
}

func TestDecodeDriverDescriptor_InvalidProtobuf(t *testing.T) {
	// Create valid base64 but invalid protobuf content
	invalidProtobuf := base64.URLEncoding.EncodeToString([]byte("invalid protobuf data"))
	invalidJSON := `{"driver": "` + invalidProtobuf + `"}`
	reader := io.NopCloser(strings.NewReader(invalidJSON))
	driver, descriptorStr, err := DecodeDriverDescriptor(reader)

	assert.Error(t, err)
	assert.Nil(t, driver)
	assert.Equal(t, invalidProtobuf, descriptorStr)
}

func TestDecodeDriverDescriptorFromString_Success(t *testing.T) {
	version := "3.0.0"
	driverType := "string-test-driver"
	originalDriver := Driver_builder{
		Spec: DriverSpec_builder{
			Version:    &version,
			DriverType: &driverType,
		}.Build(),
	}.Build()

	// Manually create base64 encoded protobuf
	protoBytes, err := proto.Marshal(originalDriver)
	require.NoError(t, err)
	base64Str := base64.URLEncoding.EncodeToString(protoBytes)

	// Decode from string
	decodedDriver, err := DecodeDriverDescriptorFromString(base64Str)
	require.NoError(t, err)
	assert.NotNil(t, decodedDriver)
	assert.Equal(t, version, decodedDriver.GetSpec().GetVersion())
	assert.Equal(t, driverType, decodedDriver.GetSpec().GetDriverType())
}

func TestDecodeDriverDescriptorFromString_EmptyString(t *testing.T) {
	driver, err := DecodeDriverDescriptorFromString("")
	// Empty string decodes to empty but valid protobuf
	require.NoError(t, err)
	assert.NotNil(t, driver)
}

func TestDecodeDriverDescriptorFromString_InvalidBase64(t *testing.T) {
	driver, err := DecodeDriverDescriptorFromString("not-valid-base64!@#$")
	assert.Error(t, err)
	assert.Nil(t, driver)
}

func TestDecodeDriverDescriptorFromString_InvalidProtobuf(t *testing.T) {
	invalidBase64 := base64.URLEncoding.EncodeToString([]byte("not a protobuf message"))
	driver, err := DecodeDriverDescriptorFromString(invalidBase64)
	assert.Error(t, err)
	assert.Nil(t, driver)
}

func TestRoundTrip_EncodeAndDecode(t *testing.T) {
	version := "4.0.0"
	driverType := "roundtrip-driver"
	originalDriver := Driver_builder{
		Spec: DriverSpec_builder{
			Version:    &version,
			DriverType: &driverType,
		}.Build(),
	}.Build()

	// Encode
	encoded, err := EncodeDriverDescriptor(originalDriver)
	require.NoError(t, err)

	// Decode via reader
	reader := io.NopCloser(strings.NewReader(encoded))
	decodedDriver1, _, err := DecodeDriverDescriptor(reader)
	require.NoError(t, err)
	assert.Equal(t, version, decodedDriver1.GetSpec().GetVersion())
	assert.Equal(t, driverType, decodedDriver1.GetSpec().GetDriverType())

	// Decode via string
	var holder driverDescriptor
	err = json.Unmarshal([]byte(encoded), &holder)
	require.NoError(t, err)

	decodedDriver2, err := DecodeDriverDescriptorFromString(holder.Descriptor)
	require.NoError(t, err)
	assert.Equal(t, version, decodedDriver2.GetSpec().GetVersion())
	assert.Equal(t, driverType, decodedDriver2.GetSpec().GetDriverType())
}

func TestDriverDescriptor_JSONStructure(t *testing.T) {
	version := "5.0.0"
	driverType := "json-test-driver"
	driver := Driver_builder{
		Spec: DriverSpec_builder{
			Version:    &version,
			DriverType: &driverType,
		}.Build(),
	}.Build()

	encoded, err := EncodeDriverDescriptor(driver)
	require.NoError(t, err)

	// Parse and verify JSON structure
	var result map[string]interface{}
	err = json.Unmarshal([]byte(encoded), &result)
	require.NoError(t, err)

	// Should have exactly one field named "driver"
	assert.Len(t, result, 1)
	assert.Contains(t, result, "driver")
	assert.IsType(t, "", result["driver"])
}

func TestDecodeDriverDescriptor_EmptyReader(t *testing.T) {
	reader := io.NopCloser(bytes.NewReader([]byte{}))
	driver, descriptorStr, err := DecodeDriverDescriptor(reader)

	assert.Error(t, err)
	assert.Nil(t, driver)
	assert.Empty(t, descriptorStr)
}

func TestDecodeDriverDescriptorFromString_ValidEmptyDriver(t *testing.T) {
	// Create an empty but valid Driver protobuf
	emptyDriver := &Driver{}
	protoBytes, err := proto.Marshal(emptyDriver)
	require.NoError(t, err)
	base64Str := base64.URLEncoding.EncodeToString(protoBytes)

	// Should decode successfully
	decodedDriver, err := DecodeDriverDescriptorFromString(base64Str)
	require.NoError(t, err)
	assert.NotNil(t, decodedDriver)
	// Empty driver will have nil spec
	if decodedDriver.GetSpec() != nil {
		assert.Empty(t, decodedDriver.GetSpec().GetVersion())
		assert.Empty(t, decodedDriver.GetSpec().GetDriverType())
	}
}

// Benchmark tests
func BenchmarkEncodeDriverDescriptor(b *testing.B) {
	version := "1.0.0"
	driverType := "benchmark-driver"
	driver := Driver_builder{
		Spec: DriverSpec_builder{
			Version:    &version,
			DriverType: &driverType,
		}.Build(),
	}.Build()
	for b.Loop() {
		_, _ = EncodeDriverDescriptor(driver)
	}
}

func BenchmarkDecodeDriverDescriptor(b *testing.B) {
	version := "1.0.0"
	driverType := "benchmark-driver"
	driver := Driver_builder{
		Spec: DriverSpec_builder{
			Version:    &version,
			DriverType: &driverType,
		}.Build(),
	}.Build()
	encoded, _ := EncodeDriverDescriptor(driver)
	for b.Loop() {
		reader := io.NopCloser(strings.NewReader(encoded))
		_, _, _ = DecodeDriverDescriptor(reader)
	}
}

func BenchmarkDecodeDriverDescriptorFromString(b *testing.B) {
	version := "1.0.0"
	driverType := "benchmark-driver"
	driver := Driver_builder{
		Spec: DriverSpec_builder{
			Version:    &version,
			DriverType: &driverType,
		}.Build(),
	}.Build()
	protoBytes, _ := proto.Marshal(driver)
	base64Str := base64.URLEncoding.EncodeToString(protoBytes)
	for b.Loop() {
		_, _ = DecodeDriverDescriptorFromString(base64Str)
	}
}
