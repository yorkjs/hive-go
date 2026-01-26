package hive

import "testing"

func TestFormatCategory(t *testing.T) {
	// Empty category
	if got := FormatCategory(ICategory{}); got != "" {
		t.Errorf("FormatCategory(empty) = %q; want \"\"", got)
	}

	// Case 1
	c1 := ICategory{
		Category1: &INode{ID: 0, Name: "美食餐饮"},
	}
	if got := FormatCategory(c1); got != "美食餐饮" {
		t.Errorf("FormatCategory(c1) = %q; want \"美食餐饮\"", got)
	}

	// Case 2
	c2 := ICategory{
		Category1: &INode{ID: 0, Name: "美食餐饮"},
		Category2: &INode{ID: 0, Name: "中餐"},
	}
	if got := FormatCategory(c2); got != "美食餐饮/中餐" {
		t.Errorf("FormatCategory(c2) = %q; want \"美食餐饮/中餐\"", got)
	}

	// Case 3
	c3 := ICategory{
		Category1: &INode{ID: 0, Name: "美食餐饮"},
		Category2: &INode{ID: 0, Name: "中餐"},
		Category3: &INode{ID: 0, Name: "烤鸭"},
	}
	if got := FormatCategory(c3); got != "美食餐饮/中餐/烤鸭" {
		t.Errorf("FormatCategory(c3) = %q; want \"美食餐饮/中餐/烤鸭\"", got)
	}
}
