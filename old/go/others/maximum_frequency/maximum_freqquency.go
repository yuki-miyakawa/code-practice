package maximumfrequency

import (
	"fmt"
)

// Input: nums = [1, 2, 3, 2, 1]
// Output: 4

func maximum_frequency(nums []int) int {
	m := make(map[int]int)
	for i := 0; i < len(nums)-1; i++ {
		diff := abs(nums[i], nums[i+1])
		if val, ok := m[diff]; ok {
			m[diff] = val + 1
		} else {
			m[diff] = 1
		}

	}
	fmt.Println(m)
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
		return b - a
	}

}
