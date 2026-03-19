package hive

import "time"

// FormatMonth 把时间戳格式化为 2020-10 格式
func FormatMonth(t time.Time, format ...string) string {
	f := MONTH_DEFAULT
	if len(format) > 0 {
		f = format[0]
	}
	return t.Format(f)
}
