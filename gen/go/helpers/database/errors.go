package database

import (
	"database/sql"
	"errors"
	"regexp"

	"github.com/cybroslabs/ouro-api-shared/gen/go/common"
	"google.golang.org/grpc/status"
)

var (
	_reSqlState = regexp.MustCompile(`\(SQLSTATE\s*([^ \)]+)\)`)
)

// Translates a database error into a common error.
func Translate(err error) error {
	return TranslateObject(err, "")
}

// Translates a database error into a common error, with an optional object name.
func TranslateObject(err error, objectName string) error {
	if err == nil {
		return nil
	}

	// Do not translate if it's already a gRPC status error.
	if _, ok := status.FromError(err); ok {
		return err
	}

	if errors.Is(err, sql.ErrNoRows) {
		if objectName == "" {
			return common.ErrOuroObjectNotFound
		}
		return errors.Join(common.ErrOuroObjectNotFound, errors.New(objectName))
	}

	// Support for PostgreSQL error codes from jackc/pgx library.
	// We don't want to import the library here, so we do simple regex matching on the error string.
	err_str := err.Error()
	if ss := _reSqlState.FindStringSubmatch(err_str); len(ss) == 2 {
		switch ss[1] {
		case "23503": // foreign key violation
			return common.ErrOuroForeignKeyViolation
		case "23505": // unique violation
			return common.ErrOuroUniqueConstraintViolation
		case "23514": // check constraint violation
			return common.ErrOuroConstraintViolation
		}
	}

	return errors.Join(common.ErrOuroGenericDatabaseError, err)
}
