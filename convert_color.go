package hive

import (
	"fmt"
	"strconv"
	"strings"
)

// Rgba 表示 RGBA 颜色
type IRgba struct {
	Red   int
	Green int
	Blue  int
	Alpha float64
}

// HexToRgbaObject 将 HEX 颜色转换为 RGBA 格式
// color: HEX 颜色值
// alpha: 透明度，取值范围 0-1
// 返回 RGBA 颜色对象和可能的错误
func HexToRgbaObject(color string, alpha float64) (*IRgba, error) {
	// 移除 # 号
	hex := strings.TrimPrefix(color, "#")

	// 处理简写格式 (#rgb 或 #rgba)
	if len(hex) == 3 || len(hex) == 4 {
		var expanded strings.Builder
		for _, char := range hex {
			expanded.WriteString(string(char) + string(char))
		}
		hex = expanded.String()
	}

	// 验证hex长度
	if len(hex) != 6 && len(hex) != 8 {
		return nil, fmt.Errorf("无效的HEX颜色格式: %s", color)
	}

	// 解析RGB值
	red, err := strconv.ParseInt(hex[0:2], 16, 64)
	if err != nil {
		return nil, fmt.Errorf("解析红色值失败: %v", err)
	}

	green, err := strconv.ParseInt(hex[2:4], 16, 64)
	if err != nil {
		return nil, fmt.Errorf("解析绿色值失败: %v", err)
	}

	blue, err := strconv.ParseInt(hex[4:6], 16, 64)
	if err != nil {
		return nil, fmt.Errorf("解析蓝色值失败: %v", err)
	}

	// 解析透明度值
	if len(hex) == 8 {
		alphaVal, err := strconv.ParseInt(hex[6:8], 16, 64)
		if err != nil {
			return nil, fmt.Errorf("解析透明度值失败: %v", err)
		}
		alpha = float64(alphaVal) / 255
	}

	// 限制透明度范围
	if alpha < 0 {
		alpha = 0
	} else if alpha > 1 {
		alpha = 1
	}

	return &IRgba{
		Red:   int(red),
		Green: int(green),
		Blue:  int(blue),
		Alpha: alpha,
	}, nil
}

// HexToRgbaString 将 HEX 颜色转换为 RGBA 格式的字符串
// color: HEX 颜色值
// alpha: 透明度，取值范围 0-1
// 返回 RGBA 颜色字符串和可能的错误
func HexToRgbaString(color string, alpha float64) (string, error) {
	rgba, err := HexToRgbaObject(color, alpha)
	if err != nil {
		return "", err
	}
	return fmt.Sprintf("rgba(%d,%d,%d,%g)", rgba.Red, rgba.Green, rgba.Blue, rgba.Alpha), nil
}
