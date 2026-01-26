package hive

import "testing"

func TestFormatCount(t *testing.T) {
	if got := FormatCount(1234567); got != "1,234,567" {
		t.Errorf("FormatCount(1234567) = %v; want 1,234,567", got)
	}
	if got := FormatCount(1234567, "个"); got != "1,234,567个" {
		t.Errorf("FormatCount(1234567, \"个\") = %v; want 1,234,567个", got)
	}
}

func TestFormatCountShortly(t *testing.T) {
	if got := FormatCountShortly(12345); got != "1.2万" {
		t.Errorf("FormatCountShortly(12345) = %v; want 1.2万", got)
	}
	if got := FormatCountShortly(12345, "个"); got != "1.2万个" {
		t.Errorf("FormatCountShortly(12345, \"个\") = %v; want 1.2万个", got)
	}
}
