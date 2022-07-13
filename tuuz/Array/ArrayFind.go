package Array

func ArrayFindMax[T int | int32 | int64 | float32 | float64](slice []T) T {
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

func ArrayFindMin[T int | int32 | int64 | float32 | float64](slice []T) T {
	if len(slice) < 1 {
		return 0
	}
	maxVal := slice[0]
	maxIndex := 0
	for i := 1; i < len(slice); i++ {
		//从第二个 元素开始循环比较，如果发现有更小的，则交换
		if maxVal > slice[i] {
			maxVal = slice[i]
			maxIndex = i
		}
	}
	return slice[maxIndex]
}
