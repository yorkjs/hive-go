package hive

import "time"

// FormatBirthday 把时间戳格式化为 10.01 格式
func FormatBirthday(value int64) string {
	return time.UnixMilli(value).Format(DATE_MONTH_DATE_DOT)
}
