package hive

import "testing"

func TestNormalizeDuration(t *testing.T) {
	data := NormalizeDuration(int64(MS_SECOND))
	if data.Days != 0 || data.Hours != 0 || data.Minutes != 0 || data.Seconds != 1 {
		t.Errorf("NormalizeDuration(MS_SECOND) = %+v; want {0 0 0 1}", data)
	}

	data = NormalizeDuration(int64(MS_MINUTE))
	if data.Days != 0 || data.Hours != 0 || data.Minutes != 1 || data.Seconds != 0 {
		t.Errorf("NormalizeDuration(MS_MINUTE) = %+v; want {0 0 1 0}", data)
	}

	data = NormalizeDuration(int64(MS_HOUR))
	if data.Days != 0 || data.Hours != 1 || data.Minutes != 0 || data.Seconds != 0 {
		t.Errorf("NormalizeDuration(MS_HOUR) = %+v; want {0 1 0 0}", data)
	}

	data = NormalizeDuration(int64(MS_DAY))
	if data.Days != 1 || data.Hours != 0 || data.Minutes != 0 || data.Seconds != 0 {
		t.Errorf("NormalizeDuration(MS_DAY) = %+v; want {1 0 0 0}", data)
	}

	data = NormalizeDuration(int64(MS_DAY + 2*MS_HOUR + 3*MS_MINUTE + 4*MS_SECOND))
	if data.Days != 1 || data.Hours != 2 || data.Minutes != 3 || data.Seconds != 4 {
		t.Errorf("NormalizeDuration(MS_DAY + 2 * MS_HOUR + 3 * MS_MINUTE + 4 * MS_SECOND) = %+v; want {1 2 3 4}", data)
	}

}
