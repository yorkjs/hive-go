package hive

import "testing"

func TestIsIdentityCardNumber(t *testing.T) {
	if got := IsIdentityCardNumber(""); got != false {
		t.Errorf(`IsIdentityCardNumber("") = %v; want false`, got)
	}
	if got := IsIdentityCardNumber("1"); got != false {
		t.Errorf(`IsIdentityCardNumber("1") = %v; want false`, got)
	}
	if got := IsIdentityCardNumber("12"); got != false {
		t.Errorf(`IsIdentityCardNumber("12") = %v; want false`, got)
	}
	if got := IsIdentityCardNumber("464561561561651"); got != false {
		t.Errorf(`IsIdentityCardNumber("464561561561651") = %v; want false`, got)
	}
}
