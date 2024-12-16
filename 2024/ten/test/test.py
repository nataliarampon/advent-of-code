import unittest

from ten.dayTen import getReacheableNodes, getPaths, getDirectedGraphAndZeros, partOneDay10, partTwoDay10

class TestDayTen(unittest.TestCase):
    def test_getDirectedGraphAndZeros(self):
        graph, zeros = getDirectedGraphAndZeros(['101240'])
        self.assertListEqual(zeros, [(0,1), (0,5)])
        self.assertDictEqual(dict(graph), {(0,1): [(0,0), (0,2)], (0,2): [(0,3)]})

    def test_getReacheableNodes(self):
        self.assertEqual(getReacheableNodes({(0,0): [(0,1)], (0,1): [(0,2)], (0,2): [(0,3)], (0,3): [(0,4)], \
            (0,4): [(0,5)], (0,5): [(0,6)], (0,6): [(0,7)], (0,7): [(0,8)], (0,8): [(0,9)]}, (0,0), [], 0), 1)
        self.assertEqual(getReacheableNodes({(0,0): [(0,1), (1,0)], (1,0): [(0,2)], (0,1): [(0,2)], (0,2): [(0,3)], (0,3): [(0,4)], \
            (0,4): [(0,5)], (0,5): [(0,6)], (0,6): [(0,7)], (0,7): [(0,8)], (0,8): [(0,9)]}, (0,0), [], 0), 1)

    def test_partOne(self):
        self.assertEqual(partOneDay10('ten/test/test.txt'), 36)
    
    def test_getPaths(self):
        self.assertEqual(getPaths({(0,0): [(0,1)], (0,1): [(0,2)], (0,2): [(0,3)], (0,3): [(0,4)], \
            (0,4): [(0,5)], (0,5): [(0,6)], (0,6): [(0,7)], (0,7): [(0,8)], (0,8): [(0,9)]}, (0,0), 0), 1)
        self.assertEqual(getPaths({(0,0): [(0,1), (1,0)], (1,0): [(0,2)], (0,1): [(0,2)], (0,2): [(0,3)], (0,3): [(0,4)], \
            (0,4): [(0,5)], (0,5): [(0,6)], (0,6): [(0,7)], (0,7): [(0,8)], (0,8): [(0,9)]}, (0,0), 0), 2)
    
    def test_partTwo(self):
        self.assertEqual(partTwoDay10('ten/test/test.txt'), 81)

if __name__ == "__main__":
    unittest.main()