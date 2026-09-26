# https://code.google.com/codejam/contest/90101/dashboard#s=p0

# another easier way to transform it into a re

FILENAME = "A-large-practice"

import sys
sys.stdin = open(FILENAME + ".in", 'r')
sys.stdout = open(FILENAME + ".out", 'w')

def get_line(): return input()
def get_int(): return int(get_line())
def get_ints(): return [int(x) for x in get_line().split()]

def solve_one(pattern, d):
    if len(pattern) == 0: return 1
    letter = pattern[0]
    if letter == '(':
        end = pattern.find(')')
        sum = 0
        for pos in range(1,end):
            letter = pattern[pos]
            if letter in d:
                sum += solve_one(pattern[end+1:], d[letter])
        return sum
    else:
        if letter in d:
            return solve_one(pattern[1:], d[letter])
        else:
            return 0
    return pattern


def solve(d):
    return solve_one(get_line(), d)


d = {}
tokens, word_count, cases = get_ints()
for word in [get_line() for x in range(word_count)]:
    dd = d
    for letter in word:
        if letter not in dd: dd[letter] = {}
        dd = dd[letter]
for case in range(cases):
    print('Case #%d: %s' % (case+1, solve(d)))


