#! /usr/bin/env python
"""
Tuenti Challenge 4, Apr 2014, Challenge 9 - Bendito Caos

This problem is an instance of The Maximum Flow Problem.
We'll use Google OR-tools library in order to not reiventing the wheel:
https://code.google.com/p/or-tools/
"""

from ortools.graph import pywrapgraph


target = "AwesomeVille"
CARS = 200  # flow of cars in 1 km/h (1000 m / 5 m)


def solve():
    source = raw_input()
    Sspeed, Dspeed = get_ints()
    Icount, Rcount = get_ints()
    max_flow = pywrapgraph.SimpleMaxFlow()
    for _ in range(Rcount):
        from_node, to_node, type, lanes = get_strs()
        # don't need roads arriving source or leaving target
        if to_node != source and from_node != target:
            total_speed = int(lanes) * (Sspeed if type == "normal" else Dspeed)
            vfrom = 0 if from_node == source else int(from_node)+1
            vto = Rcount+1 if to_node == target else int(to_node)+1
            max_flow.AddArcWithCapacity(vfrom, vto, total_speed)
    max_flow.Solve(0, Rcount+1)
    return "{} {}".format(source, max_flow.OptimalFlow() * CARS)


def get_strs(): return [x for x in raw_input().split()]
def get_ints(): return map(int, get_strs())


if __name__ == '__main__':
    for case in range(int(raw_input())):
        print(solve())
