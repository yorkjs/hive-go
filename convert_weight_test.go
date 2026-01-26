package hive

import "testing"

func TestWeightConvert(t *testing.T) {
	if got := WeightToG(1); got != 0.001 {
		t.Errorf("WeightToG(1) = %v; want 0.001", got)
	}
	if got := WeightToG(10); got != 0.01 {
		t.Errorf("WeightToG(10) = %v; want 0.01", got)
	}
	if got := WeightToG(100); got != 0.1 {
		t.Errorf("WeightToG(100) = %v; want 0.1", got)
	}
	if got := WeightToG(1000); got != 1 {
		t.Errorf("WeightToG(1000) = %v; want 1", got)
	}
	if got := WeightToG(10000); got != 10 {
		t.Errorf("WeightToG(10000) = %v; want 10", got)
	}

	if got := WeightToKG(1); got != 0.000001 {
		t.Errorf("WeightToKG(1) = %v; want 0.000001", got)
	}
	if got := WeightToKG(10); got != 0.00001 {
		t.Errorf("WeightToKG(10) = %v; want 0.00001", got)
	}
	if got := WeightToKG(100); got != 0.0001 {
		t.Errorf("WeightToKG(100) = %v; want 0.0001", got)
	}
	if got := WeightToKG(1000); got != 0.001 {
		t.Errorf("WeightToKG(1000) = %v; want 0.001", got)
	}
	if got := WeightToKG(10000); got != 0.01 {
		t.Errorf("WeightToKG(10000) = %v; want 0.01", got)
	}
	if got := WeightToKG(100000); got != 0.1 {
		t.Errorf("WeightToKG(100000) = %v; want 0.1", got)
	}
	if got := WeightToKG(1000000); got != 1 {
		t.Errorf("WeightToKG(1000000) = %v; want 1", got)
	}
	if got := WeightToKG(10000000); got != 10 {
		t.Errorf("WeightToKG(10000000) = %v; want 10", got)
	}

	if got := WeightGToBackend(0.001); got != 1 {
		t.Errorf("WeightGToBackend(0.001) = %v; want 1", got)
	}
	if got := WeightGToBackend(0.01); got != 10 {
		t.Errorf("WeightGToBackend(0.01) = %v; want 10", got)
	}
	if got := WeightGToBackend(0.1); got != 100 {
		t.Errorf("WeightGToBackend(0.1) = %v; want 100", got)
	}
	if got := WeightGToBackend(1); got != 1000 {
		t.Errorf("WeightGToBackend(1) = %v; want 1000", got)
	}
	if got := WeightGToBackend(10); got != 10000 {
		t.Errorf("WeightGToBackend(10) = %v; want 10000", got)
	}

	if got := WeightKGToBackend(0.001); got != 1000 {
		t.Errorf("WeightKGToBackend(0.001) = %v; want 1000", got)
	}
	if got := WeightKGToBackend(0.01); got != 10000 {
		t.Errorf("WeightKGToBackend(0.01) = %v; want 10000", got)
	}
	if got := WeightKGToBackend(0.1); got != 100000 {
		t.Errorf("WeightKGToBackend(0.1) = %v; want 100000", got)
	}
	if got := WeightKGToBackend(1); got != 1000000 {
		t.Errorf("WeightKGToBackend(1) = %v; want 1000000", got)
	}
	if got := WeightKGToBackend(10); got != 10000000 {
		t.Errorf("WeightKGToBackend(10) = %v; want 10000000", got)
	}
}
