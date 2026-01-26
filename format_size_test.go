package hive

import "testing"

func TestFormatSize(t *testing.T) {
	if got := FormatSize(500); got != "500B" {
		t.Errorf("FormatSize(500) = %q; want \"500B\"", got)
	}
	if got := FormatSize(1024); got != "1KB" {
		t.Errorf("FormatSize(1024) = %q; want \"1KB\"", got)
	}
	if got := FormatSize(1536); got != "1.50KB" {
		t.Errorf("FormatSize(1536) = %q; want \"1.50KB\"", got)
	}
	if got := FormatSize(8.0 * 1024); got != "8KB" {
		t.Errorf("FormatSize(8*1024) = %q; want \"8KB\"", got)
	}
	if got := FormatSize(1048576); got != "1MB" {
		t.Errorf("FormatSize(1048576) = %q; want \"1MB\"", got)
	}
	if got := FormatSize(1572864); got != "1.50MB" {
		t.Errorf("FormatSize(1572864) = %q; want \"1.50MB\"", got)
	}
	if got := FormatSize(1073741824); got != "1GB" {
		t.Errorf("FormatSize(1073741824) = %q; want \"1GB\"", got)
	}
}
