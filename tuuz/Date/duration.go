package Date

import "time"

func Date_between_hours(datetime_start, datetime_end time.Time) float64 {
	if datetime_start.Before(datetime_end) {
		return datetime_end.Sub(datetime_start).Hours()
	} else {
		return datetime_start.Sub(datetime_end).Hours()
	}
}

func Date_between_hours_int64(datetime_start, datetime_end time.Time) int64 {
	if datetime_start.Before(datetime_end) {
		return int64(datetime_end.Sub(datetime_start).Hours())
	} else {
		return int64(datetime_start.Sub(datetime_end).Hours())
	}
}

func Date_between_days(datetime_start, datetime_end time.Time) int64 {
	return int64(Date_between_hours(datetime_start, datetime_end) / 24)
}
