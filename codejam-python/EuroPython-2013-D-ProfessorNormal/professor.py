# https://code.google.com/codejam/contest/2845486/dashboard#s=p3

FILENAME = "D-small-practice-1"

import sys
sys.stdin = open(FILENAME + ".in", 'r')
sys.stdout = open(FILENAME + ".out", 'w')

def get_line(): return input()
def get_int(): return int(get_line())
def get_ints(): return [int(x) for x in get_line().split()]


import numpy as np
from itertools import product

incs = [ (-1,0), (1,0), (0,1), (0,-1) ]

def neigbours(y, x, maxy, maxx):
    ret = []
    for y,x in [(y+inc[0],x+inc[1]) for inc in incs]:
        if y>=0 and x>=0 and y<maxy and x<maxx:
            ret.append((y,x))
    return ret


def clean_players(matrix):
    M, N = matrix.shape[0] , matrix.shape[1]
    delete = []
    oldovers = 0
    newovers = 0
    has_players = False
    for y, x in product(range(M),range(N)):
        if matrix[y,x] == 0:
            oldovers += 1
        elif matrix[y,x] < 12 or sum(map(lambda pos: matrix[pos[0],pos[1]], neigbours(y, x, M, N))) == 0:
            newovers += 1
            delete.append((y,x))
        else:
            has_players = True
    for y,x in delete: matrix[y,x] = 0
    return matrix if has_players else None


def solve():
    M, N = get_int(), get_int()
    matrix = np.matrix([get_ints() for x in range(M)])
    turns = -1
    while not clean_players(matrix) == None:
        newmatrix = np.matrix(matrix)
        isolated = False
        for y, x in product(range(M),range(N)):
            if matrix[y,x]:
                good = list(filter(lambda pos: matrix[pos[0],pos[1]] > 0, neigbours(y, x, M, N)))
                newmatrix[y,x] -= 12
                count = len(good)
                if count:
                    for yy,xx in good: newmatrix[yy,xx] += 12/count
                else:
                    isolated = True
        if (matrix==newmatrix).all():
            return '%d children will play forever' % np.count_nonzero(matrix)

        dif = newmatrix - matrix
        jump = (12 - matrix) // dif
        jumppos = jump[jump > 0]
        accelerate = 1 if isolated or jumppos.size==0 else max(1, jumppos.min())
        accelerate = 1 # TODO: Get this right for all cases
        turns += accelerate

        matrix = matrix + (dif * accelerate)
    return '%d turns' % max(0,turns)


for case in range(get_int()):
    print('Case #%d: %s' % (case+1, solve()))

