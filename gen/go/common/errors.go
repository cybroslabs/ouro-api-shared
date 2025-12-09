package common

import (
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

var (
	// ErrOuroGeneric indicates a generic internal error that doesn't fit other categories.
	ErrOuroGeneric = status.New(codes.Internal, "ouro: generic error").Err()

	// ErrOuroValidation indicates that the input failed validation checks.
	ErrOuroValidation = status.New(codes.InvalidArgument, "ouro: validation error").Err()

	// ErrOuroManagedFieldsMustBeEmpty indicates that system-managed fields were provided in user input when they should be empty.
	ErrOuroManagedFieldsMustBeEmpty = status.New(codes.InvalidArgument, "ouro: managed fields must be empty").Err()

	// ErrOuroObjectNotFound indicates that the requested object does not exist in the system.
	ErrOuroObjectNotFound = status.New(codes.NotFound, "ouro: not found").Err()

	// ErrOuroUniqueConstraintViolation indicates that the operation violates a unique constraint (e.g., duplicate name or ID).
	ErrOuroUniqueConstraintViolation = status.New(codes.AlreadyExists, "ouro: unique constraint violation").Err()

	// ErrOuroForeignKeyViolation indicates that the operation violates a foreign key constraint.
	ErrOuroForeignKeyViolation = status.New(codes.InvalidArgument, "ouro: foreign key violation").Err()

	// ErrOuroConstraintViolation indicates that a referenced object does not exist or is not valid for the operation.
	ErrOuroConstraintViolation = status.New(codes.InvalidArgument, "ouro: referenced object does not exist or is not valid").Err()

	// ErrOuroGenericDatabaseError indicates a database error that doesn't fit other specific categories.
	ErrOuroGenericDatabaseError = status.New(codes.Internal, "ouro: generic database error").Err()
)
