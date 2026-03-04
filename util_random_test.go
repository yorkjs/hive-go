package hive

import (
	"fmt"
	"math"
	"regexp"
	"strings"
	"testing"
	"time"
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
	// 测试 int 类型
	for length := 2; length < 10; length++ {
		result := RandomIntegerByLength[int](length)
		if len(fmt.Sprintf("%d", result)) != length {
			t.Errorf("RandomIntegerByLength[int](%d) 长度 = %d, 期望 %d", length, len(fmt.Sprintf("%d", result)), length)
		}
	}

	// 测试 int64 类型
	for length := 2; length < 20; length++ {
		result := RandomIntegerByLength[int64](length)
		if len(fmt.Sprintf("%d", result)) != length {
			t.Errorf("RandomIntegerByLength[int64](%d) 长度 = %d, 期望 %d", length, len(fmt.Sprintf("%d", result)), length)
		}
	}
}

func TestRandomIntegerByRange(t *testing.T) {
	// 测试 int 类型
	t.Run("int type", func(t *testing.T) {
		minValue, maxValue := 10, 99
		for i := 0; i < 100; i++ {
			randomValue := RandomIntegerByRange(minValue, maxValue)
			if randomValue < minValue || randomValue > maxValue {
				t.Errorf("RandomIntegerByRange(%d, %d) 生成的整数 %d 不在范围 [%d, %d] 内", minValue, maxValue, randomValue, minValue, maxValue)
			}
		}
	})

	// 测试 int64 类型
	t.Run("int64 type", func(t *testing.T) {
		for length := 2; length < 20; length++ {
			minValue := int64(math.Pow10(length - 1))
			maxValue := int64(math.Pow10(length) - 1)
			randomValue := RandomIntegerByRange(minValue, maxValue)

			if randomValue < minValue || randomValue > maxValue {
				t.Errorf("RandomIntegerByRange(%d, %d) 生成的整数 %d 不在范围 [%d, %d] 内", minValue, maxValue, randomValue, minValue, maxValue)
			}
		}
	})
}

func TestRandomStringByCurrentTime(t *testing.T) {
	// 测试长度
	if len(RandomStringByCurrentTime(-1)) != 17 {
		t.Errorf("RandomStringByCurrentTime(-1) 长度应为 17, 实际为 %d", len(RandomStringByCurrentTime(-1)))
	}

	if len(RandomStringByCurrentTime(0)) != 17 {
		t.Errorf("RandomStringByCurrentTime(0) 长度应为 17, 实际为 %d", len(RandomStringByCurrentTime(0)))
	}

	if len(RandomStringByCurrentTime(3)) != 20 {
		t.Errorf("RandomStringByCurrentTime(3) 长度应为 20, 实际为 %d", len(RandomStringByCurrentTime(3)))
	}

	// 测试是否只包含数字
	result := RandomStringByCurrentTime(3)
	digitRegex := regexp.MustCompile(`^\d+$`)
	if !digitRegex.MatchString(result) {
		t.Errorf("RandomStringByCurrentTime(3) 应该只包含数字, 实际得到: %s", result)
	}

	// 测试是否以当前年份开头
	currentYear := time.Now().Format("2006")
	if !strings.HasPrefix(result, currentYear) {
		t.Errorf("RandomStringByCurrentTime(3) 应该以 %s 开头, 实际得到: %s", currentYear, result)
	}
}
