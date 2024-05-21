import unittest
from maximum_frequency import MaximumFrequency


class TestMaximumFrequency(unittest.TestCase):
    def test1(self):
        mf = MaximumFrequency()
        self.assertEqual(mf.maximum_frequency([1, 2, 3, 2, 1]), 4)

    def test2(self):
        mf = MaximumFrequency()
        self.assertEqual(mf.maximum_frequency([1, 5, 2, 3, 6]), 2)


if __name__ == "__main__":
    unittest.main()
