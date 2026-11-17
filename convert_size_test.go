package hive

import (
	"math"
	"testing"
)

func almostEqual(got, want, delta float64) bool {
	return math.Abs(got-want) <= delta
}

func TestSizeConvert(t *testing.T) {
	// 正向：字节 → KB/MB/GB（容差 0.01，覆盖 2 位小数）
	toUnit := []struct {
		name string
		got  float64
		want float64
	}{
		{"SizeToKB(1024)", SizeToKB[int64](1024), 1},
		{"SizeToKB(2048)", SizeToKB[int64](2048), 2},
		{"SizeToKB(10240)", SizeToKB[int64](10240), 10},
		{"SizeToMB(1048576)", SizeToMB[int64](1048576), 1},
		{"SizeToMB(2097152)", SizeToMB[int64](2097152), 2},
		{"SizeToMB(10485760)", SizeToMB[int64](10485760), 10},
		{"SizeToGB(1073741824)", SizeToGB[int64](1073741824), 1},
		{"SizeToGB(2147483648)", SizeToGB[int64](2147483648), 2},
		{"SizeToGB(10737418240)", SizeToGB[int64](10737418240), 10},
	}
	for _, c := range toUnit {
		t.Run(c.name, func(t *testing.T) {
			if !almostEqual(c.got, c.want, 0.01) {
				t.Errorf("%s = %v, want %v", c.name, c.got, c.want)
			}
		})
	}

	// 反向：前端 → 后端字节（严格相等）
	toBackend := []struct {
		name string
		got  int64
		want int64
	}{
		{"SizeKBToBackend(1)", SizeKBToBackend[int64](1), 1024},
		{"SizeKBToBackend(10)", SizeKBToBackend[int64](10), 10240},
		{"SizeKBToBackend(100)", SizeKBToBackend[int64](100), 102400},
		{"SizeMBToBackend(1)", SizeMBToBackend[int64](1), 1048576},
		{"SizeMBToBackend(10)", SizeMBToBackend[int64](10), 10485760},
		{"SizeMBToBackend(100)", SizeMBToBackend[int64](100), 104857600},
		{"SizeGBToBackend(1)", SizeGBToBackend[int64](1), 1073741824},
		{"SizeGBToBackend(10)", SizeGBToBackend[int64](10), 10737418240},
		{"SizeGBToBackend(100)", SizeGBToBackend[int64](100), 107374182400},
	}
	for _, c := range toBackend {
		t.Run(c.name, func(t *testing.T) {
			if c.got != c.want {
				t.Errorf("%s = %v, want %v", c.name, c.got, c.want)
			}
		})
	}
}
