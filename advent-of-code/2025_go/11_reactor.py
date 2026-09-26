#!/usr/bin/env python3

from functools import lru_cache


def solve11(lines: list[str], b: bool) -> int:
    nodes: dict[str, list[str]] = {}
    for line in lines:
        parts = line.split()
        key = parts[0][:-1]
        nodes[key] = parts[1:]

    @lru_cache(maxsize=None)
    def dfs(node: str, has_fft: bool, has_dac: bool) -> int:
        has_fft |= node == "fft"
        has_dac |= node == "dac"
        total = 0
        for next_node in nodes.get(node, []):
            if next_node == "out":
                if not b or (has_fft and has_dac):
                    total += 1
            else:
                total += dfs(next_node, has_fft, has_dac)
        return total
    start = "svr" if b else "you"
    return dfs(start, False, False)


if __name__ == "__main__":
    with open("inputs/11_sample_a.txt") as f:
        sa = [l for line in f if (l := line.strip())]
    with open("inputs/11_sample_b.txt") as f:
        sb = [l for line in f if (l := line.strip())]
    with open("inputs/11_input.txt") as f:
        i = [l for line in f if (l := line.strip())]

    print(f"Sample A: {solve11(sa, False)}")  # 5
    print(f"A:        {solve11(i, False)}")   # 599
    print(f"Sample B: {solve11(sb, True)}")   # 2
    print(f"B:        {solve11(i, True)}")    # 393474305030400
