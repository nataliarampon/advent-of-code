from five.dayFive import partOneDay05
from four.dayFour import partOneDay04, partTwoDay04
from three.dayThree import partOneDay03, partTwoDay03
from two.dayTwo import partOneDay02, partTwoDay02
from one.dayOne import partOneDay01, partTwoDay01

if __name__ == "__main__":
    print(f"Day 01 Part 01: {partOneDay01('one/input.txt')}")
    print(f"Day 01 Part 02: {partTwoDay01('one/input.txt')}")

    print(f"Day 02 Part 01: {partOneDay02('two/input.txt')}")
    print(f"Day 02 Part 02: {partTwoDay02('two/input.txt')}")

    print(f"Day 03 Part 01: {partOneDay03('three/input.txt')}")
    print(f"Day 03 Part 02: {partTwoDay03('three/input.txt')}")

    print(f"Day 04 Part 01: {partOneDay04('four/input.txt')}")
    print(f"Day 04 Part 02: {partTwoDay04('four/input.txt')}")

    print(f"Day 05 Part 01: {partOneDay05('five/input.txt')}")