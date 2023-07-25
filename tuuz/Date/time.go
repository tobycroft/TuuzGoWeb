package Date

import (
	"main.go/config/app_conf"
	"regexp"
	"strings"
	"time"
)

func Date_time_parser(date_time string, location *time.Location) (p time.Time, err error) {
	p, err = date_parse_time(date_time)
	if err != nil {
		return
	}
	p = date_parse_offset(p, location)
	return
}

func date_parse_time(date_time string) (p time.Time, err error) {
	if strings.Contains(date_time, "T") {
		if strings.Contains(date_time, ".") {
			return time.Parse(time.RFC3339Nano, date_time)
		} else {
			return time.Parse(time.RFC3339, date_time)
		}
	} else {
		var datetime_exp *regexp.Regexp
		datetime_exp, err = regexp.Compile(`[2][0-9][0-9][0-9]\-[0-9]+\-[0-9]+[ ][0-9]+\:[0-9]+\:[0-9]+`)
		if err != nil {
			return
		}
		if datetime_exp.MatchString(date_time) {
			return time.Parse("2006-1-2 15:4:5", date_time)
		} else {
			return time.Parse("2006-1-2", date_time)
		}
	}
}

func date_parse_offset(p time.Time, location *time.Location) time.Time {
	if location != nil {
		p = p.In(location)
	} else {
		loc, err := time.LoadLocation(app_conf.TimeZoneLocation)
		if err != nil {
			loc = app_conf.TimeZone
		}
		p = p.In(loc)
	}
	return p
}
