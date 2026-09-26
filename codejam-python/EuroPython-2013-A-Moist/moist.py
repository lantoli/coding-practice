# https://code.google.com/codejam/contest/2845486/dashboard#s=p0

FILENAME = "A-small-practice-2"

import sys
sys.stdin = open(FILENAME + ".in", 'r')
sys.stdout = open(FILENAME + ".out", 'w')

def get_line(): return input()
def get_int(): return int(get_line())


def solve():
    N = get_int()
    names = [get_line() for x in range(N)]
    inv = 0
    for i in range(1,N):
        if names[i-1] > names[i]:
            names[i-1], names[i] = names[i], names[i-1]
            inv += 1
    return inv

for case in range(get_int()):
    print('Case #%d: %s' % (case+1, solve()))

