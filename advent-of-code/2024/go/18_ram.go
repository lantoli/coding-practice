package main

import (
	"fmt"
	"math"
	"strconv"
	"strings"
)

// 268 64,11  (sample 22 6,1)
func main() {
	file, size, mem := "18_sample.txt", 7, 12
	if isInput18 {
		file, size, mem = "18_input.txt", 71, 1024
	}
	solve(file, line18, nil, func() int { return solve18a(size, mem) }, func() int { return solve18b(size) })
}

var (
	isInput18 = true
	input18   = make([][]int, 0)
)

func line18(line string) {
	nums := strings.Split(line, ",")
	x, _ := strconv.Atoi(nums[0])
	y, _ := strconv.Atoi(nums[1])
	input18 = append(input18, []int{x, y})
}

func solve18a(size, mem int) int {
	grid := make([]bool, size*size)
	for pos := range mem {
		x, y := input18[pos][0], input18[pos][1]
		grid[y*size+x] = true
	}
	return calc18(size, grid)
}

func solve18b(size int) int {
	grid := make([]bool, size*size)
	for pos := range input18 {
		x, y := input18[pos][0], input18[pos][1]
		grid[y*size+x] = true
		if calc18(size, grid) == math.MaxInt {
			fmt.Printf("%d,%d\n", x, y)
			return 0
		}
	}
	panic("no solution")
}

func calc18(size int, grid []bool) int {
	incs := [][]int{{0, -1}, {0, 1}, {-1, 0}, {1, 0}}
	minsteps := math.MaxInt
	steps := make(map[int]int)
	list := []int{0}
	steps[0] = 0
	for len(list) > 0 {
		pos := list[0]
		list = list[1:]
		y, x := pos/size, pos%size
		step, ok := steps[pos]
		if !ok {
			panic("no step")
		}
		for _, inc := range incs {
			ynext, xnext := y+inc[0], x+inc[1]
			posnext, stepnext := ynext*size+xnext, step+1
			if ynext >= 0 && ynext < size && xnext >= 0 && xnext < size && !grid[posnext] {
				if _, found := steps[posnext]; !found || stepnext < steps[posnext] {
					steps[posnext] = stepnext
					list = append(list, posnext)
				}
				if ynext == size-1 && xnext == size-1 {
					minsteps = min(minsteps, stepnext)
				}
			}
		}
	}
	return minsteps
}
