#! /usr/bin/env python3
"""
Tuenti Challenge 4, Apr 2014, Challenge 3 - The Gambler’s Club - Monkey Island 2
"""

from math import sqrt


def main():
    for case in range(get_int()):
        x, y = get_ints()
        res = sqrt(x*x + y*y)
        print(round(res,2))


def get_ints(): return [int(x) for x in input().split()]
def get_int(): return get_ints()[0]

if __name__ == '__main__':
    main()
