package common

import (
	"fmt"
	"slices"
	"strconv"
	"strings"

	"github.com/rmg/iso4217"
)

func NewFieldDescriptor(fieldId string, jsPath string, label string, groupId string, required bool, editable bool, visible bool, multiValue bool, secured bool) *FieldDescriptor {
	fieldId = strings.TrimSpace(fieldId)
	if fieldId == "" {
		panic("fieldId is required")
	}
	label = strings.TrimSpace(label)
	if label == "" {
		panic("label is required")
	}
	fd := FieldDescriptor_builder{
		FieldId:    &fieldId,
		JsPath:     &jsPath,
		Label:      &label,
		DataType:   FieldDataType_TEXT.Enum(),
		GroupId:    &groupId,
		Required:   &required,
		Editable:   &editable,
		Visible:    &visible,
		MultiValue: &multiValue,
		Secured:    &secured,
		Format:     FieldDisplayFormat_DEFAULT.Enum(),
	}.Build()
	return fd
}

var validFormats = map[FieldDataType][]FieldDisplayFormat{
	FieldDataType_TEXT:      {FieldDisplayFormat_MULTILINE, FieldDisplayFormat_PASSWORD},
	FieldDataType_INTEGER:   {FieldDisplayFormat_MONEY, FieldDisplayFormat_TIMEOFDAY},
	FieldDataType_DOUBLE:    {FieldDisplayFormat_MONEY},
	FieldDataType_TIMESTAMP: {FieldDisplayFormat_DATE_ONLY, FieldDisplayFormat_UTC_DATETIME, FieldDisplayFormat_UTC_DATE_ONLY},
	FieldDataType_BOOLEAN:   {FieldDisplayFormat_DEFAULT},
}

func validateDisplayFormat(dataType FieldDataType, format *FieldDisplayFormat) FieldDisplayFormat {
	if format == nil || *format == FieldDisplayFormat_DEFAULT {
		return FieldDisplayFormat_DEFAULT
	}
	allowed, ok := validFormats[dataType]
	if !ok {
		return FieldDisplayFormat_DEFAULT
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
	fd.SetFormat(FieldDisplayFormat_DEFAULT)
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
	if fd.GetDataType() != FieldDataType_INTEGER && fd.GetDataType() != FieldDataType_TEXT {
		panic("Options are only supported for INTEGER or TEXT fields")
	}
	if fd.GetDataType() != FieldDataType_INTEGER {
		// Validate indexes
		for k := range options {
			if _, err := strconv.Atoi(k); err != nil {
				panic("Options keys must be string-integers when FieldDataType is INTEGER")
			}
		}
	}
	fd.ensureValidation().SetOptions(options)
	return fd
}

func (fd *FieldDescriptor) WithIntegerOptions(options map[int32]string) *FieldDescriptor {
	fd.SetDataType(FieldDataType_INTEGER)
	tmp := make(map[string]string, len(options))
	for k, v := range options {
		tmp[fmt.Sprintf("%d", k)] = v
	}
	fd.ensureValidation().SetOptions(tmp)
	return fd
}
