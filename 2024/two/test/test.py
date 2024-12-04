import unittest

from two.dayTwo import isReportSafe, isLevelSafe, isReportSafeWithDampening, partOneDay02, partTwoDay02

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
    
    def test_isReportSafeWithDampening(self):
        self.assertTrue(isReportSafeWithDampening([7, 6, 4, 2, 1]))
        self.assertFalse(isReportSafeWithDampening([1, 2, 7, 8, 9]))
        self.assertFalse(isReportSafeWithDampening([9, 7, 6, 2, 1]))
        self.assertTrue(isReportSafeWithDampening([1, 3, 2, 4, 5]))
        self.assertTrue(isReportSafeWithDampening([8, 6, 4, 4, 1]))
        self.assertFalse(isReportSafeWithDampening([0, 6, 4, 4, 1]))
        self.assertFalse(isReportSafeWithDampening([1, 6, 4, 4, 4]))     

    def test_partOne(self):
        self.assertEqual(partOneDay02('two/test/test.txt'), 2)
    
    def test_partTwo(self):
        self.assertEqual(partTwoDay02('two/test/test.txt'), 4)

if __name__ == "__main__":
    unittest.main()