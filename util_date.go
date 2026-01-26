package hive

import (
	"time"
)

// StartOfDay 获取某天的开始时间
//
// t 时间
func StartOfDay(t time.Time) time.Time {
	y, m, d := t.Date()
	start := time.Date(y, m, d, 0, 0, 0, 0, t.Location())
	return start
}

// StartOfPrevDay 获取前一天的开始时间
func StartOfPrevDay(t time.Time) time.Time {
	return StartOfDay(t.AddDate(0, 0, -1))
}

// StartOfNextDay 获取后一天的开始时间
func StartOfNextDay(t time.Time) time.Time {
	return StartOfDay(t.AddDate(0, 0, 1))
}

// EndOfDay 获取某天的结束时间
func EndOfDay(t time.Time) time.Time {
	y, m, d := t.Date()
	// 23:59:59.999 -> 999ms = 999000000ns
	end := time.Date(y, m, d, 23, 59, 59, 999000000, t.Location())
	return end
}

// StartOfWeek 获取某周的开始时间
func StartOfWeek(t time.Time) time.Time {
	day := int(t.Weekday())

	// Go Weekday: Sunday=0, Monday=1...
	offset := day
	// Subtract offset days
	start := t.AddDate(0, 0, -offset)
	y, m, d := start.Date()
	return time.Date(y, m, d, 0, 0, 0, 0, t.Location())
}

// StartOfPrevWeek 获取前一周的开始时间
func StartOfPrevWeek(t time.Time) time.Time {
	return StartOfWeek(t.AddDate(0, 0, -7))
}

// StartOfNextWeek 获取后一周的开始时间
func StartOfNextWeek(t time.Time) time.Time {
	return StartOfWeek(t.AddDate(0, 0, 7))
}

// EndOfWeek 获取某周的结束时间
func EndOfWeek(t time.Time) time.Time {
	day := int(t.Weekday())

	offset := 6 - day
	end := t.AddDate(0, 0, offset)
	y, m, d := end.Date()
	return time.Date(y, m, d, 23, 59, 59, 999000000, t.Location())
}

// StartOfMonth 获取某月的开始时间
func StartOfMonth(t time.Time) time.Time {
	y, m, _ := t.Date()
	start := time.Date(y, m, 1, 0, 0, 0, 0, t.Location())
	return start
}

// StartOfPrevMonth 获取前一月的开始时间
func StartOfPrevMonth(t time.Time) time.Time {
	y, m, _ := t.Date()
	// Last month
	start := time.Date(y, m-1, 1, 0, 0, 0, 0, t.Location())
	return start
}

// StartOfNextMonth 获取下一月的开始时间
func StartOfNextMonth(t time.Time) time.Time {
	y, m, _ := t.Date()
	// Next month
	start := time.Date(y, m+1, 1, 0, 0, 0, 0, t.Location())
	return start
}

// EndOfMonth 获取某月的结束时间
func EndOfMonth(t time.Time) time.Time {
	y, m, _ := t.Date()
	// Next month 1st day minus 1ms? Or use Date(y, m+1, 0) logic
	// Go: Date(y, m+1, 0, ...) gives the day before 1st of m+1. i.e. last day of m.
	end := time.Date(y, m+1, 0, 23, 59, 59, 999000000, t.Location())
	return end
}

type ITimeRangeOptimizer struct {
	IsDay   func(day time.Time)
	IsWeek  func(week time.Time)
	IsMonth func(month time.Time)
	IsRange func(start, end time.Time)
}

// OptimizeTimeRange 优化时间范围
func OptimizeTimeRange(startTime, endTime time.Time, optimizer ITimeRangeOptimizer) {
	startDay := StartOfDay(startTime)
	endDay := EndOfDay(startTime)

	startWeek := StartOfWeek(startTime)
	endWeek := EndOfWeek(startTime)

	startMonth := StartOfMonth(startTime)
	endMonth := EndOfMonth(startTime)

	if startTime.Equal(startDay) && endTime.Equal(endDay) && optimizer.IsDay != nil {
		optimizer.IsDay(startTime)
	} else if startTime.Equal(startWeek) && endTime.Equal(endWeek) && optimizer.IsWeek != nil {
		optimizer.IsWeek(startTime)
	} else if startTime.Equal(startMonth) && endTime.Equal(endMonth) && optimizer.IsMonth != nil {
		optimizer.IsMonth(startTime)
	} else {
		optimizer.IsRange(startTime, endTime)
	}
}
