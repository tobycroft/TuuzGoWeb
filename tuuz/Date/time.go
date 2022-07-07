package Date

import (
	"regexp"
	"strings"
	"time"
)

func Date_time_parser(date_time string) (p time.Time, err error) {
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
