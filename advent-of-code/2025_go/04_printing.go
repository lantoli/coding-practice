package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	fmt.Println("Sample A", solve04("inputs/04_sample.txt", false)) // 13
	fmt.Println("A", solve04("inputs/04_input.txt", false))         // 1489
	fmt.Println("Sample B", solve04("inputs/04_sample.txt", true))  // 43
	fmt.Println("B", solve04("inputs/04_input.txt", true))          // 8890
}

func solve04(filename string, b bool) int {
	var (
		f, _    = os.Open(filename)
		scanner = bufio.NewScanner(f)
		grid    [][]bool
		total   = 0
		changed = true
	)
	defer f.Close()
	for scanner.Scan() {
		line := scanner.Text()
		if line == "" {
			continue
		}
		grid = append(grid, make([]bool, len(line)))
		for i, c := range line {
			grid[len(grid)-1][i] = c == '@'
		}
	}
	for changed {
		changed = false
		for y := range grid {
			for x := range grid[y] {
				if !grid[y][x] {
					continue
				}
				neighbors := 0
				for yinc := -1; yinc <= 1; yinc++ {
					for xinc := -1; xinc <= 1; xinc++ {
						if (yinc != 0 || xinc != 0) && y+yinc >= 0 && y+yinc < len(grid) && x+xinc >= 0 && x+xinc < len(grid[y+yinc]) && grid[y+yinc][x+xinc] {
							neighbors++
						}
					}
				}
				if neighbors < 4 {
					total++
					if b {
						grid[y][x] = false
						changed = true
					}
				}
			}
		}
	}
	return total
}

func solve04Line(line string) int {
	return 0
}
