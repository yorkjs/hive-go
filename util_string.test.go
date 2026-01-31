package hive

import "testing"

func TestTruncateString(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		maxLen   int
		expected string
	}{
		{
			name:     "英文数字截断",
			input:    "123456789",
			maxLen:   5,
			expected: "12...",
		},
		{
			name:     "英文大写字母截断",
			input:    "ABCDEFGHI",
			maxLen:   5,
			expected: "AB...",
		},
		{
			name:     "中英文混合，英文在前",
			input:    "ABC你好呀",
			maxLen:   5,
			expected: "AB...",
		},
		{
			name:     "中英文混合，中文在前",
			input:    "你好呀ABC",
			maxLen:   5,
			expected: "你好...",
		},
		{
			name:     "纯中文截断",
			input:    "你是谁你在干什么你想吃什么",
			maxLen:   9,
			expected: "你是谁你在干...",
		},
		{
			name:     "中英文混合复杂截断",
			input:    "你是谁你在干ABC想吃什么",
			maxLen:   9,
			expected: "你是谁你在干...",
		},
		// 可以添加更多边界测试
		{
			name:     "长度刚好等于最大长度",
			input:    "Hello",
			maxLen:   5,
			expected: "Hello",
		},
		{
			name:     "长度小于最大长度",
			input:    "Hi",
			maxLen:   5,
			expected: "Hi",
		},
		{
			name:     "最大长度为3",
			input:    "ABCD",
			maxLen:   3,
			expected: "ABC",
		},
		{
			name:     "最大长度为2",
			input:    "ABCD",
			maxLen:   2,
			expected: "AB",
		},
		{
			name:     "空字符串",
			input:    "",
			maxLen:   5,
			expected: "",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			actual := TruncateString(tt.input, tt.maxLen)
			if actual != tt.expected {
				t.Errorf("TruncateString(%q, %d) = %q, expected %q",
					tt.input, tt.maxLen, actual, tt.expected)
			}
		})
	}
}
