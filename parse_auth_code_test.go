package hive

import "testing"

func TestParseAuthCode(t *testing.T) {
	if got := ParseAuthCode("133619858964803511"); got != AUTH_CODE_WECHAT {
		t.Errorf("ParseAuthCode(133619858964803511) = %v; want AUTH_CODE_WECHAT", got)
	}
	if got := ParseAuthCode("283654147086344711"); got != AUTH_CODE_ALIPAY {
		t.Errorf("ParseAuthCode(283654147086344711) = %v; want AUTH_CODE_ALIPAY", got)
	}
	if got := ParseAuthCode("583654147086344711"); got != -1 {
		t.Errorf("ParseAuthCode(583654147086344711) = %v; want -1", got)
	}
	if got := ParseAuthCode("C69241878121"); got != -1 {
		t.Errorf("ParseAuthCode(C69241878121) = %v; want -1", got)
	}
}
