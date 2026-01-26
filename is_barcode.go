package hive

import "regexp"

var (
	regexCustomBarcode = regexp.MustCompile(`^[A-Z]\d{10,12}$`)
	regexAllDigits     = regexp.MustCompile(`^\d+$`)
)

// IsStandardBarcode value 是否是标准商品条形码
//
// value 条形码文本
func IsStandardBarcode(value string) bool {
	// EAN 码：通常是69开头，13位数字
	// UPC 码：通常以0开头，12位数字
	length := len(value)
	if length == 8 || length == 12 || length == 13 || length == 14 {
		return regexAllDigits.MatchString(value)
	}
	return false
}

// IsCustomBarcode value 是否是自定义商品条形码
//
// value 条形码文本
func IsCustomBarcode(value string) bool {
	// 自定义条码，规则为 大写字母开头跟 10-12 个数字
	return regexCustomBarcode.MatchString(value)
}
