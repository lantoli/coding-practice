#! /usr/bin/env python3
"""
Tuenti Challenge 4, Apr 2014, Challenge 5 - Tribblemaker
"""

from bitarray import bitarray


N = 8


def main():
    grid = [[x=='X' for x in input()] for _ in range(N)]
    list_grids = [serialize_grid(grid)]
    while True:
        grid = next_generation(grid)
        grid_ser = serialize_grid(grid)
        for idx, grid_cur in enumerate(list_grids):
            if grid_ser == grid_cur:
                print("{} {}".format(idx, len(list_grids) - idx))
                exit(0)
        list_grids.append(grid_ser)


def next_generation(grid):
    next_grid = [[False for _ in range(N)] for _ in range(N)]
    for y,x in cross_N():
        count = neighbours(grid, y, x)
        next_grid[y][x] = count == 3 or (count == 2 and grid[y][x])
    return next_grid


def neighbours(grid, y, x):
    total = 0
    for yinc,xinc in cross_inc():
        yn, xn = y+yinc, x+xinc
        if 0 <= yn < N and 0 <= xn < N and (yn != y or xn != x):
            if grid[yn][xn]: total += 1
    return total


def serialize_grid(grid):
    return bitarray([grid[y][x] for y,x in cross_N()])


def cross(r):
    for y in r:
        for x in r:
            yield (y, x)


def cross_N(): return cross(range(N))
def cross_inc(): return cross([-1, 0, 1])


if __name__ == '__main__':
    main()
