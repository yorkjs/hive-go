package hive

import (
	"math"
)

// DistanceToDisplay 米 转换为 千米
//
// value 后端的比例值
func DistanceToDisplay[T IntegerType](value T) float64 {
	result := DivideNumber(float64(value), 1000)
	// 如果小数部分为 0，返回整数部分
	if HasDecimal(result) {
		return result
	}
	return math.Floor(result)
}

// DistanceToBackend 千米 转换为 米
//
// value 前端的比例值
func DistanceToBackend[T IntegerType](value float64) T {
	return T(TimesNumber(value, 1000))
}
