package hive

import "testing"

func TestRateConvert(t *testing.T) {
	if got := RateToDisplay(1015); got != 10.15 {
		t.Errorf("RateToDisplay(1015) = %v; want 10.15", got)
	}
	if got := RateToBackend(10.15); got != 1015 {
		t.Errorf("RateToBackend(10.15) = %v; want 1015", got)
	}
	if got := CalculateRate(10, 100); got != 1000 {
		t.Errorf("CalculateRate(10, 100) = %v; want 1000", got)
	}
	if got := CalculateRate(5, 1); got != 50000 {
		t.Errorf("CalculateRate(5, 1) = %v; want 50000", got)
	}
	if got := CalculateRate(5, 5); got != 10000 {
		t.Errorf("CalculateRate(5, 5) = %v; want 10000", got)
	}
	if got := CalculateRate(5, 10); got != 5000 {
		t.Errorf("CalculateRate(5, 10) = %v; want 5000", got)
	}
	if got := CalculateRate(5, 100); got != 500 {
		t.Errorf("CalculateRate(5, 100) = %v; want 500", got)
	}
	if got := CalculateRate(5, 1000); got != 50 {
		t.Errorf("CalculateRate(5, 1000) = %v; want 50", got)
	}
	if got := CalculateRate(5, 10000); got != 5 {
		t.Errorf("CalculateRate(5, 10000) = %v; want 5", got)
	}
}
