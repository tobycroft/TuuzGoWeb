package Array

import "github.com/shopspring/decimal"

func ArrayKeyExists[T string | int | int32 | int64 | float32 | float64 | decimal.Decimal, V string | int | int32 | int64 | float32 | float64 | decimal.Decimal | any](key T, m map[T]V) bool {
	_, ok := m[key]
	return ok
}

func ArrayKey[T string | int | int32 | int64 | float32 | float64 | decimal.Decimal, V string | int | int32 | int64 | float32 | float64 | decimal.Decimal | any](m map[T]V) []T {
	keys := []T{}
	for t := range m {
		keys = append(keys, t)
	}
	return keys
}

func ArrayVal[T string | int | int32 | int64 | float32 | float64 | decimal.Decimal, V string | int | int32 | int64 | float32 | float64 | decimal.Decimal | any](m map[T]V) []V {
	vals := []V{}
	for _, v := range m {
		vals = append(vals, v)
	}
	return vals
}
