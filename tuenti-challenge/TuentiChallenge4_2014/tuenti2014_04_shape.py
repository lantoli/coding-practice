#! /usr/bin/env python3
"""
Tuenti Challenge 4, Apr 2014, Challenge 4 - Shape shifters
"""

import sys


def solve(source,target, safe):
    cache = set()
    bfs = [[source]]  # breadth-first search tree
    while bfs:
        states = bfs.pop()
        current = states[-1]
        if one_change(current, target):
            states.append(target)
            return states
        cache.add(current)
        for other in safe - set(states) - cache:
            if one_change(current, other):
                new_states = list(states)
                new_states.append(other)
                bfs.append(new_states)


def one_change(stra,strb):
    return sum([a != b for a, b in zip(stra,strb)]) == 1


def main():
    source = input()
    target = input()
    safe = set()
    for line in sys.stdin:
        safe.add(line.rstrip('\n'))
    print('->'.join(solve(source, target, safe)))


if __name__ == '__main__':
    main()
