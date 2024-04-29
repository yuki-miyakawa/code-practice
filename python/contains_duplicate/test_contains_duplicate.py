import unittest
from contains_duplicate import Solution


class TestContainsDuplicate(unittest.TestCase):
    def test_example(self):
        solution = Solution()
        self.assertEqual(solution.containsDuplicate([1, 2, 3, 1]), True)

    def test_example2(self):
        solution = Solution()
        self.assertEqual(solution.containsDuplicate([1, 2, 3, 4]), False)

    def test_example3(self):
        solution = Solution()
        self.assertEqual(
            solution.containsDuplicate([1, 1, 1, 3, 3, 4, 3, 2, 4, 2]), True
        )


if __name__ == "__main__":
    unittest.main()
