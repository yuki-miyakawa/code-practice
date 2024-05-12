import unittest
from sort import BubbleSort


class TestBubbleSort(unittest.TestCase):
    def test1(self):
        bubble_sort = BubbleSort()
        self.assertEqual(bubble_sort.bubble_sort([3, 2, 1]), [1, 2, 3])

    def test2(self):
        bubble_sort = BubbleSort()
        self.assertEqual(
            bubble_sort.bubble_sort([3, 4, 5, 8, 8, 1, 9]), [1, 3, 4, 5, 8, 8, 9]
        )


if __name__ == "__main__":
    unittest.main()
