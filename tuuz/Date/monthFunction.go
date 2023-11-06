package Date

import "time"

type MonthFunction struct {
	TheTime           time.Time
	KeepCurrentSecond bool
}

// GetFirstDay this will return the first day of the week
func (this *MonthFunction) GetFirstDay() time.Time {
	year, month, _ := this.TheTime.Date()
	if this.KeepCurrentSecond {
		return time.Date(year, month, 1, this.TheTime.Hour(), this.TheTime.Minute(), this.TheTime.Second(), this.TheTime.Nanosecond(), time.Local)
	} else {
		return time.Date(year, month, 1, 0, 0, 0, 0, time.Local)
	}
}

func (this *MonthFunction) OffsetFirstDayInMonth(month_offset int) time.Time {
	return this.GetFirstDay().AddDate(0, month_offset, 0)
}

func (this *MonthFunction) OffsetCurrentDayInMonth(month_offset int) time.Time {
	return this.TheTime.AddDate(0, month_offset, 0)
}
