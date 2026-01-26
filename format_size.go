package hive

import (
	"fmt"
)

func FormatSize(value float64) string {
	if value >= SIZE_GB {
		result := DivideNumber(value, SIZE_GB)
		d := 2
		if IsInteger(result) {
			d = 0
		}
		return fmt.Sprintf("%sGB", TruncateNumber(result, d))
	} else if value >= SIZE_MB {
		result := DivideNumber(value, SIZE_MB)
		d := 2
		if IsInteger(result) {
			d = 0
		}
		return fmt.Sprintf("%sMB", TruncateNumber(result, d))
	} else if value >= SIZE_KB {
		result := DivideNumber(value, SIZE_KB)
		d := 2
		if IsInteger(result) {
			d = 0
		}
		return fmt.Sprintf("%sKB", TruncateNumber(result, d))
	}
	return fmt.Sprintf("%vB", value)
}
