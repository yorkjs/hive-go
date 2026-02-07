package hive

import "testing"

func TestIsVerifyCode(t *testing.T) {
	tests := []struct {
		input string
		want  bool
	}{
		{"111222", true},
		{"11122", false},
		{"1112222", false},
		{"11122a", false},
		{"aaaaaa", false},
		{"", false},
	}

	for _, tt := range tests {
		if got := IsVerifyCode(tt.input); got != tt.want {
			t.Errorf("IsVerifyCode(%q) = %v; want %v", tt.input, got, tt.want)
		}
	}
}
