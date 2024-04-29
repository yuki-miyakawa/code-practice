import unittest
from good_pair import goodPair


class TestGoodPair(unittest.TestCase):
    def test_example(self):
        self.assertEqual(goodPair([1, 2, 3, 1, 1, 3]), 4)

    def test_all_same(self):
        self.assertEqual(goodPair([1, 1, 1, 1]), 6)

    def test_no_duplicates(self):
        self.assertEqual(goodPair([1, 2, 3]), 0)

    def test_empty_list(self):
        self.assertEqual(goodPair([]), 0)


if __name__ == "__main__":
    unittest.main()
