package main

import (
	"bufio"
	"fmt"
	"os"
)

const filepath16 = "inputs/16_input.txt"

func main() {
	fmt.Println(calc16(false)) // 7477
	fmt.Println(calc16(true))  // 7853
}

func calc16(all bool) int {
	f, _ := os.Open(filepath16)
	defer f.Close()
	scanner := bufio.NewScanner(f)
	grid := make([]string, 0)
	for scanner.Scan() {
		grid = append(grid, scanner.Text())
	}
	if len(grid) != len(grid[0]) {
		panic("grid is not square")
	}
	n := len(grid)
	ret := calc16Main(grid, st16{0, 0, EAST, n})
	if all {
		for i := range grid {
			ret = max(ret, calc16Main(grid, st16{i, 0, EAST, n}))
			ret = max(ret, calc16Main(grid, st16{i, n - 1, WEST, n}))
			ret = max(ret, calc16Main(grid, st16{0, i, SOUTH, n}))
			ret = max(ret, calc16Main(grid, st16{n - 1, i, NORTH, n}))
		}
	}
	return ret
}

func calc16Main(grid []string, ini st16) int {
	n := len(grid)
	seen := make([][]int, n)
	for i := range seen {
		seen[i] = make([]int, n)
	}
	states := append([]st16{}, ini)
	for len(states) > 0 {
		newStates := make([]st16, 0)
		for _, state := range states {
			if seen[state.y][state.x]&state.dir > 0 {
				continue
			}
			seen[state.y][state.x] |= state.dir
			switch grid[state.y][state.x] {
			case '\\':
				state.dir = map[int]int{NORTH: WEST, EAST: SOUTH, SOUTH: EAST, WEST: NORTH}[state.dir]
			case '/':
				state.dir = map[int]int{NORTH: EAST, EAST: NORTH, SOUTH: WEST, WEST: SOUTH}[state.dir]
			case '|':
				if state.dir == EAST || state.dir == WEST {
					copy := st16{state.y, state.x, NORTH, n}
					if copy.move() {
						newStates = append(newStates, copy)
					}
					state.dir = SOUTH
				}
			case '-':
				if state.dir == NORTH || state.dir == SOUTH {
					copy := st16{state.y, state.x, EAST, n}
					if copy.move() {
						newStates = append(newStates, copy)
					}
					state.dir = WEST
				}
			}
			if state.move() {
				newStates = append(newStates, state)
			}
		}
		states = newStates
	}
	ret := 0
	for y := range seen {
		for x := range seen[y] {
			if seen[y][x] > 0 {
				ret++
			}
		}
	}
	return ret
}

type st16 struct {
	y   int
	x   int
	dir int
	n   int
}

func (s *st16) move() bool {
	switch s.dir {
	case NORTH:
		s.y--
	case EAST:
		s.x++
	case SOUTH:
		s.y++
	case WEST:
		s.x--
	}
	return s.y >= 0 && s.y < s.n && s.x >= 0 && s.x < s.n
}

const (
	NORTH = 1
	EAST  = 2
	SOUTH = 4
	WEST  = 8
)
