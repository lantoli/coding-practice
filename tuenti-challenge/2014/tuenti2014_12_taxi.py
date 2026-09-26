#! /usr/bin/env python3
"""
Tuenti Challenge 4, Apr 2014, Challenge 12 - Taxi Driver
"""

from collections import namedtuple, deque


Coord = namedtuple('Coord', ['row', 'col'])
State = namedtuple('State', ['coord', 'dir'])


North = Coord(-1, 0)
South = Coord(1, 0)
East = Coord(0, 1)
West = Coord(0, -1)
dirs = [North, East, South, West]


def turn_right(dir):
    return dirs[ (dirs.index(dir)+1) % len(dirs) ]


def solve():
    cols, rows = get_ints()
    walls = [[False] * cols for _ in range(rows)]
    for row in range(rows):
        for col, elm in enumerate(get_line()):
            if elm == 'S':
                start = Coord(row, col)
            elif elm == 'X':
                end = Coord(row, col)
            elif elm == '#':
                walls[row][col] = True

    distances = dict()
    bfs = deque()  # breadht-first search tree
    for d in dirs:
        state = State(start, d)
        bfs.append(state)
        distances[state] = 0

    while bfs:
        state = bfs.popleft()
        next_distance = distances[state] + 1
        next_pos = Coord(state.coord.row + state.dir.row, state.coord.col + state.dir.col)

        if 0 <= next_pos.row < rows and 0 <= next_pos.col < cols:
            if next_pos == end:
                return next_distance
            elif not walls[next_pos.row][next_pos.col]:
                child1 = State(next_pos, state.dir)
                child2 = State(next_pos, turn_right(state.dir))
                if child1 not in distances:
                    bfs.append(child1)
                    distances[child1] = next_distance
                if child2 not in distances:
                    bfs.append(child2)
                    distances[child2] = next_distance

    return "ERROR"


def get_line(): return input()
def get_int(): return int(get_line())
def get_ints(): return [int(x) for x in get_line().split()]


if __name__ == '__main__':
    for case in range(get_int()):
        print('Case #{}: {}'.format(case+1, solve()))
