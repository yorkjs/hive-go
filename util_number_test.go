package hive

import (
	"math"
	"testing"
)

func TestPlusNumber(t *testing.T) {
	if got := PlusNumber(1, 2); got != 3 {
		t.Errorf("PlusNumber(1, 2) = %v; want 3", got)
	}
	if got := PlusNumber(1.1, 2.1); got != 3.2 {
		t.Errorf("PlusNumber(1.1, 2.1) = %v; want 3.2", got)
	}
	if got := PlusNumber(1.12, 2.13); got != 3.25 {
		t.Errorf("PlusNumber(1.12, 2.13) = %v; want 3.25", got)
	}
	if got := PlusNumber(1.121, 2.131); got != 3.252 {
		t.Errorf("PlusNumber(1.121, 2.131) = %v; want 3.252", got)
	}
	if got := PlusNumber(1.121321, 2.3312213); got != 3.4525423 {
		t.Errorf("PlusNumber(1.121321, 2.3312213) = %v; want 3.4525423", got)
	}
	if got := PlusNumber(1, -2); got != -1 {
		t.Errorf("PlusNumber(1, -2) = %v; want -1", got)
	}
	if got := PlusNumber(-2, -3); got != -5 {
		t.Errorf("PlusNumber(-2, -3) = %v; want -5", got)
	}
}

func TestMinusNumber(t *testing.T) {
	if got := MinusNumber(2, 1); got != 1 {
		t.Errorf("MinusNumber(2, 1) = %v; want 1", got)
	}
	if got := MinusNumber(2.1, 1.1); got != 1 {
		t.Errorf("MinusNumber(2.1, 1.1) = %v; want 1", got)
	}
	if got := MinusNumber(2.32, 1.04); got != 1.28 {
		t.Errorf("MinusNumber(2.32, 1.04) = %v; want 1.28", got)
	}
	if got := MinusNumber(4.978672, 2.131231); got != 2.847441 {
		t.Errorf("MinusNumber(4.978672, 2.131231) = %v; want 2.847441", got)
	}
	if got := MinusNumber(3, -1); got != 4 {
		t.Errorf("MinusNumber(3, -1) = %v; want 4", got)
	}
	if got := MinusNumber(-2, -3); got != 1 {
		t.Errorf("MinusNumber(-2, -3) = %v; want 1", got)
	}
}

func TestTimesNumber(t *testing.T) {
	if got := TimesNumber(2, 1); got != 2 {
		t.Errorf("TimesNumber(2, 1) = %v; want 2", got)
	}
	if got := TimesNumber(2, 3); got != 6 {
		t.Errorf("TimesNumber(2, 3) = %v; want 6", got)
	}
	if got := TimesNumber(2.1, 10); got != 21 {
		t.Errorf("TimesNumber(2.1, 10) = %v; want 21", got)
	}
	if got := TimesNumber(2.11, 10); got != 21.1 {
		t.Errorf("TimesNumber(2.11, 10) = %v; want 21.1", got)
	}
	if got := TimesNumber(2.111, 10); got != 21.11 {
		t.Errorf("TimesNumber(2.111, 10) = %v; want 21.11", got)
	}
	if got := TimesNumber(2.111, 100); got != 211.1 {
		t.Errorf("TimesNumber(2.111, 100) = %v; want 211.1", got)
	}
	if got := TimesNumber(2.111, 1000); got != 2111 {
		t.Errorf("TimesNumber(2.111, 1000) = %v; want 2111", got)
	}
	if got := TimesNumber(2.11728132706, 2); got != 4.23456265412 {
		t.Errorf("TimesNumber(2.11728132706, 2) = %v; want 4.23456265412", got)
	}
	if got := TimesNumber(2.01, 100); got != 201 {
		t.Errorf("TimesNumber(2.01, 100) = %v; want 201", got)
	}
	if got := TimesNumber(2.1, 100); got != 210 {
		t.Errorf("TimesNumber(2.1, 100) = %v; want 210", got)
	}
	if got := TimesNumber(10, 0.5); got != 5 {
		t.Errorf("TimesNumber(10, 0.5) = %v; want 5", got)
	}
	if got := TimesNumber(10, 0.2); got != 2 {
		t.Errorf("TimesNumber(10, 0.2) = %v; want 2", got)
	}
	if got := TimesNumber(3, -2); got != -6 {
		t.Errorf("TimesNumber(3, -2) = %v; want -6", got)
	}
	if got := TimesNumber(-2, -3); got != 6 {
		t.Errorf("TimesNumber(-2, -3) = %v; want 6", got)
	}
	if got := TimesNumber(-2.1, -3.2); got != 6.72 {
		t.Errorf("TimesNumber(-2.1, -3.2) = %v; want 6.72", got)
	}
}

func TestDivideNumber(t *testing.T) {
	if got := DivideNumber(4, 2); got != 2 {
		t.Errorf("DivideNumber(4, 2) = %v; want 2", got)
	}
	if got := DivideNumber(6, 3); got != 2 {
		t.Errorf("DivideNumber(6, 3) = %v; want 2", got)
	}
	if got := DivideNumber(2.2, 2); got != 1.1 {
		t.Errorf("DivideNumber(2.2, 2) = %v; want 1.1", got)
	}
	if got := DivideNumber(4.234, 2); got != 2.117 {
		t.Errorf("DivideNumber(4.234, 2) = %v; want 2.117", got)
	}
	if got := DivideNumber(4.234562, 2); got != 2.117281 {
		t.Errorf("DivideNumber(4.234562, 2) = %v; want 2.117281", got)
	}
	if got := DivideNumber(4.23456265412, 2); got != 2.11728132706 {
		t.Errorf("DivideNumber(4.23456265412, 2) = %v; want 2.11728132706", got)
	}
	if got := DivideNumber(50, 10); got != 5 {
		t.Errorf("DivideNumber(50, 10) = %v; want 5", got)
	}
	if got := DivideNumber[float64](50, 100); got != 0.5 {
		t.Errorf("DivideNumber(50, 100) = %v; want 0.5", got)
	}
	if got := DivideNumber[float64](50, 1000); got != 0.05 {
		t.Errorf("DivideNumber(50, 1000) = %v; want 0.05", got)
	}
	if got := DivideNumber(10, 0.5); got != 20 {
		t.Errorf("DivideNumber(10, 0.5) = %v; want 20", got)
	}
	if got := DivideNumber(10, 0.2); got != 50 {
		t.Errorf("DivideNumber(10, 0.2) = %v; want 50", got)
	}
	if got := DivideNumber[float64](201, 100); got != 2.01 {
		t.Errorf("DivideNumber(201, 100) = %v; want 2.01", got)
	}
	if got := DivideNumber[float64](168, 100); got != 1.68 {
		t.Errorf("DivideNumber(168, 100) = %v; want 1.68", got)
	}
	if got := DivideNumber(5000000, 100); got != 50000 {
		t.Errorf("DivideNumber(5000000, 100) = %v; want 50000", got)
	}
	if got := DivideNumber[float64](3, -2); got != -1.5 {
		t.Errorf("DivideNumber(3, -2) = %v; want -1.5", got)
	}
	if got := DivideNumber(-6, -3); got != 2 {
		t.Errorf("DivideNumber(-6, -3) = %v; want 2", got)
	}
	if got := DivideNumber(-4.4, -2.2); got != 2 {
		t.Errorf("DivideNumber(-4.4, -2.2) = %v; want 2", got)
	}
}

func TestTruncateNumber(t *testing.T) {
	if got := TruncateNumber(1, 0); got != "1" {
		t.Errorf("TruncateNumber(1) = %q; want \"1\"", got)
	}
	if got := TruncateNumber(1.98321, 0); got != "1" {
		t.Errorf("TruncateNumber(1.98321) = %q; want \"1\"", got)
	}
	if got := TruncateNumber(1.98321, 3); got != "1.983" {
		t.Errorf("TruncateNumber(1.98321, 3) = %q; want \"1.983\"", got)
	}
	if got := TruncateNumber(1.98321, 2); got != "1.98" {
		t.Errorf("TruncateNumber(1.98321, 2) = %q; want \"1.98\"", got)
	}
	if got := TruncateNumber(1.98321, 1); got != "1.9" {
		t.Errorf("TruncateNumber(1.98321, 1) = %q; want \"1.9\"", got)
	}
	if got := TruncateNumber(1.98321, 0); got != "1" {
		t.Errorf("TruncateNumber(1.98321, 0) = %q; want \"1\"", got)
	}
	if got := TruncateNumber(1234567.89, 2); got != "1234567.89" {
		t.Errorf("TruncateNumber(1234567.89, 2) = %q; want \"1234567.89\"", got)
	}
	if got := TruncateNumber(1234567.89, 1); got != "1234567.8" {
		t.Errorf("TruncateNumber(1234567.89, 1) = %q; want \"1234567.8\"", got)
	}
	if got := TruncateNumber(1234567.89, 3); got != "1234567.890" {
		t.Errorf("TruncateNumber(1234567.89, 3) = %q; want \"1234567.890\"", got)
	}
}

// 测试 ParseInteger 函数
func TestParseInteger(t *testing.T) {
	testCases := []struct {
		name    string
		input   string
		radix   []int // 可选参数
		want    interface{}
		wantErr bool
	}{
		// 不传 radix 参数（默认行为）
		{"默认-空格前缀", " 123", nil, int64(123), false},
		{"默认-空格后缀", "123 ", nil, int64(123), false},
		{"默认-空格前后", " 123 ", nil, int64(123), false},
		{"默认-正常数字", "123", nil, int64(123), false},
		{"默认-负数", "-123", nil, int64(-123), false},
		{"默认-正数", "+123", nil, int64(123), false},
		{"默认-数字带单位", "123px", nil, int64(123), false},
		{"默认-单位在前", "px123", nil, int64(0), true},
		{"默认-浮点数", "123.456", nil, int64(123), false},
		{"默认-浮点数带单位", "123.456px", nil, int64(123), false},
		{"默认-二进制字符串", "101", nil, int64(101), false},
		{"默认-二进制字符串", "101", []int{2}, int64(5), false},
		{"默认-二进制带单位", "101a", nil, int64(101), false},
		{"默认-二进制带单位", "101a", []int{2}, int64(5), false},
		{"默认-十六进制0XFF", "0XFF", nil, int64(255), false},
		{"默认-十六进制0xFF", "0xFF", []int{16}, int64(255), false},

		{"默认-十六进制0XFF", "0xFF", nil, int64(255), false},
		{"默认-十六进制0xFF", "0xFF", []int{16}, int64(255), false},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			var got int64
			var err error

			// 根据是否传 radix 参数调用函数
			if tc.radix == nil {
				got, err = ParseInteger(tc.input)
			} else {
				got, err = ParseInteger(tc.input, tc.radix[0])
			}

			// 检查错误
			if tc.wantErr {
				if err == nil {
					t.Errorf("ParseInteger(%q, %v) expected error but got nil, result=%v",
						tc.input, tc.radix, got)
				}
				return
			}

			if err != nil {
				t.Errorf("ParseInteger(%q, %v) unexpected error: %v",
					tc.input, tc.radix, err)
				return
			}

			// 检查结果
			if got != tc.want {
				t.Errorf("ParseInteger(%q, %v) = %v, want %v",
					tc.input, tc.radix, got, tc.want)
			}
		})
	}
}

// 测试 ParseNumber 函数
func TestParseNumber(t *testing.T) {
	testCases := []struct {
		name    string
		input   string
		want    float64
		wantErr bool
	}{
		{"空格前缀", " 123", 123, false},
		{"空格后缀", "123 ", 123, false},
		{"空格前后", " 123 ", 123, false},
		{"正常整数", "123", 123, false},
		{"负数", "-123", -123, false},
		{"正数", "+123", 123, false},
		{"整数带单位", "123px", 123, false},
		{"单位在前", "px123", 0, true},
		{"浮点数", "123.456", 123.456, false},
		{"浮点数带单位", "123.456px", 123.456, false},
		{"十六进制0xFF", "0xFF", 0, false},
		{"十六进制0XFF", "0XFF", 0, false},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			got, err := ParseNumber(tc.input)

			// 检查错误
			if tc.wantErr {
				if err == nil {
					t.Errorf("ParseNumber(%q) expected error but got nil, result=%v",
						tc.input, got)
				}
				return
			}

			if err != nil {
				t.Errorf("ParseNumber(%q) unexpected error: %v", tc.input, err)
				return
			}

			// 检查结果（浮点数比较需要考虑精度）
			if math.Abs(got-tc.want) > 1e-9 {
				t.Errorf("ParseNumber(%q) = %v, want %v", tc.input, got, tc.want)
			}
		})
	}
}

func TestHasDecimal(t *testing.T) {
	if got := HasDecimal(1); got != false {
		t.Errorf("HasDecimal(1) = %v; want false", got)
	}
	if got := HasDecimal(1.0); got != false {
		t.Errorf("HasDecimal(1.0) = %v; want false", got)
	}
	if got := HasDecimal(1.1); got != true {
		t.Errorf("HasDecimal(1.1) = %v; want true", got)
	}
}
