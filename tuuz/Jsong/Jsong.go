package Jsong

import (
	"fmt"
	"github.com/bytedance/sonic"
	"strings"
)

func Encode(data interface{}) (string, error) {
	jb, err := sonic.MarshalString(data)
	if err != nil {
		fmt.Println("JENCODEEncode", err)
		return "", err
	}
	return string(jb), err
}

func JArrayObject[T string | int | int32 | int64 | float32 | float64, V string | int | int32 | int64 | float32 | float64 | any](data string) ([]map[T]V, error) {
	var arr []map[T]V
	err := sonic.UnmarshalString(data, &arr)
	if err != nil {
		return nil, err
	}
	return arr, err
}

func JArray[T string | int | int32 | int64 | float32 | float64 | any](data string) ([]T, error) {
	var arr []T
	err := sonic.UnmarshalString(data, &arr)
	if err != nil {
		return nil, err
	}
	return arr, err
}

func JObject[T string | int | int32 | int64 | float32 | float64, V string | int | int32 | int64 | float32 | float64 | any](data string) (map[T]V, error) {
	var arr map[T]V
	err := sonic.UnmarshalString(data, &arr)
	if err != nil {
		return nil, err
	}
	return arr, err
}

func TCPJObject(temp *string) ([]map[string]interface{}, error) {
	var arr []map[string]interface{}

	//var strs []string

	data := *temp
	if len(*temp) > 65535 {
		*temp = ""
		return nil, fmt.Errorf("%s", "too long")
	}
	strs := strings.Split(data, "}{")
	if len(strs) > 2 {
		unable := ""
		for i, v := range strs {
			arr2 := make(map[string]interface{})
			if i == 0 {
				err := sonic.Unmarshal([]byte(v+"}"), &arr2)
				if err != nil {
					//unable += v + "}"
					fmt.Println(1, i, i+1, v+"}")
				} else {
					arr = append(arr, arr2)
				}
			} else if len(strs) == int(i+1) {
				err := sonic.Unmarshal([]byte("{"+v), &arr2)
				if err != nil {
					unable += "{" + v
					//fmt.Println(2, i, i+1, "{"+v)
				} else {
					arr = append(arr, arr2)
				}
				//fmt.Println(2, "len", len(strs), i+1, "{"+v)
			} else {
				err := sonic.Unmarshal([]byte("{"+v+"}"), &arr2)
				if err != nil {
					unable += "{" + v + "}"
					fmt.Println(3, i, "{"+v+"}")
				} else {
					arr = append(arr, arr2)
				}
			}
		}
		*temp = unable
		return arr, nil
	} else if len(strs) > 1 {
		arr2 := make(map[string]interface{})
		err := sonic.Unmarshal([]byte(strs[0]+"}"), &arr2)
		if err != nil {
			//fmt.Println("2",data)
			//fmt.Println(err)
			return nil, err
		} else {
			*temp = "{" + strs[1]
			arr = append(arr, arr2)
			return arr, err
		}
	} else {
		arr2 := make(map[string]interface{})
		err := sonic.Unmarshal([]byte(data), &arr2)
		if err != nil {
			//fmt.Println("2",data)
			//fmt.Println(err)
			return nil, err
		} else {
			*temp = ""
			arr = append(arr, arr2)
			return arr, err
		}
	}

}

func TCPJArray(temp *string) ([]interface{}, error) {

	var arr []interface{}

	//var strs []string

	data := *temp
	strs := strings.Split(data, "][")
	if len(strs) > 2 {
		unable := ""
		for i, v := range strs {
			var arr2 interface{}
			if i == 0 {
				err := sonic.Unmarshal([]byte(v+"]"), &arr2)
				if err != nil {
					unable += v + "}"
					fmt.Println(1, i, i+1, v+"]")
				} else {
					arr = append(arr, arr2)
				}
			} else if len(strs) == int(i+1) {
				err := sonic.Unmarshal([]byte("["+v), &arr2)
				if err != nil {
					unable += "{" + v
					fmt.Println(2, i, i+1, "["+v)
				} else {
					arr = append(arr, arr2)
				}
				//fmt.Println(2, "len", len(strs), i+1, "{"+v)
			} else {
				err := sonic.Unmarshal([]byte("["+v+"]"), &arr2)
				if err != nil {
					unable += "{" + v + "}"
					fmt.Println(3, i, "["+v+"]")
				} else {
					arr = append(arr, arr2)
				}
			}
		}

		*temp = unable
		return arr, nil
	} else if len(strs) > 1 {
		var arr2 interface{}
		err := sonic.Unmarshal([]byte(strs[0]+"]"), &arr2)
		if err != nil {
			//fmt.Println("2",data)
			//fmt.Println(err)
			return nil, err
		} else {
			*temp = "[" + strs[1]
			arr = append(arr, arr2)
			return arr, err
		}
	} else {
		var arr2 interface{}
		err := sonic.Unmarshal([]byte(data), &arr2)
		if err != nil {
			//fmt.Println("2",data)
			//fmt.Println(err)
			return nil, err
		} else {
			*temp = ""
			arr = append(arr, arr2)
			return arr, err
		}
	}

}

func TCP_JSON_CUT(temp *string) (string, bool) {
	var arr []map[string]interface{}
	var arr2 map[string]interface{}
	//var strs []string

	data := *temp
	strs := strings.Split(data, "}{")
	if len(strs) > 1 {
		for i, v := range strs {
			if i == 0 {
				strs[i] = v + "}"
			} else if len(strs) == i+1 {

				strs[i] = "{" + v
			} else {
				strs[i] = "{" + v + "}"
			}
			//fmt.Println(strs[i])
		}
		data = "[" + strings.Join(strs, ",") + "]"
		err := sonic.Unmarshal([]byte(data), &arr)
		if err != nil {
			fmt.Println(err)
		} else {
			//fmt.Println("ss:", arr)
		}
		*temp = ""
		return data, true
	} else {
		err := sonic.Unmarshal([]byte(data), &arr2)
		if err != nil {
			//fmt.Println("2",data)
			//fmt.Println(err)
			return "", false
		} else {
			*temp = ""
			return data, true
		}
	}
}
