package Date

import "time"

func ThisYear() string {
	return time.Now().Format("2006")
}

func ThisMonth() string {
	return time.Now().Format("2006-01")
}

func ThisMonthCombine() string {
	return time.Now().Format("200601")
}

func TodayCombine() string {
	return time.Now().Format("20060102")
}

func LastMonth() string {
	year, month, _ := time.Now().Date()
	thisMonth := time.Date(year, month, 1, 0, 0, 0, 0, time.Local)
	monthOneDay := thisMonth.AddDate(0, -1, 0).Format("2006-01")
	return monthOneDay
}

func NextMonth() string {
	year, month, _ := time.Now().Date()
	thisMonth := time.Date(year, month, 1, 0, 0, 0, 0, time.Local)
	monthOneDay := thisMonth.AddDate(0, 1, 0).Format("2006-01")
	return monthOneDay
}

func MonthDateNow(month int) string {
	monthOneDay := time.Now().AddDate(0, month, 0).Format("2006-01-02 15:04:05")
	return monthOneDay
}

func MonthDay(month int) string {
	return Offset_format_day(Offset_month1st(month))
}

func Month(month int) string {
	return Offset_format_month(Offset_month1st(0))
}

func NextMonth_1st() string {
	return Offset_format_second(Offset_month1st(1))
}

func LastMonthCombine() string {
	return Offset_month1st(-1).Format("200601")
}

func ThisMonth1st() string {
	return Offset_format_day(Offset_month1st(0))
}

func ThisMonth1st_int() int64 {
	return Offset_month1st(0).Unix()
}

func LastMonth1st() string {
	return Offset_format_second(Offset_month1st(-1))
}

func LastMonth1st_int() int64 {
	return Offset_month1st(-1).Unix()
}

func ThisWeek() string {
	return Offset_format_day(Offset_thisWeek())
}

func ThisWeek_int() int64 {
	return Offset_thisWeek().Unix()
}

func LastWeek() string {
	return Offset_format_day(Offset_week1st(-1))
}

func NextWeek() string {
	return Offset_format_day(Offset_week1st(1))
}

func LastWeek_int() int64 {
	return Offset_week1st(-1).Unix()
}

func Today() string {
	return Offset_format_second(Offset_dayZero(0))
}

func Tomorrow() string {
	monthOneDay := time.Now().AddDate(0, 0, 1).Format("2006-01-02")
	return monthOneDay
}

func Today_int() int64 {
	return Datetime2Int(Today())
}

func Yesterday() string {
	dat := time.Now().AddDate(0, 0, -1).Format("2006-01-02")
	return dat
}

func Yesterday_int() int64 {
	return Date2Int(Yesterday())
}

func WeekBefore() string {
	now := time.Now()
	offset := int(time.Monday - now.Weekday())
	if offset > 0 {
		offset = -6
	}
	weekStartDate := time.Date(now.Year(), now.Month(), now.Day(), 0, 0, 0, 0, time.Local).AddDate(0, 0, offset).AddDate(0, 0, -7)
	weekMonday := weekStartDate.Format("2006-01-02")
	return weekMonday
}

func WeekBefore_int() int64 {
	return Date2Int(WeekBefore())
}
