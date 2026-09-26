FILENAME = "C-small-attempt2"

import sys
sys.stdin = open(FILENAME + ".in", 'r')
sys.stdout = open(FILENAME + ".out", 'w')


def construct(r, c, mines):
    if mines == r * c - 1:
        result = []
        for y in range(r):
            result.append([])
            for x in range(c):
                result[y].append('*')
        result[0][0] = 'c'
        return result
    free = r * c - mines
    if r == 1:
        result = [['c']]
        for i in range(free - 1):
            result[0].append('.')
        for i in range(mines):
            result[0].append('*')
        return result
    for h in range(2, r + 1):
        for w in range(h, c + 1):
            if h * w >= free:
                if h * w == free:
                    gow = 0
                    goh = 0
                else:
                    gow = min(h * w - free, w - 2)
                    goh = h * w - free - gow + 1
                if gow > w - 2 or goh > h - 2:
                    continue
                result = []
                for y in range(r):
                    result.append([])
                    for x in range(c):
                        if y == 0 and x == 0:
                            cur = 'c'
                        elif y < h - 1 and x < w - 1:
                            cur = '.'
                        elif y == h - 1 and x < w - gow:
                            cur = '.'
                        elif x == w - 1 and y < h - goh:
                            cur = '.'
                        else:
                            cur = '*'
                        result[y].append(cur)
                return result
    return []

def transpose(x):
    if not x:
        return x
    result = []
    for i in range(len(x[0])):
        cur = []
        for j in range(len(x)):
            cur.append(x[j][i])
        result.append(cur)
    return result

d = int(raw_input())

for case_no in range(1, d + 1):
    r, c, mines = map(int, raw_input().split())
    if r > c:
        result = transpose(construct(c, r, mines))
    else:
        result = construct(r, c, mines)
    print "Case #%d:" % case_no
    if not result:
        print 'Impossible'
    else:
        for y in range(r):
            for x in range(c):
                sys.stdout.write(result[y][x])
            print
