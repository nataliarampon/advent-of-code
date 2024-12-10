from common.util import readFile, isOutOfBounds
from itertools import combinations

def getAntennas(matrix):
    antennas = {}
    for i in range(len(matrix)):
        for j in range(len(matrix[0])):
            if matrix[i][j] != '.':
                if matrix[i][j] in antennas:
                    antennas[matrix[i][j]].append((i,j))
                else:
                    antennas[matrix[i][j]] = [(i,j)]
    return antennas

def getAntinodes(antenna1, antenna2):
    diffCol = abs(antenna1[0] - antenna2[0])
    diffRow = abs(antenna1[1] - antenna2[1])

    if antenna1[0] < antenna2[0]:
        diffCol = -diffCol
    if antenna1[1] < antenna2[1]:
        diffRow = -diffRow

    return (antenna1[0] + diffCol, antenna1[1] + diffRow), (antenna2[0] - diffCol, antenna2[1] -diffRow)

def partOneDay08(filePath):
    lines = readFile(filePath)
    antennas = getAntennas(lines)
    antinodes = []
    for frequencies in antennas.values():
        if len(frequencies) > 1:
            for combination in combinations(frequencies, 2):
                antinodes += [*getAntinodes(combination[0], combination[1])]
    antinodes = set([a for a in antinodes if not isOutOfBounds(a[0], a[1], lines)])
    
    return len(antinodes)

def getHarmonicAntinodes(antenna1, antenna2, matrix):
    diffCol = abs(antenna1[0] - antenna2[0])
    diffRow = abs(antenna1[1] - antenna2[1])

    antinodes = []

    if antenna1[0] < antenna2[0]:
        diffCol = -diffCol
    if antenna1[1] < antenna2[1]:
        diffRow = -diffRow

    i = 0
    tempAntinode = (antenna1[0] + i*diffCol, antenna1[1] + i*diffRow)
    while not isOutOfBounds(tempAntinode[0], tempAntinode[1], matrix):
        antinodes.append(tempAntinode)
        i += 1
        tempAntinode = (antenna1[0] + i*diffCol, antenna1[1] + i*diffRow)
    
    i = 0
    tempAntinode = (antenna2[0] - i*diffCol, antenna2[1] - i*diffRow)
    while not isOutOfBounds(tempAntinode[0], tempAntinode[1], matrix):
        antinodes.append(tempAntinode)
        i += 1
        tempAntinode = (antenna2[0] - i*diffCol, antenna2[1] - i*diffRow)
    return antinodes

def partTwoDay08(filePath):
    lines = readFile(filePath)
    antennas = getAntennas(lines)
    antinodes = []
    for frequencies in antennas.values():
        if len(frequencies) > 1:
            for combination in combinations(frequencies, 2):
                antinodes += getHarmonicAntinodes(combination[0], combination[1], lines)
    antinodes = set([a for a in antinodes if not isOutOfBounds(a[0], a[1], lines)])
    
    return len(antinodes)