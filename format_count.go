package hive

import "fmt"

// FormatCount 格式化数量
func FormatCount(value float64, unit ...string) string {
	u := ""
	if len(unit) > 0 {
		u = unit[0]
	}
	return FormatNumberWithComma(value) + u
}

// FormatCountShortly 格式化数量（以尽可能短的方式显示数量）
func FormatCountShortly(value float64, unit ...string) string {
	u := ""
	if len(unit) > 0 {
		u = unit[0]
	}
	return ShortNumber(
		value,
		func(v float64) string {
			return fmt.Sprintf("%v", v)
		},
	) + u
}
