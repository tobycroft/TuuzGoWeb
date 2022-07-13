package Array

import "github.com/shopspring/decimal"

func InArray[T string | float64 | float32 | int | int32 | int64 | decimal.Decimal](find_str T, haystack []T) bool {
	for _, t := range haystack {
		if find_str == t {
			return true
		}
	}
	return false
}
