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

func NextMonth_1st() string {
	year, month, _ := time.Now().Date()
	thisMonth := time.Date(year, month, 1, 0, 0, 0, 0, time.Local)
	monthOneDay := thisMonth.AddDate(0, 1, 0).Format("2006-01-02")
	return monthOneDay
}

func LastMonthCombine() string {
	year, month, _ := time.Now().Date()
	thisMonth := time.Date(year, month, 1, 0, 0, 0, 0, time.Local)
	monthOneDay := thisMonth.AddDate(0, -1, 0).Format("200601")
	return monthOneDay
}

func ThisMonth1st() string {
	year, month, _ := time.Now().Date()
	thisMonth := time.Date(year, month, 1, 0, 0, 0, 0, time.Local)
	monthOneDay := thisMonth.Format("2006-01-02")
	return monthOneDay
}

func ThisMonth1st_int() int64 {
	return Date2Int(ThisMonth1st())
}

func LastMonth1st() string {
	year, month, _ := time.Now().Date()
	thisMonth := time.Date(year, month, 1, 0, 0, 0, 0, time.Local)
	monthOneDay := thisMonth.AddDate(0, -1, 0).Format("2006-01-02")
	return monthOneDay
}

func LastMonth1st_int() int64 {
	return Date2Int(LastMonth1st())
}

func ThisWeek() string {
	now := time.Now()
	offset := int(time.Monday - now.Weekday())
	if offset > 0 {
		offset = -6
	}
	weekStartDate := time.Date(now.Year(), now.Month(), now.Day(), 0, 0, 0, 0, time.Local).AddDate(0, 0, offset)
	weekMonday := weekStartDate.Format("2006-01-02")
	return weekMonday
}

func ThisWeek_int() int64 {
	return Date2Int(ThisWeek())
}

func LastWeek() string {
	thisWeekMonday := ThisWeek()
	TimeMonday, _ := time.Parse("2006-01-02", thisWeekMonday)
	lastWeekMonday := TimeMonday.AddDate(0, 0, -7)
	weekMonday := lastWeekMonday.Format("2006-01-02")
	return weekMonday
}

func NextWeek() string {
	thisWeekMonday := ThisWeek()
	TimeMonday, _ := time.Parse("2006-01-02", thisWeekMonday)
	lastWeekMonday := TimeMonday.AddDate(0, 0, 7)
	weekMonday := lastWeekMonday.Format("2006-01-02")
	return weekMonday
}

func LastWeek_int() int64 {
	return Date2Int(LastWeek())
}

func Today() string {
	monthOneDay := time.Now().Format("2006-01-02")
	return monthOneDay
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
