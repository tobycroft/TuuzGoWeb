package Calc

import (
	"fmt"
	"github.com/shopspring/decimal"
	"math/big"
	"reflect"
	"strconv"
	"strings"
)

func Chop(s string, character_mask string) string {
	return strings.TrimRight(s, character_mask)
}

//func Transform[T int | int64 | string | interface{}, K int | int64 | string | interface{}](any T) K{
//
//}

func Any2String(any interface{}) string {
	var str string
	switch any.(type) {

	case int64:
		tmp := any.(int64)
		str = Int642String(tmp)

	case float64:
		tmp := any.(float64)
		str = Float642String(tmp)

	case bool:
		tmp := any.(bool)
		if tmp == true {
			return "true"
		} else {
			return "false"
		}

	case string:
		str = any.(string)

	case nil:
		str = ""

	case int:
		tmp := any.(int)
		str = Int2String(tmp)

	case int32:
		tmp := int64(any.(int32))
		str = Int642String(tmp)

	case float32:
		tmp := Float322Float64(any.(float32))
		str = Float642String(tmp)

	case *big.Int:
		tmp := any.(*big.Int)
		str = tmp.String()

	case decimal.Decimal:
		tmp := any.(decimal.Decimal)
		str = tmp.String()

	default:
		fmt.Println("any2string", any, reflect.TypeOf(any))
		str = ""
	}
	return str
}

func String2Int(str string) (int, error) {
	return strconv.Atoi(str)
}

func String2Int64(str string) (int64, error) {
	return strconv.ParseInt(str, 10, 64)
}

func String2Float64(str string) (float64, error) {
	float, err := strconv.ParseFloat(str, 64)
	return float, err
}

func PhoneSafe(phone string) string {
	old := ""
	for k, v := range phone {
		if k >= 3 && k <= 6 {
			old = old + string(v)
		}
	}
	phone = strings.Replace(phone, old, "****", -1)
	return phone
}

func Interface2String(inter []interface{}) []string {
	strs := []string{}
	for it := range inter {
		strs = append(strs, Any2String(it))
	}
	return strs
}

func AnyJoin[T int | int64 | float64 | decimal.Decimal](joins []T, sep string) string {
	strs := []string{}
	for _, join := range joins {
		strs = append(strs, Any2String(join))
	}
	return strings.Join(strs, sep)
}
