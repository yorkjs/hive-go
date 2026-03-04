package hive

import "fmt"

// FormatDiscount 把万分比格式化为折扣
func FormatDiscount[T IntegerType](value T) string {
	return fmt.Sprintf("%v折", DiscountToDisplay(value))
}
