from common.util import readFile
from collections import Counter

def extractColumns(filePath):
    column1, column2 = [], []
    for line in readFile(filePath):
        value1, value2 = map(int, line.split())
        column1.append(value1)
        column2.append(value2)
    return column1, column2

def partOneDay01(filePath):
    locations1, locations2 = map(sorted, extractColumns(filePath))
    difference = sum(abs(x - y) for x, y in zip(locations1, locations2))
    return difference

def partTwoDay01(filePath):
    locations1, locations2 = extractColumns(filePath)
    counter = Counter(locations2)
    similarity_score = sum(x * counter[x] for x in locations1)
    return similarity_score