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
		return PadStringStart(tokens[0], 12)
	case 2:
		return PadStringStart(tokens[0], 6) + PadStringStart(tokens[1], 6)
	case 3:
		return PadStringStart(tokens[0], 4) + PadStringStart(tokens[1], 4) + PadStringStart(tokens[2], 4)
	}
	return "000000000000"
}
