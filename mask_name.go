package hive

// 脱敏姓名
func MaskName(name string) string {
	nameRune := []rune(name)
	length := len(nameRune)
	if length <= 1 {
		return "***"
	}
	return "***" + string(nameRune[length-1:length])
}
