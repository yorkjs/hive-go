package hive

import (
	"testing"
	"time"
)

func TestDateDay(t *testing.T) {
	loc := time.Local

	// 2020-10-10 10:01:01
	date := time.Date(2020, 10, 10, 10, 1, 1, 0, loc).UnixMilli()

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
	date = time.Date(2020, 4, 1, 10, 1, 1, 0, loc).UnixMilli()
	if got := FormatDateTime(StartOfPrevDay(date), DATE_TIME_YEAR_MONTH_DATE_HOUR_MINUTE_SECOND); got != "2020-03-31 00:00:00" {
		t.Errorf("StartOfPrevDay(2020-04-01) = %v; want 2020-03-31 00:00:00", got)
	}
	if got := FormatDateTime(StartOfNextDay(date), DATE_TIME_YEAR_MONTH_DATE_HOUR_MINUTE_SECOND); got != "2020-04-02 00:00:00" {
		t.Errorf("StartOfNextDay(2020-04-01) = %v; want 2020-04-02 00:00:00", got)
	}

	// 2020-01-01 10:01:01
	date = time.Date(2020, 1, 1, 10, 1, 1, 0, loc).UnixMilli()
	if got := FormatDateTime(StartOfPrevDay(date), DATE_TIME_YEAR_MONTH_DATE_HOUR_MINUTE_SECOND); got != "2019-12-31 00:00:00" {
		t.Errorf("StartOfPrevDay(2020-01-01) = %v; want 2019-12-31 00:00:00", got)
	}
	if got := FormatDateTime(StartOfNextDay(date), DATE_TIME_YEAR_MONTH_DATE_HOUR_MINUTE_SECOND); got != "2020-01-02 00:00:00" {
		t.Errorf("StartOfNextDay(2020-01-01) = %v; want 2020-01-02 00:00:00", got)
	}

	// 2024-03-01 10:01:01 (Leap year)
	date = time.Date(2024, 3, 1, 10, 1, 1, 0, loc).UnixMilli()
	if got := FormatDateTime(StartOfPrevDay(date), DATE_TIME_YEAR_MONTH_DATE_HOUR_MINUTE_SECOND); got != "2024-02-29 00:00:00" {
		t.Errorf("StartOfPrevDay(2024-03-01) = %v; want 2024-02-29 00:00:00", got)
	}
	if got := FormatDateTime(StartOfNextDay(date), DATE_TIME_YEAR_MONTH_DATE_HOUR_MINUTE_SECOND); got != "2024-03-02 00:00:00" {
		t.Errorf("StartOfNextDay(2024-03-01) = %v; want 2024-03-02 00:00:00", got)
	}

	// 2025-03-01 10:01:01 (Non-leap year)
	date = time.Date(2025, 3, 1, 10, 1, 1, 0, loc).UnixMilli()
	if got := FormatDateTime(StartOfPrevDay(date), DATE_TIME_YEAR_MONTH_DATE_HOUR_MINUTE_SECOND); got != "2025-02-28 00:00:00" {
		t.Errorf("StartOfPrevDay(2025-03-01) = %v; want 2025-02-28 00:00:00", got)
	}
	if got := FormatDateTime(StartOfNextDay(date), DATE_TIME_YEAR_MONTH_DATE_HOUR_MINUTE_SECOND); got != "2025-03-02 00:00:00" {
		t.Errorf("StartOfNextDay(2025-03-01) = %v; want 2025-03-02 00:00:00", got)
	}
}

func TestDateWeek(t *testing.T) {
	loc := time.Local

	// 2025-07-27 10:01:01 (Sunday)
	date := time.Date(2025, 7, 27, 10, 1, 1, 0, loc).UnixMilli()

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
	date = time.Date(2025, 7, 29, 10, 1, 1, 0, loc).UnixMilli()
	if got := FormatDateTime(StartOfWeek(date), DATE_TIME_YEAR_MONTH_DATE_HOUR_MINUTE_SECOND); got != "2025-07-27 00:00:00" {
		t.Errorf("StartOfWeek(2025-07-29) = %v; want 2025-07-27 00:00:00", got)
	}
	if got := FormatDateTime(EndOfWeek(date), DATE_TIME_YEAR_MONTH_DATE_HOUR_MINUTE_SECOND); got != "2025-08-02 23:59:59" {
		t.Errorf("EndOfWeek(2025-07-29) = %v; want 2025-08-02 23:59:59", got)
	}

	// 2025-08-01 10:01:01 (Friday)
	date = time.Date(2025, 8, 1, 10, 1, 1, 0, loc).UnixMilli()
	if got := FormatDateTime(StartOfWeek(date), DATE_TIME_YEAR_MONTH_DATE_HOUR_MINUTE_SECOND); got != "2025-07-27 00:00:00" {
		t.Errorf("StartOfWeek(2025-08-01) = %v; want 2025-07-27 00:00:00", got)
	}
	if got := FormatDateTime(EndOfWeek(date), DATE_TIME_YEAR_MONTH_DATE_HOUR_MINUTE_SECOND); got != "2025-08-02 23:59:59" {
		t.Errorf("EndOfWeek(2025-08-01) = %v; want 2025-08-02 23:59:59", got)
	}

	// 2025-08-02 10:01:01 (Saturday)
	date = time.Date(2025, 8, 2, 10, 1, 1, 0, loc).UnixMilli()
	if got := FormatDateTime(StartOfWeek(date), DATE_TIME_YEAR_MONTH_DATE_HOUR_MINUTE_SECOND); got != "2025-07-27 00:00:00" {
		t.Errorf("StartOfWeek(2025-08-02) = %v; want 2025-07-27 00:00:00", got)
	}
	if got := FormatDateTime(EndOfWeek(date), DATE_TIME_YEAR_MONTH_DATE_HOUR_MINUTE_SECOND); got != "2025-08-02 23:59:59" {
		t.Errorf("EndOfWeek(2025-08-02) = %v; want 2025-08-02 23:59:59", got)
	}
}

func TestDateMonth(t *testing.T) {
	loc := time.Local

	// 2025-02-10 10:01:01
	date := time.Date(2025, 2, 10, 10, 1, 1, 0, loc).UnixMilli()

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
	date = time.Date(2025, 6, 29, 10, 1, 1, 0, loc).UnixMilli()
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
	date = time.Date(2025, 7, 29, 10, 1, 1, 0, loc).UnixMilli()
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
	date = time.Date(2025, 12, 19, 10, 1, 1, 0, loc).UnixMilli()
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
	date = time.Date(2025, 4, 19, 10, 1, 1, 0, loc).UnixMilli()
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
	date = time.Date(2025, 1, 19, 10, 1, 1, 0, loc).UnixMilli()
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
	date := time.Date(2025, 2, 10, 10, 1, 1, 0, loc).UnixMilli()

	// isDay
	startTime := StartOfDay(date)
	endTime := EndOfDay(date)
	isDay := false
	isWeek := false
	isMonth := false
	isRange := false

	OptimizeTimeRange(startTime, endTime, ITimeRangeOptimizer{
		IsDay: func(day int64) {
			isDay = true
			if day != startTime {
				t.Errorf("isDay start time mismatch")
			}
		},
		IsWeek:  func(week int64) { isWeek = true },
		IsMonth: func(month int64) { isMonth = true },
		IsRange: func(start, end int64) { isRange = true },
	})

	if !isDay || isWeek || isMonth || isRange {
		t.Errorf("isDay check failed: day=%v week=%v month=%v range=%v", isDay, isWeek, isMonth, isRange)
	}

	// isDay but without IsDay handler
	startTime = StartOfDay(date)
	endTime = EndOfDay(date)
	isDay = false
	isWeek = false
	isMonth = false
	isRange = false

	OptimizeTimeRange(startTime, endTime, ITimeRangeOptimizer{
		IsWeek:  func(week int64) { isWeek = true },
		IsMonth: func(month int64) { isMonth = true },
		IsRange: func(start, end int64) {
			isRange = true
			if start != startTime || end != endTime {
				t.Errorf("isRange time mismatch")
			}
		},
	})

	if isDay || isWeek || isMonth || !isRange {
		t.Errorf("isDay without handler check failed: day=%v week=%v month=%v range=%v", isDay, isWeek, isMonth, isRange)
	}

	// Partial day
	startTime = time.Date(2025, 10, 10, 10, 0, 0, 0, loc).UnixMilli()
	endTime = time.Date(2025, 10, 10, 12, 0, 0, 0, loc).UnixMilli()
	isDay = false
	isWeek = false
	isMonth = false
	isRange = false

	OptimizeTimeRange(startTime, endTime, ITimeRangeOptimizer{
		IsDay:   func(day int64) { isDay = true },
		IsWeek:  func(week int64) { isWeek = true },
		IsMonth: func(month int64) { isMonth = true },
		IsRange: func(start, end int64) {
			isRange = true
			if start != startTime || end != endTime {
				t.Errorf("isRange time mismatch")
			}
		},
	})
	if isDay || isWeek || isMonth || !isRange {
		t.Errorf("Partial day check failed: day=%v week=%v month=%v range=%v", isDay, isWeek, isMonth, isRange)
	}

	// Cross day
	startTime = StartOfDay(time.Date(2025, 10, 10, 10, 0, 0, 0, loc).UnixMilli())
	endTime = EndOfDay(time.Date(2025, 10, 12, 12, 0, 0, 0, loc).UnixMilli())
	isDay = false
	isWeek = false
	isMonth = false
	isRange = false

	OptimizeTimeRange(startTime, endTime, ITimeRangeOptimizer{
		IsDay:   func(day int64) { isDay = true },
		IsWeek:  func(week int64) { isWeek = true },
		IsMonth: func(month int64) { isMonth = true },
		IsRange: func(start, end int64) {
			isRange = true
			if start != startTime || end != endTime {
				t.Errorf("isRange time mismatch")
			}
		},
	})
	if isDay || isWeek || isMonth || !isRange {
		t.Errorf("Cross day check failed: day=%v week=%v month=%v range=%v", isDay, isWeek, isMonth, isRange)
	}

	// isWeek
	startTime = StartOfDay(time.Date(2026, 1, 4, 10, 0, 0, 0, loc).UnixMilli()) // 2026-01-04 is Sunday
	endTime = EndOfDay(time.Date(2026, 1, 10, 12, 0, 0, 0, loc).UnixMilli())    // 2026-01-10 is Saturday
	isDay = false
	isWeek = false
	isMonth = false
	isRange = false

	OptimizeTimeRange(startTime, endTime, ITimeRangeOptimizer{
		IsDay: func(day int64) { isDay = true },
		IsWeek: func(week int64) {
			isWeek = true
			if week != startTime {
				t.Errorf("isWeek start time mismatch")
			}
		},
		IsMonth: func(month int64) { isMonth = true },
		IsRange: func(start, end int64) { isRange = true },
	})
	if isDay || !isWeek || isMonth || isRange {
		t.Errorf("isWeek check failed: day=%v week=%v month=%v range=%v", isDay, isWeek, isMonth, isRange)
	}

	// isMonth
	startTime = StartOfDay(time.Date(2026, 1, 1, 10, 0, 0, 0, loc).UnixMilli())
	endTime = EndOfDay(time.Date(2026, 1, 31, 12, 0, 0, 0, loc).UnixMilli())
	isDay = false
	isWeek = false
	isMonth = false
	isRange = false

	OptimizeTimeRange(startTime, endTime, ITimeRangeOptimizer{
		IsDay:  func(day int64) { isDay = true },
		IsWeek: func(week int64) { isWeek = true },
		IsMonth: func(month int64) {
			isMonth = true
			if month != startTime {
				t.Errorf("isMonth start time mismatch")
			}
		},
		IsRange: func(start, end int64) { isRange = true },
	})
	if isDay || isWeek || !isMonth || isRange {
		t.Errorf("isMonth check failed: day=%v week=%v month=%v range=%v", isDay, isWeek, isMonth, isRange)
	}
}
