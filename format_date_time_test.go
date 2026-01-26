package hive

import (
	"testing"
	"time"
)

func TestFormatDateTime(t *testing.T) {
	// 2020-11-11 10:01:02
	t1 := time.Date(2020, 11, 11, 10, 1, 2, 0, time.Local)
	if got := FormatDateTime(t1); got != "2020-11-11 10:01" {
		t.Errorf("FormatDateTime(2020-11-11 10:01) = %q; want \"2020-11-11 10:01\"", got)
	}

	// 2020-10-01 01:01:02
	t2 := time.Date(2020, 10, 1, 1, 1, 2, 0, time.Local)
	if got := FormatDateTime(t2); got != "2020-10-01 01:01" {
		t.Errorf("FormatDateTime(2020-10-01 01:01) = %q; want \"2020-10-01 01:01\"", got)
	}

	// 2020-01-10 10:10:02
	t3 := time.Date(2020, 1, 10, 10, 10, 2, 0, time.Local)
	if got := FormatDateTime(t3); got != "2020-01-10 10:10" {
		t.Errorf("FormatDateTime(2020-01-10 10:10) = %q; want \"2020-01-10 10:10\"", got)
	}

	// 2020-01-01 10:10:02
	t4 := time.Date(2020, 1, 1, 10, 10, 2, 0, time.Local)
	if got := FormatDateTime(t4); got != "2020-01-01 10:10" {
		t.Errorf("FormatDateTime(2020-01-01 10:10) = %q; want \"2020-01-01 10:10\"", got)
	}

	// With format DATE_TIME_YEAR_MONTH_DATE_HOUR_MINUTE_SLASH "2006/01/02 15:04"
	if got := FormatDateTime(t1, DATE_TIME_YEAR_MONTH_DATE_HOUR_MINUTE_SLASH); got != "2020/11/11 10:01" {
		t.Errorf("FormatDateTime(..., SLASH) = %q; want \"2020/11/11 10:01\"", got)
	}

	// With format DATE_TIME_YEAR_MONTH_DATE_HOUR_MINUTE_CHINESE "2006年1月2日 15:04"
	// Note: Go layout "1月2日" uses non-padded month/day if "1" is used, padded if "01" is used.
	// The constant in TS might imply behavior that Go layout handles differently.
	// TS: YYYY年M月D日 -> 2020年11月11日, 2020年1月1日
	// Go constant: "2006年1月2日 15:04" -> This layout means month/day are NOT zero-padded.
	// Let's check TS expectation:
	// '2020年11月11日 10:01'
	// '2020年10月1日 01:01' (October 1st, Month 10, Day 1)
	// '2020年1月10日 10:10'
	// '2020年1月1日 10:10'
	// Go's "1" matches "1", "10", "11". "2" matches "2", "11", "20".
	// So "2006年1月2日" should produce the correct output for these cases.

	if got := FormatDateTime(t1, DATE_TIME_YEAR_MONTH_DATE_HOUR_MINUTE_CHINESE); got != "2020年11月11日 10:01" {
		t.Errorf("FormatDateTime(..., CHINESE) = %q; want \"2020年11月11日 10:01\"", got)
	}
	if got := FormatDateTime(t2, DATE_TIME_YEAR_MONTH_DATE_HOUR_MINUTE_CHINESE); got != "2020年10月1日 01:01" {
		t.Errorf("FormatDateTime(..., CHINESE) = %q; want \"2020年10月1日 01:01\"", got)
	}
	if got := FormatDateTime(t3, DATE_TIME_YEAR_MONTH_DATE_HOUR_MINUTE_CHINESE); got != "2020年1月10日 10:10" {
		t.Errorf("FormatDateTime(..., CHINESE) = %q; want \"2020年1月10日 10:10\"", got)
	}
	if got := FormatDateTime(t4, DATE_TIME_YEAR_MONTH_DATE_HOUR_MINUTE_CHINESE); got != "2020年1月1日 10:10" {
		t.Errorf("FormatDateTime(..., CHINESE) = %q; want \"2020年1月1日 10:10\"", got)
	}

	// FormatDateTimeShortly
	// TS: formatDateTimeShortly(new Date('2026-10-01 10:01:02').getTime()) -> '10-01 10:01' (Current year is 2026? Env says 2026-01-26)
	// If input year == current year, show "MM-DD HH:MM". Else show "YYYY-MM-DD HH:MM".
	// Env says Today's date: 2026-01-26.
	// Case 1: 2026-10-01 (Same year)
	tCurrentYear := time.Date(2026, 10, 1, 10, 1, 2, 0, time.Local)
	if got := FormatDateTimeShortly(tCurrentYear); got != "10-01 10:01" {
		t.Errorf("FormatDateTimeShortly(2026...) = %q; want \"10-01 10:01\"", got)
	}

	// Case 2: 2020-10-01 (Different year)
	tOtherYear := time.Date(2020, 10, 1, 10, 1, 2, 0, time.Local)
	if got := FormatDateTimeShortly(tOtherYear); got != "2020-10-01 10:01" {
		t.Errorf("FormatDateTimeShortly(2020...) = %q; want \"2020-10-01 10:01\"", got)
	}

	// FormatDateTimeRange
	// '2020-10-01 10:01 至 2020-10-01 10:03'
	tStart := time.Date(2020, 10, 1, 10, 1, 2, 0, time.Local)
	tEnd := time.Date(2020, 10, 1, 10, 3, 2, 0, time.Local)
	if got := FormatDateTimeRange(tStart, tEnd); got != "2020-10-01 10:01 至 2020-10-01 10:03" {
		t.Errorf("FormatDateTimeRange(...) = %q; want \"2020-10-01 10:01 至 2020-10-01 10:03\"", got)
	}
}
