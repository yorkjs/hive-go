package hive

import "testing"

func TestDistanceConvert(t *testing.T) {
	if got := DistanceToDisplay(8000); got != 8 {
		t.Errorf("DistanceToDisplay(8000) = %v; want 8", got)
	}
	if got := DistanceToDisplay(8800); got != 8.8 {
		t.Errorf("DistanceToDisplay(8800) = %v; want 8.8", got)
	}
	if got := DistanceToDisplay(8880); got != 8.88 {
		t.Errorf("DistanceToDisplay(8880) = %v; want 8.88", got)
	}
	if got := DistanceToDisplay(8888); got != 8.888 {
		t.Errorf("DistanceToDisplay(8888) = %v; want 8.888", got)
	}

	if got := DistanceToBackend[int](8); got != 8000 {
		t.Errorf("DistanceToBackend(8) = %v; want 8000", got)
	}
	if got := DistanceToBackend[int](8.8); got != 8800 {
		t.Errorf("DistanceToBackend(8.8) = %v; want 8800", got)
	}
	if got := DistanceToBackend[int](8.88); got != 8880 {
		t.Errorf("DistanceToBackend(8.88) = %v; want 8880", got)
	}
	if got := DistanceToBackend[int](8.888); got != 8888 {
		t.Errorf("DistanceToBackend(8.888) = %v; want 8888", got)
	}

}
