package sort

func bubbleSort(nums []int) []int {
	for i := 0; i < len(nums); i++ {
		for j := 0; j < len(nums)-i-1; j++ {
			if nums[j] > nums[j+1] {
				nums[j], nums[j+1] = nums[j+1], nums[j]
			}
		}
	}
	return nums
}

func mergeSort(nums []int, left, right int) []int {
	if left >= right {
		return nums
	}
	mid := (left + right) / 2
	mergeSort(nums, left, right)
	mergeSort(nums, mid+1, right)

	merge(nums, left, mid, right)

	return nums
}

func merge(nums []int, left, mid, right int) {
	left_nums := nums[left : mid+1]
	right_nums := nums[mid+1 : right+1]

	i, j, k := 0, 0, left

	for i < len(left_nums) && j < len(right_nums) {
		if left_nums[i] <= right_nums[j] {
			nums[k] = left_nums[i]
			i++
		} else {
			nums[k] = right_nums[j]
			j++
		}
		k++
	}

	for i < len(left_nums) {
		nums[k] = left_nums[i]
		i++
		k++
	}

	for j < len(right_nums) {
		nums[k] = right_nums[j]
		j++
		k++
	}
}
