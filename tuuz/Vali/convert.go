package Vali

import "strconv"

func int2String(num int) string {
	return strconv.Itoa(num)
}

func int642String(num int64) string {
	return strconv.FormatInt(num, 10)
}

func float642String(f64 float64) string {
	return strconv.FormatFloat(f64, 'f', -1, 64)
}
