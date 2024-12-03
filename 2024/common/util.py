def readFile(fileLocation):
    with open(fileLocation, 'r') as file:
        return file.readlines()
        