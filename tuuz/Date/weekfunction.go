package Date

import "time"

type WeekFunction struct {
	TheTime           time.Time
	KeepCurrentSecond bool
}

// GetFirstDay this will return the first day of the week
func (this *WeekFunction) GetFirstDay() time.Time {
	now := this.TheTime
	offset := int(time.Monday - now.Weekday())
	if offset > 0 {
		offset = -6
	}
	if this.KeepCurrentSecond {
		return time.Date(now.Year(), now.Month(), now.Day(), now.Hour(), now.Minute(), now.Second(), now.Nanosecond(), time.Local).AddDate(0, 0, offset)
	} else {
		return time.Date(now.Year(), now.Month(), now.Day(), 0, 0, 0, 0, time.Local).AddDate(0, 0, offset)
	}
}

// Date_offset_week1st 0 means the first day of the week,plus 1 will return the first day of the next week,minus 1 will return the first day of the previous week
func (this *WeekFunction) OffsetFirstDayInWeek(week_offset int) time.Time {
	return this.GetFirstDay().AddDate(0, 0, 7*week_offset)
}
