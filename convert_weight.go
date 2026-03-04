package hive

// WeightToG 毫克 转换为 克
//
// value 后端的重量值，单位毫克
func WeightToG[T IntegerType](value T) float64 {
	return DivideNumber(float64(value), 1000)
}

// WeightToKG 毫克 转换为 千克
//
// value 后端的重量值，单位毫克
func WeightToKG[T IntegerType](value T) float64 {
	return DivideNumber(float64(value), 1000000)
}

// WeightGToBackend 克 转为后端使用的 毫克
//
// value 前端的重量值，单位是克
func WeightGToBackend[T IntegerType](value float64) T {
	return T(TimesNumber(value, 1000))
}

// WeightKGToBackend 千克 转为后端使用的 毫克
//
// value 前端的重量值，单位是千克
func WeightKGToBackend[T IntegerType](value float64) T {
	return T(TimesNumber(value, 1000000))
}
