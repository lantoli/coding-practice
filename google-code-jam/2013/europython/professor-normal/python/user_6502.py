import sys

FILENAME = "D-small-practice-2"

import sys
sys.stdin = open(FILENAME + ".in", 'r')
sys.stdout = open(FILENAME + ".out", 'w')



n = int(sys.stdin.readline())

for case in range(n):
    m = int(sys.stdin.readline())
    n = int(sys.stdin.readline())
    incompatible = {}
    matrix = []
    for r in range(m):
        matrix.append(map(int, sys.stdin.readline().split(" ")))

    number_of_exchanges = 0
    while True:
        ok = 0
        for y in range(m):
            for x in range(n):
                if matrix[y][x] < 12:
                    matrix[y][x] = -1
        nhc = [[0]*n for i in range(m)]
        for y in range(m):
            for x in range(n):
                if matrix[y][x] > 0:
                    nh = ((y > 0 and matrix[y-1][x] >= 0) +
                          (y < m-1 and matrix[y+1][x] >= 0) +
                          (x > 0 and matrix[y][x-1] >= 0) +
                          (x < n-1 and matrix[y][x+1] >= 0) +
                          0)
                    if nh == 0:
                        matrix[y][x] = -1
                    else:
                        ok += 1
                        nhc[y][x] = nh
        if ok == 0:
            break

        before = [x[:] for x in matrix]

        for y in range(m):
            for x in range(n):
                if matrix[y][x] > 0:
                    c = 12 / nhc[y][x]
                    matrix[y][x] -= 12
                    if y > 0 and matrix[y-1][x] >= 0:
                        matrix[y-1][x] += c
                    if y < m-1 and matrix[y+1][x] >= 0:
                        matrix[y+1][x] += c
                    if x > 0 and matrix[y][x-1] >= 0:
                        matrix[y][x-1] += c
                    if x < n-1 and matrix[y][x+1] >= 0:
                        matrix[y][x+1] += c

        number_of_exchanges += 1
        if matrix == before:
            number_of_exchanges = "forever"
            break

        delta = [[matrix[y][x] - before[y][x] for x in range(n)]
                 for y in range(m)]

        def chcount(y, x):
            if before[y][x] == 0 or delta[y][x] >= 0:
                return 1E100
            else:
                return (matrix[y][x] - 12) // -delta[y][x]

        k = min(chcount(y, x) for y in range(m) for x in range(n))

        if k > 0:
            number_of_exchanges += k
            for y in xrange(m):
                for x in xrange(n):
                    if matrix[y][x] >= 0:
                        matrix[y][x] += k * delta[y][x]
                        assert matrix[y][x] >= 0

    print ("Case #%i: %s" % (case + 1,
                            "%i children will play forever" % ok
                            if number_of_exchanges == "forever"
                            else
                            "%i turns" % number_of_exchanges))
