from common.util import readFile

def isCalculationPossible(res, runningSum, op, operands):
    if operands == []:
        return runningSum == res
    else:
        if op == '*':
            return isCalculationPossible(res, runningSum * int(operands[0]), '*', operands[1:]) or isCalculationPossible(res, runningSum * int(operands[0]), '+', operands[1:])
        else:
            return isCalculationPossible(res, runningSum + int(operands[0]), '+', operands[1:]) or isCalculationPossible(res, runningSum + int(operands[0]), '*', operands[1:])

def partOneDay07(filePath):
    lines = readFile(filePath)
    calibrationResult = 0

    for line in lines:
        result, operands = line.split(':')
        operands = operands[1:].split(' ')
        if isCalculationPossible(int(result), int(operands[0]), '*', operands[1:]) or isCalculationPossible(int(result), int(operands[0]), '+', operands[1:]):
            calibrationResult += int(result)
    return calibrationResult

def partTwoDay07(filePath):
    lines = readFile(filePath)
    return 0