package hive

import (
	"fmt"
	"time"
)

// FormatWeek 把时间戳格式化为 周开始 ~ 周结束 格式
func FormatWeek(t time.Time) string {
	day := int(t.Weekday())
	offset := day
	startTime := t.AddDate(0, 0, -offset)
	return fmt.Sprintf("%s 至 %s", FormatDateShortly(startTime), FormatDateShortly(startTime.AddDate(0, 0, 6)))
}
