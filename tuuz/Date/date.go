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
