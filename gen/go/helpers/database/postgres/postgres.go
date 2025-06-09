package postgres

import (
	"errors"
	"fmt"
	"strings"

	"github.com/cybroslabs/ouro-api-shared/gen/go/common"
	"github.com/cybroslabs/ouro-api-shared/gen/go/helpers/database"
)

var (
	_NO_ARGS = make([]any, 0)
)

// PathToDbPathFunc is a function type that maps a object path to its corresponding database column name and or JSONB path within a JSONB column.
type PathToDbPathFunc func(path string) (dbPath string, ok bool)

// PrepareWOL prepares the SQL query with WHERE, ORDER BY, and LIMIT clauses based on the provided DbSelector and the path map.
// The function returns the WHERE clause, ORDER BY clause, LIMIT clause, and any arguments needed for the query.
// The pathToDbPath is used to map field paths to their corresponding database column names and or JSONB paths.
// It uses the specified modelColumn for the JSONB column and idColumn for the ID column.
// If the DbSelector is nil or has no filters, it returns an empty WHERE clause and a default LIMIT of 100.
// If the DbSelector contains IDs, it constructs a WHERE clause to filter by those IDs.
// If fixedWhere is provided, it appends those conditions to the WHERE clause.
// The function also handles errors related to invalid field IDs or unsupported data types.
// It is designed to be used in a PostgreSQL context where JSONB fields are queried.
// The function returns an error if there are issues with the input parameters or if the query cannot be constructed.
func PrepareWOL(in *database.DbSelector, pathToDbPath PathToDbPathFunc, modelColumn string, idColumn string, fixedWhere ...database.PersistentWhere) (qWhere string, qOrderBy string, qLimit string, qArgs []any, err error) {
	if in == nil {
		qArgs = appendFixedWhere(fixedWhere, &qWhere, qArgs)
		return
	}
	if pathToDbPath == nil {
		return "", "", "", nil, errors.New("pathToDbPath cannot be nil")
	}
	if err = in.Err(); err != nil {
		return
	}
	if cnt := len(in.Id); cnt > 0 {
		if cnt == 1 {
			qArgs = []any{in.Id[0]}
			qWhere = fmt.Sprintf("WHERE (%s = $1)", idColumn)
		} else {
			qArgs = make([]any, cnt)
			qWhere = fmt.Sprintf("WHERE (%s IN (", idColumn)
			for i, id := range in.Id {
				if i > 0 {
					qWhere += ", "
				}
				qWhere += fmt.Sprintf("$%d", i+1)
				qArgs[i] = id
			}
			qWhere += "))"
		}
		qArgs = appendFixedWhere(fixedWhere, &qWhere, qArgs)
		// This should return single rocord, so no need for ORDER BY or LIMIT
		return
	}

	if in.FilterBy != nil {
		qWhere, qArgs, err = getWhere(in, pathToDbPath, modelColumn)
		if err != nil {
			return
		}
	}

	qArgs = appendFixedWhere(fixedWhere, &qWhere, qArgs)

	qOrderBy, err = getOrderBy(in, pathToDbPath, modelColumn)
	if err != nil {
		return
	}

	qLimit, err = getLimitOffset(in)
	return
}

func appendFixedWhere(fixedWhere []database.PersistentWhere, qWhere *string, qArgsIn []any) (qArgs []any) {
	if len(fixedWhere) == 0 {
		if qArgsIn == nil {
			qArgs = _NO_ARGS
		} else {
			qArgs = qArgsIn
		}
		return
	}
	if len(*qWhere) == 0 {
		*qWhere = "WHERE "
	} else {
		*qWhere += " AND "
	}
	qArgs = qArgsIn
	for idx, item := range fixedWhere {
		if idx > 0 {
			*qWhere += " AND "
		}
		*qWhere += fmt.Sprintf("(%s = $%d)", item.Query, len(qArgs)+1)
		qArgs = append(qArgs, item.Arg)
	}
	return
}

func getWhere(in *database.DbSelector, pathToDbPath PathToDbPathFunc, modelColumn string) (string, []any, error) {
	var err error
	parts := make([]string, 0, len(in.GetFilterBy()))
	values := make([]any, 0, len(in.GetFilterBy()))
	for _, f := range in.GetFilterBy() {
		if !f.HasOperator() {
			continue
		}

		path := (&common.FieldDescriptor{}).ConvertJsPathToPath(f.GetPath())
		path, ok := pathToDbPath(path)
		if !ok {
			return "", nil, errors.New("unknwon path: " + f.GetPath())
		}

		col := dbPathToColumn(path, modelColumn, true)
		if len(col) == 0 {
			return "", nil, errors.New("invalid field id")
		}

		switch f.GetOperator() {
		case common.FilterOperator_EQUAL:
			makeOpVal := func(operand string) string { return " = " + operand }
			err = addSingleOperandOperator(&parts, &values, col, f, makeOpVal)
		case common.FilterOperator_NOT_EQUAL:
			makeOpVal := func(operand string) string { return " <> " + operand }
			err = addSingleOperandOperator(&parts, &values, col, f, makeOpVal)
		case common.FilterOperator_GREATER_THAN:
			makeOpVal := func(operand string) string { return " > " + operand }
			err = addSingleOperandOperator(&parts, &values, col, f, makeOpVal)
		case common.FilterOperator_GREATER_THAN_OR_EQUAL:
			makeOpVal := func(operand string) string { return " >= " + operand }
			err = addSingleOperandOperator(&parts, &values, col, f, makeOpVal)
		case common.FilterOperator_LESS_THAN:
			makeOpVal := func(operand string) string { return " < " + operand }
			err = addSingleOperandOperator(&parts, &values, col, f, makeOpVal)
		case common.FilterOperator_LESS_THAN_OR_EQUAL:
			makeOpVal := func(operand string) string { return " <= " + operand }
			err = addSingleOperandOperator(&parts, &values, col, f, makeOpVal)
		case common.FilterOperator_CONTAINS:
			makeOpVal := func(operand string) string { return " LIKE '%' || " + operand + " || '%' " }
			err = addSingleOperandOperator(&parts, &values, col, f, makeOpVal)
		case common.FilterOperator_NOT_CONTAINS:
			makeOpVal := func(operand string) string { return " NOT LIKE '%' || " + operand + " || '%' " }
			err = addSingleOperandOperator(&parts, &values, col, f, makeOpVal)
		case common.FilterOperator_STARTS_WITH:
			makeOpVal := func(operand string) string { return " LIKE " + operand + " || '%' " }
			err = addSingleOperandOperator(&parts, &values, col, f, makeOpVal)
		case common.FilterOperator_ENDS_WITH:
			makeOpVal := func(operand string) string { return " LIKE '%' || " + operand }
			err = addSingleOperandOperator(&parts, &values, col, f, makeOpVal)
		// Multi-operand operators
		case common.FilterOperator_IN:
			err = addMultiOperandOperator(&parts, &values, col, f, "IN")
		case common.FilterOperator_NOT_IN:
			err = addMultiOperandOperator(&parts, &values, col, f, "NOT IN")
			// 2-operand operators
		case common.FilterOperator_BETWEEN:
			if f.GetDataType() == common.FieldDataType_INTEGER {
				if t := f.GetInteger(); len(t) != 2 {
					return "", nil, errors.New("invalid number of operands")
				} else {
					if dpath := fieldToPath(path); len(dpath) > 0 {
						// Use jsonb_path_exists to optimize the query
						parts = append(parts, fmt.Sprintf("jsonb_path_exists(%s, '%s ? (@ >= %d && @ <= %d)')", modelColumn, dpath, t[0], t[1]))
					} else {
						parts = append(parts, fmt.Sprintf("%s >= $%d AND %s <= $%d", col, len(values)+1, col, len(values)+2))
						values = append(values, t[0], t[1])
					}
				}
			} else if f.GetDataType() == common.FieldDataType_DOUBLE {
				if t := f.GetNumber(); len(t) != 2 {
					return "", nil, errors.New("invalid number of operands")
				} else {
					if dpath := fieldToPath(path); len(dpath) > 0 {
						parts = append(parts, fmt.Sprintf("jsonb_path_exists(%s, '%s ? (@ >= %f && @ <= %f)')", modelColumn, dpath, t[0], t[1]))
					} else {
						parts = append(parts, fmt.Sprintf("%s >= $%d AND %s <= $%d", col, len(values)+1, col, len(values)+2))
						values = append(values, t[0], t[1])
					}
				}
			} else {
				return "", nil, errors.New("unsupported data type")
			}
			// No-operand operators
		case common.FilterOperator_IS_NULL:
			parts = append(parts, col+" IS NULL")
		case common.FilterOperator_IS_NOT_NULL:
			parts = append(parts, col+" IS NOT NULL")
		}

		if err != nil {
			return "", nil, err
		}
	}

	if len(parts) == 0 {
		return "", values, nil
	} else {
		return "WHERE (" + strings.Join(parts, ") AND (") + ")", values, nil
	}
}

func getOrderBy(in *database.DbSelector, pathToDbPath PathToDbPathFunc, modelColumn string) (string, error) {
	if in == nil {
		return "", nil
	}
	fields := in.GetSortBy()
	if len(fields) == 0 {
		return "", nil
	}

	tmp := strings.Builder{}
	tmp.WriteString("ORDER BY ")
	for i, s := range fields {
		path := (&common.FieldDescriptor{}).ConvertJsPathToPath(s.GetPath())
		path, ok := pathToDbPath(path)
		if !ok {
			return "", errors.New("unknwon path: " + s.GetPath())
		}

		col := dbPathToColumn(path, modelColumn, false)
		if len(col) == 0 {
			return "", errors.New("invalid field id")
		}
		if i > 0 {
			tmp.WriteString(", ")
		}
		tmp.WriteString(col)
		if s.GetDesc() {
			tmp.WriteString(" DESC")
		}
	}
	return tmp.String(), nil
}

func getLimitOffset(in *database.DbSelector) (string, error) {
	if in == nil {
		return " LIMIT 100", nil
	}
	limit := in.GetPageSize()
	if limit == 0 {
		limit = 100
	} else if limit > 10000 {
		return "", errors.New("limit too high")
	}
	return fmt.Sprintf(" LIMIT %d OFFSET %d", limit, in.GetOffset()), nil
}

func addMultiOperandOperator(parts *[]string, values *[]any, col string, in *common.ListSelectorFilterBy, operator string) error {
	base_id := len(*values) + 1
	switch in.GetDataType() {
	case common.FieldDataType_TEXT:
		for _, t := range in.GetText() {
			*values = append(*values, t)
		}
	case common.FieldDataType_INTEGER:
		for _, t := range in.GetInteger() {
			*values = append(*values, t)
		}
		col = "(" + col + ")::int"
	case common.FieldDataType_DOUBLE:
		for _, t := range in.GetNumber() {
			*values = append(*values, t)
		}
		col = "(" + col + ")::numeric"
	default:
		return errors.New("unsupported data type")
	}
	if len(*values)+1 == base_id {
		return errors.New("invalid number of operands")
	}

	tmp := strings.Builder{}
	tmp.WriteString(col)
	tmp.WriteString(" ")
	tmp.WriteString(operator)
	tmp.WriteString(" (")
	for i := base_id; i <= len(*values); i++ {
		if i > base_id {
			tmp.WriteString(", ")
		}
		tmp.WriteString(fmt.Sprintf("$%d", i))
	}
	tmp.WriteString(")")
	*parts = append(*parts, tmp.String())
	return nil
}

func addSingleOperandOperator(parts *[]string, values *[]any, col string, in *common.ListSelectorFilterBy, composeOpVal func(operand string) string) error {
	switch in.GetDataType() {
	case common.FieldDataType_TEXT:
		if t := in.GetText(); len(t) != 1 {
			return errors.New("invalid number of operands")
		} else {
			*values = append(*values, t[0])
			*parts = append(*parts, col+composeOpVal(fmt.Sprintf("$%d", len(*values))))
		}
	case common.FieldDataType_INTEGER:
		if t := in.GetInteger(); len(t) != 1 {
			return errors.New("invalid number of operands")
		} else {
			*values = append(*values, t[0])
			*parts = append(*parts, "("+col+")::int"+composeOpVal(fmt.Sprintf("$%d", len(*values))))
		}
	case common.FieldDataType_BOOLEAN:
		if t := in.GetBoolean(); len(t) != 1 {
			return errors.New("invalid number of operands")
		} else {
			*values = append(*values, t[0])
			*parts = append(*parts, "("+col+")::bool"+composeOpVal(fmt.Sprintf("$%d", len(*values))))
		}
	case common.FieldDataType_DOUBLE:
		if t := in.GetNumber(); len(t) != 1 {
			return errors.New("invalid number of operands")
		} else {
			*values = append(*values, t[0])
			*parts = append(*parts, "("+col+")::numeric"+composeOpVal(fmt.Sprintf("$%d", len(*values))))
		}
	default:
		return errors.New("unsupported data type")
	}

	return nil
}

func fieldToPath(field string) string {
	parts := strings.Split(field, ".")
	if len(parts) >= 2 && (parts[0] == "$") {
		return "$." + strings.Join(parts[1:], ".")
	} else {
		// Neither a valid object field nor a valid column name
		return ""
	}
}

func dbPathToColumn(dbPath string, modelColumn string, useDoubleArrow bool) string {
	parts := strings.Split(dbPath, ".")
	if len(parts) >= 2 && (parts[0] == "$") {
		if !useDoubleArrow {
			return modelColumn + "->'" + strings.Join(parts[1:], "'->'") + "'"
		}
		if len(parts) > 2 {
			return modelColumn + "->'" + strings.Join(parts[1:len(parts)-1], "'->'") + "'->>'" + parts[len(parts)-1] + "'"
		} else {
			return modelColumn + "->>'" + parts[len(parts)-1] + "'"
		}
	} else if len(parts) == 1 {
		return dbPath
	} else {
		// Neither a valid object field nor a valid column name
		return ""
	}
}
