package hive

import "testing"

func TestMoneyConvert(t *testing.T) {
	if got := MoneyToDisplay(100); got != 1 {
		t.Errorf("MoneyToDisplay(100) = %v; want 1", got)
	}
	if got := MoneyToDisplay(100, MONEY_YUAN_TO_CENT); got != 1 {
		t.Errorf("MoneyToDisplay(100, MONEY_YUAN_TO_CENT) = %v; want 1", got)
	}
	if got := MoneyToDisplay(1000000, MONEY_TEN_THOUSAND_YUAN_TO_CENT); got != 1 {
		t.Errorf("MoneyToDisplay(1000000, MONEY_TEN_THOUSAND_YUAN_TO_CENT) = %v; want 1", got)
	}
	if got := MoneyToDisplay(1100000, MONEY_TEN_THOUSAND_YUAN_TO_CENT); got != 1.1 {
		t.Errorf("MoneyToDisplay(1100000, MONEY_TEN_THOUSAND_YUAN_TO_CENT) = %v; want 1.1", got)
	}

	if got := MoneyToBackend(1.01); got != 101 {
		t.Errorf("MoneyToBackend(1.01) = %v; want 101", got)
	}
	if got := MoneyToBackend(1.01, MONEY_YUAN_TO_CENT); got != 101 {
		t.Errorf("MoneyToBackend(1.01, MONEY_YUAN_TO_CENT) = %v; want 101", got)
	}
	if got := MoneyToBackend(1.01, MONEY_TEN_THOUSAND_YUAN_TO_CENT); got != 1010000 {
		t.Errorf("MoneyToBackend(1.01, MONEY_TEN_THOUSAND_YUAN_TO_CENT) = %v; want 1010000", got)
	}
}
