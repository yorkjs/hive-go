package hive

// MoneyToDisplay 把金额转换为前端显示所用的格式
//
// value 后端的金额值，单位是分
func MoneyToDisplay(value int, unit ...int) float64 {
	u := MONEY_YUAN_TO_CENT
	if len(unit) > 0 {
		u = unit[0]
	}
	return DivideNumber(float64(value), float64(u))
}

// MoneyToBackend 把金额转换为后端接口所用的格式
//
// value 前端的金额值
func MoneyToBackend(value float64, unit ...int) int {
	u := MONEY_YUAN_TO_CENT
	if len(unit) > 0 {
		u = unit[0]
	}
	return int(TimesNumber(value, float64(u)))
}
