package hive

import (
	"testing"
	"time"
)

func TestFormatYear(t *testing.T) {
	ts := time.Date(2023, 10, 1, 0, 0, 0, 0, time.Local).UnixMilli()
	if got := FormatYear(ts); got != "2023" {
		t.Errorf("FormatYear(ts) = %v; want 2023", got)
	}
}
