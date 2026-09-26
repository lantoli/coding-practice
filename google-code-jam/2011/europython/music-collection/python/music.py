#! /usr/bin/env python3

"""
Solution to: https://code.google.com/codejam/contest/1277486/dashboard#s=p1
"""

FILENAME = "B-small-practice-2"

import sys
sys.stdin = open(FILENAME + ".in", 'r')
sys.stdout = open(FILENAME + ".out", 'w')


def optimal(song, others):
    if not others: return '""'
    for size in range(1, len(song) + 1):
        candidates = [song[i:(i + size)] for i in range(len(song) + 1 - size)
                      if all([song[i:(i + size)] not in other for other in others])]
        if candidates:
            return '"' + sorted(candidates)[0] + '"'
    return ':('


def solve(songs):
    return [optimal(song, songs[:i] + songs[(i + 1):]) for i, song in enumerate(songs)]


def get_line(): return input()
def get_int(): return int(get_line())


if __name__ == '__main__':
    for case in range(get_int()):
        print('Case #{}:\n{}'.format(case + 1,
                    '\n'.join(solve([get_line().upper() for _ in range(get_int())]))))
