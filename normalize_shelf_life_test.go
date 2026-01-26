package hive

import "testing"

func TestNormalizeShelfLifeEdgeCases(t *testing.T) {
	// 测试边界条件
	tests := []struct {
		name     string
		input    int
		expected IShelfLife
	}{
		{
			name:  "0小时",
			input: 0,
			expected: IShelfLife{
				Years:  0,
				Months: 0,
				Days:   0,
				Hours:  0,
			},
		},
		{
			name:  "负数小时",
			input: -10,
			expected: IShelfLife{
				Years:  0,
				Months: 0,
				Days:   0,
				Hours:  0,
			},
		},
		{
			name:  "正好两年",
			input: 2 * 24 * 365,
			expected: IShelfLife{
				Years:  2,
				Months: 0,
				Days:   0,
				Hours:  0,
			},
		},
		{
			name:  "复杂的组合",
			input: 2*24*365 + 3*24*30 + 5*24 + 12,
			expected: IShelfLife{
				Years:  2,
				Months: 3,
				Days:   5,
				Hours:  12,
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := NormalizeShelfLife(tt.input)

			if result.Years != tt.expected.Years ||
				result.Months != tt.expected.Months ||
				result.Days != tt.expected.Days ||
				result.Hours != tt.expected.Hours {
				t.Errorf("%s: 期望 %+v, 实际 %+v", tt.name, tt.expected, result)
			}
		})
	}
}
