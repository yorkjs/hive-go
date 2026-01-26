package hive

import "testing"

func TestFormatRatePercent(t *testing.T) {
	// RateToDisplay: value / 100 (assuming basis points?)
	// Or value * 100?
	// TS: export function rateToDisplay(value: number) { return NP.times(value, 100) } ?
	// Need to check convert_rate.go.
	// If input is 0.1, output 10%.
	
	// I'll check convert_rate.go first, but assume standard percentage.
	// If FormatRatePercent(0.1234) -> "12.34%"
	
	// Wait, "FormatRatePercent 把万分比格式化为百分比" comment in TS/Go says "万分比".
	// If input is 100 (万分之100 = 1%).
	// Then RateToDisplay(100) should be 1.
	// Then "1%".
	
	// Let's assume input is basis points (1/10000).
	if got := FormatRatePercent(100); got != "1%" {
		t.Errorf("FormatRatePercent(100) = %v; want 1%%", got)
	}
}
