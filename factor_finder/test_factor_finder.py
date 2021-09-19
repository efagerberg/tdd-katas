import unittest

from factor_finder import FactorFinder


class TestFactorFinder(unittest.TestCase):

    def setUp(self):
        self.finder = FactorFinder()

    def assert_gcf(self, expected, numbers):
        self.assertEquals(expected,
                          self.finder.greatest_common_factor(numbers))

    def assert_common_factors(self, expected, numbers):
        self.assertEquals(expected,
                          self.finder.get_common_factors(numbers))

    def assert_factors(self, expected, number):
        self.assertEquals(expected, self.finder.get_factors(number))

    def test_returns_correct_gcf(self):
        self.assert_gcf(1, [1])
        self.assert_gcf(5, [5])
        self.assert_gcf(1, [7, 5])
        self.assert_gcf(1, [7, 5, 1])
        self.assert_gcf(3, [6, 3])
        self.assert_gcf(1, [6, 3, 10])
        self.assert_gcf(10, [10, 100, 10])
        self.assert_gcf(1, [-1])
        self.assert_gcf(10, [-10, 100, -100])

    def test_returns_correct_factors(self):
        self.assert_factors(set([1]), 1)
        self.assert_factors(set([1, 3, 5, 15]), 15)
        self.assert_factors(set([-1, -2, -3, -6, 1, 2, 3, 6]), -6)

    def test_returns_correct_common_factors(self):
        self.assert_common_factors(set([1]), [1])
        self.assert_common_factors(set([1, 3, 5, 15]), [15])
        self.assert_common_factors(set([1]), [1, 2])
        self.assert_common_factors(set([1, 5]), [15, 5])


if __name__ == '__main__':
    unittest.main()
