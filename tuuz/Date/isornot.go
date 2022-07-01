package Date

import "time"

func Date_between_time(the_time time.Time, start_time time.Time, end_time time.Time) bool {
	if the_time.Before(start_time) {
		return false
	}
	if the_time.After(end_time) {
		return false
	}
	return true
}

func Date_is_in_thisMonth(the_time time.Time) bool {
	this_month := Date_offset_month1st(0)
	next_month := Date_offset_month1st(1)
	return Date_between_time(the_time, this_month, next_month)
}

func Date_is_in_thisWeek(the_time time.Time) bool {
	this_week := Date_offset_week1st(0)
	next_week := Date_offset_week1st(1)
	return Date_between_time(the_time, this_week, next_week)
}
