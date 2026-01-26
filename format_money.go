package hive

// FormatAmount 格式化金额（元），保留 2 位小数
func FormatAmount(value int, unit ...string) string {
	u := "元"
	if len(unit) > 0 {
		u = unit[0]
	}
	return FormatNumberWithComma(MoneyToDisplay(value), 2) + u
}

// FormatPenny 格式化金额（厘），保留 3 位小数
func FormatPenny(value int, unit ...string) string {
	u := "元"
	if len(unit) > 0 {
		u = unit[0]
	}
	return FormatNumberWithComma(MoneyToDisplay(value, MONEY_YUAN_TO_PENNY), 3) + u
}

// FormatAmountShortly 格式化金额（元），以较短的方式返回
func FormatAmountShortly(value int, unit ...string) string {
	u := "元"
	if len(unit) > 0 {
		u = unit[0]
	}
	return ShortNumber(
		MoneyToDisplay(value),
		func(v float64) string {
			return TruncateNumber(v, 2)
		},
	) + u
}
