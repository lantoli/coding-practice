# https://code.google.com/codejam/contest/2974486/dashboard#s=p3

FILENAME = "D-large"

import sys
sys.stdin = open(FILENAME + ".in", 'r')
sys.stdout = open(FILENAME + ".out", 'w')

def get_line(): return input()
def get_int(): return int(get_line())
def get_floats(): return [float(x) for x in get_line().split()]

def solve_deceitful(naomi, ken):
    count = 0
    for n in sorted(naomi):
        kk = [k for k in ken if k<n]
        if len(kk) > 0:
            count += 1
            k = max(kk)
        else:
            k = max(ken)
        ken.remove(k)
    return count

def solve_war(naomi, ken):
    count = 0
    for n in sorted(naomi, reverse=True):
        kk = [k for k in ken if k>n]
        if len(kk) > 0:
            k = min(kk)
        else:
            k = min(ken)
            count += 1
        ken.remove(k)
    return count


def solve():
    N = get_int()
    naomi = get_floats()
    ken = get_floats()
    return "{} {}".format(solve_deceitful(list(naomi), list(ken)), solve_war(list(naomi), list(ken)))


for case in range(get_int()):
    print('Case #%d: %s' % (case+1, solve()))


