package Date

import "time"

func Date_format_year(T time.Time) string {
	return T.Format("2006")
}

func Date_format_month(T time.Time) string {
	return T.Format("2006-01")
}

func Date_format_day(T time.Time) string {
	return T.Format("2006-01-02")
}

func Date_format_hour(T time.Time) string {
	return T.Format("2006-01-02 15")
}

func Date_format_minute(T time.Time) string {
	return T.Format("2006-01-02 15:04")
}

func Date_format_second(T time.Time) string {
	return T.Format("2006-01-02 15:04:05")
}
