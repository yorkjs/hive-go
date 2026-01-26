package hive

import "testing"

func TestNormalizeVersion(t *testing.T) {
	if got := NormalizeVersion("1"); got != "000000000001" {
		t.Errorf("NormalizeVersion(\"1\") = %v; want 000000000001", got)
	}
	if got := NormalizeVersion("1.2"); got != "000001000002" {
		t.Errorf("NormalizeVersion(\"1.2\") = %v; want 000001000002", got)
	}
	if got := NormalizeVersion("1.2.3"); got != "000100020003" {
		t.Errorf("NormalizeVersion(\"1.2.3\") = %v; want 000100020003", got)
	}
	if got := NormalizeVersion(""); got != "000000000000" {
		t.Errorf("NormalizeVersion(\"\") = %v; want 000000000000", got)
	}
}
