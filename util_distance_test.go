package hive

import "testing"

func TestCalculateDistance(t *testing.T) {
	if got := CalculateDistance[int](
		116.4074, // 北京经度
		39.9042,  // 北京纬度
		121.4737, // 上海经度
		31.2304,  // 上海纬度
	); got != 1067310 {
		t.Errorf("CalculateDistance(...) = %v; want 1067310", got)
	}
}
