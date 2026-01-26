package hive

import "time"

// FormatYear 把时间戳格式化为 2020 格式
func FormatYear(t time.Time) string {
	return t.Format(YEAR_DEFAULT)
}
