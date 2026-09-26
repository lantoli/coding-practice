#! /usr/bin/env python3
"""
Tuenti Challenge 4, Apr 2014, Challenge 1 - Anonymous Poll
"""

def solve():
    info = input()
    return ','.join(sorted(data[info])) if info in data else "NONE"


data = dict()
with open('students', 'r') as f:
    for line in f:
        comma = line.index(",")
        name, info = line[:comma], line[comma+1:-1]
        if info not in data:
            data[info] = set()
        data[info].add(name)

for case in range(int(input())):
    print('Case #{}: {}'.format(case+1, solve()))

