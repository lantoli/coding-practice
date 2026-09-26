# https://code.google.com/codejam/contest/2845486/dashboard#s=p1

FILENAME = "B-small-practice"

import math
import sys
sys.stdin = open(FILENAME + ".in", 'r')
sys.stdout = open(FILENAME + ".out", 'w')

def get_line(): return input()
def get_int(): return int(get_line())
def get_ints(): return [int(x) for x in get_line().split()]

def solve():
    # D  =  v * v * sin(2*ANG) / g
    V, D = get_ints()
    calc = D * 9.8 / (V*V)
    return math.degrees(math.asin( min(1.0,calc) ) / 2)

for case in range(get_int()):
    print('Case #%d: %s' % (case+1, solve()))

