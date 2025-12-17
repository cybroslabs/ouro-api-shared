package common

import (
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestNewFormattedMessage_Simple(t *testing.T) {
	fm := NewFormattedMessage("Hello, World!")
	assert.NotNil(t, fm)
	assert.Equal(t, "Hello, World!", fm.GetMessage())
	assert.Empty(t, fm.GetArgs())
}

func TestNewFormattedMessage_WithStringArg(t *testing.T) {
	fm := NewFormattedMessage("Hello, %s!", "Alice")
	assert.NotNil(t, fm)
	assert.Equal(t, "Hello, %s!", fm.GetMessage())
	assert.Len(t, fm.GetArgs(), 1)
	assert.Equal(t, "Alice", fm.GetArgs()[0])
}

func TestNewFormattedMessage_WithMultipleArgs(t *testing.T) {
	fm := NewFormattedMessage("User %s has %d points", "Bob", 42)
	assert.NotNil(t, fm)
	assert.Equal(t, "User %s has %s points", fm.GetMessage())
	assert.Len(t, fm.GetArgs(), 2)
	assert.Equal(t, "Bob", fm.GetArgs()[0])
	assert.Equal(t, "42", fm.GetArgs()[1])
}

func TestNewFormattedMessage_IntegerFormatting(t *testing.T) {
	tests := []struct {
		name     string
		format   string
		arg      int
		expected string
	}{
		{"decimal", "Value: %d", 42, "42"},
		{"octal", "Value: %o", 8, "10"},
		{"hex lowercase", "Value: %x", 255, "ff"},
		{"hex uppercase", "Value: %X", 255, "FF"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			fm := NewFormattedMessage(tt.format, tt.arg)
			assert.Equal(t, "Value: %s", fm.GetMessage())
			assert.Len(t, fm.GetArgs(), 1)
			assert.Equal(t, tt.expected, fm.GetArgs()[0])
		})
	}
}

func TestNewFormattedMessage_FloatFormatting(t *testing.T) {
	tests := []struct {
		name     string
		format   string
		arg      float64
		expected string
	}{
		{"default", "Pi: %f", 3.14159, "3.141590"},
		{"scientific", "Pi: %e", 3.14159, "3.141590e+00"},
		{"compact", "Pi: %g", 3.14159, "3.14159"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			fm := NewFormattedMessage(tt.format, tt.arg)
			assert.Equal(t, "Pi: %s", fm.GetMessage())
			assert.Len(t, fm.GetArgs(), 1)
			assert.Equal(t, tt.expected, fm.GetArgs()[0])
		})
	}
}

func TestNewFormattedMessage_BooleanFormatting(t *testing.T) {
	fm := NewFormattedMessage("Status: %t", true)
	assert.Equal(t, "Status: %s", fm.GetMessage())
	assert.Len(t, fm.GetArgs(), 1)
	assert.Equal(t, "true", fm.GetArgs()[0])
}

func TestNewFormattedMessage_PercentEscape(t *testing.T) {
	fm := NewFormattedMessage("100%% complete")
	assert.Equal(t, "100% complete", fm.GetMessage())
	assert.Empty(t, fm.GetArgs())
}

func TestNewFormattedMessage_PercentWithArg(t *testing.T) {
	fm := NewFormattedMessage("Progress: %d%% done", 75)
	assert.Equal(t, "Progress: %s% done", fm.GetMessage())
	assert.Len(t, fm.GetArgs(), 1)
	assert.Equal(t, "75", fm.GetArgs()[0])
}

func TestNewFormattedMessage_WithError(t *testing.T) {
	err := errors.New("something went wrong")
	fm := NewFormattedMessage("Error: %w", err)
	assert.Equal(t, "Error: %s", fm.GetMessage())
	assert.Len(t, fm.GetArgs(), 1)
	assert.Equal(t, "something went wrong", fm.GetArgs()[0])
}

func TestNewFormattedMessage_MissingArgs(t *testing.T) {
	// More format specifiers than args
	fm := NewFormattedMessage("User %s has %d points", "Alice")
	assert.Equal(t, "User %s has %s points", fm.GetMessage())
	assert.Len(t, fm.GetArgs(), 2)
	assert.Equal(t, "Alice", fm.GetArgs()[0])
	assert.Equal(t, "invalid", fm.GetArgs()[1])
}

func TestNewFormattedMessage_ExtraArgs(t *testing.T) {
	// More args than format specifiers
	fm := NewFormattedMessage("Hello, %s!", "Alice", "Bob")
	assert.Equal(t, "Hello, %s!", fm.GetMessage())
	assert.Len(t, fm.GetArgs(), 1)
	assert.Equal(t, "Alice", fm.GetArgs()[0])
	// "Bob" is ignored
}

func TestNewFormattedMessage_ComplexFormat(t *testing.T) {
	fm := NewFormattedMessage("User %s (ID: %d) scored %.2f points", "Charlie", 123, 98.5)
	assert.Equal(t, "User %s (ID: %s) scored %s points", fm.GetMessage())
	assert.Len(t, fm.GetArgs(), 3)
	assert.Equal(t, "Charlie", fm.GetArgs()[0])
	assert.Equal(t, "123", fm.GetArgs()[1])
	assert.Equal(t, "98.50", fm.GetArgs()[2])
}

func TestNewFormattedMessage_EmptyFormat(t *testing.T) {
	fm := NewFormattedMessage("")
	assert.NotNil(t, fm)
	assert.Equal(t, "", fm.GetMessage())
	assert.Empty(t, fm.GetArgs())
}

func TestNewFormattedMessage_NoArgs(t *testing.T) {
	fm := NewFormattedMessage("Static message")
	assert.NotNil(t, fm)
	assert.Equal(t, "Static message", fm.GetMessage())
	assert.Empty(t, fm.GetArgs())
}

func TestFormattedMessage_AddParams_Success(t *testing.T) {
	fm := NewFormattedMessage("Hello, World!")

	err := fm.AddParams("key1", "value1", "key2", 42)
	require.NoError(t, err)

	params := fm.GetParams().AsMap()
	assert.Len(t, params, 2)
	assert.Equal(t, "value1", params["key1"])
	assert.Equal(t, float64(42), params["key2"]) // structpb converts numbers to float64
}

func TestFormattedMessage_AddParams_MultipleCallsAccumulate(t *testing.T) {
	fm := NewFormattedMessage("Message")

	err := fm.AddParams("key1", "value1")
	require.NoError(t, err)

	err = fm.AddParams("key2", "value2")
	require.NoError(t, err)

	params := fm.GetParams().AsMap()
	assert.Len(t, params, 2)
	assert.Equal(t, "value1", params["key1"])
	assert.Equal(t, "value2", params["key2"])
}

func TestFormattedMessage_AddParams_OverwriteExisting(t *testing.T) {
	fm := NewFormattedMessage("Message")

	err := fm.AddParams("key1", "original")
	require.NoError(t, err)

	err = fm.AddParams("key1", "updated")
	require.NoError(t, err)

	params := fm.GetParams().AsMap()
	assert.Len(t, params, 1)
	assert.Equal(t, "updated", params["key1"])
}

func TestFormattedMessage_AddParams_OddNumberOfArgs(t *testing.T) {
	fm := NewFormattedMessage("Message")

	err := fm.AddParams("key1", "value1", "key2")
	assert.Error(t, err)
	assert.Contains(t, err.Error(), "even number")
	assert.Contains(t, err.Error(), "3")
}

func TestFormattedMessage_AddParams_NonStringKey(t *testing.T) {
	fm := NewFormattedMessage("Message")

	err := fm.AddParams(123, "value1")
	assert.Error(t, err)
	assert.Contains(t, err.Error(), "string keys")
	assert.Contains(t, err.Error(), "int")
}

func TestFormattedMessage_AddParams_VariousValueTypes(t *testing.T) {
	fm := NewFormattedMessage("Message")

	err := fm.AddParams(
		"string", "text",
		"int", 42,
		"float", 3.14,
		"bool", true,
		"nil", nil,
	)
	require.NoError(t, err)

	params := fm.GetParams().AsMap()
	assert.Len(t, params, 5)
	assert.Equal(t, "text", params["string"])
	assert.Equal(t, float64(42), params["int"])
	assert.Equal(t, 3.14, params["float"])
	assert.Equal(t, true, params["bool"])
	assert.Nil(t, params["nil"])
}

func TestFormattedMessage_FormattedString_Simple(t *testing.T) {
	fm := NewFormattedMessage("Hello, World!")
	result := fm.FormattedString()
	assert.Equal(t, "Hello, World!", result)
}

func TestFormattedMessage_FormattedString_WithArgs(t *testing.T) {
	fm := NewFormattedMessage("Hello, %s!", "Alice")
	result := fm.FormattedString()
	assert.Equal(t, "Hello, Alice!", result)
}

func TestFormattedMessage_FormattedString_MultipleArgs(t *testing.T) {
	fm := NewFormattedMessage("User %s has %d points", "Bob", 42)
	result := fm.FormattedString()
	assert.Equal(t, "User Bob has 42 points", result)
}

func TestFormattedMessage_FormattedString_EmptyMessage(t *testing.T) {
	fm := &FormattedMessage{}
	result := fm.FormattedString()
	assert.Equal(t, "", result)
}

func TestFormattedMessage_FormattedString_NoArgs(t *testing.T) {
	fm := FormattedMessage_builder{
		Message: stringPtr("Static message"),
		Args:    []string{},
	}.Build()

	result := fm.FormattedString()
	assert.Equal(t, "Static message", result)
}

func TestFormattedMessage_FormattedString_ComplexFormatting(t *testing.T) {
	fm := NewFormattedMessage("User %s (ID: %d) scored %.2f points", "Charlie", 123, 98.5)
	result := fm.FormattedString()
	assert.Equal(t, "User Charlie (ID: 123) scored 98.50 points", result)
}

func TestFormattedMessage_RoundTrip(t *testing.T) {
	// Create message, format it, verify output
	original := "Error: %s occurred at line %d"
	fm := NewFormattedMessage(original, "timeout", 42)

	formatted := fm.FormattedString()
	assert.Equal(t, "Error: timeout occurred at line 42", formatted)

	// Verify internal structure
	assert.Equal(t, "Error: %s occurred at line %s", fm.GetMessage())
	assert.Len(t, fm.GetArgs(), 2)
}

func TestFormattedMessage_SpecialCharacters(t *testing.T) {
	fm := NewFormattedMessage("Path: %s", "/usr/local/bin")
	result := fm.FormattedString()
	assert.Equal(t, "Path: /usr/local/bin", result)
}

func TestFormattedMessage_Unicode(t *testing.T) {
	fm := NewFormattedMessage("Message: %s", "Hello 世界 🌍")
	result := fm.FormattedString()
	assert.Equal(t, "Message: Hello 世界 🌍", result)
}

func TestNewFormattedMessage_AllFormatVerbs(t *testing.T) {
	// Test coverage for all format verb cases
	tests := []struct {
		format string
		arg    any
	}{
		{"%v", "value"},
		{"%T", "type"},
		{"%t", true},
		{"%b", 5},
		{"%c", 'A'},
		{"%d", 42},
		{"%o", 8},
		{"%O", 8},
		{"%q", "quoted"},
		{"%x", 255},
		{"%X", 255},
		{"%U", 'A'},
		{"%e", 1.23},
		{"%E", 1.23},
		{"%f", 3.14},
		{"%F", 3.14},
		{"%g", 1.5},
		{"%G", 1.5},
		{"%s", "string"},
		{"%p", "pointer"},
	}

	for _, tt := range tests {
		t.Run(tt.format, func(t *testing.T) {
			fm := NewFormattedMessage(tt.format, tt.arg)
			assert.Equal(t, "%s", fm.GetMessage())
			assert.Len(t, fm.GetArgs(), 1)
			assert.NotEmpty(t, fm.GetArgs()[0])
		})
	}
}

func TestFormattedMessage_AddParams_EmptyCall(t *testing.T) {
	fm := NewFormattedMessage("Message")
	err := fm.AddParams()
	require.NoError(t, err)
	// Should succeed with no changes
}

// Benchmark tests
func BenchmarkNewFormattedMessage_Simple(b *testing.B) {
	for b.Loop() {
		_ = NewFormattedMessage("Hello, World!")
	}
}

func BenchmarkNewFormattedMessage_WithArgs(b *testing.B) {
	for b.Loop() {
		_ = NewFormattedMessage("User %s has %d points", "Alice", 42)
	}
}

func BenchmarkFormattedMessage_FormattedString(b *testing.B) {
	fm := NewFormattedMessage("User %s has %d points", "Alice", 42)

	for b.Loop() {
		_ = fm.FormattedString()
	}
}

func BenchmarkFormattedMessage_AddParams(b *testing.B) {
	fm := NewFormattedMessage("Message")

	for b.Loop() {
		_ = fm.AddParams("key", "value")
	}
}
