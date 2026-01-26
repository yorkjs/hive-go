package hive

import "testing"

func TestIsInteger(t *testing.T) {
	if got := IsInteger(1); got != true {
		t.Errorf("IsInteger(1) = %v; want true", got)
	}
	if got := IsInteger(1.0); got != true {
		t.Errorf("IsInteger(1.0) = %v; want true", got)
	}
	if got := IsInteger(1.1); got != false {
		t.Errorf("IsInteger(1.1) = %v; want false", got)
	}
}
