package containtsduplicate

// time complexity O(n^2), Spacial complexity O(1)
func containsDuplicate(nums []int) bool {
	for i := 0; i < len(nums); i++ {
		for j := i + 1; j < len(nums); j++ {
			if nums[i] == nums[j] {
				return true
			}
		}
	}
	return false
}

// time complexity O(n), Spacial complexity O(n)
func containsDuplicate2(nums []int) bool {
	m := make(map[int]bool, len(nums))
	for _, num := range nums {
		if _, exists := m[num]; exists {
			return true
		}
		m[num] = true
	}
	return false
}
