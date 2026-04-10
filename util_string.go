package hive

import (
	"fmt"
	"regexp"
	"strings"
	"unicode/utf8"
)

// GetStringLength 获取字符串字符数量
//
// 注意：中文和英文都算 1 个字符
func GetStringLength(str string) int {
	return utf8.RuneCountInString(str)
}

// GetStringWidth 获取字符串宽度，此函数常用于排版辅助计算
//
// 注意：中文算 2 个单位，英文数字算 1 个单位
func GetStringWidth(str string) int {
	if str == "" {
		return 0
	}

	var (
		wideCount int
		length    int
	)

	for i := 0; i < len(str); {
		r, size := utf8.DecodeRuneInString(str[i:])
		if r == utf8.RuneError {
			// 处理无效 UTF-8 字符
			i++
			continue
		}

		if isWideChar(r) {
			wideCount++
		}
		length++
		i += size
	}

	return wideCount*2 + (length - wideCount)
}

// isWideChar 判断字符是否为宽字符（中文字符、全角标点等）
func isWideChar(r rune) bool {
	// 判断是否为非 ASCII 字符（大于 0xFF）
	// 或者是否为全角片假名字符范围（FF61-FF9F）
	return r > 0xFF || (r >= 0xFF61 && r <= 0xFF9F)
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

func PadStringStart(str string, length int) string {
	if len(str) >= length {
		return str
	}
	return strings.Repeat("0", length-len(str)) + str
}

var (
	// 允许的字符模式：字母、数字、中文、常见标点符号
	// \u4e00-\u9fa5 匹配所有中文字符
	specialCharacterPattern = regexp.MustCompile(`[^ \x{4e00}-\x{9fa5}a-zA-Z0-9，。、；：！π“”‘’（）【】《》？～·—…\.,;:!?"'()\[\]{}<>@#&%￥$_\+/*-]`)
)

// HasSpecialCharacter 判断字符串是否包含特殊字符
func HasSpecialCharacter(str string) bool {
	if str == "" {
		return false
	}

	// 查找是否包含不允许的字符
	return specialCharacterPattern.MatchString(str)
}

// RemoveSpecialCharacter 移除字符串中的特殊字符
func RemoveSpecialCharacter(str string) string {
	if str == "" {
		return ""
	}

	// 移除所有不允许的字符
	return specialCharacterPattern.ReplaceAllString(str, "")
}
