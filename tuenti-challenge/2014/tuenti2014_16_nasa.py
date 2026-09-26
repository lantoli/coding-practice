#! /usr/bin/env python3
"""
Tuenti Challenge 4, Apr 2014, Challenge 16 - NASA

To avoid sqrt we can check (x2-x1)**2 + (y1-y2)**2 < (r1+r2)**2
We divide the space in squares so we can limit collision detection to close squares only.
"""

from collections import namedtuple

Point = namedtuple('Point', ['pos','x', 'y', 'r'])

MAX_XY = 100000
MAX_R = 500
CELL_SIZE = MAX_R
CELL_NUM = MAX_XY // CELL_SIZE


def main():
    start, count = map(int, input().split(','))
    start -= 1 # 0-index

    cells = [ [set() for _ in range(CELL_NUM)] for _ in range(CELL_NUM) ]
    points = []
    with open('points', 'r') as f:
        for _ in range(start):
            next(f)
        for pos in range(count):
            x, y, r = map(int, next(f).split())
            p = Point(pos, x, y, r)
            points.append(p)
            for cell in get_cells(cells, p):
                cell.add(p)

    total = 0
    for p1 in points:
        collect = set()
        for cell in get_cells(cells, p1):
            cell.remove(p1)
            for p2 in cell:
                if p1.pos < p2.pos and p2 not in collect and collision(p1, p2):
                    collect.add(p2)
        total += len(collect)
    print(total)


def collision(p1, p2):
    x = p1.x - p2.x
    y = p1.y - p2.y
    r = p1.r + p2.r
    return x*x + y*y < r*r


def get_cells(cells, p):
    ymin = max(0, (p.y - p.r) // CELL_SIZE)
    ymax = min(CELL_NUM-1, (p.y + p.r) // CELL_SIZE)

    xmin = max(0, (p.x - p.r) // CELL_SIZE)
    xmax = min(CELL_NUM-1, (p.x + p.r) // CELL_SIZE)

    for y in range(ymin, ymax + 1):
        for x in range(xmin, xmax + 1):
            yield cells[y][x]


if __name__ == '__main__':
    main()
