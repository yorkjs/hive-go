package hive

import (
	"crypto/rand"
	"math"
	"math/big"
)

// RandomIntegerByLength 生成指定长度的随机整数
func RandomIntegerByLength[T IntegerType](length int) T {
	min := T(math.Pow10(length - 1))
	max := T(math.Pow10(length)) - 1

	return RandomIntegerByRange(min, max+1)
}

// RandomIntegerByRange 生成指定范围内的随机整数
// 支持 int 和 int64 类型，返回类型与参数类型保持一致
func RandomIntegerByRange[T IntegerType](min, max T) T {
	if min >= max {
		return min
	}

	n, err := rand.Int(rand.Reader, big.NewInt(int64(max-min)))
	if err != nil {
		return min
	}

	return min + T(n.Int64())
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
