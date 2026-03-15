package hive

import (
	"fmt"
	"math"
	"strconv"
	"strings"
)

// RGBA 颜色对象
type IRgba struct {
	Red   int
	Green int
	Blue  int
	Alpha float64
}

// HSL 颜色对象
type IHsl struct {
	Hue        float64
	Saturation float64
	Lightness  float64
}

// HexToRgbaObject 将 HEX 颜色转换为 RGBA 对象
// color: HEX 颜色值
// 返回 RGBA 颜色对象
func HexToRgbaObject(color string) (*IRgba, error) {
	// 移除 # 号
	hex := strings.ReplaceAll(color, "#", "")

	// 处理简写格式 (#rgb 或 #rgba)
	if len(hex) == 3 || len(hex) == 4 {
		newHex := ""
		for _, char := range hex {
			newHex += string(char) + string(char)
		}
		hex = newHex
	}

	// 验证hex长度
	if len(hex) != 6 && len(hex) != 8 {
		return nil, fmt.Errorf("无效的HEX颜色格式: %s", color)
	}

	// 解析RGB值

	red, err := ParseInteger(hex[0:2], 16)
	if err != nil {
		return nil, err
	}
	green, err := ParseInteger(hex[2:4], 16)
	if err != nil {
		return nil, err
	}
	blue, err := ParseInteger(hex[4:6], 16)
	if err != nil {
		return nil, err
	}

	result := &IRgba{
		Red:   int(red),
		Green: int(green),
		Blue:  int(blue),
		Alpha: 1,
	}

	// 解析透明度
	if len(hex) == 8 {
		alpha, err := ParseInteger(hex[6:8], 16)
		if err != nil {
			return nil, err
		}
		result.Alpha = float64(alpha) / 255
	}

	return result, nil
}

// HexToHslObject 将 HEX 颜色转换为 HSL 对象
// color: HEX 颜色值
// 返回 HSL 颜色对象
func HexToHslObject(color string) (*IHsl, error) {
	rgba, err := HexToRgbaObject(color)
	if err != nil {
		return nil, err
	}
	return RgbToHsl(rgba), nil
}

// HexToRgbaString 将 HEX 颜色转换为 RGBA 格式
// color: HEX 颜色值
// alpha: 透明度，取值范围 0-1
// 返回 RGBA 颜色字符串
func HexToRgbaString(color string, alpha float64) (string, error) {
	rgba, err := HexToRgbaObject(color)
	if err != nil {
		return "", err
	}
	return fmt.Sprintf("rgba(%d,%d,%d,%g)", rgba.Red, rgba.Green, rgba.Blue, alpha), nil
}

// DarkenColor 加深颜色亮度
// color: HEX 颜色值
// offset: 加深幅度，取值范围 0-1
// 返回新的 hex 颜色
func DarkenColor(color string, offset float64) (string, error) {
	return adjustColorBrightness(color, -offset)
}

// LightenColor 减淡颜色亮度
// color: HEX 颜色值
// offset: 减淡幅度，取值范围 0-1
// 返回新的 hex 颜色
func LightenColor(color string, offset float64) (string, error) {
	return adjustColorBrightness(color, offset)
}

// adjustColorBrightness 调整颜色亮度
// hex: 原始颜色
// offset: 取值范围 0-1
// 返回新的 hex 颜色字符串
func adjustColorBrightness(hex string, offset float64) (string, error) {
	rgba, err := HexToRgbaObject(hex)
	if err != nil {
		return "", err
	}
	hsl := RgbToHsl(rgba)

	// 调整亮度，限制在 0-100 之间
	newL := hsl.Lightness + (offset * 100)
	hsl.Lightness = math.Max(0, math.Min(100, newL))

	newRgb := HslToRgb(hsl)

	// 如果原颜色有透明度，返回值保留该透明度
	result := fmt.Sprintf("#%s%s%s",
		toHex(newRgb.Red),
		toHex(newRgb.Green),
		toHex(newRgb.Blue))

	if rgba.Alpha < 1 {
		result += toHex(int(rgba.Alpha * 255))
	}
	return result, nil
}

// RgbToHsl 将 RGB 转换为 HSL
// r, g, b: 0-255
// 返回 h: 0-360, s: 0-100, l: 0-100
func RgbToHsl(rgb *IRgba) *IHsl {
	r := float64(rgb.Red) / 255.0
	g := float64(rgb.Green) / 255.0
	b := float64(rgb.Blue) / 255.0

	max := math.Max(r, math.Max(g, b))
	min := math.Min(r, math.Min(g, b))
	h := 0.0
	s := 0.0
	l := (max + min) / 2.0

	if max != min {
		d := max - min
		if l > 0.5 {
			s = d / (2.0 - max - min)
		} else {
			s = d / (max + min)
		}

		switch {
		case max == r:
			h = (g - b) / d
			if g < b {
				h += 6.0
			}
		case max == g:
			h = (b-r)/d + 2.0
		case max == b:
			h = (r-g)/d + 4.0
		}
		h *= 60.0
	}

	return &IHsl{
		Hue:        h,
		Saturation: s * 100.0,
		Lightness:  l * 100.0,
	}
}

// HslToRgb 将 HSL 转换为 RGB
// 返回 r, g, b: 0-255
func HslToRgb(hsl *IHsl) *IRgba {
	h := hsl.Hue / 360.0
	s := hsl.Saturation / 100.0
	l := hsl.Lightness / 100.0

	var q float64
	if l < 0.5 {
		q = l * (1.0 + s)
	} else {
		q = l + s - l*s
	}
	p := 2.0*l - q

	r := hue2rgb(p, q, h+1.0/3.0)
	g := hue2rgb(p, q, h)
	b := hue2rgb(p, q, h-1.0/3.0)

	return &IRgba{
		Red:   int(math.Round(r * 255.0)),
		Green: int(math.Round(g * 255.0)),
		Blue:  int(math.Round(b * 255.0)),
		Alpha: 1,
	}
}

// hue2rgb 辅助函数：将色调转换为RGB
func hue2rgb(p, q, t float64) float64 {
	tt := t
	if tt < 0 {
		tt += 1.0
	}
	if tt > 1.0 {
		tt -= 1.0
	}
	if tt < 1.0/6.0 {
		return p + (q-p)*6.0*tt
	}
	if tt < 1.0/2.0 {
		return q
	}
	if tt < 2.0/3.0 {
		return p + (q-p)*(2.0/3.0-tt)*6.0
	}
	return p
}

// toHex 将颜色值转换为十六进制字符串
func toHex(color int) string {
	return PadStringStart(strings.ToUpper(strconv.FormatInt(int64(color), 16)), 2)
}
