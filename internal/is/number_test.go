package is

import "testing"

func TestNumber(t *testing.T) {
	if !Integer(1) {
		t.Error("error")
	}
	if !Integer(1.0) {
		t.Error("error")
	}
	if Integer(1.1) {
		t.Error("error")
	}
}
