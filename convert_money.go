package hive

// MoneyToDisplay 把金额转换为前端显示所用的格式
//
// value 后端的金额值，单位是分
func MoneyToDisplay[T IntegerType](value T, unit ...int) float64 {
	u := MONEY_YUAN_TO_CENT
	if len(unit) > 0 {
		u = unit[0]
	}
	return DivideNumber(float64(value), float64(u))
}

// MoneyToBackend 把金额转换为后端接口所用的格式
//
// value 前端的金额值
func MoneyToBackend[T IntegerType](value float64, unit ...int) T {
	u := MONEY_YUAN_TO_CENT
	if len(unit) > 0 {
		u = unit[0]
	}
	return T(TimesNumber(value, float64(u)))
}
