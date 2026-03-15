package hive

// 此文件的函数仅在内部使用，不对外暴露

// shortNumber 以较短的方式返回数字，避免 UI 层显示不下所有数字
func shortNumber[T NumberType](value T, formatUnshort func(T) string) string {
	if value >= 1000000000000 {
		trillion := DivideNumber(float64(value), 1000000000000)
		decimals := 0
		if HasDecimal(trillion) {
			decimals = 1
		}
		return TruncateNumber(trillion, decimals) + "万亿"
	}
	if value >= 100000000 {
		billion := DivideNumber(float64(value), 100000000)
		decimals := 0
		if HasDecimal(billion) {
			decimals = 1
		}
		return TruncateNumber(billion, decimals) + "亿"
	}
	if value >= 10000 {
		tenThousand := DivideNumber(float64(value), 10000)
		decimals := 0
		if HasDecimal(tenThousand) {
			decimals = 1
		}
		return TruncateNumber(tenThousand, decimals) + "万"
	}
	return formatUnshort(value)
}
