package hive

import "testing"

func TestIsStandardBarcode(t *testing.T) {
	if got := IsStandardBarcode("6924187812697"); got != true {
		t.Errorf("IsStandardBarcode(6924187812697) = %v; want true", got)
	}
	if got := IsStandardBarcode("C6924187812697"); got != false {
		t.Errorf("IsStandardBarcode(C6924187812697) = %v; want false", got)
	}
}

func TestIsCustomBarcode(t *testing.T) {
	if got := IsCustomBarcode("6924187812697"); got != false {
		t.Errorf("IsCustomBarcode(6924187812697) = %v; want false", got)
	}
	if got := IsCustomBarcode("C6924187812"); got != true {
		t.Errorf("IsCustomBarcode(C6924187812) = %v; want true", got)
	}
	if got := IsCustomBarcode("C69241878121"); got != true {
		t.Errorf("IsCustomBarcode(C69241878121) = %v; want true", got)
	}
	if got := IsCustomBarcode("C692418781212"); got != true {
		t.Errorf("IsCustomBarcode(C692418781212) = %v; want true", got)
	}
}
