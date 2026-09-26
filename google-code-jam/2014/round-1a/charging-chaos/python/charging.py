# https://code.google.com/codejam/contest/2984486/dashboard#s=p0

FILENAME = "A-large-practice"

import sys
sys.stdin = open(FILENAME + ".in", 'r')
sys.stdout = open(FILENAME + ".out", 'w')

def get_line(): return input()
def get_int(): return int(get_line())
def get_ints(): return [int(x) for x in get_line().split()]
def get_strs(): return [x for x in get_line().split()]

def bit_count(a):
    return bin(a).count("1")

def solve():
    N, L = get_ints()
    outlets = sorted([int(x,2) for x in get_strs()])
    devices = sorted([int(x,2) for x in get_strs()])

    table = sorted(list(set([o^d for d in devices for o in outlets])),key=bit_count)
    for t in table:
        outlets_new = sorted([t^o for o in outlets])
        if outlets_new == devices:
            return bit_count(t)
    return 'NOT POSSIBLE'

for case in range(get_int()):
    print('Case #%d: %s' % (case+1, solve()))


