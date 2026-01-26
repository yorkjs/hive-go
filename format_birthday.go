package hive

import "time"

// FormatBirthday 把时间戳格式化为 10.01 格式
func FormatBirthday(timestamp int64) string {
	return time.UnixMilli(timestamp).Format(DATE_MONTH_DATE_DOT)
}
