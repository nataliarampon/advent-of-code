import unittest
from one.dayOne import extractColumns, partOneDay01, partTwoDay01

class TestDayOne(unittest.TestCase):
    def test_extractColumns(self):
        locations1, locations2 = extractColumns('one/test/test.txt')
        self.assertEqual(locations1, [5, 3])
        self.assertEqual(locations2, [1, 2])
    
    def test_partOne(self):
        self.assertEqual(partOneDay01('one/test/test.txt'), 5)
    
    def test_partTwo(self):
        self.assertEqual(partTwoDay01('one/test/test2.txt'), 31)

if __name__ == "__main__":
    unittest.main()