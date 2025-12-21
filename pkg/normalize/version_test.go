package normalize

import (
	"testing"
)

func TestNormalizeVersion(t *testing.T) {
	if NormalizeVersion("1.0.10") <= NormalizeVersion("1.0.9") {
		t.Error("error")
	}
	if NormalizeVersion("1.0.09") >= NormalizeVersion("1.0.10") {
		t.Error("error")
	}
	if NormalizeVersion("1.0.09") != "000100000009" {
		t.Error("error")
	}
}
