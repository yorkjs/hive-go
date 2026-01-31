package hive

import (
	"testing"
	"time"
)

func TestTimeHour(t *testing.T) {
	loc := time.Local

	// 2020-10-10 10:01:01
	date := time.Date(2020, 10, 10, 10, 1, 1, 0, loc)

	if got := FormatDateTime(StartOfHour(date), DATE_TIME_YEAR_MONTH_DATE_HOUR_MINUTE_SECOND); got != "2020-10-10 10:00:00" {
		t.Errorf("StartOfHour(2020-10-10) = %v; want 2020-10-10 10:00:00", got)
	}
	if got := FormatDateTime(StartOfPrevHour(date), DATE_TIME_YEAR_MONTH_DATE_HOUR_MINUTE_SECOND); got != "2020-10-10 09:00:00" {
		t.Errorf("StartOfPrevHour(2020-10-10) = %v; want 2020-10-10 09:00:00", got)
	}
	if got := FormatDateTime(StartOfNextHour(date), DATE_TIME_YEAR_MONTH_DATE_HOUR_MINUTE_SECOND); got != "2020-10-10 11:00:00" {
		t.Errorf("StartOfNextHour(2020-10-10) = %v; want 2020-10-10 11:00:00", got)
	}
	if got := FormatDateTime(EndOfHour(date), DATE_TIME_YEAR_MONTH_DATE_HOUR_MINUTE_SECOND); got != "2020-10-10 10:59:59" {
		t.Errorf("EndOfHour(2020-10-10) = %v; want 2020-10-10 10:59:59", got)
	}

}

func TestTimeDay(t *testing.T) {
	loc := time.Local

	// 2020-10-10 10:01:01
	date := time.Date(2020, 10, 10, 10, 1, 1, 0, loc)

	if got := FormatDateTime(StartOfDay(date), DATE_TIME_YEAR_MONTH_DATE_HOUR_MINUTE_SECOND); got != "2020-10-10 00:00:00" {
		t.Errorf("StartOfDay(2020-10-10) = %v; want 2020-10-10 00:00:00", got)
	}
	if got := FormatDateTime(StartOfPrevDay(date), DATE_TIME_YEAR_MONTH_DATE_HOUR_MINUTE_SECOND); got != "2020-10-09 00:00:00" {
		t.Errorf("StartOfPrevDay(2020-10-10) = %v; want 2020-10-09 00:00:00", got)
	}
	if got := FormatDateTime(StartOfNextDay(date), DATE_TIME_YEAR_MONTH_DATE_HOUR_MINUTE_SECOND); got != "2020-10-11 00:00:00" {
		t.Errorf("StartOfNextDay(2020-10-10) = %v; want 2020-10-11 00:00:00", got)
	}
	if got := FormatDateTime(EndOfDay(date), DATE_TIME_YEAR_MONTH_DATE_HOUR_MINUTE_SECOND); got != "2020-10-10 23:59:59" {
		t.Errorf("EndOfDay(2020-10-10) = %v; want 2020-10-10 23:59:59", got)
	}

	// 2020-04-01 10:01:01
	date = time.Date(2020, 4, 1, 10, 1, 1, 0, loc)
	if got := FormatDateTime(StartOfPrevDay(date), DATE_TIME_YEAR_MONTH_DATE_HOUR_MINUTE_SECOND); got != "2020-03-31 00:00:00" {
		t.Errorf("StartOfPrevDay(2020-04-01) = %v; want 2020-03-31 00:00:00", got)
	}
	if got := FormatDateTime(StartOfNextDay(date), DATE_TIME_YEAR_MONTH_DATE_HOUR_MINUTE_SECOND); got != "2020-04-02 00:00:00" {
		t.Errorf("StartOfNextDay(2020-04-01) = %v; want 2020-04-02 00:00:00", got)
	}

	// 2020-01-01 10:01:01
	date = time.Date(2020, 1, 1, 10, 1, 1, 0, loc)
	if got := FormatDateTime(StartOfPrevDay(date), DATE_TIME_YEAR_MONTH_DATE_HOUR_MINUTE_SECOND); got != "2019-12-31 00:00:00" {
		t.Errorf("StartOfPrevDay(2020-01-01) = %v; want 2019-12-31 00:00:00", got)
	}
	if got := FormatDateTime(StartOfNextDay(date), DATE_TIME_YEAR_MONTH_DATE_HOUR_MINUTE_SECOND); got != "2020-01-02 00:00:00" {
		t.Errorf("StartOfNextDay(2020-01-01) = %v; want 2020-01-02 00:00:00", got)
	}

	// 2024-03-01 10:01:01 (Leap year)
	date = time.Date(2024, 3, 1, 10, 1, 1, 0, loc)
	if got := FormatDateTime(StartOfPrevDay(date), DATE_TIME_YEAR_MONTH_DATE_HOUR_MINUTE_SECOND); got != "2024-02-29 00:00:00" {
		t.Errorf("StartOfPrevDay(2024-03-01) = %v; want 2024-02-29 00:00:00", got)
	}
	if got := FormatDateTime(StartOfNextDay(date), DATE_TIME_YEAR_MONTH_DATE_HOUR_MINUTE_SECOND); got != "2024-03-02 00:00:00" {
		t.Errorf("StartOfNextDay(2024-03-01) = %v; want 2024-03-02 00:00:00", got)
	}

	// 2025-03-01 10:01:01 (Non-leap year)
	date = time.Date(2025, 3, 1, 10, 1, 1, 0, loc)
	if got := FormatDateTime(StartOfPrevDay(date), DATE_TIME_YEAR_MONTH_DATE_HOUR_MINUTE_SECOND); got != "2025-02-28 00:00:00" {
		t.Errorf("StartOfPrevDay(2025-03-01) = %v; want 2025-02-28 00:00:00", got)
	}
	if got := FormatDateTime(StartOfNextDay(date), DATE_TIME_YEAR_MONTH_DATE_HOUR_MINUTE_SECOND); got != "2025-03-02 00:00:00" {
		t.Errorf("StartOfNextDay(2025-03-01) = %v; want 2025-03-02 00:00:00", got)
	}
}

func TestTimeWeek(t *testing.T) {
	loc := time.Local

	// 2025-07-27 10:01:01 (Sunday)
	date := time.Date(2025, 7, 27, 10, 1, 1, 0, loc)

	if got := FormatDateTime(StartOfWeek(date), DATE_TIME_YEAR_MONTH_DATE_HOUR_MINUTE_SECOND); got != "2025-07-27 00:00:00" {
		t.Errorf("StartOfWeek(2025-07-27) = %v; want 2025-07-27 00:00:00", got)
	}
	if got := FormatDateTime(StartOfPrevWeek(date), DATE_TIME_YEAR_MONTH_DATE_HOUR_MINUTE_SECOND); got != "2025-07-20 00:00:00" {
		t.Errorf("StartOfPrevWeek(2025-07-27) = %v; want 2025-07-20 00:00:00", got)
	}
	if got := FormatDateTime(StartOfNextWeek(date), DATE_TIME_YEAR_MONTH_DATE_HOUR_MINUTE_SECOND); got != "2025-08-03 00:00:00" {
		t.Errorf("StartOfNextWeek(2025-07-27) = %v; want 2025-08-03 00:00:00", got)
	}
	if got := FormatDateTime(EndOfWeek(date), DATE_TIME_YEAR_MONTH_DATE_HOUR_MINUTE_SECOND); got != "2025-08-02 23:59:59" {
		t.Errorf("EndOfWeek(2025-07-27) = %v; want 2025-08-02 23:59:59", got)
	}

	// 2025-07-29 10:01:01 (Tuesday)
	date = time.Date(2025, 7, 29, 10, 1, 1, 0, loc)
	if got := FormatDateTime(StartOfWeek(date), DATE_TIME_YEAR_MONTH_DATE_HOUR_MINUTE_SECOND); got != "2025-07-27 00:00:00" {
		t.Errorf("StartOfWeek(2025-07-29) = %v; want 2025-07-27 00:00:00", got)
	}
	if got := FormatDateTime(EndOfWeek(date), DATE_TIME_YEAR_MONTH_DATE_HOUR_MINUTE_SECOND); got != "2025-08-02 23:59:59" {
		t.Errorf("EndOfWeek(2025-07-29) = %v; want 2025-08-02 23:59:59", got)
	}

	// 2025-08-01 10:01:01 (Friday)
	date = time.Date(2025, 8, 1, 10, 1, 1, 0, loc)
	if got := FormatDateTime(StartOfWeek(date), DATE_TIME_YEAR_MONTH_DATE_HOUR_MINUTE_SECOND); got != "2025-07-27 00:00:00" {
		t.Errorf("StartOfWeek(2025-08-01) = %v; want 2025-07-27 00:00:00", got)
	}
	if got := FormatDateTime(EndOfWeek(date), DATE_TIME_YEAR_MONTH_DATE_HOUR_MINUTE_SECOND); got != "2025-08-02 23:59:59" {
		t.Errorf("EndOfWeek(2025-08-01) = %v; want 2025-08-02 23:59:59", got)
	}

	// 2025-08-02 10:01:01 (Saturday)
	date = time.Date(2025, 8, 2, 10, 1, 1, 0, loc)
	if got := FormatDateTime(StartOfWeek(date), DATE_TIME_YEAR_MONTH_DATE_HOUR_MINUTE_SECOND); got != "2025-07-27 00:00:00" {
		t.Errorf("StartOfWeek(2025-08-02) = %v; want 2025-07-27 00:00:00", got)
	}
	if got := FormatDateTime(EndOfWeek(date), DATE_TIME_YEAR_MONTH_DATE_HOUR_MINUTE_SECOND); got != "2025-08-02 23:59:59" {
		t.Errorf("EndOfWeek(2025-08-02) = %v; want 2025-08-02 23:59:59", got)
	}
}

func TestTimeMonth(t *testing.T) {
	loc := time.Local

	// 2025-02-10 10:01:01
	date := time.Date(2025, 2, 10, 10, 1, 1, 0, loc)

	if got := FormatDateTime(StartOfMonth(date), DATE_TIME_YEAR_MONTH_DATE_HOUR_MINUTE_SECOND); got != "2025-02-01 00:00:00" {
		t.Errorf("StartOfMonth(2025-02-10) = %v; want 2025-02-01 00:00:00", got)
	}
	if got := FormatDateTime(StartOfPrevMonth(date), DATE_TIME_YEAR_MONTH_DATE_HOUR_MINUTE_SECOND); got != "2025-01-01 00:00:00" {
		t.Errorf("StartOfPrevMonth(2025-02-10) = %v; want 2025-01-01 00:00:00", got)
	}
	if got := FormatDateTime(StartOfNextMonth(date), DATE_TIME_YEAR_MONTH_DATE_HOUR_MINUTE_SECOND); got != "2025-03-01 00:00:00" {
		t.Errorf("StartOfNextMonth(2025-02-10) = %v; want 2025-03-01 00:00:00", got)
	}
	if got := FormatDateTime(EndOfMonth(date), DATE_TIME_YEAR_MONTH_DATE_HOUR_MINUTE_SECOND); got != "2025-02-28 23:59:59" {
		t.Errorf("EndOfMonth(2025-02-10) = %v; want 2025-02-28 23:59:59", got)
	}

	// 2025-06-29 10:01:01
	date = time.Date(2025, 6, 29, 10, 1, 1, 0, loc)
	if got := FormatDateTime(StartOfMonth(date), DATE_TIME_YEAR_MONTH_DATE_HOUR_MINUTE_SECOND); got != "2025-06-01 00:00:00" {
		t.Errorf("StartOfMonth(2025-06-29) = %v; want 2025-06-01 00:00:00", got)
	}
	if got := FormatDateTime(StartOfPrevMonth(date), DATE_TIME_YEAR_MONTH_DATE_HOUR_MINUTE_SECOND); got != "2025-05-01 00:00:00" {
		t.Errorf("StartOfPrevMonth(2025-06-29) = %v; want 2025-05-01 00:00:00", got)
	}
	if got := FormatDateTime(StartOfNextMonth(date), DATE_TIME_YEAR_MONTH_DATE_HOUR_MINUTE_SECOND); got != "2025-07-01 00:00:00" {
		t.Errorf("StartOfNextMonth(2025-06-29) = %v; want 2025-07-01 00:00:00", got)
	}
	if got := FormatDateTime(EndOfMonth(date), DATE_TIME_YEAR_MONTH_DATE_HOUR_MINUTE_SECOND); got != "2025-06-30 23:59:59" {
		t.Errorf("EndOfMonth(2025-06-29) = %v; want 2025-06-30 23:59:59", got)
	}

	// 2025-07-29 10:01:01
	date = time.Date(2025, 7, 29, 10, 1, 1, 0, loc)
	if got := FormatDateTime(StartOfMonth(date), DATE_TIME_YEAR_MONTH_DATE_HOUR_MINUTE_SECOND); got != "2025-07-01 00:00:00" {
		t.Errorf("StartOfMonth(2025-07-29) = %v; want 2025-07-01 00:00:00", got)
	}
	if got := FormatDateTime(StartOfPrevMonth(date), DATE_TIME_YEAR_MONTH_DATE_HOUR_MINUTE_SECOND); got != "2025-06-01 00:00:00" {
		t.Errorf("StartOfPrevMonth(2025-07-29) = %v; want 2025-06-01 00:00:00", got)
	}
	if got := FormatDateTime(StartOfNextMonth(date), DATE_TIME_YEAR_MONTH_DATE_HOUR_MINUTE_SECOND); got != "2025-08-01 00:00:00" {
		t.Errorf("StartOfNextMonth(2025-07-29) = %v; want 2025-08-01 00:00:00", got)
	}
	if got := FormatDateTime(EndOfMonth(date), DATE_TIME_YEAR_MONTH_DATE_HOUR_MINUTE_SECOND); got != "2025-07-31 23:59:59" {
		t.Errorf("EndOfMonth(2025-07-29) = %v; want 2025-07-31 23:59:59", got)
	}

	// 2025-12-19 10:01:01
	date = time.Date(2025, 12, 19, 10, 1, 1, 0, loc)
	if got := FormatDateTime(StartOfMonth(date), DATE_TIME_YEAR_MONTH_DATE_HOUR_MINUTE_SECOND); got != "2025-12-01 00:00:00" {
		t.Errorf("StartOfMonth(2025-12-19) = %v; want 2025-12-01 00:00:00", got)
	}
	if got := FormatDateTime(StartOfPrevMonth(date), DATE_TIME_YEAR_MONTH_DATE_HOUR_MINUTE_SECOND); got != "2025-11-01 00:00:00" {
		t.Errorf("StartOfPrevMonth(2025-12-19) = %v; want 2025-11-01 00:00:00", got)
	}
	if got := FormatDateTime(StartOfNextMonth(date), DATE_TIME_YEAR_MONTH_DATE_HOUR_MINUTE_SECOND); got != "2026-01-01 00:00:00" {
		t.Errorf("StartOfNextMonth(2025-12-19) = %v; want 2026-01-01 00:00:00", got)
	}
	if got := FormatDateTime(EndOfMonth(date), DATE_TIME_YEAR_MONTH_DATE_HOUR_MINUTE_SECOND); got != "2025-12-31 23:59:59" {
		t.Errorf("EndOfMonth(2025-12-19) = %v; want 2025-12-31 23:59:59", got)
	}

	// 2025-04-19 10:01:01
	date = time.Date(2025, 4, 19, 10, 1, 1, 0, loc)
	if got := FormatDateTime(StartOfMonth(date), DATE_TIME_YEAR_MONTH_DATE_HOUR_MINUTE_SECOND); got != "2025-04-01 00:00:00" {
		t.Errorf("StartOfMonth(2025-04-19) = %v; want 2025-04-01 00:00:00", got)
	}
	if got := FormatDateTime(StartOfPrevMonth(date), DATE_TIME_YEAR_MONTH_DATE_HOUR_MINUTE_SECOND); got != "2025-03-01 00:00:00" {
		t.Errorf("StartOfPrevMonth(2025-04-19) = %v; want 2025-03-01 00:00:00", got)
	}
	if got := FormatDateTime(StartOfNextMonth(date), DATE_TIME_YEAR_MONTH_DATE_HOUR_MINUTE_SECOND); got != "2025-05-01 00:00:00" {
		t.Errorf("StartOfNextMonth(2025-04-19) = %v; want 2025-05-01 00:00:00", got)
	}
	if got := FormatDateTime(EndOfMonth(date), DATE_TIME_YEAR_MONTH_DATE_HOUR_MINUTE_SECOND); got != "2025-04-30 23:59:59" {
		t.Errorf("EndOfMonth(2025-04-19) = %v; want 2025-04-30 23:59:59", got)
	}

	// 2025-01-19 10:01:01
	date = time.Date(2025, 1, 19, 10, 1, 1, 0, loc)
	if got := FormatDateTime(StartOfMonth(date), DATE_TIME_YEAR_MONTH_DATE_HOUR_MINUTE_SECOND); got != "2025-01-01 00:00:00" {
		t.Errorf("StartOfMonth(2025-01-19) = %v; want 2025-01-01 00:00:00", got)
	}
	if got := FormatDateTime(StartOfPrevMonth(date), DATE_TIME_YEAR_MONTH_DATE_HOUR_MINUTE_SECOND); got != "2024-12-01 00:00:00" {
		t.Errorf("StartOfPrevMonth(2025-01-19) = %v; want 2024-12-01 00:00:00", got)
	}
	if got := FormatDateTime(StartOfNextMonth(date), DATE_TIME_YEAR_MONTH_DATE_HOUR_MINUTE_SECOND); got != "2025-02-01 00:00:00" {
		t.Errorf("StartOfNextMonth(2025-01-19) = %v; want 2025-02-01 00:00:00", got)
	}
	if got := FormatDateTime(EndOfMonth(date), DATE_TIME_YEAR_MONTH_DATE_HOUR_MINUTE_SECOND); got != "2025-01-31 23:59:59" {
		t.Errorf("EndOfMonth(2025-01-19) = %v; want 2025-01-31 23:59:59", got)
	}
}

func TestTimeRangeOptimize(t *testing.T) {
	loc := time.Local

	// 2025-02-10 10:01:01
	date := time.Date(2025, 2, 10, 10, 1, 1, 0, loc)

	// isHour
	startTime := StartOfDay(date)
	endTime := EndOfDay(date)
	isHour := false
	isDay := false
	isWeek := false
	isMonth := false
	isRange := false

	OptimizeTimeRange(startTime, endTime, ITimeRangeOptimizer{
		IsHour: func(hour time.Time) {
			isHour = true
			if hour != startTime {
				t.Errorf("IsHour: expected %v, got %v", startTime, hour)
			}
		},
		IsDay: func(day time.Time) {
			isDay = true
		},
		IsWeek: func(week time.Time) {
			isWeek = true
		},
		IsMonth: func(month time.Time) {
			isMonth = true
		},
		IsRange: func(start, end time.Time) {
			isRange = true
		},
	})

	if isHour {
		t.Errorf("expected isHour to be false, got true")
	}
	if !isDay {
		t.Errorf("expected isDay to be true, got false")
	}
	if isWeek {
		t.Errorf("expected isWeek to be false, got true")
	}
	if isMonth {
		t.Errorf("expected isMonth to be false, got true")
	}
	if isRange {
		t.Errorf("expected isRange to be false, got true")
	}

	// isHour 但是不传 isHour 函数
	startTime = StartOfDay(date)
	endTime = EndOfDay(date)
	isHour = false
	isDay = false
	isWeek = false
	isMonth = false
	isRange = false

	OptimizeTimeRange(startTime, endTime, ITimeRangeOptimizer{
		IsDay: func(day time.Time) {
			isDay = true
			if day != startTime {
				t.Errorf("IsDay: expected %v, got %v", startTime, day)
			}
		},
		IsWeek: func(week time.Time) {
			isWeek = true
		},
		IsMonth: func(month time.Time) {
			isMonth = true
		},
		IsRange: func(start, end time.Time) {
			isRange = true
		},
	})

	if isHour {
		t.Errorf("expected isHour to be false, got true")
	}
	if !isDay {
		t.Errorf("expected isDay to be true, got false")
	}
	if isWeek {
		t.Errorf("expected isWeek to be false, got true")
	}
	if isMonth {
		t.Errorf("expected isMonth to be false, got true")
	}
	if isRange {
		t.Errorf("expected isRange to be false, got true")
	}

	// isDay 但是不传 isDay 函数
	startTime = StartOfDay(date)
	endTime = EndOfDay(date)
	isDay = false
	isWeek = false
	isMonth = false
	isRange = false

	OptimizeTimeRange(startTime, endTime, ITimeRangeOptimizer{
		IsWeek: func(week time.Time) {
			isWeek = true
		},
		IsMonth: func(month time.Time) {
			isMonth = true
		},
		IsRange: func(start, end time.Time) {
			isRange = true
			if start != startTime {
				t.Errorf("IsRange start: expected %v, got %v", startTime, start)
			}
			if end != endTime {
				t.Errorf("IsRange end: expected %v, got %v", endTime, end)
			}
		},
	})

	if isDay {
		t.Errorf("expected isDay to be false, got true")
	}
	if isWeek {
		t.Errorf("expected isWeek to be false, got true")
	}
	if isMonth {
		t.Errorf("expected isMonth to be false, got true")
	}
	if !isRange {
		t.Errorf("expected isRange to be true, got false")
	}

	// 截取日期中间的一段时间
	startTime = time.Date(2025, 10, 10, 10, 0, 0, 0, loc)
	endTime = time.Date(2025, 10, 10, 12, 0, 0, 0, loc)
	isHour = false
	isDay = false
	isWeek = false
	isMonth = false
	isRange = false

	OptimizeTimeRange(startTime, endTime, ITimeRangeOptimizer{
		IsDay: func(day time.Time) {
			isDay = true
		},
		IsWeek: func(week time.Time) {
			isWeek = true
		},
		IsMonth: func(month time.Time) {
			isMonth = true
		},
		IsRange: func(start, end time.Time) {
			isRange = true
			if start != startTime {
				t.Errorf("IsRange start: expected %v, got %v", startTime, start)
			}
			if end != endTime {
				t.Errorf("IsRange end: expected %v, got %v", endTime, end)
			}
		},
	})

	if isDay {
		t.Errorf("expected isDay to be false, got true")
	}
	if isWeek {
		t.Errorf("expected isWeek to be false, got true")
	}
	if isMonth {
		t.Errorf("expected isMonth to be false, got true")
	}
	if !isRange {
		t.Errorf("expected isRange to be true, got false")
	}

	// 跨天
	startTime = StartOfDay(time.Date(2025, 10, 10, 10, 0, 0, 0, loc))
	endTime = EndOfDay(time.Date(2025, 10, 12, 12, 0, 0, 0, loc))
	isDay = false
	isWeek = false
	isMonth = false
	isRange = false

	OptimizeTimeRange(startTime, endTime, ITimeRangeOptimizer{
		IsDay: func(day time.Time) {
			isDay = true
		},
		IsWeek: func(week time.Time) {
			isWeek = true
		},
		IsMonth: func(month time.Time) {
			isMonth = true
		},
		IsRange: func(start, end time.Time) {
			isRange = true
			if start != startTime {
				t.Errorf("IsRange start: expected %v, got %v", startTime, start)
			}
			if end != endTime {
				t.Errorf("IsRange end: expected %v, got %v", endTime, end)
			}
		},
	})

	if isDay {
		t.Errorf("expected isDay to be false, got true")
	}
	if isWeek {
		t.Errorf("expected isWeek to be false, got true")
	}
	if isMonth {
		t.Errorf("expected isMonth to be false, got true")
	}
	if !isRange {
		t.Errorf("expected isRange to be true, got false")
	}

	// isWeek
	startTime = StartOfDay(time.Date(2026, 1, 4, 10, 0, 0, 0, loc))
	endTime = EndOfDay(time.Date(2026, 1, 10, 12, 0, 0, 0, loc))
	isDay = false
	isWeek = false
	isMonth = false
	isRange = false

	// 注意：这里需要根据实际的 OptimizeTimeRange 逻辑调整
	// 为了测试，我们假设时间范围正好是一周
	OptimizeTimeRange(startTime, endTime, ITimeRangeOptimizer{
		IsDay: func(day time.Time) {
			isDay = true
		},
		IsWeek: func(week time.Time) {
			isWeek = true
			if week != startTime {
				t.Errorf("IsWeek: expected %v, got %v", startTime, week)
			}
		},
		IsMonth: func(month time.Time) {
			isMonth = true
		},
		IsRange: func(start, end time.Time) {
			isRange = true
		},
	})

	if isDay {
		t.Errorf("expected isDay to be false, got true")
	}
	if !isWeek {
		t.Errorf("expected isWeek to be true, got false")
	}
	if isMonth {
		t.Errorf("expected isMonth to be false, got true")
	}
	if isRange {
		t.Errorf("expected isRange to be false, got true")
	}

	// isMonth
	startTime = StartOfDay(time.Date(2026, 1, 1, 10, 0, 0, 0, loc))
	endTime = EndOfDay(time.Date(2026, 1, 31, 12, 0, 0, 0, loc))
	isDay = false
	isWeek = false
	isMonth = false
	isRange = false

	// 注意：这里需要根据实际的 OptimizeTimeRange 逻辑调整
	// 为了测试，我们假设时间范围正好是一个月
	OptimizeTimeRange(startTime, endTime, ITimeRangeOptimizer{
		IsDay: func(day time.Time) {
			isDay = true
		},
		IsWeek: func(week time.Time) {
			isWeek = true
		},
		IsMonth: func(month time.Time) {
			isMonth = true
			if month != startTime {
				t.Errorf("IsMonth: expected %v, got %v", startTime, month)
			}
		},
		IsRange: func(start, end time.Time) {
			isRange = true
		},
	})

	if isDay {
		t.Errorf("expected isDay to be false, got true")
	}
	if isWeek {
		t.Errorf("expected isWeek to be false, got true")
	}
	if !isMonth {
		t.Errorf("expected isMonth to be true, got false")
	}
	if isRange {
		t.Errorf("expected isRange to be false, got true")
	}
}
