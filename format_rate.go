package hive

import "fmt"

// FormatRatePercent 把万分比格式化为百分比
func FormatRatePercent(value int) string {
	return fmt.Sprintf("%v%%", RateToDisplay(value))
}
