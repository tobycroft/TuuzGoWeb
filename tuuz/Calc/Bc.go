package Calc

import (
	"github.com/shopspring/decimal"
)

func todecimal(number interface{}) decimal.Decimal {
	switch number.(type) {
	case int:
		return decimal.NewFromInt(int64(number.(int)))

	case int64:
		return decimal.NewFromInt(number.(int64))

	case float64:
		return decimal.NewFromFloat(number.(float64))

	case float32:
		return decimal.NewFromFloat32(number.(float32))

	case int32:
		return decimal.NewFromInt32(number.(int32))

	case string:
		ret, err := decimal.NewFromString(number.(string))
		if err != nil {
			return decimal.NewFromInt(0)
		}
		return ret

	case decimal.Decimal:
		return number.(decimal.Decimal)

	default:
		return decimal.NewFromFloat(Any2Float64(number))
	}
}

func Bc_add(num1, num2 interface{}) decimal.Decimal {
	return Bc_sum(num1, num2)
}

func Bc_sum(num1, num2 interface{}) decimal.Decimal {
	return todecimal(num1).Add(todecimal(num2))
}

func Bc_min(num1, num2 interface{}) decimal.Decimal {
	return todecimal(num1).Sub(todecimal(num2))
}

func Bc_mul(num1, num2 interface{}) decimal.Decimal {
	return todecimal(num1).Mul(todecimal(num2))
}

func Bc_div(num1, num2 interface{}) decimal.Decimal {
	return todecimal(num1).Div(todecimal(num2))
}

func Bc_pow(num1, num2 interface{}) decimal.Decimal {
	return todecimal(num1).Pow(todecimal(num2))
}

func Bc_round(num1 interface{}, round int) decimal.Decimal {
	return todecimal(num1).Round(int32(round))
}

func Bc_div_round(num1 interface{}, num2 interface{}, round int) decimal.Decimal {
	return todecimal(num1).DivRound(todecimal(num2), int32(round))
}

func Bc_abs(num1 interface{}) decimal.Decimal {
	return todecimal(num1).Abs()
}

func Bc_mod(num1, num2 interface{}) decimal.Decimal {
	return todecimal(num1).Mod(todecimal(num2))
}
