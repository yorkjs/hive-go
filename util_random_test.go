package hive

import (
	"fmt"
	"math"
	"testing"
)

// TestRandomStringByLength 测试随机字符串生成
func TestRandomStringByLength(t *testing.T) {
	for length := 5; length < 100; length++ {
		if len(RandomStringByLength(length)) != length {
			t.Errorf("RandomStringByLength(%d) 长度 = %d, 期望 %d", length, len(RandomStringByLength(length)), length)
		}
		if RandomStringByLength(length) == RandomStringByLength(length) {
			t.Errorf("RandomStringByLength(%d) 生成相同的字符串: %q", length, RandomStringByLength(length))
		}
	}
}

func TestRandomIntegerByLength(t *testing.T) {
	for length := 2; length < 20; length++ {
		if len(fmt.Sprintf("%d", RandomIntegerByLength(length))) != length {
			t.Errorf("RandomIntegerByLength(%d) 长度 = %d, 期望 %d", length, len(fmt.Sprintf("%d", RandomIntegerByLength(length))), length)
		}
		if RandomIntegerByLength(length) == RandomIntegerByLength(length) {
			t.Errorf("RandomIntegerByLength(%d) 生成相同的整数: %d", length, RandomIntegerByLength(length))
		}
	}
}

func TestRandomIntegerByRange(t *testing.T) {
	for length := 2; length < 20; length++ {
		minValue := int64(math.Pow10(length - 1))
		maxValue := int64(math.Pow10(length) - 1)
		randomValue := RandomIntegerByRange(minValue, maxValue)

		if randomValue < minValue || randomValue >= maxValue {
			t.Errorf("RandomIntegerByRange(%d) 生成的整数 %d 不在范围 [%d, %d) 内", length, randomValue, minValue, maxValue)
		}
	}
}
