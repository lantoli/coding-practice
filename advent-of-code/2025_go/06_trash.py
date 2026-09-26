#!/usr/bin/env python3

import math


def solve06a(lines: list[str]) -> int:
    nums = [list(map(int, line.split())) for line in lines[:-1]]
    ops = lines[-1].split()
    return sum(
        math.prod(sub) if op == '*'
        else sum(sub)
        for x, op in enumerate(ops)
        for sub in [[nums[y][x] for y in range(len(nums))]]
    )


def solve06b(lines: list[str]) -> int:
    total, nums, x = 0, [], len(lines[0]) - 1

    while x >= 0:
        nums.append(int(''.join(c for y in range(len(lines) - 1)
                                if (c := lines[y][x]).isdigit()) or '0'))
        op = lines[-1][x]
        if op in ['+', '*']:
            x -= 1
            total += sum(nums) if op == '+' else math.prod(nums)
            nums = []
        x -= 1
    return total


if __name__ == "__main__":
    with open("inputs/06_sample.txt") as f:
        s = [line for line in f if line.strip()]
    with open("inputs/06_input.txt") as f:
        i = [line for line in f if line.strip()]

    print(f"Sample A: {solve06a(s)}")  # 4277556
    print(f"A:        {solve06a(i)}")  # 5524274308182
    print(f"Sample B: {solve06b(s)}")  # 3263827
    print(f"B:        {solve06b(i)}")  # 8843673199391
