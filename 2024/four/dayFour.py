from common.util import readFile

WORD = 'XMAS'

def rotateMatrixClockwise(matrix):
    return [''.join(row) for row in zip(*matrix[::-1])]

def getDiagonal(matrix, i, j):
    return ''.join([matrix[i+d][j+d] for d in range(len(WORD))])

def partOneDay04(filePath):
    NB_ROTATIONS = 4

    lines = readFile(filePath)
    matches = 0
    matrix = lines

    for _ in range(NB_ROTATIONS):
        for i in range(len(matrix)):
            for j in range(len(matrix[0]) - len(WORD) + 1):
                if matrix[i][j:j+len(WORD)] == 'XMAS':
                    matches += 1
                if i <= (len(matrix) - len(WORD)):
                    if getDiagonal(matrix, i, j) == 'XMAS':
                        matches += 1
        matrix = rotateMatrixClockwise(matrix)
    return matches

def partTwoDay04(filePath):
    lines = readFile(filePath)
    return 0
