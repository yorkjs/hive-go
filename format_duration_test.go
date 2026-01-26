package hive

import (
	"testing"
)

func TestFormatDuration(t *testing.T) {
	// 1 day 1 hour 1 min 1 sec = 24*3600 + 3600 + 60 + 1 = 86400 + 3600 + 61 = 90061 sec = 90061000 ms
	ms := int64(90061000)
	if got := FormatDuration(ms); got != "1天1小时1分钟1秒" {
		t.Errorf("FormatDuration(%d) = %v; want 1天1小时1分钟1秒", ms, got)
	}

	// 1 hour
	ms = int64(3600000)
	if got := FormatDuration(ms); got != "1小时" {
		t.Errorf("FormatDuration(%d) = %v; want 1小时", ms, got)
	}
}
