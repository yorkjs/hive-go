package hive

import (
	"fmt"
	"strings"
)

// value 毫秒
func FormatDuration(value int64) string {
	var result []string

	data := NormalizeDuration(value)
	if data.Days > 0 {
		result = append(result, fmt.Sprintf("%d日", data.Days))
	}
	if data.Hours > 0 {
		result = append(result, fmt.Sprintf("%d时", data.Hours))
	}
	if data.Minutes > 0 {
		result = append(result, fmt.Sprintf("%d分", data.Minutes))
	}
	if data.Seconds > 0 {
		result = append(result, fmt.Sprintf("%d秒", data.Seconds))
	}

	return strings.Join(result, "")
}
