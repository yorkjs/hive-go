package hive

import "testing"

func TestRateUtils(t *testing.T) {
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

	if got := ApplyRateFloor(1000, 0); got != 0 {
		t.Errorf("ApplyRateFloor(1000, 0) = %v; want 0", got)
	}
	if got := ApplyRateFloor(1000, 1000); got != 100 {
		t.Errorf("ApplyRateFloor(1000, 1000) = %v; want 100", got)
	}
	if got := ApplyRateFloor(1000, 10000); got != 1000 {
		t.Errorf("ApplyRateFloor(1000, 10000) = %v; want 1000", got)
	}
	if got := ApplyRateFloor(1000, 245); got != 24 {
		t.Errorf("ApplyRateFloor(1000, 245) = %v; want 24", got)
	}

	if got := ApplyRateCeil(1000, 0); got != 0 {
		t.Errorf("ApplyRateCeil(1000, 0) = %v; want 0", got)
	}
	if got := ApplyRateCeil(1000, 1000); got != 100 {
		t.Errorf("ApplyRateCeil(1000, 1000) = %v; want 100", got)
	}
	if got := ApplyRateCeil(1000, 10000); got != 1000 {
		t.Errorf("ApplyRateCeil(1000, 10000) = %v; want 1000", got)
	}
	if got := ApplyRateCeil(1000, 245); got != 25 {
		t.Errorf("ApplyRateCeil(1000, 245) = %v; want 25", got)
	}

	if got := ApplyRateRound(1000, 0); got != 0 {
		t.Errorf("ApplyRateRound(1000, 0) = %v; want 0", got)
	}
	if got := ApplyRateRound(1000, 1000); got != 100 {
		t.Errorf("ApplyRateRound(1000, 1000) = %v; want 100", got)
	}
	if got := ApplyRateRound(1000, 10000); got != 1000 {
		t.Errorf("ApplyRateRound(1000, 10000) = %v; want 1000", got)
	}
	if got := ApplyRateRound(1000, 245); got != 25 {
		t.Errorf("ApplyRateRound(1000, 245) = %v; want 25", got)
	}
	if got := ApplyRateRound(1000, 244); got != 24 {
		t.Errorf("ApplyRateRound(1000, 244) = %v; want 24", got)
	}
}
