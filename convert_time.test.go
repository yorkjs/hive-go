package hive

import (
	"testing"
	"time"
)

func TestTimeConvert(t *testing.T) {

	date := time.Now()
	timestamp := TimeToTimestamp(date)

	if TimeToTimestamp(date) != timestamp {
		t.Errorf("timeToTimestamp 失败: 期望 %d, 得到 %d", timestamp, TimeToTimestamp(date))
	}

	resultTime := TimestampToTime(timestamp)
	if TimeToTimestamp(resultTime) != timestamp {
		t.Errorf("timestampToTime 失败: 期望 %d, 得到 %d", timestamp, TimeToTimestamp(resultTime))
	}

	// 测试无效日期
	zeroTime, err := StringToTime("-", DATE_YEAR_MONTH_DATE)
	if err != nil {
		t.Fatalf("解析无效日期失败: %v", err)
	}
	if TimeToTimestamp(zeroTime) != 0 {
		t.Errorf("无效日期转换失败: 期望 0, 得到 %d", TimeToTimestamp(zeroTime))
	}

	time1, err := StringToTime(
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
	time2, err := StringToTime(
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
	time3, err := StringToTime(
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

	time1, err = StringToTime(
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
	time2, err = StringToTime(
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
	time3, err = StringToTime(
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

	time1, err = StringToTime(
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
	time2, err = StringToTime(
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
	time3, err = StringToTime(
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
