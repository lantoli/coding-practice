#! /usr/bin/env python3

"""
https://www.facebook.com/hackercup/problems.php?pid=403525256396727&round=185564241586420
"""

FILENAME = "practice"

import sys
sys.stdin = open(FILENAME + ".in", 'r')
sys.stdout = open(FILENAME + ".out", 'w')


def solve(text):
    min_pars, max_pars = 0, 0
    smiley = False
    for ch in text:
        if ch == '(':
            max_pars += 1
            if not smiley:
                min_pars += 1
        elif ch == ')':
            if min_pars > 0: min_pars -= 1
            if not smiley:
                max_pars -= 1
                if max_pars < 0:
                    return False
        smiley = ch == ':'
    return min_pars == 0


def get_line(): return input()
def get_int(): return int(get_line())


if __name__ == '__main__':
    for case in range(get_int()):
        print('Case #{}: {}'.format(case + 1, "YES" if solve(get_line()) else "NO"))
