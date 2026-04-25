package hive

import (
	"testing"
	"time"
)

func TestTimeParse(t *testing.T) {

	// 测试无效日期
	zeroTime, err := ParseTime("-", DATE_YEAR_MONTH_DATE)
	if err == nil {
		t.Error("解析无效日期应该报错")
	}
	if TimeToTimestamp(zeroTime) != 0 {
		t.Errorf("无效日期转换失败: 期望 0, 得到 %d", TimeToTimestamp(zeroTime))
	}

	time1, err := ParseTime(
		"2020-10-01 10:00:00",
		DATE_TIME_YEAR_MONTH_DATE_HOUR_MINUTE_SECOND,
	)
	if err != nil {
		t.Fatalf("解析日期时间(带秒)失败: %v", err)
	}

	formatted := FormatDateTime(
		time1,
		DATE_TIME_YEAR_MONTH_DATE_HOUR_MINUTE_SECOND,
	)
	if formatted != "2020-10-01 10:00:00" {
		t.Errorf("日期时间(带秒)格式化失败: 期望 2020-10-01 10:00:00, 得到 %s", formatted)
	}

	// 日期时间（不带秒）
	time2, err := ParseTime(
		"2020-10-01 10:00",
		DATE_TIME_YEAR_MONTH_DATE_HOUR_MINUTE,
	)
	if err != nil {
		t.Fatalf("解析日期时间(不带秒)失败: %v", err)
	}

	formatted2 := FormatDateTime(
		time2,
		DATE_TIME_YEAR_MONTH_DATE_HOUR_MINUTE,
	)
	if formatted2 != "2020-10-01 10:00" {
		t.Errorf("日期时间(不带秒)格式化失败: 期望 2020-10-01 10:00, 得到 %s", formatted2)
	}

	// 纯日期
	time3, err := ParseTime(
		"2020-10-01",
		DATE_YEAR_MONTH_DATE,
	)
	if err != nil {
		t.Fatalf("解析纯日期失败: %v", err)
	}

	formatted3 := FormatDateTime(
		time3,
		DATE_YEAR_MONTH_DATE,
	)
	if formatted3 != "2020-10-01" {
		t.Errorf("纯日期格式化失败: 期望 2020-10-01, 得到 %s", formatted3)
	}

	time1, err = ParseTime(
		"2020.10.01 10:00:00",
		DATE_TIME_YEAR_MONTH_DATE_HOUR_MINUTE_SECOND_DOT,
	)
	if err != nil {
		t.Fatalf("解析点号格式日期时间(带秒)失败: %v", err)
	}

	formatted = FormatDateTime(
		time1,
		DATE_TIME_YEAR_MONTH_DATE_HOUR_MINUTE_SECOND,
	)
	if formatted != "2020-10-01 10:00:00" {
		t.Errorf("点号格式日期时间(带秒)格式化失败: 期望 2020-10-01 10:00:00, 得到 %s", formatted)
	}

	// 日期时间（不带秒）
	time2, err = ParseTime(
		"2020.10.01 10:00",
		DATE_TIME_YEAR_MONTH_DATE_HOUR_MINUTE_DOT,
	)
	if err != nil {
		t.Fatalf("解析点号格式日期时间(不带秒)失败: %v", err)
	}

	formatted2 = FormatDateTime(
		time2,
		DATE_TIME_YEAR_MONTH_DATE_HOUR_MINUTE,
	)
	if formatted2 != "2020-10-01 10:00" {
		t.Errorf("点号格式日期时间(不带秒)格式化失败: 期望 2020-10-01 10:00, 得到 %s", formatted2)
	}

	// 纯日期
	time3, err = ParseTime(
		"2020.10.01",
		DATE_YEAR_MONTH_DATE_DOT,
	)
	if err != nil {
		t.Fatalf("解析点号格式纯日期失败: %v", err)
	}

	formatted3 = FormatDateTime(
		time3,
		DATE_YEAR_MONTH_DATE,
	)
	if formatted3 != "2020-10-01" {
		t.Errorf("点号格式纯日期格式化失败: 期望 2020-10-01, 得到 %s", formatted3)
	}

	time1, err = ParseTime(
		"2020/10/01 10:00:00",
		DATE_TIME_YEAR_MONTH_DATE_HOUR_MINUTE_SECOND_SLASH,
	)
	if err != nil {
		t.Fatalf("解析斜杠格式日期时间(带秒)失败: %v", err)
	}

	formatted = FormatDateTime(
		time1,
		DATE_TIME_YEAR_MONTH_DATE_HOUR_MINUTE_SECOND,
	)
	if formatted != "2020-10-01 10:00:00" {
		t.Errorf("斜杠格式日期时间(带秒)格式化失败: 期望 2020-10-01 10:00:00, 得到 %s", formatted)
	}

	// 日期时间（不带秒）
	time2, err = ParseTime(
		"2020/10/01 10:00",
		DATE_TIME_YEAR_MONTH_DATE_HOUR_MINUTE_SLASH,
	)
	if err != nil {
		t.Fatalf("解析斜杠格式日期时间(不带秒)失败: %v", err)
	}

	formatted2 = FormatDateTime(
		time2,
		DATE_TIME_YEAR_MONTH_DATE_HOUR_MINUTE,
	)
	if formatted2 != "2020-10-01 10:00" {
		t.Errorf("斜杠格式日期时间(不带秒)格式化失败: 期望 2020-10-01 10:00, 得到 %s", formatted2)
	}

	// 纯日期
	time3, err = ParseTime(
		"2020/10/01",
		DATE_YEAR_MONTH_DATE_SLASH,
	)
	if err != nil {
		t.Fatalf("解析斜杠格式纯日期失败: %v", err)
	}

	formatted3 = FormatDateTime(
		time3,
		DATE_YEAR_MONTH_DATE,
	)
	if formatted3 != "2020-10-01" {
		t.Errorf("斜杠格式纯日期格式化失败: 期望 2020-10-01, 得到 %s", formatted3)
	}

}

func TestTimeHour(t *testing.T) {
	loc := time.Local

	// 2020-10-10 10:01:01
	date := time.Date(2020, 10, 10, 10, 1, 1, 0, loc)

	if got := FormatDateTime(StartOfHour(date), DATE_TIME_YEAR_MONTH_DATE_HOUR_MINUTE_SECOND); got != "2020-10-10 10:00:00" {
		t.Errorf("StartOfHour(2020-10-10) = %v; want 2020-10-10 10:00:00", got)
	}
	if got := FormatDateTime(StartOfPrevHour(date), DATE_TIME_YEAR_MONTH_DATE_HOUR_MINUTE_SECOND); got != "2020-10-10 09:00:00" {
		t.Errorf("StartOfPrevHour(2020-10-10) = %v; want 2020-10-10 09:00:00", got)
	}
	if got := FormatDateTime(StartOfNextHour(date), DATE_TIME_YEAR_MONTH_DATE_HOUR_MINUTE_SECOND); got != "2020-10-10 11:00:00" {
		t.Errorf("StartOfNextHour(2020-10-10) = %v; want 2020-10-10 11:00:00", got)
	}
	if got := FormatDateTime(EndOfHour(date), DATE_TIME_YEAR_MONTH_DATE_HOUR_MINUTE_SECOND); got != "2020-10-10 10:59:59" {
		t.Errorf("EndOfHour(2020-10-10) = %v; want 2020-10-10 10:59:59", got)
	}

}

func TestTimeDay(t *testing.T) {
	loc := time.Local

	// 2020-10-10 10:01:01
	date := time.Date(2020, 10, 10, 10, 1, 1, 0, loc)

	if got := FormatDateTime(StartOfDay(date), DATE_TIME_YEAR_MONTH_DATE_HOUR_MINUTE_SECOND); got != "2020-10-10 00:00:00" {
		t.Errorf("StartOfDay(2020-10-10) = %v; want 2020-10-10 00:00:00", got)
	}
	if got := FormatDateTime(StartOfPrevDay(date), DATE_TIME_YEAR_MONTH_DATE_HOUR_MINUTE_SECOND); got != "2020-10-09 00:00:00" {
		t.Errorf("StartOfPrevDay(2020-10-10) = %v; want 2020-10-09 00:00:00", got)
	}
	if got := FormatDateTime(StartOfNextDay(date), DATE_TIME_YEAR_MONTH_DATE_HOUR_MINUTE_SECOND); got != "2020-10-11 00:00:00" {
		t.Errorf("StartOfNextDay(2020-10-10) = %v; want 2020-10-11 00:00:00", got)
	}
	if got := FormatDateTime(EndOfDay(date), DATE_TIME_YEAR_MONTH_DATE_HOUR_MINUTE_SECOND); got != "2020-10-10 23:59:59" {
		t.Errorf("EndOfDay(2020-10-10) = %v; want 2020-10-10 23:59:59", got)
	}

	// 2020-04-01 10:01:01
	date = time.Date(2020, 4, 1, 10, 1, 1, 0, loc)
	if got := FormatDateTime(StartOfPrevDay(date), DATE_TIME_YEAR_MONTH_DATE_HOUR_MINUTE_SECOND); got != "2020-03-31 00:00:00" {
		t.Errorf("StartOfPrevDay(2020-04-01) = %v; want 2020-03-31 00:00:00", got)
	}
	if got := FormatDateTime(StartOfNextDay(date), DATE_TIME_YEAR_MONTH_DATE_HOUR_MINUTE_SECOND); got != "2020-04-02 00:00:00" {
		t.Errorf("StartOfNextDay(2020-04-01) = %v; want 2020-04-02 00:00:00", got)
	}

	// 2020-01-01 10:01:01
	date = time.Date(2020, 1, 1, 10, 1, 1, 0, loc)
	if got := FormatDateTime(StartOfPrevDay(date), DATE_TIME_YEAR_MONTH_DATE_HOUR_MINUTE_SECOND); got != "2019-12-31 00:00:00" {
		t.Errorf("StartOfPrevDay(2020-01-01) = %v; want 2019-12-31 00:00:00", got)
	}
	if got := FormatDateTime(StartOfNextDay(date), DATE_TIME_YEAR_MONTH_DATE_HOUR_MINUTE_SECOND); got != "2020-01-02 00:00:00" {
		t.Errorf("StartOfNextDay(2020-01-01) = %v; want 2020-01-02 00:00:00", got)
	}

	// 2024-03-01 10:01:01 (Leap year)
	date = time.Date(2024, 3, 1, 10, 1, 1, 0, loc)
	if got := FormatDateTime(StartOfPrevDay(date), DATE_TIME_YEAR_MONTH_DATE_HOUR_MINUTE_SECOND); got != "2024-02-29 00:00:00" {
		t.Errorf("StartOfPrevDay(2024-03-01) = %v; want 2024-02-29 00:00:00", got)
	}
	if got := FormatDateTime(StartOfNextDay(date), DATE_TIME_YEAR_MONTH_DATE_HOUR_MINUTE_SECOND); got != "2024-03-02 00:00:00" {
		t.Errorf("StartOfNextDay(2024-03-01) = %v; want 2024-03-02 00:00:00", got)
	}

	// 2025-03-01 10:01:01 (Non-leap year)
	date = time.Date(2025, 3, 1, 10, 1, 1, 0, loc)
	if got := FormatDateTime(StartOfPrevDay(date), DATE_TIME_YEAR_MONTH_DATE_HOUR_MINUTE_SECOND); got != "2025-02-28 00:00:00" {
		t.Errorf("StartOfPrevDay(2025-03-01) = %v; want 2025-02-28 00:00:00", got)
	}
	if got := FormatDateTime(StartOfNextDay(date), DATE_TIME_YEAR_MONTH_DATE_HOUR_MINUTE_SECOND); got != "2025-03-02 00:00:00" {
		t.Errorf("StartOfNextDay(2025-03-01) = %v; want 2025-03-02 00:00:00", got)
	}
}

func TestTimeWeek(t *testing.T) {
	loc := time.Local

	// 2025-07-27 10:01:01 (Sunday)
	date := time.Date(2025, 7, 27, 10, 1, 1, 0, loc)

	if got := FormatDateTime(StartOfWeek(date), DATE_TIME_YEAR_MONTH_DATE_HOUR_MINUTE_SECOND); got != "2025-07-27 00:00:00" {
		t.Errorf("StartOfWeek(2025-07-27) = %v; want 2025-07-27 00:00:00", got)
	}
	if got := FormatDateTime(StartOfPrevWeek(date), DATE_TIME_YEAR_MONTH_DATE_HOUR_MINUTE_SECOND); got != "2025-07-20 00:00:00" {
		t.Errorf("StartOfPrevWeek(2025-07-27) = %v; want 2025-07-20 00:00:00", got)
	}
	if got := FormatDateTime(StartOfNextWeek(date), DATE_TIME_YEAR_MONTH_DATE_HOUR_MINUTE_SECOND); got != "2025-08-03 00:00:00" {
		t.Errorf("StartOfNextWeek(2025-07-27) = %v; want 2025-08-03 00:00:00", got)
	}
	if got := FormatDateTime(EndOfWeek(date), DATE_TIME_YEAR_MONTH_DATE_HOUR_MINUTE_SECOND); got != "2025-08-02 23:59:59" {
		t.Errorf("EndOfWeek(2025-07-27) = %v; want 2025-08-02 23:59:59", got)
	}

	// 2025-07-29 10:01:01 (Tuesday)
	date = time.Date(2025, 7, 29, 10, 1, 1, 0, loc)
	if got := FormatDateTime(StartOfWeek(date), DATE_TIME_YEAR_MONTH_DATE_HOUR_MINUTE_SECOND); got != "2025-07-27 00:00:00" {
		t.Errorf("StartOfWeek(2025-07-29) = %v; want 2025-07-27 00:00:00", got)
	}
	if got := FormatDateTime(EndOfWeek(date), DATE_TIME_YEAR_MONTH_DATE_HOUR_MINUTE_SECOND); got != "2025-08-02 23:59:59" {
		t.Errorf("EndOfWeek(2025-07-29) = %v; want 2025-08-02 23:59:59", got)
	}

	// 2025-08-01 10:01:01 (Friday)
	date = time.Date(2025, 8, 1, 10, 1, 1, 0, loc)
	if got := FormatDateTime(StartOfWeek(date), DATE_TIME_YEAR_MONTH_DATE_HOUR_MINUTE_SECOND); got != "2025-07-27 00:00:00" {
		t.Errorf("StartOfWeek(2025-08-01) = %v; want 2025-07-27 00:00:00", got)
	}
	if got := FormatDateTime(EndOfWeek(date), DATE_TIME_YEAR_MONTH_DATE_HOUR_MINUTE_SECOND); got != "2025-08-02 23:59:59" {
		t.Errorf("EndOfWeek(2025-08-01) = %v; want 2025-08-02 23:59:59", got)
	}

	// 2025-08-02 10:01:01 (Saturday)
	date = time.Date(2025, 8, 2, 10, 1, 1, 0, loc)
	if got := FormatDateTime(StartOfWeek(date), DATE_TIME_YEAR_MONTH_DATE_HOUR_MINUTE_SECOND); got != "2025-07-27 00:00:00" {
		t.Errorf("StartOfWeek(2025-08-02) = %v; want 2025-07-27 00:00:00", got)
	}
	if got := FormatDateTime(EndOfWeek(date), DATE_TIME_YEAR_MONTH_DATE_HOUR_MINUTE_SECOND); got != "2025-08-02 23:59:59" {
		t.Errorf("EndOfWeek(2025-08-02) = %v; want 2025-08-02 23:59:59", got)
	}
}

func TestTimeMonth(t *testing.T) {
	loc := time.Local

	// 2025-02-10 10:01:01
	date := time.Date(2025, 2, 10, 10, 1, 1, 0, loc)

	if got := FormatDateTime(StartOfMonth(date), DATE_TIME_YEAR_MONTH_DATE_HOUR_MINUTE_SECOND); got != "2025-02-01 00:00:00" {
		t.Errorf("StartOfMonth(2025-02-10) = %v; want 2025-02-01 00:00:00", got)
	}
	if got := FormatDateTime(StartOfPrevMonth(date), DATE_TIME_YEAR_MONTH_DATE_HOUR_MINUTE_SECOND); got != "2025-01-01 00:00:00" {
		t.Errorf("StartOfPrevMonth(2025-02-10) = %v; want 2025-01-01 00:00:00", got)
	}
	if got := FormatDateTime(StartOfNextMonth(date), DATE_TIME_YEAR_MONTH_DATE_HOUR_MINUTE_SECOND); got != "2025-03-01 00:00:00" {
		t.Errorf("StartOfNextMonth(2025-02-10) = %v; want 2025-03-01 00:00:00", got)
	}
	if got := FormatDateTime(EndOfMonth(date), DATE_TIME_YEAR_MONTH_DATE_HOUR_MINUTE_SECOND); got != "2025-02-28 23:59:59" {
		t.Errorf("EndOfMonth(2025-02-10) = %v; want 2025-02-28 23:59:59", got)
	}

	// 2025-06-29 10:01:01
	date = time.Date(2025, 6, 29, 10, 1, 1, 0, loc)
	if got := FormatDateTime(StartOfMonth(date), DATE_TIME_YEAR_MONTH_DATE_HOUR_MINUTE_SECOND); got != "2025-06-01 00:00:00" {
		t.Errorf("StartOfMonth(2025-06-29) = %v; want 2025-06-01 00:00:00", got)
	}
	if got := FormatDateTime(StartOfPrevMonth(date), DATE_TIME_YEAR_MONTH_DATE_HOUR_MINUTE_SECOND); got != "2025-05-01 00:00:00" {
		t.Errorf("StartOfPrevMonth(2025-06-29) = %v; want 2025-05-01 00:00:00", got)
	}
	if got := FormatDateTime(StartOfNextMonth(date), DATE_TIME_YEAR_MONTH_DATE_HOUR_MINUTE_SECOND); got != "2025-07-01 00:00:00" {
		t.Errorf("StartOfNextMonth(2025-06-29) = %v; want 2025-07-01 00:00:00", got)
	}
	if got := FormatDateTime(EndOfMonth(date), DATE_TIME_YEAR_MONTH_DATE_HOUR_MINUTE_SECOND); got != "2025-06-30 23:59:59" {
		t.Errorf("EndOfMonth(2025-06-29) = %v; want 2025-06-30 23:59:59", got)
	}

	// 2025-07-29 10:01:01
	date = time.Date(2025, 7, 29, 10, 1, 1, 0, loc)
	if got := FormatDateTime(StartOfMonth(date), DATE_TIME_YEAR_MONTH_DATE_HOUR_MINUTE_SECOND); got != "2025-07-01 00:00:00" {
		t.Errorf("StartOfMonth(2025-07-29) = %v; want 2025-07-01 00:00:00", got)
	}
	if got := FormatDateTime(StartOfPrevMonth(date), DATE_TIME_YEAR_MONTH_DATE_HOUR_MINUTE_SECOND); got != "2025-06-01 00:00:00" {
		t.Errorf("StartOfPrevMonth(2025-07-29) = %v; want 2025-06-01 00:00:00", got)
	}
	if got := FormatDateTime(StartOfNextMonth(date), DATE_TIME_YEAR_MONTH_DATE_HOUR_MINUTE_SECOND); got != "2025-08-01 00:00:00" {
		t.Errorf("StartOfNextMonth(2025-07-29) = %v; want 2025-08-01 00:00:00", got)
	}
	if got := FormatDateTime(EndOfMonth(date), DATE_TIME_YEAR_MONTH_DATE_HOUR_MINUTE_SECOND); got != "2025-07-31 23:59:59" {
		t.Errorf("EndOfMonth(2025-07-29) = %v; want 2025-07-31 23:59:59", got)
	}

	// 2025-12-19 10:01:01
	date = time.Date(2025, 12, 19, 10, 1, 1, 0, loc)
	if got := FormatDateTime(StartOfMonth(date), DATE_TIME_YEAR_MONTH_DATE_HOUR_MINUTE_SECOND); got != "2025-12-01 00:00:00" {
		t.Errorf("StartOfMonth(2025-12-19) = %v; want 2025-12-01 00:00:00", got)
	}
	if got := FormatDateTime(StartOfPrevMonth(date), DATE_TIME_YEAR_MONTH_DATE_HOUR_MINUTE_SECOND); got != "2025-11-01 00:00:00" {
		t.Errorf("StartOfPrevMonth(2025-12-19) = %v; want 2025-11-01 00:00:00", got)
	}
	if got := FormatDateTime(StartOfNextMonth(date), DATE_TIME_YEAR_MONTH_DATE_HOUR_MINUTE_SECOND); got != "2026-01-01 00:00:00" {
		t.Errorf("StartOfNextMonth(2025-12-19) = %v; want 2026-01-01 00:00:00", got)
	}
	if got := FormatDateTime(EndOfMonth(date), DATE_TIME_YEAR_MONTH_DATE_HOUR_MINUTE_SECOND); got != "2025-12-31 23:59:59" {
		t.Errorf("EndOfMonth(2025-12-19) = %v; want 2025-12-31 23:59:59", got)
	}

	// 2025-04-19 10:01:01
	date = time.Date(2025, 4, 19, 10, 1, 1, 0, loc)
	if got := FormatDateTime(StartOfMonth(date), DATE_TIME_YEAR_MONTH_DATE_HOUR_MINUTE_SECOND); got != "2025-04-01 00:00:00" {
		t.Errorf("StartOfMonth(2025-04-19) = %v; want 2025-04-01 00:00:00", got)
	}
	if got := FormatDateTime(StartOfPrevMonth(date), DATE_TIME_YEAR_MONTH_DATE_HOUR_MINUTE_SECOND); got != "2025-03-01 00:00:00" {
		t.Errorf("StartOfPrevMonth(2025-04-19) = %v; want 2025-03-01 00:00:00", got)
	}
	if got := FormatDateTime(StartOfNextMonth(date), DATE_TIME_YEAR_MONTH_DATE_HOUR_MINUTE_SECOND); got != "2025-05-01 00:00:00" {
		t.Errorf("StartOfNextMonth(2025-04-19) = %v; want 2025-05-01 00:00:00", got)
	}
	if got := FormatDateTime(EndOfMonth(date), DATE_TIME_YEAR_MONTH_DATE_HOUR_MINUTE_SECOND); got != "2025-04-30 23:59:59" {
		t.Errorf("EndOfMonth(2025-04-19) = %v; want 2025-04-30 23:59:59", got)
	}

	// 2025-01-19 10:01:01
	date = time.Date(2025, 1, 19, 10, 1, 1, 0, loc)
	if got := FormatDateTime(StartOfMonth(date), DATE_TIME_YEAR_MONTH_DATE_HOUR_MINUTE_SECOND); got != "2025-01-01 00:00:00" {
		t.Errorf("StartOfMonth(2025-01-19) = %v; want 2025-01-01 00:00:00", got)
	}
	if got := FormatDateTime(StartOfPrevMonth(date), DATE_TIME_YEAR_MONTH_DATE_HOUR_MINUTE_SECOND); got != "2024-12-01 00:00:00" {
		t.Errorf("StartOfPrevMonth(2025-01-19) = %v; want 2024-12-01 00:00:00", got)
	}
	if got := FormatDateTime(StartOfNextMonth(date), DATE_TIME_YEAR_MONTH_DATE_HOUR_MINUTE_SECOND); got != "2025-02-01 00:00:00" {
		t.Errorf("StartOfNextMonth(2025-01-19) = %v; want 2025-02-01 00:00:00", got)
	}
	if got := FormatDateTime(EndOfMonth(date), DATE_TIME_YEAR_MONTH_DATE_HOUR_MINUTE_SECOND); got != "2025-01-31 23:59:59" {
		t.Errorf("EndOfMonth(2025-01-19) = %v; want 2025-01-31 23:59:59", got)
	}
}

func TestSameOfPrevDay(t *testing.T) {
	tests := []struct {
		name string
		now  time.Time
		want time.Time
	}{
		{
			name: "普通日期",
			now:  time.Date(2026, 4, 27, 15, 30, 45, 123456789, time.Local),
			want: time.Date(2026, 4, 26, 15, 30, 45, 123456789, time.Local),
		},
		{
			name: "跨月边界-月初",
			now:  time.Date(2026, 5, 1, 10, 0, 0, 0, time.Local),
			want: time.Date(2026, 4, 30, 10, 0, 0, 0, time.Local),
		},
		{
			name: "跨年边界-元旦",
			now:  time.Date(2026, 1, 1, 8, 30, 0, 0, time.Local),
			want: time.Date(2025, 12, 31, 8, 30, 0, 0, time.Local),
		},
		{
			name: "闰年边界-3月1日",
			now:  time.Date(2024, 3, 1, 12, 0, 0, 0, time.Local),
			want: time.Date(2024, 2, 29, 12, 0, 0, 0, time.Local),
		},
		{
			name: "闰年边界-非闰年3月1日",
			now:  time.Date(2023, 3, 1, 12, 0, 0, 0, time.Local),
			want: time.Date(2023, 2, 28, 12, 0, 0, 0, time.Local),
		},
		{
			name: "包含纳秒",
			now:  time.Date(2026, 12, 31, 23, 59, 59, 999999999, time.UTC),
			want: time.Date(2026, 12, 30, 23, 59, 59, 999999999, time.UTC),
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := SameOfPrevDay(tt.now)
			if !got.Equal(tt.want) {
				t.Errorf("SameOfPrevDay() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestSameOfPrevWeek(t *testing.T) {
	tests := []struct {
		name string
		now  time.Time
		want time.Time
	}{
		{
			name: "普通日期",
			now:  time.Date(2026, 4, 27, 15, 30, 45, 123456789, time.Local),
			want: time.Date(2026, 4, 20, 15, 30, 45, 123456789, time.Local),
		},
		{
			name: "跨月边界",
			now:  time.Date(2026, 4, 5, 10, 0, 0, 0, time.Local),
			want: time.Date(2026, 3, 29, 10, 0, 0, 0, time.Local),
		},
		{
			name: "跨年边界",
			now:  time.Date(2026, 1, 10, 8, 30, 0, 0, time.Local),
			want: time.Date(2026, 1, 3, 8, 30, 0, 0, time.Local),
		},
		{
			name: "跨年边界-年初",
			now:  time.Date(2026, 1, 5, 12, 0, 0, 0, time.Local),
			want: time.Date(2025, 12, 29, 12, 0, 0, 0, time.Local),
		},
		{
			name: "闰年边界",
			now:  time.Date(2024, 3, 7, 14, 20, 0, 0, time.Local),
			want: time.Date(2024, 2, 29, 14, 20, 0, 0, time.Local),
		},
		{
			name: "包含纳秒",
			now:  time.Date(2026, 12, 31, 23, 59, 59, 999999999, time.UTC),
			want: time.Date(2026, 12, 24, 23, 59, 59, 999999999, time.UTC),
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := SameOfPrevWeek(tt.now)
			if !got.Equal(tt.want) {
				t.Errorf("SameOfPrevWeek() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestSameOfPrevMonth(t *testing.T) {
	tests := []struct {
		name string
		now  time.Time
		want time.Time
	}{
		{
			name: "普通日期-月中",
			now:  time.Date(2026, 4, 15, 15, 30, 45, 123456789, time.Local),
			want: time.Date(2026, 3, 15, 15, 30, 45, 123456789, time.Local),
		},
		{
			name: "月末-上个月有31天",
			now:  time.Date(2026, 4, 30, 10, 0, 0, 0, time.Local),
			want: time.Date(2026, 3, 30, 10, 0, 0, 0, time.Local),
		},
		{
			name: "月末-上个月有30天",
			now:  time.Date(2026, 5, 31, 12, 0, 0, 0, time.Local),
			want: time.Date(2026, 4, 30, 23, 59, 59, 999000000, time.Local),
		},
		{
			name: "月末-上个月是二月非闰年28天",
			now:  time.Date(2026, 3, 31, 8, 30, 0, 0, time.Local),
			want: time.Date(2026, 2, 28, 23, 59, 59, 999000000, time.Local),
		},
		{
			name: "月末-上个月是二月闰年29天",
			now:  time.Date(2024, 3, 31, 14, 20, 0, 0, time.Local),
			want: time.Date(2024, 2, 29, 23, 59, 59, 999000000, time.Local),
		},
		{
			name: "月初-1号",
			now:  time.Date(2026, 4, 1, 9, 15, 0, 0, time.Local),
			want: time.Date(2026, 3, 1, 9, 15, 0, 0, time.Local),
		},
		{
			name: "跨年-1月",
			now:  time.Date(2026, 1, 15, 20, 30, 0, 0, time.Local),
			want: time.Date(2025, 12, 15, 20, 30, 0, 0, time.Local),
		},
		{
			name: "跨年-1月31日",
			now:  time.Date(2026, 1, 31, 22, 45, 0, 0, time.Local),
			want: time.Date(2025, 12, 31, 22, 45, 0, 0, time.Local),
		},
		{
			name: "包含纳秒",
			now:  time.Date(2026, 6, 30, 15, 30, 45, 123456789, time.Local),
			want: time.Date(2026, 5, 30, 15, 30, 45, 123456789, time.Local),
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := SameOfPrevMonth(tt.now)
			if !got.Equal(tt.want) {
				t.Errorf("SameOfPrevMonth() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestSameOfPrevYear(t *testing.T) {
	tests := []struct {
		name string
		now  time.Time
		want time.Time
	}{
		{
			name: "普通日期",
			now:  time.Date(2026, 4, 15, 15, 30, 45, 123456789, time.Local),
			want: time.Date(2025, 4, 15, 15, 30, 45, 123456789, time.Local),
		},
		{
			name: "闰年2月29日-去年是闰年",
			now:  time.Date(2024, 2, 29, 10, 0, 0, 0, time.Local),
			want: time.Date(2023, 2, 28, 23, 59, 59, 999000000, time.Local),
		},
		{
			name: "闰年2月29日-去年是闰年但今年不是",
			now:  time.Date(2020, 2, 29, 14, 30, 0, 0, time.Local),
			want: time.Date(2019, 2, 28, 23, 59, 59, 999000000, time.Local),
		},
		{
			name: "闰年2月28日-去年是闰年",
			now:  time.Date(2024, 2, 28, 12, 0, 0, 0, time.Local),
			want: time.Date(2023, 2, 28, 12, 0, 0, 0, time.Local),
		},
		{
			name: "闰年3月1日-去年是闰年",
			now:  time.Date(2024, 3, 1, 8, 30, 0, 0, time.Local),
			want: time.Date(2023, 3, 1, 8, 30, 0, 0, time.Local),
		},
		{
			name: "闰年2月29日-去年是闰年但从闰年跨闰年",
			now:  time.Date(2024, 2, 29, 20, 15, 0, 0, time.Local),
			want: time.Date(2023, 2, 28, 23, 59, 59, 999000000, time.Local),
		},
		{
			name: "平年2月28日-去年是平年",
			now:  time.Date(2023, 2, 28, 9, 0, 0, 0, time.Local),
			want: time.Date(2022, 2, 28, 9, 0, 0, 0, time.Local),
		},
		{
			name: "跨世纪-2000年闰年",
			now:  time.Date(2000, 2, 29, 13, 45, 0, 0, time.Local),
			want: time.Date(1999, 2, 28, 23, 59, 59, 999000000, time.Local),
		},
		{
			name: "12月31日",
			now:  time.Date(2026, 12, 31, 23, 59, 59, 0, time.Local),
			want: time.Date(2025, 12, 31, 23, 59, 59, 0, time.Local),
		},
		{
			name: "1月1日",
			now:  time.Date(2026, 1, 1, 0, 0, 0, 0, time.Local),
			want: time.Date(2025, 1, 1, 0, 0, 0, 0, time.Local),
		},
		{
			name: "包含纳秒",
			now:  time.Date(2026, 6, 15, 12, 30, 45, 123456789, time.UTC),
			want: time.Date(2025, 6, 15, 12, 30, 45, 123456789, time.UTC),
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := SameOfPrevYear(tt.now)
			if !got.Equal(tt.want) {
				t.Errorf("SameOfPrevYear() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestTimeOptimizeTimeRange(t *testing.T) {
	loc := time.Local

	// 2025-02-10 10:01:01
	date := time.Date(2025, 2, 10, 10, 1, 1, 0, loc)

	// isHour
	startTime := StartOfHour(date)
	endTime := EndOfHour(date)
	isHour := false
	isDay := false
	isWeek := false
	isMonth := false
	isRange := false

	OptimizeTimeRange(startTime, endTime, ITimeRangeOptimizer{
		IsHour: func(hour time.Time) {
			isHour = true
			if hour != startTime {
				t.Errorf("IsHour: expected %v, got %v", startTime, hour)
			}
		},
		IsDay: func(day time.Time) {
			isDay = true
		},
		IsWeek: func(week time.Time) {
			isWeek = true
		},
		IsMonth: func(month time.Time) {
			isMonth = true
		},
		IsRange: func(start, end time.Time) {
			isRange = true
		},
	})

	if !isHour {
		t.Errorf("expected isHour to be false, got true")
	}
	if isDay {
		t.Errorf("expected isDay to be true, got false")
	}
	if isWeek {
		t.Errorf("expected isWeek to be false, got true")
	}
	if isMonth {
		t.Errorf("expected isMonth to be false, got true")
	}
	if isRange {
		t.Errorf("expected isRange to be false, got true")
	}

	// isHour 但是不传 isHour 函数
	startTime = StartOfHour(date)
	endTime = EndOfHour(date)
	isHour = false
	isDay = false
	isWeek = false
	isMonth = false
	isRange = false

	OptimizeTimeRange(startTime, endTime, ITimeRangeOptimizer{
		IsDay: func(day time.Time) {
			isDay = true
		},
		IsWeek: func(week time.Time) {
			isWeek = true
		},
		IsMonth: func(month time.Time) {
			isMonth = true
		},
		IsRange: func(start, end time.Time) {
			isRange = true
			if start != startTime {
				t.Errorf("IsRange start: expected %v, got %v", startTime, start)
			}
			if end != endTime {
				t.Errorf("IsRange end: expected %v, got %v", endTime, end)
			}
		},
	})

	if isHour {
		t.Errorf("expected isHour to be false, got true")
	}
	if isDay {
		t.Errorf("expected isDay to be true, got false")
	}
	if isWeek {
		t.Errorf("expected isWeek to be false, got true")
	}
	if isMonth {
		t.Errorf("expected isMonth to be false, got true")
	}
	if !isRange {
		t.Errorf("expected isRange to be false, got true")
	}

	// isDay

	startTime = StartOfDay(date)
	endTime = EndOfDay(date)
	isHour = false
	isDay = false
	isWeek = false
	isMonth = false
	isRange = false

	OptimizeTimeRange(startTime, endTime, ITimeRangeOptimizer{
		IsHour: func(hour time.Time) {
			isHour = true
		},
		IsDay: func(day time.Time) {
			isDay = true
			if day != startTime {
				t.Errorf("IsDay: expected %v, got %v", startTime, day)
			}
		},
		IsWeek: func(week time.Time) {
			isWeek = true
		},
		IsMonth: func(month time.Time) {
			isMonth = true
		},
		IsRange: func(start, end time.Time) {
			isRange = true
		},
	})

	if isHour {
		t.Errorf("expected isHour to be false, got true")
	}
	if !isDay {
		t.Errorf("expected isDay to be true, got false")
	}
	if isWeek {
		t.Errorf("expected isWeek to be false, got true")
	}
	if isMonth {
		t.Errorf("expected isMonth to be false, got true")
	}
	if isRange {
		t.Errorf("expected isRange to be false, got true")
	}

	// isDay 但是不传 isDay 函数
	startTime = StartOfDay(date)
	endTime = EndOfDay(date)
	isDay = false
	isWeek = false
	isMonth = false
	isRange = false

	OptimizeTimeRange(startTime, endTime, ITimeRangeOptimizer{
		IsWeek: func(week time.Time) {
			isWeek = true
		},
		IsMonth: func(month time.Time) {
			isMonth = true
		},
		IsRange: func(start, end time.Time) {
			isRange = true
			if start != startTime {
				t.Errorf("IsRange start: expected %v, got %v", startTime, start)
			}
			if end != endTime {
				t.Errorf("IsRange end: expected %v, got %v", endTime, end)
			}
		},
	})

	if isDay {
		t.Errorf("expected isDay to be false, got true")
	}
	if isWeek {
		t.Errorf("expected isWeek to be false, got true")
	}
	if isMonth {
		t.Errorf("expected isMonth to be false, got true")
	}
	if !isRange {
		t.Errorf("expected isRange to be true, got false")
	}

	// 截取日期中间的一段时间
	startTime = time.Date(2025, 10, 10, 10, 0, 0, 0, loc)
	endTime = time.Date(2025, 10, 10, 12, 0, 0, 0, loc)
	isHour = false
	isDay = false
	isWeek = false
	isMonth = false
	isRange = false

	OptimizeTimeRange(startTime, endTime, ITimeRangeOptimizer{
		IsDay: func(day time.Time) {
			isDay = true
		},
		IsWeek: func(week time.Time) {
			isWeek = true
		},
		IsMonth: func(month time.Time) {
			isMonth = true
		},
		IsRange: func(start, end time.Time) {
			isRange = true
			if start != startTime {
				t.Errorf("IsRange start: expected %v, got %v", startTime, start)
			}
			if end != endTime {
				t.Errorf("IsRange end: expected %v, got %v", endTime, end)
			}
		},
	})

	if isDay {
		t.Errorf("expected isDay to be false, got true")
	}
	if isWeek {
		t.Errorf("expected isWeek to be false, got true")
	}
	if isMonth {
		t.Errorf("expected isMonth to be false, got true")
	}
	if !isRange {
		t.Errorf("expected isRange to be true, got false")
	}

	// 跨天
	startTime = StartOfDay(time.Date(2025, 10, 10, 10, 0, 0, 0, loc))
	endTime = EndOfDay(time.Date(2025, 10, 12, 12, 0, 0, 0, loc))
	isDay = false
	isWeek = false
	isMonth = false
	isRange = false

	OptimizeTimeRange(startTime, endTime, ITimeRangeOptimizer{
		IsDay: func(day time.Time) {
			isDay = true
		},
		IsWeek: func(week time.Time) {
			isWeek = true
		},
		IsMonth: func(month time.Time) {
			isMonth = true
		},
		IsRange: func(start, end time.Time) {
			isRange = true
			if start != startTime {
				t.Errorf("IsRange start: expected %v, got %v", startTime, start)
			}
			if end != endTime {
				t.Errorf("IsRange end: expected %v, got %v", endTime, end)
			}
		},
	})

	if isDay {
		t.Errorf("expected isDay to be false, got true")
	}
	if isWeek {
		t.Errorf("expected isWeek to be false, got true")
	}
	if isMonth {
		t.Errorf("expected isMonth to be false, got true")
	}
	if !isRange {
		t.Errorf("expected isRange to be true, got false")
	}

	// isWeek
	startTime = StartOfDay(time.Date(2026, 1, 4, 10, 0, 0, 0, loc))
	endTime = EndOfDay(time.Date(2026, 1, 10, 12, 0, 0, 0, loc))
	isDay = false
	isWeek = false
	isMonth = false
	isRange = false

	// 注意：这里需要根据实际的 OptimizeTimeRange 逻辑调整
	// 为了测试，我们假设时间范围正好是一周
	OptimizeTimeRange(startTime, endTime, ITimeRangeOptimizer{
		IsDay: func(day time.Time) {
			isDay = true
		},
		IsWeek: func(week time.Time) {
			isWeek = true
			if week != startTime {
				t.Errorf("IsWeek: expected %v, got %v", startTime, week)
			}
		},
		IsMonth: func(month time.Time) {
			isMonth = true
		},
		IsRange: func(start, end time.Time) {
			isRange = true
		},
	})

	if isDay {
		t.Errorf("expected isDay to be false, got true")
	}
	if !isWeek {
		t.Errorf("expected isWeek to be true, got false")
	}
	if isMonth {
		t.Errorf("expected isMonth to be false, got true")
	}
	if isRange {
		t.Errorf("expected isRange to be false, got true")
	}

	// isMonth
	startTime = StartOfDay(time.Date(2026, 1, 1, 10, 0, 0, 0, loc))
	endTime = EndOfDay(time.Date(2026, 1, 31, 12, 0, 0, 0, loc))
	isDay = false
	isWeek = false
	isMonth = false
	isRange = false

	// 注意：这里需要根据实际的 OptimizeTimeRange 逻辑调整
	// 为了测试，我们假设时间范围正好是一个月
	OptimizeTimeRange(startTime, endTime, ITimeRangeOptimizer{
		IsDay: func(day time.Time) {
			isDay = true
		},
		IsWeek: func(week time.Time) {
			isWeek = true
		},
		IsMonth: func(month time.Time) {
			isMonth = true
			if month != startTime {
				t.Errorf("IsMonth: expected %v, got %v", startTime, month)
			}
		},
		IsRange: func(start, end time.Time) {
			isRange = true
		},
	})

	if isDay {
		t.Errorf("expected isDay to be false, got true")
	}
	if isWeek {
		t.Errorf("expected isWeek to be false, got true")
	}
	if !isMonth {
		t.Errorf("expected isMonth to be true, got false")
	}
	if isRange {
		t.Errorf("expected isRange to be false, got true")
	}
}
