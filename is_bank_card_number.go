package hive

import (
	"regexp"
	"strconv"
)

// IsBankCardNumber 是否为银行卡号码
//
// value 要校验的值
// 返回是否为银行卡号码
func IsBankCardNumber(value string) bool {
	// 1. 基础验证：只能是数字
	match, _ := regexp.MatchString(`^\d+$`, value)
	if !match {
		return false
	}

	// 2. 长度验证：银行卡号长度通常在 13-19 位之间
	if len(value) < 13 || len(value) > 19 {
		return false
	}

	// 3. Luhn 算法验证
	return luhnCheck(value)
}

// luhnCheck Luhn 算法验证（模10算法）
//
// 算法步骤：
// 1. 从右向左，将每一位数字隔位乘以2
// 2. 如果乘以2后的数字大于9，则减去9（或相加各位数字）
// 3. 将所有数字相加
// 4. 如果总和能被10整除，则有效
//
// digits 数字字符串
// 返回是否通过校验
func luhnCheck(digits string) bool {
	var (
		sum    int
		isEven bool
	)

	// 从右向左遍历
	for i := len(digits) - 1; i >= 0; i-- {
		digit, err := strconv.Atoi(string(digits[i]))
		if err != nil {
			return false
		}

		// 隔位乘以2
		if isEven {
			digit *= 2
			// 如果大于9，减去9（等同于各位数字相加）
			if digit > 9 {
				digit -= 9
			}
		}

		sum += digit
		isEven = !isEven
	}

	return sum%10 == 0
}
