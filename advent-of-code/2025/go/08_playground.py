#!/usr/bin/env python3

from dataclasses import dataclass
from itertools import combinations


@dataclass
class Edge:
    pos1: int
    pos2: int
    sqdist: int


def find_root(parent: list[int], i: int) -> int:
    while i != parent[i]:
        i = parent[i]
    return i


def get_counts(parent: list[int]) -> list[int]:
    counts = [0] * len(parent)
    for i in range(len(parent)):
        counts[find_root(parent, i)] += 1
    return sorted(counts, reverse=True)


def all_together(parent: list[int]) -> bool:
    return get_counts(parent)[0] == len(parent)


def solve08(lines: list[str], connections: int, b: bool) -> int:
    coords = [tuple(map(int, line.split(','))) for line in lines]
    n = len(coords)

    edges = [
        Edge(i, j, sum((a - b) ** 2 for a, b in zip(coords[i], coords[j])))
        for i, j in combinations(range(n), 2)
    ]

    if b:
        connections = len(edges) + 1

    edges.sort(key=lambda e: e.sqdist)
    parent = list(range(n))

    for edge in edges[:connections]:
        root1 = find_root(parent, edge.pos1)
        root2 = find_root(parent, edge.pos2)
        if root1 != root2:
            parent[root2] = root1
            if b and all_together(parent):
                return coords[edge.pos1][0] * coords[edge.pos2][0]

    counts = get_counts(parent)
    return counts[0] * counts[1] * counts[2]


if __name__ == "__main__":
    with open("inputs/08_sample.txt") as f:
        s = [l for line in f if (l := line.strip())]
    with open("inputs/08_input.txt") as f:
        i = [l for line in f if (l := line.strip())]

    print(f"Sample A: {solve08(s, 10, False)}")  # 40
    print(f"A:        {solve08(i, 1000, False)}")  # 123930
    print(f"Sample B: {solve08(s, 10, True)}")  # 25272
    print(f"B:        {solve08(i, 1000, True)}")  # 27338688
