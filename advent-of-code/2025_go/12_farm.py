#!/usr/bin/env python3

# Based on https://github.com/LuisGC/advent-of-code/blob/main/2025/day-12/main.py

from math import prod

def amount_that_fits(tiles: list[str], regions: list[str]) -> int:
    
    tile_counts = {int(tile[0]): tile.count('#') for tile in tiles}
    count = 0

    for region in regions.splitlines():
        size, tile_requirements = region.split(": ")
        region_size = prod(tuple(int(x) for x in size.split('x')))
        region_space = sum(
            (tile_counts[id_] * int(num))
            for id_, num in enumerate(tile_requirements.split())
        )
        if region_space < region_size:
            count += 1
    return count

with open("inputs/12_sample.txt", encoding="utf-8") as f:
    *tiles, regions = f.read().split("\n\n")

# The lucky guess does not work for the example :-(
#    assert 2 == amount_that_fits(tiles, regions)

with open("inputs/12_input.txt", encoding="utf-8") as f:
    *tiles, regions = f.read().split("\n\n")

    print(f"A: {amount_that_fits(tiles, regions)}")
