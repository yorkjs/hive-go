package normalize

import (
	"strings"
)

func Version(version string) string {

	tokens := strings.Split(version, ".")

	switch len(tokens) {
	case 1:
		// 单个数字，填充到12位
		return padLeft(tokens[0], 12, '0')
	case 2:
		// 两个数字，各填充到6位
		return padLeft(tokens[0], 6, '0') + padLeft(tokens[1], 6, '0')
	case 3:
		// 三个数字，各填充到4位
		return padLeft(tokens[0], 4, '0') + padLeft(tokens[1], 4, '0') + padLeft(tokens[2], 4, '0')
	}
	// 超过3个部分或其他情况
	return "000000000000"
}

// padLeft 在字符串左侧填充指定字符到指定长度
func padLeft(s string, totalLength int, padChar rune) string {
	if len(s) >= totalLength {
		return s
	}
	padding := strings.Repeat(string(padChar), totalLength-len(s))
	return padding + s
}
