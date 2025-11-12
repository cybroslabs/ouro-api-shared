package database

import (
	"database/sql"
	"errors"
	"fmt"
	"regexp"

	"github.com/cybroslabs/ouro-api-shared/gen/go/common"
)

var (
	_reSqlState = regexp.MustCompile(`\(SQLSTATE\s*([^ \)]+)\)`)
)

func Translate(err error, objectName string) error {
	if err == nil {
		return nil
	}

	if errors.Is(err, sql.ErrNoRows) {
		return errors.Join(common.ErrOuroObjectNotFound, fmt.Errorf("%s not found", objectName))
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
