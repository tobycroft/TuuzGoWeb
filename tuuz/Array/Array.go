package Array

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
