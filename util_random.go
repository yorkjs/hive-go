package hive

import (
	"crypto/rand"
	"math"
	"math/big"
)

// RandomIntegerByLength 生成指定长度的随机整数
func RandomIntegerByLength(length int) int64 {
	min := int64(math.Pow10(length - 1))
	max := int64(math.Pow10(length)) - 1

	return RandomIntegerByRange(min, max+1)
}

// RandomIntegerByRange 生成指定范围内的随机整数
func RandomIntegerByRange(min, max int64) int64 {
	if min > max {
		return min
	}

	if min == max {
		return min
	}

	n, err := rand.Int(rand.Reader, big.NewInt(max-min))
	if err != nil {
		return min
	}

	return min + n.Int64()
}

// RandomStringByLength 生成指定长度的随机字符串
func RandomStringByLength(length int, chars ...string) string {
	charSet := "ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz0123456789"
	if len(chars) > 0 && chars[0] != "" {
		charSet = chars[0]
	}

	result := make([]byte, length)
	charLength := big.NewInt(int64(len(charSet)))

	for i := range length {
		// 使用 crypto/rand 生成安全的随机数
		randomIndex, err := rand.Int(rand.Reader, charLength)
		if err != nil {
			// 降级使用简单的随机数（仅用于演示，实际应用中应处理错误）
			randomIndex = big.NewInt(int64(i % len(charSet)))
		}
		result[i] = charSet[randomIndex.Int64()]
	}

	return string(result)
}
