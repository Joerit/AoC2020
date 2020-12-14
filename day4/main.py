#!/bin/python3
from pprint import *

class Passport:
    fields: dict

    def __init__(self, entry):
        self.fields = {}
        ent = entry.split()
        for field in ent:
            spl = field.split(":")
            self.fields[spl[0]] = spl[1]
    
    
    # switch dict
    validators = {
        "byr" : (lambda val: int(val) > 1919 and int(val) < 2003),
        "iyr" : (lambda val: int(val) > 2009 and int(val) < 2021),
        "eyr" : (lambda val: int(val) > 2019 and int(val) < 2031),
        "hgt" : (lambda val: (val[-2:] == "cm" and 
                            ( int(val[:-2]) > 149 and int(val[:-2]) < 194) )
                        or val[-2:] == "in" and 
                            ( int(val[:-2]) > 58 and int(val[:-2]) < 77) ),
        "hcl" : (lambda val: val[0] == "#"
                        and val[1:].isalnum()),
        "ecl" : (lambda val: val in ["amb", "blu", "brn", "gry", "grn", "hzl", "oth"]),
        "pid" : (lambda val: len(val) == 9),
        "cid" : (lambda val: True),
    }
    def validateField(self, key, val):
        return self.validators.get(key, lambda val: False)(val)

    #def validateField(self, key, val):
    #    switch key:
    #        "byr" : 
    #            return int(val) > 1919 and int(val) < 2003
    #        "iyr" : 
    #            return int(val) > 2009 and int(val) < 2021
    #        "eyr" :
    #            return int(val) > 2019 and int(val) < 2031
     #       "hgt" : 
     #           return (val[-2:] == "cm" and 
     ##                       ( int(val[:-2]) > 149 and int(val[:-2]) < 194))
     #                   or ( val[-2:] == "in" and 
     #                       ( int(val[:-2]) > 58 and int(val[:-2]) < 77))
     #       "hcl" : 
     #           return val[0] == "#" and val[1:].isalnum()
     #       "ecl" : 
     #           return val in ["amb", "blu", "brn", "gry", "grn", "hzl", "oth"]
      #      "pid" : 
      #          return len(val) == 9
      #      "cid" : 
      ##          return True
      #      default:
      #          return False

    def isValid(self):
        present = "byr" in self.fields and "iyr" in self.fields and "eyr" in self.fields and "hgt" in self.fields and "hcl" in self.fields and "ecl" in self.fields and "pid" in self.fields
        if not present:
            return False
        for key in self.fields:
            if not self.validateField(key, self.fields[key]):
                return False
        return True



with open("input") as f:
    entries = f.read().split("\n\n")
    
total = 0
valids = 0
for entry in entries:
    total = total + 1
    pas = Passport(entry)
    if pas.isValid():
        valids = valids + 1

pprint(total)
pprint(valids)
