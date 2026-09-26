package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

// 1338 975376 (sample 44)
func main() {
	solve20("20_sample.txt", 2, 2)
	solve20("20_input.txt", 2, 100)
	solve20("20_input.txt", 20, 100)
}

func solve20(file string, cheats, saves int) {
	var (
		f, _       = os.Open("inputs/" + file)
		grid       []string
		rows       int
		start, end st20
		total      int
	)
	defer f.Close()
	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		line := scanner.Text()
		if pos := strings.Index(line, "S"); pos != -1 {
			start = st20{rows, pos}
		}
		if pos := strings.Index(line, "E"); pos != -1 {
			end = st20{rows, pos}
		}
		grid = append(grid, line)
		rows++
	}

	trackCount := make(map[st20]int)
	for pos, count := start, 0; ; count++ {
		trackCount[pos] = count
		if pos == end {
			break
		}
		for _, dir := range dirs20 {
			newpos := st20{pos[0] + dir[0], pos[1] + dir[1]}
			if _, found := trackCount[newpos]; !found && grid[newpos[0]][newpos[1]] != '#' {
				pos = newpos
				break
			}
		}
	}

	for pos, count := range trackCount {
		for yinc := -cheats; yinc <= cheats; yinc++ {
			for xinc := -cheats + abs(yinc); xinc <= cheats-abs(yinc); xinc++ {
				posNew := st20{pos[0] + yinc, pos[1] + xinc}
				countNew, found := trackCount[posNew]
				if found && countNew > count {
					inc := abs(yinc) + abs(xinc)
					if inc+saves <= countNew-count {
						total++
					}
				}
			}
		}
	}
	fmt.Println(total)
}

type st20 [2]int

var dirs20 = []st20{{-1, 0}, {1, 0}, {0, -1}, {0, 1}}
