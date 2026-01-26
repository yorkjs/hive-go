package hive

import (
	"fmt"
	"strings"
)

func FormatShelfLife(value int) string {
	var result []string

	data := NormalizeShelfLife(value)
	if data.Years > 0 {
		result = append(result, fmt.Sprintf("%d年", data.Years))
	}
	if data.Months > 0 {
		result = append(result, fmt.Sprintf("%d个月", data.Months))
	}
	if data.Days > 0 {
		result = append(result, fmt.Sprintf("%d天", data.Days))
	}
	if data.Hours > 0 {
		result = append(result, fmt.Sprintf("%d小时", data.Hours))
	}

	return strings.Join(result, "")

}
