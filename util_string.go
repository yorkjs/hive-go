package hive

import (
	"crypto/rand"
	"fmt"
	"math/big"
	"regexp"
	"strings"
)

// GetStringLength 获取字符串长度
// 注意：中文算 1 个字符
func GetStringLength(str string) int {
	return len([]rune(str))
}

// TrimString 移除字符串开头和结尾的空白符
func TrimString(str string) string {
	return strings.TrimSpace(str)
}

// SliceString 截取字符串
func SliceString(str string, start, end int) string {
	runes := []rune(str)
	if start < 0 {
		start = 0
	}
	if end > len(runes) {
		end = len(runes)
	}
	if start > end {
		return ""
	}
	return string(runes[start:end])
}

// TruncateString 截断字符串，最多显示 maxLength 个字符，超过部分用省略号表示
func TruncateString(str string, maxLength int) string {
	runes := []rune(str)
	if len(runes) <= maxLength {
		return str
	}

	if maxLength <= 3 {
		return string(runes[:maxLength])
	}

	return string(runes[:maxLength-3]) + "..."
}

// RandomString 生成指定长度的随机字符串
func RandomString(length int, chars ...string) string {
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

// RenderStringTemplate 渲染字符串模板
// str: 字符串模板，例如：'你好，${name}'
// data: 数据对象，例如：map[string]interface{}{"name": "张三"}
// 返回渲染后的字符串，例如：'你好，张三'
func RenderStringTemplate(str string, data map[string]interface{}) string {
	// 编译正则表达式
	re := regexp.MustCompile(`\${(.*?)}`)
	return re.ReplaceAllStringFunc(str, func(match string) string {
		// 提取变量名（去掉 ${ 和 }）
		key := match[2 : len(match)-1]
		// 去除两端空白
		key = TrimString(key)

		// 获取对应的值
		value, ok := data[key]
		if value == nil || !ok {
			// 如果找不到对应的值，返回原字符串
			return match
		}

		// 将值转换为字符串
		return fmt.Sprintf("%v", value)
	})
}
