package hive

import "testing"

func TestMaskName(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		expected string
	}{
		{
			name:     "empty string",
			input:    "",
			expected: "***",
		},
		{
			name:     "single character",
			input:    "你",
			expected: "***",
		},
		{
			name:     "two-character name",
			input:    "你好",
			expected: "***好",
		},
		{
			name:     "three-character name",
			input:    "你好呀",
			expected: "***呀",
		},
		{
			name:     "mixed Chinese and alphanumeric",
			input:    "你好呀123",
			expected: "***3",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := MaskName(tt.input)
			if result != tt.expected {
				t.Errorf("MaskName(%q) = %q, want %q", tt.input, result, tt.expected)
			}
		})
	}
}
