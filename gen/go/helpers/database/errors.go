package database

import (
	"database/sql"
	"errors"
	"fmt"
)

var (
	ErrObjectNotFound            = errors.New("not found")
	ErrUniqueConstraintViolation = errors.New("unique constraint violation")
	ErrForeignKeyViolation       = errors.New("foreign key violation")
	ErrConstraintViolation       = errors.New("referenced object does not exist or is not valid")
	ErrGenericDatabaseError      = errors.New("generic database error")
)

func Translate(err error, objectName string) error {
	if err == nil {
		return nil
	}

	if errors.Is(err, sql.ErrNoRows) {
		return errors.Join(ErrObjectNotFound, fmt.Errorf("%s not found", objectName))
	}

	// var pqErr *pq.Error
	// if errors.As(err, &pqErr) {
	// 	switch pqErr.Code {
	// 	case "23503": // foreign key violation
	// 		return ErrForeignKeyViolation
	// 	case "23505": // unique violation
	// 		return ErrUniqueConstraintViolation
	// 	case "23514": // check constraint violation
	// 		return ErrConstraintViolation
	// 	}
	// }

	return errors.Join(ErrGenericDatabaseError, fmt.Errorf("unexpected error: %w", err))
}
