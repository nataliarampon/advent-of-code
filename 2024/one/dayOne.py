from common.util import readFile
from functools import reduce

def extractColumns(filePath):
    column1, column2 = [], []
    for line in readFile(filePath):
        value1, value2 = map(int, line.split())
        column1.append(value1)
        column2.append(value2)
    return column1, column2

def partOne(filePath):
    locations1, locations2 = map(sorted, extractColumns(filePath))
    difference = sum(abs(x - y) for x, y in zip(locations1, locations2))
    return difference 