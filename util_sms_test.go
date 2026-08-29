package hive

import "testing"

func TestCalculateSmsCount(t *testing.T) {
	tests := []struct {
		name            string
		signatureName   string
		templateContent string
		perSmsCharCount int
		wantCharCount   int
		wantSmsCount    int
	}{
		{
			name:            "空内容",
			signatureName:   "",
			templateContent: "",
			perSmsCharCount: 70,
			wantCharCount:   0,
			wantSmsCount:    0,
		},
		{
			name:            "仅模板内容且不足一条",
			signatureName:   "",
			templateContent: "你好",
			perSmsCharCount: 70,
			wantCharCount:   2,
			wantSmsCount:    1,
		},
		{
			name:            "刚好一条",
			signatureName:   "",
			templateContent: "1234567890",
			perSmsCharCount: 10,
			wantCharCount:   10,
			wantSmsCount:    1,
		},
		{
			name:            "超过一条一个字符",
			signatureName:   "",
			templateContent: "12345678901",
			perSmsCharCount: 10,
			wantCharCount:   11,
			wantSmsCount:    2,
		},
		{
			name:            "刚好两条",
			signatureName:   "",
			templateContent: "12345678901234567890",
			perSmsCharCount: 10,
			wantCharCount:   20,
			wantSmsCount:    2,
		},
		{
			name:            "三条不足",
			signatureName:   "",
			templateContent: "123456789012345678901",
			perSmsCharCount: 10,
			wantCharCount:   21,
			wantSmsCount:    3,
		},
		{
			name:            "包含签名",
			signatureName:   "测试",
			templateContent: "你好",
			perSmsCharCount: 70,
			// 【测试】= 4 个字符，模板=2 个字符
			wantCharCount: 6,
			wantSmsCount:  1,
		},
		{
			name:            "签名导致跨条数",
			signatureName:   "测试",
			templateContent: "123456789",
			perSmsCharCount: 10,
			// 【测试】= 4，模板=9，总计=13
			wantCharCount: 13,
			wantSmsCount:  2,
		},
		{
			name:            "英文和中文都按一个字符计算",
			signatureName:   "签名",
			templateContent: "Hello你好😀😀",
			perSmsCharCount: 10,
			// Hello=5，你好=2，😀😀=2，签名=2，两个括号=2
			// 总计 5 + 2 + 2 + 2 + 2 = 13
			wantCharCount: 13,
			wantSmsCount:  2,
		},
		{
			name:            "短信条数配置为1",
			signatureName:   "",
			templateContent: "你好世界",
			perSmsCharCount: 1,
			wantCharCount:   4,
			wantSmsCount:    4,
		},
		{
			name:            "短信条数配置为0",
			signatureName:   "",
			templateContent: "你好",
			perSmsCharCount: 0,
			wantCharCount:   2,
			wantSmsCount:    0,
		},
		{
			name:            "短信条数配置为负数",
			signatureName:   "",
			templateContent: "你好",
			perSmsCharCount: -1,
			wantCharCount:   2,
			wantSmsCount:    0,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotCharCount, gotSmsCount := CalculateSmsCount(
				tt.signatureName,
				tt.templateContent,
				tt.perSmsCharCount,
			)

			if gotCharCount != tt.wantCharCount {
				t.Errorf(
					"charCount = %d, want %d",
					gotCharCount,
					tt.wantCharCount,
				)
			}

			if gotSmsCount != tt.wantSmsCount {
				t.Errorf(
					"smsCount = %d, want %d",
					gotSmsCount,
					tt.wantSmsCount,
				)
			}
		})
	}
}
