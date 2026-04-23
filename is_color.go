package hive

// IsHexColor value 是否是十六进制颜色
//
// value 十六进制颜色值
func IsHexColor(value string) bool {
	_, err := HexToRgbaObject(value)
	return err == nil
}
