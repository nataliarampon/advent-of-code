from common.util import readFile

def isLevelSafe(before, after, isAscending):
    step = abs(before - after)
    result = True
    if  step <= 3 and step >= 1:
        if isAscending == None:
            isAscending = True if before < after else False
        if (isAscending and before > after) or (not isAscending and before < after):
            result = False
    else:
        result = False
    return result, isAscending

def isReportSafe(report):
    isAscending = None
    for i in range(len(report) - 1):
        safety, isAscending = isLevelSafe(report[i], report[i+1], isAscending)
        if not safety:
            return safety
    return safety

def isReportSafeWithDampening(report):
    isAscending = None
    for i in range(len(report) - 1):
        step = abs(report[i] - report[i+1])
        if  step <= 3 and step >= 1:
            if isAscending == None:
                isAscending = True if report[i] < report[i+1] else False
            if (isAscending and report[i] > report[i+1]) or (not isAscending and report[i] < report[i+1]):
                return False
        else:
            return False
    return True

def partOneDay02(filePath):
    lines = readFile(filePath)
    numberSafeReports = sum([1 if isReportSafe(list(map(int, line.split()))) else 0 for line in lines])
    return numberSafeReports