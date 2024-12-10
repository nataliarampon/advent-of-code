import unittest

from nine.dayNine import defragmentMemory, expandMemory, partOneDay09, partTwoDay09

class TestDayNine(unittest.TestCase):
    def test_expandMemory(self):
        self.assertEqual(expandMemory('12345'), [0,-1,-1,1,1,1,-1,-1,-1,-1,2,2,2,2,2])
    
    def test_defragmentMemory(self):
        self.assertEqual(defragmentMemory([0,-1,-1,1,1,1,-1,-1,-1,-1,2,2,2,2,2]), [0, 2, 2, 1, 1, 1, 2, 2, 2])

    def test_partOne(self):
        self.assertEqual(partOneDay09('nine/test/test.txt'), 1928)
    
    def test_partTwo(self):
        self.assertEqual(partTwoDay09('nine/test/test.txt'), 0)

if __name__ == "__main__":
    unittest.main()