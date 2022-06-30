package Date

import "time"

func Offset_month1st(month_offset int) time.Time {
	year, month, _ := time.Now().Date()
	thisMonth := time.Date(year, month, 1, 0, 0, 0, 0, time.Local)
	return thisMonth.AddDate(0, month_offset, 0)
}

func Offset_month_withTimeZero(month_offset int) time.Time {
	year, month, day := time.Now().Date()
	thisMonth := time.Date(year, month, day, 0, 0, 0, 0, time.Local)
	return thisMonth.AddDate(0, month_offset, 0)
}

func Offset_month_withCurrentTime(month_offset int) time.Time {
	return time.Now().AddDate(0, month_offset, 0)
}

func Offset_thisWeek() time.Time {
	now := time.Now()
	offset := int(time.Monday - now.Weekday())
	if offset > 0 {
		offset = -6
	}
	weekStartDate := time.Date(now.Year(), now.Month(), now.Day(), 0, 0, 0, 0, time.Local).AddDate(0, 0, offset)
	return weekStartDate
}

func Offset_week1st(week_offset int) time.Time {
	return Offset_thisWeek().AddDate(0, 0, -7*week_offset)
}

func Offset_format_year(T time.Time) string {
	return T.Format("2006")
}

func Offset_format_month(T time.Time) string {
	return T.Format("2006-01")
}

func Offset_format_day(T time.Time) string {
	return T.Format("2006-01-02")
}

func Offset_format_hour(T time.Time) string {
	return T.Format("2006-01-02 15")
}

func Offset_format_minute(T time.Time) string {
	return T.Format("2006-01-02 15:04")
}

func Offset_format_second(T time.Time) string {
	return T.Format("2006-01-02 15:04:05")
}
