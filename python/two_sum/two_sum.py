def twoSum(nums, target):
    for i in range(len(nums)):
        for j in range(i + 1, len(nums)):
            if nums[i] + nums[j] == target:
                return [i, j]
    return None


def twoSum2(nums, target):
    for i in range(len(nums)):
        diff = target - nums[i]
        if diff in nums:
            return [i, nums.index(diff)]
    return None
