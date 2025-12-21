package is

import "testing"

func TestStandardBarcode(t *testing.T) {
	if !StandardBarcode("6924187812697") {
		t.Error("error")
	}
	if StandardBarcode("C6924187812697") {
		t.Error("error")
	}
}

func TestCustomBarcode(t *testing.T) {
	if CustomBarcode("6924187812697") {
		t.Error("error")
	}
	if !CustomBarcode("C6924187812") {
		t.Error("error")
	}
	if !CustomBarcode("C69241878121") {
		t.Error("error")
	}
	if !CustomBarcode("C692418781212") {
		t.Error("error")
	}
}
