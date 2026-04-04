package hive

import "testing"

// TestIsLocationInChina 单元测试
func TestIsLocationInChina(t *testing.T) {
	tests := []struct {
		name      string
		longitude float64
		latitude  float64
		want      bool
	}{
		{
			name:      "北京坐标应该返回true",
			longitude: 116.4074,
			latitude:  39.9042,
			want:      true,
		},
		{
			name:      "上海坐标应该返回true",
			longitude: 121.4874,
			latitude:  31.2242,
			want:      true,
		},
		{
			name:      "经度超出范围应该返回false",
			longitude: 160.4874,
			latitude:  31.2242,
			want:      false,
		},
		{
			name:      "纬度超出范围应该返回false",
			longitude: 121.4874,
			latitude:  80,
			want:      false,
		},
		{
			name:      "原点坐标应该返回false",
			longitude: 0,
			latitude:  0,
			want:      false,
		},
		{
			name:      "负坐标应该返回false",
			longitude: -120,
			latitude:  -30,
			want:      false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := IsLocationInChina(tt.longitude, tt.latitude); got != tt.want {
				t.Errorf("IsLocationInChina(%v, %v) = %v, want %v",
					tt.longitude, tt.latitude, got, tt.want)
			}
		})
	}
}
