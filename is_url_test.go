package hive

import "testing"

func TestIsUrl(t *testing.T) {

	tests := []struct {
		input string
		want  bool
	}{
		{"", false},
		{"https://www.baidu.com", true},
		{"http://www.baidu.com", true},
		{"htt://www.baidu.com", false},
		{"ftp://www.baidu.com", false},
		{"https://www.baidu.com/", true},
		{"https://www.baidu.com/a", true},
		{"https://www.baidu.com/a/b", true},
		{"https://www.baidu.com/a/b/c", true},
		{"https://www.baidu.com/a/b/c?a=1", true},
		{"https://www.baidu.com/a/b/c?a=1&b=2", true},
		{"https://www.baidu.com/a/b/c?a=1&b=2#hash", true},
	}

	for _, tt := range tests {
		if got := IsUrl(tt.input); got != tt.want {
			t.Errorf("IsUrl(%q) = %v; want %v", tt.input, got, tt.want)
		}
	}
}
