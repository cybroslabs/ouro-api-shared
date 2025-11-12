package common

import (
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

var (
	ErrOuroGeneric                   = status.New(codes.Internal, "ouro: generic error").Err()
	ErrOuroValidation                = status.New(codes.InvalidArgument, "ouro: validation error").Err()
	ErrOuroManagedFieldsMustBeEmpty  = status.New(codes.InvalidArgument, "ouro: managed fields must be empty").Err()
	ErrOuroObjectNotFound            = status.New(codes.NotFound, "ouro: not found").Err()
	ErrOuroUniqueConstraintViolation = status.New(codes.AlreadyExists, "ouro: unique constraint violation").Err()
	ErrOuroForeignKeyViolation       = status.New(codes.InvalidArgument, "ouro: foreign key violation").Err()
	ErrOuroConstraintViolation       = status.New(codes.InvalidArgument, "ouro: referenced object does not exist or is not valid").Err()
	ErrOuroGenericDatabaseError      = status.New(codes.Internal, "ouro: generic database error").Err()
)
