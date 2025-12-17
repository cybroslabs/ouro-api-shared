package acquisition

import (
	"bytes"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestProfileValuesEncoder_NewAndBasics(t *testing.T) {
	encoder := NewProfileValuesEncoder(60, "celsius")
	assert.NotNil(t, encoder)
	assert.Equal(t, int32(60), encoder.periodseconds)
	assert.Equal(t, "celsius", encoder.unit)

	// Empty encoder should produce header only
	data := encoder.Bytes()
	assert.Greater(t, len(data), 0)

	// Verify we can decode the header
	decoder, err := NewProfileValuesDecoder(data)
	require.NoError(t, err)
	assert.Equal(t, time.Duration(60)*time.Second, decoder.GetPeriod())
	assert.Equal(t, "celsius", decoder.GetUnit())
}

func TestProfileValuesEncoder_Reset(t *testing.T) {
	encoder := NewProfileValuesEncoder(60, "celsius")
	ts := time.Date(2024, 1, 1, 0, 0, 0, 0, time.UTC)
	encoder.AppendDouble(ts, 0, 0, 0, 23.5, nil)

	// Reset with different parameters
	encoder.Reset(30, "fahrenheit")
	assert.Equal(t, int32(30), encoder.periodseconds)
	assert.Equal(t, "fahrenheit", encoder.unit)
	assert.Equal(t, 0, encoder.items)

	// Should be able to encode new data
	encoder.AppendDouble(ts, 0, 0, 0, 75.2, nil)
	data := encoder.Bytes()

	decoder, err := NewProfileValuesDecoder(data)
	require.NoError(t, err)
	assert.Equal(t, int32(30), decoder.GetPeriodSeconds())
	assert.Equal(t, "fahrenheit", decoder.GetUnit())
}

func TestProfileValuesEncoder_AppendDouble(t *testing.T) {
	encoder := NewProfileValuesEncoder(60, "celsius")
	ts := time.Date(2024, 1, 1, 0, 0, 0, 0, time.UTC)

	encoder.AppendDouble(ts, 100, 200, -2, 23.5, nil)
	encoder.AppendDouble(ts.Add(60*time.Second), 100, 200, -2, 24.7, nil)
	encoder.AppendDouble(ts.Add(120*time.Second), 100, 200, -2, 22.3, nil)

	data := encoder.Bytes()
	assert.Greater(t, len(data), 0)

	// Decode and verify
	decoder, err := NewProfileValuesDecoder(data)
	require.NoError(t, err)

	values := collectValues(t, decoder)
	require.Len(t, values, 3)

	assert.Equal(t, ts, values[0].Timestamp)
	assert.Equal(t, 23.5, values[0].Value.GetDoubleValue())
	assert.Equal(t, int64(100), values[0].Value.GetStatus())
	assert.Equal(t, uint64(200), values[0].Value.GetNstatus())
	assert.Equal(t, int32(-2), values[0].Value.GetExponent())

	assert.Equal(t, 24.7, values[1].Value.GetDoubleValue())
	assert.Equal(t, 22.3, values[2].Value.GetDoubleValue())
}

func TestProfileValuesEncoder_AppendInteger(t *testing.T) {
	tests := []struct {
		name  string
		value int64
	}{
		{"zero", 0},
		{"small positive", 42},
		{"small negative", -42},
		{"int8 max", 127},
		{"int8 min", -128},
		{"int16 max", 32767},
		{"int16 min", -32768},
		{"int32 max", 2147483647},
		{"int32 min", -2147483648},
		{"int64 large", 9223372036854775807},
		{"int64 large negative", -9223372036854775808},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			encoder := NewProfileValuesEncoder(60, "count")
			ts := time.Date(2024, 1, 1, 0, 0, 0, 0, time.UTC)

			encoder.AppendInteger(ts, 0, 0, 0, tt.value, nil)
			data := encoder.Bytes()

			decoder, err := NewProfileValuesDecoder(data)
			require.NoError(t, err)

			values := collectValues(t, decoder)
			require.Len(t, values, 1)

			assert.Equal(t, tt.value, values[0].Value.GetIntegerValue())
		})
	}
}

func TestProfileValuesEncoder_AppendInteger_PreviousValueOptimization(t *testing.T) {
	encoder := NewProfileValuesEncoder(60, "count")
	ts := time.Date(2024, 1, 1, 0, 0, 0, 0, time.UTC)

	// Add same value multiple times - should use previous value optimization
	encoder.AppendInteger(ts, 0, 0, 0, 42, nil)
	encoder.AppendInteger(ts.Add(60*time.Second), 0, 0, 0, 42, nil)
	encoder.AppendInteger(ts.Add(120*time.Second), 0, 0, 0, 42, nil)

	data := encoder.Bytes()

	decoder, err := NewProfileValuesDecoder(data)
	require.NoError(t, err)

	values := collectValues(t, decoder)
	require.Len(t, values, 3)

	for i, v := range values {
		assert.Equal(t, int64(42), v.Value.GetIntegerValue(), "value %d", i)
	}
}

func TestProfileValuesEncoder_AppendString(t *testing.T) {
	tests := []struct {
		name  string
		value string
	}{
		{"empty", ""},
		{"short", "ok"},
		{"medium", "this is a medium length string"},
		{"long", string(bytes.Repeat([]byte("a"), 300))},
		{"unicode", "Hello 世界 🌍"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			encoder := NewProfileValuesEncoder(60, "status")
			ts := time.Date(2024, 1, 1, 0, 0, 0, 0, time.UTC)

			encoder.AppendString(ts, 0, 0, 0, tt.value, nil)
			data := encoder.Bytes()

			decoder, err := NewProfileValuesDecoder(data)
			require.NoError(t, err)

			values := collectValues(t, decoder)
			require.Len(t, values, 1)

			assert.Equal(t, tt.value, values[0].Value.GetStringValue())
		})
	}
}

func TestProfileValuesEncoder_AppendString_PreviousValueOptimization(t *testing.T) {
	encoder := NewProfileValuesEncoder(60, "status")
	ts := time.Date(2024, 1, 1, 0, 0, 0, 0, time.UTC)

	// Add same string multiple times
	encoder.AppendString(ts, 0, 0, 0, "active", nil)
	encoder.AppendString(ts.Add(60*time.Second), 0, 0, 0, "active", nil)
	encoder.AppendString(ts.Add(120*time.Second), 0, 0, 0, "inactive", nil)

	data := encoder.Bytes()

	decoder, err := NewProfileValuesDecoder(data)
	require.NoError(t, err)

	values := collectValues(t, decoder)
	require.Len(t, values, 3)

	assert.Equal(t, "active", values[0].Value.GetStringValue())
	assert.Equal(t, "active", values[1].Value.GetStringValue())
	assert.Equal(t, "inactive", values[2].Value.GetStringValue())
}

func TestProfileValuesEncoder_AppendTimestamp(t *testing.T) {
	encoder := NewProfileValuesEncoder(60, "")
	ts := time.Date(2024, 1, 1, 0, 0, 0, 0, time.UTC)
	valueTs := time.Date(2024, 6, 15, 12, 30, 45, 0, time.UTC)

	encoder.AppendTimestamp(ts, 0, 0, 0, valueTs, nil)
	data := encoder.Bytes()

	decoder, err := NewProfileValuesDecoder(data)
	require.NoError(t, err)

	values := collectValues(t, decoder)
	require.Len(t, values, 1)

	assert.NotNil(t, values[0].Value.GetTimestampValue())
	assert.Equal(t, valueTs.Unix(), values[0].Value.GetTimestampValue().Seconds)
}

func TestProfileValuesEncoder_AppendTimestampWithTz(t *testing.T) {
	encoder := NewProfileValuesEncoder(60, "")
	ts := time.Date(2024, 1, 1, 0, 0, 0, 0, time.UTC)
	valueTs := time.Date(2024, 6, 15, 12, 30, 45, 0, time.UTC)

	encoder.AppendTimestampWithTz(ts, 0, 0, 0, valueTs, nil)
	data := encoder.Bytes()

	decoder, err := NewProfileValuesDecoder(data)
	require.NoError(t, err)

	values := collectValues(t, decoder)
	require.Len(t, values, 1)

	expectedStr := valueTs.Format(time.RFC3339)
	assert.Equal(t, expectedStr, values[0].Value.GetTimestampTzValue())
}

func TestProfileValuesEncoder_AppendBoolean(t *testing.T) {
	encoder := NewProfileValuesEncoder(60, "")
	ts := time.Date(2024, 1, 1, 0, 0, 0, 0, time.UTC)

	encoder.AppendBoolean(ts, 0, 0, 0, true, nil)
	encoder.AppendBoolean(ts.Add(60*time.Second), 0, 0, 0, false, nil)
	encoder.AppendBoolean(ts.Add(120*time.Second), 0, 0, 0, true, nil)

	data := encoder.Bytes()

	decoder, err := NewProfileValuesDecoder(data)
	require.NoError(t, err)

	values := collectValues(t, decoder)
	require.Len(t, values, 3)

	assert.True(t, values[0].Value.GetBoolValue())
	assert.False(t, values[1].Value.GetBoolValue())
	assert.True(t, values[2].Value.GetBoolValue())
}

func TestProfileValuesEncoder_AppendValue_AllTypes(t *testing.T) {
	ts := time.Date(2024, 1, 1, 0, 0, 0, 0, time.UTC)

	tests := []struct {
		value    any
		expected string
	}{
		{float64(23.5), "double"},
		{int64(42), "integer"},
		{"hello", "string"},
		{time.Date(2024, 1, 1, 0, 0, 0, 0, time.UTC), "timestamp"},
		{TimeWithTimeZone{time.Date(2024, 1, 1, 0, 0, 0, 0, time.UTC)}, "timestamptz"},
		{true, "bool"},
	}

	for _, tt := range tests {
		t.Run(tt.expected, func(t *testing.T) {
			encoder := NewProfileValuesEncoder(60, "test")
			err := encoder.AppendValue(ts, 0, 0, 0, tt.value, nil)
			require.NoError(t, err)

			data := encoder.Bytes()
			decoder, err := NewProfileValuesDecoder(data)
			require.NoError(t, err)

			values := collectValues(t, decoder)
			require.Len(t, values, 1)
			assert.NotNil(t, values[0].Value)
		})
	}
}

func TestProfileValuesEncoder_AppendValue_UnknownType(t *testing.T) {
	encoder := NewProfileValuesEncoder(60, "test")
	ts := time.Date(2024, 1, 1, 0, 0, 0, 0, time.UTC)

	err := encoder.AppendValue(ts, 0, 0, 0, struct{}{}, nil)
	assert.Error(t, err)
	assert.Contains(t, err.Error(), "unknown type")
}

func TestProfileValuesEncoder_PeakTimestamp(t *testing.T) {
	encoder := NewProfileValuesEncoder(60, "celsius")
	ts := time.Date(2024, 1, 1, 0, 0, 0, 0, time.UTC)
	peakTs := time.Date(2024, 1, 1, 0, 0, 30, 0, time.UTC)

	encoder.AppendDouble(ts, 0, 0, 0, 23.5, &peakTs)
	data := encoder.Bytes()

	decoder, err := NewProfileValuesDecoder(data)
	require.NoError(t, err)

	values := collectValues(t, decoder)
	require.Len(t, values, 1)

	assert.NotNil(t, values[0].Value.GetPeakTs())
	assert.Equal(t, peakTs.Unix(), values[0].Value.GetPeakTs().Seconds)
}

func TestProfileValuesEncoder_StatusAndNStatus(t *testing.T) {
	encoder := NewProfileValuesEncoder(60, "test")
	ts := time.Date(2024, 1, 1, 0, 0, 0, 0, time.UTC)

	// First value with status
	encoder.AppendDouble(ts, 100, 200, 0, 1.0, nil)
	// Same status - should be optimized
	encoder.AppendDouble(ts.Add(60*time.Second), 100, 200, 0, 2.0, nil)
	// Different status
	encoder.AppendDouble(ts.Add(120*time.Second), 300, 400, 0, 3.0, nil)

	data := encoder.Bytes()

	decoder, err := NewProfileValuesDecoder(data)
	require.NoError(t, err)

	values := collectValues(t, decoder)
	require.Len(t, values, 3)

	assert.Equal(t, int64(100), values[0].Value.GetStatus())
	assert.Equal(t, uint64(200), values[0].Value.GetNstatus())
	assert.Equal(t, int64(100), values[1].Value.GetStatus())
	assert.Equal(t, uint64(200), values[1].Value.GetNstatus())
	assert.Equal(t, int64(300), values[2].Value.GetStatus())
	assert.Equal(t, uint64(400), values[2].Value.GetNstatus())
}

func TestProfileValuesEncoder_Exponent(t *testing.T) {
	encoder := NewProfileValuesEncoder(60, "test")
	ts := time.Date(2024, 1, 1, 0, 0, 0, 0, time.UTC)

	encoder.AppendDouble(ts, 0, 0, -2, 1.0, nil)
	encoder.AppendDouble(ts.Add(60*time.Second), 0, 0, -2, 2.0, nil)
	encoder.AppendDouble(ts.Add(120*time.Second), 0, 0, -3, 3.0, nil)

	data := encoder.Bytes()

	decoder, err := NewProfileValuesDecoder(data)
	require.NoError(t, err)

	values := collectValues(t, decoder)
	require.Len(t, values, 3)

	assert.Equal(t, int32(-2), values[0].Value.GetExponent())
	assert.Equal(t, int32(-2), values[1].Value.GetExponent())
	assert.Equal(t, int32(-3), values[2].Value.GetExponent())
}

func TestProfileValuesEncoder_BlockBoundary(t *testing.T) {
	encoder := NewProfileValuesEncoder(60, "test")
	ts := time.Date(2024, 1, 1, 0, 0, 0, 0, time.UTC)

	// Add 257 values to force block boundary (max 256 items per block)
	for i := range 257 {
		encoder.AppendDouble(ts.Add(time.Duration(i*60)*time.Second), 0, 0, 0, float64(i), nil)
	}

	data := encoder.Bytes()

	decoder, err := NewProfileValuesDecoder(data)
	require.NoError(t, err)

	values := collectValues(t, decoder)
	require.Len(t, values, 257)

	for i := range 257 {
		assert.Equal(t, float64(i), values[i].Value.GetDoubleValue())
	}
}

func TestProfileValuesEncoder_MixedTypes_SeparateBlocks(t *testing.T) {
	encoder := NewProfileValuesEncoder(60, "mixed")
	ts := time.Date(2024, 1, 1, 0, 0, 0, 0, time.UTC)

	// Different types should create different blocks
	encoder.AppendDouble(ts, 0, 0, 0, 1.5, nil)
	encoder.AppendInteger(ts.Add(60*time.Second), 0, 0, 0, 42, nil)
	encoder.AppendString(ts.Add(120*time.Second), 0, 0, 0, "hello", nil)

	data := encoder.Bytes()

	decoder, err := NewProfileValuesDecoder(data)
	require.NoError(t, err)

	values := collectValues(t, decoder)
	require.Len(t, values, 3)

	// Just verify values are set correctly by type
	assert.NotEqual(t, 0.0, values[0].Value.GetDoubleValue())
	assert.NotEqual(t, int64(0), values[1].Value.GetIntegerValue())
	assert.NotEqual(t, "", values[2].Value.GetStringValue())
}

func TestProfileValuesDecoder_EmptyData(t *testing.T) {
	decoder, err := NewProfileValuesDecoder([]byte{})
	require.NoError(t, err)
	assert.True(t, decoder.IsNil())
	assert.Equal(t, time.Duration(0), decoder.GetPeriod())
	assert.Equal(t, "", decoder.GetUnit())

	// Should handle empty iterator
	count := 0
	for range decoder.Values() {
		count++
	}
	assert.Equal(t, 0, count)
}

func TestProfileValuesDecoder_InvalidVersion(t *testing.T) {
	// Need enough data to pass the initial read
	data := []byte{99, 0, 0, 0, 60, 0, 0, 0, 0} // version 99, period, unit length
	_, err := NewProfileValuesDecoder(data)
	assert.Error(t, err)
	assert.Contains(t, err.Error(), "invalid data version")
}

func TestProfileValuesDecoder_GetInfo(t *testing.T) {
	encoder := NewProfileValuesEncoder(60, "celsius")
	ts1 := time.Date(2024, 1, 1, 0, 0, 0, 0, time.UTC)
	ts2 := ts1.Add(60 * time.Second)
	ts3 := ts1.Add(120 * time.Second)

	encoder.AppendDouble(ts1, 0, 0, 0, 1.0, nil)
	encoder.AppendDouble(ts2, 0, 0, 0, 2.0, nil)
	encoder.AppendDouble(ts3, 0, 0, 0, 3.0, nil)

	data := encoder.Bytes()
	decoder, err := NewProfileValuesDecoder(data)
	require.NoError(t, err)

	firstTs, lastTs, count, err := decoder.GetInfo()
	require.NoError(t, err)

	assert.Equal(t, ts1, firstTs)
	assert.Equal(t, ts3, lastTs)
	assert.Equal(t, 3, count)
}

func TestProfileValuesDecoder_GetInfo_EmptyData(t *testing.T) {
	decoder, err := NewProfileValuesDecoder([]byte{})
	require.NoError(t, err)

	firstTs, lastTs, count, err := decoder.GetInfo()
	require.NoError(t, err)

	assert.True(t, firstTs.IsZero())
	assert.True(t, lastTs.IsZero())
	assert.Equal(t, 0, count)
}

func TestProfileValuesDecoder_GetInfo_MultipleBlocks(t *testing.T) {
	encoder := NewProfileValuesEncoder(60, "test")
	ts1 := time.Date(2024, 1, 1, 0, 0, 0, 0, time.UTC)

	// Create multiple blocks by changing types
	encoder.AppendDouble(ts1, 0, 0, 0, 1.0, nil)
	encoder.AppendDouble(ts1.Add(60*time.Second), 0, 0, 0, 2.0, nil)
	encoder.AppendInteger(ts1.Add(120*time.Second), 0, 0, 0, 3, nil)
	encoder.AppendInteger(ts1.Add(180*time.Second), 0, 0, 0, 4, nil)

	data := encoder.Bytes()
	decoder, err := NewProfileValuesDecoder(data)
	require.NoError(t, err)

	firstTs, lastTs, count, err := decoder.GetInfo()
	require.NoError(t, err)

	assert.Equal(t, ts1, firstTs)
	assert.Equal(t, ts1.Add(180*time.Second), lastTs)
	assert.Equal(t, 4, count)
}

func TestMergeProfileBlobs_BothEmpty(t *testing.T) {
	dst := make([]byte, 0)
	result, err := MergeProfileBlobs(dst, []byte{}, []byte{})
	require.NoError(t, err)
	assert.Empty(t, result)
}

func TestMergeProfileBlobs_FirstEmpty(t *testing.T) {
	encoder := NewProfileValuesEncoder(60, "test")
	ts := time.Date(2024, 1, 1, 0, 0, 0, 0, time.UTC)
	encoder.AppendDouble(ts, 0, 0, 0, 1.0, nil)
	data2 := encoder.Bytes()

	dst := make([]byte, 0)
	result, err := MergeProfileBlobs(dst, []byte{}, data2)
	require.NoError(t, err)
	assert.Equal(t, data2, result)
}

func TestMergeProfileBlobs_SecondEmpty(t *testing.T) {
	encoder := NewProfileValuesEncoder(60, "test")
	ts := time.Date(2024, 1, 1, 0, 0, 0, 0, time.UTC)
	encoder.AppendDouble(ts, 0, 0, 0, 1.0, nil)
	data1 := encoder.Bytes()

	dst := make([]byte, 0)
	result, err := MergeProfileBlobs(dst, data1, []byte{})
	require.NoError(t, err)
	assert.Equal(t, data1, result)
}

func TestMergeProfileBlobs_Success(t *testing.T) {
	encoder1 := NewProfileValuesEncoder(60, "test")
	ts := time.Date(2024, 1, 1, 0, 0, 0, 0, time.UTC)
	encoder1.AppendDouble(ts, 0, 0, 0, 1.0, nil)
	data1 := encoder1.Bytes()

	encoder2 := NewProfileValuesEncoder(60, "test")
	encoder2.AppendDouble(ts.Add(60*time.Second), 0, 0, 0, 2.0, nil)
	data2 := encoder2.Bytes()

	dst := make([]byte, 0)
	result, err := MergeProfileBlobs(dst, data1, data2)
	require.NoError(t, err)

	decoder, err := NewProfileValuesDecoder(result)
	require.NoError(t, err)

	values := collectValues(t, decoder)
	require.Len(t, values, 2)
	assert.Equal(t, 1.0, values[0].Value.GetDoubleValue())
	assert.Equal(t, 2.0, values[1].Value.GetDoubleValue())
}

func TestMergeProfileBlobs_DifferentPeriod(t *testing.T) {
	encoder1 := NewProfileValuesEncoder(60, "test")
	ts := time.Date(2024, 1, 1, 0, 0, 0, 0, time.UTC)
	encoder1.AppendDouble(ts, 0, 0, 0, 1.0, nil)
	data1 := encoder1.Bytes()

	encoder2 := NewProfileValuesEncoder(30, "test")
	encoder2.AppendDouble(ts, 0, 0, 0, 2.0, nil)
	data2 := encoder2.Bytes()

	dst := make([]byte, 0)
	_, err := MergeProfileBlobs(dst, data1, data2)
	assert.Error(t, err)
	assert.Contains(t, err.Error(), "different period")
}

func TestMergeProfileBlobs_DifferentUnit(t *testing.T) {
	encoder1 := NewProfileValuesEncoder(60, "celsius")
	ts := time.Date(2024, 1, 1, 0, 0, 0, 0, time.UTC)
	encoder1.AppendDouble(ts, 0, 0, 0, 1.0, nil)
	data1 := encoder1.Bytes()

	encoder2 := NewProfileValuesEncoder(60, "fahrenheit")
	encoder2.AppendDouble(ts, 0, 0, 0, 2.0, nil)
	data2 := encoder2.Bytes()

	dst := make([]byte, 0)
	_, err := MergeProfileBlobs(dst, data1, data2)
	assert.Error(t, err)
	assert.Contains(t, err.Error(), "different unit")
}

func TestProfileValuesDecoder_CorruptData(t *testing.T) {
	encoder := NewProfileValuesEncoder(60, "test")
	ts := time.Date(2024, 1, 1, 0, 0, 0, 0, time.UTC)
	encoder.AppendDouble(ts, 0, 0, 0, 1.0, nil)
	data := encoder.Bytes()

	// Corrupt the data by truncating it
	corruptData := data[:len(data)/2]

	decoder, err := NewProfileValuesDecoder(corruptData)
	if err != nil {
		// Corruption detected during decoding
		return
	}

	// Should get error when iterating
	hasError := false
	for item := range decoder.Values() {
		if item.Err != nil {
			hasError = true
			break
		}
	}
	assert.True(t, hasError, "Expected error when reading corrupt data")
}

func TestProfileValuesEncoder_LargeDataset(t *testing.T) {
	encoder := NewProfileValuesEncoder(1, "test")
	ts := time.Date(2024, 1, 1, 0, 0, 0, 0, time.UTC)

	// Add 10000 values
	count := 10000
	for i := range count {
		encoder.AppendDouble(ts.Add(time.Duration(i)*time.Second), 0, 0, 0, float64(i), nil)
	}

	data := encoder.Bytes()
	assert.Greater(t, len(data), 0)

	decoder, err := NewProfileValuesDecoder(data)
	require.NoError(t, err)

	values := collectValues(t, decoder)
	assert.Equal(t, count, len(values))

	// Spot check some values
	assert.Equal(t, 0.0, values[0].Value.GetDoubleValue())
	assert.Equal(t, float64(count/2), values[count/2].Value.GetDoubleValue())
	assert.Equal(t, float64(count-1), values[count-1].Value.GetDoubleValue())
}

func TestProfileValuesEncoder_EmptyUnit(t *testing.T) {
	encoder := NewProfileValuesEncoder(60, "")
	ts := time.Date(2024, 1, 1, 0, 0, 0, 0, time.UTC)
	encoder.AppendDouble(ts, 0, 0, 0, 1.0, nil)

	data := encoder.Bytes()
	decoder, err := NewProfileValuesDecoder(data)
	require.NoError(t, err)

	assert.Equal(t, "", decoder.GetUnit())
}

// Helper function to collect all values from decoder
func collectValues(t *testing.T, decoder *ProfileValuesDecoder) []ProfileValueItem {
	t.Helper()
	var values []ProfileValueItem
	for item := range decoder.Values() {
		if item.Err != nil {
			require.NoError(t, item.Err)
		}
		values = append(values, item)
	}
	return values
}

// Benchmark tests
func BenchmarkProfileValuesEncoder_AppendDouble(b *testing.B) {
	encoder := NewProfileValuesEncoder(60, "test")
	ts := time.Date(2024, 1, 1, 0, 0, 0, 0, time.UTC)

	for i := 0; b.Loop(); i++ {
		encoder.AppendDouble(ts.Add(time.Duration(i)*time.Second), 0, 0, 0, float64(i), nil)
	}
}

func BenchmarkProfileValuesEncoder_AppendInteger(b *testing.B) {
	encoder := NewProfileValuesEncoder(60, "test")
	ts := time.Date(2024, 1, 1, 0, 0, 0, 0, time.UTC)

	for i := 0; b.Loop(); i++ {
		encoder.AppendInteger(ts.Add(time.Duration(i)*time.Second), 0, 0, 0, int64(i), nil)
	}
}

func BenchmarkProfileValuesEncoder_AppendString(b *testing.B) {
	encoder := NewProfileValuesEncoder(60, "test")
	ts := time.Date(2024, 1, 1, 0, 0, 0, 0, time.UTC)

	for i := 0; b.Loop(); i++ {
		encoder.AppendString(ts.Add(time.Duration(i)*time.Second), 0, 0, 0, "test string", nil)
	}
}

func BenchmarkProfileValuesDecoder_Values(b *testing.B) {
	encoder := NewProfileValuesEncoder(60, "test")
	ts := time.Date(2024, 1, 1, 0, 0, 0, 0, time.UTC)

	for i := range 1000 {
		encoder.AppendDouble(ts.Add(time.Duration(i)*time.Second), 0, 0, 0, float64(i), nil)
	}
	data := encoder.Bytes()

	for b.Loop() {
		decoder, _ := NewProfileValuesDecoder(data)
		for range decoder.Values() {
		}
	}
}
