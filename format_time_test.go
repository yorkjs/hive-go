package hive

import (
	"testing"
	"time"
)

func TestFormatTime(t *testing.T) {
	ts := time.Date(2023, 10, 1, 12, 21, 25, 0, time.Local)
	if got := FormatTime(ts); got != "12:21:25" {
		t.Errorf("FormatTime(ts) = %v; want 12:21:25", got)
	}
}
