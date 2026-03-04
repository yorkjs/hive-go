package hive

import (
	"encoding/hex"
	"fmt"
	"regexp"
	"strings"
)

// EncodeUriComponent 编码 URI 组件
func EncodeUriComponent(str string) string {
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

// DecodeUriComponent 解码 URI 组件
func DecodeUriComponent(str string) (string, error) {
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

var httpProtocolPattern = regexp.MustCompile(`^https?://`)

// NormalizeUrl 标准化 URL：确保包含协议部分
func NormalizeUrl(urlStr string) string {
	if urlStr == "" {
		return ""
	}

	if httpProtocolPattern.MatchString(urlStr) {
		return urlStr
	}

	if strings.HasPrefix(urlStr, "//") {
		return "https:" + urlStr
	}

	return "https://" + urlStr
}

// ToProtocolRelativeUrl 将 URL 转换为协议相对路径（以 // 开头）
func ToProtocolRelativeUrl(urlStr string) string {
	if urlStr == "" {
		return ""
	}

	if httpProtocolPattern.MatchString(urlStr) {
		return httpProtocolPattern.ReplaceAllString(urlStr, "//")
	}

	if !strings.HasPrefix(urlStr, "//") {
		return "//" + urlStr
	}

	return urlStr
}
