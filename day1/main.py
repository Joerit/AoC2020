#!/bin/python3

def solve2(intlist):
    for int1 in ints:
        for int2 in ints:
            if int1 + int2 >= 2020:
                continue
            for int3 in ints:
                if (int1 + int2 + int3) == 2020:
                    print(str(int1 * int2 * int3))
                    quit()

def solve1(intlist):
    for int1 in ints:
        for int2 in ints:
           if (int1 + int2) == 2020:
               print(str(int1 * int2))
               quit()


with open("input") as f:
    print("opened file")
    ints = []
    for line in f:
        ints.append(int(line))

solve2(ints)
