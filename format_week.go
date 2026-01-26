package hive

import (
	"fmt"
	"time"
)

// FormatWeek 把时间戳格式化为 周开始 ~ 周结束 格式
func FormatWeek(timestamp int64) string {
	t := time.UnixMilli(timestamp)
	day := int(t.Weekday())
	offset := -1 * day
	startTimestamp := timestamp + int64(offset)*MS_DAY
	return fmt.Sprintf("%s 至 %s", FormatDateShortly(startTimestamp), FormatDateShortly(startTimestamp+6*MS_DAY))
}
