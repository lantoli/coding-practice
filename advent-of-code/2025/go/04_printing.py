#!/usr/bin/env python3

NEIGHBORS = [(dy, dx) for dy in range(-1, 2) for dx in range(-1, 2) if (dy, dx) != (0, 0)]


def solve04(grid: list[list[bool]], part_b: bool) -> int:
    def count_neighbors(y: int, x: int) -> int:
        return sum(
            1 for dy, dx in NEIGHBORS
            if 0 <= (ny := y + dy) < len(grid)
            and 0 <= (nx := x + dx) < len(grid[ny])
            and grid[ny][nx]
        )

    total = 0
    changed = True
    while changed:
        changed = False
        for y, row in enumerate(grid):
            for x, is_active in enumerate(row):
                if is_active and count_neighbors(y, x) < 4:
                    total += 1
                    if part_b:
                        grid[y][x] = False
                        changed = True
    return total


if __name__ == "__main__":
    with open("inputs/04_sample.txt") as f:
        s = [[c == '@' for c in line.strip()] for line in f]
    with open("inputs/04_input.txt") as f:
        i = [[c == '@' for c in line.strip()] for line in f]

    print(f"Sample A: {solve04(s, False)}") # 13
    print(f"A:        {solve04(i, False)}") # 1489
    print(f"Sample B: {solve04(s, True)}")  # 43
    print(f"B:        {solve04(i, True)}")  # 8890
