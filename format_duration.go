package hive

import (
	"fmt"
	"strings"
)

func FormatDuration(milliseconds int64) string {
	var result []string

	data := NormalizeDuration(milliseconds)
	if data.Days > 0 {
		result = append(result, fmt.Sprintf("%d天", data.Days))
	}
	if data.Hours > 0 {
		result = append(result, fmt.Sprintf("%d小时", data.Hours))
	}
	if data.Minutes > 0 {
		result = append(result, fmt.Sprintf("%d分钟", data.Minutes))
	}
	if data.Seconds > 0 {
		result = append(result, fmt.Sprintf("%d秒", data.Seconds))
	}

	return strings.Join(result, "")
}
