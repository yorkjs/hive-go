package hive

import (
	"testing"
	"time"
)

func TestFormatWeek(t *testing.T) {
	// 2000-01-01 is Saturday.
	// Week starts on Sunday (1999-12-26).
	// End of week (Saturday) is 2000-01-01.
	// Format: 1999-12-26 至 2000-01-01
	
	ts := time.Date(2000, 1, 1, 0, 0, 0, 0, time.Local).UnixMilli()
	want := "1999-12-26 至 2000-01-01"
	if got := FormatWeek(ts); got != want {
		t.Errorf("FormatWeek(2000-01-01) = %v; want %v", got, want)
	}
}
