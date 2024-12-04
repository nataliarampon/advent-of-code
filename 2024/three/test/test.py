import unittest

from three.dayThree import getMultiplyOps, getMultiplyOpsWithConditionals, partOneDay03, partTwoDay03

class TestDayThree(unittest.TestCase):
    def test_getMultiplyOps(self):
        self.assertEqual(getMultiplyOps('xmul(2,4)%&mul[3,7]!@^do_not_mul(5,5)+mul(32,64]then(mul(11,8)mul(8,5))'), 161)
    
    def test_getMultiplyOpsWithConditionals(self):
        self.assertEqual(getMultiplyOpsWithConditionals("xmul(2,4)&mul[3,7]!^don't()_mul(5,5)+mul(32,64](mul(11,8)undo()?mul(8,5))"), 48)

    def test_partOne(self):
        self.assertEqual(partOneDay03('three/test/test.txt'), 163)
    
    def test_partTwo(self):
        self.assertEqual(partTwoDay03('three/test/test.txt'), 163)

if __name__ == "__main__":
    unittest.main()