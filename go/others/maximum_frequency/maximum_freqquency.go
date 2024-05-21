package maximumfrequency

// Input: nums = [1, 2, 3, 2, 1]
// Output: 4

func maximum_frequency(nums []int) int {
	m := make(map[int]int)
	for i := 0; i < len(nums)+1; i++ {
		if _, ok := m[abs(nums[i], nums[i+1])]; ok {
			m[abs(nums[i], nums[i+1])]++
		}
		m[abs(nums[i], nums[i+1])] = 0
	}
	var ans int
	for _, n := range m {
		if ans < n {
			ans = n
		}
	}
	return ans
}

func abs(a, b int) int {
	if a > b {
		return a - b
	} else {
		return a - b
	}

}
