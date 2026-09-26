# https://code.google.com/codejam/contest/2845486/dashboard#s=p2

FILENAME = "C-small-practice-2"

import sys
sys.stdin = open(FILENAME + ".in", 'r')
sys.stdout = open(FILENAME + ".out", 'w')

def get_line(): return input()
def get_int(): return int(get_line())
def get_strs(): return [x for x in get_line().split()]


def solve_rec(list, set1, set2):
    if len(list) == 0: return True
    newlist = []
    for a,b in list:
        if a in set1:
            if b in set1: return False
            else: set2.add(b)
        elif a in set2:
            if b in set2: return False
            else: set1.add(b)
        elif b in set1: set2.add(a)
        elif b in set2: set1.add(a)
        else: newlist.append((a,b))
    if len(list) == len(newlist):
        a, b = list.pop()
        set1.add(a)
        set2.add(b)
        return solve_rec(list, set1, set2)
    else:
        return solve_rec(newlist, set1, set2) if newlist else True


def solve():
    list = [get_strs() for x in range(get_int())]
    return "Yes" if solve_rec(list, set(), set()) else "No"

for case in range(get_int()):
    print('Case #%d: %s' % (case+1, solve()))

