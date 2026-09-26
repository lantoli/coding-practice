# https://code.google.com/codejam/contest/90101/dashboard#s=p2

FILENAME = "C-large-practice"

import sys

sys.stdin = open(FILENAME + ".in", 'r')
sys.stdout = open(FILENAME + ".out", 'w')

def get_line(): return input()
def get_int(): return int(get_line())

welcome = "welcome to code jam"

def solve():
    sentence = get_line()
    table = [0] * len(welcome)
    for i in range(0, len(sentence)):
        for j in range(len(welcome)):
            if sentence[i] == welcome[j]:
                table[j] += table[j-1] if j>0 else 1
                table[j] %= 10000
    return '%04d' % table[len(welcome)-1]


for case in range(get_int()):
    print('Case #%d: %s' % (case + 1, solve()))

