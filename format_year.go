package hive

import "time"

// FormatYear 把时间戳格式化为 2020 格式
func FormatYear(timestamp int64) string {
	return time.UnixMilli(timestamp).Format(YEAR_DEFAULT)
}
