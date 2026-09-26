package main

import (
	"slices"
	"strings"
)

// 5564 1976 (sample 41 6)
func main() {
	solve("06_input.txt", line06a, line06b, solve06a, solve06b)
}

var (
	grid06 = make([][]rune, 0)
	y06    = 0
	x06    = 0
	dirs06 = []struct {
		yinc, xinc int
	}{
		{-1, 0},
		{0, 1},
		{1, 0},
		{0, -1},
	}
)

func line06a(line string) {
	b := []rune(line)
	if index := strings.Index(line, "^"); index != -1 {
		y06 = len(grid06)
		x06 = index
	}
	grid06 = append(grid06, b)
}

func solve06a() int {
	grid := getGrid06()
	walk06(grid, true)
	total := 0
	for y := range grid {
		for x := range grid[y] {
			if grid[y][x] == 'X' {
				total++
			}
		}
	}
	return total
}

func line06b(line string) {
}

func solve06b() int {
	total := 0
	grid := getGrid06()
	for y := range grid06 {
		for x := range grid06[y] {
			if grid[y][x] != '#' {
				old := grid[y][x]
				grid[y][x] = '#'
				if walk06(grid, false) {
					total++
				}
				grid[y][x] = old
			}
		}
	}
	return total
}

func getGrid06() [][]rune {
	ret := make([][]rune, len(grid06))
	for i := range grid06 {
		ret[i] = make([]rune, len(grid06[i]))
		copy(ret[i], grid06[i])
	}
	return ret
}

func walk06(grid [][]rune, alterGrid bool) (isLoop bool) {
	cols := len(grid[0])
	visited := make(map[int][]int)
	if alterGrid {
		grid[y06][x06] = 'X'
	}
	dir := 0
	for y, x := y06, x06; ; {
		yfut, xfut := y+dirs06[dir].yinc, x+dirs06[dir].xinc
		if yfut < 0 || yfut >= len(grid) || xfut < 0 || xfut >= len(grid[yfut]) {
			break
		}
		if grid[yfut][xfut] == '#' {
			dir = (dir + 1) % 4
			continue
		}
		y, x = yfut, xfut
		if slices.Contains(visited[y*cols+x], dir) {
			return true
		}
		visited[y*cols+x] = append(visited[y*cols+x], dir)
		if alterGrid {
			grid[y][x] = 'X'
		}
	}
	return
}
