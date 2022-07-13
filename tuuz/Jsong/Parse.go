package Jsong

func ParseObject[T string | int | int32 | int64 | float32 | float64, V string | int | int32 | int64 | float32 | float64 | any](data interface{}) (map[T]V, error) {
	ret, err := Encode(data)
	if err != nil {
		return nil, err
	}
	return JObject[T, V](ret)
}

func ParseSlice[T string | int | int32 | int64 | float32 | float64 | any](data interface{}) ([]T, error) {
	ret, err := Encode(data)
	if err != nil {
		return nil, err
	}
	return JArray[T](ret)
}

func ParseArrayObject[T string | int | int32 | int64 | float32 | float64, V string | int | int32 | int64 | float32 | float64 | any](data interface{}) ([]map[T]V, error) {
	ret, err := Encode(data)
	if err != nil {
		return nil, err
	}
	return JArrayObject[T, V](ret)
}
