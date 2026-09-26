

import sympy

sympy.init_printing(use_unicode=True)

x = sympy.Symbol('x')
print(sympy.limit(sympy.sin(x)/x, x, 0))
print(sympy.integrate(1/x,x))


print(sympy.sqrt(4))
print(sympy.sqrt(4))

y = ((x-2)*(x-3))/((1-2)*(1-3))*1 + (((x-1)*(x-3))/((2-1)*(2-3)))*3 + (((x-1)*(x-2))/((3-1)*(3-2)))*6
print(y)
print(sympy.expand(y))
print(sympy.simplify(y))
print(y.subs(x,1))
print(y.subs(x,2))
print(y.subs(x,3))
print(y.subs(x,4))

