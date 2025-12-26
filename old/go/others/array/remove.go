package array

func remove(array []int, index int) []int {
	return append(array[:index], array[index+1:]...)
}
