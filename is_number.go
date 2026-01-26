package hive

import "math"

// IsInteger 是否为整数
func IsInteger(value float64) bool {
	return math.Mod(value, 1) == 0
}
