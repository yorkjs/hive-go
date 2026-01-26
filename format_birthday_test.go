package hive

import (
	"testing"
	"time"
)

func TestFormatBirthday(t *testing.T) {
	ts := time.Date(2023, 10, 1, 0, 0, 0, 0, time.Local).UnixMilli()
	if got := FormatBirthday(ts); got != "10.01" {
		t.Errorf("FormatBirthday(ts) = %v; want 10.01", got)
	}
}
