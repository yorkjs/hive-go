package hive

import (
	"strings"
)

// NormalizeVersion 标准化版本号，方便后续进行比较
//
// version 如 '1.2.3'，最多支持三段，每段的子版本号不超过 9999
// returns 标准化后的版本号，12 位长度
func NormalizeVersion(version string) string {
	tokens := strings.Split(version, ".")
	switch len(tokens) {
	case 1:
		return padStart(tokens[0], 12, '0')
	case 2:
		return padStart(tokens[0], 6, '0') + padStart(tokens[1], 6, '0')
	case 3:
		return padStart(tokens[0], 4, '0') + padStart(tokens[1], 4, '0') + padStart(tokens[2], 4, '0')
	}
	return "000000000000"
}

func padStart(s string, length int, padChar byte) string {
	if len(s) >= length {
		return s
	}
	return strings.Repeat(string(padChar), length-len(s)) + s
}
