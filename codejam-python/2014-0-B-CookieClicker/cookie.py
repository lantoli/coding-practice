# https://code.google.com/codejam/contest/2974486/dashboard#s=p1

FILENAME = "B-large"

import sys
sys.stdin = open(FILENAME + ".in", 'r')
sys.stdout = open(FILENAME + ".out", 'w')

def get_line(): return input()
def get_int(): return int(get_line())
def get_floats(): return [float(x) for x in get_line().split()]

def solve():
    C_farm_cost, F_farm_ratio, X_goal = get_floats()

    time = 0.0
    ratio = 2.0
    while True:
        time_keep = time + (X_goal / ratio)
        time_buy = time + (C_farm_cost / ratio) + (X_goal / (ratio + F_farm_ratio))
        if time_keep <= time_buy:
            return "{:.7f}".format(time_keep)
        time += C_farm_cost / ratio
        ratio += F_farm_ratio


for case in range(get_int()):
    print('Case #%d: %s' % (case+1, solve()))


