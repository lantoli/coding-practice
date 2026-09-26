# https://code.google.com/codejam/contest/90101/dashboard#s=p1

FILENAME = "B-large-practice"

import sys

sys.stdin = open(FILENAME + ".in", 'r')
sys.stdout = open(FILENAME + ".out", 'w')

def get_line(): return input()
def get_int(): return int(get_line())
def get_ints(): return [int(x) for x in get_line().split()]


import numpy as np


def cross(a, b):
    for i in range(a):
        for j in range(b):
            yield (i, j)

def neighbours(ui, uj, m, n):
    if ui - 1 >= 0: yield (ui - 1, uj)
    if uj - 1 >= 0: yield (ui, uj - 1)
    if uj + 1 < n: yield (ui, uj + 1)
    if ui + 1 < m: yield (ui + 1, uj)


def solve_basin(matrix, res, H, W, yini, xini, color):
    if res[yini,xini] > 0:
        return res[yini,xini]
    min = matrix[yini,xini]
    for y,x in neighbours(yini, xini, H, W):
        if matrix[y,x] < min:
            min = matrix[y,x]
            ymin, xmin = y,x
    if min ==  matrix[yini,xini]:
        res[yini,xini] = color+1
    elif res[ymin,xmin] > 0:
        res[yini,xini] = res[ymin,xmin]
    else:
        res[yini,xini] = solve_basin(matrix, res, H, W, ymin, xmin, color)

    return res[yini,xini]


def solve():
    H, W = get_ints()
    matrix = np.matrix([get_ints() for x in range(H)])
    res = np.zeros_like(matrix)

    color = 0
    for y,x in cross(H,W):
        color = solve_basin(matrix, res, H, W, y, x, color)

    return '\n' + '\n'.join(' '.join(chr(ord('a') - 1 + res[row, col]) for col in range(W)) for row in range(H))


for case in range(get_int()):
    print('Case #%d: %s' % (case + 1, solve()))

