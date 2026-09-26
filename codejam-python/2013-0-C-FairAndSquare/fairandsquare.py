# https://code.google.com/codejam/contest/2270488/dashboard#s=p2

FILENAME = "C-large-practice-1"

import sys
sys.stdin = open(FILENAME + ".in", 'r')
sys.stdout = open(FILENAME + ".out", 'w')

def get_line(): return input()
def get_int(): return int(get_line())
def get_ints(): return [int(x) for x in get_line().split()]


def good(number):
    n = str(number)
    l = len(n)//2
    part1 = n[l-1::-1]
    part2 = n[l+1:]
    return part1 == part2



def next_comb(a):
    while a.count(2) != len(a):
        pos = len(a) - 1
        while pos >= 0:
            a[pos] += 1
            if a[pos] == 3:
                a[pos] = 0
                pos -= 1
            else:
                break
        return True
    return False


numbers = [1, 4, 9]
for digits in range(1,5):

    a = [0] + [2] * (digits-1)
    while next_comb(a):
        part = ''.join(map(str,a))
        number = int(part + part[::-1])
        nnumber = number*number
        if len(str(nnumber)) % 2 == 1 and good(nnumber):
            numbers.append(nnumber)

    a = [0] + [2] * (digits-1)
    while next_comb(a):
        part = ''.join(map(str,a))
        for middle in range(0,10):
            number = int(str(part) + str(middle) +  str(part)[::-1])
            nnumber = number*number
            if len(str(nnumber)) % 2 == 1 and good(nnumber):
                numbers.append(nnumber)

import bisect
def solve():
    a,b = get_ints()


    return bisect.bisect_right(numbers, b) - bisect.bisect_left(numbers, a)


for case in range(get_int()):
    print('Case #%d: %s' % (case+1, solve()))


