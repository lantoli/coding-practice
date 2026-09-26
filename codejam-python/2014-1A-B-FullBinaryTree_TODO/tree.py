# https://code.google.com/codejam/contest/2984486/dashboard#s=p1

FILENAME = "sample"

import sys
sys.stdin = open(FILENAME + ".in", 'r')
#sys.stdout = open(FILENAME + ".out", 'w')

def get_line(): return input()
def get_int(): return int(get_line())
def get_ints(): return [int(x) for x in get_line().split()]
def get_strs(): return [x for x in get_line().split()]

def cut(ad, root, forbidden=frozenset()):
    children = ad[root]
    num_children = len(children)
    if num_children == 0 or num_children == 2: return accum

    return root

def solve():
    N = get_int()
    ad = [set() for _ in range(N)]
    for _ in range(N-1):
        X,Y = get_ints()
        ad[X-1].add(Y-1)
        ad[Y-1].add(X-1)

    hola1 = frozenset()
    hola2 = frozenset(hola1, 124)
    hola3 = hola2.add(15)

    return min(len(cut(ad, root)) for root in range(N))

for case in range(get_int()):
    print('Case #%d: %s' % (case+1, solve()))


