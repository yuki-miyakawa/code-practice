import unittest
from longest_consecutive_sequence import LongestConsecutive


class TestLongestConsecutive(unittest.TestCase):
    def test1(self):
        lc = LongestConsecutive()
        self.assertEqual(lc.longest_consecutive([100, 4, 200, 1, 3, 2]), 4)


if __name__ == "__main__":
    unittest.main()
