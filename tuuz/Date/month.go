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
	return Date_format_month(Offset_month1st(-1))
}

func NextMonth() string {
	return Date_format_month(Offset_month1st(1))
}

func MonthDateNow(month int) string {
	return Date_format_second(Offset_month_withCurrentTime(month))
}

func MonthDay(month int) string {
	return Date_format_day(Offset_month1st(month))
}

func Month(month int) string {
	return Date_format_month(Offset_month1st(0))
}

func NextMonth_1st() string {
	return Date_format_second(Offset_month1st(1))
}

func LastMonthCombine() string {
	return Offset_month1st(-1).Format("200601")
}

func ThisMonth1st() string {
	return Date_format_day(Offset_month1st(0))
}

func ThisMonth1st_int() int64 {
	return Offset_month1st(0).Unix()
}

func LastMonth1st() string {
	return Date_format_second(Offset_month1st(-1))
}

func LastMonth1st_int() int64 {
	return Offset_month1st(-1).Unix()
}

func ThisWeek() string {
	return Date_format_day(Offset_thisWeek())
}

func ThisWeek_int() int64 {
	return Offset_thisWeek().Unix()
}

func LastWeek() string {
	return Date_format_day(Offset_week1st(-1))
}

func NextWeek() string {
	return Date_format_day(Offset_week1st(1))
}

func LastWeek_int() int64 {
	return Offset_week1st(-1).Unix()
}

func Today() string {
	return Date_format_day(Offset_dayZero(0))
}

func Tomorrow() string {
	return Date_format_day(Offset_dayZero(1))
}

func Today_int() int64 {
	return Offset_dayZero(0).Unix()
}

func Yesterday() string {
	return Date_format_day(Offset_dayZero(-1))
}

func Yesterday_int() int64 {
	return Offset_dayZero(-1).Unix()
}

func WeekBefore() string {
	return Date_format_day(Offset_week1st(-1))
}

func WeekBefore_int() int64 {
	return Offset_week1st(-1).Unix()
}
