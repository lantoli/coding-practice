# https://code.google.com/codejam/contest/3004486/dashboard#s=p0

FILENAME = "A-large-practice"

import sys
sys.stdin = open(FILENAME + ".in", 'r')
sys.stdout = open(FILENAME + ".out", 'w')

from fractions import Fraction

def get_line(): return input()
def get_int(): return int(get_line())


def solve():
    num = Fraction(get_line())

    if bin(num.denominator).count("1") != 1:  # Q must be a power of 2.
        return "impossible"

    for x in range(1,41):
        if not num.denominator % 2 == 0:
            return "impossible"
        if num.numerator >= num.denominator//2:
            return x
        num = Fraction(num.numerator, num.denominator//2)

    return "impossible"

for case in range(get_int()):
    print('Case #%d: %s' % (case+1, solve()))


