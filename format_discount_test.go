package hive

import "testing"

func TestFormatDiscount(t *testing.T) {
	// DiscountToDisplay: value / 1000
	// 9000 -> 9折
	if got := FormatDiscount(9000); got != "9折" {
		t.Errorf("FormatDiscount(9000) = %v; want 9折", got)
	}
	// 9500 -> 9.5折
	if got := FormatDiscount(9500); got != "9.5折" {
		t.Errorf("FormatDiscount(9500) = %v; want 9.5折", got)
	}
}
