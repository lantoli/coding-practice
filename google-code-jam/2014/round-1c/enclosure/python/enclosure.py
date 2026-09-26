# https://code.google.com/codejam/contest/3004486/dashboard#s=p2

FILENAME = "C-large-practice"
#FILENAME = "sample"

import sys
import math
sys.stdin = open(FILENAME + ".in", 'r')
sys.stdout = open(FILENAME + ".out", 'w')

def get_line(): return input()
def get_int(): return int(get_line())
def get_ints(): return [int(x) for x in get_line().split()]
def get_strs(): return [x for x in get_line().split()]

def count_unique(board, N, M):
    total = 0
    for y in range(1, N-1):
        for x in range(1, M-1):
            if board[y][x] and board[y-1][x] and board[y+1][x] and board[y][x-1] and board[y][x+1]:
                total +=1

    return count_non(board) - total

def count_non(board):
    return sum([item for sublist in board for item in sublist])

def solve():
    N, M, K = get_ints()
    M, N = min(N,M), max(N,M)

    board = [ [False for _ in range(M)] for _ in range(N) ]

    sq = int(math.sqrt(K))
#    sq = int(sq) + 1

    for y in range(1, N-1):
        for x in range(1, min(sq+1, M-1)):
            board[y][x] = True
            board[y-1][x] = True
            board[y+1][x] = True
            board[y][x-1] = True
            board[y][x+1] = True
            if count_non(board) >= K: return count_unique(board, N, M)

    for y in range(N):
        for x in range(M):
            if not board[y][x]:
                board[y][x] = True
                if count_non(board) >= K: return count_unique(board, N, M)

for case in range(get_int()):
    print('Case #%d: %s' % (case+1, solve()))


