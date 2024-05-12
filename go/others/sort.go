package sort

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

	while i<len(left_nums) && j <len(right_nums){
		
	}
}
