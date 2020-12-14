#!/bin/python3
from pprint import *

def validate1( line ):
    sline = line.split(":")
    ssline = sline[0].split(" ")
    char = ssline[1]
    charrange = ssline[0].split("-")
    charrange = list(map(int, charrange))
    pprint(charrange)
    print(char)
    charcount = 0
    print(sline[1])
    for c in sline[1]:
        if c == char:
            charcount = charcount + 1
    if (charcount >= charrange[0]) and (charcount <= charrange[1]):
        return 1
    else:
        return 0

def validate2( line ):
    sline = line.split(":")
    ssline = sline[0].split(" ")
    char = ssline[1]
    charrange = ssline[0].split("-")
    charrange = list(map(int, charrange))
    pprint(charrange)
    print(char)
    charcount = 0
    print(sline[1])
    if (sline[1][charrange[0]] == char) != (sline[1][charrange[1]] == char):
        print("valid")
        return 1
    else:
        return 0


with open("input") as f:
    valids = 0
    for line in f:
        valids = valids + validate2(line)
    print(valids)
