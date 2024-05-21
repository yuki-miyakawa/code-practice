from typing import List
from collections import defaultdict


class MaximumFrequency:
    def maximum_frequency(self, nums: List[int]) -> int:
        freq_map = defaultdict(int)
        ans = 0
        for i in range(len(nums) - 1):
            diff = abs(nums[i] - nums[i + 1])
            freq_map[diff] += 1
            ans = max(ans, freq_map[diff])
        return ans
