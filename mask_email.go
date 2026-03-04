package hive

import "strings"

// 脱敏邮箱
func MaskEmail(email string) string {
	// 查找 @ 符号
	atIndex := strings.Index(email, "@")
	if atIndex <= 0 {
		// 没有 @ 符号或 @ 在开头，不是有效邮箱，原样返回
		return email
	}

	// 分割用户名和域名
	username := email[:atIndex]
	domain := email[atIndex:]

	switch {
	case len(username) <= 1:
		return "***" + domain
	default:
		return "***" + username[len(username)-1:] + domain
	}
}
