package hive

import (
	"time"
)

// FormatDate 把时间戳格式化为 2020-10-01 格式
//
// timestamp 毫秒时间戳
func FormatDate(timestamp int64, format ...string) string {
	f := DATE_YEAR_MONTH_DATE
	if len(format) > 0 {
		f = format[0]
	}
	return time.UnixMilli(timestamp).Format(f)
}

// FormatDateRange 把时间戳格式化为 2020-10-01 至 2020-10-02 形式
func FormatDateRange(startTimestamp, endTimestamp int64) string {
	return FormatDate(startTimestamp) + " 至 " + FormatDate(endTimestamp)
}

// FormatDateShortly 把同年份的时间戳格式化为 10-01 格式，不同年份的时间戳格式化成 2020-10-01 格式
func FormatDateShortly(timestamp int64) string {
	t1 := time.UnixMilli(timestamp)
	t2 := time.Now()

	if t1.Year() == t2.Year() {
		return t1.Format(DATE_MONTH_DATE)
	}
	return t1.Format(DATE_YEAR_MONTH_DATE)
}
