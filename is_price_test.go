package hive

import "testing"

func TestIsPrice(t *testing.T) {
	tests := []struct {
		input string
		want  bool
	}{
		{"0", true},
		{"1", true},
		{"1.0", true},
		{"1.1", true},
		{"1.12", true},
		{"123.12", true},
		{"1.123", false},
		{"+1", false},
		{"-1", false},
		{"abc", false},
		{"1.", false},
		{"", false},
	}

	for _, tt := range tests {
		if got := IsPrice(tt.input); got != tt.want {
			t.Errorf("IsPrice(%q) = %v; want %v", tt.input, got, tt.want)
		}
	}
}
