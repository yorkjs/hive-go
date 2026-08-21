package hive

import "time"

type ITimeField struct {
	Year        int
	Month       int
	Date        int
	Hour        int
	Minute      int
	Second      int
	Millisecond int
}

func TimeToTimeField(t time.Time) ITimeField {
	return ITimeField{
		Year:        t.Year(),
		Month:       int(t.Month()),
		Date:        t.Day(),
		Hour:        t.Hour(),
		Minute:      t.Minute(),
		Second:      t.Second(),
		Millisecond: int(t.Nanosecond() / 1000000),
	}
}

func TimeFieldToTime(t ITimeField) time.Time {
	return time.Date(
		t.Year, time.Month(t.Month), t.Date,
		t.Hour, t.Minute, t.Second, t.Millisecond*1000000,
		time.Local,
	)
}

// TimeToTimestamp 时间对象转成时间戳（毫秒）
func TimeToTimestamp(t time.Time) int64 {
	if t.IsZero() {
		return 0
	}
	return t.UnixMilli()
}

// TimestampToTime 时间戳转成时间对象
func TimestampToTime(timestamp int64) time.Time {
	if timestamp <= 0 {
		return time.Time{}
	}
	return time.UnixMilli(timestamp)
}

// TimeToMinuteSegment 返回给定时间在当天从 0 点开始的分钟数（0 ~ 1439）
func TimeToMinuteSegment(t time.Time) int {
	hours := t.Hour()
	minutes := t.Minute()
	return hours*60 + minutes
}
