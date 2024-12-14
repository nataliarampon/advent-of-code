import unittest

from nine.dayNine import File, defragmentMemory, defragmentMemoryForFile, expandMemory, expandMemoryForFile, getIndexByFileSize, partOneDay09, partTwoDay09

class TestDayNine(unittest.TestCase):
    def test_expandMemory(self):
        self.assertListEqual(expandMemory('12345'), [0,-1,-1,1,1,1,-1,-1,-1,-1,2,2,2,2,2])
    
    def test_expandMemoryForFile(self):
        self.assertListEqual(expandMemoryForFile('12345'), [File(1,0), File(2), File(3,1), File(4), File(5,2)])
    
    def test_defragmentMemory(self):
        self.assertListEqual(defragmentMemory([0,-1,-1,1,1,1,-1,-1,-1,-1,2,2,2,2,2]), [0, 2, 2, 1, 1, 1, 2, 2, 2])
    
    def test_defragmentMemoryForFile(self):
        self.assertListEqual(defragmentMemoryForFile([File(1,0), File(2), File(3,1), File(4), File(5,2)]), [File(1,0), File(2), File(3,1), File(4), File(5,2)])
        self.assertListEqual(defragmentMemoryForFile([File(1,0), File(4), File(3,1), File(4), File(5,2)]), [File(1,0), File(3,1), File(1), File(3), File(4), File(5,2)])

    def test_getIndexByFileSize(self):
        self.assertEqual(getIndexByFileSize([File(1,0), File(2), File(3,1), File(4), File(5,2)], 3, 2), -1)
        self.assertEqual(getIndexByFileSize([File(1,0), File(3), File(3,1), File(4), File(1,2)], 3, 5), 1)
        self.assertEqual(getIndexByFileSize([File(1,0), File(2), File(3,1), File(4), File(5,2)], 3, 4), 3)

    def test_partOne(self):
        self.assertEqual(partOneDay09('nine/test/test.txt'), 1928)
    
    def test_partTwo(self):
        self.assertEqual(partTwoDay09('nine/test/test.txt'), 2858)

if __name__ == "__main__":
    unittest.main()