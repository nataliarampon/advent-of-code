import unittest

from eleven.dayEleven import blink, partOneDay11

class TestDayEleven(unittest.TestCase):
    def test_blink(self):
        self.assertEqual(blink(''), '')
        self.assertEqual(blink(['0','1','10','99','999']), ['1','2024','1','0','9','9','2021976'])

    def test_partOne(self):
        self.assertEqual(partOneDay11('eleven/test/test.txt', 25), 55312)

if __name__ == "__main__":
    unittest.main()