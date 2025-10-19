package system

import (
	"os"
	"runtime/debug"
	"strconv"
)

func init() {
	ml, ok := os.LookupEnv("GOMEMLIMIT")
	if !ok {
		return
	}
	if ml == "off" {
		return
	}
	ml_percent, ok := os.LookupEnv("GOMEMLIMIT_PERCENTAGE")
	if !ok {
		return
	}
	limit, err := strconv.ParseFloat(ml, 64)
	if err != nil || limit <= 0 {
		return
	}
	percent, err := strconv.ParseFloat(ml_percent, 64)
	if err != nil || percent <= 0 || percent >= 100 {
		return
	}
	limit = limit * percent / 100.0
	if limit <= 1 {
		return
	}
	debug.SetMemoryLimit(int64(limit * percent / 100.0))
}
