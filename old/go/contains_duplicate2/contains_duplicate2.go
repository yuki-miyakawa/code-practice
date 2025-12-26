package containsduplicate2

// create my own
func containsNearbyDuplicate(nums []int, k int) bool {
	m := make(map[int]int)
	var res bool
	for i := 0; i < len(nums); i++ {
		if _, ok := m[nums[i]]; ok {
			if abs(m[nums[i]]-i) <= k {
				res = true
			}
		}
		m[nums[i]] = i
	}
	return res
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

// better solution
func containsNearbyDuplicate2(nums []int, k int) bool {
	if k == 0 {
		return false
	}
	m := make(map[int]int, len(nums))
	for i := 0; i < len(nums); i++ {
		_, ok := m[nums[i]]
		if ok && i-m[nums[i]] <= k {
			return true
		}
		m[nums[i]] = i
	}
	return false
}
