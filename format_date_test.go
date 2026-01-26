package hive

import (
	"testing"
	"time"
)

func TestFormatDate(t *testing.T) {
	// 2020-11-11 00:00:00 UTC
	t1 := time.Date(2020, 11, 11, 0, 0, 0, 0, time.UTC).UnixMilli()
	if got := FormatDate(t1); got != "2020-11-11" {
		t.Errorf("FormatDate(2020-11-11) = %v; want 2020-11-11", got)
	}

	t2 := time.Date(2020, 10, 1, 0, 0, 0, 0, time.UTC).UnixMilli()
	if got := FormatDate(t2); got != "2020-10-01" {
		t.Errorf("FormatDate(2020-10-01) = %v; want 2020-10-01", got)
	}

	// Slash format
	if got := FormatDate(t1, DATE_YEAR_MONTH_DATE_SLASH); got != "2020/11/11" {
		t.Errorf("FormatDate(..., SLASH) = %v; want 2020/11/11", got)
	}

	// Chinese format
	// 2020-01-01
	t3 := time.Date(2020, 1, 1, 0, 0, 0, 0, time.UTC).UnixMilli()
	if got := FormatDate(t3, DATE_YEAR_MONTH_DATE_CHINESE); got != "2020年1月1日" {
		t.Errorf("FormatDate(2020-01-01, CHINESE) = %v; want 2020年1月1日", got)
	}
	if got := FormatDate(t2, DATE_YEAR_MONTH_DATE_CHINESE); got != "2020年10月1日" {
		t.Errorf("FormatDate(2020-10-01, CHINESE) = %v; want 2020年10月1日", got)
	}
}

func TestFormatDateShortly(t *testing.T) {
	// Current year
	now := time.Now()
	t1 := time.Date(now.Year(), 10, 1, 0, 0, 0, 0, time.Local).UnixMilli()
	if got := FormatDateShortly(t1); got != "10-01" {
		t.Errorf("FormatDateShortly(CurrentYear) = %v; want 10-01", got)
	}

	// Past year
	t2 := time.Date(2020, 10, 1, 0, 0, 0, 0, time.Local).UnixMilli()
	if got := FormatDateShortly(t2); got != "2020-10-01" {
		t.Errorf("FormatDateShortly(2020) = %v; want 2020-10-01", got)
	}
}

func TestFormatDateRange(t *testing.T) {
	t1 := time.Date(2020, 10, 1, 10, 1, 2, 0, time.UTC).UnixMilli()
	t2 := time.Date(2020, 10, 2, 10, 3, 2, 0, time.UTC).UnixMilli()
	
	if got := FormatDateRange(t1, t2); got != "2020-10-01 至 2020-10-02" {
		t.Errorf("FormatDateRange(...) = %v; want 2020-10-01 至 2020-10-02", got)
	}
}
