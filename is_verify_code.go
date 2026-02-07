package hive

import (
	"regexp"
)

// IsVerifyCode 是否为验证码
func IsVerifyCode(value string) bool {
	regexVerifyCode := regexp.MustCompile(`^\d{6}$`)
	return regexVerifyCode.MatchString(value)
}
