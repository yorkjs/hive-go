package hive

import "regexp"

var (
	regexMobile   = regexp.MustCompile(`^1\d{10}$`)
	regexLandline = regexp.MustCompile(`^(0\d{9,11}|0\d{2,3}-\d{7,8})$`)
	regex400      = regexp.MustCompile(`^(400\d{7}|400-\d{3}-\d{4})$`)
)

// ParsePhoneNumber 解析电话号码
func ParsePhoneNumber(value string) int {
	// 手机号码
	if regexMobile.MatchString(value) {
		return PHONE_NUMBER_MOBILE
	}

	// 固定电话
	if regexLandline.MatchString(value) {
		return PHONE_NUMBER_LANDLINE
	}

	// 400 电话
	if regex400.MatchString(value) {
		return PHONE_NUMBER_400
	}

	return -1
}
