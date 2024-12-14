from common.util import readFile
from dataclasses import dataclass
from copy import copy

FREE_MEMORY = -1

@dataclass
class File:
    size: int
    content: int = FREE_MEMORY

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

def expandMemoryForFile(memory):
    isFree = False
    memoryId = 0
    expanded = []

    for n in memory:
        if isFree:
            expanded += [File(int(n))]
        else:
            expanded += [File(int(n), memoryId)]
            memoryId += 1
        isFree = not isFree
    return expanded

def getIndexByFileSize(memory, size, limit):
    i = 0
    while i < limit:
        if memory[i].content == FREE_MEMORY and memory[i].size >= size:
            return i
        i += 1
    return -1

def defragmentMemoryForFile(memory):
    i = len(memory) - 1

    while i > 0:
        if memory[i].content != FREE_MEMORY:
            freeSlotIndex = getIndexByFileSize(memory, memory[i].size, i)
            if freeSlotIndex >= 0:
                memory[freeSlotIndex].size -= memory[i].size
                memory.insert(freeSlotIndex, copy(memory[i]))
                memory[i+1].content = FREE_MEMORY
                if memory[freeSlotIndex + 1].size == 0:
                    del memory[freeSlotIndex+1]
        i -= 1
    return memory

def partTwoDay09(filePath):
    lines = readFile(filePath)
    memory = expandMemoryForFile(lines[0])
    defragmentedMemory = defragmentMemoryForFile(memory)

    i = 0
    checksum = 0
    for slot in defragmentedMemory:
        if slot.content != FREE_MEMORY:
            checksum += slot.size / 2 * (2*i + slot.size - 1) * slot.content
        i += slot.size
    return int(checksum)