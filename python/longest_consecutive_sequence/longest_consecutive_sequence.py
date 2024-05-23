from typing import List


class LongestConsecutive:
    def longest_consecutive(self, nums: List[int]) -> int:
        ans, nums_set = 0, set(nums)

        for num in nums:
            if num not in nums_set:
                continue
            nums_set.remove(num)
            current_ans = 1
            i = 1

            while num + i in nums_set:
                nums_set.remove(num + i)
                current_ans += 1
                i += 1

            while num - i in nums_set:
                nums_set.remove(num - i)
                current_ans += 1
                i += 1
            ans = max(ans, current_ans)
        return ans
