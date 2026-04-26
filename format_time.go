package hive

import "time"

// FormatTime 把时间戳格式化为 12:00:00 格式
func FormatTime(t time.Time, format ...string) string {
	f := TIME_DEFAULT
	if len(format) > 0 {
		f = format[0]
	}
	return t.Format(f)
}
