package hive

import "fmt"

// FormatDiscount 把万分比格式化为折扣
func FormatDiscount(value int) string {
	return fmt.Sprintf("%v折", DiscountToDisplay(value))
}
