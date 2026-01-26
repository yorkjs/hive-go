package hive

import "regexp"

var regexPrice = regexp.MustCompile(`^(?:[1-9]\d*|0)(?:\.\d{1,2})?$`)

// IsPrice 是否为价格
func IsPrice(value string) bool {
	return regexPrice.MatchString(value)
}
