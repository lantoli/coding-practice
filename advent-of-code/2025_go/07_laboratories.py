#!/usr/bin/env python3


def solve07(lines: list[str], b: bool) -> int:
    count = [0] * len(lines[0])
    splits = 0

    for line in lines:
        for x, c in enumerate(line):
            if c == 'S':
                count[x] = 1
            elif c == '^':
                count[x-1] += count[x]
                count[x+1] += count[x]
                splits += count[x] > 0
                count[x] = 0

    return sum(count) if b else splits


if __name__ == "__main__":
    with open("inputs/07_sample.txt") as f:
        s = [l for line in f if (l := line.strip())]
    with open("inputs/07_input.txt") as f:
        i = [l for line in f if (l := line.strip())]

    print(f"Sample A: {solve07(s, False)}")  # 21
    print(f"A:        {solve07(i, False)}")  # 1678
    print(f"Sample B: {solve07(s, True)}")  # 40
    print(f"B:        {solve07(i, True)}")  # 357525737893560
