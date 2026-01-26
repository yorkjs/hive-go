package hive

import (
	"fmt"
	"strings"
)

func FormatHourMinutes(value int) string {
	hours := value / 60
	minutes := value % 60
	return fmt.Sprintf("%02d:%02d", hours, minutes)
}

// FormatBusinessTimes 营业时间时段范围为 [0, 2880] 可跨天, 0-1440 为当天，1440-2880 为次日
func FormatBusinessTimes(businessTimes []int) string {
	length := len(businessTimes)
	if length == 0 || length%2 != 0 {
		return ""
	}

	var timeRanges []string
	for i := 0; i < length; i += 2 {
		start := businessTimes[i]
		end := businessTimes[i+1]
		startTime := start % 1440
		endTime := end % 1440

		// 判断是否是全天
		if startTime == 0 && endTime == 0 && end > start {
			timeRanges = append(timeRanges, "全天")
			continue
		}

		startTimeStr := FormatHourMinutes(startTime)
		if start > 1440 {
			startTimeStr = "次日" + startTimeStr
		}

		endTimeStr := FormatHourMinutes(endTime)
		if end > 1440 {
			endTimeStr = "次日" + endTimeStr
		}

		timeRanges = append(timeRanges, fmt.Sprintf("%s-%s", startTimeStr, endTimeStr))
	}

	return strings.Join(timeRanges, "、")
}
