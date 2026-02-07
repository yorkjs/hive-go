package hive

import "net/url"

// IsUrl 是否为 URL
func IsUrl(value string) bool {
	u, err := url.Parse(value)
	if err != nil {
		return false
	}
	return u.Scheme == "http" || u.Scheme == "https"
}
