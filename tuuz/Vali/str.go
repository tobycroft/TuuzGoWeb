package Vali

import (
	"fmt"
	"math/big"
	"reflect"
)

func any2string(any interface{}) string {
	var str string
	switch any.(type) {
	case string:
		str = any.(string)

	case int:
		tmp := any.(int)
		str = int2String(tmp)

	case int32:
		tmp := int64(any.(int32))
		str = int642String(tmp)

	case int64:
		tmp := any.(int64)
		str = int642String(tmp)

	case float64:
		tmp := any.(float64)
		str = float642String(tmp)

	case float32:
		tmp := float64(any.(float32))
		str = float642String(tmp)

	case *big.Int:
		tmp := any.(*big.Int)
		str = tmp.String()

	case nil:
		str = ""

	case bool:
		tmp := any.(bool)
		if tmp == true {
			return "true"
		} else {
			return "false"
		}

	default:
		fmt.Println("any2string", any, reflect.TypeOf(any))
		str = ""
	}
	return str
}
