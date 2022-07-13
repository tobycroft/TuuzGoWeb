package Array

import "github.com/tobycroft/Calc"

func MapAny2MapString[T int | string | int64 | float64, V string | int64 | float64 | int | any](maps map[T]V) map[T]string {
	strs := make(map[T]string)
	for key, value := range maps {
		strs[key] = Calc.Any2String(value)
	}
	return strs
}

func MapAny2MapInterface[T int | string | int64 | float64, V string | int64 | float64 | int | any](maps map[T]V) map[T]interface{} {
	arr := make(map[T]interface{})
	for key, value := range maps {
		arr[key] = value
	}
	return arr
}
