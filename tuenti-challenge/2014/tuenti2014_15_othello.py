#! /usr/bin/env python3
"""
Tuenti Challenge 4, Apr 2014, Challenge 15 - Take a corner

To make it mor efficient boards are formatted in a 128-bit int, first 64 bits are for empty/occupied squares.
Next 64 bits are for white (clear) or black (set), applicable if it's occupied.
Order is a1 b1 ... h1 a2 ... a8 ... h8
We always start with White, if input is Black we switch the square colors to make it White.
"""


def solve():
    color, moves = get_line().split(' in ')
    board = get_board(switchColors=(color == 'Black'))
    for move_white, board_white in play(board, black=False, nodict=False).items():
        if can_take_corners(board_white, int(moves)-2):
            return label(move_white)
    return "Impossible"


def get_board(switchColors):
    pos_fill = 1
    pos_color = 1 << SIZE
    board = 0
    for square in [x for _ in range(DIM) for x in get_line()]:
        if not square == '.':
            board |= pos_fill
            if (square == 'X') != switchColors:
                board |= pos_color
        pos_fill <<= 1
        pos_color <<= 1
    return board


def play(board, black, nodict=True):
    ret = [] if nodict else dict()
    for y, x in cross(range(DIM)):
        if empty(board, y, x): # must be in empty square
            boardnew = board
            for ynew, xnew in neighbours(y, x):
                if filled(board, ynew, xnew, not black): # must sandwich opponent discs
                    good = False
                    boardtentative = boardnew
                    boardtentative = fill(boardtentative, y, x, black) # square selected
                    for yline, xline in line(y, x, ynew, xnew):
                        if empty(board, yline, xline):
                            break
                        if filled(board, yline, xline, black):
                            good  = True
                            break
                        boardtentative = fill(boardtentative, yline, xline, black)
                    if good:
                        boardnew = boardtentative
                        pass
            if not board == boardnew:
                if nodict:
                    ret.append(boardnew)
                else:
                    ret[pos(y,x)] = boardnew
    if (not ret) and nodict : # no legal movements so we skip
        ret.append(board)
    return ret


def can_take_corners(board, moves):
    for board_black in play(board, black=True):
        found = False
        if moves:
            for board_white in play(board_black, black=False):
                if can_take_corners(board_white, moves-1):
                    found = True
                    break
        else:
            for pos_white, _ in play(board_black, black=False, nodict=False).items():
                if pos_white in pos_corners:
                    found = True
                    break
        if not found: return False
    return True


def empty(board, y, x):
    return not (board & (1 << pos(y,x)))


def filled(board, y, x, black):
    return (not empty(board, y, x)) and bool(board & (1 << (pos(y,x)+SIZE))) == black


def fill(board, y, x, black):
    board |= (1 << pos(y,x)) # set occupied bit
    if black:
        board |= (1 << (pos(y,x))+SIZE)  # set bit
    else:
        board &= ~(1 << (pos(y,x))+SIZE) # clear bit
    return board


def neighbours(y, x):
    for yinc, xinc in cross([-1, 0, 1]):
        ynew, xnew = y+yinc, x+xinc
        if 0 <= ynew < DIM and 0 <= xnew < DIM and (yinc != 0 or xinc != 0):
            yield ynew, xnew


def line(y, x, ynew, xnew):
    yinc, xinc = ynew-y, xnew-x
    while True:
        y += yinc
        x += xinc
        if 0 <= y < DIM and 0 <= x < DIM:
            yield y, x
        else:
            return


DIM = 8
SIZE = DIM * DIM
pos_corners = tuple([0, DIM-1, DIM*(DIM-1), DIM*DIM-1])
def pos(y, x): return y * DIM + x
def ycoord(pos): return pos // DIM
def xcoord(pos): return pos % DIM
def label(pos): return 'abcdefgh'[xcoord(pos)] + '12345678'[ycoord(pos)]


def get_line(): return input()
def get_int(): return int(get_line())


def cross(r):
    for y in r:
        for x in r:
            yield (y, x)


def print_board(board):
    pos_fill = 1
    pos_color = 1 << SIZE
    for elm in range(SIZE):
        if board & pos_fill:
            if board & pos_color:
                print('X', end=' ')
            else:
                print('O', end=' ')
        else:
            print('.', end=' ')
        if elm % DIM == DIM - 1: print()
        pos_fill <<= 1
        pos_color <<= 1


if __name__ == '__main__':
    for case in range(get_int()):
        print(solve())
