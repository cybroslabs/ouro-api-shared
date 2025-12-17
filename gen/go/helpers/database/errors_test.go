package database

import (
	"database/sql"
	"errors"
	"testing"

	"github.com/cybroslabs/ouro-api-shared/gen/go/common"
	"github.com/stretchr/testify/assert"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func TestTranslate_NilError(t *testing.T) {
	result := Translate(nil)
	assert.NoError(t, result)
}

func TestTranslate_SqlErrNoRows(t *testing.T) {
	result := Translate(sql.ErrNoRows)
	assert.Error(t, result)
	assert.ErrorIs(t, result, common.ErrOuroObjectNotFound)
}

func TestTranslateObject_SqlErrNoRowsWithObjectName(t *testing.T) {
	result := TranslateObject(sql.ErrNoRows, "user")
	assert.Error(t, result)
	assert.ErrorIs(t, result, common.ErrOuroObjectNotFound)
	// Should contain the object name
	assert.Contains(t, result.Error(), "user")
}

func TestTranslateObject_SqlErrNoRowsWithoutObjectName(t *testing.T) {
	result := TranslateObject(sql.ErrNoRows, "")
	assert.Error(t, result)
	assert.ErrorIs(t, result, common.ErrOuroObjectNotFound)
}

func TestTranslate_GRPCStatusError(t *testing.T) {
	grpcErr := status.Error(codes.NotFound, "not found")
	result := Translate(grpcErr)

	// Should not be translated, return as-is
	assert.Equal(t, grpcErr, result)
}

func TestTranslateObject_GRPCStatusError(t *testing.T) {
	grpcErr := status.Error(codes.Internal, "internal error")
	result := TranslateObject(grpcErr, "test")

	// Should not be translated, return as-is
	assert.Equal(t, grpcErr, result)
}

func TestTranslate_PostgreSQLForeignKeyViolation(t *testing.T) {
	// Simulate PostgreSQL foreign key violation error
	pgErr := errors.New("ERROR: insert or update on table violates foreign key constraint (SQLSTATE 23503)")
	result := Translate(pgErr)

	assert.Error(t, result)
	assert.ErrorIs(t, result, common.ErrOuroForeignKeyViolation)
}

func TestTranslate_PostgreSQLUniqueViolation(t *testing.T) {
	// Simulate PostgreSQL unique constraint violation error
	pgErr := errors.New("ERROR: duplicate key value violates unique constraint (SQLSTATE 23505)")
	result := Translate(pgErr)

	assert.Error(t, result)
	assert.ErrorIs(t, result, common.ErrOuroUniqueConstraintViolation)
}

func TestTranslate_PostgreSQLCheckConstraintViolation(t *testing.T) {
	// Simulate PostgreSQL check constraint violation error
	pgErr := errors.New("ERROR: new row violates check constraint (SQLSTATE 23514)")
	result := Translate(pgErr)

	assert.Error(t, result)
	assert.ErrorIs(t, result, common.ErrOuroConstraintViolation)
}

func TestTranslate_PostgreSQLErrorVariations(t *testing.T) {
	tests := []struct {
		name          string
		errorMsg      string
		expectedError error
	}{
		{
			name:          "foreign key with extra spaces",
			errorMsg:      "ERROR: constraint violation (SQLSTATE  23503)",
			expectedError: common.ErrOuroForeignKeyViolation,
		},
		{
			name:          "unique violation with parenthesis",
			errorMsg:      "some error text (SQLSTATE 23505) more text",
			expectedError: common.ErrOuroUniqueConstraintViolation,
		},
		{
			name:          "check constraint at end",
			errorMsg:      "constraint failed (SQLSTATE 23514)",
			expectedError: common.ErrOuroConstraintViolation,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			pgErr := errors.New(tt.errorMsg)
			result := Translate(pgErr)
			assert.ErrorIs(t, result, tt.expectedError)
		})
	}
}

func TestTranslate_UnknownPostgreSQLError(t *testing.T) {
	// Unknown SQLSTATE code should fall through to generic error
	pgErr := errors.New("ERROR: some error (SQLSTATE 99999)")
	result := Translate(pgErr)

	assert.Error(t, result)
	assert.ErrorIs(t, result, common.ErrOuroGenericDatabaseError)
	// Original error should be joined
	assert.Contains(t, result.Error(), "99999")
}

func TestTranslate_GenericDatabaseError(t *testing.T) {
	// Generic error without SQLSTATE
	dbErr := errors.New("connection timeout")
	result := Translate(dbErr)

	assert.Error(t, result)
	assert.ErrorIs(t, result, common.ErrOuroGenericDatabaseError)
	// Original error should be joined
	assert.Contains(t, result.Error(), "connection timeout")
}

func TestTranslateObject_GenericErrorWithObjectName(t *testing.T) {
	dbErr := errors.New("some database error")
	result := TranslateObject(dbErr, "table_name")

	assert.Error(t, result)
	assert.ErrorIs(t, result, common.ErrOuroGenericDatabaseError)
	// Object name is not added for generic errors
	assert.Contains(t, result.Error(), "some database error")
}

func TestTranslate_NoSQLSTATE(t *testing.T) {
	// Error message without SQLSTATE pattern
	dbErr := errors.New("ERROR: database is locked")
	result := Translate(dbErr)

	assert.Error(t, result)
	assert.ErrorIs(t, result, common.ErrOuroGenericDatabaseError)
}

func TestTranslate_MalformedSQLSTATE(t *testing.T) {
	tests := []struct {
		name     string
		errorMsg string
	}{
		{
			name:     "missing closing parenthesis",
			errorMsg: "ERROR: test (SQLSTATE 23505",
		},
		{
			name:     "no code",
			errorMsg: "ERROR: test (SQLSTATE )",
		},
		{
			name:     "incomplete",
			errorMsg: "ERROR: test SQLSTATE 23505)",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			dbErr := errors.New(tt.errorMsg)
			result := Translate(dbErr)
			// Should fall through to generic error
			assert.ErrorIs(t, result, common.ErrOuroGenericDatabaseError)
		})
	}
}

func TestTranslate_WrappedSqlErrNoRows(t *testing.T) {
	wrapped := errors.Join(errors.New("query failed"), sql.ErrNoRows)
	result := Translate(wrapped)

	assert.Error(t, result)
	assert.ErrorIs(t, result, common.ErrOuroObjectNotFound)
}

func TestTranslateObject_PreservesOriginalError(t *testing.T) {
	originalErr := errors.New("original database error with details")
	result := TranslateObject(originalErr, "")

	assert.Error(t, result)
	// Original error message should be preserved
	assert.Contains(t, result.Error(), "original database error with details")
}

func TestTranslate_CaseInsensitivity(t *testing.T) {
	// The regex should work with different cases/formatting
	tests := []struct {
		name          string
		errorMsg      string
		expectedError error
	}{
		{
			name:          "uppercase SQLSTATE",
			errorMsg:      "error (SQLSTATE 23505)",
			expectedError: common.ErrOuroUniqueConstraintViolation,
		},
		{
			name:          "with newlines",
			errorMsg:      "error\n(SQLSTATE 23503)\ndetails",
			expectedError: common.ErrOuroForeignKeyViolation,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			dbErr := errors.New(tt.errorMsg)
			result := Translate(dbErr)
			assert.ErrorIs(t, result, tt.expectedError)
		})
	}
}

// Benchmark tests
func BenchmarkTranslate_SqlErrNoRows(b *testing.B) {
	for b.Loop() {
		_ = Translate(sql.ErrNoRows)
	}
}

func BenchmarkTranslate_PostgreSQLError(b *testing.B) {
	pgErr := errors.New("ERROR: duplicate key (SQLSTATE 23505)")

	for b.Loop() {
		_ = Translate(pgErr)
	}
}

func BenchmarkTranslate_GenericError(b *testing.B) {
	dbErr := errors.New("connection timeout")

	for b.Loop() {
		_ = Translate(dbErr)
	}
}

func BenchmarkTranslateObject_WithObjectName(b *testing.B) {
	for b.Loop() {
		_ = TranslateObject(sql.ErrNoRows, "user")
	}
}
