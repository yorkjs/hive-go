package hive

import "testing"

func TestFormatDistance(t *testing.T) {
	// DistanceToDisplay: value / 1000
	if got := FormatDistance(1500); got != "1.5公里" {
		t.Errorf("FormatDistance(1500) = %v; want 1.5公里", got)
	}
	if got := FormatDistance(500); got != "0.5公里" {
		t.Errorf("FormatDistance(500) = %v; want 0.5公里", got)
	}
}
