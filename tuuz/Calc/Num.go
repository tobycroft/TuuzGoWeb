package Calc

func Max_from_string(str []string) float64 {
	max := float64(0)
	lock := false
	for _, v := range str {
		f, err := String2Float64(v)
		if err != nil {

		} else {
			if f > max || lock == false {
				max = f
				lock = true
			}
		}
	}
	return max
}

func Min_from_string(str []string) float64 {
	min := float64(0)
	lock := false
	for _, v := range str {
		f, err := String2Float64(v)
		if err != nil {

		} else {
			if f < min || lock == false {
				min = f
				lock = true
			}
		}
	}
	return min
}
