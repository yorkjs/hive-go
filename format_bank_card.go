package hive

import (
	"math"
	"strings"
)

// FormatBankCardNumber 格式化银行卡号
//
// value
// masked 是否脱敏显示 (default true)
func FormatBankCardNumber(value string, masked ...bool) string {
	isMasked := true
	if len(masked) > 0 {
		isMasked = masked[0]
	}

	length := len(value)
	var parts []string
	for i := 0; i < length; i += 4 {
		end := int(math.Min(float64(i+4), float64(length)))
		parts = append(parts, value[i:end])
	}

	if isMasked {
		for i := 0; i < len(parts)-1; i++ {
			parts[i] = "****"
		}
	}

	return strings.Join(parts, " ")
}
