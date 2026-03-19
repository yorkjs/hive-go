package hive

import (
	"testing"
	"time"
)

func TestFormatBirthday(t *testing.T) {
	t1 := time.Date(2023, 10, 1, 0, 0, 0, 0, time.Local)
	if got := FormatBirthday(t1); got != "10.01" {
		t.Errorf("FormatBirthday(ts) = %v; want 10.01", got)
	}
}
