package hive

import "testing"

func TestFormatArea(t *testing.T) {
	tests := []struct {
		name    string
		area    IArea
		options []IFormatAreaOptions
		want    string
	}{
		{
			name: "Beijing Haidian (Default)",
			area: IArea{
				Country:  &INode{ID: 0, Name: "中国"},
				Province: &INode{ID: 0, Name: "北京市"},
				City:     &INode{ID: 0, Name: "市辖区"},
				District: &INode{ID: 0, Name: "海淀区"},
			},
			want: "北京 海淀区",
		},
		{
			name: "Beijing Haidian (Simplify: false)",
			area: IArea{
				Country:  &INode{ID: 0, Name: "中国"},
				Province: &INode{ID: 0, Name: "北京市"},
				City:     &INode{ID: 0, Name: "市辖区"},
				District: &INode{ID: 0, Name: "海淀区"},
			},
			options: []IFormatAreaOptions{{Simplify: false}},
			want:    "北京市海淀区",
		},
		{
			name: "Henan Zhengzhou (Default)",
			area: IArea{
				Country:  &INode{ID: 0, Name: "中国"},
				Province: &INode{ID: 0, Name: "河南省"},
				City:     &INode{ID: 0, Name: "郑州市"},
				District: &INode{ID: 0, Name: "管城回族区"},
			},
			want: "河南 郑州 管城回族区",
		},
		{
			name: "Henan Zhengzhou + Address (Default)",
			area: IArea{
				Country:  &INode{ID: 0, Name: "中国"},
				Province: &INode{ID: 0, Name: "河南省"},
				City:     &INode{ID: 0, Name: "郑州市"},
				District: &INode{ID: 0, Name: "管城回族区"},
				Address:  "测试地址",
			},
			want: "河南 郑州 管城回族区 测试地址",
		},
		{
			name: "Henan Zhengzhou (Simplify: false)",
			area: IArea{
				Country:  &INode{ID: 0, Name: "中国"},
				Province: &INode{ID: 0, Name: "河南省"},
				City:     &INode{ID: 0, Name: "郑州市"},
				District: &INode{ID: 0, Name: "管城回族区"},
			},
			options: []IFormatAreaOptions{{Simplify: false}},
			want:    "河南省郑州市管城回族区",
		},
		{
			name: "Henan Zhengzhou + Address (Simplify: false)",
			area: IArea{
				Country:  &INode{ID: 0, Name: "中国"},
				Province: &INode{ID: 0, Name: "河南省"},
				City:     &INode{ID: 0, Name: "郑州市"},
				District: &INode{ID: 0, Name: "管城回族区"},
				Address:  "测试地址",
			},
			options: []IFormatAreaOptions{{Simplify: false}},
			want:    "河南省郑州市管城回族区测试地址",
		},
		{
			name: "Hong Kong",
			area: IArea{
				Country: &INode{ID: 0, Name: "中国香港"},
			},
			want: "中国香港",
		},
		{
			name: "Macau",
			area: IArea{
				Country: &INode{ID: 0, Name: "中国澳门"},
			},
			want: "中国澳门",
		},
		{
			name: "Taiwan",
			area: IArea{
				Country: &INode{ID: 0, Name: "中国台湾"},
			},
			want: "中国台湾",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var got string
			if len(tt.options) > 0 {
				got = FormatArea(tt.area, tt.options...)
			} else {
				got = FormatArea(tt.area)
			}
			if got != tt.want {
				t.Errorf("FormatArea() = %v, want %v", got, tt.want)
			}
		})
	}

	if got := FormatCity("中国台湾"); got != "台湾" {
		t.Errorf("FormatCity(中国台湾) = %v; want 台湾", got)
	}
}
