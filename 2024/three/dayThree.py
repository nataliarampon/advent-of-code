from common.util import readFile
import re

MUL_OP_REGEX = 'mul\((\d+),(\d+)\)'
MUL_COND_OP_REGEX = "do\(\)|don't\(\)|mul\(\d+,\d+\)"

def getMultiplyOps(memory):
    return sum(int(mul[0]) * int(mul[1]) for mul in re.findall(MUL_OP_REGEX, memory))

def partOneDay03(filePath):
    lines = readFile(filePath)
    return sum([getMultiplyOps(line) for line in lines])

def getMultiplyOpsWithConditionals(memory):
    enabled = True
    result = 0
    for match in re.findall(MUL_COND_OP_REGEX, memory):
        if 'mul' in match and enabled:
            mul_match = re.search(MUL_OP_REGEX, match)
            result += int(mul_match.group(1)) * int(mul_match.group(2))
        if "don't()" == match:
            enabled = False
        if 'do()' == match:
            enabled = True
    return result

def partTwoDay03(filePath):
    lines = readFile(filePath)
    return getMultiplyOpsWithConditionals(''.join(lines))
