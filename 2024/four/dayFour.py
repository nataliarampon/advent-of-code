from common.util import readFile

WORD = 'XMAS'
X_MAS_WORD = 'MAS'

def rotateMatrixClockwise(matrix):
    return [''.join(row) for row in zip(*matrix[::-1])]

def getDiagonal(matrix, i, j, word):
    return ''.join([matrix[i+d][j+d] for d in range(len(word))])

def getInverseDiagonal(matrix, i, j, word):
    return ''.join([matrix[i+d][j-d] for d in range(len(word))])

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
                    if getDiagonal(matrix, i, j, WORD) == 'XMAS':
                        matches += 1
        matrix = rotateMatrixClockwise(matrix)
    return matches

def partTwoDay04(filePath):
    lines = readFile(filePath)
    matches = 0
    matrix = lines

    for i in range(1, len(matrix) - 1):
        for j in range(1, len(matrix[0]) - 1):
            if matrix[i][j] == 'A':
                diag = getDiagonal(matrix, i-1, j-1, X_MAS_WORD)
                inverseDiag = getInverseDiagonal(matrix, i-1, j+1, X_MAS_WORD)
                if (diag == 'SAM' or diag == 'MAS') and (inverseDiag == 'MAS' or inverseDiag == 'SAM'):
                    matches += 1

    return matches
