package hive

import "time"

// FormatBirthday 把时间戳格式化为 10.01 格式
func FormatBirthday(t time.Time, format ...string) string {
	f := DATE_MONTH_DATE_DOT
	if len(format) > 0 {
		f = format[0]
	}
	return t.Format(f)
}
