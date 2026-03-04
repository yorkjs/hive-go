package hive

// IntegerType 定义支持的整数类型约束
type IntegerType interface {
	~int | ~int64
}

// NumberType 定义支持的数字类型约束
type NumberType interface {
	~int | ~int64 | ~float64
}
