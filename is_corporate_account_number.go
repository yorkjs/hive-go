package hive

import "regexp"

// IsCorporateAccountNumber 是否为对公账户号码
// value: 对公账户号码
// 返回: 是否为对公账户
func IsCorporateAccountNumber(value string) bool {
	matched, _ := regexp.MatchString(`^\d{9,32}$`, value)
	return matched
}
