package main

import (
	"bufio"
	"fmt"
	"os"
)

const filepath14 = "inputs/14_input.txt"

func main() {
	fmt.Println(calc14(false)) // 109665
	fmt.Println(calc14(true))  // 96061
}

func calc14(cycle bool) int {
	f, _ := os.Open(filepath14)
	defer f.Close()
	scanner := bufio.NewScanner(f)
	m := make([][]byte, 0)
	for scanner.Scan() {
		m = append(m, []byte(scanner.Text()))
	}
	calc14Move(m, cycle)
	return calc14Total(m)
}

func calc14Move(m [][]byte, cycle bool) {
	movs := [][]int{{-1, 0}}
	count := 1
	if cycle {
		movs = [][]int{{-1, 0}, {0, -1}, {1, 0}, {0, 1}}
		count = 1_000_000_000
	}
	cache := make(map[string]int)
	rows, cols := len(m), len(m[0])
	for i := 0; i < count; i++ {
		for _, mov := range movs {
			yini, xini, yinc, xinc := 0, 0, 1, 1
			if mov[0] > 0 {
				yini, yinc = rows-1, -1
			}
			if mov[1] > 0 {
				xini, xinc = cols-1, -1
			}
			for y := yini; y >= 0 && y < rows; y += yinc {
				for x := xini; x >= 0 && x < cols; x += xinc {
					if m[y][x] == 'O' {
						for mul := 1; ; mul++ {
							yy := y + mov[0]*mul
							xx := x + mov[1]*mul
							if yy < 0 || yy >= rows || xx < 0 || xx >= cols || m[yy][xx] != '.' {
								break
							}
							m[yy][xx] = 'O'
							m[yy-mov[0]][xx-mov[1]] = '.'
						}
					}
				}
			}
		}
		val := cache[serialize14(m)]
		if val != 0 {
			mod := i - val
			for ; i+mod < count; i += mod {
			}
		} else {
			cache[serialize14(m)] = i
		}
	}
}

func calc14Total(m [][]byte) int {
	rows, cols := len(m), len(m[0])
	total := 0
	for y := 0; y < rows; y++ {
		for x := 0; x < cols; x++ {
			if m[y][x] == 'O' {
				total += rows - y
			}
		}
	}
	return total
}

func serialize14(m [][]byte) string {
	rows, cols := len(m), len(m[0])
	s := make([]byte, 0, rows*cols)
	for y := 0; y < rows; y++ {
		s = append(s, m[y]...)
	}
	return string(s)
}
