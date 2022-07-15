package Calc

import (
	"errors"
	"fmt"
	"github.com/shopspring/decimal"
	"math"
	"math/big"
	"strconv"
)

func Gene2Float64[T string | int | int32 | int64 | float32 | float64 | decimal.Decimal | *big.Int](v T) (float64, error) {
	switch any(v).(type) {
	case string:
		float, err := strconv.ParseFloat(any(v).(string), 64)
		return float, err

	case int:
		return decimal.NewFromInt(int64(any(v).(int))).InexactFloat64(), nil

	case int64:
		return decimal.NewFromInt(any(v).(int64)).InexactFloat64(), nil

	case float32:
		return strconv.ParseFloat(fmt.Sprint(v), 64)

	case float64:
		return any(v).(float64), nil

	case decimal.Decimal:
		return any(v).(decimal.Decimal).InexactFloat64(), nil

	case *big.Int:
		return float64(any(v).(*big.Int).Int64()), nil
	}
	return 0, errors.New("unable to match")
}

func Gene2Int64[T string | int | int32 | int64 | float32 | float64 | decimal.Decimal | *big.Int](v T) (int64, error) {
	switch any(v).(type) {
	case string:
		return strconv.ParseInt(any(v).(string), 10, 64)

	case int:
		return int64(any(v).(int)), nil

	case int64:
		return any(v).(int64), nil

	case float32:
		return int64(math.Round(float64(any(v).(float32)))), nil

	case float64:
		return int64(math.Round(any(v).(float64))), nil

	case decimal.Decimal:
		return any(v).(decimal.Decimal).IntPart(), nil

	case *big.Int:
		return any(v).(*big.Int).Int64(), nil
	}
	return 0, errors.New("unable to match")
}

func Gene2Int[T string | int | int32 | int64 | float32 | float64 | decimal.Decimal | *big.Int](v T) (int, error) {
	ret, err := Gene2Int64(v)
	if err != nil {
		return 0, err
	}
	if math.MaxInt > ret {
		return int(ret), nil
	} else {
		return 0, errors.New("too large for int")
	}
}

func Gene2String[T string | int | int32 | int64 | float64](v T) (string, error) {
	switch any(v).(type) {
	case string:
		return any(v).(string), nil

	case int:
		return strconv.Itoa(any(v).(int)), nil

	case int64:
		return strconv.FormatInt(any(v).(int64), 10), nil

	case float64:
		return strconv.FormatFloat(any(v).(float64), 'f', -1, 64), nil

	case decimal.Decimal:
		return any(v).(decimal.Decimal).String(), nil

	case *big.Int:
		return any(v).(*big.Int).String(), nil
	}
	return "", errors.New("unable to match")
}
