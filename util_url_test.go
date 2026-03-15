package hive

import "testing"

func TestEncodeUriComponent(t *testing.T) {
	if EncodeUriComponent("key=123 啊啊+-*/_.!~()'") != "key%3D123%20%E5%95%8A%E5%95%8A%2B-*%2F_.!~()'" {
		t.Errorf("EncodeUriComponent(%q) = %q, want %q",
			"key=123 啊啊+-*/_.!~()'", EncodeUriComponent("key=123 啊啊+-*/_.!~()'"), "key%3D123%20%E5%95%8A%E5%95%8A%2B-*%2F_.!~()'")
	}
}

func TestDecodeUriComponent(t *testing.T) {
	original, _ := DecodeUriComponent("key%3D123%20%E5%95%8A%E5%95%8A%2B-*%2F_.!~()'")
	if original != "key=123 啊啊+-*/_.!~()'" {
		t.Errorf("DecodeUriComponent(%q) = %q, want %q",
			"key%3D123%20%E5%95%8A%E5%95%8A%2B-*%2F_.!~()'", original, "key=123 啊啊+-*/_.!~()'")
	}
}

// 更简洁的写法：使用匿名结构体
func TestToHttpProtocolUrl(t *testing.T) {
	testCases := []struct {
		in  string
		out string
	}{
		{"http://example.com", "http://example.com"},
		{"https://example.com", "https://example.com"},
		{"//example.com", "https://example.com"},
		{"example.com", "https://example.com"},
	}

	for _, tc := range testCases {
		t.Run(tc.in, func(t *testing.T) {
			got := ToHttpProtocolUrl(tc.in)
			if got != tc.out {
				t.Errorf("ToHttpProtocolUrl(%q) = %q; 期望 %q", tc.in, got, tc.out)
			}
		})
	}
}

func TestToRelativeProtocolUrl(t *testing.T) {
	testCases := []struct {
		in  string
		out string
	}{
		{"http://example.com", "//example.com"},
		{"https://example.com", "//example.com"},
		{"//example.com", "//example.com"},
		{"example.com", "//example.com"},
	}

	for _, tc := range testCases {
		t.Run(tc.in, func(t *testing.T) {
			got := ToRelativeProtocolUrl(tc.in)
			if got != tc.out {
				t.Errorf("ToRelativeProtocolUrl(%q) = %q; 期望 %q", tc.in, got, tc.out)
			}
		})
	}
}
