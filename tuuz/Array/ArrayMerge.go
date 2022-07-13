package Array

import "github.com/shopspring/decimal"

func Merge[T string | int64 | float64, K string | interface{} | int64 | float64 | decimal.Decimal](args ...map[T]K) map[T]K {
	arr := make(map[T]K)
	for _, arrs := range args {
		for key, value := range arrs {
			arr[key] = value
		}
	}
	return arr
}
