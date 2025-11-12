package common

import "errors"

var (
	ErrOuroGeneric                   = errors.New("ouro: generic error")
	ErrOuroValidation                = errors.New("ouro: validation error")
	ErrOuroManagedFieldsMustBeEmpty  = errors.New("ouro: managed fields must be empty")
	ErrOuroObjectNotFound            = errors.New("ouro: not found")
	ErrOuroUniqueConstraintViolation = errors.New("ouro: unique constraint violation")
	ErrOuroForeignKeyViolation       = errors.New("ouro: foreign key violation")
	ErrOuroConstraintViolation       = errors.New("ouro: referenced object does not exist or is not valid")
	ErrOuroGenericDatabaseError      = errors.New("ouro: generic database error")
)
