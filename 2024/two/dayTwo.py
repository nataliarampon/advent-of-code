from common.util import readFile

def isReportSafe(report):
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