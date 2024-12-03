import unittest
from one.dayOne import extractColumns, partOne

class TestDaytOne(unittest.TestCase):
    def test_extractColumns(self):
        locations1, locations2 = extractColumns('one/test/test.txt')
        self.assertEqual(locations1, [5, 3])
        self.assertEqual(locations2, [1, 2])
    
    def test_partOne(self):
        self.assertEqual(partOne('one/test/test.txt'), 5)

if __name__ == "__main__":
    unittest.main()