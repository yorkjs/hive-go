package hive

// 这些常量需要你在实际使用中定义
// const SHELF_LIFE_YEAR = 365 * 24 // 假设一年是365天的总小时数
// const SHELF_LIFE_MONTH = 30 * 24 // 假设一月是30天的总小时数
// const SHELF_LIFE_DAY = 24

type IShelfLife struct {
	Years  int
	Months int
	Days   int
	Hours  int
}

func NormalizeShelfLife(value int) *IShelfLife {
	result := &IShelfLife{}

	if value <= 0 {
		return result
	}

	remainingValue := value

	years := remainingValue / SHELF_LIFE_YEAR
	if years > 0 {
		result.Years = years
		remainingValue -= SHELF_LIFE_YEAR * years
	}

	months := remainingValue / SHELF_LIFE_MONTH
	if months > 0 {
		result.Months = months
		remainingValue -= SHELF_LIFE_MONTH * months
	}

	days := remainingValue / SHELF_LIFE_DAY
	if days > 0 {
		result.Days = days
		remainingValue -= SHELF_LIFE_DAY * days
	}

	if remainingValue > 0 {
		result.Hours = remainingValue
	}

	return result
}
