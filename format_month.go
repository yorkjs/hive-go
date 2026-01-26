package hive

import "time"

// FormatMonth 把时间戳格式化为 2020-10 格式
func FormatMonth(timestamp int64) string {
	return time.UnixMilli(timestamp).Format(MONTH_DEFAULT)
}
