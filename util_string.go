package hive

import (
	"encoding/hex"
	"fmt"
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

// EncodeURIComponent 编码 URI 组件
func EncodeURIComponent(str string) string {
	var result strings.Builder

	for _, r := range str {
		switch {
		// 字母和数字不编码
		case (r >= 'a' && r <= 'z') || (r >= 'A' && r <= 'Z') || (r >= '0' && r <= '9'):
			result.WriteRune(r)

		// encodeURIComponent 不编码的字符（根据 MDN）
		case r == '-' || r == '_' || r == '.' || r == '!' || r == '~' || r == '*' || r == '\'' || r == '(' || r == ')':
			result.WriteRune(r)

		// 空格需要编码为 %20
		case r == ' ':
			result.WriteString("%20")

		// 其他所有字符都需要编码
		default:
			// 将字符转换为 UTF-8 字节序列并编码
			bytes := []byte(string(r))
			for _, b := range bytes {
				result.WriteString(fmt.Sprintf("%%%02X", b))
			}
		}
	}

	return result.String()
}

// DecodeURIComponent 解码 URI 组件
func DecodeURIComponent(str string) (string, error) {
	var (
		result strings.Builder
		i      int
	)

	for i < len(str) {
		// 处理 %XX 编码
		if str[i] == '%' && i+2 < len(str) {
			// 尝试解析两个十六进制字符
			hexStr := str[i+1 : i+3]
			b, err := hex.DecodeString(hexStr)
			if err != nil {
				return "", err
			}

			result.Write(b)
			i += 3
		} else {
			// 普通字符直接写入
			result.WriteByte(str[i])
			i++
		}
	}

	return result.String(), nil
}
