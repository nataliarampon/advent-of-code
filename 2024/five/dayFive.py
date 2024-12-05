from common.util import readFile

def processRules(file):
    rulesBefore = {}
    rulesAfter = {}
    i = 0

    while file[i] != '':
        before, after = file[i].split('|')
        rulesBefore[before] = [after] if before not in rulesBefore else rulesBefore[before] + [after]
        rulesAfter[after] = [before] if after not in rulesAfter else rulesAfter[after] + [before]
        i += 1
    return rulesBefore, rulesAfter, i + 1

def isUpdateValid(rulesBefore, rulesAfter, update):
    for i in range(len(update) - 1):
        if (update[i] in rulesBefore and update[i+1] not in rulesBefore[update[i]]) or \
            (update[i+1] in rulesAfter and update[i] not in rulesAfter[update[i+1]]):
            return False
    return True

def partOneDay05(filePath):
    lines = readFile(filePath)
    rulesBefore, rulesAfter, current_file_index = processRules(lines)

    sum = 0
    while current_file_index < len(lines):
        update = lines[current_file_index].split(',')
        if isUpdateValid(rulesBefore, rulesAfter, update):
            sum += int(update[len(update)//2])
        current_file_index += 1
    return sum

def partTwoDay05(filePath):
    lines = readFile(filePath)

    return 0
