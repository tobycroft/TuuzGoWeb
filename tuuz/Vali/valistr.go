package Vali

import (
	"regexp"
)

func Length(str string, minlen int, maxlen int) (bool, string) {
	if len(str) > maxlen {
		return false, "长度需要小于" + any2string(maxlen)
	}
	if len(str) < minlen {
		return false, "长度需要大于" + any2string(minlen)
	}
	return true, ""
}

func Complex(str string, need_number bool, need_letter, need_upcase, need_lowercase bool) (bool, string) {
	if need_number {
		number := `[0-9]`
		number_x := regexp.MustCompile(number)
		if !number_x.MatchString(str) {
			return false, "\"" + str + "\"" + "需要包含字符"
		}
	}
	if need_letter {
		if need_upcase {
			up := `[A-Z]`
			up_x := regexp.MustCompile(up)
			if !up_x.MatchString(str) {
				return false, "\"" + str + "\"" + "需要包含小写字符"
			}
		}
		if need_lowercase {
			low := `[a-z]`
			low_x := regexp.MustCompile(low)
			if !low_x.MatchString(str) {
				return false, "\"" + str + "\"" + "需要包含大写字符"
			}
		}
	}
	return true, ""
}
