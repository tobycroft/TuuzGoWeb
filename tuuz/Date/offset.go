package Date

import (
	"time"
)

func Date_offset_month1st(month_offset int) time.Time {
	month := MonthFunction{TheTime: time.Now(), KeepCurrentSecond: false}
	return month.OffsetFirstDayInMonth(month_offset)
}

func Date_offset_month_todayWithTimeZero(month_offset int) time.Time {
	month := MonthFunction{TheTime: time.Now(), KeepCurrentSecond: false}
	return month.OffsetCurrentDayInMonth(month_offset)
}

func Date_offset_month_withCurrentTime(month_offset int) time.Time {
	month := MonthFunction{TheTime: time.Now(), KeepCurrentSecond: true}
	return month.OffsetCurrentDayInMonth(month_offset)
}

func Date_offset_thisWeek(percise_to_second bool) time.Time {
	week := WeekFunction{TheTime: time.Now(), KeepCurrentSecond: percise_to_second}
	return week.GetFirstDay()
}

func Date_offset_week1st(week_offset int) time.Time {
	week := WeekFunction{TheTime: time.Now(), KeepCurrentSecond: false}
	return week.OffsetFirstDayInWeek(week_offset)
}

func Date_offset_dayZero(day_offset int) time.Time {
	day := DayFunction{TheTime: time.Now(), KeepCurrentSecond: false}
	return day.OffsetDayInZero(day_offset)
}
