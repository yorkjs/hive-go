package hive

import (
	"time"
)

// StartOfDay 获取某天的开始时间
//
// timestamp 毫秒时间戳
func StartOfDay(timestamp int64) int64 {
	t := time.UnixMilli(timestamp)
	y, m, d := t.Date()
	start := time.Date(y, m, d, 0, 0, 0, 0, t.Location())
	return start.UnixMilli()
}

// StartOfPrevDay 获取前一天的开始时间
func StartOfPrevDay(timestamp int64) int64 {
	return StartOfDay(timestamp - MS_DAY)
}

// StartOfNextDay 获取后一天的开始时间
func StartOfNextDay(timestamp int64) int64 {
	return StartOfDay(timestamp + MS_DAY)
}

// EndOfDay 获取某天的结束时间
func EndOfDay(timestamp int64) int64 {
	t := time.UnixMilli(timestamp)
	y, m, d := t.Date()
	// 23:59:59.999 -> 999ms = 999000000ns
	end := time.Date(y, m, d, 23, 59, 59, 999000000, t.Location())
	return end.UnixMilli()
}

// StartOfWeek 获取某周的开始时间
func StartOfWeek(timestamp int64) int64 {
	t := time.UnixMilli(timestamp)
	day := int(t.Weekday())

	// Go Weekday: Sunday=0, Monday=1...
	offset := day
	// Subtract offset days
	start := t.AddDate(0, 0, -offset)
	y, m, d := start.Date()
	return time.Date(y, m, d, 0, 0, 0, 0, t.Location()).UnixMilli()
}

// StartOfPrevWeek 获取前一周的开始时间
func StartOfPrevWeek(timestamp int64) int64 {
	return StartOfWeek(timestamp - MS_WEEK)
}

// StartOfNextWeek 获取后一周的开始时间
func StartOfNextWeek(timestamp int64) int64 {
	return StartOfWeek(timestamp + MS_WEEK)
}

// EndOfWeek 获取某周的结束时间
func EndOfWeek(timestamp int64) int64 {
	t := time.UnixMilli(timestamp)
	day := int(t.Weekday())

	offset := 6 - day
	end := t.AddDate(0, 0, offset)
	y, m, d := end.Date()
	return time.Date(y, m, d, 23, 59, 59, 999000000, t.Location()).UnixMilli()
}

// StartOfMonth 获取某月的开始时间
func StartOfMonth(timestamp int64) int64 {
	t := time.UnixMilli(timestamp)
	y, m, _ := t.Date()
	start := time.Date(y, m, 1, 0, 0, 0, 0, t.Location())
	return start.UnixMilli()
}

// StartOfPrevMonth 获取前一月的开始时间
func StartOfPrevMonth(timestamp int64) int64 {
	t := time.UnixMilli(timestamp)
	y, m, _ := t.Date()
	// Last month
	start := time.Date(y, m-1, 1, 0, 0, 0, 0, t.Location())
	return start.UnixMilli()
}

// StartOfNextMonth 获取下一月的开始时间
func StartOfNextMonth(timestamp int64) int64 {
	t := time.UnixMilli(timestamp)
	y, m, _ := t.Date()
	// Next month
	start := time.Date(y, m+1, 1, 0, 0, 0, 0, t.Location())
	return start.UnixMilli()
}

// EndOfMonth 获取某月的结束时间
func EndOfMonth(timestamp int64) int64 {
	t := time.UnixMilli(timestamp)
	y, m, _ := t.Date()
	// Next month 1st day minus 1ms? Or use Date(y, m+1, 0) logic
	// Go: Date(y, m+1, 0, ...) gives the day before 1st of m+1. i.e. last day of m.
	end := time.Date(y, m+1, 0, 23, 59, 59, 999000000, t.Location())
	return end.UnixMilli()
}

type ITimeRangeOptimizer struct {
	IsDay   func(day int64)
	IsWeek  func(week int64)
	IsMonth func(month int64)
	IsRange func(start, end int64)
}

// OptimizeTimeRange 优化时间范围
func OptimizeTimeRange(startTimestamp, endTimestamp int64, optimizer ITimeRangeOptimizer) {
	startDay := StartOfDay(startTimestamp)
	endDay := EndOfDay(startTimestamp)

	startWeek := StartOfWeek(startTimestamp)
	endWeek := EndOfWeek(startTimestamp)

	startMonth := StartOfMonth(startTimestamp)
	endMonth := EndOfMonth(startTimestamp)

	if startTimestamp == startDay && endTimestamp == endDay && optimizer.IsDay != nil {
		optimizer.IsDay(startTimestamp)
	} else if startTimestamp == startWeek && endTimestamp == endWeek && optimizer.IsWeek != nil {
		optimizer.IsWeek(startTimestamp)
	} else if startTimestamp == startMonth && endTimestamp == endMonth && optimizer.IsMonth != nil {
		optimizer.IsMonth(startTimestamp)
	} else {
		optimizer.IsRange(startTimestamp, endTimestamp)
	}
}
