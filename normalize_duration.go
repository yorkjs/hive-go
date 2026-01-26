package hive

import "math"

type IDuration struct {
	Days    int
	Hours   int
	Minutes int
	Seconds int
}

func NormalizeDuration(milliseconds int64) IDuration {
	result := IDuration{}

	if milliseconds <= 0 {
		return result
	}

	seconds := math.Ceil(float64(milliseconds) / float64(MS_SECOND))
	minutes := math.Floor(float64(milliseconds) / float64(MS_MINUTE))
	hours := math.Floor(float64(milliseconds) / float64(MS_HOUR))
	days := math.Floor(float64(milliseconds) / float64(MS_DAY))

	if days > 0 {
		result.Days = int(days)
	}
	if int(hours)%24 > 0 {
		result.Hours = int(hours) % 24
	}
	if int(minutes)%60 > 0 {
		result.Minutes = int(minutes) % 60
	}
	if int(seconds)%60 > 0 {
		result.Seconds = int(seconds) % 60
	}
	return result
}
