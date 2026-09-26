# https://code.google.com/codejam/contest/351101/dashboard#s=p0

FILENAME = "A-large-practice"

import sys
sys.stdin = open(FILENAME + ".in", 'r')
sys.stdout = open(FILENAME + ".out", 'w')

def get_line(): return input()
def get_int(): return int(get_line())
def get_ints(): return [int(x) for x in get_line().split()]


def solve():
    C = get_int()
    I = get_int()
    P = get_ints()
    for i, x in enumerate(P):
        for j, y in enumerate(P):
            if i != j and x+y == C: return "%d %d" % (i+1, j+1)


for case in range(get_int()):
    print('Case #%d: %s' % (case+1, solve()))

