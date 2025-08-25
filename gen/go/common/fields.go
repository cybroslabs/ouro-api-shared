package common

import (
	"errors"
	"fmt"
	"regexp"
	"slices"
	"strconv"
	"strings"

	"github.com/rmg/iso4217"
	"k8s.io/utils/ptr"
)

var (
	reCapUnderscores = regexp.MustCompile(`^[A-Z]+(?:_[A-Z]+)*$`)
	reCapitals       = regexp.MustCompile(`([A-Z]+)`)
)

// NewFieldDescriptorInternal creates a new FieldDescriptorInternal with the given parameters.
// It is used internally to create a FieldDescriptor with additional metadata like ID, group, and database path.
// The group is a field descriptor group key used to clean up removed descriptors.
// The dbPath is the database column name or JSONB path. The JSON path must start with '$.' to be registered as a JSONB path.
func NewFieldDescriptorInternal(dbPath string, descriptor *FieldDescriptor, customGidSuffix *string, customGroupSuffix *string) (*FieldDescriptorInternal, error) {
	if strings.Contains(dbPath, ".{") {
		return nil, errors.New("dbPath must not contain case selector brackets, use jsPath in FieldDescriptor instead")
	}

	if len(dbPath) == 0 {
		return nil, errors.New("dbPath must not be empty")
	}

	var group_suffix string
	if descriptor.GetIsUserDefined() {
		group_suffix = "user"
	} else {
		group_suffix = "sys"
	}

	object_type := descriptor.GetObjectType()

	// Auto-generate the system-wide unique identifier (gid) for the field descriptor.
	var gid string
	if cgs := ptr.Deref(customGidSuffix, ""); len(cgs) > 0 {
		gid = strings.ToLower(fmt.Sprintf("%s#%s#%s", object_type, descriptor.GetPath(), cgs))
	} else {
		gid = strings.ToLower(fmt.Sprintf("%s#%s", object_type, descriptor.GetPath()))
	}
	descriptor.SetGid(gid)

	// Generate the field descriptor group. It's used to combine built-in vs user-defined fields for specific object type.
	var group string
	if cgs := ptr.Deref(customGroupSuffix, ""); len(cgs) > 0 {
		group = strings.ToLower(fmt.Sprintf("%s#%s#%s", object_type, cgs, group_suffix))
	} else {
		group = strings.ToLower(fmt.Sprintf("%s#%s", object_type, group_suffix))
	}

	if !descriptor.HasSortable() {
		descriptor.SetSortable(true) // Default is true, but we set it explicitly.
	}
	if !descriptor.HasFilterable() {
		descriptor.SetFilterable(true) // Default is true, but we set it explicitly.
	}

	return FieldDescriptorInternal_builder{
		Group:           ptr.To(group),
		DbPath:          ptr.To(dbPath),
		FieldDescriptor: descriptor,
	}.Build(), nil
}

// NewFieldDescriptorInternal creates a new FieldDescriptorInternal with the given parameters.
// It is used internally to create a FieldDescriptorInternal wrapper over the FieldDescriptor with additional metadata like group or database path.
// The dbPath is the database column name or JSONB path. The JSON path must start with '$.' to be registered as a JSONB path.
// MustNewFieldDescriptorInternal creates a new FieldDescriptorInternal and panics if there is an error.
func MustNewFieldDescriptorInternal(dbPath string, descriptor *FieldDescriptor, customGidSuffix *string, customGroupSuffix *string) *FieldDescriptorInternal {
	fd, err := NewFieldDescriptorInternal(dbPath, descriptor, customGidSuffix, customGroupSuffix)
	if err != nil {
		panic(err.Error())
	}
	return fd
}

// NewFieldDescriptor creates a new FieldDescriptor with the given parameters.
// It auto-generates the fieldId from the jsPath if not provided.
// The jsPath must not be empty, and the fieldId must not start with '$.'.
// The label must not be empty.
// The jsPath is a JavaScript/TypeScript path to the field in given object type. It is a segment path like "name.first" or "address.street". Additionally, it shall contain case selector value in the brackets, like "seg.seg.holder.{option}.more".
// The path is later automatically converted from the jsPath where holder is replaced directly with the option value.
func NewFieldDescriptor(objectType ObjectType, fieldId string, jsPath string, label string, groupId string, required bool, editable bool, visible bool, multiValue bool, secured bool, sortable bool, filterable bool) *FieldDescriptor {
	if _, known := ObjectType_name[int32(objectType)]; !known {
		panic(fmt.Sprintf("unknown objectType: %s", objectType))
	}

	jsPath = strings.TrimSpace(jsPath)
	if jsPath == "" {
		panic("jsPath is required")
	}

	label = strings.TrimSpace(label)
	if label == "" {
		panic("label is required")
	}

	fd := FieldDescriptor_builder{
		ObjectType: objectType.Enum(),
		JsPath:     &jsPath,
		Label:      &label,
		DataType:   FieldDataType_TEXT.Enum(),
		GroupId:    &groupId,
		Required:   &required,
		Editable:   &editable,
		Visible:    &visible,
		MultiValue: &multiValue,
		Secured:    &secured,
		Format:     FieldDisplayFormat_DISPLAY_FORMAT_UNSPECIFIED.Enum(),
		Sortable:   &sortable,
		Filterable: &filterable,
	}.Build()

	path := fd.ConvertJsPathToPath(jsPath)
	fd.SetPath(path)

	fieldId = strings.TrimSpace(fieldId)
	if fieldId == "" {
		// Auto-generate fieldId as the path
		fieldId = path
	}
	fd.SetFieldId(fieldId)

	return fd
}

var validFormats = map[FieldDataType][]FieldDisplayFormat{
	FieldDataType_TEXT:      {FieldDisplayFormat_MULTILINE, FieldDisplayFormat_PASSWORD, FieldDisplayFormat_COMBO},
	FieldDataType_INTEGER:   {FieldDisplayFormat_MONEY, FieldDisplayFormat_TIMEOFDAY, FieldDisplayFormat_COMBO},
	FieldDataType_DOUBLE:    {FieldDisplayFormat_MONEY},
	FieldDataType_TIMESTAMP: {FieldDisplayFormat_DATE_ONLY, FieldDisplayFormat_UTC_DATETIME, FieldDisplayFormat_UTC_DATE_ONLY},
	FieldDataType_BOOLEAN:   {},
}

func validateDisplayFormat(dataType FieldDataType, format *FieldDisplayFormat) FieldDisplayFormat {
	if ptr.Deref(format, FieldDisplayFormat_DISPLAY_FORMAT_UNSPECIFIED) == FieldDisplayFormat_DISPLAY_FORMAT_UNSPECIFIED {
		return FieldDisplayFormat_DISPLAY_FORMAT_UNSPECIFIED
	}
	allowed, ok := validFormats[dataType]
	if !ok {
		return FieldDisplayFormat_DISPLAY_FORMAT_UNSPECIFIED
	}
	if slices.Contains(allowed, *format) {
		return *format
	}
	panic("displayFormat is not supported")
}

// WithTooltip sets the field's tooltip.
func (fd *FieldDescriptor) WithTooltip(tooltip string) *FieldDescriptor {
	fd.SetTooltip(tooltip)
	return fd
}

// WithDouble sets the field to a double type with the given precision and unit.
// The precision is the number of decimal places, if it's less than -15, it's ignored.
// The unit is the unit of the value, it can be empty.
// The displayFormat is the format of the value, it can be nil.
// If the displayFormat is MONEY, the unit must be a valid ISO 4217 currency code.
func (fd *FieldDescriptor) WithDouble(precision int32, unit string, displayFormat *FieldDisplayFormat) *FieldDescriptor {
	fd.SetDataType(FieldDataType_DOUBLE)
	if precision > -15 {
		fd.SetPrecision(precision)
	} else {
		fd.ClearPrecision()
	}
	if unit != "" {
		fd.SetUnit(unit)
	} else {
		fd.ClearUnit()
	}
	fmt := validateDisplayFormat(FieldDataType_DOUBLE, displayFormat)
	if fmt == FieldDisplayFormat_MONEY {
		if code, _ := iso4217.ByName(unit); code == 0 {
			panic("unit is not a valid ISO 4217 currency code")
		}
	}
	fd.SetFormat(fmt)
	return fd
}

func (fd *FieldDescriptor) WithInteger(unit string, displayFormat *FieldDisplayFormat) *FieldDescriptor {
	fd.SetDataType(FieldDataType_INTEGER)
	fd.ClearPrecision()
	if unit != "" {
		fd.SetUnit(unit)
	} else {
		fd.ClearUnit()
	}
	fmt := validateDisplayFormat(FieldDataType_INTEGER, displayFormat)
	if fmt == FieldDisplayFormat_MONEY {
		if code, _ := iso4217.ByName(unit); code == 0 {
			panic("unit is not a valid ISO 4217 currency code")
		}
	}
	fd.SetFormat(fmt)
	return fd
}

func (fd *FieldDescriptor) WithTimestamp(displayFormat *FieldDisplayFormat) *FieldDescriptor {
	fd.SetDataType(FieldDataType_TIMESTAMP)
	fd.ClearPrecision()
	fd.ClearUnit()
	fd.SetFormat(validateDisplayFormat(FieldDataType_TIMESTAMP, displayFormat))
	return fd
}

func (fd *FieldDescriptor) WithString(displayFormat *FieldDisplayFormat) *FieldDescriptor {
	fd.SetDataType(FieldDataType_TEXT)
	fd.ClearPrecision()
	fd.ClearUnit()
	fd.SetFormat(validateDisplayFormat(FieldDataType_TEXT, displayFormat))
	return fd
}

func (fd *FieldDescriptor) WithBool() *FieldDescriptor {
	fd.SetDataType(FieldDataType_BOOLEAN)
	fd.ClearPrecision()
	fd.ClearUnit()
	fd.ClearFormat()
	return fd
}

func (fd *FieldDescriptor) WithDateTime(displayFormat *FieldDisplayFormat) *FieldDescriptor {
	fd.SetDataType(FieldDataType_TIMESTAMP)
	fd.ClearPrecision()
	fd.ClearUnit()
	fd.SetFormat(validateDisplayFormat(FieldDataType_TIMESTAMP, displayFormat))
	return fd
}

func (fd *FieldDescriptor) WithDuration(displayFormat *FieldDisplayFormat) *FieldDescriptor {
	fd.SetDataType(FieldDataType_DURATION)
	fd.ClearPrecision()
	fd.ClearUnit()
	fd.SetFormat(validateDisplayFormat(FieldDataType_DURATION, displayFormat))
	return fd
}

func (fd *FieldDescriptor) ensureValidation() *FieldValidation {
	v := fd.GetValidation()
	if v != nil {
		return v
	}
	v = FieldValidation_builder{}.Build()
	fd.SetValidation(v)
	return v
}

func (fd *FieldDescriptor) WithRe(re string) *FieldDescriptor {
	if len(re) == 0 {
		fd.ensureValidation().ClearRe()
	} else {
		fd.ensureValidation().SetRe(re)
	}
	return fd
}

func (fd *FieldDescriptor) WithMin(min int) *FieldDescriptor {
	if fd.GetDataType() == FieldDataType_INTEGER {
		fd.ensureValidation().SetMinInteger(int64(min))
	} else if fd.GetDataType() == FieldDataType_DOUBLE {
		fd.ensureValidation().SetMinNumber(float64(min))
	} else {
		panic("FieldDataType is not INTEGER or DOUBLE")
	}
	return fd
}

func (fd *FieldDescriptor) WithMax(max int) *FieldDescriptor {
	if fd.GetDataType() == FieldDataType_INTEGER {
		fd.ensureValidation().SetMaxInteger(int64(max))
	} else if fd.GetDataType() == FieldDataType_DOUBLE {
		fd.ensureValidation().SetMaxNumber(float64(max))
	} else {
		panic("FieldDataType is not INTEGER or DOUBLE")
	}
	return fd
}

func (fd *FieldDescriptor) WithMinNumber(min float64) *FieldDescriptor {
	if fd.GetDataType() != FieldDataType_DOUBLE {
		panic("FieldDataType is not DOUBLE")
	}
	fd.ensureValidation().SetMinNumber(min)
	return fd
}

func (fd *FieldDescriptor) WithMaxNumber(max float64) *FieldDescriptor {
	if fd.GetDataType() != FieldDataType_DOUBLE {
		panic("FieldDataType is not DOUBLE")
	}
	fd.ensureValidation().SetMaxNumber(max)
	return fd
}

func (fd *FieldDescriptor) WithMaxLength(maxLength int) *FieldDescriptor {
	if maxLength == 0 {
		fd.ensureValidation().ClearMaxLength()
	} else {
		fd.ensureValidation().SetMaxLength(int32(maxLength))
	}
	return fd
}

func (fd *FieldDescriptor) WithOptions(options map[string]string) *FieldDescriptor {
	if options == nil {
		fd.ClearFormat()
		fd.ensureValidation().SetOptions(nil)
		return fd
	}

	if fd.GetDataType() != FieldDataType_INTEGER && fd.GetDataType() != FieldDataType_TEXT {
		panic("Options are only supported for INTEGER or TEXT fields")
	}
	if fd.GetDataType() == FieldDataType_INTEGER {
		// Validate indexes
		for k := range options {
			if _, err := strconv.Atoi(k); err != nil {
				panic("Options keys must be string-integers when FieldDataType is INTEGER")
			}
		}
	}

	fd.SetFormat(FieldDisplayFormat_COMBO)
	validation := fd.ensureValidation()
	validation.SetOptions(options)
	validation.ClearOptionsSource()
	return fd
}

// WithIntegerOptions sets the field to an INTEGER type with the given options.
// The options are a map of integer values to string labels.
// The keys of the map must be integers, and the values are the labels for the options
// It's typically used for protobuf enum values.
func (fd *FieldDescriptor) WithIntegerOptions(options map[int32]string) *FieldDescriptor {
	if options == nil {
		fd.ClearFormat()
		fd.ensureValidation().SetOptions(nil)
		return fd
	}

	fd.SetDataType(FieldDataType_INTEGER)
	tmp := make(map[string]string, len(options))

	// Detect enum names format.
	// If all options are in the format of ALL_CAPS_UNDERSCORES, then we will convert them to a human-readable format with spaces.
	// If all options are in the format of PascalCase, we will prefix all capital letters with a space to make them more readable.
	all_cap_underscores := true
	no_space := true
	for k, v := range options {
		tmp[strconv.FormatInt(int64(k), 10)] = v
		if all_cap_underscores && !reCapUnderscores.MatchString(v) {
			all_cap_underscores = false
		}
		if no_space && strings.Contains(v, " ") {
			no_space = false
		}
	}

	if all_cap_underscores {
		for k := range tmp {
			parts := strings.Split(tmp[k], "_")
			for i, part := range parts {
				// Capitalize the first letter of each part.
				if len(part) > 0 {
					parts[i] = strings.ToUpper(part[:1]) + part[1:]
				}
			}
			tmp[k] = strings.Join(parts, " ")
		}
	} else if no_space {
		for k := range tmp {
			// Replace PascalCase values 'v' by prefixing all capital letters with a space.
			tmp[k] = strings.TrimSpace(reCapitals.ReplaceAllString(tmp[k], " $1"))
		}
	}

	fd.SetFormat(FieldDisplayFormat_COMBO)
	validation := fd.ensureValidation()
	validation.SetOptions(tmp)
	validation.ClearOptionsSource()
	return fd
}

type EnumWithString interface {
	fmt.Stringer
	~int32
}

func CreateOptions[T EnumWithString](enumMap map[int32]string) map[string]string {
	if enumMap == nil {
		return nil
	}

	result := make(map[string]string, len(enumMap))
	for val, name := range enumMap {
		x := T(val)
		result[x.String()] = name
	}
	return result
}

func (fd *FieldDescriptor) WithOptionsSource(source string) *FieldDescriptor {
	if source == "" {
		fd.ClearFormat()
		fd.ensureValidation().ClearOptionsSource()
		return fd
	}

	if fd.GetDataType() != FieldDataType_TEXT {
		panic("Options source is only supported for TEXT fields")
	}

	fd.SetFormat(FieldDisplayFormat_COMBO)
	validation := fd.ensureValidation()
	validation.SetOptions(nil)
	validation.SetOptionsSource(source)
	return fd
}

func (fd *FieldDescriptor) WithDefaultValue(value *FieldValue) *FieldDescriptor {
	if value == nil {
		fd.ClearDefaultValue()
		return fd
	}

	if err := fd.Validate(value); err != nil {
		panic(fmt.Sprintf("default value is invalid: %s", err.Error()))
	}

	fd.SetDefaultValue(value)
	return fd
}

func (fd *FieldDescriptor) WithValidation(validation *FieldValidation) *FieldDescriptor {
	if validation == nil {
		fd.ClearValidation()
		return fd
	}

	fd.SetValidation(validation)
	return fd
}

func (fd *FieldDescriptor) Validate(value *FieldValue) error {
	if fd == nil {
		return nil
	}

	switch fd.GetDataType() {
	case FieldDataType_TEXT:
		if value.WhichKind() != FieldValue_StringValue_case {
			return errors.New("the value must be a string")
		}
		if validation := fd.GetValidation(); validation != nil {
			if validation.HasMaxLength() && (len(value.GetStringValue()) > int(validation.GetMaxLength())) {
				return errors.New("the value is too long")
			}
			if validation.HasMinLength() && (len(value.GetStringValue()) < int(validation.GetMinLength())) {
				return errors.New("the value is too short")
			}
			if validation.HasRe() {
				if re, err := regexp.Compile(validation.GetRe()); err == nil && !re.MatchString(value.GetStringValue()) {
					return errors.New("the value does not match the regular expression format")
				}
			}
		}

	case FieldDataType_INTEGER:
		if value.WhichKind() != FieldValue_IntegerValue_case {
			return errors.New("the value must be an integer")
		}
		if validation := fd.GetValidation(); validation != nil {
			if validation.HasMaxInteger() && (value.GetIntegerValue() > validation.GetMaxInteger()) {
				return errors.New("the value is too high")
			}
			if validation.HasMinInteger() && (value.GetIntegerValue() < validation.GetMinInteger()) {
				return errors.New("the value is too low")
			}
			if validation.HasRe() {
				if re, err := regexp.Compile(validation.GetRe()); err == nil && !re.MatchString(fmt.Sprintf("%d", value.GetIntegerValue())) {
					return errors.New("the value does not match the regular expression format")
				}
			}
		}

	case FieldDataType_DOUBLE:
		if value.WhichKind() != FieldValue_DoubleValue_case {
			return errors.New("the value must be a number")
		}
		if validation := fd.GetValidation(); validation != nil {
			if validation.HasMaxNumber() && (value.GetDoubleValue() > validation.GetMaxNumber()) {
				return errors.New("the value is too high")
			}
			if validation.HasMinNumber() && (value.GetDoubleValue() < validation.GetMinNumber()) {
				return errors.New("the value is too low")
			}
			if validation.HasRe() {
				if re, err := regexp.Compile(validation.GetRe()); err == nil && !re.MatchString(fmt.Sprintf("%f", value.GetDoubleValue())) {
					return errors.New("the value does not match the regular expression format")
				}
			}
		}

	case FieldDataType_DURATION:
		if value.WhichKind() != FieldValue_DurationValue_case {
			return errors.New("the value must be a duration")
		}

	case FieldDataType_BOOLEAN:
		if value.WhichKind() != FieldValue_BoolValue_case {
			return errors.New("the value must be a boolean")
		}

	case FieldDataType_TIMESTAMP:
		if value.WhichKind() != FieldValue_DateValue_case {
			return errors.New("the value must be a date")
		}

	case FieldDataType_BINARY:
		if value.WhichKind() != FieldValue_BinaryValue_case {
			return errors.New("the value must be a binary value")
		}

	default:
		return errors.ErrUnsupported
	}

	return nil
}

// GenerateJsPath generates a JavaScript/TypeScript path for the field descriptor.
func (fd *FieldDescriptor) GenerateJsPath(prefix string) string {
	if fd == nil {
		return ""
	}
	tmp := prefix + fd.GetFieldId()
	switch fd.GetDataType() {
	case FieldDataType_TEXT:
		tmp += ".kind.{stringValue}"
	case FieldDataType_INTEGER:
		tmp += ".kind.{integerValue}"
	case FieldDataType_BOOLEAN:
		tmp += ".kind.{boolValue}"
	case FieldDataType_DOUBLE:
		tmp += ".kind.{doubleValue}"
	case FieldDataType_BINARY:
		tmp += ".kind.{binaryValue}"
	case FieldDataType_TIMESTAMP:
		tmp += ".kind.{dateValue}"
	case FieldDataType_DURATION:
		tmp += ".kind.{durationValue}"
	default:
		return ""
	}
	return tmp
}

var (
	_re_js_path = regexp.MustCompile(`(^|\.)[^\.\{\}]+\.\{([^\}]+)\}`)
)

// The method converts a JS path to a standard path format.
// The JS path is extended covering case selector where value is defined in the brackets and the case-wrapper is the previous path segment.
// In the path the wrapper is not present and the case-value is directly used as a path segment.
// Example: 'any.dot.path.hide.{selector}.more' -> 'any.dot.path.selector.more'
func (fd *FieldDescriptor) ConvertJsPathToPath(jsPath string) string {
	return _re_js_path.ReplaceAllString(jsPath, "$1$2")
}

// The method validates the field values against the field descriptors.
// It returns an error if any of the field values are invalid or if any required fields are missing.
// It also sets the default values for any fields that are not present in the values map.
// If the values map is nil, it will be initialized with the default values for all fields that have them and the initialized flag will be set to true.
func ValidateFields(descriptors []*FieldDescriptor, values *map[string]*FieldValue) (initialized bool, err error) {
	if values == nil {
		return false, errors.New("unable to validate nil value map")
	}

	fd_map := make(map[string]*FieldDescriptor, len(descriptors))
	for _, fd := range descriptors {
		fd_map[fd.GetFieldId()] = fd
	}

	for field_id, field := range *values {
		fd, ok := fd_map[field_id]
		if !ok {
			return false, fmt.Errorf("field %s is not defined", field_id)
		}
		if err := fd.Validate(field); err != nil {
			return false, fmt.Errorf("field %s: %w", field_id, err)
		}
		delete(fd_map, field_id)
	}

	for field_id, descriptor := range fd_map {
		if descriptor.HasDefaultValue() {
			if *values == nil {
				initialized = true
				*values = map[string]*FieldValue{
					field_id: descriptor.GetDefaultValue(),
				}
			} else {
				(*values)[field_id] = descriptor.GetDefaultValue()
			}
			continue
		}
		if descriptor.GetRequired() {
			return initialized, fmt.Errorf("field %s is required", field_id)
		}
	}

	return
}

// GetAnyValue returns the value of the FieldValue as an interface{}.
func (x *FieldValue) GetAnyValue() any {
	if x != nil {
		switch x.WhichKind() {
		case FieldValue_StringValue_case:
			return x.GetStringValue()
		case FieldValue_IntegerValue_case:
			return x.GetIntegerValue()
		case FieldValue_DoubleValue_case:
			return x.GetDoubleValue()
		case FieldValue_BoolValue_case:
			return x.GetBoolValue()
		case FieldValue_DateValue_case:
			return x.GetDateValue()
		case FieldValue_DurationValue_case:
			return x.GetDurationValue()
		case FieldValue_BinaryValue_case:
			return x.GetBinaryValue()
		default:
			return nil
		}
	}
	return nil
}
