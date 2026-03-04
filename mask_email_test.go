package hive

import "testing"

func TestMaskEmail(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		expected string
	}{
		{
			name:     "empty string",
			input:    "",
			expected: "",
		},
		{
			name:     "single character",
			input:    "你",
			expected: "你",
		},
		{
			name:     "valid email",
			input:    "user@example.com",
			expected: "***r@example.com",
		},
		{
			name:     "valid email",
			input:    "a@example.com",
			expected: "***@example.com",
		},
		{
			name:     "invalid email",
			input:    "@example.com",
			expected: "@example.com",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := MaskEmail(tt.input)
			if result != tt.expected {
				t.Errorf("MaskEmail(%q) = %q, want %q", tt.input, result, tt.expected)
			}
		})
	}
}
