# https://code.google.com/codejam/contest/2270488/dashboard#s=p3

FILENAME = "sample"

import sys
sys.stdin = open(FILENAME + ".in", 'r')
#sys.stdout = open(FILENAME + ".out", 'w')

def get_line(): return input()
def get_int(): return int(get_line())
def get_ints(): return [int(x) for x in get_line().split()]

def solve():
    a,b = get_ints()
    return "hola"


for case in range(get_int()):
    print('Case #%d: %s' % (case+1, solve()))


