package hive

import (
	"time"
)

func ParseTime(str string, format string) (time.Time, error) {
	return time.ParseInLocation(format, str, time.Local)
}

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

// StartOfYear 获取某年的开始时间
func StartOfYear(t time.Time) time.Time {
	return time.Date(t.Year(), 1, 1, 0, 0, 0, 0, t.Location())
}

// StartOfPrevYear 获取前一年的开始时间
func StartOfPrevYear(t time.Time) time.Time {
	return time.Date(t.Year()-1, 1, 1, 0, 0, 0, 0, t.Location())
}

// StartOfNextYear 获取下一年的开始时间
func StartOfNextYear(t time.Time) time.Time {
	return time.Date(t.Year()+1, 1, 1, 0, 0, 0, 0, t.Location())
}

// EndOfYear 获取某年的结束时间
func EndOfYear(t time.Time) time.Time {
	return time.Date(t.Year(), 12, 31, 23, 59, 59, 999000000, t.Location())
}

// SameOfPrevDay 获取昨天的同时刻
func SameOfPrevDay(t time.Time) time.Time {
	return t.AddDate(0, 0, -1)
}

// SameOfPrevWeek 获取前一周的同时刻
func SameOfPrevWeek(t time.Time) time.Time {
	return t.AddDate(0, 0, -7)
}

// SameOfPrevMonth 获取上个月同时刻
func SameOfPrevMonth(t time.Time) time.Time {

	day := t.Day()

	prevMonthStart := StartOfPrevMonth(t)
	prevMonthEnd := EndOfMonth(prevMonthStart)

	// 如果当前日期大于上个月的最后一天，则使用上个月的最后一天
	if day > prevMonthEnd.Day() {
		return prevMonthEnd
	}

	// 构建上个月同时刻的时间
	return time.Date(
		prevMonthStart.Year(),
		prevMonthStart.Month(),
		day,
		t.Hour(),
		t.Minute(),
		t.Second(),
		t.Nanosecond(),
		t.Location(),
	)
}

// SameOfPrevYear 获取去年的同时刻
func SameOfPrevYear(t time.Time) time.Time {

	// 获取当前日期
	month := t.Month()
	day := t.Day()

	// 计算去年的同月同日
	prevYearMonthStart := time.Date(
		t.Year()-1,
		month,
		1,
		0, 0, 0, 0,
		t.Location(),
	)

	prevYearMonthEnd := EndOfMonth(prevYearMonthStart)

	// 如果当前日期大于去年该月的最后一天，则使用该月的最后一天
	if day > prevYearMonthEnd.Day() {
		return prevYearMonthEnd
	}

	// 构建去年同时刻的时间
	return time.Date(
		t.Year()-1,
		month,
		day,
		t.Hour(),
		t.Minute(),
		t.Second(),
		t.Nanosecond(),
		t.Location(),
	)
}

// MatchMinuteSegments 判断目标时间是否落在给定的分钟段区间内（左闭右开）
func MatchMinuteSegments(segments []int, target time.Time) bool {
	length := len(segments)
	if length == 0 || length%2 != 0 {
		return false
	}

	firstValue := TimeToMinuteSegment(target) // 当日分钟数
	secondValue := firstValue + 24*60         // 次日分钟数（偏移一天）

	for i := 0; i < length; i += 2 {
		startTime := segments[i]
		endTime := segments[i+1]
		if firstValue >= startTime && firstValue < endTime {
			return true
		}
		if secondValue >= startTime && secondValue < endTime {
			return true
		}
	}
	return false
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
