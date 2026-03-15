package hive

import "math"

// 定义地球半径（单位：米）
const earthRadiusM = 6371 * 1000

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
func CalculateDistance[T IntegerType](longitude1, latitude1, longitude2, latitude2 float64) T {
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
	return T(earthRadiusM * c)
}
