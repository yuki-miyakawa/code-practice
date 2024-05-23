from typing import List


def test(nums: List[int]):
    num_set = set(nums)
    print(type(num_set))


test([1, 2, 3])
