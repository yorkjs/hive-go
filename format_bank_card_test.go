package hive

import "testing"

func TestFormatBankCardNumber(t *testing.T) {
	if got := FormatBankCardNumber("9879764467646788"); got != "**** **** **** 6788" {
		t.Errorf("FormatBankCardNumber(9879764467646788) = %q; want \"**** **** **** 6788\"", got)
	}
	if got := FormatBankCardNumber("98797644676467881"); got != "**** **** **** **** 1" {
		t.Errorf("FormatBankCardNumber(98797644676467881) = %q; want \"**** **** **** **** 1\"", got)
	}
	if got := FormatBankCardNumber("987976446764678812"); got != "**** **** **** **** 12" {
		t.Errorf("FormatBankCardNumber(987976446764678812) = %q; want \"**** **** **** **** 12\"", got)
	}
	if got := FormatBankCardNumber("9879764467646788123"); got != "**** **** **** **** 123" {
		t.Errorf("FormatBankCardNumber(9879764467646788123) = %q; want \"**** **** **** **** 123\"", got)
	}
	if got := FormatBankCardNumber("98797644676467881234"); got != "**** **** **** **** 1234" {
		t.Errorf("FormatBankCardNumber(98797644676467881234) = %q; want \"**** **** **** **** 1234\"", got)
	}
	if got := FormatBankCardNumber("987976446764678812345"); got != "**** **** **** **** **** 5" {
		t.Errorf("FormatBankCardNumber(987976446764678812345) = %q; want \"**** **** **** **** **** 5\"", got)
	}

	if got := FormatBankCardNumber("9879764467646788", false); got != "9879 7644 6764 6788" {
		t.Errorf("FormatBankCardNumber(..., false) = %q; want \"9879 7644 6764 6788\"", got)
	}
	if got := FormatBankCardNumber("98797644676467881", false); got != "9879 7644 6764 6788 1" {
		t.Errorf("FormatBankCardNumber(..., false) = %q; want \"9879 7644 6764 6788 1\"", got)
	}
	if got := FormatBankCardNumber("987976446764678812", false); got != "9879 7644 6764 6788 12" {
		t.Errorf("FormatBankCardNumber(..., false) = %q; want \"9879 7644 6764 6788 12\"", got)
	}
	if got := FormatBankCardNumber("9879764467646788123", false); got != "9879 7644 6764 6788 123" {
		t.Errorf("FormatBankCardNumber(..., false) = %q; want \"9879 7644 6764 6788 123\"", got)
	}
	if got := FormatBankCardNumber("98797644676467881234", false); got != "9879 7644 6764 6788 1234" {
		t.Errorf("FormatBankCardNumber(..., false) = %q; want \"9879 7644 6764 6788 1234\"", got)
	}
	if got := FormatBankCardNumber("987976446764678812345", false); got != "9879 7644 6764 6788 1234 5" {
		t.Errorf("FormatBankCardNumber(..., false) = %q; want \"9879 7644 6764 6788 1234 5\"", got)
	}
}
