package hive

import "testing"

func TestIsBankCardNumber(t *testing.T) {
	if got := IsBankCardNumber(""); got != false {
		t.Errorf(`IsBankCardNumber("") = %v; want false`, got)
	}
	if got := IsBankCardNumber("1"); got != false {
		t.Errorf(`IsBankCardNumber("1") = %v; want false`, got)
	}
	if got := IsBankCardNumber("12"); got != false {
		t.Errorf(`IsBankCardNumber("12") = %v; want false`, got)
	}
	if got := IsBankCardNumber("6228480012345678"); got != false {
		t.Errorf(`IsBankCardNumber("6228480012345678") = %v; want false`, got)
	}
	if got := IsBankCardNumber("1234567890123"); got != false {
		t.Errorf(`IsBankCardNumber("1234567890123") = %v; want false`, got)
	}
}
