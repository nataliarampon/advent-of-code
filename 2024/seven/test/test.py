import unittest

from seven.daySeven import isCalculationPossible, isCalculationPossibleThreeOps, partOneDay07, partTwoDay07

class TestDaySeven(unittest.TestCase):
    def test_isCalculationPossible(self):
        self.assertTrue(isCalculationPossible(190, 19, '*', ['10']))
        self.assertFalse(isCalculationPossible(190, 19, '+', ['10']))

        self.assertFalse(isCalculationPossible(83, 17, '+', ['5']))
        self.assertFalse(isCalculationPossible(83, 17, '*', ['5']))

        self.assertTrue(isCalculationPossible(3267, 81, '*', ['40', '27']))
        self.assertTrue(isCalculationPossible(3267, 81, '+', ['40', '27']))
    
    def test_isCalculationPossibleThreeOps(self):
        self.assertTrue(isCalculationPossibleThreeOps(190, 19, '*', ['10']))
        self.assertFalse(isCalculationPossibleThreeOps(190, 19, '+', ['10']))

        self.assertFalse(isCalculationPossibleThreeOps(83, 17, '+', ['5']))
        self.assertFalse(isCalculationPossibleThreeOps(83, 17, '*', ['5']))

        self.assertTrue(isCalculationPossibleThreeOps(156, 15, '||', ['6']))
        self.assertTrue(isCalculationPossibleThreeOps(7290, 6, '*', ['8', '6', '15']))

    def test_partOne(self):
        self.assertEqual(partOneDay07('seven/test/test.txt'), 3749)
    
    def test_partTwo(self):
        self.assertEqual(partTwoDay07('seven/test/test.txt'), 11387)

if __name__ == "__main__":
    unittest.main()