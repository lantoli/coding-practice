#! /usr/bin/env python3
"""
Tuenti Challenge 4, Apr 2014, Challenge 20 - forfor

Challenge not submitted to the contest.


Got the idea from ajimenezh that it can be interpolated by a polynomial:
https://github.com/ajimenezh/Programing-Contests/tree/master/Tuenti%20Challenge/2014/Problema%2020

Thanks also to @alvarogonzalez for insigts to get to this solution.

It can represented by any of these polynomials using Lagrange polynomial:
https://github.com/lantoli/codejam-python/blob/master/Examples/sympy/polinomial_interpolation.py

x*(-x + 8)*(x/6 - 1/6)*(x/5 - 2/5)*(x/4 - 3/4)*(x/3 - 4/3)*(x/2 - 5/2)*(x - 6) + 35*x*(x/7 - 1/7)*(x/6 - 1/3)*(x/5 - 3/5)*(x/4 - 1)*(x/3 - 5/3)*(x/2 - 3)*(x - 7)/4
x**8/2880 - x**7/120 + 119*x**6/1440 - 7*x**5/16 + 3829*x**4/2880 - 553*x**3/240 + 167*x**2/80 - 3*x/4
x*(x - 6)*(x - 5)*(x - 4)*(x - 3)**2*(x - 2)*(x - 1)/2880

"""

import sys

MOD = 3211123

def f(x): return (x - 6)*(x - 5)*(x - 4)*(x - 3)*(x - 3)*(x - 2)*(x - 1)*x // 2880

if __name__ == '__main__':
    res = 0
    for line in sys.stdin:
        num = int(line.rstrip('\n'))
        res = (res + f(num % MOD)) % MOD
    print(res)
