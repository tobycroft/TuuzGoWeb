package Array

import (
	"github.com/shopspring/decimal"
	"github.com/tobycroft/Calc"
	"strings"
)

func Join[T int | int64 | float64 | float32 | decimal.Decimal | string | interface{}](slices []T) string {
	strs := []string{}
	for _, slice := range slices {
		strs = append(strs, Calc.Any2String(slice))
	}
	return strings.Join(strs, ",")
}
