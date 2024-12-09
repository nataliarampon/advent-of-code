from enum import Enum

class Direction(Enum):
    UP = 0
    RIGHT = 1
    DOWN = 2
    LEFT = 3

def readFile(fileLocation):
    with open(fileLocation, 'r') as file:
        return file.read().splitlines()

def isOutOfBounds(row, col, matrix):
    boundRow = len(matrix)
    boundCol = len(matrix[0])

    return row < 0 or col < 0 or row >= boundRow or col >= boundCol

def move(row, col, direction):
    match direction:
        case Direction.UP.value: return (row - 1, col)
        case Direction.RIGHT.value: return (row, col + 1)
        case Direction.DOWN.value: return (row + 1, col)
        case Direction.LEFT.value: return (row, col - 1)
