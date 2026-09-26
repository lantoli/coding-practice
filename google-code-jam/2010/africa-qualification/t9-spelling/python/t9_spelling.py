# https://code.google.com/codejam/contest/351101/dashboard#s=p2

FILENAME = "C-large-practice"

import sys
sys.stdin = open(FILENAME + ".in", 'r')
sys.stdout = open(FILENAME + ".out", 'w')

def get_line(): return input()
def get_int(): return int(get_line())


keyboard = {
    'a': '2', 'b': '22', 'c': '222',
    'd': '3', 'e': '33', 'f': '333',
    'g': '4', 'h': '44', 'i': '444',
    'j': '5', 'k': '55', 'l': '555',
    'm': '6', 'n': '66', 'o': '666',
    'p': '7', 'q': '77', 'r': '777', 's': '7777',
    't': '8', 'u': '88', 'v': '888',
    'w': '9', 'x': '99', 'y': '999', 'z': '9999',
    ' ': '0'
}


def solve():
    ret = ''
    for ch in get_line():
        out = keyboard[ch]
        if ret and ret[-1] == out[0]: ret += ' '
        ret += out
    return ret


for case in range(get_int()):
    print('Case #%d: %s' % (case+1, solve()))
