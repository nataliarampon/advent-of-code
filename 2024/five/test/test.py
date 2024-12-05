import unittest

from common.util import readFile
from five.dayFive import isUpdateValid, partOneDay05, partTwoDay05, processRules

class TestDayFive(unittest.TestCase):
    def test_processRules(self):
        rulesBefore, rulesAfter, index = processRules(['1|2', '3|4', '1|56', '', '1|34'])
        self.assertDictEqual(rulesBefore, {'1': ['2', '56'], '3': ['4']})
        self.assertDictEqual(rulesAfter, {'56': ['1'], '2': ['1'], '4': ['3']})
        self.assertEqual(index, 4)
    
    def test_isUpdateValid(self):
        lines = readFile('five/test/test.txt')
        rulesBefore, rulesAfter, _ = processRules(lines)

        self.assertTrue(isUpdateValid(rulesBefore, rulesAfter, ['75','47','61','53','29']))
        self.assertTrue(isUpdateValid(rulesBefore, rulesAfter, ['97','61','53','29','13']))
        self.assertTrue(isUpdateValid(rulesBefore, rulesAfter, ['75','29','13']))
        self.assertFalse(isUpdateValid(rulesBefore, rulesAfter, ['75','97','47','61','53']))
        self.assertFalse(isUpdateValid(rulesBefore, rulesAfter, ['61','13','29']))
        self.assertFalse(isUpdateValid(rulesBefore, rulesAfter, ['97','13','75','29','47']))

    def test_partOne(self):
        self.assertEqual(partOneDay05('five/test/test.txt'), 143)
    
    def test_partTwo(self):
        self.assertEqual(partTwoDay05('five/test/test.txt'), 0)

if __name__ == "__main__":
    unittest.main()