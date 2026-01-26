package hive

import (
    "testing"
    "time"
)

func TestFormatMonth(t *testing.T) {
    t1 := time.Date(2020, 10, 1, 0, 0, 0, 0, time.Local).UnixMilli()
	if got := FormatMonth(t1); got != "2020-10" {
		t.Errorf("FormatMonth(2020-10-01) = %q; want \"2020-10\"", got)
	}
}
