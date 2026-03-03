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

func TestEncodeURIComponent(t *testing.T) {
	if EncodeURIComponent("key=123 啊啊+-*/_.!~()'") != "key%3D123%20%E5%95%8A%E5%95%8A%2B-*%2F_.!~()'" {
		t.Errorf("EncodeURIComponent(%q) = %q, want %q",
			"key=123 啊啊+-*/_.!~()'", EncodeURIComponent("key=123 啊啊+-*/_.!~()'"), "key%3D123%20%E5%95%8A%E5%95%8A%2B-*%2F_.!~()'")
	}
}

func TestDecodeURIComponent(t *testing.T) {
	original, _ := DecodeURIComponent("key%3D123%20%E5%95%8A%E5%95%8A%2B-*%2F_.!~()'")
	if original != "key=123 啊啊+-*/_.!~()'" {
		t.Errorf("DecodeURIComponent(%q) = %q, want %q",
			"key%3D123%20%E5%95%8A%E5%95%8A%2B-*%2F_.!~()'", original, "key=123 啊啊+-*/_.!~()'")
	}
}
