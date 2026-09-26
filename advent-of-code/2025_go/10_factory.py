#!/usr/bin/env python3

# Based on https://github.com/LuisGC/advent-of-code/blob/main/2025/day-10/main.py

from collections import deque

import z3


def parse_line(line: str) -> tuple[list[bool], list[list[int]], list[int]]:
    """Parse a line into lights, buttons, and joltages."""
    parts = line.split(" ")

    # Parse lights (target state)
    lights = [c == "#" for c in parts[0][1:-1]]

    # Parse joltages
    joltages = [int(x) for x in parts[-1][1:-1].split(",")]

    # Parse buttons
    buttons = [
        [int(x) for x in part[1:-1].split(",")]
        for part in parts[1:-1]
    ]

    return lights, buttons, joltages


def solve10a(lights: list[bool], buttons: list[list[int]]) -> int:
    """Find minimum button presses to turn all lights on (BFS)."""
    initial = tuple([False] * len(lights))
    target = tuple(lights)
    seen = {initial}
    queue = deque([(initial, 0)])

    while queue:
        state, presses = queue.popleft()

        if state == target:
            return presses

        for button in buttons:
            new_state = list(state)
            for idx in button:
                if 0 <= idx < len(new_state):
                    new_state[idx] = not new_state[idx]

            new_tuple = tuple(new_state)
            if new_tuple not in seen:
                seen.add(new_tuple)
                queue.append((new_tuple, presses + 1))

    return -1


def solve10b(joltages: list[int], buttons: list[list[int]]) -> int:
    """Find minimum button presses to achieve joltages (ILP with Z3)."""
    n = len(buttons)
    m = len(joltages)

    presses = [z3.Int(f"x{i}") for i in range(n)]
    opt = z3.Optimize()

    # Non-negative constraints
    for p in presses:
        opt.add(p >= 0)

    # Joltage constraints: sum of presses affecting position i = joltages[i]
    for i in range(m):
        affecting = [presses[j] for j, btn in enumerate(buttons) if i in btn]
        if affecting:
            opt.add(z3.Sum(affecting) == joltages[i])
        elif joltages[i] != 0:
            return -1

    # Minimize total presses
    opt.minimize(z3.Sum(presses))
    opt.check()
    model = opt.model()

    return sum(model[p].as_long() for p in presses)


def solve10(lines: list[str], b: bool) -> int:
    """Solve Day 10: Part A (lights) or Part B (joltages)."""
    total = 0

    for line in lines:
        lights, buttons, joltages = parse_line(line)
        if b:
            total += solve10b(joltages, buttons)
        else:
            total += solve10a(lights, buttons)

    return total


if __name__ == "__main__":
    with open("inputs/10_sample.txt") as f:
        s = [l for line in f if (l := line.strip())]
    with open("inputs/10_input.txt") as f:
        i = [l for line in f if (l := line.strip())]

    print(f"Sample A: {solve10(s, False)}")  # 7
    print(f"A:        {solve10(i, False)}")  # 524
    print(f"Sample B: {solve10(s, True)}")   # 33
    print(f"B:        {solve10(i, True)}")   # 21696
