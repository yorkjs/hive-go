package hive

func TruncateString(str string, maxLength int) string {
	// 将字符串转换为 rune 切片，以正确计算字符数
	runes := []rune(str)

	if len(runes) <= maxLength {
		return str
	}

	if maxLength <= 3 {
		return string(runes[:maxLength])
	}

	return string(runes[:maxLength-3]) + "..."
}
