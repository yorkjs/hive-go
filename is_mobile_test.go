package hive

import "testing"

func TestIsMobile(t *testing.T) {
	testCases := []struct {
		name     string
		input    string
		expected bool
	}{
		{"valid mobile", "13512345678", true},
		{"too short", "1351234567", false},
		{"too long", "135123456789", false},
		{"leading space", " 13512345678", false},
		{"trailing space", "13512345678 ", false},
		{"both spaces", " 13512345678 ", false},
		{"empty string", "", false},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			result := IsMobile(tc.input)
			if result != tc.expected {
				t.Errorf("IsMobile(%q) = %v; expected %v", tc.input, result, tc.expected)
			}
		})
	}
}
