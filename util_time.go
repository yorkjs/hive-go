package hive

import (
	"time"
)

// StartOfHour 获取某个小时开始时间
func StartOfHour(t time.Time) time.Time {
	return time.Date(t.Year(), t.Month(), t.Day(), t.Hour(), 0, 0, 0, t.Location())
}

// StartOfPrevHour 获取前一个小时开始时间
func StartOfPrevHour(t time.Time) time.Time {
	return StartOfHour(t.Add(-1 * time.Hour))
}

// StartOfNextHour 获取下个小时开始时间
func StartOfNextHour(t time.Time) time.Time {
	return StartOfHour(t.Add(time.Hour))
}

// EndOfHour 获取某个小时结束时间
func EndOfHour(t time.Time) time.Time {
	return time.Date(t.Year(), t.Month(), t.Day(), t.Hour(), 59, 59, 999000000, t.Location())
}

// StartOfDay 获取某天的开始时间
func StartOfDay(t time.Time) time.Time {
	return time.Date(t.Year(), t.Month(), t.Day(), 0, 0, 0, 0, t.Location())
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
	// 23:59:59.999 -> 999ms = 999000000ns
	return time.Date(t.Year(), t.Month(), t.Day(), 23, 59, 59, 999000000, t.Location())
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
	return time.Date(t.Year(), t.Month(), 1, 0, 0, 0, 0, t.Location())
}

// StartOfPrevMonth 获取前一月的开始时间
func StartOfPrevMonth(t time.Time) time.Time {
	return time.Date(t.Year(), t.Month()-1, 1, 0, 0, 0, 0, t.Location())
}

// StartOfNextMonth 获取下一月的开始时间
func StartOfNextMonth(t time.Time) time.Time {
	return time.Date(t.Year(), t.Month()+1, 1, 0, 0, 0, 0, t.Location())
}

// EndOfMonth 获取某月的结束时间
func EndOfMonth(t time.Time) time.Time {
	return time.Date(t.Year(), t.Month()+1, 0, 23, 59, 59, 999000000, t.Location())
}

type ITimeRangeOptimizer struct {
	IsHour  func(hour time.Time)
	IsDay   func(day time.Time)
	IsWeek  func(week time.Time)
	IsMonth func(month time.Time)
	IsRange func(start, end time.Time)
}

// OptimizeTimeRange 优化时间范围
func OptimizeTimeRange(startTime, endTime time.Time, optimizer ITimeRangeOptimizer) {

	startHour := StartOfHour(startTime)
	endHour := EndOfHour(startTime)

	startDay := StartOfDay(startTime)
	endDay := EndOfDay(startTime)

	startWeek := StartOfWeek(startTime)
	endWeek := EndOfWeek(startTime)

	startMonth := StartOfMonth(startTime)
	endMonth := EndOfMonth(startTime)

	if startTime.Equal(startHour) && endTime.Equal(endHour) && optimizer.IsHour != nil {
		optimizer.IsHour(startTime)
	} else if startTime.Equal(startDay) && endTime.Equal(endDay) && optimizer.IsDay != nil {
		optimizer.IsDay(startTime)
	} else if startTime.Equal(startWeek) && endTime.Equal(endWeek) && optimizer.IsWeek != nil {
		optimizer.IsWeek(startTime)
	} else if startTime.Equal(startMonth) && endTime.Equal(endMonth) && optimizer.IsMonth != nil {
		optimizer.IsMonth(startTime)
	} else {
		optimizer.IsRange(startTime, endTime)
	}
}
