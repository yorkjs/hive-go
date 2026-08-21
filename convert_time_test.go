package hive

import (
	"testing"
	"time"
)

func TestTimeConvert(t *testing.T) {

	date := time.Now()
	timestamp := TimeToTimestamp(date)

	timeField := TimeToTimeField(date)
	convertedTime := TimeFieldToTime(timeField)
	if timestamp != TimeToTimestamp(convertedTime) {
		t.Errorf("timeToTimestamp 失败: 期望 %d, 得到 %d", TimeToTimestamp(date), TimeToTimestamp(convertedTime))
	}

	if TimeToTimestamp(date) != timestamp {
		t.Errorf("timeToTimestamp 失败: 期望 %d, 得到 %d", timestamp, TimeToTimestamp(date))
	}
	if TimeToTimestamp(time.Time{}) != 0 {
		t.Errorf("timeToTimestamp 失败: 期望 %d, 得到 %d", 0, TimeToTimestamp(time.Time{}))
	}
	if TimeToTimestamp(TimestampToTime(timestamp)) != timestamp {
		t.Errorf("timestampToTime 失败: 期望 %d, 得到 %d", timestamp, TimeToTimestamp(TimestampToTime(timestamp)))
	}

	if !TimestampToTime(0).IsZero() {
		t.Errorf("timestampToTime 失败: 期望 %v, 得到 %v", true, false)
	}

	t1 := TimeFieldToTime(ITimeField{
		Year:  2025,
		Month: 5,
		Date:  5,
	})
	if TimeToMinuteSegment(t1) != 0 {
		t.Errorf("timeToMinuteSegment 失败: 期望 %d, 得到 %d", 0, TimeToMinuteSegment(t1))
	}

	t2 := TimeFieldToTime(ITimeField{
		Year:   2025,
		Month:  5,
		Date:   5,
		Hour:   10,
		Minute: 1,
	})
	if TimeToMinuteSegment(t2) != 601 {
		t.Errorf("timeToMinuteSegment 失败: 期望 %d, 得到 %d", 601, TimeToMinuteSegment(t2))
	}

	t3 := TimeFieldToTime(ITimeField{
		Year:   2025,
		Month:  5,
		Date:   5,
		Hour:   23,
		Minute: 59,
		Second: 59,
	})
	if TimeToMinuteSegment(t3) != 1439 {
		t.Errorf("timeToMinuteSegment 失败: 期望 %d, 得到 %d", 1439, TimeToMinuteSegment(t3))
	}

}
