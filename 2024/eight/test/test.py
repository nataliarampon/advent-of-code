import unittest

from eight.dayEight import getAntennas, getAntinodes, partOneDay08, partTwoDay08

class TestDayEight(unittest.TestCase):
    def test_getAntennas(self):
        self.assertDictEqual(getAntennas(['.a..4..A', '.a......']), {'a': [(0,1), (1,1)], '4': [(0,4)], 'A': [(0,7)]})

    def test_getAntinodes(self):
        self.assertEqual(getAntinodes((1,8), (2,5)), ((0,11), (3,2)))
        self.assertEqual(getAntinodes((2,5), (3,7)), ((1,3), (4,9)))

    def test_partOne(self):
        self.assertEqual(partOneDay08('eight/test/test.txt'), 14)
    
    def test_partTwo(self):
        self.assertEqual(partTwoDay08('eight/test/test.txt'), 0)

if __name__ == "__main__":
    unittest.main()