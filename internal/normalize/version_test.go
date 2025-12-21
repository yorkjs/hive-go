package normalize

import (
	"testing"
)

func TestVersion(t *testing.T) {
	if Version("1.0.10") <= Version("1.0.9") {
		t.Error("error")
	}
	if Version("1.0.09") >= Version("1.0.10") {
		t.Error("error")
	}
	if Version("1.0.09") != "000100000009" {
		t.Error("error")
	}
}
