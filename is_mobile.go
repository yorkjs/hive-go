package hive

// isMobile 是否为手机号码
func IsMobile(value string) bool {
	return ParsePhoneNumber(value) == PHONE_NUMBER_MOBILE
}
