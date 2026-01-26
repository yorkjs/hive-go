package hive

import "regexp"

var regexEmail = regexp.MustCompile(`^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}$`)

// IsEmail 是否为邮箱
func IsEmail(value string) bool {
	return regexEmail.MatchString(value)
}
