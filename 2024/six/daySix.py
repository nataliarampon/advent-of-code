from common.util import readFile, isOutOfBounds, move, Direction

NB_DIRECTIONS = 4

def getGuardCoordinates(matrix):
    GUARD_SYMBOL = '^'
    for i, row in enumerate(matrix):
        col = row.find(GUARD_SYMBOL)
        if col != -1:
            return i, col
    return -1, -1

def partOneDay06(filePath):
    lines = readFile(filePath)
    guardRow, guardCol = getGuardCoordinates(lines)
    lines[guardRow] = lines[guardRow][:guardCol] + 'X' + lines[guardRow][guardCol+1:]
    steps = [(guardRow, guardCol)]
    direction = Direction.UP.value
    while not isOutOfBounds(guardRow, guardCol, lines):
        tempRow, tempCol = move(guardRow, guardCol, direction)
        if isOutOfBounds(tempRow, tempCol, lines):
            break
        if lines[tempRow][tempCol] == '#':
            direction = (direction + 1) % NB_DIRECTIONS
        else:
            lines[guardRow] = lines[guardRow][:guardCol] + 'X' + lines[guardRow][guardCol+1:]
            guardRow, guardCol = tempRow, tempCol
        steps += [(guardRow, guardCol)] if lines[tempRow][tempCol] == '.' else []
    return steps

def isLoop(lines, guardRow, guardCol):
    visited = {}
    direction = Direction.UP.value

    while not isOutOfBounds(guardRow, guardCol, lines):
        tempRow, tempCol = move(guardRow, guardCol, direction)
        if isOutOfBounds(tempRow, tempCol, lines):
            return False
        if lines[tempRow][tempCol] == '#':
            if (tempRow, tempCol) not in visited:
                visited[(tempRow, tempCol)] = [False] * 4 
            elif visited[(tempRow, tempCol)][direction]:
                    return True
            else:
                visited[(tempRow, tempCol)][direction] = True
            direction = (direction + 1) % NB_DIRECTIONS
        else:
            lines[guardRow] = lines[guardRow][:guardCol] + 'X' + lines[guardRow][guardCol+1:]
            guardRow, guardCol = tempRow, tempCol
    return False

def partTwoDay06(filePath):
    lines = readFile(filePath)
    originalPath = partOneDay06(filePath)
    nbCycles = 0

    # skip first one because it's the initial position
    for obstaclePlacement in originalPath[1:]:
        guardRow, guardCol = originalPath[0]
        lines[obstaclePlacement[0]] = lines[obstaclePlacement[0]][:obstaclePlacement[1]] + '#' + lines[obstaclePlacement[0]][obstaclePlacement[1]+1:]
        nbCycles += 1 if isLoop(lines, guardRow, guardCol) else 0
        lines[obstaclePlacement[0]] = lines[obstaclePlacement[0]][:obstaclePlacement[1]] + '.' + lines[obstaclePlacement[0]][obstaclePlacement[1]+1:]
    return nbCycles
        