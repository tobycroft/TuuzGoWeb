package Array

import "github.com/shopspring/decimal"

func ArrayUnique[T string | int | int32 | int64 | float32 | float64 | byte | decimal.Decimal](slice []T) []T {
	n64 := []T{}
	for _, s1 := range slice {
		temp := true
		for _, s2 := range n64 {
			if s2 == s1 {
				temp = false
				break
			}
		}
		if temp {
			n64 = append(n64, s1)
		}
	}
	return n64
}

func ArrayDiff[T string | int | int32 | int64 | float32 | float64 | byte | decimal.Decimal | any](slice, slice2 []T) []T {
	n64 := []T{}
	for _, s1 := range slice {
		temp := true
		for _, s2 := range slice2 {
			if any(s2) == any(s1) {
				temp = false
				break
			}
		}
		if temp {
			n64 = append(n64, s1)
		}
	}
	return n64
}

func ArrayTrim[T string | int | int32 | int64 | float32 | float64 | byte | decimal.Decimal | any](slice []T, exp T) []T {
	n64 := []T{}
	for _, s1 := range slice {
		temp := true
		if any(exp) == any(s1) {
			temp = false
			break
		}
		if temp {
			n64 = append(n64, s1)
		}
	}
	return n64
}
