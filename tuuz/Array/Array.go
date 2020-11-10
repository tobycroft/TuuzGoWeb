package Array

import (
	"main.go/tuuz/Calc"
	"main.go/tuuz/Jsong"
)

func Merge(args ...map[string]interface{}) map[string]interface{} {
	arr := make(map[string]interface{})
	for _, arrs := range args {
		for key, value := range arrs {
			arr[key] = value
		}
	}
	return arr
}

func Mapinterface2MapString(maps map[string]interface{}) map[string]string {
	strs := make(map[string]string)
	for key, value := range maps {
		switch value.(type) {
		case string:
			strs[key] = value.(string)
			break
		case int:
			tmp := value.(int)
			strs[key] = Calc.Int2String(tmp)
			break
		case int64:
			tmp := value.(int64)
			strs[key] = Calc.Int642String(tmp)
			break
		case float64:
			tmp := value.(float64)
			strs[key] = Calc.Float642String(tmp)
			break

		case float32:
			tmp := value.(float64)
			strs[key] = Calc.Float642String(tmp)
			break

		default:
			strs[key] = value.(string)
			break
		}

	}
	return strs
}

func MapString2MapInterface(maps map[string]string) map[string]interface{} {
	arr := make(map[string]interface{})
	for key, value := range maps {
		arr[key] = value
	}
	return arr
}

func MapString2Interface(maps map[string]interface{}) map[string]interface{} {
	strs := make(map[string]interface{})
	for key, value := range maps {
		switch value.(type) {
		case string:
			strs[key] = value.(string)
		case int:
			tmp := value.(int)
			strs[key] = Calc.Int2String(tmp)
		case int64:
			tmp := value.(int64)
			strs[key] = Calc.Int642String(tmp)
		}
	}
	return strs
}
func InArray(str interface{}, haystack []interface{}) bool {
	str, _ = Jsong.Encode(str)
	for _, v := range haystack {
		v, _ = Jsong.Encode(v)
		if str == v {
			return true
		}
	}
	return false
}

func InArrayF64(str float64, haystack []float64) bool {
	for _, v := range haystack {
		if str == v {
			return true
		}
	}
	return false
}

func InArrayInt(str int, haystack []int) bool {
	for _, v := range haystack {
		if str == v {
			return true
		}
	}
	return false
}

func InArrayInt64(str int64, haystack []int64) bool {
	for _, v := range haystack {
		if str == v {
			return true
		}
	}
	return false
}

func InArrayString(str string, haystack []string) bool {
	for _, v := range haystack {
		if str == v {
			return true
		}
	}
	return false
}

func ArrayKeyExists(key interface{}, m map[interface{}]interface{}) bool {
	_, ok := m[key]
	return ok
}

func ArrayUnique_int64(slice []int64) []int64 {
	n64 := []int64{}
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

func ArrayUnique_float64(slice []float64) []float64 {
	n64 := []float64{}
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

func ArrayUnique_string(slice []string) []string {
	n64 := []string{}
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

func Array_find_max_int64(slice []int64) int64 {
	if len(slice) < 1 {
		return 0
	}
	maxVal := slice[0]
	maxIndex := 0
	for i := 1; i < len(slice); i++ {
		//从第二个 元素开始循环比较，如果发现有更大的，则交换
		if maxVal < slice[i] {
			maxVal = slice[i]
			maxIndex = i
		}
	}
	return slice[maxIndex]
}

func Array_find_max_float64(slice []float64) float64 {
	if len(slice) < 1 {
		return 0
	}
	maxVal := slice[0]
	maxIndex := 0
	for i := 1; i < len(slice); i++ {
		//从第二个 元素开始循环比较，如果发现有更大的，则交换
		if maxVal < slice[i] {
			maxVal = slice[i]
			maxIndex = i
		}
	}
	return slice[maxIndex]
}

func Array_diff_float64(slice, slice2 []float64) []float64 {
	n64 := []float64{}
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

func Array_diff_int64(slice, slice2 []int64) []int64 {
	n64 := []int64{}
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

func Array_diff_string(slice, slice2 []string) []string {
	n64 := []string{}
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
