from typing import List


class BubbleSort:
    def bubble_sort(self, nums: List[int]) -> List[int]:
        n = len(nums)
        for i in range(n):
            for j in range(n - i - 1):
                if nums[j] > nums[j + 1]:
                    nums[j], nums[j + 1] = nums[j + 1], nums[j]
        return nums


class MergeSort:
    def merge_sort(self, nums: List[int], left: int, right: int) -> List[int]:
        if left >= right:
            return nums
        mid = (left + right) // 2  # // is floor division(整数除算)
        print(f"nums: {nums}, left: {left}, mid: {mid}, right: {right}")
        self.merge_sort(nums, left, mid)
        self.merge_sort(nums, mid + 1, right)

        self.merge(nums, left, mid, right)

        return nums

    def merge(self, nums: List[int], left: int, mid: int, right: int):
        left_nums = nums[left : mid + 1]
        right_nums = nums[mid + 1 : right + 1]

        i, j, k = 0, 0, left

        while i < len(left_nums) and j < len(right_nums):
            if left_nums[i] <= right_nums[j]:
                nums[k] = left_nums[i]
                i += 1
            else:
                nums[k] = right_nums[j]
                j += 1
            k += 1

        while i < len(left_nums):
            nums[k] = left_nums[i]
            i += 1
            k += 1

        while j < len(right_nums):
            nums[k] = right_nums[j]
            j += 1
            k += 1


if __name__ == "__main__":
    mergesort = MergeSort()
    print(mergesort.merge_sort([5, 2, 3, 2, 1], 0, 4))
