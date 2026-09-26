#! /usr/bin/env python3
"""
Tuenti Challenge 4, Apr 2014, Challenge 4 - Shape shifters
"""

from sys import stdin
from itertools import product
import networkx as nx


def solve(source,target, safe):
    G = nx.Graph()
    G.add_edges_from(((x,y) for x,y in product(safe, safe) if one_change(x,y)))
    return nx.shortest_path(G, source, target)


def one_change(stra,strb):
    return sum([a != b for a, b in zip(stra,strb)]) == 1


def main():
    source = input()
    target = input()
    safe = set([source, target])
    for line in stdin:
        safe.add(line.rstrip('\n'))
    print('->'.join(solve(source, target, safe)))


if __name__ == '__main__':
    main()
