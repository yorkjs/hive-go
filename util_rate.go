package hive

// CalculateRate 计算万分比，即 value1 / value2 得到一个万分比
// value1: 除数
// value2: 被除数
// 返回万分比比例
func CalculateRate[T NumberType](value1, value2 T) int {
	if value2 == 0 {
		return 0
	}
	return int(DivideNumber(TimesNumber(value1, 10000), value2))
}

// ApplyRate 根据万分比比例计算数值
// value: 原始数值
// rate: 万分比比例
// 返回计算后的数值，仅返回整数部分
func ApplyRate[T NumberType](value, rate T) int {
	return int(DivideNumber(TimesNumber(value, rate), 10000))
}
