package hive

import (
	"math"
	"testing"
)

func TestHexToRgbaString(t *testing.T) {
	tests := []struct {
		name     string
		color    string
		alpha    float64
		expected string
	}{
		{"简写红色", "#F00", 0.5, "rgba(255,0,0,0.5)"},
		{"完整红色", "#FF0000", 0.5, "rgba(255,0,0,0.5)"},
		{"简写绿色", "#0F0", 0.7, "rgba(0,255,0,0.7)"},
		{"完整绿色", "#00FF00", 0.7, "rgba(0,255,0,0.7)"},
		{"简写蓝色", "#00F", 0.9, "rgba(0,0,255,0.9)"},
		{"完整蓝色", "#0000FF", 0.9, "rgba(0,0,255,0.9)"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result, err := HexToRgbaString(tt.color, tt.alpha)
			if err != nil {
				t.Errorf("HexToRgbaString() error = %v", err)
				return
			}
			if result != tt.expected {
				t.Errorf("HexToRgbaString() = %v, want %v", result, tt.expected)
			}
		})
	}
}

func TestHexToRgbaObject(t *testing.T) {
	tests := []struct {
		name     string
		color    string
		expected *IRgba
	}{
		{"简写红色", "#F00", &IRgba{Red: 255, Green: 0, Blue: 0, Alpha: 1.0}},
		{"简写红色带透明", "#F000", &IRgba{Red: 255, Green: 0, Blue: 0, Alpha: 0.0}},
		{"简写红色带透明(FF)", "#F00F", &IRgba{Red: 255, Green: 0, Blue: 0, Alpha: 1.0}},
		{"完整红色", "#FF0000", &IRgba{Red: 255, Green: 0, Blue: 0, Alpha: 1.0}},
		{"完整红色带透明00", "#FF000000", &IRgba{Red: 255, Green: 0, Blue: 0, Alpha: 0.0}},
		{"完整红色带透明FF", "#FF0000FF", &IRgba{Red: 255, Green: 0, Blue: 0, Alpha: 1.0}},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result, err := HexToRgbaObject(tt.color)
			if err != nil {
				t.Errorf("HexToRgbaObject() error = %v", err)
				return
			}

			if result.Red != tt.expected.Red ||
				result.Green != tt.expected.Green ||
				result.Blue != tt.expected.Blue ||
				math.Abs(result.Alpha-tt.expected.Alpha) > 0.001 {
				t.Errorf("HexToRgbaObject() = %+v, want %+v", result, tt.expected)
			}
		})
	}
}

func TestDarkenColor(t *testing.T) {
	tests := []struct {
		name     string
		color    string
		offset   float64
		expected string
	}{
		{"加深红色", "#ff0000", 0.2, "#990000"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result, err := DarkenColor(tt.color, tt.offset)
			if err != nil {
				t.Errorf("DarkenColor() error = %v", err)
				return
			}
			if result != tt.expected {
				t.Errorf("DarkenColor() = %v, want %v", result, tt.expected)
			}
		})
	}
}

func TestLightenColor(t *testing.T) {
	tests := []struct {
		name     string
		color    string
		offset   float64
		expected string
	}{
		{"减淡红色", "#ff0000", 0.2, "#ff6666"},
		{"减淡红色带透明00", "#ff000000", 0.2, "#ff666600"},
		{"减淡红色带透明FF", "#ff0000FF", 0.2, "#ff6666"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result, err := LightenColor(tt.color, tt.offset)
			if err != nil {
				t.Errorf("LightenColor() error = %v", err)
				return
			}
			if result != tt.expected {
				t.Errorf("LightenColor() = %v, want %v", result, tt.expected)
			}
		})
	}
}
