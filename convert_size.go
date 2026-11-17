package hive

// SizeToKB 字节 转换为 KB
//
// value 后端的 size 值，单位字节
func SizeToKB[T IntegerType](value T) float64 {
	return DivideNumber(float64(value), SIZE_KB)
}

// SizeToMB 字节 转换为 MB
//
// value 后端的 size 值，单位字节
func SizeToMB[T IntegerType](value T) float64 {
	return DivideNumber(float64(value), SIZE_MB)
}

// SizeToGB 字节 转换为 GB
//
// value 后端的 size 值，单位字节
func SizeToGB[T IntegerType](value T) float64 {
	return DivideNumber(float64(value), SIZE_GB)
}

// SizeKBToBackend KB 转为后端使用的 字节
//
// value 前端的 size 值，单位是 KB
func SizeKBToBackend[T IntegerType](value float64) T {
	return T(TimesNumber(value, SIZE_KB))
}

// SizeMBToBackend MB 转为后端使用的 字节
//
// value 前端的 size 值，单位是 MB
func SizeMBToBackend[T IntegerType](value float64) T {
	return T(TimesNumber(value, SIZE_MB))
}

// SizeGBToBackend GB 转为后端使用的 字节
//
// value 前端的 size 值，单位是 GB
func SizeGBToBackend[T IntegerType](value float64) T {
	return T(TimesNumber(value, SIZE_GB))
}
