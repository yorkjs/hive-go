package hive

import "time"

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

func StringToTime(str string, format string) (time.Time, error) {
	return time.Parse(format, str)
}
