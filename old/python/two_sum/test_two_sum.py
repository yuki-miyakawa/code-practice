import unittest
from two_sum import twoSum, twoSum2


class TestTowSum(unittest.TestCase):
    def test_example(self):
        self.assertEqual(twoSum([2, 7, 11, 15], 9), [0, 1])

    def test_example2(self):
        self.assertEqual(twoSum([3, 2, 4], 6), [1, 2])

    def test_empty_list(self):
        self.assertEqual(twoSum([], 10), None)

    def test_one_element(self):
        self.assertEqual(twoSum([1], 10), None)

    def test_two_elements(self):
        self.assertEqual(twoSum([1, 2], 3), [0, 1])


class TestTwoSum2(unittest.TestCase):
    def test_example(self):
        self.assertEqual(twoSum2([2, 7, 11, 15], 9), [0, 1])

    def test_example2(self):
        self.assertEqual(twoSum([3, 2, 4], 6), [1, 2])

    def test_empty_list(self):
        self.assertEqual(twoSum2([], 10), None)

    def test_one_element(self):
        self.assertEqual(twoSum2([1], 10), None)

    def test_two_elements(self):
        self.assertEqual(twoSum2([1, 2], 3), [0, 1])


if __name__ == "__main__":
    unittest.main()
