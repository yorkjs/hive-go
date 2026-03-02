package hive

import "testing"

func TestHexToRgbaString(t *testing.T) {
	tests := []struct {
		name  string
		color string
		alpha float64
		want  string
	}{
		{"红色简写", "#F00", 0.5, "rgba(255,0,0,0.5)"},
		{"红色完整", "#FF0000", 0.5, "rgba(255,0,0,0.5)"},
		{"绿色简写", "#0F0", 0.7, "rgba(0,255,0,0.7)"},
		{"绿色完整", "#00FF00", 0.7, "rgba(0,255,0,0.7)"},
		{"蓝色简写", "#00F", 0.9, "rgba(0,0,255,0.9)"},
		{"蓝色完整", "#0000FF", 0.9, "rgba(0,0,255,0.9)"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := HexToRgbaString(tt.color, tt.alpha)
			if err != nil {
				t.Errorf("HexToRgbaString() error = %v", err)
				return
			}
			if got != tt.want {
				t.Errorf("HexToRgbaString() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestHexToRgbaObject(t *testing.T) {
	tests := []struct {
		name  string
		color string
		alpha float64
		want  *IRgba
	}{
		{
			name:  "红色简写",
			color: "#F00",
			alpha: 0.5,
			want:  &IRgba{Red: 255, Green: 0, Blue: 0, Alpha: 0.5},
		},
		{
			name:  "红色完整",
			color: "#FF0000",
			alpha: 0.5,
			want:  &IRgba{Red: 255, Green: 0, Blue: 0, Alpha: 0.5},
		},
		{
			name:  "绿色简写",
			color: "#0F0",
			alpha: 0.7,
			want:  &IRgba{Red: 0, Green: 255, Blue: 0, Alpha: 0.7},
		},
		{
			name:  "绿色完整",
			color: "#00FF00",
			alpha: 0.7,
			want:  &IRgba{Red: 0, Green: 255, Blue: 0, Alpha: 0.7},
		},
		{
			name:  "蓝色简写",
			color: "#00F",
			alpha: 0.9,
			want:  &IRgba{Red: 0, Green: 0, Blue: 255, Alpha: 0.9},
		},
		{
			name:  "蓝色完整",
			color: "#0000FF",
			alpha: 0.9,
			want:  &IRgba{Red: 0, Green: 0, Blue: 255, Alpha: 0.9},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := HexToRgbaObject(tt.color, tt.alpha)
			if err != nil {
				t.Errorf("HexToRgbaObject() error = %v", err)
				return
			}
			compareRgba(t, got, tt.want)
		})
	}
}

// 辅助函数，用于比较 IRgba 结构体
func compareRgba(t *testing.T, got, want *IRgba) {
	t.Helper()
	if got.Red != want.Red {
		t.Errorf("Red: got %d, want %d", got.Red, want.Red)
	}
	if got.Green != want.Green {
		t.Errorf("Green: got %d, want %d", got.Green, want.Green)
	}
	if got.Blue != want.Blue {
		t.Errorf("Blue: got %d, want %d", got.Blue, want.Blue)
	}
	if got.Alpha != want.Alpha {
		t.Errorf("Alpha: got %f, want %f", got.Alpha, want.Alpha)
	}
}
