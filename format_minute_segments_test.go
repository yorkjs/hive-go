package hive

import (
	"testing"
)

func TestFormatMinuteSegments(t *testing.T) {
	tests := []struct {
		name     string
		segments []int
		sep      string // 空字符串表示使用默认分隔符
		want     string
	}{
		{"empty", []int{}, "", ""},
		{"odd_length", []int{0}, "", ""},
		{"full_day", []int{0, 1440}, "", "全天"},
		{"two_full_days", []int{0, 1440, 1440, 2880}, "", "全天、全天"},
		{"mixed", []int{540, 960, 1200, 1560}, "", "09:00-16:00、20:00-次日02:00"},
		{"mixed_with_plus", []int{540, 960, 1200, 1560}, "+", "09:00-16:00+20:00-次日02:00"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var got string
			if tt.sep == "" {
				got = FormatMinuteSegments(tt.segments) // 使用默认分隔符
			} else {
				got = FormatMinuteSegments(tt.segments, tt.sep) // 指定分隔符
			}
			if got != tt.want {
				t.Errorf("FormatMinuteSegments(%v, %q) = %q, want %q", tt.segments, tt.sep, got, tt.want)
			}
		})
	}
}
