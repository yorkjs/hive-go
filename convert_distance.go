package hive

import (
	"math"
)

// DistanceToDisplay 米 转换为 千米
//
// value 后端的比例值
func DistanceToDisplay(value int) float64 {
	result := DivideNumber(float64(value), 1000)
	// 如果小数部分为 0，返回整数部分
	if IsInteger(result) {
		return math.Floor(result)
	}
	return result
}

// DistanceToBackend 千米 转换为 米
//
// value 前端的比例值
func DistanceToBackend(value float64) int {
	return int(TimesNumber(value, 1000))
}

// 定义地球半径（单位：米）
const EARTH_RADIUS_M = 6371 * 1000

// 将角度转换为弧度
func toRadians(degrees float64) float64 {
	return degrees * math.Pi / 180
}

// CalculateDistance 计算两个点之间的距离，返回距离单位是米
//
// longitude1 第一个点的经度
// latitude1 第一个点的纬度
// longitude2 第二个点的经度
// latitude2 第二个点的纬度
func CalculateDistance(longitude1, latitude1, longitude2, latitude2 float64) int {
	// 将经纬度转换为弧度
	lat1 := toRadians(latitude1)
	lon1 := toRadians(longitude1)
	lat2 := toRadians(latitude2)
	lon2 := toRadians(longitude2)

	// 计算差值
	dLat := lat2 - lat1
	dLon := lon2 - lon1

	// Haversine 公式
	a := math.Pow(math.Sin(dLat/2), 2) +
		math.Cos(lat1)*math.Cos(lat2)*
			math.Pow(math.Sin(dLon/2), 2)
	c := 2 * math.Atan2(math.Sqrt(a), math.Sqrt(1-a))

	// 计算距离
	return int(EARTH_RADIUS_M * c)
}
