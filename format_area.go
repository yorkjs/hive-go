package hive

import (
	"strings"
)

var provinceMap = map[string]string{
	"北京市": "北京",
	"上海市": "上海",
	"天津市": "天津",
	"重庆市": "重庆",
}

var cityMap = map[string]string{
	"省直辖县级行政区划":   "省直辖",
	"自治区直辖县级行政区划": "省直辖",

	// 云南
	"楚雄彝族自治州":    "楚雄",
	"红河哈尼族彝族自治州": "红河",
	"文山壮族苗族自治州":  "文山",
	"西双版纳傣族自治州":  "西双版纳",
	"大理白族自治州":    "大理",
	"德宏傣族景颇族自治州": "德宏",
	"怒江傈僳族自治州":   "怒江",
	"迪庆藏族自治州":    "迪庆",

	// 湖北
	"恩施土家族苗族自治州": "恩施",

	// 湖南
	"湘西土家族苗族自治州": "湘西",

	// 甘肃
	"临夏回族自治州": "临夏",
	"甘南藏族自治州": "甘南",

	// 内蒙古
	"兴安盟":   "兴安",
	"锡林郭勒盟": "锡林郭勒",
	"阿拉善盟":  "阿拉善",

	// 青海
	"海北藏族自治州":    "海北",
	"黄南藏族自治州":    "黄南",
	"海南藏族自治州":    "海南州",
	"果洛藏族自治州":    "果洛",
	"玉树藏族自治州":    "玉树",
	"海西蒙古族藏族自治州": "海西",

	// 吉林
	"延边朝鲜族自治州": "延边",

	// 四川
	"阿坝藏族羌族自治州": "阿坝",
	"甘孜藏族自治州":   "甘孜",
	"凉山彝族自治州":   "凉山",

	// 新疆
	"昌吉回族自治州":     "昌吉",
	"博尔塔拉蒙古自治州":   "博尔塔拉",
	"巴音郭楞蒙古自治州":   "巴音郭楞",
	"阿克苏地区":       "阿克苏",
	"克孜勒苏柯尔克孜自治州": "克孜勒苏柯尔克孜",
	"喀什地区":        "喀什",
	"和田地区":        "和田",
	"伊犁哈萨克自治州":    "伊犁",
	"塔城地区":        "塔城",
	"阿勒泰地区":       "阿勒泰",

	// 贵州
	"黔西南布依族苗族自治州": "黔西南",
	"黔东南苗族侗族自治州":  "黔东南",
	"黔南布依族苗族自治州":  "黔南",
}

type INode struct {
	ID   int64
	Name string
}

type IArea struct {
	Country  *INode
	Province *INode
	City     *INode
	District *INode
	Address  string
}

type IFormatAreaOptions struct {
	Simplify bool
}

func FormatArea(area IArea, options ...IFormatAreaOptions) string {
	simplify := true
	if len(options) > 0 {
		simplify = options[0].Simplify
	}

	var list []string
	prevItem := ""

	appendItem := func(item string) {
		list = append(list, item)
		prevItem = item
	}

	isChina := false
	separator := ""
	if simplify {
		separator = " "
	}

	if area.Country != nil {
		item := area.Country.Name
		appendItem(item)
		if item == "中国" {
			isChina = true
		}
	}

	if area.Province != nil {
		name := area.Province.Name
		if simplify {
			name = FormatProvince(name)
		}
		appendItem(name)
	}

	if area.City != nil {
		item := area.City.Name
		if simplify {
			item = FormatCity(item)
		}
		if item != "" &&
			item != "市辖区" &&
			item != "县" &&
			item != "省直辖" &&
			item != prevItem {
			appendItem(item)
		}
	}

	if area.District != nil {
		item := area.District.Name
		if simplify {
			item = FormatDistrict(item)
		}
		if item != "" {
			appendItem(item)
		}
	}

	if area.Address != "" {
		appendItem(area.Address)
	}

	if len(list) > 1 && isChina {
		// 删掉中国
		list = list[1:]
	}

	return strings.Join(list, separator)
}

func FormatProvince(name string) string {
	if val, ok := provinceMap[name]; ok {
		return val
	}
	if strings.HasSuffix(name, "省") || strings.HasSuffix(name, "州") {
		return name[:len(name)-len("省")]
	}

	runes := []rune(name)
	if strings.HasPrefix(name, "新疆") ||
		strings.HasPrefix(name, "西藏") ||
		strings.HasPrefix(name, "宁夏") ||
		strings.HasPrefix(name, "广西") ||
		strings.HasPrefix(name, "香港") ||
		strings.HasPrefix(name, "澳门") {
		if len(runes) >= 2 {
			return string(runes[:2])
		}
	}
	if strings.HasPrefix(name, "内蒙古") {
		if len(runes) >= 3 {
			return string(runes[:3])
		}
	}
	return name
}

func FormatCity(name string) string {
	// 有时候会把港澳台放到城市这级显示
	if after, ok := strings.CutPrefix(name, "中国"); ok {
		name = after
	}

	if val, ok := cityMap[name]; ok {
		return val
	}

	if strings.HasSuffix(name, "市") {
		return strings.TrimSuffix(name, "市")
	}
	if strings.HasSuffix(name, "地区") {
		return strings.TrimSuffix(name, "地区")
	}
	if strings.HasSuffix(name, "自治州") {
		return strings.TrimSuffix(name, "自治州")
	}
	return name
}

func FormatDistrict(name string) string {
	if name == "市辖区" {
		return ""
	}
	// 不要处理区，会影响有效信息展示
	return name
}
