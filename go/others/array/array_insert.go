package array

func insert(array []int, index, value int) []int {
	result := make([]int, len(array)+1)
	for k, v := range array {
		if k < index {
			result[k] = v
		}
		if k == index {
			result[k] = value
			result[k+1] = v
		}
		if k > index {
			result[k+1] = v
		}
	}
	return result
}

func insert2(array []int, index, value int) []int {
	result := make([]int, len(array)+1)
	copy(result[:index], array[:index])
	result[index] = value
	copy(result[index+1:], array[index:])
	return result
}
