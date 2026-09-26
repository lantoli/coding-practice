# https://code.google.com/codejam/contest/3004486/dashboard#s=p1

# ONLY WORKS FOR SMALL SETS


FILENAME = "B-small-practice"

import operator
import sys
sys.stdin = open(FILENAME + ".in", 'r')
sys.stdout = open(FILENAME + ".out", 'w')

from itertools import groupby
from functools import reduce
from math import factorial

def get_line(): return input()
def get_int(): return int(get_line())
def get_ints(): return [int(x) for x in get_line().split()]
def get_strs(): return [x for x in get_line().split()]


def prod(iterable): return reduce(operator.mul, iterable, 1)

MOD = 1000000007


def recur(str, cars):
    cars.discard(str[-1]+str[-1])
    if len(cars) == 0:
        return 1

    candidates = [car for car in cars if car[0] == str[-1] and car[1] not in str]
    if len(candidates) > 1:
        return 0
    elif len(candidates) == 1:
        car = candidates.pop()
        sublist = set(cars)
        sublist.remove(car)
        return recur(str + car, sublist)
    else:
        total = 0
        for car in cars:
            if car[0] not in str and car[1] not in str:
                sublist = set(cars)
                sublist.remove(car)
                total += recur(str + car, sublist)
        return total


def good_car(car):
    for idx1 in range(len(car)-1):
        for idx2 in range(idx1+1, len(car)):
            if car[idx1] == car[idx2]:
                for id in range(idx1+1, idx2-2):
                    if car[idx1] != car[id]:
                        return False
    return True


def solve():
    get_line()
    cars = [x for x in get_line().split()]

    for idx, original_car in enumerate(cars):
        car = original_car
        if not good_car(car):
            return 0
        while len(car) > 2 and car[0] == car[1]:
            car = car[1:]
        while len(car) > 2 and car[-1] == car[-2]:
            car = car[:-1]
        letters = set([ch for mycar in cars for ch in mycar if mycar != original_car])
        for i in range(1, len(car)-1):
            if car[i] in letters:
                return 0    # middle letter can't match another car
        cars[idx] = str(car[0]) + str(car[-1])
    cars = sorted(cars)

    if [True for key, group in groupby(cars) if key[0] != key[-1] and len(list(group)) > 1]:
        return 0

    mult = prod([factorial(len(list(group))) % MOD for key, group in groupby(cars)]) % MOD

    return (mult * recur("A", set(cars))) % MOD

for case in range(get_int()):
    print('Case #%d: %s' % (case+1, solve()))

