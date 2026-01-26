package hive

import "testing"

func TestFormatShelfLife(t *testing.T) {
	// 5 hours
	if got := FormatShelfLife(5); got != "5小时" {
		t.Errorf("FormatShelfLife(5) = %v; want 5小时", got)
	}

	// 24 hours = 1 day
	if got := FormatShelfLife(24); got != "1天" {
		t.Errorf("FormatShelfLife(24) = %v; want 1天", got)
	}

	// 30 hours = 1 day 6 hours
	if got := FormatShelfLife(30); got != "1天6小时" {
		t.Errorf("FormatShelfLife(30) = %v; want 1天6小时", got)
	}

	// SHELF_LIFE_MONTH
	if got := FormatShelfLife(SHELF_LIFE_MONTH); got != "1个月" {
		t.Errorf("FormatShelfLife(SHELF_LIFE_MONTH) = %v; want 1个月", got)
	}

	// 35 * SHELF_LIFE_DAY = 35 * 24 = 840 hours
	// 1 month (720) + 5 days (120) = 840
	if got := FormatShelfLife(35 * SHELF_LIFE_DAY); got != "1个月5天" {
		t.Errorf("FormatShelfLife(35*DAY) = %v; want 1个月5天", got)
	}

	// 365 * SHELF_LIFE_DAY = 1 year
	if got := FormatShelfLife(365 * SHELF_LIFE_DAY); got != "1年" {
		t.Errorf("FormatShelfLife(365*DAY) = %v; want 1年", got)
	}

	// SHELF_LIFE_YEAR + 6 * SHELF_LIFE_MONTH
	if got := FormatShelfLife(SHELF_LIFE_YEAR + 6*SHELF_LIFE_MONTH); got != "1年6个月" {
		t.Errorf("FormatShelfLife(YEAR + 6*MONTH) = %v; want 1年6个月", got)
	}
}
