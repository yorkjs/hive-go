package hive

import "testing"

func TestMaskMobile(t *testing.T) {
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
			name:     "valid 11-digit mobile",
			input:    "13812345678",
			expected: "138****5678",
		},
		{
			name:     "10-digit number",
			input:    "1234567890",
			expected: "1234567890",
		},
		{
			name:     "5-digit number",
			input:    "12345",
			expected: "12345",
		},
		{
			name:     "formatted phone number",
			input:    "138-0013-8000",
			expected: "138-0013-8000",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := MaskMobile(tt.input)
			if result != tt.expected {
				t.Errorf("MaskMobile(%q) = %q, want %q", tt.input, result, tt.expected)
			}
		})
	}
}
