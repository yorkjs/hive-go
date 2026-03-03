package hive

import (
	"regexp"
	"strconv"
	"strings"
)

// IsIdentityCardNumber 是否为身份证号码
// value 要校验的值
// 返回是否为身份证号码
func IsIdentityCardNumber(value string) bool {
	// 1. 基础格式验证
	pattern := `^\d{6}(18|19|20)?\d{2}(0[1-9]|1[012])(0[1-9]|[12]\d|3[01])\d{3}(\d|X)$`
	match, _ := regexp.MatchString(pattern, value)
	if value == "" || !match {
		return false
	}

	// 2. 地址编码验证
	areaCode := value[0:2]
	if _, ok := areaMap[areaCode]; !ok {
		return false
	}

	// 3. 18 位身份证需要验证最后一位校验位
	if len(value) == 18 {
		// 加权因子
		factor := []int{7, 9, 10, 5, 8, 4, 2, 1, 6, 3, 7, 9, 10, 5, 8, 4, 2}
		// 校验位
		parity := []string{"1", "0", "X", "9", "8", "7", "6", "5", "4", "3", "2"}

		var sum int
		for i := range 17 {
			ai, err := strconv.Atoi(string(value[i]))
			if err != nil {
				return false
			}
			wi := factor[i]
			sum += ai * wi
		}

		last := parity[sum%11]
		if last != strings.ToUpper(string(value[17])) {
			return false
		}
	}

	return true
}

// 地区编码映射
var areaMap = map[string]string{
	"11": "北京",
	"12": "天津",
	"13": "河北",
	"14": "山西",
	"15": "内蒙古",
	"21": "辽宁",
	"22": "吉林",
	"23": "黑龙江",
	"31": "上海",
	"32": "江苏",
	"33": "浙江",
	"34": "安徽",
	"35": "福建",
	"36": "江西",
	"37": "山东",
	"41": "河南",
	"42": "湖北",
	"43": "湖南",
	"44": "广东",
	"45": "广西",
	"46": "海南",
	"50": "重庆",
	"51": "四川",
	"52": "贵州",
	"53": "云南",
	"54": "西藏",
	"61": "陕西",
	"62": "甘肃",
	"63": "青海",
	"64": "宁夏",
	"65": "新疆",
	"71": "台湾",
	"81": "香港",
	"82": "澳门",
	"91": "国外",
}
