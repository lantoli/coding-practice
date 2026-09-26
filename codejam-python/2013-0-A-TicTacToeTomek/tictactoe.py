# https://code.google.com/codejam/contest/2270488/dashboard#s=p0

FILENAME = "A-large-practice"

import sys
sys.stdin = open(FILENAME + ".in", 'r')
sys.stdout = open(FILENAME + ".out", 'w')

def get_line(): return input()
def get_int(): return int(get_line())


def won(player, board):
    for row in board:
        if row.count(player) + row.count('T') == 4: return True

    for col in [ ''.join([board[y][x] for y in range(4)]) for x in range(4)]:
        if col.count(player) + col.count('T') == 4: return True

    diag1 = ''.join([board[x][x] for x in range(4)])
    if diag1.count(player) + diag1.count('T') == 4: return True

    diag2 = ''.join([board[3-x][x] for x in range(4)])
    if diag2.count(player) + diag2.count('T') == 4: return True

    return False

def solve():
    board = [ get_line() for x in range(4) ]
    get_line()
    if won('X', board): return "X won"
    elif won('O', board): return "O won"
    elif '.' in ''.join(board): return "Game has not completed"
    else: return 'Draw'


for case in range(get_int()):
    print('Case #%d: %s' % (case+1, solve()))
