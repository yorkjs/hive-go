package hive

import "fmt"

// FormatCount 格式化数量
func FormatCount[T IntegerType](value T, unit ...string) string {
	u := ""
	if len(unit) > 0 {
		u = unit[0]
	}
	return FormatNumberWithComma(float64(value)) + u
}

// FormatCountShortly 格式化数量（以尽可能短的方式显示数量）
func FormatCountShortly[T IntegerType](value T, unit ...string) string {
	u := ""
	if len(unit) > 0 {
		u = unit[0]
	}
	return shortNumber(
		value,
		func(v T) string {
			return fmt.Sprintf("%v", v)
		},
	) + u
}
