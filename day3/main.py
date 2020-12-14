#!/bin/python3
from pprint import *

field = []
with open("input") as f:
    for line in f:
        field.append(line[0:-1])

pprint(field)

def checkfield(field, xstep, ystep):
    x = 0
    y = 0
    trees = 0

    while y < (len(field) - 1):
        if field[y][x] == "#":
            trees = trees + 1
        x = x + xstep
        y = y + ystep
        if x >= len(field[y]):
            x = x - len(field[y])
    return trees

print(checkfield(field, 1, 1))
print(checkfield(field, 3, 1))
print(checkfield(field, 5, 1))
print(checkfield(field, 7, 1))
print(checkfield(field, 1, 2))
