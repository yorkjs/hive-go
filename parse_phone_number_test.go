package hive

import "testing"

func TestParsePhoneNumber(t *testing.T) {
	if got := ParsePhoneNumber("15812347890"); got != PHONE_NUMBER_MOBILE {
		t.Errorf("ParsePhoneNumber(15812347890) = %v; want PHONE_NUMBER_MOBILE", got)
	}
	if got := ParsePhoneNumber("01088427865"); got != PHONE_NUMBER_LANDLINE {
		t.Errorf("ParsePhoneNumber(01088427865) = %v; want PHONE_NUMBER_LANDLINE", got)
	}
	if got := ParsePhoneNumber("010-88427865"); got != PHONE_NUMBER_LANDLINE {
		t.Errorf("ParsePhoneNumber(010-88427865) = %v; want PHONE_NUMBER_LANDLINE", got)
	}
	if got := ParsePhoneNumber("091288427865"); got != PHONE_NUMBER_LANDLINE {
		t.Errorf("ParsePhoneNumber(091288427865) = %v; want PHONE_NUMBER_LANDLINE", got)
	}
	if got := ParsePhoneNumber("0912-87654321"); got != PHONE_NUMBER_LANDLINE {
		t.Errorf("ParsePhoneNumber(0912-87654321) = %v; want PHONE_NUMBER_LANDLINE", got)
	}
	if got := ParsePhoneNumber("4008997651"); got != PHONE_NUMBER_400 {
		t.Errorf("ParsePhoneNumber(4008997651) = %v; want PHONE_NUMBER_400", got)
	}
	if got := ParsePhoneNumber("400-899-7651"); got != PHONE_NUMBER_400 {
		t.Errorf("ParsePhoneNumber(400-899-7651) = %v; want PHONE_NUMBER_400", got)
	}
	if got := ParsePhoneNumber("40089976512"); got != -1 {
		t.Errorf("ParsePhoneNumber(40089976512) = %v; want -1", got)
	}
	if got := ParsePhoneNumber("C69241878121"); got != -1 {
		t.Errorf("ParsePhoneNumber(C69241878121) = %v; want -1", got)
	}
}
