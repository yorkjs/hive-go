package hive

import "fmt"

// FormatDistance 把距离格式化为千米单位
func FormatDistance[T IntegerType](distance T) string {
	return fmt.Sprintf("%v公里", DistanceToDisplay(distance))
}
