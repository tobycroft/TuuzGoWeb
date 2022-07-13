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

func ArrayDiff[T string | int | int32 | int64 | float32 | float64 | byte | decimal.Decimal](slice, slice2 []T) []T {
	n64 := []T{}
	for _, s1 := range slice {
		temp := true
		for _, s2 := range slice2 {
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
