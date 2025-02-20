package common

import (
	"strconv"
	"strings"
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

func (fd *FieldDescriptor) WithTooltip(tooltip string) *FieldDescriptor {
	fd.SetTooltip(tooltip)
	return fd
}

func (fd *FieldDescriptor) WithDouble(precision *uint, unit string, displayFormat *FieldDisplayFormat) *FieldDescriptor {
	fd.SetDataType(FieldDataType_DOUBLE)
	if precision != nil {
		fd.SetPrecision(int32(*precision))
	} else {
		fd.ClearPrecision()
	}
	if unit != "" {
		fd.SetUnit(unit)
	} else {
		fd.ClearUnit()
	}
	if displayFormat != nil {
		switch *displayFormat {
		case FieldDisplayFormat_DEFAULT:
		case FieldDisplayFormat_DURATION:
		case FieldDisplayFormat_INTERVAL:
		default:
			panic("displayFormat is not supported for double")
		}
		fd.SetFormat(*displayFormat)
	} else {
		fd.SetFormat(FieldDisplayFormat_DEFAULT)
	}
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
	if displayFormat != nil {
		switch *displayFormat {
		case FieldDisplayFormat_DEFAULT:
		case FieldDisplayFormat_DURATION:
		case FieldDisplayFormat_INTERVAL:
		case FieldDisplayFormat_MONTH:
		default:
			panic("displayFormat is not supported for double")
		}
		fd.SetFormat(*displayFormat)
	} else {
		fd.SetFormat(FieldDisplayFormat_DEFAULT)
	}
	return fd
}

func (fd *FieldDescriptor) WithTimestamp(displayFormat *FieldDisplayFormat) *FieldDescriptor {
	fd.SetDataType(FieldDataType_TIMESTAMP)
	fd.ClearPrecision()
	fd.ClearUnit()
	if displayFormat != nil {
		switch *displayFormat {
		case FieldDisplayFormat_DEFAULT:
		case FieldDisplayFormat_DATE:
		case FieldDisplayFormat_UTC_DATE:
		case FieldDisplayFormat_DAYOFWEEK:
		case FieldDisplayFormat_TIMEOFDAY:
		default:
			panic("displayFormat is not supported for double")
		}
		fd.SetFormat(*displayFormat)
	} else {
		fd.SetFormat(FieldDisplayFormat_DEFAULT)
	}
	return fd
}

func (fd *FieldDescriptor) WithString(displayFormat *FieldDisplayFormat) *FieldDescriptor {
	fd.SetDataType(FieldDataType_TEXT)
	fd.ClearPrecision()
	fd.ClearUnit()
	if displayFormat != nil {
		switch *displayFormat {
		case FieldDisplayFormat_DEFAULT:
		case FieldDisplayFormat_MULTILINE:
		case FieldDisplayFormat_DATE:
		case FieldDisplayFormat_UTC_DATE:
		case FieldDisplayFormat_TIMEOFDAY:
		case FieldDisplayFormat_PASSWORD:
		default:
			panic("displayFormat is not supported for double")
		}
		fd.SetFormat(*displayFormat)
	} else {
		fd.SetFormat(FieldDisplayFormat_DEFAULT)
	}
	return fd
}

func (fd *FieldDescriptor) WithInterval() *FieldDescriptor {
	fd.SetDataType(FieldDataType_TIMESTAMP)
	fd.ClearPrecision()
	fd.ClearUnit()
	fd.SetFormat(FieldDisplayFormat_INTERVAL)
	fd.SetMultiValue(true)
	return fd
}

func (fd *FieldDescriptor) WithRe(re string) *FieldDescriptor {
	v := fd.GetValidation()
	if v == nil {
		v = FieldValidation_builder{}.Build()
		fd.SetValidation(v)
	}
	if len(re) == 0 {
		v.ClearRe()
	} else {
		v.SetRe(re)
	}
	return fd
}

func (fd *FieldDescriptor) WithMinInteger(min int) *FieldDescriptor {
	v := fd.GetValidation()
	if v == nil {
		v = FieldValidation_builder{}.Build()
		fd.SetValidation(v)
	}
	v.SetMinInteger(int64(min))
	return fd
}

func (fd *FieldDescriptor) WithMaxInteger(max int) *FieldDescriptor {
	v := fd.GetValidation()
	if v == nil {
		v = FieldValidation_builder{}.Build()
		fd.SetValidation(v)
	}
	v.SetMaxInteger(int64(max))
	return fd
}

func (fd *FieldDescriptor) WithMinNumber(min int) *FieldDescriptor {
	v := fd.GetValidation()
	if v == nil {
		v = FieldValidation_builder{}.Build()
		fd.SetValidation(v)
	}
	v.SetMinNumber(float64(min))
	return fd
}

func (fd *FieldDescriptor) WithMaxNumber(max int) *FieldDescriptor {
	v := fd.GetValidation()
	if v == nil {
		v = FieldValidation_builder{}.Build()
		fd.SetValidation(v)
	}
	v.SetMaxNumber(float64(max))
	return fd
}

func (fd *FieldDescriptor) WithMaxLength(maxLength int) *FieldDescriptor {
	v := fd.GetValidation()
	if v == nil {
		v = FieldValidation_builder{}.Build()
		fd.SetValidation(v)
	}
	if maxLength == 0 {
		v.ClearMaxLength()
	} else {
		v.SetMaxLength(int32(maxLength))
	}
	return fd
}

func (fd *FieldDescriptor) WithOptions(options map[string]string) *FieldDescriptor {
	v := fd.GetValidation()
	if v == nil {
		v = FieldValidation_builder{}.Build()
		fd.SetValidation(v)
	}
	v.SetOptions(options)
	return fd
}

func (fd *FieldDescriptor) WithIntegerOptions(options map[int32]string) *FieldDescriptor {
	fd.SetDataType(FieldDataType_INTEGER)
	v := fd.GetValidation()
	if v == nil {
		v = FieldValidation_builder{}.Build()
		fd.SetValidation(v)
	}
	tmp := make(map[string]string, len(options))
	for k, v := range options {
		tmp[strconv.Itoa(int(k))] = v
	}
	v.SetOptions(tmp)
	return fd
}
