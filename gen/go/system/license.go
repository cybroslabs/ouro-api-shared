package system

import (
	"encoding/json"
	"slices"
	"strconv"
	"time"
)

type LicensedItem string

const (
	LicensedItemParallelAcqusitionChannels LicensedItem = "parallel_acqusition_channels"
	LicensedItemDeviceDrivers              LicensedItem = "device_drivers"
)

// IsLicensed returns true if the licensed item is present in the license.
func (l *License) IsLicensed(item LicensedItem) bool {
	if opts := l.GetOptions(); opts == nil {
		return false
	} else {
		_, ok := opts[string(item)]
		return ok
	}
}

// IsLicensedStringArray checks if the option is present in the license for the given item.
func (l *License) IsLicensedOneOf(item LicensedItem, option string) bool {
	if opts := l.GetOptions(); opts != nil {
		if v, ok := opts[string(item)]; ok {
			var result []string
			if err := json.Unmarshal([]byte(v), &result); err == nil {
				return slices.Contains(result, option)
			}
		}
	}
	return false
}

// GetLicensedString returns the numberic count of the licensed item, or 0 if not licensed.
func (l *License) GetLicensedCount(item LicensedItem) int {
	if opts := l.GetOptions(); opts != nil {
		if v, ok := opts[string(item)]; ok {
			if v_int, err := strconv.Atoi(v); err == nil {
				return v_int
			}
		}
	}
	return 0
}

// GetLicensedString returns the array of string as set in the license, or an empty array if not licensed.
func (l *License) GetLicensedStringArray(item LicensedItem) []string {
	result := make([]string, 0)
	if opts := l.GetOptions(); opts != nil {
		if v, ok := opts[string(item)]; ok {
			if err := json.Unmarshal([]byte(v), &result); err == nil {
				return result
			}
		}
	}
	return result
}

// GetLicensedTimestamp returns the timestamp as set in the license, or zero time if not licensed.
func (l *License) GetLicensedTimestamp(item LicensedItem) time.Time {
	if opts := l.GetOptions(); opts != nil {
		if v, ok := opts[string(item)]; ok {
			if v_int, err := strconv.Atoi(v); err == nil {
				return time.Unix(int64(v_int), 0)
			} else if v_time, err := time.Parse(time.RFC3339, v); err == nil {
				return v_time
			}
		}
	}
	return time.Time{}
}
