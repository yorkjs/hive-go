package is

import (
	"regexp"
)

// 是否是标准商品条形码
func StandardBarcode(value string) bool {
	// EAN 码：通常是69开头，13位数字
	// UPC 码：通常以0开头，12位数字
	length := len(value)
	if length == 8 || length == 12 || length == 13 || length == 14 {
		re := regexp.MustCompile(`^\d+$`)
		return re.MatchString(value)
	}
	return false
}

// 检查是否是自定义商品条形码
func CustomBarcode(value string) bool {
	// 自定义条码，规则为 大写字母开头跟 10-12 个数字
	re := regexp.MustCompile(`^[A-Z]\d{10,12}$`)
	return re.MatchString(value)
}
