from common.util import readFile
from dataclasses import dataclass


FREE_MEMORY = -1

def expandMemory(memory):
    isFree = False
    memoryId = 0
    expanded = []

    for n in memory:
        if isFree:
            expanded += [FREE_MEMORY] * int(n)
        else:
            expanded += [memoryId] * int(n)
            memoryId += 1
        isFree = not isFree
    return expanded

def defragmentMemory(memory):
    i = 0
    defragmentedMemory = []

    while i < len(memory):
        if memory[i] == FREE_MEMORY:
            if memory[-1] != FREE_MEMORY:
                defragmentedMemory.append(memory[-1])
            else:
                i -= 1
            memory = memory[:-1]
        else:
            defragmentedMemory.append(memory[i])
        i += 1
    return defragmentedMemory

def partOneDay09(filePath):
    lines = readFile(filePath)
    memory = expandMemory(lines[0])
    defragmentedMemory = defragmentMemory(memory)

    return sum([i * defragmentedMemory[i] for i in range(len(defragmentedMemory))])

def partTwoDay09(filePath):
    lines = readFile(filePath)
    return 0