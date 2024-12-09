import unittest

from six.daySix import isLoop, partOneDay06, partTwoDay06, getGuardCoordinates

class TestDaySix(unittest.TestCase):
    def test_getGuardCoordinates(self):
        self.assertEqual(getGuardCoordinates(['.......','...^...']), (1,3))
        self.assertEqual(getGuardCoordinates(['.......','.......']), (-1,-1))

    def test_partOne(self):
        self.assertEqual(len(partOneDay06('six/test/test.txt')), 41)
    
    def test_partTwo(self):
        self.assertEqual(partTwoDay06('six/test/test.txt'), 6)

    def test_isLoop(self):
        self.assertTrue(isLoop(['..#.', '.#^#.', '..#.'], 1, 2))
        self.assertFalse(isLoop(['..#.', '.#^..', '..#.'], 1, 2))

if __name__ == "__main__":
    unittest.main()