#!/usr/bin/env python3


def solve05a(ranges: list[tuple[int, int]], items: list[int]) -> int:
    return sum(1 for item in items
               if any(start <= item <= end for start, end in ranges))


def solve05b(ranges: list[tuple[int, int]]) -> int:
    ranges.sort()
    merged = []
    for start, end in ranges:
        if not merged or merged[-1][1] < start:
            merged.append([start, end])
        elif end > merged[-1][1]:
            merged[-1][1] = end
    return sum(end - start + 1 for start, end in merged)


def parse(filename):
    with open(filename) as f:
        ranges_part, items_part = f.read().split("\n\n")
    ranges = [tuple(map(int, line.split("-"))) for line in ranges_part.strip().split("\n")]
    items = [int(line) for line in items_part.strip().split("\n") if line]
    return ranges, items


if __name__ == "__main__":
    s_ranges, s_items = parse("inputs/05_sample.txt")
    i_ranges, i_items = parse("inputs/05_input.txt")

    print(f"Sample A: {solve05a(s_ranges, s_items)}")  # 3
    print(f"A:        {solve05a(i_ranges, i_items)}")  # 821
    print(f"Sample B: {solve05b(s_ranges)}")  # 14
    print(f"B:        {solve05b(i_ranges)}")  # 344771884978261
