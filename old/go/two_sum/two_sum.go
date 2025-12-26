package twosum

// 時間計算量O(n^2), 空間計算量O(1)
func twoSum(nums []int, target int) []int {
	for i := 0; i < len(nums); i++ {
		for j := i + 1; j < len(nums); j++ {
			if (nums[i] + nums[j]) == target {
				return []int{i, j}
			}
		}
	}
	return []int{}
}

// 時間計算量O(n), 空間計算量O(n)
func twoSum2(nums []int, target int) []int {
	m := make(map[int]int)
	for i := 0; i < len(nums); i++ {
		diff := target - nums[i]
		if value, exists := m[diff]; exists {
			return []int{value, i}
		} else {
			m[nums[i]] = i
		}
	}
	return nil
}
