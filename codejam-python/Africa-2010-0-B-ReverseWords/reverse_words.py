# https://code.google.com/codejam/contest/351101/dashboard#s=p1

FILENAME = "B-large-practice"

import sys
sys.stdin = open(FILENAME + ".in", 'r')
sys.stdout = open(FILENAME + ".out", 'w')

def get_line(): return input()
def get_int(): return int(get_line())
def get_strs(): return [x for x in get_line().split()]


def solve(): return ' '.join(reversed(get_strs()))

for case in range(get_int()):
    print('Case #%d: %s' % (case+1, solve()))

