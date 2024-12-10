from common.util import readFile

def isCalculationPossible(res, runningSum, op, operands):
    if operands == []:
        return runningSum == res
    else:
        if op == '*':
            sum = runningSum * int(operands[0])
        else:
            sum = runningSum + int(operands[0])
        return isCalculationPossible(res, sum, '+', operands[1:]) or isCalculationPossible(res, sum, '*', operands[1:])

def isCalculationPossibleThreeOps(res, runningSum, op, operands):
    if operands == []:
        return runningSum == res
    else:
        if op == '*':
            sum = runningSum * int(operands[0])
        elif op == '+':
            sum = runningSum + int(operands[0])
        else:
            sum = int(str(runningSum) + operands[0])
        return isCalculationPossibleThreeOps(res, sum, '+', operands[1:]) or \
            isCalculationPossibleThreeOps(res, sum, '*', operands[1:]) or \
            isCalculationPossibleThreeOps(res, sum, '||', operands[1:])


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
    calibrationResult = 0

    for line in lines:
        result, operands = line.split(':')
        operands = operands[1:].split(' ')
        if isCalculationPossibleThreeOps(int(result), int(operands[0]), '*', operands[1:]) or \
            isCalculationPossibleThreeOps(int(result), int(operands[0]), '+', operands[1:]) or \
            isCalculationPossibleThreeOps(int(result), int(operands[0]), '||', operands[1:]):
            calibrationResult += int(result)
    return calibrationResult