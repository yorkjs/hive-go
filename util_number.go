package hive

import (
	"regexp"
	"strconv"
	"strings"
	"unicode"

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

// parseInteger 解析字符串中的整型
func ParseInteger(str string, radix ...int) (int64, error) {
	// 处理可选参数radix
	var r *int
	if len(radix) > 0 {
		r = &radix[0]
	}

	// 步骤1：去除首尾空格
	str = strings.TrimSpace(str)
	if str == "" {
		return 0, strconv.ErrSyntax
	}

	// 步骤2：提取符号（+/-）和数字主体
	sign := ""
	numStr := str
	if strings.HasPrefix(str, "-") || strings.HasPrefix(str, "+") {
		sign = string(str[0])
		numStr = str[1:]
		if numStr == "" {
			return 0, strconv.ErrSyntax
		}
	}

	// 步骤3：提取有效数字前缀（直到第一个非数字字符为止）
	validCharRegex := regexp.MustCompile(`^[0-9a-zA-Z]+`)
	numMatch := validCharRegex.FindString(numStr)
	if numMatch == "" {
		return 0, strconv.ErrSyntax
	}
	rawNumPart := numMatch

	// 步骤4：处理radix的默认值和自动推断逻辑
	finalRadix := 10
	if r == nil || *r == 0 {
		if strings.HasPrefix(strings.ToLower(rawNumPart), "0x") {
			finalRadix = 16
		} else {
			finalRadix = 10
		}
	} else {
		finalRadix = *r
	}

	// 步骤5：校验radix的合法性（2-36之间）
	if finalRadix < 2 || finalRadix > 36 {
		return 0, strconv.ErrSyntax
	}

	// 步骤6：处理16进制前缀（0x/0X）
	actualNumPart := rawNumPart
	if finalRadix == 16 && strings.HasPrefix(strings.ToLower(rawNumPart), "0x") {
		actualNumPart = rawNumPart[2:]
		if actualNumPart == "" {
			return 0, strconv.ErrSyntax
		}
	}

	// 步骤7：解析数字（处理非法字符，只解析到第一个非法字符前）
	var validNumPart string
	for _, char := range actualNumPart {
		// 转小写
		lowerChar := unicode.ToLower(char)
		var charValue int

		// 转换字符为数值（0-9→0-9，a-z→10-35）
		if lowerChar >= '0' && lowerChar <= '9' {
			charValue = int(lowerChar - '0')
		} else if lowerChar >= 'a' && lowerChar <= 'z' {
			charValue = 10 + int(lowerChar-'a')
		} else {
			break // 非法字符，停止解析
		}

		// 字符数值超过进制范围，停止解析
		if charValue >= finalRadix {
			break
		}

		validNumPart += string(lowerChar)
	}

	// 无有效数字部分
	if validNumPart == "" {
		return 0, strconv.ErrSyntax
	}

	// 步骤8：解析最终的数字并加上符号
	result, err := strconv.ParseInt(validNumPart, finalRadix, 64)
	if err != nil {
		return 0, err
	}

	// 处理符号
	if sign == "-" {
		result = -result
	}

	return result, nil
}

// ParseNumber 解析字符串中的浮点数
func ParseNumber(str string) (float64, error) {
	// 去除首尾空格
	str = strings.TrimSpace(str)
	if str == "" {
		return 0, strconv.ErrSyntax
	}

	// 匹配浮点数字符串（支持正负、整数、小数）
	regex := regexp.MustCompile(`^[+-]?\d*\.?\d+`)
	match := regex.FindString(str)
	if match == "" {
		return 0, strconv.ErrSyntax
	}

	// 解析浮点数
	result, err := strconv.ParseFloat(match, 64)
	if err != nil {
		return 0, err
	}

	return result, nil
}
