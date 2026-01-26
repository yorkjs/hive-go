package hive

import "testing"

func TestIsEmail(t *testing.T) {
	tests := []struct {
		input string
		want  bool
	}{
		{"abc@xxx.com", true},
		{"a@163.com", true},
		{"abc@qq.com", true},
		{"abc@qq.cn", true},
		{"abc@qq.com.cn", true},
		{"abc@qq.net", true},
		{"abc-123@qq.net", true},
		{"abc-123qq.net", false},
		{"abc-123@qq", false},
		{"@qq", false},
		{"@qq.com", false},
		{"中文@qq.com", false},
	}

	for _, tt := range tests {
		if got := IsEmail(tt.input); got != tt.want {
			t.Errorf("IsEmail(%q) = %v; want %v", tt.input, got, tt.want)
		}
	}
}
