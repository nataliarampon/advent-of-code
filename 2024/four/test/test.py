import unittest

from four.dayFour import getDiagonal, partOneDay04, partTwoDay04, rotateMatrixClockwise

class TestDayFour(unittest.TestCase):
    def test_rotateMatrixClockwise(self):
        self.assertEqual(rotateMatrixClockwise(['XM','BI']), ['BX','IM'])
    
    def test_getDiagonal(self):
        self.assertEqual(getDiagonal(['X...','.M..', '..A.', '...S'], 0, 0), 'XMAS')

    def test_partOne(self):
        self.assertEqual(partOneDay04('four/test/test.txt'), 18)
    
    # def test_partTwo(self):
    #     self.assertEqual(partTwoDay04('four/test/test.txt'), 0)

if __name__ == "__main__":
    unittest.main()