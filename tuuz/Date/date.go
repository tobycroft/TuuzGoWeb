package Date

import (
	"time"
)

func Date2Int(date string) int64 {
	p, err := time.Parse("2006-01-02 15:04:05", date)
	if err != nil {
		return 0
	} else {
		return p.Unix()
	}
}

func Datetime2Int(date string) int64 {
	p, err := time.Parse("2006-01-02", date)
	if err != nil {
		return 0
	} else {
		return p.Unix()
	}
}

func Int2Date(t int64) string {
	timing := time.Unix(t, 0)
	return timing.Format("2006-01-02")
}

func Int2Datetime(t int64) string {
	timing := time.Unix(t, 0)
	return timing.Format("2006-01-02 15:04:05")
}

func Time2Datetime(t time.Time) string {
	return t.Format("2006-01-02 15:04:05")
}

func Time2Int64(t time.Time) int64 {
	return t.Unix()
}

func Datetime2Date(datetime string) string {
	p, err := time.Parse("2006-01-02 15:04:05", datetime)
	if err != nil {
		return ""
	} else {
		return p.Format("2006-01-02")
	}
}

func Date2Time(year int, month time.Month, day int) time.Time {
	return time.Date(year, month, day, 0, 0, 0, 0, time.Local)
}

func Date2DateTime(year int, month time.Month, day int, hour, min, sec int) time.Time {
	return time.Date(year, month, day, hour, min, sec, 0, time.Local)
}

func ThisYear2Time(year int) time.Time {
	now := time.Now()
	return time.Date(year, now.Month(), now.Day(), now.Hour(), now.Minute(), now.Second(), now.Nanosecond(), time.Local)
}

func YearMonth2Time(year int, month time.Month) time.Time {
	return time.Date(year, month, 1, 0, 0, 0, 0, time.Local)
}
