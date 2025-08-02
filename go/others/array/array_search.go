package array

func arraySearch(array []int, target int) int {
	for k, v := range array {
		if v == target {
			return k
		}
	}
	return -1
}
