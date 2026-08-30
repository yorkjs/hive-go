package hive

import (
	"time"
)

// FormatDateTime 把时间戳格式化为 2020-10-01 10:00 格式
func FormatDateTime(t time.Time, format ...string) string {
	f := DATE_TIME_YEAR_MONTH_DATE_HOUR_MINUTE
	if len(format) > 0 {
		f = format[0]
	}
	return t.Format(f)
}

// FormatDateTimeRange 把时间戳格式化为 2020-10-01 00:00 至 2020-10-02 00:00 形式
func FormatDateTimeRange(start, end time.Time) string {
	// 日期相同，写一次日期即可
	if FormatDate(start) == FormatDate(end) {
		return FormatDateTime(start, DATE_TIME_YEAR_MONTH_DATE_HOUR_MINUTE) + " 至 " + FormatTime(end, TIME_HOUR_MINUTE)
	}
	return FormatDateTime(start, DATE_TIME_YEAR_MONTH_DATE_HOUR_MINUTE) + " 至 " + FormatDateTime(end, DATE_TIME_YEAR_MONTH_DATE_HOUR_MINUTE)
}

// FormatDateTimeShortly 把同年份的时间戳格式化为 10-01 10:00 格式，不同年份的时间戳格式化成 2020-10-01 10:00 格式
func FormatDateTimeShortly(t time.Time) string {
	t1 := t
	t2 := time.Now()

	if t1.Year() == t2.Year() {
		return t1.Format(DATE_TIME_MONTH_DATE_HOUR_MINUTE)
	}
	return t1.Format(DATE_TIME_YEAR_MONTH_DATE_HOUR_MINUTE)
}
