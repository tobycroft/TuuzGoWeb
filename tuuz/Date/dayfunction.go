package Date

import "time"

type DayFunction struct {
	TheTime           time.Time
	KeepCurrentSecond bool
}

func (this *DayFunction) OffsetDayInZero(day_offset int) time.Time {
	now := this.TheTime
	if this.KeepCurrentSecond {
		return now.AddDate(0, 0, day_offset)
	} else {
		return time.Date(now.Year(), now.Month(), now.Day(), 0, 0, 0, 0, time.Local).AddDate(0, 0, day_offset)
	}
}

func (this *DayFunction) GetDayZero() time.Time {
	return time.Date(this.TheTime.Year(), this.TheTime.Month(), this.TheTime.Day(), 0, 0, 0, 0, time.Local)
}
