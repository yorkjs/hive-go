package hive

import "testing"

func TestFormatNumberWithComma(t *testing.T) {
	tests := []struct {
		input    float64
		decimals []int // Use slice to handle optional argument
		want     string
	}{
		{input: -228.80, decimals: []int{2}, want: "-228.80"},
		{input: 1000000, decimals: nil, want: "1,000,000"},
		{input: 1000000.981, decimals: []int{2}, want: "1,000,000.98"},
		{input: 10000, decimals: []int{1}, want: "10,000.0"},
		{input: 10000, decimals: []int{0}, want: "10,000"},
		{input: 10, decimals: []int{2}, want: "10.00"},
		{input: 1234567.89, decimals: nil, want: "1,234,567"},
		{input: 1234567.89, decimals: []int{0}, want: "1,234,567"},
		{input: 1234567.89, decimals: []int{3}, want: "1,234,567.890"},
		{input: 1234567.89, decimals: []int{5}, want: "1,234,567.89000"},
		{input: -1234.56, decimals: []int{2}, want: "-1,234.56"},
	}

	for _, tt := range tests {
		var got string
		if tt.decimals != nil {
			got = FormatNumberWithComma(tt.input, tt.decimals[0])
		} else {
			got = FormatNumberWithComma(tt.input)
		}
		if got != tt.want {
			t.Errorf("FormatNumberWithComma(%v, %v) = %v; want %v", tt.input, tt.decimals, got, tt.want)
		}
	}
}
