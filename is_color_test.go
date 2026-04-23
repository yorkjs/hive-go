package hive

import "testing"

func TestIsHexColor(t *testing.T) {
	if got := IsHexColor("#666"); got != true {
		t.Errorf("IsHexColor(#666) = %v; want true", got)
	}
	if got := IsHexColor("#616161"); got != true {
		t.Errorf("IsHexColor(#616161) = %v; want true", got)
	}
}
