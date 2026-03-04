package hive

import (
	"strconv"
	"strings"

	"github.com/shopspring/decimal"
)

// PlusNumber 精确加法，比如 PlusNumber(3, 1) === 4
func PlusNumber[T NumberType](value1, value2 T) T {
	d1 := decimal.NewFromFloat(float64(value1))
	d2 := decimal.NewFromFloat(float64(value2))
	f, _ := d1.Add(d2).Float64()
	return T(f)
}

// MinusNumber 精确减法，比如 MinusNumber(3, 1) === 2
func MinusNumber[T NumberType](value1, value2 T) T {
	d1 := decimal.NewFromFloat(float64(value1))
	d2 := decimal.NewFromFloat(float64(value2))
	f, _ := d1.Sub(d2).Float64()
	return T(f)
}

// TimesNumber 精确乘法，比如 TimesNumber(3, 2) === 6
func TimesNumber[T NumberType](value1, value2 T) T {
	d1 := decimal.NewFromFloat(float64(value1))
	d2 := decimal.NewFromFloat(float64(value2))
	f, _ := d1.Mul(d2).Float64()
	return T(f)
}

// DivideNumber 精确除法，比如 DivideNumber(6, 2) === 3
func DivideNumber[T NumberType](value1, value2 T) T {
	if value2 == 0 {
		panic("Divisor cannot be zero")
	}
	d1 := decimal.NewFromFloat(float64(value1))
	d2 := decimal.NewFromFloat(float64(value2))
	f, _ := d1.Div(d2).Float64()
	return T(f)
}

// TruncateNumber 截断数字，解决 1.983.toFixed(1) 为 2.0 的问题
func TruncateNumber(value float64, decimals int) string {
	s := strconv.FormatFloat(value, 'f', -1, 64)
	parts := strings.Split(s, ".")
	integerPart := parts[0]
	var decimalPart string
	if len(parts) > 1 {
		decimalPart = parts[1]
	}

	if decimals == 0 {
		return integerPart
	}
	if decimalPart == "" {
		return integerPart + "." + strings.Repeat("0", decimals)
	}

	var truncatedDecimal string
	if len(decimalPart) > decimals {
		truncatedDecimal = decimalPart[:decimals]
	} else {
		truncatedDecimal = decimalPart + strings.Repeat("0", decimals-len(decimalPart))
	}

	return integerPart + "." + truncatedDecimal
}

// ShortNumber 以较短的方式返回数字，避免 UI 层显示不下所有数字
func ShortNumber[T NumberType](value T, formatUnshort func(T) string) string {
	if value >= 1000000000000 {
		trillion := DivideNumber(float64(value), 1000000000000)
		decimals := 1
		if IsInteger(trillion) {
			decimals = 0
		}
		return TruncateNumber(trillion, decimals) + "万亿"
	}
	if value >= 100000000 {
		billion := DivideNumber(float64(value), 100000000)
		decimals := 1
		if IsInteger(billion) {
			decimals = 0
		}
		return TruncateNumber(billion, decimals) + "亿"
	}
	if value >= 10000 {
		tenThousand := DivideNumber(float64(value), 10000)
		decimals := 1
		if IsInteger(tenThousand) {
			decimals = 0
		}
		return TruncateNumber(tenThousand, decimals) + "万"
	}
	return formatUnshort(value)
}

func stringToFloat64(str string) float64 {
	f, _ := strconv.ParseFloat(str, 64)
	return f
}
