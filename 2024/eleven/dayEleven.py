from common.util import readFile
from collections import Counter

memoization = {}

def blink(pebbles):
    i = 0
    while i < len(pebbles):
        pebble = pebbles[i]
        length = len(pebble)
        if pebble == '0':
            pebbles[i] = '1'
        elif length % 2 == 0:
            pebbles[i] = str(int(pebble[length // 2:]))
            pebbles.insert(i, str(int(pebble[:length // 2])))
            i += 1
        else:
            pebbles[i] = str(int(pebble) * 2024)
        i += 1
    return pebbles

def blinkMemoization(pebbles):
    newState = Counter()
    for pebble, freq in pebbles.items():
        if pebble:
            if pebble not in memoization:
                length = len(pebble)
                if pebble == '0':
                    memoization[pebble] = ['1', None]
                elif length % 2 == 0:
                    memoization[pebble] = [str(int(pebble[length // 2:])), str(int(pebble[:length // 2]))]
                else:
                    memoization[pebble] = [str(int(pebble) * 2024), None]
            if memoization[pebble][0] != memoization[pebble][1]:
                newState.update(Counter({memoization[pebble][0]:freq, memoization[pebble][1]: freq}))
            else:
                newState.update(Counter({memoization[pebble][0]:2*freq}))
    del newState[None]
    return newState

def partOneDay11(filePath, n_blinks):
    pebbles = readFile(filePath)[0].split()
    pebbles = Counter(pebbles)

    for i in range(n_blinks):
        pebbles = blinkMemoization(pebbles)
    return pebbles.total()