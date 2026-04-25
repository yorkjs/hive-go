package hive

import "testing"

func TestGetStringLength(t *testing.T) {
	tests := []struct {
		input    string
		expected int
	}{
		{"1234", 4},
		{"1234你", 5},
		{"1234你好", 6},
	}

	for _, tt := range tests {
		t.Run(tt.input, func(t *testing.T) {
			result := GetStringLength(tt.input)
			if result != tt.expected {
				t.Errorf("GetStringLength(%q) = %d, want %d", tt.input, result, tt.expected)
			}
		})
	}
}

func TestGetStringWidth(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		expected int
	}{
		{"纯数字", "12", 2},
		{"数字+大写字母", "12A", 3},
		{"数字+大小写字母", "12Aa", 4},
		{"数字+字母+中文", "12Aa啊", 6},
		{"数字+字母+中文+下划线", "12Aa啊_", 7},
		{"数字+字母+中文+下划线+逗号", "12Aa啊_，", 9},
		{"数字+字母+中文+下划线+逗号+句号", "12Aa啊_，。", 11},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := GetStringWidth(tt.input)
			if result != tt.expected {
				t.Errorf("GetStringWidth(%q) = %v, 期望 %v", tt.input, result, tt.expected)
			}
		})
	}
}

func TestTrimString(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		expected string
	}{
		// 前导空白
		{"leading tab", "\t1234", "1234"},
		{"leading space", " 1234", "1234"},
		{"leading tab and space", "\t 1234", "1234"},

		// 尾部空白
		{"trailing tab", "1234\t", "1234"},
		{"trailing space", "1234 ", "1234"},
		{"trailing tab and space", "1234\t ", "1234"},

		// 两端空白
		{"both sides tab and space", "\t1234\t ", "1234"},
		{"with newline", "\t1234\n    ", "1234"},
		{"with newline both sides", "\n    1234\n    ", "1234"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := TrimString(tt.input)
			if result != tt.expected {
				t.Errorf("TrimString(%q) = %q, want %q", tt.input, result, tt.expected)
			}
		})
	}
}

func TestSliceString(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		start    int
		end      int
		expected string
	}{
		{"ascii only", "12345678", 4, 6, "56"},
		{"mixed 1", "1234你好", 3, 5, "4你"},
		{"mixed 2", "1234你好5", 5, 7, "好5"},
		{"mixed 3", "1234你好", 4, 6, "你好"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := SliceString(tt.input, tt.start, tt.end)
			if result != tt.expected {
				t.Errorf("SliceString(%q, %d, %d) = %q, want %q",
					tt.input, tt.start, tt.end, result, tt.expected)
			}
		})
	}
}

func TestTruncateString(t *testing.T) {
	tests := []struct {
		name      string
		input     string
		maxLength int
		expected  string
	}{
		{"ascii", "123456789", 5, "12..."},
		{"uppercase", "ABCDEFGHI", 5, "AB..."},
		{"mixed 1", "ABC你好呀", 5, "AB..."},
		{"mixed 2", "你好呀ABC", 5, "你好..."},
		{"chinese 1", "你是谁你在干什么你想吃什么", 9, "你是谁你在干..."},
		{"chinese 2", "你是谁你在干ABC想吃什么", 9, "你是谁你在干..."},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := TruncateString(tt.input, tt.maxLength)
			if result != tt.expected {
				t.Errorf("TruncateString(%q, %d) = %q, want %q",
					tt.input, tt.maxLength, result, tt.expected)
			}
		})
	}
}

func TestRenderStringTemplate(t *testing.T) {
	tests := []struct {
		name     string
		template string
		data     map[string]interface{}
		expected string
	}{
		{
			name:     "basic string template",
			template: "你好，${name}",
			data:     map[string]interface{}{"name": "张三"},
			expected: "你好，张三",
		},
		{
			name:     "multiple variables",
			template: "你好，${name1}，${name2}",
			data: map[string]interface{}{
				"name1": "张三",
				"name2": "李四",
			},
			expected: "你好，张三，李四",
		},
		{
			name:     "keep placeholder for missing variables",
			template: "你好，${name1}，${name3}。",
			data: map[string]interface{}{
				"name1": "张三",
			},
			expected: "你好，张三，${name3}。",
		},
		{
			name:     "render integer value",
			template: "你好，${value}。",
			data:     map[string]interface{}{"value": 10},
			expected: "你好，10。",
		},
		{
			name:     "render double value",
			template: "你好，${value}。",
			data:     map[string]interface{}{"value": 10.11},
			expected: "你好，10.11。",
		},
		{
			name:     "render boolean true",
			template: "你好，${value}。",
			data:     map[string]interface{}{"value": true},
			expected: "你好，true。",
		},
		{
			name:     "render boolean false",
			template: "你好，${value}。",
			data:     map[string]interface{}{"value": false},
			expected: "你好，false。",
		},
		{
			name:     "handle nil value",
			template: "你好，${value}。",
			data:     map[string]interface{}{"value": nil},
			expected: "你好，${value}。",
		},
		{
			name:     "multiple variable types",
			template: "值：${int}，${double}，${negative}",
			data: map[string]interface{}{
				"int":      42,
				"double":   3.14159,
				"negative": -100,
			},
			expected: "值：42，3.14159，-100",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := RenderStringTemplate(tt.template, tt.data)
			if result != tt.expected {
				t.Errorf("RenderStringTemplate(%q, %v) = %q, want %q",
					tt.template, tt.data, result, tt.expected)
			}
		})
	}
}

func TestPadStringStart(t *testing.T) {
	if PadStringStart("1", 3) != "001" {
		t.Errorf("PadStringStart(%q, %d) = %q, want %q",
			"1", 3, PadStringStart("1", 3), "001")
	}
	if PadStringStart("11", 3) != "011" {
		t.Errorf("PadStringStart(%q, %d) = %q, want %q",
			"11", 3, PadStringStart("11", 3), "011")
	}
	if PadStringStart("111", 3) != "111" {
		t.Errorf("PadStringStart(%q, %d) = %q, want %q",
			"111", 3, PadStringStart("111", 3), "111")
	}
}

func TestHasSpecialCharacter(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		expected bool
	}{
		{
			name:     "允许的字符组合（包含中文和标点）",
			input:    "abc,[1]23. 你好，【世界】！%",
			expected: false,
		},
		{
			name:     "包含制表符和换行符",
			input:    "abc,123. \t\n",
			expected: true,
		},
		{
			name:     "包含/",
			input:    "炸鸡风味瓜子/原味鸡风味花生二选一",
			expected: false,
		},
		{
			name:     "包含+",
			input:    "老北京鸡肉卷+吮指原味鸡",
			expected: false,
		},
		{
			name:     "包含-",
			input:    "老北京鸡肉卷-吮指原味鸡",
			expected: false,
		},
		{
			name:     "包含￥",
			input:    "老北京鸡肉卷-吮指原味鸡￥10",
			expected: false,
		},
		{
			name:     "包含表情符号",
			input:    "abc,123.☺️",
			expected: true,
		},
		{
			name:     "包含*",
			input:    "14*16",
			expected: false,
		},
		{
			name:     "包含π",
			input:    "茶π",
			expected: false,
		},
		{
			name:     "包含𰻝𰻝面",
			input:    "𰻝𰻝面",
			expected: false,
		},
		{
			name:     "普通文本包含空格",
			input:    " abc,  123. ",
			expected: false,
		},
		{
			name:     "空字符串",
			input:    "",
			expected: false,
		},
		{
			name:     "只包含允许的英文标点",
			input:    ".,!?;:-_",
			expected: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := HasSpecialCharacter(tt.input)
			if result != tt.expected {
				t.Errorf("HasSpecialCharacter(%q) = %v, 期望 %v", tt.input, result, tt.expected)
			}
		})
	}
}

func TestRemoveSpecialCharacter(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		expected string
	}{
		{
			name:     "允许的字符组合保持不变",
			input:    "abc,[1]23. 你好，【世界】！",
			expected: "abc,[1]23. 你好，【世界】！",
		},
		{
			name:     "移除制表符和换行符",
			input:    "abc,123. \t\n",
			expected: "abc,123. ",
		},
		{
			name:     "移除表情符号",
			input:    "a☺️bc,☺️123.^456☺️%",
			expected: "abc,123.456%",
		},
		{
			name:     "保留空格但去除首尾空格",
			input:    " abc,  123. ",
			expected: " abc,  123. ",
		},
		{
			name:     "空字符串",
			input:    "",
			expected: "",
		},
		{
			name:     "只包含特殊字符",
			input:    "@#$%^&*π",
			expected: "@#$%&*π",
		},
		{
			name:     "包含𰻝𰻝面",
			input:    "𰻝𰻝面",
			expected: "𰻝𰻝面",
		},
		{
			name:     "混合允许和不允许字符",
			input:    "Hello☺️@World#2024$Test",
			expected: "Hello@World#2024$Test",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := RemoveSpecialCharacter(tt.input)
			if result != tt.expected {
				t.Errorf("RemoveSpecialCharacter(%q) = %q, 期望 %q", tt.input, result, tt.expected)
			}
		})
	}
}
