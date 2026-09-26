# https://code.google.com/codejam/contest/2270488/dashboard#s=p1

FILENAME = "B-large-practice"

import numpy as np
from itertools import product

import sys
sys.stdin = open(FILENAME + ".in", 'r')
sys.stdout = open(FILENAME + ".out", 'w')

def get_line(): return input()
def get_int(): return int(get_line())
def get_ints(): return [int(x) for x in get_line().split()]


def solve():
    N, M = get_ints()
    matrix = np.matrix([get_ints() for x in range(N)])
    for y, x  in product(range(N),range(M)):
        val = matrix[y,x]
        maxrow = np.max(matrix[y,:])
        maxcol = np.max(matrix[:,x])
        if val < maxrow and val < maxcol: return "NO"
    return "YES"


for case in range(get_int()):
    print('Case #%d: %s' % (case+1, solve()))


