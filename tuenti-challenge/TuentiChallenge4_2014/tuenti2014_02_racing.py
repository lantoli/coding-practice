#! /usr/bin/env python3
"""
Tuenti Challenge 4, Apr 2014, Challenge 2 - Tuenti Movil F1 Racing Team
"""

from collections import namedtuple

Coord = namedtuple('Coord', ['y', 'x'])
initial_direction = Coord(y=0, x=1)


def main():
    line = input()
    min_pos, max_pos = get_min_max_track_coordinates(line)
    track = build_empty_track(min_pos, max_pos)
    ini_pos = Coord(y=-min_pos.y, x=-min_pos.x)
    fill_track(track, line, ini_pos)
    for row in track:
        print(''.join(row))


def next_direction_and_pos(dir, pos, ch):
    if ch == '/':
        dir = Coord(y=-dir.x, x=-dir.y)
    elif ch == '\\':
        dir = Coord(y=dir.x, x=dir.y)
    elif ch != '#' and ch != '-':
        raise ValueError('invalid char')
    pos = Coord(y=pos.y + dir.y, x=pos.x + dir.x)
    return dir, pos


def get_min_max_track_coordinates(line):
    dir = initial_direction
    pos = min_pos = max_pos = Coord(y=0, x=0)
    for ch in line:
        dir, pos = next_direction_and_pos(dir, pos, ch)
        min_pos = Coord(y=min(min_pos.y, pos.y), x=min(min_pos.x, pos.x))
        max_pos = Coord(y=max(max_pos.y, pos.y), x=max(max_pos.x, pos.x))
    return min_pos, max_pos


def build_empty_track(min_pos, max_pos):
    size = Coord(y=max_pos.y - min_pos.y + 1, x=max_pos.x - min_pos.x + 1)
    return [ [' '] * size.x for _ in range(size.y)]


def fill_track(track, line, pos):
    dir = initial_direction
    for ch in line:
        track[pos.y][pos.x] = '|' if ch == '-' and dir.y else ch
        dir, pos = next_direction_and_pos(dir, pos, ch)


if __name__ == '__main__':
    main()
