package hive

import "regexp"

var (
	regexWechatStart = regexp.MustCompile(`^1[0-5]`)
	regexWechatLen   = regexp.MustCompile(`^\d{18}$`)
	regexAlipayStart = regexp.MustCompile(`^(2[5-9]|30)`)
	regexAlipayLen   = regexp.MustCompile(`^\d{18,25}$`)
)

// ParseAuthCode 解析付款码
func ParseAuthCode(value string) int {
	// 微信支付通常以 10-15 开头、18 位纯数字
	if regexWechatStart.MatchString(value) && regexWechatLen.MatchString(value) {
		return AUTH_CODE_WECHAT
	}

	// 支付宝通常以 25-30 开头、18-25 位纯数字
	if regexAlipayStart.MatchString(value) && regexAlipayLen.MatchString(value) {
		return AUTH_CODE_ALIPAY
	}

	return -1
}
