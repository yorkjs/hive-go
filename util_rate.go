package hive

import "math"

// CalculateRate 计算万分比，即 value1 / value2 得到一个万分比
// value1: 除数
// value2: 被除数
// 返回万分比比例
func CalculateRate[T NumberType](value1, value2 T) int {
	if value2 == 0 {
		return 0
	}
	return int(DivideNumber(float64(TimesNumber(value1, 10000)), float64(value2)))
}

// ApplyRateFloor 根据万分比计算数值，策略是向下取整
// value: 原始数值
// rate: 万分比比例
func ApplyRateFloor[T NumberType](value T, rate int) int64 {
	return int64(DivideNumber(float64(TimesNumber(value, T(rate))), 10000))
}

// ApplyRateCeil 根据万分比计算数值，策略是向上取整
// value: 原始数值
// rate: 万分比比例
func ApplyRateCeil[T NumberType](value T, rate int) int64 {
	return int64(math.Ceil(DivideNumber(float64(TimesNumber(value, T(rate))), 10000)))
}
