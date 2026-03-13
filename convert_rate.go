package hive

import "math"

// RateToDisplay 万分比 转换为 百分比
//
// value 后端的比例值
func RateToDisplay[T IntegerType](value T) float64 {
	result := DivideNumber(float64(value), 100)
	// 如果小数部分为 0，返回整数部分
	if HasDecimal(result) {
		return result
	}
	return math.Floor(result)
}

// RateToBackend 百分比 转换为 万分比
//
// value 前端的比例值
func RateToBackend[T NumberType](value T) int {
	return int(TimesNumber(value, 100))
}
