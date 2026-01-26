package hive

import "time"

// FormatMonth 把时间戳格式化为 2020-10 格式
func FormatMonth(t time.Time) string {
	return t.Format(MONTH_DEFAULT)
}
