package hive

import "time"

// FormatYear 把时间戳格式化为 2020 格式
func FormatYear(t time.Time, format ...string) string {
	f := YEAR_DEFAULT
	if len(format) > 0 {
		f = format[0]
	}
	return t.Format(f)
}
