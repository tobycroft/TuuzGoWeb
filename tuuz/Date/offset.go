package Date

import (
	"time"
)

func Date_offset_month1st(month_offset int) time.Time {
	year, month, _ := time.Now().Date()
	thisMonth := time.Date(year, month, 1, 0, 0, 0, 0, time.Local)
	return thisMonth.AddDate(0, month_offset, 0)
}

func Date_offset_month_todayWithTimeZero(month_offset int) time.Time {
	year, month, day := time.Now().Date()
	thisMonth := time.Date(year, month, day, 0, 0, 0, 0, time.Local)
	return thisMonth.AddDate(0, month_offset, 0)
}

func Date_offset_month_withCurrentTime(month_offset int) time.Time {
	return time.Now().AddDate(0, month_offset, 0)
}

func Date_offset_thisWeek(need_now bool) time.Time {
	now := time.Now()
	offset := int(time.Monday - now.Weekday())
	if offset > 0 {
		offset = -6
	}
	if need_now {
		return time.Date(now.Year(), now.Month(), now.Day(), now.Hour(), now.Minute(), now.Second(), now.Nanosecond(), time.Local).AddDate(0, 0, offset)
	} else {
		return time.Date(now.Year(), now.Month(), now.Day(), 0, 0, 0, 0, time.Local).AddDate(0, 0, offset)
	}
}

func Date_offset_week1st(week_offset int) time.Time {
	return Date_offset_thisWeek(false).AddDate(0, 0, 7*week_offset)
}

func Date_offset_dayZero(day_offset int) time.Time {
	now := time.Now()
	return time.Date(now.Year(), now.Month(), now.Day(), 0, 0, 0, 0, time.Local).AddDate(0, 0, day_offset)
}
