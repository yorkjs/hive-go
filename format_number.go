package hive

import (
	"fmt"
	"strings"
)

// FormatNumberWithComma 把数字的整数部分格式化为以千为段拆分，以逗号为分隔符
//
// value 数字
// decimals 保留几位小数
func FormatNumberWithComma(value float64, decimals ...int) string {
	d := 0
	if len(decimals) > 0 {
		d = decimals[0]
	}

	isNegative := value < 0
	if isNegative {
		value *= -1
	}

	var newValue string
	if d >= 0 {
		newValue = TruncateNumber(value, d)
	} else {
		newValue = fmt.Sprintf("%v", value)
	}

	parts := strings.Split(newValue, ".")
	integerPart := parts[0]
	var decimalPart string
	if len(parts) > 1 {
		decimalPart = parts[1]
	}

	var list []string
	end := len(integerPart) - 1
	for i := end; i >= 0; i-- {
		if i != end && (end-i)%3 == 0 {
			list = append(list, ",")
		}
		list = append(list, string(integerPart[i]))
	}

	// Reverse list
	for i, j := 0, len(list)-1; i < j; i, j = i+1, j-1 {
		list[i], list[j] = list[j], list[i]
	}
	result := strings.Join(list, "")

	if decimalPart != "" {
		if d > 0 {
			// padEnd(decimals, '0')
			if len(decimalPart) < d {
				decimalPart += strings.Repeat("0", d-len(decimalPart))
			} else if len(decimalPart) > d {
				decimalPart = decimalPart[:d]
			}
			result += "." + decimalPart
		}
	} else if d > 0 {
		result += "." + strings.Repeat("0", d)
	}

	if isNegative {
		return "-" + result
	}
	return result
}
