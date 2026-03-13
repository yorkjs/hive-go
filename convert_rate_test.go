package hive

import "testing"

func TestRateConvert(t *testing.T) {
	if got := RateToDisplay(1015); got != 10.15 {
		t.Errorf("RateToDisplay(1015) = %v; want 10.15", got)
	}
	if got := RateToBackend(10.15); got != 1015 {
		t.Errorf("RateToBackend(10.15) = %v; want 1015", got)
	}
}
