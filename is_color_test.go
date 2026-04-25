package hive

import "testing"

func TestIsHexColor(t *testing.T) {
	tests := []struct {
		input string
		want  bool
		desc  string
	}{
		{"#666", true, "3位hex"},
		{"#616161", true, "6位hex"},
		{"#6", false, "1位hex"},
		{"#66", false, "2位hex"},
		{"#6666", true, "4位hex (rgba)"},
		{"#66666", false, "5位hex"},
		{"#6666666", false, "7位hex"},
		{"#66666666", true, "8位hex (rgba)"},
		{"#666666666", false, "9位hex"},
		{"666", false, "无#的3位"},
		{"666666", false, "无#的6位"},
	}

	for _, tt := range tests {
		t.Run(tt.desc, func(t *testing.T) {
			got := IsHexColor(tt.input)
			if got != tt.want {
				t.Errorf("IsHexColor(%q) = %v, want %v", tt.input, got, tt.want)
			}
		})
	}
}
