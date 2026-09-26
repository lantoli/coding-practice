#! /usr/bin/env python3
"""
Polynomial interpolation

Input is a number of points (x1,y1) ... (xn, yn) and print the result N-grade polynomial.

It uses Lagrange polynomials: http://en.wikipedia.org/wiki/Polynomial_interpolation
"""

points = [
    (0, 0),
    (1, 0),
    (2, 0),
    (3, 0),
    (4, 0),
    (5, 0),
    (6, 0),
    (7, 7),
    (8, 70),
]


import sympy

sympy.init_printing(use_unicode=True)

xsym = sympy.Symbol('x')

fsym = 0

for indx, (x, y) in enumerate(points):
    prod = y
    for iprod, (xprod, _) in enumerate(points):
        if iprod != indx:
            prod *= (xsym - xprod) / (x - xprod);
    fsym += prod

print(fsym)
print(sympy.expand(fsym))
print(sympy.simplify(fsym))

# Just to verify it's really aproximated by a polynomial
print(fsym.subs(xsym, 7))
print(fsym.subs(xsym, 8))
print(fsym.subs(xsym, 9))
print(fsym.subs(xsym, 10))
print(fsym.subs(xsym, 11))
print(fsym.subs(xsym, 12))
