#!/usr/bin/env python3

from itertools import combinations
from typing import NamedTuple


class Seg(NamedTuple):
    fixed: int
    lo: int
    hi: int


def build_segments(coords: list[tuple[int, int]]) -> tuple[list[Seg], list[Seg]]:
    h_segs, v_segs = [], []
    for (x1, y1), (x2, y2) in zip(coords, coords[1:] + coords[:1]):
        if y1 == y2:
            h_segs.append(Seg(y1, min(x1, x2), max(x1, x2)))
        else:
            v_segs.append(Seg(x1, min(y1, y2), max(y1, y2)))
    return h_segs, v_segs


def is_rect_inside(x1: int, y1: int, x2: int, y2: int, h_segs: list[Seg], v_segs: list[Seg]) -> bool:
    if any(y1 < h.fixed < y2 and h.lo < x2 and x1 < h.hi for h in h_segs):
        return False
    if any(x1 < v.fixed < x2 and v.lo < y2 and y1 < v.hi for v in v_segs):
        return False
    return True


def solve09(lines: list[str], b: bool) -> int:
    coords = [tuple(map(int, line.split(','))) for line in lines]
    h_segs, v_segs = build_segments(coords)

    rects = [
        (min(c1[0], c2[0]), min(c1[1], c2[1]), max(c1[0], c2[0]), max(c1[1], c2[1]))
        for c1, c2 in combinations(coords, 2)
    ]
    rects.sort(key=lambda r: (r[2] - r[0] + 1) * (r[3] - r[1] + 1), reverse=True)

    for x1, y1, x2, y2 in rects:
        if not b or is_rect_inside(x1, y1, x2, y2, h_segs, v_segs):
            return (x2 - x1 + 1) * (y2 - y1 + 1)

    raise ValueError("no solution found")


if __name__ == "__main__":
    with open("inputs/09_sample.txt") as f:
        s = [l for line in f if (l := line.strip())]
    with open("inputs/09_input.txt") as f:
        i = [l for line in f if (l := line.strip())]

    print(f"Sample A: {solve09(s, False)}")  # 50
    print(f"A:        {solve09(i, False)}")  # 4776100539
    print(f"Sample B: {solve09(s, True)}")  # 24
    print(f"B:        {solve09(i, True)}")  # 1476550548
