package postgres

import (
	"errors"
	"fmt"
	"strconv"
	"strings"

	"github.com/cybroslabs/ouro-api-shared/gen/go/common"
	"github.com/cybroslabs/ouro-api-shared/gen/go/helpers/database"
)

var (
	_NO_ARGS  = make([]any, 0)
	_DUMMY_FD = &common.FieldDescriptor{}
)

// PathToDbPathFunc is a function type that maps a object path to its corresponding database column name and or JSONB path within a JSONB column.
// dbPath examples:
//
//	schema.table.col:$.level1.level2.levelN
//	schema.table.col:$.level1.level2.levelN[*]@.propertyName
//
// If the dbPath contains ':' then json select shall be used, otherwise it's a direct column name.
// If the dbPath contains '@' then JSONB_PATH_EXISTS function shall be used, otherwise the JSONB column is used directly.
type PathToDbPathFunc func(path string) (dbPath string, ok bool)

// PrepareWOL prepares the SQL query with WHERE, ORDER BY, and LIMIT clauses based on the provided DbSelector and the path map.
// The function returns the WHERE clause, ORDER BY clause, LIMIT clause, and any arguments needed for the query.
// The pathToDbPath is used to map field paths to their corresponding database column names and or JSONB paths.
// If the DbSelector is nil or has no filters, it returns an empty WHERE clause and a default LIMIT of 100.
// If the DbSelector contains IDs, it constructs a WHERE clause to filter by those IDs.
// If fixedWhere is provided, it appends those conditions to the WHERE clause.
// The function also handles errors related to invalid field IDs or unsupported data types.
// It is designed to be used in a PostgreSQL context where JSONB fields are queried.
// The function returns an error if there are issues with the input parameters or if the query cannot be constructed.
func PrepareWOL(in *database.DbSelector, pathToDbPath PathToDbPathFunc, idColumn string, fixedWhere ...database.PersistentWhere) (qWhere string, qOrderBy string, qLimit string, qArgs []any, err error) {
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
		qWhere, qArgs, err = getWhere(in, pathToDbPath)
		if err != nil {
			return
		}
	}

	qArgs = appendFixedWhere(fixedWhere, &qWhere, qArgs)

	qOrderBy, err = getOrderBy(in, pathToDbPath)
	if err != nil {
		return
	}

	qLimit, err = getLimitOffset(in)
	return
}

func escapeForRegex(s string) string {
	if s == "" {
		return s
	}
	s = strings.ReplaceAll(s, `\`, `\\`) // escape backslash first
	s = strings.ReplaceAll(s, `"`, `\"`) // then escape double quote
	return s
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

func getWhere(in *database.DbSelector, pathToDbPath PathToDbPathFunc) (string, []any, error) {
	var err error
	parts := make([]string, 0, len(in.GetFilterBy()))
	values := make([]any, 0, len(in.GetFilterBy()))
	for _, f := range in.GetFilterBy() {
		if !f.HasOperator() {
			continue
		}

		raw_path := f.GetPath()
		path, ok := pathToDbPath(_DUMMY_FD.ConvertJsPathToPath(raw_path))
		if !ok {
			return "", nil, errors.New("unknown path: " + raw_path)
		}

		var col, json_path, json_property string
		col, json_path, json_property, err = dbPathToDbSelector(path, true)
		if err != nil {
			return "", nil, err
		}
		use_jsonb_func := len(json_property) > 0

		switch f.GetOperator() {
		case common.FilterOperator_EQUAL:
			if !use_jsonb_func {
				makeOpVal := func(operand string) string { return " = " + operand }
				err = addSingleOperandOperator(&parts, &values, col, f, makeOpVal)
			} else {
				makeOpVal := func(value string) (string, bool) {
					if len(value) > 0 {
						return ` == "` + escapeForRegex(value) + `"`, false
					}
					return " == ", false
				}
				err = addSingleOperandOperatorJson(&parts, col, json_path, json_property, f, makeOpVal)
			}
		case common.FilterOperator_NOT_EQUAL:
			if !use_jsonb_func {
				makeOpVal := func(operand string) string { return " <> " + operand }
				err = addSingleOperandOperator(&parts, &values, col, f, makeOpVal)
			} else {
				makeOpVal := func(value string) (string, bool) {
					if len(value) > 0 {
						return ` <> "` + escapeForRegex(value) + `"`, false
					}
					return " <> ", false
				}
				err = addSingleOperandOperatorJson(&parts, col, json_path, json_property, f, makeOpVal)
			}
		case common.FilterOperator_GREATER_THAN:
			if !use_jsonb_func {
				makeOpVal := func(operand string) string { return " > " + operand }
				err = addSingleOperandOperator(&parts, &values, col, f, makeOpVal)
			} else {
				makeOpVal := func(value string) (string, bool) {
					if len(value) > 0 {
						return ` > "` + escapeForRegex(value) + `"`, false
					}
					return " > ", false
				}
				err = addSingleOperandOperatorJson(&parts, col, json_path, json_property, f, makeOpVal)
			}
		case common.FilterOperator_GREATER_THAN_OR_EQUAL:
			if !use_jsonb_func {
				makeOpVal := func(operand string) string { return " >= " + operand }
				err = addSingleOperandOperator(&parts, &values, col, f, makeOpVal)
			} else {
				makeOpVal := func(value string) (string, bool) {
					if len(value) > 0 {
						return ` >= "` + escapeForRegex(value) + `"`, false
					}
					return " >= ", false
				}
				err = addSingleOperandOperatorJson(&parts, col, json_path, json_property, f, makeOpVal)
			}
		case common.FilterOperator_LESS_THAN:
			if !use_jsonb_func {
				makeOpVal := func(operand string) string { return " < " + operand }
				err = addSingleOperandOperator(&parts, &values, col, f, makeOpVal)
			} else {
				makeOpVal := func(value string) (string, bool) {
					if len(value) > 0 {
						return ` < "` + escapeForRegex(value) + `"`, false
					}
					return " < ", false
				}
				err = addSingleOperandOperatorJson(&parts, col, json_path, json_property, f, makeOpVal)
			}
		case common.FilterOperator_LESS_THAN_OR_EQUAL:
			if !use_jsonb_func {
				makeOpVal := func(operand string) string { return " <= " + operand }
				err = addSingleOperandOperator(&parts, &values, col, f, makeOpVal)
			} else {
				makeOpVal := func(value string) (string, bool) {
					if len(value) > 0 {
						return ` <= "` + escapeForRegex(value) + `"`, false
					}
					return " <= ", false
				}
				err = addSingleOperandOperatorJson(&parts, col, json_path, json_property, f, makeOpVal)
			}
		case common.FilterOperator_CONTAINS:
			if !use_jsonb_func {
				makeOpVal := func(operand string) string { return " LIKE '%' || " + operand + " || '%' " }
				err = addSingleOperandOperator(&parts, &values, col, f, makeOpVal)
			} else {
				makeOpVal := func(value string) (string, bool) {
					if len(value) > 0 {
						return ` like_regex "` + escapeForRegex(value) + `"`, false
					}
					return "", false
				}
				err = addSingleOperandOperatorJson(&parts, col, json_path, json_property, f, makeOpVal)
			}
		case common.FilterOperator_NOT_CONTAINS:
			if !use_jsonb_func {
				makeOpVal := func(operand string) string { return " NOT LIKE '%' || " + operand + " || '%' " }
				err = addSingleOperandOperator(&parts, &values, col, f, makeOpVal)
			} else {
				makeOpVal := func(value string) (string, bool) {
					if len(value) > 0 {
						return ` like_regex "` + escapeForRegex(value) + `"`, true
					}
					return "", true
				}
				err = addSingleOperandOperatorJson(&parts, col, json_path, json_property, f, makeOpVal)
			}
		case common.FilterOperator_STARTS_WITH:
			if !use_jsonb_func {
				makeOpVal := func(operand string) string { return " LIKE " + operand + " || '%' " }
				err = addSingleOperandOperator(&parts, &values, col, f, makeOpVal)
			} else {
				makeOpVal := func(value string) (string, bool) {
					if len(value) > 0 {
						return ` like_regex "^` + escapeForRegex(value) + `"`, false
					}
					return "", false
				}
				err = addSingleOperandOperatorJson(&parts, col, json_path, json_property, f, makeOpVal)
			}
		case common.FilterOperator_ENDS_WITH:
			if !use_jsonb_func {
				makeOpVal := func(operand string) string { return " LIKE '%' || " + operand }
				err = addSingleOperandOperator(&parts, &values, col, f, makeOpVal)
			} else {
				makeOpVal := func(value string) (string, bool) {
					if len(value) > 0 {
						return ` like_regex "` + escapeForRegex(value) + `$"`, false
					}
					return "", false
				}
				err = addSingleOperandOperatorJson(&parts, col, json_path, json_property, f, makeOpVal)
			}
		// Multi-operand operators
		case common.FilterOperator_IN:
			if !use_jsonb_func {
				err = addMultiOperandOperator(&parts, &values, col, f, "IN")
			} else {
				err = addMultiOperandOperatorJson(&parts, col, json_path, json_property, f, "==")
			}
		case common.FilterOperator_NOT_IN:
			err = addMultiOperandOperator(&parts, &values, col, f, "NOT IN")
			// 2-operand operators
		case common.FilterOperator_BETWEEN:
			if f.GetDataType() == common.FieldDataType_INTEGER {
				if t := f.GetInteger(); len(t) != 2 {
					return "", nil, errors.New("invalid number of operands")
				} else {
					if len(json_path) > 0 {
						// Use jsonb_path_exists to optimize the query
						parts = append(parts, fmt.Sprintf("jsonb_path_exists(%s, '%s ? (%s >= %d && %s <= %d)')", col, json_path, json_property, t[0], json_property, t[1]))
					} else {
						parts = append(parts, fmt.Sprintf("%s >= $%d AND %s <= $%d", col, len(values)+1, col, len(values)+2))
						values = append(values, t[0], t[1])
					}
				}
			} else if f.GetDataType() == common.FieldDataType_DOUBLE {
				if t := f.GetNumber(); len(t) != 2 {
					return "", nil, errors.New("invalid number of operands")
				} else {
					if len(json_path) > 0 {
						parts = append(parts, fmt.Sprintf("jsonb_path_exists(%s, '%s ? (%s >= %f && %s <= %f)')", col, json_path, json_property, t[0], json_property, t[1]))
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

func getOrderBy(in *database.DbSelector, pathToDbPath PathToDbPathFunc) (string, error) {
	if in == nil {
		return "", nil
	}
	fields := in.GetSortBy()
	if len(fields) == 0 {
		return "", nil
	}

	var tmp strings.Builder
	tmp.WriteString("ORDER BY ")
	for i, s := range fields {
		path, ok := pathToDbPath(_DUMMY_FD.ConvertJsPathToPath(s.GetPath()))
		if !ok {
			return "", errors.New("unknown path: " + s.GetPath())
		}

		col, _, _, err := dbPathToDbSelector(path, false)
		if err != nil {
			return "", err
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

func addSingleOperandOperatorJson(parts *[]string, modelColumn string, jsonPath string, jsonProperty string, in *common.ListSelectorFilterBy, composeOpVal func(value string) (result string, invert bool)) error {
	switch in.GetDataType() {
	case common.FieldDataType_TEXT:
		if t := in.GetText(); len(t) != 1 {
			return errors.New("invalid number of operands")
		} else {
			composed, invert := composeOpVal(t[0])
			if invert {
				*parts = append(*parts, fmt.Sprintf("JSONB_PATH_EXISTS(%s, '%s ? (!(%s %s))')", modelColumn, jsonPath, jsonProperty, composed))
			} else {
				*parts = append(*parts, fmt.Sprintf("JSONB_PATH_EXISTS(%s, '%s ? (%s %s)')", modelColumn, jsonPath, jsonProperty, composed))
			}
		}
	case common.FieldDataType_INTEGER:
		if t := in.GetInteger(); len(t) != 1 {
			return errors.New("invalid number of operands")
		} else {
			op, _ := composeOpVal("")
			*parts = append(*parts, fmt.Sprintf("JSONB_PATH_EXISTS(%s, '%s ? (%s %s %d)')", modelColumn, jsonPath, jsonProperty, op, t[0]))
		}
	case common.FieldDataType_BOOLEAN:
		if t := in.GetBoolean(); len(t) != 1 {
			return errors.New("invalid number of operands")
		} else {
			b := "false"
			if t[0] {
				b = "true"
			}
			op, _ := composeOpVal("")
			*parts = append(*parts, fmt.Sprintf("JSONB_PATH_EXISTS(%s, '%s ? (%s %s %s)')", modelColumn, jsonPath, jsonProperty, op, b))
		}
	case common.FieldDataType_DOUBLE:
		if t := in.GetNumber(); len(t) != 1 {
			return errors.New("invalid number of operands")
		} else {
			op, _ := composeOpVal("")
			*parts = append(*parts, fmt.Sprintf("JSONB_PATH_EXISTS(%s, '%s ? (%s %s %f)')", modelColumn, jsonPath, jsonProperty, op, t[0]))
		}
	default:
		return errors.New("unsupported data type")
	}

	return nil
}

func addMultiOperandOperatorJson(parts *[]string, modelColumn string, jsonPath string, jsonProperty string, in *common.ListSelectorFilterBy, operator string) error {
	var vals string
	switch in.GetDataType() {
	case common.FieldDataType_TEXT:
		stringified := make([]string, 0, len(in.GetText()))
		for _, v := range in.GetText() {
			stringified = append(stringified, escapeForRegex(v))
		}
		vals = strings.Join(stringified, `", "`)
		if len(vals) > 0 {
			vals = `["` + vals + `"]`
		} else {
			vals = "[]"
		}
	case common.FieldDataType_INTEGER:
		stringified := make([]string, 0, len(in.GetInteger()))
		for _, v := range in.GetInteger() {
			stringified = append(stringified, strconv.FormatInt(v, 10))
		}
		vals = "[" + strings.Join(stringified, `, `) + "]"
	case common.FieldDataType_BOOLEAN:
		stringified := make([]string, 0, len(in.GetBoolean()))
		for _, v := range in.GetBoolean() {
			if v {
				stringified = append(stringified, "true")
			} else {
				stringified = append(stringified, "false")
			}
		}
		vals = "[" + strings.Join(stringified, `, `) + "]"
	case common.FieldDataType_DOUBLE:
		stringified := make([]string, 0, len(in.GetNumber()))
		for _, v := range in.GetNumber() {
			stringified = append(stringified, fmt.Sprintf("%f", v))
		}
		vals = "[" + strings.Join(stringified, `, `) + "]"
	default:
		return errors.New("unsupported data type")
	}

	*parts = append(*parts, fmt.Sprintf("JSONB_PATH_EXISTS(%s, '%s ? (%s %s $vals)', '{\"vals\": %s}')", modelColumn, jsonPath, jsonProperty, operator, vals))

	return nil
}

// dbPathToDbSelector converts a dbPath string into a column with optional -> selector, or column, jsonPath with propertyName triplet.
// Examples:
//
//	    -:$.level1.field.path         						...   N/A					$.level1.field.path        	@
//		column:$.level1.field.path        					...   column				$.level1.field.path        	@
//		alias:$.level1.field.path@.xx   					...   alias					$.level1.field.path	     	@.xx
//		table.column:$.level1.field.path@.xx  				...   table.column			$.level1.field.path	     	@.xx
//		schema.table.column:$.level1.field.path[*]@.xx  	...   schema.table.column	$.level1.field.path[*]	    @.xx
func dbPathToDbSelector(dbPath string, useDoubleArrow bool) (columnReference string, jsonPath string, propertyName string, err error) {
	parts := strings.SplitN(dbPath, ":", 2)
	if len(parts[0]) == 0 {
		err = errors.New("the dbPath must contain column name reference, got: " + dbPath)
		return
	} else if len(parts) == 1 {
		columnReference = dbPath
		return
	}

	column := parts[0]
	if column == "-" {
		err = errors.New("the field descriptor can't be used for filtering or sorting")
		return
	}

	path := parts[1]

	subSelectorParts := strings.SplitN(path, "@", 2)

	switch len(subSelectorParts) {
	case 1:
		// Simple -> or ->> path (no JSONPath sub-selector)
		if !strings.HasPrefix(path, "$.") {
			err = errors.New("the path must start with '$.', got: " + path)
			return
		}

		pathParts := strings.Split(path, ".")
		pathParts = pathParts[1:]

		switch {
		case len(pathParts) == 0:
			err = errors.New("the path must contain at least one part after '$.', got: " + path)
			return
		case !useDoubleArrow:
			columnReference = column + "->'" + strings.Join(pathParts, "'->'") + "'"
			return
		case len(pathParts) == 1:
			columnReference = column + "->>'" + pathParts[0] + "'"
			return
		default:
			columnReference = column + "->'" + strings.Join(pathParts[:len(pathParts)-1], "'->'") + "'->>'" + pathParts[len(pathParts)-1] + "'"
			return
		}

	case 2:
		// JSONPath query with embedded sub-selector (e.g., `column:$.path.to[*]@.sub`)
		columnReference = column
		jsonPath = subSelectorParts[0]
		propertyName = "@" + subSelectorParts[1]

		if !strings.HasPrefix(jsonPath, "$.") && jsonPath != "$" {
			err = errors.New("the JSONPath must start with '$.' or be equal to '$', got: " + jsonPath)
			return
		}

		return

	default:
		// More than one '@' not supported
		err = errors.New("the path must contain at most one '@', got: " + path)
		return
	}
}
