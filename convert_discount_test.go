package hive

import "testing"

func TestDiscountConvert(t *testing.T) {
	if got := DiscountToDisplay(8000); got != 8 {
		t.Errorf("DiscountToDisplay(8000) = %v; want 8", got)
	}
	if got := DiscountToDisplay(8800); got != 8.8 {
		t.Errorf("DiscountToDisplay(8800) = %v; want 8.8", got)
	}
	if got := DiscountToDisplay(8880); got != 8.8 {
		t.Errorf("DiscountToDisplay(8880) = %v; want 8.8", got)
	}

	if got := DiscountToBackend(8); got != 8000 {
		t.Errorf("DiscountToBackend(8) = %v; want 8000", got)
	}
	if got := DiscountToBackend(8.8); got != 8800 {
		t.Errorf("DiscountToBackend(8.8) = %v; want 8800", got)
	}
	if got := DiscountToBackend(8.88); got != 8800 {
		t.Errorf("DiscountToBackend(8.88) = %v; want 8800", got)
	}
}
