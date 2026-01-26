package hive

import "testing"

func TestFormatAmount(t *testing.T) {
	tests := []struct {
		input int
		want  string
	}{
		{1000, "10.00元"},
		{1230, "12.30元"},
		{1234, "12.34元"},
		{123456, "1,234.56元"}, // Original Go test case
	}

	for _, tt := range tests {
		if got := FormatAmount(tt.input); got != tt.want {
			t.Errorf("FormatAmount(%v) = %v; want %v", tt.input, got, tt.want)
		}
	}

	if got := FormatAmount(123456, "美元"); got != "1,234.56美元" {
		t.Errorf("FormatAmount(123456, \"美元\") = %v; want 1,234.56美元", got)
	}
}

func TestFormatPenny(t *testing.T) {
	tests := []struct {
		input int
		want  string
	}{
		{80, "0.080元"},
		{81, "0.081元"},
		{181, "0.181元"},
		{1234560, "1,234.560元"}, // Original Go test case
	}

	for _, tt := range tests {
		if got := FormatPenny(tt.input); got != tt.want {
			t.Errorf("FormatPenny(%v) = %v; want %v", tt.input, got, tt.want)
		}
	}
}

func TestFormatAmountShortly(t *testing.T) {
	tests := []struct {
		input int
		unit  []string
		want  string
	}{
		{100000000000000, nil, "1万亿元"},
		{110000000000000, nil, "1.1万亿元"},
		{111000000000000, nil, "1.1万亿元"},
		{100000000000, nil, "10亿元"},
		{10000000000, nil, "1亿元"},
		{11000000000, nil, "1.1亿元"},
		{11100000000, nil, "1.1亿元"},
		{10000000, nil, "10万元"},
		{1000000, nil, "1万元"},
		{1100000, nil, "1.1万元"},
		{1110000, nil, "1.1万元"},
		{1000, nil, "10.00元"},
		{1000, []string{""}, "10.00"},
		{1234500, nil, "1.2万元"},   // Original Go test case
		{999900, nil, "9999.00元"}, // Original Go test case
	}

	for _, tt := range tests {
		var got string
		if tt.unit != nil {
			got = FormatAmountShortly(tt.input, tt.unit...)
		} else {
			got = FormatAmountShortly(tt.input)
		}
		if got != tt.want {
			t.Errorf("FormatAmountShortly(%v) = %v; want %v", tt.input, got, tt.want)
		}
	}
}
