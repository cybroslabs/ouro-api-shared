package common

import (
	"strings"
)

func NewFieldDescriptor(fieldId string, label string, groupId string, required bool, editable bool, visible bool, multiValue bool, secured bool) *FieldDescriptor {
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
