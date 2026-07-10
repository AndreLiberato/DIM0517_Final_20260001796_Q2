import unittest

from Calc import add, sub


class TestCalc(unittest.TestCase):

    def test_add(self):
        self.assertEqual(add(3, 4), 7)

    def test_sub_4_3(self):
        self.assertEqual(sub(4, 3), 1)

    def test_sub_3_4(self):
        self.assertEqual(sub(3, 4), -1)


if __name__ == "__main__":
    unittest.main()
