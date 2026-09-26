# https://code.google.com/codejam/contest/2974486/dashboard#s=p2

FILENAME = "C-small-attempt2"

import sys
sys.stdin = open(FILENAME + ".in", 'r')
sys.stdout = open(FILENAME + ".out", 'w')

def get_line(): return input()
def get_int(): return int(get_line())
def get_ints(): return [int(x) for x in get_line().split()]

def solve():
    R, C, M = get_ints()

    res = [['.' for _ in range(C)] for _ in range(R)]
    res[R-1][C-1] = 'c'
    mines = 0
    for row in range(R):
        for col in range(C):
            if mines < M and (row < R-2 or col < C-2 or M+1 == R*C):
                mines += 1
                res[row][col] = '*'
    return res if mines == M else None

for case in range(get_int()):
    print('Case #%d:' % (case+1))
    ans = solve()
    if ans:
        for r in ans: print (" ".join(r))
    else:
        print('Impossible')


