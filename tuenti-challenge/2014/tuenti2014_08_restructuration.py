#! /usr/bin/env python3
"""
Tuenti Challenge 4, Apr 2014, Challenge 8 - Tuenti Restructuration

All tables are changed so source is always [0, 1, 2, 3, 4, 5, 6, 7, 8] and
target is recalculated to take that into account.
Unidimensional arrays are used in order to have efficient copies to generate neighbours tables (swaping persons).
All possible targets are precalculated before starting to process input.
"""

from collections import deque


DIM = 3
SIZE = DIM * DIM
solutions = dict()


def main():
    calculate_solutions()
    for case in range(int(input())):
        source, target = get_table(), get_table()
        board = [source.index(x) for x in target]
        print(solve(board))


def get_table():
    input()
    return [x.strip() for _ in range(DIM) for x in input().split(',')]


def solve(board):
    return solutions[str(board)] if str(board) in solutions else -1


def calculate_solutions():
    board_ini = list(range(SIZE))
    solutions[str(board_ini)] = 0
    search_tree = deque([ board_ini ])
    while search_tree:
        board = search_tree.popleft()
        next_day = solutions[str(board)] + 1
        for next in neighbours(board):
            if str(next) not in solutions:
                solutions[str(next)] = next_day
                search_tree.append(next)


def neighbours(board):
    for pos in range(SIZE):
        if (pos+1) % DIM != 0:  # horizontal swap
            yield new_board(board, pos, pos+1)
        if pos+DIM < SIZE:      # vertical swap
            yield new_board(board, pos, pos+DIM)


def new_board(board, pos1, pos2):
    ret = list(board)
    ret[pos1], ret[pos2] = ret[pos2], ret[pos1]
    return ret


if __name__ == '__main__':
    main()

