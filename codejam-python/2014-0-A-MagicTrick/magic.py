# https://code.google.com/codejam/contest/2974486/dashboard#s=p0

FILENAME = "A-small-attempt0"

import sys
sys.stdin = open(FILENAME + ".in", 'r')
sys.stdout = open(FILENAME + ".out", 'w')

def get_line(): return input()
def get_int(): return int(get_line())
def get_ints(): return [int(x) for x in get_line().split()]

def solve():
    row1_sel = get_int() - 1
    row1 = [get_ints() for _ in range(4)]
    row2_sel = get_int() - 1
    row2 = [get_ints() for _ in range(4)]
    set_together = set(row1[row1_sel]) & set(row2[row2_sel])
    if len(set_together) == 0:
        return "Volunteer cheated!"
    elif len(set_together) > 1:
        return "Bad magician!"
    else:
        return list(set_together)[0]

for case in range(get_int()):
    print('Case #%d: %s' % (case+1, solve()))


