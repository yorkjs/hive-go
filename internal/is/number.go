package is

func Integer(value float64) bool {
	return value == float64(int(value))
}
