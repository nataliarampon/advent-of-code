import unittest
from two.dayTwo import isReportSafe, isLevelSafe, partOneDay02

class TestDayTwo(unittest.TestCase):
    def test_isLevelSafe(self):
        self.assertEqual(isLevelSafe(7,6, None), (True, False))
        self.assertEqual(isLevelSafe(6,8, None), (True, True))
        self.assertEqual(isLevelSafe(6,10, True), (False, True))
        self.assertEqual(isLevelSafe(6,6, True), (False, True))
        self.assertEqual(isLevelSafe(6,8, False), (False, False))

    def test_isReportSafe(self):
        self.assertTrue(isReportSafe([7, 6, 4, 2, 1]))
        self.assertFalse(isReportSafe([1, 2, 7, 8, 9]))
        self.assertFalse(isReportSafe([1, 3, 2, 4, 5]))
        self.assertFalse(isReportSafe([8, 6, 4, 4, 1]))

    def test_partOne(self):
        self.assertEqual(partOneDay02('two/test/test.txt'), 2)
    
    # def test_partTwo(self):
    #     self.assertEqual(partTwo('one/test/test2.txt'), 31)

if __name__ == "__main__":
    unittest.main()