package longestconsecutivesequence

import "fmt"

// Given an unsorted array of integers nums, return the length of the longest consecutive elements sequence.

// You must write an algorithm that runs in O(n) time.

func longestConsecutive(nums []int) int {
	var m = make(map[int]int)
	var max int
	for _, num := range nums {
		if _, exists := m[num-1]; exists {
			m[num] = m[num-1] + 1
		} else {
			m[num] = 1
		}
		fmt.Println(m)
	}
	for _, v := range m {
		if v > max {
			max = v
		}
	}
	return max
}
