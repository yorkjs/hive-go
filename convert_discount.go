package hive

import (
	"math"
	"strconv"
)

// DiscountToDisplay 万分比 转换为 折扣，最多保留 1 位小数
//
// value 后端的比例值
func DiscountToDisplay[T IntegerType](value T) float64 {
	result := DivideNumber(float64(value), 1000)
	// 如果小数部分为 0，返回整数部分
	// 如果有小数，保留 1 位小数
	if IsInteger(result) {
		return math.Floor(result)
	}
	result, _ = strconv.ParseFloat(TruncateNumber(result, 1), 64)
	return result
}

// DiscountToBackend 折扣 转换为 万分比
//
// value 前端的比例值
func DiscountToBackend[T IntegerType](value float64) T {
	var v float64
	if IsInteger(value) {
		v = value
	} else {
		v, _ = strconv.ParseFloat(TruncateNumber(value, 1), 64)
	}
	return T(TimesNumber(v, 1000))
}
