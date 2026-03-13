package hive

import "testing"

func TestIsCorporateAccountNumber(t *testing.T) {
	testCases := []struct {
		name     string
		input    string
		expected bool
	}{
		{"empty string", "", false},
		{"single digit", "1", false},
		{"two digits", "12", false},
		{"too long", "6228480012123123123123123345678", false},
		{"valid 13 digits", "1234567890123", true},
		{"non-numeric characters", "asdasdasdasd", false},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			result := IsCorporateAccountNumber(tc.input)
			if result != tc.expected {
				t.Errorf("IsCorporateAccountNumber(%q) = %v; expected %v", tc.input, result, tc.expected)
			}
		})
	}
}
