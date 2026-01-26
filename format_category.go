package hive

import "strings"

type ICategory struct {
	Category1 *INode
	Category2 *INode
	Category3 *INode
}

func FormatCategory(category ICategory) string {
	var list []string

	if category.Category1 != nil {
		list = append(list, category.Category1.Name)
	}
	if category.Category2 != nil {
		list = append(list, category.Category2.Name)
	}
	if category.Category3 != nil {
		list = append(list, category.Category3.Name)
	}

	return strings.Join(list, "/")
}
