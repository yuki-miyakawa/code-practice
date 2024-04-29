package goodpair

func numIdenticalPairs(nums []int) int {
	var goodPairs int
	for i := 0; i < len(nums); i++ {
		for j := i + 1; j < len(nums); j++ {
			if nums[i] == nums[j] {
				goodPairs++
			}
		}
	}
	return goodPairs
}
