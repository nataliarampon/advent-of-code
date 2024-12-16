from common.util import readFile
from collections import defaultdict

def getDirectedGraphAndZeros(matrix):
    zeros = []
    graph = defaultdict(list)
    for i in range(len(matrix)):
        for j in range(len(matrix[0])):
            current_nb = int(matrix[i][j])
            if current_nb == 0:
                zeros.append((i,j))
            if i > 0:
                if int(matrix[i-1][j]) - current_nb == 1:
                    graph[(i,j)].append((i-1,j))
            if j > 0:
                if int(matrix[i][j-1]) - current_nb == 1:
                    graph[(i,j)].append((i,j-1))
            if j < len(matrix[0]) - 1:
                if int(matrix[i][j+1]) - current_nb == 1:
                    graph[(i,j)].append((i,j+1))
            if i < len(matrix) - 1:
                if int(matrix[i+1][j]) - current_nb == 1:
                    graph[(i,j)].append((i+1,j))
    return graph, zeros

def getReacheableNodes(graph, node, visited, pathLength):
    paths = 0
    visited.append(node)
    if pathLength == 9:
        return  1
    for nextNode in graph[node]:
        if nextNode not in visited:
            paths += getReacheableNodes(graph, nextNode, visited, pathLength + 1)
    return paths

def partOneDay10(filePath):
    lines = readFile(filePath)
    graph, zeros = getDirectedGraphAndZeros(lines)
    return sum([getReacheableNodes(graph, zero, [], 0) for zero in zeros])

def getPaths(graph, node, pathLength):
    paths = 0
    if pathLength == 9:
        return  1
    for nextNode in graph[node]:
            paths += getPaths(graph, nextNode, pathLength + 1)
    return paths

def partTwoDay10(filePath):
    lines = readFile(filePath)
    graph, zeros = getDirectedGraphAndZeros(lines)
    return sum([getPaths(graph, zero, 0) for zero in zeros])