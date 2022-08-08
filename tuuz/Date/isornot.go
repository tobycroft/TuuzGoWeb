package Date

import (
	"time"
)

func Date_between_time(the_time time.Time, start_time time.Time, end_time time.Time) bool {
	if the_time.Before(start_time) {
		return false
	}
	if the_time.After(end_time) {
		return false
	}
	return true
}

func Date_is_in_this_Month(the_time time.Time) bool {
	return Date_is_in_offset_free_Month(the_time, 0, 1)
}

func Date_is_in_offset_Month(the_time time.Time, offset int) bool {
	return Date_is_in_offset_free_Month(the_time, offset, offset+1)
}

func Date_is_in_offset_free_Month(the_time time.Time, offset_start, offset_end int) bool {
	this_month := Date_offset_month1st(offset_start)
	next_month := Date_offset_month1st(offset_end)
	return Date_between_time(the_time, this_month, next_month)
}

func Date_is_in_this_Week(the_time time.Time) bool {
	return Date_is_in_offset_free_Week(the_time, 0, 1)
}

func Date_is_in_offset_Week(the_time time.Time, offset int) bool {
	return Date_is_in_offset_free_Week(the_time, offset, offset+1)
}

func Date_is_in_offset_free_Week(the_time time.Time, offset_start, offset_end int) bool {
	this_week := Date_offset_week1st(offset_start)
	next_week := Date_offset_week1st(offset_end)
	return Date_between_time(the_time, this_week, next_week)
}

func Date_is_in_today(the_time time.Time) bool {
	return Date_is_in_offset_free_Day(the_time, 0, 1)
}

func Date_is_in_offset_Day(the_time time.Time, offset int) bool {
	return Date_is_in_offset_free_Day(the_time, offset, offset+1)
}

func Date_is_in_offset_free_Day(the_time time.Time, offset_start, offset_end int) bool {
	today := Date_offset_dayZero(offset_start)
	tomorrow := Date_offset_dayZero(offset_end)
	return Date_between_time(the_time, today, tomorrow)
}
