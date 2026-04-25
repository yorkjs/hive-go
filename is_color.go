package hive

import "strings"

// IsHexColor value 是否是十六进制颜色
//
// value 十六进制颜色值
func IsHexColor(value string) bool {
	if !strings.HasPrefix(value, "#") {
		return false
	}
	_, err := HexToRgbaObject(value)
	return err == nil
}
