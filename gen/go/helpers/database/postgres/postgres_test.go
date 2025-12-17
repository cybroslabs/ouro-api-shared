package postgres

import (
	"testing"
)

func Test_1_dbPathToDbSelector(t *testing.T) {
	in := "col:$.abc.def"
	expected := []string{"col->'abc'->'def'", "", ""}
	if a, b, c, err := dbPathToDbSelector(in, false); a != expected[0] || b != expected[1] || c != expected[2] || err != nil {
		t.Errorf("dbPathToDbSelector(%q) = %q, %q, %q, %v; want %q, %q, %q", in, a, b, c, err, expected[0], expected[1], expected[2])
	}
}

func Test_2_dbPathToDbSelector(t *testing.T) {
	in := "col:$.abc.def"
	expected := []string{"col->'abc'->>'def'", "", ""}
	if a, b, c, err := dbPathToDbSelector(in, true); a != expected[0] || b != expected[1] || c != expected[2] || err != nil {
		t.Errorf("dbPathToDbSelector(%q) = %q, %q, %q, %v; want %q, %q, %q", in, a, b, c, err, expected[0], expected[1], expected[2])
	}
}

func Test_3_dbPathToDbSelector(t *testing.T) {
	in := "col:$.abc.def[*]@.p1"
	expected := []string{"col", "$.abc.def[*]", "@.p1"}
	if a, b, c, err := dbPathToDbSelector(in, false); a != expected[0] || b != expected[1] || c != expected[2] || err != nil {
		t.Errorf("dbPathToDbSelector(%q) = %q, %q, %q, %v; want %q, %q, %q", in, a, b, c, err, expected[0], expected[1], expected[2])
	}
}

func Test_4_dbPathToDbSelector(t *testing.T) {
	in := "\"table.col\":$.abc.def@.p1"
	// The input "table.col" contains quotes and a dot. QuoteIdentifiers will split on the dot
	// and quote each part independently. Since the quotes are inside the string parts after split,
	// each part will be re-quoted: "table -> ""table", "col"" -> ""col"""
	expected := []string{"\"\"\"table\".\"col\"\"\"", "$.abc.def", "@.p1"}
	if a, b, c, err := dbPathToDbSelector(in, false); a != expected[0] || b != expected[1] || c != expected[2] || err != nil {
		t.Errorf("dbPathToDbSelector(%q) = %q, %q, %q, %v; want %q, %q, %q", in, a, b, c, err, expected[0], expected[1], expected[2])
	}
}

// Test SQL injection attempt in identifier
func Test_5_dbPathToDbSelector_SQLInjection(t *testing.T) {
	in := "col; DROP TABLE users--:$.abc.def"
	_, _, _, err := dbPathToDbSelector(in, false)
	if err == nil {
		t.Errorf("dbPathToDbSelector(%q) should return error for SQL injection attempt, got nil", in)
	}
}

// Test SQL injection attempt in JSONB path
func Test_6_dbPathToDbSelector_JSONBPathInjection(t *testing.T) {
	in := "col:$') OR 1=1 --"
	_, _, _, err := dbPathToDbSelector(in, false)
	if err == nil {
		t.Errorf("dbPathToDbSelector(%q) should return error for JSONB path injection attempt, got nil", in)
	}
}

// Test invalid JSONB path format
func Test_7_dbPathToDbSelector_InvalidJSONBPath(t *testing.T) {
	in := "col:$.abc..def"
	_, _, _, err := dbPathToDbSelector(in, false)
	if err == nil {
		t.Errorf("dbPathToDbSelector(%q) should return error for invalid JSONB path format, got nil", in)
	}
}

// Test null byte in identifier
func Test_8_sanitizeIdentifier_NullByte(t *testing.T) {
	in := "col\x00umn"
	err := sanitizeIdentifier(in)
	if err == nil {
		t.Errorf("sanitizeIdentifier(%q) should return error for null byte, got nil", in)
	}
}

// Test validateJSONBPath with valid paths
func Test_9_validateJSONBPath_Valid(t *testing.T) {
	validPaths := []string{
		"$",
		"$.field",
		"$.field.nested",
		"$.field[0]",
		"$.field[*]",
		"$.field[*]@",
	}
	for _, path := range validPaths {
		if err := validateJSONBPath(path); err != nil {
			t.Errorf("validateJSONBPath(%q) should be valid, got error: %v", path, err)
		}
	}
}

// Test validateJSONBPath with invalid paths
func Test_10_validateJSONBPath_Invalid(t *testing.T) {
	invalidPaths := []string{
		"$'); DROP TABLE users--",
		"$' OR 1=1 --",
		"$/* comment */",
	}
	for _, path := range invalidPaths {
		if err := validateJSONBPath(path); err == nil {
			t.Errorf("validateJSONBPath(%q) should be invalid, got nil error", path)
		}
	}
}

// Test QuoteIdentifier with identifiers that don't need quoting
func Test_11_QuoteIdentifier_NoQuotingNeeded(t *testing.T) {
	tests := []struct {
		input    string
		expected string
	}{
		{"user_id", "user_id"},
		{"my_table", "my_table"},
		{"column123", "column123"},
		{"abc_def_123", "abc_def_123"},
		{"_private", "_private"},
	}
	for _, test := range tests {
		result := quoteSimpleIdentifier(test.input)
		if result != test.expected {
			t.Errorf("QuoteIdentifier(%q) = %q; want %q", test.input, result, test.expected)
		}
	}
}

// Test QuoteIdentifier with identifiers that need quoting
func Test_12_QuoteIdentifier_QuotingNeeded(t *testing.T) {
	tests := []struct {
		input    string
		expected string
	}{
		{"User", "\"User\""},               // uppercase
		{"user-id", "\"user-id\""},         // hyphen
		{"user.id", "\"user.id\""},         // dot
		{"user id", "\"user id\""},         // space
		{"123column", "\"123column\""},     // starts with digit
		{"select", "\"select\""},           // reserved keyword
		{"from", "\"from\""},               // reserved keyword
		{"WHERE", "\"WHERE\""},             // reserved keyword (uppercase)
		{"user\"name", "\"user\"\"name\""}, // contains quote
	}
	for _, test := range tests {
		result := quoteSimpleIdentifier(test.input)
		if result != test.expected {
			t.Errorf("QuoteIdentifier(%q) = %q; want %q", test.input, result, test.expected)
		}
	}
}

// Test QuoteIdentifier with already quoted identifiers
func Test_13_QuoteIdentifier_AlreadyQuoted(t *testing.T) {
	tests := []struct {
		input    string
		expected string
	}{
		{"\"User\"", "\"User\""},
		{"\"user-id\"", "\"user-id\""},
		{"\"my table\"", "\"my table\""},
		{"\"SELECT\"", "\"SELECT\""},
	}
	for _, test := range tests {
		result := quoteSimpleIdentifier(test.input)
		if result != test.expected {
			t.Errorf("QuoteIdentifier(%q) = %q; want %q (should not double-quote)", test.input, result, test.expected)
		}
	}
}

// Test QuoteIdentifiers with dotted identifiers
func Test_14_QuoteIdentifiers_DottedNames(t *testing.T) {
	tests := []struct {
		input    string
		expected string
	}{
		{"schema.table.column", "schema.\"table\".\"column\""},                 // "table" and "column" are reserved
		{"public.users.user_id", "public.users.user_id"},                       // none are keywords
		{"my_schema.my_table.my_column", "my_schema.my_table.my_column"},       // all simple
		{"MySchema.MyTable.MyColumn", "\"MySchema\".\"MyTable\".\"MyColumn\""}, // all uppercase
		{"schema.\"table\".column", "schema.\"table\".\"column\""},             // "column" is reserved, "table" already quoted (but split will re-evaluate)
	}
	for _, test := range tests {
		result := QuoteIdentifier(test.input)
		if result != test.expected {
			t.Errorf("QuoteIdentifiers(%q) = %q; want %q", test.input, result, test.expected)
		}
	}
}

// Test QuoteIdentifier with edge cases
func Test_15_QuoteIdentifier_EdgeCases(t *testing.T) {
	tests := []struct {
		input    string
		expected string
	}{
		{"", ""},             // empty string
		{"\"\"", "\"\""},     // already quoted empty
		{"_", "_"},           // just underscore
		{"__test", "__test"}, // double underscore start
	}
	for _, test := range tests {
		result := quoteSimpleIdentifier(test.input)
		if result != test.expected {
			t.Errorf("QuoteIdentifier(%q) = %q; want %q", test.input, result, test.expected)
		}
	}
}

// Test real-world usage examples showing smart quoting
func Test_16_QuoteIdentifier_RealWorldExamples(t *testing.T) {
	// Example 1: Simple table and column names - no quotes needed
	query1 := "SELECT " + quoteSimpleIdentifier("user_id") + ", " + quoteSimpleIdentifier("email") +
		" FROM " + QuoteIdentifier("public.users")
	expected1 := "SELECT user_id, email FROM public.users"
	if query1 != expected1 {
		t.Errorf("Query 1 = %q; want %q", query1, expected1)
	}

	// Example 2: Reserved keywords and special characters - quotes added automatically
	query2 := "SELECT " + quoteSimpleIdentifier("order") + ", " + quoteSimpleIdentifier("user-name") +
		" FROM " + QuoteIdentifier("schema.table")
	expected2 := "SELECT \"order\", \"user-name\" FROM schema.\"table\""
	if query2 != expected2 {
		t.Errorf("Query 2 = %q; want %q", query2, expected2)
	}

	// Example 3: Already quoted identifiers - no double quoting
	alreadyQuoted := "\"MyTable\""
	query3 := "SELECT * FROM " + quoteSimpleIdentifier(alreadyQuoted)
	expected3 := "SELECT * FROM \"MyTable\""
	if query3 != expected3 {
		t.Errorf("Query 3 = %q; want %q", query3, expected3)
	}

	// Example 4: Mixed case requires quoting
	query4 := "SELECT " + QuoteIdentifier("MySchema.MyTable.MyColumn")
	expected4 := "SELECT \"MySchema\".\"MyTable\".\"MyColumn\""
	if query4 != expected4 {
		t.Errorf("Query 4 = %q; want %q", query4, expected4)
	}
}
