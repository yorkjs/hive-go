package hive

import (
	"math"
)

// DiscountToDisplay 万分比 转换为 折扣，最多保留 1 位小数
//
// value 后端的比例值
func DiscountToDisplay(value int) float64 {
	result := DivideNumber(float64(value), 1000)
	// 如果小数部分为 0，返回整数部分
	// 如果有小数，保留 1 位小数
	if IsInteger(result) {
		return math.Floor(result)
	}
	return stringToFloat64(
		TruncateNumber(result, 1),
	)
}

// DiscountToBackend 折扣 转换为 万分比
//
// value 前端的比例值
func DiscountToBackend(value float64) int {
	var v float64
	if IsInteger(value) {
		v = value
	} else {
		v = stringToFloat64(
			TruncateNumber(value, 1),
		)
	}
	return int(TimesNumber(v, 1000))
}
