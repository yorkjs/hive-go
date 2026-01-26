package hive

import "math"

// RateToDisplay 万分比 转换为 百分比
//
// value 后端的比例值
func RateToDisplay(value int) float64 {
	result := DivideNumber(float64(value), 100)
	// 如果小数部分为 0，返回整数部分
	if IsInteger(result) {
		return math.Floor(result)
	}
	return result
}

// RateToBackend 百分比 转换为 万分比
//
// value 前端的比例值
func RateToBackend(value float64) int {
	return int(TimesNumber(value, 100))
}

// CalculateRate 计算万分比
//
// value1 除数
// value2 被除数
func CalculateRate(value1, value2 float64) int {
	if value2 == 0 {
		return 0
	}
	return int(DivideNumber(value1*10000, value2))
}
