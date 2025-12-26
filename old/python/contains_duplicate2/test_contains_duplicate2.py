import unittest
from contains_duplicate2 import Solution


class TestContainsDuplicate2(unittest.TestCase):
    def test1(self):
        solution = Solution()
        self.assertEqual(solution.containsDuplicate2([1, 2, 3, 1], 3), True)

    def test2(self):
        solution = Solution()
        self.assertEqual(solution.containsDuplicate2([1, 0, 1, 1], 1), True)

    def test3(self):
        solution = Solution()
        self.assertEqual(solution.containsDuplicate2([1, 2, 3, 1, 2, 3], 2), False)


if __name__ == "__main__":
    unittest.main()
