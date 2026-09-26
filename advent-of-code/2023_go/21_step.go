package main

import (
	"bufio"
	"fmt"
	"os"
)

const filepath21 = "inputs/21_sample.txt"

func main() {
	fmt.Println(calc21(false)) // 3677
	fmt.Println(calc21(true))  // XXX
}

func calc21(part2 bool) int {
	f, _ := os.Open(filepath21)
	defer f.Close()
	scanner := bufio.NewScanner(f)
	grid := make([]string, 0)
	for scanner.Scan() {
		grid = append(grid, scanner.Text())
	}
	if len(grid) != len(grid[0]) {
		panic("grid is not square")
	}
	return calc21Main(grid, part2)
}

func calc21Main(grid []string, part2 bool) int {
	n := len(grid)
	var y, x int
outer:
	for y = 0; y < n; y++ {
		for x = 0; x < n; x++ {
			if grid[y][x] == 'S' {
				break outer
			}
		}
	}
	v := map[st21]any{st21{y, x}: nil}
	for i := 0; i < 5000; i++ {
		newv := make(map[st21]any)
		for k := range v {
			for _, inc := range [][2]int{{-1, 0}, {1, 0}, {0, 1}, {0, -1}} {
				y, x := k.y+inc[0], k.x+inc[1]
				if part2 {
					if grid[(y%n+n)%n][(x%n+n)%n] != '#' {
						newv[st21{y, x}] = nil
					}
				} else {
					if y >= 0 && y < n && x >= 0 && x < n && grid[y][x] != '#' {
						newv[st21{y, x}] = nil
					}
				}
			}
		}
		old := len(v)
		v = newv
		fmt.Println(i, len(v), old, len(v)-old)
	}
	return len(v)
}

type st21 struct {
	y, x int
}
