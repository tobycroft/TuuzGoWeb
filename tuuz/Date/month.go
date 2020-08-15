package Date

import "time"

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
